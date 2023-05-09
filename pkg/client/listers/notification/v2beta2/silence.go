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

package v2beta2

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v2beta2 "imaginekube.com/api/notification/v2beta2"
)

// SilenceLister helps list Silences.
// All objects returned here must be treated as read-only.
type SilenceLister interface {
	// List lists all Silences in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2beta2.Silence, err error)
	// Get retrieves the Silence from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2beta2.Silence, error)
	SilenceListerExpansion
}

// silenceLister implements the SilenceLister interface.
type silenceLister struct {
	indexer cache.Indexer
}

// NewSilenceLister returns a new SilenceLister.
func NewSilenceLister(indexer cache.Indexer) SilenceLister {
	return &silenceLister{indexer: indexer}
}

// List lists all Silences in the indexer.
func (s *silenceLister) List(selector labels.Selector) (ret []*v2beta2.Silence, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2beta2.Silence))
	})
	return ret, err
}

// Get retrieves the Silence from the index for a given name.
func (s *silenceLister) Get(name string) (*v2beta2.Silence, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2beta2.Resource("silence"), name)
	}
	return obj.(*v2beta2.Silence), nil
}
