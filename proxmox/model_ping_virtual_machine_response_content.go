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

// checks if the PingVirtualMachineResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PingVirtualMachineResponseContent{}

// PingVirtualMachineResponseContent struct for PingVirtualMachineResponseContent
type PingVirtualMachineResponseContent struct {
	Data map[string]interface{} `json:"data"`
}

// NewPingVirtualMachineResponseContent instantiates a new PingVirtualMachineResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingVirtualMachineResponseContent(data map[string]interface{}) *PingVirtualMachineResponseContent {
	this := PingVirtualMachineResponseContent{}
	this.Data = data
	return &this
}

// NewPingVirtualMachineResponseContentWithDefaults instantiates a new PingVirtualMachineResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingVirtualMachineResponseContentWithDefaults() *PingVirtualMachineResponseContent {
	this := PingVirtualMachineResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *PingVirtualMachineResponseContent) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PingVirtualMachineResponseContent) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PingVirtualMachineResponseContent) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o PingVirtualMachineResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PingVirtualMachineResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePingVirtualMachineResponseContent struct {
	value *PingVirtualMachineResponseContent
	isSet bool
}

func (v NullablePingVirtualMachineResponseContent) Get() *PingVirtualMachineResponseContent {
	return v.value
}

func (v *NullablePingVirtualMachineResponseContent) Set(val *PingVirtualMachineResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullablePingVirtualMachineResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullablePingVirtualMachineResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingVirtualMachineResponseContent(val *PingVirtualMachineResponseContent) *NullablePingVirtualMachineResponseContent {
	return &NullablePingVirtualMachineResponseContent{value: val, isSet: true}
}

func (v NullablePingVirtualMachineResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingVirtualMachineResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


