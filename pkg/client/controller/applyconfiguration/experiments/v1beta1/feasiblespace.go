/*
Copyright 2022 The Kubeflow Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/kubeflow/katib/pkg/apis/controller/experiments/v1beta1"
)

// FeasibleSpaceApplyConfiguration represents a declarative configuration of the FeasibleSpace type for use
// with apply.
type FeasibleSpaceApplyConfiguration struct {
	Max          *string               `json:"max,omitempty"`
	Min          *string               `json:"min,omitempty"`
	List         []string              `json:"list,omitempty"`
	Step         *string               `json:"step,omitempty"`
	Distribution *v1beta1.Distribution `json:"distribution,omitempty"`
}

// FeasibleSpaceApplyConfiguration constructs a declarative configuration of the FeasibleSpace type for use with
// apply.
func FeasibleSpace() *FeasibleSpaceApplyConfiguration {
	return &FeasibleSpaceApplyConfiguration{}
}

// WithMax sets the Max field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Max field is set to the value of the last call.
func (b *FeasibleSpaceApplyConfiguration) WithMax(value string) *FeasibleSpaceApplyConfiguration {
	b.Max = &value
	return b
}

// WithMin sets the Min field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Min field is set to the value of the last call.
func (b *FeasibleSpaceApplyConfiguration) WithMin(value string) *FeasibleSpaceApplyConfiguration {
	b.Min = &value
	return b
}

// WithList adds the given value to the List field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the List field.
func (b *FeasibleSpaceApplyConfiguration) WithList(values ...string) *FeasibleSpaceApplyConfiguration {
	for i := range values {
		b.List = append(b.List, values[i])
	}
	return b
}

// WithStep sets the Step field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Step field is set to the value of the last call.
func (b *FeasibleSpaceApplyConfiguration) WithStep(value string) *FeasibleSpaceApplyConfiguration {
	b.Step = &value
	return b
}

// WithDistribution sets the Distribution field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Distribution field is set to the value of the last call.
func (b *FeasibleSpaceApplyConfiguration) WithDistribution(value v1beta1.Distribution) *FeasibleSpaceApplyConfiguration {
	b.Distribution = &value
	return b
}
