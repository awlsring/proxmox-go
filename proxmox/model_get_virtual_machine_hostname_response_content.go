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

// checks if the GetVirtualMachineHostnameResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVirtualMachineHostnameResponseContent{}

// GetVirtualMachineHostnameResponseContent struct for GetVirtualMachineHostnameResponseContent
type GetVirtualMachineHostnameResponseContent struct {
	Data GetHostnameResult `json:"data"`
}

// NewGetVirtualMachineHostnameResponseContent instantiates a new GetVirtualMachineHostnameResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVirtualMachineHostnameResponseContent(data GetHostnameResult) *GetVirtualMachineHostnameResponseContent {
	this := GetVirtualMachineHostnameResponseContent{}
	this.Data = data
	return &this
}

// NewGetVirtualMachineHostnameResponseContentWithDefaults instantiates a new GetVirtualMachineHostnameResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVirtualMachineHostnameResponseContentWithDefaults() *GetVirtualMachineHostnameResponseContent {
	this := GetVirtualMachineHostnameResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *GetVirtualMachineHostnameResponseContent) GetData() GetHostnameResult {
	if o == nil {
		var ret GetHostnameResult
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetVirtualMachineHostnameResponseContent) GetDataOk() (*GetHostnameResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetVirtualMachineHostnameResponseContent) SetData(v GetHostnameResult) {
	o.Data = v
}

func (o GetVirtualMachineHostnameResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVirtualMachineHostnameResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetVirtualMachineHostnameResponseContent struct {
	value *GetVirtualMachineHostnameResponseContent
	isSet bool
}

func (v NullableGetVirtualMachineHostnameResponseContent) Get() *GetVirtualMachineHostnameResponseContent {
	return v.value
}

func (v *NullableGetVirtualMachineHostnameResponseContent) Set(val *GetVirtualMachineHostnameResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVirtualMachineHostnameResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVirtualMachineHostnameResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVirtualMachineHostnameResponseContent(val *GetVirtualMachineHostnameResponseContent) *NullableGetVirtualMachineHostnameResponseContent {
	return &NullableGetVirtualMachineHostnameResponseContent{value: val, isSet: true}
}

func (v NullableGetVirtualMachineHostnameResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVirtualMachineHostnameResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


