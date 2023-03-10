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

// checks if the ApplyNetworkInterfaceConfigurationResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplyNetworkInterfaceConfigurationResponseContent{}

// ApplyNetworkInterfaceConfigurationResponseContent struct for ApplyNetworkInterfaceConfigurationResponseContent
type ApplyNetworkInterfaceConfigurationResponseContent struct {
	Data string `json:"data"`
}

// NewApplyNetworkInterfaceConfigurationResponseContent instantiates a new ApplyNetworkInterfaceConfigurationResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyNetworkInterfaceConfigurationResponseContent(data string) *ApplyNetworkInterfaceConfigurationResponseContent {
	this := ApplyNetworkInterfaceConfigurationResponseContent{}
	this.Data = data
	return &this
}

// NewApplyNetworkInterfaceConfigurationResponseContentWithDefaults instantiates a new ApplyNetworkInterfaceConfigurationResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyNetworkInterfaceConfigurationResponseContentWithDefaults() *ApplyNetworkInterfaceConfigurationResponseContent {
	this := ApplyNetworkInterfaceConfigurationResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *ApplyNetworkInterfaceConfigurationResponseContent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ApplyNetworkInterfaceConfigurationResponseContent) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ApplyNetworkInterfaceConfigurationResponseContent) SetData(v string) {
	o.Data = v
}

func (o ApplyNetworkInterfaceConfigurationResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyNetworkInterfaceConfigurationResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableApplyNetworkInterfaceConfigurationResponseContent struct {
	value *ApplyNetworkInterfaceConfigurationResponseContent
	isSet bool
}

func (v NullableApplyNetworkInterfaceConfigurationResponseContent) Get() *ApplyNetworkInterfaceConfigurationResponseContent {
	return v.value
}

func (v *NullableApplyNetworkInterfaceConfigurationResponseContent) Set(val *ApplyNetworkInterfaceConfigurationResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyNetworkInterfaceConfigurationResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyNetworkInterfaceConfigurationResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyNetworkInterfaceConfigurationResponseContent(val *ApplyNetworkInterfaceConfigurationResponseContent) *NullableApplyNetworkInterfaceConfigurationResponseContent {
	return &NullableApplyNetworkInterfaceConfigurationResponseContent{value: val, isSet: true}
}

func (v NullableApplyNetworkInterfaceConfigurationResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyNetworkInterfaceConfigurationResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


