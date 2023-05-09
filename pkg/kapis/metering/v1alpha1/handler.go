/*

 Copyright 2019 The ImagineKube Authors.

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

package v1alpha1

import (
	"github.com/emicklei/go-restful"
	"k8s.io/client-go/kubernetes"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"imaginekube.com/imaginekube/pkg/models/openpitrix"

	"imaginekube.com/imaginekube/pkg/informers"
	monitorhle "imaginekube.com/imaginekube/pkg/kapis/monitoring/v1alpha3"
	resourcev1alpha3 "imaginekube.com/imaginekube/pkg/models/resources/v1alpha3/resource"
	meteringclient "imaginekube.com/imaginekube/pkg/simple/client/metering"
	"imaginekube.com/imaginekube/pkg/simple/client/monitoring"
)

type meterHandler interface {
	HandleClusterMeterQuery(req *restful.Request, resp *restful.Response)
	HandleNodeMeterQuery(req *restful.Request, resp *restful.Response)
	HandleWorkspaceMeterQuery(req *restful.Request, resp *restful.Response)
	HandleNamespaceMeterQuery(re *restful.Request, resp *restful.Response)
	HandleOpenpitrixMeterQuery(req *restful.Request, resp *restful.Response)
	HandleWorkloadMeterQuery(req *restful.Request, resp *restful.Response)
	HandleApplicationMeterQuery(req *restful.Request, resp *restful.Response)
	HandlePodMeterQuery(req *restful.Request, resp *restful.Response)
	HandleServiceMeterQuery(req *restful.Request, resp *restful.Response)
	HandlePVCMeterQuery(req *restful.Request, resp *restful.Response)
}

func newHandler(k kubernetes.Interface, m monitoring.Interface, f informers.InformerFactory, resourceGetter *resourcev1alpha3.ResourceGetter, meteringOptions *meteringclient.Options, opClient openpitrix.Interface, rtClient runtimeclient.Client) meterHandler {
	return monitorhle.NewHandler(k, m, nil, f, resourceGetter, meteringOptions, opClient, rtClient)
}
