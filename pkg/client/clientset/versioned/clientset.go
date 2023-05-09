/*
Copyright 2023 The ImagineKube Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	alertingv2beta1 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/alerting/v2beta1"
	applicationv1alpha1 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/application/v1alpha1"
	auditingv1alpha1 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/auditing/v1alpha1"
	clusterv1alpha1 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/cluster/v1alpha1"
	devopsv1alpha1 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/devops/v1alpha1"
	devopsv1alpha3 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/devops/v1alpha3"
	iamv1alpha2 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/iam/v1alpha2"
	networkv1alpha1 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/network/v1alpha1"
	notificationv2beta1 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/notification/v2beta1"
	notificationv2beta2 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/notification/v2beta2"
	quotav1alpha2 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/quota/v1alpha2"
	servicemeshv1alpha2 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/servicemesh/v1alpha2"
	storagev1alpha1 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/storage/v1alpha1"
	tenantv1alpha1 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/tenant/v1alpha1"
	tenantv1alpha2 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/tenant/v1alpha2"
	typesv1beta1 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/types/v1beta1"
	typesv1beta2 "imaginekube.com/imaginekube/pkg/client/clientset/versioned/typed/types/v1beta2"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AlertingV2beta1() alertingv2beta1.AlertingV2beta1Interface
	ApplicationV1alpha1() applicationv1alpha1.ApplicationV1alpha1Interface
	AuditingV1alpha1() auditingv1alpha1.AuditingV1alpha1Interface
	ClusterV1alpha1() clusterv1alpha1.ClusterV1alpha1Interface
	DevopsV1alpha1() devopsv1alpha1.DevopsV1alpha1Interface
	DevopsV1alpha3() devopsv1alpha3.DevopsV1alpha3Interface
	IamV1alpha2() iamv1alpha2.IamV1alpha2Interface
	NetworkV1alpha1() networkv1alpha1.NetworkV1alpha1Interface
	NotificationV2beta1() notificationv2beta1.NotificationV2beta1Interface
	NotificationV2beta2() notificationv2beta2.NotificationV2beta2Interface
	QuotaV1alpha2() quotav1alpha2.QuotaV1alpha2Interface
	ServicemeshV1alpha2() servicemeshv1alpha2.ServicemeshV1alpha2Interface
	StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface
	TenantV1alpha1() tenantv1alpha1.TenantV1alpha1Interface
	TenantV1alpha2() tenantv1alpha2.TenantV1alpha2Interface
	TypesV1beta1() typesv1beta1.TypesV1beta1Interface
	TypesV1beta2() typesv1beta2.TypesV1beta2Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	alertingV2beta1     *alertingv2beta1.AlertingV2beta1Client
	applicationV1alpha1 *applicationv1alpha1.ApplicationV1alpha1Client
	auditingV1alpha1    *auditingv1alpha1.AuditingV1alpha1Client
	clusterV1alpha1     *clusterv1alpha1.ClusterV1alpha1Client
	devopsV1alpha1      *devopsv1alpha1.DevopsV1alpha1Client
	devopsV1alpha3      *devopsv1alpha3.DevopsV1alpha3Client
	iamV1alpha2         *iamv1alpha2.IamV1alpha2Client
	networkV1alpha1     *networkv1alpha1.NetworkV1alpha1Client
	notificationV2beta1 *notificationv2beta1.NotificationV2beta1Client
	notificationV2beta2 *notificationv2beta2.NotificationV2beta2Client
	quotaV1alpha2       *quotav1alpha2.QuotaV1alpha2Client
	servicemeshV1alpha2 *servicemeshv1alpha2.ServicemeshV1alpha2Client
	storageV1alpha1     *storagev1alpha1.StorageV1alpha1Client
	tenantV1alpha1      *tenantv1alpha1.TenantV1alpha1Client
	tenantV1alpha2      *tenantv1alpha2.TenantV1alpha2Client
	typesV1beta1        *typesv1beta1.TypesV1beta1Client
	typesV1beta2        *typesv1beta2.TypesV1beta2Client
}

// AlertingV2beta1 retrieves the AlertingV2beta1Client
func (c *Clientset) AlertingV2beta1() alertingv2beta1.AlertingV2beta1Interface {
	return c.alertingV2beta1
}

// ApplicationV1alpha1 retrieves the ApplicationV1alpha1Client
func (c *Clientset) ApplicationV1alpha1() applicationv1alpha1.ApplicationV1alpha1Interface {
	return c.applicationV1alpha1
}

// AuditingV1alpha1 retrieves the AuditingV1alpha1Client
func (c *Clientset) AuditingV1alpha1() auditingv1alpha1.AuditingV1alpha1Interface {
	return c.auditingV1alpha1
}

// ClusterV1alpha1 retrieves the ClusterV1alpha1Client
func (c *Clientset) ClusterV1alpha1() clusterv1alpha1.ClusterV1alpha1Interface {
	return c.clusterV1alpha1
}

// DevopsV1alpha1 retrieves the DevopsV1alpha1Client
func (c *Clientset) DevopsV1alpha1() devopsv1alpha1.DevopsV1alpha1Interface {
	return c.devopsV1alpha1
}

// DevopsV1alpha3 retrieves the DevopsV1alpha3Client
func (c *Clientset) DevopsV1alpha3() devopsv1alpha3.DevopsV1alpha3Interface {
	return c.devopsV1alpha3
}

// IamV1alpha2 retrieves the IamV1alpha2Client
func (c *Clientset) IamV1alpha2() iamv1alpha2.IamV1alpha2Interface {
	return c.iamV1alpha2
}

// NetworkV1alpha1 retrieves the NetworkV1alpha1Client
func (c *Clientset) NetworkV1alpha1() networkv1alpha1.NetworkV1alpha1Interface {
	return c.networkV1alpha1
}

// NotificationV2beta1 retrieves the NotificationV2beta1Client
func (c *Clientset) NotificationV2beta1() notificationv2beta1.NotificationV2beta1Interface {
	return c.notificationV2beta1
}

// NotificationV2beta2 retrieves the NotificationV2beta2Client
func (c *Clientset) NotificationV2beta2() notificationv2beta2.NotificationV2beta2Interface {
	return c.notificationV2beta2
}

// QuotaV1alpha2 retrieves the QuotaV1alpha2Client
func (c *Clientset) QuotaV1alpha2() quotav1alpha2.QuotaV1alpha2Interface {
	return c.quotaV1alpha2
}

// ServicemeshV1alpha2 retrieves the ServicemeshV1alpha2Client
func (c *Clientset) ServicemeshV1alpha2() servicemeshv1alpha2.ServicemeshV1alpha2Interface {
	return c.servicemeshV1alpha2
}

// StorageV1alpha1 retrieves the StorageV1alpha1Client
func (c *Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return c.storageV1alpha1
}

// TenantV1alpha1 retrieves the TenantV1alpha1Client
func (c *Clientset) TenantV1alpha1() tenantv1alpha1.TenantV1alpha1Interface {
	return c.tenantV1alpha1
}

// TenantV1alpha2 retrieves the TenantV1alpha2Client
func (c *Clientset) TenantV1alpha2() tenantv1alpha2.TenantV1alpha2Interface {
	return c.tenantV1alpha2
}

// TypesV1beta1 retrieves the TypesV1beta1Client
func (c *Clientset) TypesV1beta1() typesv1beta1.TypesV1beta1Interface {
	return c.typesV1beta1
}

// TypesV1beta2 retrieves the TypesV1beta2Client
func (c *Clientset) TypesV1beta2() typesv1beta2.TypesV1beta2Interface {
	return c.typesV1beta2
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.alertingV2beta1, err = alertingv2beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.applicationV1alpha1, err = applicationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.auditingV1alpha1, err = auditingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.clusterV1alpha1, err = clusterv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.devopsV1alpha1, err = devopsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.devopsV1alpha3, err = devopsv1alpha3.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.iamV1alpha2, err = iamv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkV1alpha1, err = networkv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.notificationV2beta1, err = notificationv2beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.notificationV2beta2, err = notificationv2beta2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.quotaV1alpha2, err = quotav1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.servicemeshV1alpha2, err = servicemeshv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storageV1alpha1, err = storagev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.tenantV1alpha1, err = tenantv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.tenantV1alpha2, err = tenantv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.typesV1beta1, err = typesv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.typesV1beta2, err = typesv1beta2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.alertingV2beta1 = alertingv2beta1.NewForConfigOrDie(c)
	cs.applicationV1alpha1 = applicationv1alpha1.NewForConfigOrDie(c)
	cs.auditingV1alpha1 = auditingv1alpha1.NewForConfigOrDie(c)
	cs.clusterV1alpha1 = clusterv1alpha1.NewForConfigOrDie(c)
	cs.devopsV1alpha1 = devopsv1alpha1.NewForConfigOrDie(c)
	cs.devopsV1alpha3 = devopsv1alpha3.NewForConfigOrDie(c)
	cs.iamV1alpha2 = iamv1alpha2.NewForConfigOrDie(c)
	cs.networkV1alpha1 = networkv1alpha1.NewForConfigOrDie(c)
	cs.notificationV2beta1 = notificationv2beta1.NewForConfigOrDie(c)
	cs.notificationV2beta2 = notificationv2beta2.NewForConfigOrDie(c)
	cs.quotaV1alpha2 = quotav1alpha2.NewForConfigOrDie(c)
	cs.servicemeshV1alpha2 = servicemeshv1alpha2.NewForConfigOrDie(c)
	cs.storageV1alpha1 = storagev1alpha1.NewForConfigOrDie(c)
	cs.tenantV1alpha1 = tenantv1alpha1.NewForConfigOrDie(c)
	cs.tenantV1alpha2 = tenantv1alpha2.NewForConfigOrDie(c)
	cs.typesV1beta1 = typesv1beta1.NewForConfigOrDie(c)
	cs.typesV1beta2 = typesv1beta2.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.alertingV2beta1 = alertingv2beta1.New(c)
	cs.applicationV1alpha1 = applicationv1alpha1.New(c)
	cs.auditingV1alpha1 = auditingv1alpha1.New(c)
	cs.clusterV1alpha1 = clusterv1alpha1.New(c)
	cs.devopsV1alpha1 = devopsv1alpha1.New(c)
	cs.devopsV1alpha3 = devopsv1alpha3.New(c)
	cs.iamV1alpha2 = iamv1alpha2.New(c)
	cs.networkV1alpha1 = networkv1alpha1.New(c)
	cs.notificationV2beta1 = notificationv2beta1.New(c)
	cs.notificationV2beta2 = notificationv2beta2.New(c)
	cs.quotaV1alpha2 = quotav1alpha2.New(c)
	cs.servicemeshV1alpha2 = servicemeshv1alpha2.New(c)
	cs.storageV1alpha1 = storagev1alpha1.New(c)
	cs.tenantV1alpha1 = tenantv1alpha1.New(c)
	cs.tenantV1alpha2 = tenantv1alpha2.New(c)
	cs.typesV1beta1 = typesv1beta1.New(c)
	cs.typesV1beta2 = typesv1beta2.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
