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

// checks if the ListMachineCapabilitiesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMachineCapabilitiesResponseContent{}

// ListMachineCapabilitiesResponseContent struct for ListMachineCapabilitiesResponseContent
type ListMachineCapabilitiesResponseContent struct {
	Data []MachineCapabilitySummary `json:"data"`
}

// NewListMachineCapabilitiesResponseContent instantiates a new ListMachineCapabilitiesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMachineCapabilitiesResponseContent(data []MachineCapabilitySummary) *ListMachineCapabilitiesResponseContent {
	this := ListMachineCapabilitiesResponseContent{}
	this.Data = data
	return &this
}

// NewListMachineCapabilitiesResponseContentWithDefaults instantiates a new ListMachineCapabilitiesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMachineCapabilitiesResponseContentWithDefaults() *ListMachineCapabilitiesResponseContent {
	this := ListMachineCapabilitiesResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *ListMachineCapabilitiesResponseContent) GetData() []MachineCapabilitySummary {
	if o == nil {
		var ret []MachineCapabilitySummary
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListMachineCapabilitiesResponseContent) GetDataOk() ([]MachineCapabilitySummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListMachineCapabilitiesResponseContent) SetData(v []MachineCapabilitySummary) {
	o.Data = v
}

func (o ListMachineCapabilitiesResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMachineCapabilitiesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableListMachineCapabilitiesResponseContent struct {
	value *ListMachineCapabilitiesResponseContent
	isSet bool
}

func (v NullableListMachineCapabilitiesResponseContent) Get() *ListMachineCapabilitiesResponseContent {
	return v.value
}

func (v *NullableListMachineCapabilitiesResponseContent) Set(val *ListMachineCapabilitiesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListMachineCapabilitiesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListMachineCapabilitiesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMachineCapabilitiesResponseContent(val *ListMachineCapabilitiesResponseContent) *NullableListMachineCapabilitiesResponseContent {
	return &NullableListMachineCapabilitiesResponseContent{value: val, isSet: true}
}

func (v NullableListMachineCapabilitiesResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMachineCapabilitiesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


