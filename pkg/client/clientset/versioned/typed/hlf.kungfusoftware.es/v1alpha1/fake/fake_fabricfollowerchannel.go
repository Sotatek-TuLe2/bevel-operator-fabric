/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	hlfkungfusoftwareesv1alpha1 "github.com/kfsoftware/hlf-operator/pkg/client/applyconfiguration/hlf.kungfusoftware.es/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFabricFollowerChannels implements FabricFollowerChannelInterface
type FakeFabricFollowerChannels struct {
	Fake *FakeHlfV1alpha1
}

var fabricfollowerchannelsResource = v1alpha1.SchemeGroupVersion.WithResource("fabricfollowerchannels")

var fabricfollowerchannelsKind = v1alpha1.SchemeGroupVersion.WithKind("FabricFollowerChannel")

// Get takes name of the fabricFollowerChannel, and returns the corresponding fabricFollowerChannel object, and an error if there is any.
func (c *FakeFabricFollowerChannels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FabricFollowerChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(fabricfollowerchannelsResource, name), &v1alpha1.FabricFollowerChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricFollowerChannel), err
}

// List takes label and field selectors, and returns the list of FabricFollowerChannels that match those selectors.
func (c *FakeFabricFollowerChannels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FabricFollowerChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(fabricfollowerchannelsResource, fabricfollowerchannelsKind, opts), &v1alpha1.FabricFollowerChannelList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FabricFollowerChannelList{ListMeta: obj.(*v1alpha1.FabricFollowerChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.FabricFollowerChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fabricFollowerChannels.
func (c *FakeFabricFollowerChannels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(fabricfollowerchannelsResource, opts))
}

// Create takes the representation of a fabricFollowerChannel and creates it.  Returns the server's representation of the fabricFollowerChannel, and an error, if there is any.
func (c *FakeFabricFollowerChannels) Create(ctx context.Context, fabricFollowerChannel *v1alpha1.FabricFollowerChannel, opts v1.CreateOptions) (result *v1alpha1.FabricFollowerChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(fabricfollowerchannelsResource, fabricFollowerChannel), &v1alpha1.FabricFollowerChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricFollowerChannel), err
}

// Update takes the representation of a fabricFollowerChannel and updates it. Returns the server's representation of the fabricFollowerChannel, and an error, if there is any.
func (c *FakeFabricFollowerChannels) Update(ctx context.Context, fabricFollowerChannel *v1alpha1.FabricFollowerChannel, opts v1.UpdateOptions) (result *v1alpha1.FabricFollowerChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(fabricfollowerchannelsResource, fabricFollowerChannel), &v1alpha1.FabricFollowerChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricFollowerChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFabricFollowerChannels) UpdateStatus(ctx context.Context, fabricFollowerChannel *v1alpha1.FabricFollowerChannel, opts v1.UpdateOptions) (*v1alpha1.FabricFollowerChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(fabricfollowerchannelsResource, "status", fabricFollowerChannel), &v1alpha1.FabricFollowerChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricFollowerChannel), err
}

// Delete takes name of the fabricFollowerChannel and deletes it. Returns an error if one occurs.
func (c *FakeFabricFollowerChannels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(fabricfollowerchannelsResource, name, opts), &v1alpha1.FabricFollowerChannel{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFabricFollowerChannels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(fabricfollowerchannelsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FabricFollowerChannelList{})
	return err
}

// Patch applies the patch and returns the patched fabricFollowerChannel.
func (c *FakeFabricFollowerChannels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricFollowerChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(fabricfollowerchannelsResource, name, pt, data, subresources...), &v1alpha1.FabricFollowerChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricFollowerChannel), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied fabricFollowerChannel.
func (c *FakeFabricFollowerChannels) Apply(ctx context.Context, fabricFollowerChannel *hlfkungfusoftwareesv1alpha1.FabricFollowerChannelApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricFollowerChannel, err error) {
	if fabricFollowerChannel == nil {
		return nil, fmt.Errorf("fabricFollowerChannel provided to Apply must not be nil")
	}
	data, err := json.Marshal(fabricFollowerChannel)
	if err != nil {
		return nil, err
	}
	name := fabricFollowerChannel.Name
	if name == nil {
		return nil, fmt.Errorf("fabricFollowerChannel.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(fabricfollowerchannelsResource, *name, types.ApplyPatchType, data), &v1alpha1.FabricFollowerChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricFollowerChannel), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeFabricFollowerChannels) ApplyStatus(ctx context.Context, fabricFollowerChannel *hlfkungfusoftwareesv1alpha1.FabricFollowerChannelApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricFollowerChannel, err error) {
	if fabricFollowerChannel == nil {
		return nil, fmt.Errorf("fabricFollowerChannel provided to Apply must not be nil")
	}
	data, err := json.Marshal(fabricFollowerChannel)
	if err != nil {
		return nil, err
	}
	name := fabricFollowerChannel.Name
	if name == nil {
		return nil, fmt.Errorf("fabricFollowerChannel.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(fabricfollowerchannelsResource, *name, types.ApplyPatchType, data, "status"), &v1alpha1.FabricFollowerChannel{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricFollowerChannel), err
}
