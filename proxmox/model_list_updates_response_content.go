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

// ListUpdatesResponseContent struct for ListUpdatesResponseContent
type ListUpdatesResponseContent struct {
	Data []UpdateSummary `json:"data"`
}

// NewListUpdatesResponseContent instantiates a new ListUpdatesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUpdatesResponseContent(data []UpdateSummary) *ListUpdatesResponseContent {
	this := ListUpdatesResponseContent{}
	this.Data = data
	return &this
}

// NewListUpdatesResponseContentWithDefaults instantiates a new ListUpdatesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUpdatesResponseContentWithDefaults() *ListUpdatesResponseContent {
	this := ListUpdatesResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *ListUpdatesResponseContent) GetData() []UpdateSummary {
	if o == nil {
		var ret []UpdateSummary
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListUpdatesResponseContent) GetDataOk() ([]UpdateSummary, bool) {
	if o == nil {
    return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListUpdatesResponseContent) SetData(v []UpdateSummary) {
	o.Data = v
}

func (o ListUpdatesResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListUpdatesResponseContent struct {
	value *ListUpdatesResponseContent
	isSet bool
}

func (v NullableListUpdatesResponseContent) Get() *ListUpdatesResponseContent {
	return v.value
}

func (v *NullableListUpdatesResponseContent) Set(val *ListUpdatesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListUpdatesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListUpdatesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUpdatesResponseContent(val *ListUpdatesResponseContent) *NullableListUpdatesResponseContent {
	return &NullableListUpdatesResponseContent{value: val, isSet: true}
}

func (v NullableListUpdatesResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUpdatesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

