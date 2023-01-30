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

// DeleteDirectoryResponseContent struct for DeleteDirectoryResponseContent
type DeleteDirectoryResponseContent struct {
	Data string `json:"data"`
}

// NewDeleteDirectoryResponseContent instantiates a new DeleteDirectoryResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDirectoryResponseContent(data string) *DeleteDirectoryResponseContent {
	this := DeleteDirectoryResponseContent{}
	this.Data = data
	return &this
}

// NewDeleteDirectoryResponseContentWithDefaults instantiates a new DeleteDirectoryResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDirectoryResponseContentWithDefaults() *DeleteDirectoryResponseContent {
	this := DeleteDirectoryResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *DeleteDirectoryResponseContent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteDirectoryResponseContent) GetDataOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteDirectoryResponseContent) SetData(v string) {
	o.Data = v
}

func (o DeleteDirectoryResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteDirectoryResponseContent struct {
	value *DeleteDirectoryResponseContent
	isSet bool
}

func (v NullableDeleteDirectoryResponseContent) Get() *DeleteDirectoryResponseContent {
	return v.value
}

func (v *NullableDeleteDirectoryResponseContent) Set(val *DeleteDirectoryResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDirectoryResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDirectoryResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDirectoryResponseContent(val *DeleteDirectoryResponseContent) *NullableDeleteDirectoryResponseContent {
	return &NullableDeleteDirectoryResponseContent{value: val, isSet: true}
}

func (v NullableDeleteDirectoryResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDirectoryResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


