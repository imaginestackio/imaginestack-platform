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

package v1beta2

import (
	rest "k8s.io/client-go/rest"
	v1beta2 "imaginekube.com/api/types/v1beta2"
	"imaginekube.com/imaginekube/pkg/client/clientset/versioned/scheme"
)

type TypesV1beta2Interface interface {
	RESTClient() rest.Interface
	FederatedNotificationConfigsGetter
	FederatedNotificationManagersGetter
	FederatedNotificationReceiversGetter
	FederatedNotificationRoutersGetter
	FederatedNotificationSilencesGetter
}

// TypesV1beta2Client is used to interact with features provided by the types.kubefed.io group.
type TypesV1beta2Client struct {
	restClient rest.Interface
}

func (c *TypesV1beta2Client) FederatedNotificationConfigs() FederatedNotificationConfigInterface {
	return newFederatedNotificationConfigs(c)
}

func (c *TypesV1beta2Client) FederatedNotificationManagers() FederatedNotificationManagerInterface {
	return newFederatedNotificationManagers(c)
}

func (c *TypesV1beta2Client) FederatedNotificationReceivers() FederatedNotificationReceiverInterface {
	return newFederatedNotificationReceivers(c)
}

func (c *TypesV1beta2Client) FederatedNotificationRouters() FederatedNotificationRouterInterface {
	return newFederatedNotificationRouters(c)
}

func (c *TypesV1beta2Client) FederatedNotificationSilences() FederatedNotificationSilenceInterface {
	return newFederatedNotificationSilences(c)
}

// NewForConfig creates a new TypesV1beta2Client for the given config.
func NewForConfig(c *rest.Config) (*TypesV1beta2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TypesV1beta2Client{client}, nil
}

// NewForConfigOrDie creates a new TypesV1beta2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TypesV1beta2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TypesV1beta2Client for the given RESTClient.
func New(c rest.Interface) *TypesV1beta2Client {
	return &TypesV1beta2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TypesV1beta2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
