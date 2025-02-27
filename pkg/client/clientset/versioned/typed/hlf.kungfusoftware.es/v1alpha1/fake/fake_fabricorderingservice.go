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

// FakeFabricOrderingServices implements FabricOrderingServiceInterface
type FakeFabricOrderingServices struct {
	Fake *FakeHlfV1alpha1
	ns   string
}

var fabricorderingservicesResource = v1alpha1.SchemeGroupVersion.WithResource("fabricorderingservices")

var fabricorderingservicesKind = v1alpha1.SchemeGroupVersion.WithKind("FabricOrderingService")

// Get takes name of the fabricOrderingService, and returns the corresponding fabricOrderingService object, and an error if there is any.
func (c *FakeFabricOrderingServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FabricOrderingService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fabricorderingservicesResource, c.ns, name), &v1alpha1.FabricOrderingService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOrderingService), err
}

// List takes label and field selectors, and returns the list of FabricOrderingServices that match those selectors.
func (c *FakeFabricOrderingServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FabricOrderingServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fabricorderingservicesResource, fabricorderingservicesKind, c.ns, opts), &v1alpha1.FabricOrderingServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FabricOrderingServiceList{ListMeta: obj.(*v1alpha1.FabricOrderingServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.FabricOrderingServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fabricOrderingServices.
func (c *FakeFabricOrderingServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fabricorderingservicesResource, c.ns, opts))

}

// Create takes the representation of a fabricOrderingService and creates it.  Returns the server's representation of the fabricOrderingService, and an error, if there is any.
func (c *FakeFabricOrderingServices) Create(ctx context.Context, fabricOrderingService *v1alpha1.FabricOrderingService, opts v1.CreateOptions) (result *v1alpha1.FabricOrderingService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fabricorderingservicesResource, c.ns, fabricOrderingService), &v1alpha1.FabricOrderingService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOrderingService), err
}

// Update takes the representation of a fabricOrderingService and updates it. Returns the server's representation of the fabricOrderingService, and an error, if there is any.
func (c *FakeFabricOrderingServices) Update(ctx context.Context, fabricOrderingService *v1alpha1.FabricOrderingService, opts v1.UpdateOptions) (result *v1alpha1.FabricOrderingService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fabricorderingservicesResource, c.ns, fabricOrderingService), &v1alpha1.FabricOrderingService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOrderingService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFabricOrderingServices) UpdateStatus(ctx context.Context, fabricOrderingService *v1alpha1.FabricOrderingService, opts v1.UpdateOptions) (*v1alpha1.FabricOrderingService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fabricorderingservicesResource, "status", c.ns, fabricOrderingService), &v1alpha1.FabricOrderingService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOrderingService), err
}

// Delete takes name of the fabricOrderingService and deletes it. Returns an error if one occurs.
func (c *FakeFabricOrderingServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(fabricorderingservicesResource, c.ns, name, opts), &v1alpha1.FabricOrderingService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFabricOrderingServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fabricorderingservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FabricOrderingServiceList{})
	return err
}

// Patch applies the patch and returns the patched fabricOrderingService.
func (c *FakeFabricOrderingServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FabricOrderingService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricorderingservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FabricOrderingService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOrderingService), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied fabricOrderingService.
func (c *FakeFabricOrderingServices) Apply(ctx context.Context, fabricOrderingService *hlfkungfusoftwareesv1alpha1.FabricOrderingServiceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricOrderingService, err error) {
	if fabricOrderingService == nil {
		return nil, fmt.Errorf("fabricOrderingService provided to Apply must not be nil")
	}
	data, err := json.Marshal(fabricOrderingService)
	if err != nil {
		return nil, err
	}
	name := fabricOrderingService.Name
	if name == nil {
		return nil, fmt.Errorf("fabricOrderingService.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricorderingservicesResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.FabricOrderingService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOrderingService), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeFabricOrderingServices) ApplyStatus(ctx context.Context, fabricOrderingService *hlfkungfusoftwareesv1alpha1.FabricOrderingServiceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.FabricOrderingService, err error) {
	if fabricOrderingService == nil {
		return nil, fmt.Errorf("fabricOrderingService provided to Apply must not be nil")
	}
	data, err := json.Marshal(fabricOrderingService)
	if err != nil {
		return nil, err
	}
	name := fabricOrderingService.Name
	if name == nil {
		return nil, fmt.Errorf("fabricOrderingService.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fabricorderingservicesResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.FabricOrderingService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FabricOrderingService), err
}
