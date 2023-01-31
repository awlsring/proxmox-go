/*
Proxmox

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2023-01-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetVirtualMachineCloudInitResponseContent struct for GetVirtualMachineCloudInitResponseContent
type GetVirtualMachineCloudInitResponseContent struct {
	Data string `json:"data"`
}

// NewGetVirtualMachineCloudInitResponseContent instantiates a new GetVirtualMachineCloudInitResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVirtualMachineCloudInitResponseContent(data string) *GetVirtualMachineCloudInitResponseContent {
	this := GetVirtualMachineCloudInitResponseContent{}
	this.Data = data
	return &this
}

// NewGetVirtualMachineCloudInitResponseContentWithDefaults instantiates a new GetVirtualMachineCloudInitResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVirtualMachineCloudInitResponseContentWithDefaults() *GetVirtualMachineCloudInitResponseContent {
	this := GetVirtualMachineCloudInitResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *GetVirtualMachineCloudInitResponseContent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetVirtualMachineCloudInitResponseContent) GetDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetVirtualMachineCloudInitResponseContent) SetData(v string) {
	o.Data = v
}

func (o GetVirtualMachineCloudInitResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetVirtualMachineCloudInitResponseContent struct {
	value *GetVirtualMachineCloudInitResponseContent
	isSet bool
}

func (v NullableGetVirtualMachineCloudInitResponseContent) Get() *GetVirtualMachineCloudInitResponseContent {
	return v.value
}

func (v *NullableGetVirtualMachineCloudInitResponseContent) Set(val *GetVirtualMachineCloudInitResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVirtualMachineCloudInitResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVirtualMachineCloudInitResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVirtualMachineCloudInitResponseContent(val *GetVirtualMachineCloudInitResponseContent) *NullableGetVirtualMachineCloudInitResponseContent {
	return &NullableGetVirtualMachineCloudInitResponseContent{value: val, isSet: true}
}

func (v NullableGetVirtualMachineCloudInitResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVirtualMachineCloudInitResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


