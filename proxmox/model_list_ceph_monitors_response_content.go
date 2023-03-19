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

// checks if the ListCephMonitorsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCephMonitorsResponseContent{}

// ListCephMonitorsResponseContent struct for ListCephMonitorsResponseContent
type ListCephMonitorsResponseContent struct {
	Data []CephMonitorSummary `json:"data"`
}

// NewListCephMonitorsResponseContent instantiates a new ListCephMonitorsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCephMonitorsResponseContent(data []CephMonitorSummary) *ListCephMonitorsResponseContent {
	this := ListCephMonitorsResponseContent{}
	this.Data = data
	return &this
}

// NewListCephMonitorsResponseContentWithDefaults instantiates a new ListCephMonitorsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCephMonitorsResponseContentWithDefaults() *ListCephMonitorsResponseContent {
	this := ListCephMonitorsResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *ListCephMonitorsResponseContent) GetData() []CephMonitorSummary {
	if o == nil {
		var ret []CephMonitorSummary
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListCephMonitorsResponseContent) GetDataOk() ([]CephMonitorSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListCephMonitorsResponseContent) SetData(v []CephMonitorSummary) {
	o.Data = v
}

func (o ListCephMonitorsResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCephMonitorsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableListCephMonitorsResponseContent struct {
	value *ListCephMonitorsResponseContent
	isSet bool
}

func (v NullableListCephMonitorsResponseContent) Get() *ListCephMonitorsResponseContent {
	return v.value
}

func (v *NullableListCephMonitorsResponseContent) Set(val *ListCephMonitorsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListCephMonitorsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListCephMonitorsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCephMonitorsResponseContent(val *ListCephMonitorsResponseContent) *NullableListCephMonitorsResponseContent {
	return &NullableListCephMonitorsResponseContent{value: val, isSet: true}
}

func (v NullableListCephMonitorsResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCephMonitorsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


