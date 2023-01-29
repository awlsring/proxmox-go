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

// InitializeGPTResponseContent struct for InitializeGPTResponseContent
type InitializeGPTResponseContent struct {
	Data string `json:"data"`
}

// NewInitializeGPTResponseContent instantiates a new InitializeGPTResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitializeGPTResponseContent(data string) *InitializeGPTResponseContent {
	this := InitializeGPTResponseContent{}
	this.Data = data
	return &this
}

// NewInitializeGPTResponseContentWithDefaults instantiates a new InitializeGPTResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitializeGPTResponseContentWithDefaults() *InitializeGPTResponseContent {
	this := InitializeGPTResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *InitializeGPTResponseContent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InitializeGPTResponseContent) GetDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InitializeGPTResponseContent) SetData(v string) {
	o.Data = v
}

func (o InitializeGPTResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInitializeGPTResponseContent struct {
	value *InitializeGPTResponseContent
	isSet bool
}

func (v NullableInitializeGPTResponseContent) Get() *InitializeGPTResponseContent {
	return v.value
}

func (v *NullableInitializeGPTResponseContent) Set(val *InitializeGPTResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInitializeGPTResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInitializeGPTResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitializeGPTResponseContent(val *InitializeGPTResponseContent) *NullableInitializeGPTResponseContent {
	return &NullableInitializeGPTResponseContent{value: val, isSet: true}
}

func (v NullableInitializeGPTResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitializeGPTResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


