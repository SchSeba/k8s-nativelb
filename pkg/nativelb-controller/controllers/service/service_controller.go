/*
Copyright 2018 Sebastian Sch.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service_controller

import (
	"context"
	"github.com/k8s-nativelb/pkg/kubecli"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"time"

	"github.com/k8s-nativelb/pkg/apis/nativelb/v1"
	"github.com/k8s-nativelb/pkg/log"
	"github.com/k8s-nativelb/pkg/nativelb-controller/controllers/farm"
	"k8s.io/client-go/kubernetes"
)

type ServiceController struct {
	Controller       controller.Controller
	ReconcileService *ReconcileService
}

func (s *ServiceController) UpdateAllServices() {
	services := &corev1.ServiceList{}
	err := s.ReconcileService.GetClient().List(context.Background(), &client.ListOptions{}, services)
	if err != nil {
		log.Log.Errorf("Fail to get all services error: %v", err)
	}

	for _, service := range services.Items {
		s.ReconcileService.Reconcile(reconcile.Request{NamespacedName: types.NamespacedName{Namespace: service.Namespace, Name: service.Name}})
	}
}

func NewServiceController(nativelbClient kubecli.NativelbClient, farmController *farm_controller.FarmController) (*ServiceController, error) {
	reconcileService := newReconciler(nativelbClient, farmController)

	controllerInstance, err := newController(nativelbClient, reconcileService)
	if err != nil {
		return nil, err
	}
	serviceController := &ServiceController{Controller: controllerInstance,
		ReconcileService: reconcileService}

	go reconcileService.reSyncProcess()

	return serviceController, nil

}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(nativelbClient kubecli.NativelbClient, farmController *farm_controller.FarmController) *ReconcileService {
	return &ReconcileService{NativelbClient: nativelbClient,
		scheme:         nativelbClient.GetScheme(),
		Event:          nativelbClient.GetRecorder(v1.EventRecorderName),
		FarmController: farmController,
		kubeClient:     nativelbClient.GetKubeClient()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func newController(nativelbClient kubecli.NativelbClient, r reconcile.Reconciler) (controller.Controller, error) {
	// Create a new controller
	c, err := controller.New("service-controller", nativelbClient.GetManager(), controller.Options{Reconciler: r})
	if err != nil {
		return nil, err
	}

	// Watch for changes to service
	err = c.Watch(&source.Kind{Type: &corev1.Service{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return nil, err
	}

	return c, nil
}

var _ reconcile.Reconciler = &ReconcileService{}

// ReconcileService reconciles a Service object
type ReconcileService struct {
	kubecli.NativelbClient
	Event          record.EventRecorder
	FarmController *farm_controller.FarmController
	scheme         *runtime.Scheme
	kubeClient     *kubernetes.Clientset
}

// Reconcile reads that state of the cluster for a Service object and makes changes based on the state read
// and what is in the Service.Spec
// +kubebuilder:rbac:groups=core,resources=services,verbs=create;get;list;watch;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;update;delete;patch
func (r *ReconcileService) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Fetch the Service instance
	service := &corev1.Service{}
	err := r.GetClient().Get(context.TODO(), request.NamespacedName, service)
	if err != nil {
		if errors.IsNotFound(err) {
			r.FarmController.DeleteFarm(request.Namespace, request.Name)
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	if service.Spec.Type != "LoadBalancer" || len(service.Finalizers) != 0 {
		return reconcile.Result{}, nil
	}

	log.Log.V(2).Infof("Service event, service name: %s from namespace %s", service.Name, service.Namespace)
	if r.FarmController.CreateOrUpdateFarm(service, nil) {
		_, err := r.kubeClient.CoreV1().Services(service.Namespace).UpdateStatus(service)
		if err != nil {
			log.Log.Errorf("Fail to update service status error message: %s", err.Error())
		}
	}
	return reconcile.Result{}, nil
}

func (r *ReconcileService) UpdateEndpoints(service *corev1.Service, endpoint *corev1.Endpoints) {
	if r.FarmController.CreateOrUpdateFarm(service, endpoint) {
		_, err := r.kubeClient.CoreV1().Services(service.Namespace).UpdateStatus(service)
		if err != nil {
			log.Log.Errorf("Fail to update service status error message: %s", err.Error())
		}
	}
}

func (r *ReconcileService) getServiceFromEndpoint(endpointInstance *corev1.Endpoints) (*corev1.Service, error) {
	return r.kubeClient.CoreV1().Services(endpointInstance.Namespace).Get(endpointInstance.Name, metav1.GetOptions{})
}

func (r *ReconcileService) reSyncProcess() {
	resyncTick := time.Tick(30 * time.Second)

	labelSelector := labels.Set{}
	labelSelector[v1.ServiceStatusLabel] = v1.ServiceStatusLabelFailed

	for range resyncTick {
		var serviceList corev1.ServiceList
		err := r.GetClient().List(context.TODO(), &client.ListOptions{LabelSelector: labelSelector.AsSelector()}, &serviceList)
		if err != nil {
			log.Log.Error("reSyncProcess: Fail to get Service list")
		} else {
			for _, service := range serviceList.Items {
				r.Reconcile(reconcile.Request{NamespacedName: types.NamespacedName{Namespace: service.Namespace, Name: service.Name}})
			}
		}
	}
}
