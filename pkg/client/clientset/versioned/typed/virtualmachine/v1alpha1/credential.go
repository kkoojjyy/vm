/*
Copyright 2018 Rancher Labs, Inc.

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
package v1alpha1

import (
	v1alpha1 "github.com/rancher/vm/pkg/apis/ranchervm/v1alpha1"
	scheme "github.com/rancher/vm/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CredentialsGetter has a method to return a CredentialInterface.
// A group's client should implement this interface.
type CredentialsGetter interface {
	Credentials() CredentialInterface
}

// CredentialInterface has methods to work with Credential resources.
type CredentialInterface interface {
	Create(*v1alpha1.Credential) (*v1alpha1.Credential, error)
	Update(*v1alpha1.Credential) (*v1alpha1.Credential, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Credential, error)
	List(opts v1.ListOptions) (*v1alpha1.CredentialList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Credential, err error)
	CredentialExpansion
}

// credentials implements CredentialInterface
type credentials struct {
	client rest.Interface
}

// newCredentials returns a Credentials
func newCredentials(c *VirtualmachineV1alpha1Client) *credentials {
	return &credentials{
		client: c.RESTClient(),
	}
}

// Get takes name of the credential, and returns the corresponding credential object, and an error if there is any.
func (c *credentials) Get(name string, options v1.GetOptions) (result *v1alpha1.Credential, err error) {
	result = &v1alpha1.Credential{}
	err = c.client.Get().
		Resource("credentials").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Credentials that match those selectors.
func (c *credentials) List(opts v1.ListOptions) (result *v1alpha1.CredentialList, err error) {
	result = &v1alpha1.CredentialList{}
	err = c.client.Get().
		Resource("credentials").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested credentials.
func (c *credentials) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("credentials").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a credential and creates it.  Returns the server's representation of the credential, and an error, if there is any.
func (c *credentials) Create(credential *v1alpha1.Credential) (result *v1alpha1.Credential, err error) {
	result = &v1alpha1.Credential{}
	err = c.client.Post().
		Resource("credentials").
		Body(credential).
		Do().
		Into(result)
	return
}

// Update takes the representation of a credential and updates it. Returns the server's representation of the credential, and an error, if there is any.
func (c *credentials) Update(credential *v1alpha1.Credential) (result *v1alpha1.Credential, err error) {
	result = &v1alpha1.Credential{}
	err = c.client.Put().
		Resource("credentials").
		Name(credential.Name).
		Body(credential).
		Do().
		Into(result)
	return
}

// Delete takes name of the credential and deletes it. Returns an error if one occurs.
func (c *credentials) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("credentials").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *credentials) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("credentials").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched credential.
func (c *credentials) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Credential, err error) {
	result = &v1alpha1.Credential{}
	err = c.client.Patch(pt).
		Resource("credentials").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
