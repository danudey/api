#!/bin/bash

# Copyright 2015 The Kubernetes Authors.
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

# The only argument this script should ever be called with is '--verify-only'

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT=$(realpath $(dirname "${BASH_SOURCE}")/..)
BINDIR=${REPO_ROOT}/bin

# Generate the versioned clientset (pkg/client/clientset_generated/clientset)
client-gen "$@" \
		--go-header-file "${REPO_ROOT}/hack/boilerplate/boilerplate.go.txt" \
		--plural-exceptions "UISettings:UISettings" \
		--input-base "github.com/tigera/api/pkg/apis/" \
		--input "projectcalico/v3" \
		--output-dir "${REPO_ROOT}/pkg/client/clientset_generated" \
		--clientset-path "github.com/tigera/api/pkg/client/clientset_generated/" \
		--clientset-name "clientset"
# generate lister
lister-gen "$@" \
		--go-header-file "${REPO_ROOT}/hack/boilerplate/boilerplate.go.txt" \
		--output-dir "${REPO_ROOT}/pkg/client/listers_generated" \
		--output-pkg "github.com/tigera/api/pkg/client/listers_generated" \
		--plural-exceptions "UISettings:UISettings" \
		"github.com/tigera/api/pkg/apis/projectcalico/v3"
# generate informer
informer-gen "$@" \
		--go-header-file "${REPO_ROOT}/hack/boilerplate/boilerplate.go.txt" \
		--versioned-clientset-package "github.com/tigera/api/pkg/client/clientset_generated/clientset" \
		--listers-package "github.com/tigera/api/pkg/client/listers_generated" \
		--output-dir "${REPO_ROOT}/pkg/client/informers_generated" \
		--output-pkg "github.com/tigera/api/pkg/client/informers_generated" \
		--plural-exceptions "UISettings:UISettings" \
		"github.com/tigera/api/pkg/apis/projectcalico/v3"

# Patch the fake Calico client to return a fake.RESTClient instead of nil.
# This prevents flakes in some tests where a background goroutine tries to
# use the REST client.
git apply -p2 < return-fake-rest-client.patch
