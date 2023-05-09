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

package v1alpha1

import (
	rest "k8s.io/client-go/rest"
	v1alpha1 "imaginekube.com/api/auditing/v1alpha1"
	"imaginekube.com/imaginekube/pkg/client/clientset/versioned/scheme"
)

type AuditingV1alpha1Interface interface {
	RESTClient() rest.Interface
	RulesGetter
	WebhooksGetter
}

// AuditingV1alpha1Client is used to interact with features provided by the auditing.imaginekube.com group.
type AuditingV1alpha1Client struct {
	restClient rest.Interface
}

func (c *AuditingV1alpha1Client) Rules() RuleInterface {
	return newRules(c)
}

func (c *AuditingV1alpha1Client) Webhooks() WebhookInterface {
	return newWebhooks(c)
}

// NewForConfig creates a new AuditingV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*AuditingV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &AuditingV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new AuditingV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AuditingV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AuditingV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *AuditingV1alpha1Client {
	return &AuditingV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
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
func (c *AuditingV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
