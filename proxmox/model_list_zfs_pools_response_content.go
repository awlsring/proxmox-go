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

// ListZFSPoolsResponseContent struct for ListZFSPoolsResponseContent
type ListZFSPoolsResponseContent struct {
	Data []ZFSPoolSummary `json:"data"`
}

// NewListZFSPoolsResponseContent instantiates a new ListZFSPoolsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListZFSPoolsResponseContent(data []ZFSPoolSummary) *ListZFSPoolsResponseContent {
	this := ListZFSPoolsResponseContent{}
	this.Data = data
	return &this
}

// NewListZFSPoolsResponseContentWithDefaults instantiates a new ListZFSPoolsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListZFSPoolsResponseContentWithDefaults() *ListZFSPoolsResponseContent {
	this := ListZFSPoolsResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *ListZFSPoolsResponseContent) GetData() []ZFSPoolSummary {
	if o == nil {
		var ret []ZFSPoolSummary
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListZFSPoolsResponseContent) GetDataOk() ([]ZFSPoolSummary, bool) {
	if o == nil {
    return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListZFSPoolsResponseContent) SetData(v []ZFSPoolSummary) {
	o.Data = v
}

func (o ListZFSPoolsResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListZFSPoolsResponseContent struct {
	value *ListZFSPoolsResponseContent
	isSet bool
}

func (v NullableListZFSPoolsResponseContent) Get() *ListZFSPoolsResponseContent {
	return v.value
}

func (v *NullableListZFSPoolsResponseContent) Set(val *ListZFSPoolsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListZFSPoolsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListZFSPoolsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListZFSPoolsResponseContent(val *ListZFSPoolsResponseContent) *NullableListZFSPoolsResponseContent {
	return &NullableListZFSPoolsResponseContent{value: val, isSet: true}
}

func (v NullableListZFSPoolsResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListZFSPoolsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

