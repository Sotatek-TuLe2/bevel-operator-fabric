/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	"github.com/kfsoftware/hlf-operator/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type HlfV1alpha1Interface interface {
	RESTClient() rest.Interface
	FabricCAsGetter
	FabricChaincodesGetter
	FabricExplorersGetter
	FabricFollowerChannelsGetter
	FabricIdentitiesGetter
	FabricMainChannelsGetter
	FabricNetworkConfigsGetter
	FabricOperationsConsolesGetter
	FabricOperatorAPIsGetter
	FabricOperatorUIsGetter
	FabricOrdererNodesGetter
	FabricOrderingServicesGetter
	FabricPeersGetter
}

// HlfV1alpha1Client is used to interact with features provided by the hlf.kungfusoftware.es group.
type HlfV1alpha1Client struct {
	restClient rest.Interface
}

func (c *HlfV1alpha1Client) FabricCAs(namespace string) FabricCAInterface {
	return newFabricCAs(c, namespace)
}

func (c *HlfV1alpha1Client) FabricChaincodes(namespace string) FabricChaincodeInterface {
	return newFabricChaincodes(c, namespace)
}

func (c *HlfV1alpha1Client) FabricExplorers(namespace string) FabricExplorerInterface {
	return newFabricExplorers(c, namespace)
}

func (c *HlfV1alpha1Client) FabricFollowerChannels() FabricFollowerChannelInterface {
	return newFabricFollowerChannels(c)
}

func (c *HlfV1alpha1Client) FabricIdentities(namespace string) FabricIdentityInterface {
	return newFabricIdentities(c, namespace)
}

func (c *HlfV1alpha1Client) FabricMainChannels() FabricMainChannelInterface {
	return newFabricMainChannels(c)
}

func (c *HlfV1alpha1Client) FabricNetworkConfigs(namespace string) FabricNetworkConfigInterface {
	return newFabricNetworkConfigs(c, namespace)
}

func (c *HlfV1alpha1Client) FabricOperationsConsoles(namespace string) FabricOperationsConsoleInterface {
	return newFabricOperationsConsoles(c, namespace)
}

func (c *HlfV1alpha1Client) FabricOperatorAPIs(namespace string) FabricOperatorAPIInterface {
	return newFabricOperatorAPIs(c, namespace)
}

func (c *HlfV1alpha1Client) FabricOperatorUIs(namespace string) FabricOperatorUIInterface {
	return newFabricOperatorUIs(c, namespace)
}

func (c *HlfV1alpha1Client) FabricOrdererNodes(namespace string) FabricOrdererNodeInterface {
	return newFabricOrdererNodes(c, namespace)
}

func (c *HlfV1alpha1Client) FabricOrderingServices(namespace string) FabricOrderingServiceInterface {
	return newFabricOrderingServices(c, namespace)
}

func (c *HlfV1alpha1Client) FabricPeers(namespace string) FabricPeerInterface {
	return newFabricPeers(c, namespace)
}

// NewForConfig creates a new HlfV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*HlfV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new HlfV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*HlfV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &HlfV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new HlfV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *HlfV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new HlfV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *HlfV1alpha1Client {
	return &HlfV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *HlfV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
