#!/bin/bash
#
# Copyright 2018 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

set -ex

source hack/common.sh
source cluster/$NATIVELB_PROVIDER/provider.sh
source hack/config.sh
source ./cluster/gocli.sh

registry_port=$($gocli --prefix $provider_prefix ports registry | tr -d '\r')
registry=localhost:$registry_port

REGISTRY=$registry make docker-build
REGISTRY=$registry make docker-push

./cluster/kubectl.sh delete --ignore-not-found -f ./config/crds/
./cluster/kubectl.sh delete --ignore-not-found -f ./config/rbac/
./cluster/kubectl.sh delete --ignore-not-found ns nativelb

# Wait until all objects are deleted
until [[ `./cluster/kubectl.sh get ns | grep nativelb | wc -l` -eq 0 ]]; do
    sleep 5
done

./cluster/kubectl.sh apply -f config/crds/
./cluster/kubectl.sh apply -f config/rbac/
./cluster/kubectl.sh create ns nativelb
./cluster/kubectl.sh apply -f config/develop/
#./cluster/kubectl.sh apply -f config/ci/
