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

package calicov3

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	calicov3 "imaginekube.com/api/network/calicov3"
)

// IPPoolLister helps list IPPools.
type IPPoolLister interface {
	// List lists all IPPools in the indexer.
	List(selector labels.Selector) (ret []*calicov3.IPPool, err error)
	// Get retrieves the IPPool from the index for a given name.
	Get(name string) (*calicov3.IPPool, error)
	IPPoolListerExpansion
}

// iPPoolLister implements the IPPoolLister interface.
type iPPoolLister struct {
	indexer cache.Indexer
}

// NewIPPoolLister returns a new IPPoolLister.
func NewIPPoolLister(indexer cache.Indexer) IPPoolLister {
	return &iPPoolLister{indexer: indexer}
}

// List lists all IPPools in the indexer.
func (s *iPPoolLister) List(selector labels.Selector) (ret []*calicov3.IPPool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*calicov3.IPPool))
	})
	return ret, err
}

// Get retrieves the IPPool from the index for a given name.
func (s *iPPoolLister) Get(name string) (*calicov3.IPPool, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(calicov3.Resource("ippool"), name)
	}
	return obj.(*calicov3.IPPool), nil
}
