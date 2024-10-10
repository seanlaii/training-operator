// Copyright 2024 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	time "time"

	kubefloworgv2alpha1 "github.com/kubeflow/training-operator/pkg/apis/kubeflow.org/v2alpha1"
	versioned "github.com/kubeflow/training-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeflow/training-operator/pkg/client/informers/externalversions/internalinterfaces"
	v2alpha1 "github.com/kubeflow/training-operator/pkg/client/listers/kubeflow.org/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TrainingRuntimeInformer provides access to a shared informer and lister for
// TrainingRuntimes.
type TrainingRuntimeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2alpha1.TrainingRuntimeLister
}

type trainingRuntimeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTrainingRuntimeInformer constructs a new informer for TrainingRuntime type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTrainingRuntimeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTrainingRuntimeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTrainingRuntimeInformer constructs a new informer for TrainingRuntime type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTrainingRuntimeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV2alpha1().TrainingRuntimes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV2alpha1().TrainingRuntimes(namespace).Watch(context.TODO(), options)
			},
		},
		&kubefloworgv2alpha1.TrainingRuntime{},
		resyncPeriod,
		indexers,
	)
}

func (f *trainingRuntimeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTrainingRuntimeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *trainingRuntimeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubefloworgv2alpha1.TrainingRuntime{}, f.defaultInformer)
}

func (f *trainingRuntimeInformer) Lister() v2alpha1.TrainingRuntimeLister {
	return v2alpha1.NewTrainingRuntimeLister(f.Informer().GetIndexer())
}
