/*
Copyright 2022 Nutanix, Inc

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

package interfaces

import (
	prismClientV3 "github.com/nutanix-cloud-native/prism-go-client/v3"
	"k8s.io/client-go/informers"
)

type Client interface {
	Get() (Prism, error)
	SetInformers(sharedInformers informers.SharedInformerFactory)
}

type Prism interface {
	GetVM(vmUUID string) (*prismClientV3.VMIntentResponse, error)
	GetCluster(clusterUUID string) (*prismClientV3.ClusterIntentResponse, error)
	ListAllCluster(filter string) (*prismClientV3.ClusterListIntentResponse, error)
}
