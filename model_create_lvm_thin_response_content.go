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

// CreateLVMThinResponseContent struct for CreateLVMThinResponseContent
type CreateLVMThinResponseContent struct {
	Data string `json:"data"`
}

// NewCreateLVMThinResponseContent instantiates a new CreateLVMThinResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLVMThinResponseContent(data string) *CreateLVMThinResponseContent {
	this := CreateLVMThinResponseContent{}
	this.Data = data
	return &this
}

// NewCreateLVMThinResponseContentWithDefaults instantiates a new CreateLVMThinResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLVMThinResponseContentWithDefaults() *CreateLVMThinResponseContent {
	this := CreateLVMThinResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *CreateLVMThinResponseContent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateLVMThinResponseContent) GetDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateLVMThinResponseContent) SetData(v string) {
	o.Data = v
}

func (o CreateLVMThinResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLVMThinResponseContent struct {
	value *CreateLVMThinResponseContent
	isSet bool
}

func (v NullableCreateLVMThinResponseContent) Get() *CreateLVMThinResponseContent {
	return v.value
}

func (v *NullableCreateLVMThinResponseContent) Set(val *CreateLVMThinResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLVMThinResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLVMThinResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLVMThinResponseContent(val *CreateLVMThinResponseContent) *NullableCreateLVMThinResponseContent {
	return &NullableCreateLVMThinResponseContent{value: val, isSet: true}
}

func (v NullableCreateLVMThinResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLVMThinResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


