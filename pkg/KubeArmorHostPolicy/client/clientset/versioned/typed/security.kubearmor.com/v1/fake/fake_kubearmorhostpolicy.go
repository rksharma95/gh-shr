// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of KubeArmor

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	securitykubearmorcomv1 "github.com/kubearmor/KubeArmor/pkg/KubeArmorHostPolicy/api/security.kubearmor.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeArmorHostPolicies implements KubeArmorHostPolicyInterface
type FakeKubeArmorHostPolicies struct {
	Fake *FakeSecurityV1
}

var kubearmorhostpoliciesResource = schema.GroupVersionResource{Group: "security.kubearmor.com", Version: "v1", Resource: "kubearmorhostpolicies"}

var kubearmorhostpoliciesKind = schema.GroupVersionKind{Group: "security.kubearmor.com", Version: "v1", Kind: "KubeArmorHostPolicy"}

// Get takes name of the kubeArmorHostPolicy, and returns the corresponding kubeArmorHostPolicy object, and an error if there is any.
func (c *FakeKubeArmorHostPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *securitykubearmorcomv1.KubeArmorHostPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kubearmorhostpoliciesResource, name), &securitykubearmorcomv1.KubeArmorHostPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*securitykubearmorcomv1.KubeArmorHostPolicy), err
}

// List takes label and field selectors, and returns the list of KubeArmorHostPolicies that match those selectors.
func (c *FakeKubeArmorHostPolicies) List(ctx context.Context, opts v1.ListOptions) (result *securitykubearmorcomv1.KubeArmorHostPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kubearmorhostpoliciesResource, kubearmorhostpoliciesKind, opts), &securitykubearmorcomv1.KubeArmorHostPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &securitykubearmorcomv1.KubeArmorHostPolicyList{ListMeta: obj.(*securitykubearmorcomv1.KubeArmorHostPolicyList).ListMeta}
	for _, item := range obj.(*securitykubearmorcomv1.KubeArmorHostPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeArmorHostPolicies.
func (c *FakeKubeArmorHostPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kubearmorhostpoliciesResource, opts))
}

// Create takes the representation of a kubeArmorHostPolicy and creates it.  Returns the server's representation of the kubeArmorHostPolicy, and an error, if there is any.
func (c *FakeKubeArmorHostPolicies) Create(ctx context.Context, kubeArmorHostPolicy *securitykubearmorcomv1.KubeArmorHostPolicy, opts v1.CreateOptions) (result *securitykubearmorcomv1.KubeArmorHostPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kubearmorhostpoliciesResource, kubeArmorHostPolicy), &securitykubearmorcomv1.KubeArmorHostPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*securitykubearmorcomv1.KubeArmorHostPolicy), err
}

// Update takes the representation of a kubeArmorHostPolicy and updates it. Returns the server's representation of the kubeArmorHostPolicy, and an error, if there is any.
func (c *FakeKubeArmorHostPolicies) Update(ctx context.Context, kubeArmorHostPolicy *securitykubearmorcomv1.KubeArmorHostPolicy, opts v1.UpdateOptions) (result *securitykubearmorcomv1.KubeArmorHostPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kubearmorhostpoliciesResource, kubeArmorHostPolicy), &securitykubearmorcomv1.KubeArmorHostPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*securitykubearmorcomv1.KubeArmorHostPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeArmorHostPolicies) UpdateStatus(ctx context.Context, kubeArmorHostPolicy *securitykubearmorcomv1.KubeArmorHostPolicy, opts v1.UpdateOptions) (*securitykubearmorcomv1.KubeArmorHostPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kubearmorhostpoliciesResource, "status", kubeArmorHostPolicy), &securitykubearmorcomv1.KubeArmorHostPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*securitykubearmorcomv1.KubeArmorHostPolicy), err
}

// Delete takes name of the kubeArmorHostPolicy and deletes it. Returns an error if one occurs.
func (c *FakeKubeArmorHostPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(kubearmorhostpoliciesResource, name), &securitykubearmorcomv1.KubeArmorHostPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeArmorHostPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kubearmorhostpoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &securitykubearmorcomv1.KubeArmorHostPolicyList{})
	return err
}

// Patch applies the patch and returns the patched kubeArmorHostPolicy.
func (c *FakeKubeArmorHostPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *securitykubearmorcomv1.KubeArmorHostPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubearmorhostpoliciesResource, name, pt, data, subresources...), &securitykubearmorcomv1.KubeArmorHostPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*securitykubearmorcomv1.KubeArmorHostPolicy), err
}
