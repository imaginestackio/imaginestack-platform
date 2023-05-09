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

package v2beta2

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	notificationv2beta2 "imaginekube.com/api/notification/v2beta2"
	versioned "imaginekube.com/imaginekube/pkg/client/clientset/versioned"
	internalinterfaces "imaginekube.com/imaginekube/pkg/client/informers/externalversions/internalinterfaces"
	v2beta2 "imaginekube.com/imaginekube/pkg/client/listers/notification/v2beta2"
)

// SilenceInformer provides access to a shared informer and lister for
// Silences.
type SilenceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2beta2.SilenceLister
}

type silenceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewSilenceInformer constructs a new informer for Silence type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSilenceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSilenceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredSilenceInformer constructs a new informer for Silence type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSilenceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NotificationV2beta2().Silences().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NotificationV2beta2().Silences().Watch(context.TODO(), options)
			},
		},
		&notificationv2beta2.Silence{},
		resyncPeriod,
		indexers,
	)
}

func (f *silenceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSilenceInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *silenceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&notificationv2beta2.Silence{}, f.defaultInformer)
}

func (f *silenceInformer) Lister() v2beta2.SilenceLister {
	return v2beta2.NewSilenceLister(f.Informer().GetIndexer())
}
