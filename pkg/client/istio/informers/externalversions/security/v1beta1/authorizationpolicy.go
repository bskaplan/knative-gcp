/*
Copyright 2020 Google LLC

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

package v1beta1

import (
	time "time"

	versioned "github.com/google/knative-gcp/pkg/client/istio/clientset/versioned"
	internalinterfaces "github.com/google/knative-gcp/pkg/client/istio/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/google/knative-gcp/pkg/client/istio/listers/security/v1beta1"
	securityv1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AuthorizationPolicyInformer provides access to a shared informer and lister for
// AuthorizationPolicies.
type AuthorizationPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.AuthorizationPolicyLister
}

type authorizationPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAuthorizationPolicyInformer constructs a new informer for AuthorizationPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAuthorizationPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAuthorizationPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAuthorizationPolicyInformer constructs a new informer for AuthorizationPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAuthorizationPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecurityV1beta1().AuthorizationPolicies(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecurityV1beta1().AuthorizationPolicies(namespace).Watch(options)
			},
		},
		&securityv1beta1.AuthorizationPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *authorizationPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAuthorizationPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *authorizationPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&securityv1beta1.AuthorizationPolicy{}, f.defaultInformer)
}

func (f *authorizationPolicyInformer) Lister() v1beta1.AuthorizationPolicyLister {
	return v1beta1.NewAuthorizationPolicyLister(f.Informer().GetIndexer())
}
