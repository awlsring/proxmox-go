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

// checks if the ListPendingVirtualMachineConfigurationChangesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPendingVirtualMachineConfigurationChangesResponseContent{}

// ListPendingVirtualMachineConfigurationChangesResponseContent struct for ListPendingVirtualMachineConfigurationChangesResponseContent
type ListPendingVirtualMachineConfigurationChangesResponseContent struct {
	Data string `json:"data"`
}

// NewListPendingVirtualMachineConfigurationChangesResponseContent instantiates a new ListPendingVirtualMachineConfigurationChangesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPendingVirtualMachineConfigurationChangesResponseContent(data string) *ListPendingVirtualMachineConfigurationChangesResponseContent {
	this := ListPendingVirtualMachineConfigurationChangesResponseContent{}
	this.Data = data
	return &this
}

// NewListPendingVirtualMachineConfigurationChangesResponseContentWithDefaults instantiates a new ListPendingVirtualMachineConfigurationChangesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPendingVirtualMachineConfigurationChangesResponseContentWithDefaults() *ListPendingVirtualMachineConfigurationChangesResponseContent {
	this := ListPendingVirtualMachineConfigurationChangesResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *ListPendingVirtualMachineConfigurationChangesResponseContent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListPendingVirtualMachineConfigurationChangesResponseContent) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ListPendingVirtualMachineConfigurationChangesResponseContent) SetData(v string) {
	o.Data = v
}

func (o ListPendingVirtualMachineConfigurationChangesResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPendingVirtualMachineConfigurationChangesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableListPendingVirtualMachineConfigurationChangesResponseContent struct {
	value *ListPendingVirtualMachineConfigurationChangesResponseContent
	isSet bool
}

func (v NullableListPendingVirtualMachineConfigurationChangesResponseContent) Get() *ListPendingVirtualMachineConfigurationChangesResponseContent {
	return v.value
}

func (v *NullableListPendingVirtualMachineConfigurationChangesResponseContent) Set(val *ListPendingVirtualMachineConfigurationChangesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListPendingVirtualMachineConfigurationChangesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListPendingVirtualMachineConfigurationChangesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPendingVirtualMachineConfigurationChangesResponseContent(val *ListPendingVirtualMachineConfigurationChangesResponseContent) *NullableListPendingVirtualMachineConfigurationChangesResponseContent {
	return &NullableListPendingVirtualMachineConfigurationChangesResponseContent{value: val, isSet: true}
}

func (v NullableListPendingVirtualMachineConfigurationChangesResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPendingVirtualMachineConfigurationChangesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


