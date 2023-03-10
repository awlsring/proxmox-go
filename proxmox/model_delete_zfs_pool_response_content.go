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

// checks if the DeleteZFSPoolResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteZFSPoolResponseContent{}

// DeleteZFSPoolResponseContent struct for DeleteZFSPoolResponseContent
type DeleteZFSPoolResponseContent struct {
	Data string `json:"data"`
}

// NewDeleteZFSPoolResponseContent instantiates a new DeleteZFSPoolResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteZFSPoolResponseContent(data string) *DeleteZFSPoolResponseContent {
	this := DeleteZFSPoolResponseContent{}
	this.Data = data
	return &this
}

// NewDeleteZFSPoolResponseContentWithDefaults instantiates a new DeleteZFSPoolResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteZFSPoolResponseContentWithDefaults() *DeleteZFSPoolResponseContent {
	this := DeleteZFSPoolResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *DeleteZFSPoolResponseContent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteZFSPoolResponseContent) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteZFSPoolResponseContent) SetData(v string) {
	o.Data = v
}

func (o DeleteZFSPoolResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteZFSPoolResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableDeleteZFSPoolResponseContent struct {
	value *DeleteZFSPoolResponseContent
	isSet bool
}

func (v NullableDeleteZFSPoolResponseContent) Get() *DeleteZFSPoolResponseContent {
	return v.value
}

func (v *NullableDeleteZFSPoolResponseContent) Set(val *DeleteZFSPoolResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteZFSPoolResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteZFSPoolResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteZFSPoolResponseContent(val *DeleteZFSPoolResponseContent) *NullableDeleteZFSPoolResponseContent {
	return &NullableDeleteZFSPoolResponseContent{value: val, isSet: true}
}

func (v NullableDeleteZFSPoolResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteZFSPoolResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


