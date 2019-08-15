// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/contiv/vpp/plugins/crd/pkg/apis/contivppio/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CustomConfigurationLister helps list CustomConfigurations.
type CustomConfigurationLister interface {
	// List lists all CustomConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1.CustomConfiguration, err error)
	// CustomConfigurations returns an object that can list and get CustomConfigurations.
	CustomConfigurations(namespace string) CustomConfigurationNamespaceLister
	CustomConfigurationListerExpansion
}

// customConfigurationLister implements the CustomConfigurationLister interface.
type customConfigurationLister struct {
	indexer cache.Indexer
}

// NewCustomConfigurationLister returns a new CustomConfigurationLister.
func NewCustomConfigurationLister(indexer cache.Indexer) CustomConfigurationLister {
	return &customConfigurationLister{indexer: indexer}
}

// List lists all CustomConfigurations in the indexer.
func (s *customConfigurationLister) List(selector labels.Selector) (ret []*v1.CustomConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CustomConfiguration))
	})
	return ret, err
}

// CustomConfigurations returns an object that can list and get CustomConfigurations.
func (s *customConfigurationLister) CustomConfigurations(namespace string) CustomConfigurationNamespaceLister {
	return customConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CustomConfigurationNamespaceLister helps list and get CustomConfigurations.
type CustomConfigurationNamespaceLister interface {
	// List lists all CustomConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.CustomConfiguration, err error)
	// Get retrieves the CustomConfiguration from the indexer for a given namespace and name.
	Get(name string) (*v1.CustomConfiguration, error)
	CustomConfigurationNamespaceListerExpansion
}

// customConfigurationNamespaceLister implements the CustomConfigurationNamespaceLister
// interface.
type customConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CustomConfigurations in the indexer for a given namespace.
func (s customConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1.CustomConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CustomConfiguration))
	})
	return ret, err
}

// Get retrieves the CustomConfiguration from the indexer for a given namespace and name.
func (s customConfigurationNamespaceLister) Get(name string) (*v1.CustomConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("customconfiguration"), name)
	}
	return obj.(*v1.CustomConfiguration), nil
}