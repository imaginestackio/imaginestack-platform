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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	networkv1alpha1 "imaginekube.com/api/network/v1alpha1"
	versioned "imaginekube.com/imaginekube/pkg/client/clientset/versioned"
	internalinterfaces "imaginekube.com/imaginekube/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "imaginekube.com/imaginekube/pkg/client/listers/network/v1alpha1"
)

// IPAMHandleInformer provides access to a shared informer and lister for
// IPAMHandles.
type IPAMHandleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.IPAMHandleLister
}

type iPAMHandleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewIPAMHandleInformer constructs a new informer for IPAMHandle type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIPAMHandleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIPAMHandleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredIPAMHandleInformer constructs a new informer for IPAMHandle type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIPAMHandleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1alpha1().IPAMHandles().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1alpha1().IPAMHandles().Watch(context.TODO(), options)
			},
		},
		&networkv1alpha1.IPAMHandle{},
		resyncPeriod,
		indexers,
	)
}

func (f *iPAMHandleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIPAMHandleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *iPAMHandleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkv1alpha1.IPAMHandle{}, f.defaultInformer)
}

func (f *iPAMHandleInformer) Lister() v1alpha1.IPAMHandleLister {
	return v1alpha1.NewIPAMHandleLister(f.Informer().GetIndexer())
}
