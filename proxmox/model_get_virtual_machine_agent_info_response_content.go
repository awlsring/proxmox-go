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

// checks if the GetVirtualMachineAgentInfoResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVirtualMachineAgentInfoResponseContent{}

// GetVirtualMachineAgentInfoResponseContent struct for GetVirtualMachineAgentInfoResponseContent
type GetVirtualMachineAgentInfoResponseContent struct {
	Data GetInfoResult `json:"data"`
}

// NewGetVirtualMachineAgentInfoResponseContent instantiates a new GetVirtualMachineAgentInfoResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVirtualMachineAgentInfoResponseContent(data GetInfoResult) *GetVirtualMachineAgentInfoResponseContent {
	this := GetVirtualMachineAgentInfoResponseContent{}
	this.Data = data
	return &this
}

// NewGetVirtualMachineAgentInfoResponseContentWithDefaults instantiates a new GetVirtualMachineAgentInfoResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVirtualMachineAgentInfoResponseContentWithDefaults() *GetVirtualMachineAgentInfoResponseContent {
	this := GetVirtualMachineAgentInfoResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *GetVirtualMachineAgentInfoResponseContent) GetData() GetInfoResult {
	if o == nil {
		var ret GetInfoResult
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetVirtualMachineAgentInfoResponseContent) GetDataOk() (*GetInfoResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetVirtualMachineAgentInfoResponseContent) SetData(v GetInfoResult) {
	o.Data = v
}

func (o GetVirtualMachineAgentInfoResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVirtualMachineAgentInfoResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetVirtualMachineAgentInfoResponseContent struct {
	value *GetVirtualMachineAgentInfoResponseContent
	isSet bool
}

func (v NullableGetVirtualMachineAgentInfoResponseContent) Get() *GetVirtualMachineAgentInfoResponseContent {
	return v.value
}

func (v *NullableGetVirtualMachineAgentInfoResponseContent) Set(val *GetVirtualMachineAgentInfoResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVirtualMachineAgentInfoResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVirtualMachineAgentInfoResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVirtualMachineAgentInfoResponseContent(val *GetVirtualMachineAgentInfoResponseContent) *NullableGetVirtualMachineAgentInfoResponseContent {
	return &NullableGetVirtualMachineAgentInfoResponseContent{value: val, isSet: true}
}

func (v NullableGetVirtualMachineAgentInfoResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVirtualMachineAgentInfoResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

