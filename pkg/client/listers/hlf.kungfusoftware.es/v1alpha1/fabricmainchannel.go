/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FabricMainChannelLister helps list FabricMainChannels.
// All objects returned here must be treated as read-only.
type FabricMainChannelLister interface {
	// List lists all FabricMainChannels in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricMainChannel, err error)
	// Get retrieves the FabricMainChannel from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FabricMainChannel, error)
	FabricMainChannelListerExpansion
}

// fabricMainChannelLister implements the FabricMainChannelLister interface.
type fabricMainChannelLister struct {
	indexer cache.Indexer
}

// NewFabricMainChannelLister returns a new FabricMainChannelLister.
func NewFabricMainChannelLister(indexer cache.Indexer) FabricMainChannelLister {
	return &fabricMainChannelLister{indexer: indexer}
}

// List lists all FabricMainChannels in the indexer.
func (s *fabricMainChannelLister) List(selector labels.Selector) (ret []*v1alpha1.FabricMainChannel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricMainChannel))
	})
	return ret, err
}

// Get retrieves the FabricMainChannel from the index for a given name.
func (s *fabricMainChannelLister) Get(name string) (*v1alpha1.FabricMainChannel, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fabricmainchannel"), name)
	}
	return obj.(*v1alpha1.FabricMainChannel), nil
}
