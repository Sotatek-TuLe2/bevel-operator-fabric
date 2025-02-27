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

// FakeFabricOperatorUIs implements FabricOperatorUIInterface
type FakeFabricOperatorUIs struct {
	Fake *FakeHlfV1alpha1
	ns   string
}

var fabricoperatoruisResource = v1alpha1.SchemeGroupVersion.WithResource("fabricoperatoruis")

var fabricoperatoruisKind = v1alpha1.SchemeGroupVersion.WithKind("FabricOperatorUI")

// Get takes name of the fabricOperatorUI, and returns the corresponding fabricOperatorUI object, and an error if there is any.
func (c *FakeFabricOperatorUIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FabricOperatorUI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fabricoperatoruisResource, c.ns, name), &v1alpha1.FabricOperatorUI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperatorUI), err
}

// List takes label and field selectors, and returns the list of FabricOperatorUIs that match those selectors.
func (c *FakeFabricOperatorUIs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FabricOperatorUIList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fabricoperatoruisResource, fabricoperatoruisKind, c.ns, opts), &v1alpha1.FabricOperatorUIList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FabricOperatorUIList{ListMeta: obj.(*v1alpha1.FabricOperatorUIList).ListMeta}
	for _, item := range obj.(*v1alpha1.FabricOperatorUIList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fabricOperatorUIs.
func (c *FakeFabricOperatorUIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fabricoperatoruisResource, c.ns, opts))

}

// Create takes the representation of a fabricOperatorUI and creates it.  Returns the server's representation of the fabricOperatorUI, and an error, if there is any.
func (c *FakeFabricOperatorUIs) Create(ctx context.Context, fabricOperatorUI *v1alpha1.FabricOperatorUI, opts v1.CreateOptions) (result *v1alpha1.FabricOperatorUI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fabricoperatoruisResource, c.ns, fabricOperatorUI), &v1alpha1.FabricOperatorUI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperatorUI), err
}

// Update takes the representation of a fabricOperatorUI and updates it. Returns the server's representation of the fabricOperatorUI, and an error, if there is any.
func (c *FakeFabricOperatorUIs) Update(ctx context.Context, fabricOperatorUI *v1alpha1.FabricOperatorUI, opts v1.UpdateOptions) (result *v1alpha1.FabricOperatorUI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fabricoperatoruisResource, c.ns, fabricOperatorUI), &v1alpha1.FabricOperatorUI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperatorUI), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFabricOperatorUIs) UpdateStatus(ctx context.Context, fabricOperatorUI *v1alpha1.FabricOperatorUI, opts v1.UpdateOptions) (*v1alpha1.FabricOperatorUI, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fabricoperatoruisResource, "status", c.ns, fabricOperatorUI), &v1alpha1.FabricOperatorUI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperatorUI), err
}

// Delete takes name of the fabricOperatorUI and deletes it. Returns an error if one occurs.
func (c *FakeFabricOperatorUIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(fabricoperatoruisResource, c.ns, name, opts), &v1alpha1.FabricOperatorUI{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFabricOperatorUIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fabricoperatoruisResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FabricOperatorUIList{})
	return err
}

// Patch applies the patch and returns the patched fabricOperatorUI.
func (c *FakeFabricOperatorUIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricOperatorUI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricoperatoruisResource, c.ns, name, pt, data, subresources...), &v1alpha1.FabricOperatorUI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperatorUI), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied fabricOperatorUI.
func (c *FakeFabricOperatorUIs) Apply(ctx context.Context, fabricOperatorUI *hlfkungfusoftwareesv1alpha1.FabricOperatorUIApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricOperatorUI, err error) {
	if fabricOperatorUI == nil {
		return nil, fmt.Errorf("fabricOperatorUI provided to Apply must not be nil")
	}
	data, err := json.Marshal(fabricOperatorUI)
	if err != nil {
		return nil, err
	}
	name := fabricOperatorUI.Name
	if name == nil {
		return nil, fmt.Errorf("fabricOperatorUI.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricoperatoruisResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.FabricOperatorUI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperatorUI), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeFabricOperatorUIs) ApplyStatus(ctx context.Context, fabricOperatorUI *hlfkungfusoftwareesv1alpha1.FabricOperatorUIApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricOperatorUI, err error) {
	if fabricOperatorUI == nil {
		return nil, fmt.Errorf("fabricOperatorUI provided to Apply must not be nil")
	}
	data, err := json.Marshal(fabricOperatorUI)
	if err != nil {
		return nil, err
	}
	name := fabricOperatorUI.Name
	if name == nil {
		return nil, fmt.Errorf("fabricOperatorUI.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricoperatoruisResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.FabricOperatorUI{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOperatorUI), err
}
