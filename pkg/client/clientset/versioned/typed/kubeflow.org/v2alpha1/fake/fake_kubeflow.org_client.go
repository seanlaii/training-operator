// Copyright 2024 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2alpha1 "github.com/kubeflow/training-operator/pkg/client/clientset/versioned/typed/kubeflow.org/v2alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKubeflowV2alpha1 struct {
	*testing.Fake
}

func (c *FakeKubeflowV2alpha1) ClusterTrainingRuntimes() v2alpha1.ClusterTrainingRuntimeInterface {
	return &FakeClusterTrainingRuntimes{c}
}

func (c *FakeKubeflowV2alpha1) TrainJobs(namespace string) v2alpha1.TrainJobInterface {
	return &FakeTrainJobs{c, namespace}
}

func (c *FakeKubeflowV2alpha1) TrainingRuntimes(namespace string) v2alpha1.TrainingRuntimeInterface {
	return &FakeTrainingRuntimes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKubeflowV2alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
