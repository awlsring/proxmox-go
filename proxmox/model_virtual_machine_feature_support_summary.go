/*
Proxmox

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2023-01-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package proxmox

import (
	"encoding/json"
)

// checks if the VirtualMachineFeatureSupportSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualMachineFeatureSupportSummary{}

// VirtualMachineFeatureSupportSummary struct for VirtualMachineFeatureSupportSummary
type VirtualMachineFeatureSupportSummary struct {
	// An integer used to represent a boolean. 0 is false, 1 is true.
	HasFeature float32 `json:"hasFeature"`
	Nodes []string `json:"nodes"`
}

// NewVirtualMachineFeatureSupportSummary instantiates a new VirtualMachineFeatureSupportSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMachineFeatureSupportSummary(hasFeature float32, nodes []string) *VirtualMachineFeatureSupportSummary {
	this := VirtualMachineFeatureSupportSummary{}
	this.HasFeature = hasFeature
	this.Nodes = nodes
	return &this
}

// NewVirtualMachineFeatureSupportSummaryWithDefaults instantiates a new VirtualMachineFeatureSupportSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMachineFeatureSupportSummaryWithDefaults() *VirtualMachineFeatureSupportSummary {
	this := VirtualMachineFeatureSupportSummary{}
	return &this
}

// GetHasFeature returns the HasFeature field value
func (o *VirtualMachineFeatureSupportSummary) GetHasFeature() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HasFeature
}

// GetHasFeatureOk returns a tuple with the HasFeature field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineFeatureSupportSummary) GetHasFeatureOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasFeature, true
}

// SetHasFeature sets field value
func (o *VirtualMachineFeatureSupportSummary) SetHasFeature(v float32) {
	o.HasFeature = v
}

// GetNodes returns the Nodes field value
func (o *VirtualMachineFeatureSupportSummary) GetNodes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineFeatureSupportSummary) GetNodesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *VirtualMachineFeatureSupportSummary) SetNodes(v []string) {
	o.Nodes = v
}

func (o VirtualMachineFeatureSupportSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualMachineFeatureSupportSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hasFeature"] = o.HasFeature
	toSerialize["nodes"] = o.Nodes
	return toSerialize, nil
}

type NullableVirtualMachineFeatureSupportSummary struct {
	value *VirtualMachineFeatureSupportSummary
	isSet bool
}

func (v NullableVirtualMachineFeatureSupportSummary) Get() *VirtualMachineFeatureSupportSummary {
	return v.value
}

func (v *NullableVirtualMachineFeatureSupportSummary) Set(val *VirtualMachineFeatureSupportSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineFeatureSupportSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineFeatureSupportSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineFeatureSupportSummary(val *VirtualMachineFeatureSupportSummary) *NullableVirtualMachineFeatureSupportSummary {
	return &NullableVirtualMachineFeatureSupportSummary{value: val, isSet: true}
}

func (v NullableVirtualMachineFeatureSupportSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineFeatureSupportSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


