/*
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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha3

import (
	v1alpha3 "github.com/deislabs/smi-sdk-go/pkg/apis/split/v1alpha3"
	"github.com/deislabs/smi-sdk-go/pkg/gen/client/split/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type SplitV1alpha3Interface interface {
	RESTClient() rest.Interface
	TrafficSplitsGetter
}

// SplitV1alpha3Client is used to interact with features provided by the split.smi-spec.io group.
type SplitV1alpha3Client struct {
	restClient rest.Interface
}

func (c *SplitV1alpha3Client) TrafficSplits(namespace string) TrafficSplitInterface {
	return newTrafficSplits(c, namespace)
}

// NewForConfig creates a new SplitV1alpha3Client for the given config.
func NewForConfig(c *rest.Config) (*SplitV1alpha3Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SplitV1alpha3Client{client}, nil
}

// NewForConfigOrDie creates a new SplitV1alpha3Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SplitV1alpha3Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SplitV1alpha3Client for the given RESTClient.
func New(c rest.Interface) *SplitV1alpha3Client {
	return &SplitV1alpha3Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha3.SchemeGroupVersion
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
func (c *SplitV1alpha3Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
