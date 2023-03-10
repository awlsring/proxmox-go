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

// checks if the GetRealmResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRealmResponseContent{}

// GetRealmResponseContent struct for GetRealmResponseContent
type GetRealmResponseContent struct {
	Data RealmConfigurationSummary `json:"data"`
}

// NewGetRealmResponseContent instantiates a new GetRealmResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRealmResponseContent(data RealmConfigurationSummary) *GetRealmResponseContent {
	this := GetRealmResponseContent{}
	this.Data = data
	return &this
}

// NewGetRealmResponseContentWithDefaults instantiates a new GetRealmResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRealmResponseContentWithDefaults() *GetRealmResponseContent {
	this := GetRealmResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *GetRealmResponseContent) GetData() RealmConfigurationSummary {
	if o == nil {
		var ret RealmConfigurationSummary
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetRealmResponseContent) GetDataOk() (*RealmConfigurationSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetRealmResponseContent) SetData(v RealmConfigurationSummary) {
	o.Data = v
}

func (o GetRealmResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRealmResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetRealmResponseContent struct {
	value *GetRealmResponseContent
	isSet bool
}

func (v NullableGetRealmResponseContent) Get() *GetRealmResponseContent {
	return v.value
}

func (v *NullableGetRealmResponseContent) Set(val *GetRealmResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRealmResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRealmResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRealmResponseContent(val *GetRealmResponseContent) *NullableGetRealmResponseContent {
	return &NullableGetRealmResponseContent{value: val, isSet: true}
}

func (v NullableGetRealmResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRealmResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


