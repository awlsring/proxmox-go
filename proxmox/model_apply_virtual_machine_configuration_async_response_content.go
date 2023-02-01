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

// ApplyVirtualMachineConfigurationAsyncResponseContent struct for ApplyVirtualMachineConfigurationAsyncResponseContent
type ApplyVirtualMachineConfigurationAsyncResponseContent struct {
	Data string `json:"data"`
}

// NewApplyVirtualMachineConfigurationAsyncResponseContent instantiates a new ApplyVirtualMachineConfigurationAsyncResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyVirtualMachineConfigurationAsyncResponseContent(data string) *ApplyVirtualMachineConfigurationAsyncResponseContent {
	this := ApplyVirtualMachineConfigurationAsyncResponseContent{}
	this.Data = data
	return &this
}

// NewApplyVirtualMachineConfigurationAsyncResponseContentWithDefaults instantiates a new ApplyVirtualMachineConfigurationAsyncResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyVirtualMachineConfigurationAsyncResponseContentWithDefaults() *ApplyVirtualMachineConfigurationAsyncResponseContent {
	this := ApplyVirtualMachineConfigurationAsyncResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *ApplyVirtualMachineConfigurationAsyncResponseContent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ApplyVirtualMachineConfigurationAsyncResponseContent) GetDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ApplyVirtualMachineConfigurationAsyncResponseContent) SetData(v string) {
	o.Data = v
}

func (o ApplyVirtualMachineConfigurationAsyncResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableApplyVirtualMachineConfigurationAsyncResponseContent struct {
	value *ApplyVirtualMachineConfigurationAsyncResponseContent
	isSet bool
}

func (v NullableApplyVirtualMachineConfigurationAsyncResponseContent) Get() *ApplyVirtualMachineConfigurationAsyncResponseContent {
	return v.value
}

func (v *NullableApplyVirtualMachineConfigurationAsyncResponseContent) Set(val *ApplyVirtualMachineConfigurationAsyncResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyVirtualMachineConfigurationAsyncResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyVirtualMachineConfigurationAsyncResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyVirtualMachineConfigurationAsyncResponseContent(val *ApplyVirtualMachineConfigurationAsyncResponseContent) *NullableApplyVirtualMachineConfigurationAsyncResponseContent {
	return &NullableApplyVirtualMachineConfigurationAsyncResponseContent{value: val, isSet: true}
}

func (v NullableApplyVirtualMachineConfigurationAsyncResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyVirtualMachineConfigurationAsyncResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

