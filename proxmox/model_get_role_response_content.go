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

// checks if the GetRoleResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRoleResponseContent{}

// GetRoleResponseContent struct for GetRoleResponseContent
type GetRoleResponseContent struct {
	Data RolePermissionSummary `json:"data"`
}

// NewGetRoleResponseContent instantiates a new GetRoleResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRoleResponseContent(data RolePermissionSummary) *GetRoleResponseContent {
	this := GetRoleResponseContent{}
	this.Data = data
	return &this
}

// NewGetRoleResponseContentWithDefaults instantiates a new GetRoleResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRoleResponseContentWithDefaults() *GetRoleResponseContent {
	this := GetRoleResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *GetRoleResponseContent) GetData() RolePermissionSummary {
	if o == nil {
		var ret RolePermissionSummary
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetRoleResponseContent) GetDataOk() (*RolePermissionSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetRoleResponseContent) SetData(v RolePermissionSummary) {
	o.Data = v
}

func (o GetRoleResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRoleResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetRoleResponseContent struct {
	value *GetRoleResponseContent
	isSet bool
}

func (v NullableGetRoleResponseContent) Get() *GetRoleResponseContent {
	return v.value
}

func (v *NullableGetRoleResponseContent) Set(val *GetRoleResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRoleResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRoleResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRoleResponseContent(val *GetRoleResponseContent) *NullableGetRoleResponseContent {
	return &NullableGetRoleResponseContent{value: val, isSet: true}
}

func (v NullableGetRoleResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRoleResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


