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

// checks if the AddCorosyncNodeResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCorosyncNodeResponseContent{}

// AddCorosyncNodeResponseContent struct for AddCorosyncNodeResponseContent
type AddCorosyncNodeResponseContent struct {
	Data CorosyncSettings `json:"data"`
}

// NewAddCorosyncNodeResponseContent instantiates a new AddCorosyncNodeResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCorosyncNodeResponseContent(data CorosyncSettings) *AddCorosyncNodeResponseContent {
	this := AddCorosyncNodeResponseContent{}
	this.Data = data
	return &this
}

// NewAddCorosyncNodeResponseContentWithDefaults instantiates a new AddCorosyncNodeResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCorosyncNodeResponseContentWithDefaults() *AddCorosyncNodeResponseContent {
	this := AddCorosyncNodeResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *AddCorosyncNodeResponseContent) GetData() CorosyncSettings {
	if o == nil {
		var ret CorosyncSettings
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AddCorosyncNodeResponseContent) GetDataOk() (*CorosyncSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AddCorosyncNodeResponseContent) SetData(v CorosyncSettings) {
	o.Data = v
}

func (o AddCorosyncNodeResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCorosyncNodeResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAddCorosyncNodeResponseContent struct {
	value *AddCorosyncNodeResponseContent
	isSet bool
}

func (v NullableAddCorosyncNodeResponseContent) Get() *AddCorosyncNodeResponseContent {
	return v.value
}

func (v *NullableAddCorosyncNodeResponseContent) Set(val *AddCorosyncNodeResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCorosyncNodeResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCorosyncNodeResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCorosyncNodeResponseContent(val *AddCorosyncNodeResponseContent) *NullableAddCorosyncNodeResponseContent {
	return &NullableAddCorosyncNodeResponseContent{value: val, isSet: true}
}

func (v NullableAddCorosyncNodeResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCorosyncNodeResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


