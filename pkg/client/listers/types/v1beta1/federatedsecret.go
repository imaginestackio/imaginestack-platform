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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "imaginekube.com/api/types/v1beta1"
)

// FederatedSecretLister helps list FederatedSecrets.
// All objects returned here must be treated as read-only.
type FederatedSecretLister interface {
	// List lists all FederatedSecrets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.FederatedSecret, err error)
	// FederatedSecrets returns an object that can list and get FederatedSecrets.
	FederatedSecrets(namespace string) FederatedSecretNamespaceLister
	FederatedSecretListerExpansion
}

// federatedSecretLister implements the FederatedSecretLister interface.
type federatedSecretLister struct {
	indexer cache.Indexer
}

// NewFederatedSecretLister returns a new FederatedSecretLister.
func NewFederatedSecretLister(indexer cache.Indexer) FederatedSecretLister {
	return &federatedSecretLister{indexer: indexer}
}

// List lists all FederatedSecrets in the indexer.
func (s *federatedSecretLister) List(selector labels.Selector) (ret []*v1beta1.FederatedSecret, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedSecret))
	})
	return ret, err
}

// FederatedSecrets returns an object that can list and get FederatedSecrets.
func (s *federatedSecretLister) FederatedSecrets(namespace string) FederatedSecretNamespaceLister {
	return federatedSecretNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedSecretNamespaceLister helps list and get FederatedSecrets.
// All objects returned here must be treated as read-only.
type FederatedSecretNamespaceLister interface {
	// List lists all FederatedSecrets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.FederatedSecret, err error)
	// Get retrieves the FederatedSecret from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.FederatedSecret, error)
	FederatedSecretNamespaceListerExpansion
}

// federatedSecretNamespaceLister implements the FederatedSecretNamespaceLister
// interface.
type federatedSecretNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedSecrets in the indexer for a given namespace.
func (s federatedSecretNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.FederatedSecret, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.FederatedSecret))
	})
	return ret, err
}

// Get retrieves the FederatedSecret from the indexer for a given namespace and name.
func (s federatedSecretNamespaceLister) Get(name string) (*v1beta1.FederatedSecret, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("federatedsecret"), name)
	}
	return obj.(*v1beta1.FederatedSecret), nil
}
