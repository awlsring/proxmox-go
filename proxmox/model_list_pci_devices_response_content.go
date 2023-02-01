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

// ListPciDevicesResponseContent struct for ListPciDevicesResponseContent
type ListPciDevicesResponseContent struct {
	Data []UsbDeviceSummary `json:"data"`
}

// NewListPciDevicesResponseContent instantiates a new ListPciDevicesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPciDevicesResponseContent(data []UsbDeviceSummary) *ListPciDevicesResponseContent {
	this := ListPciDevicesResponseContent{}
	this.Data = data
	return &this
}

// NewListPciDevicesResponseContentWithDefaults instantiates a new ListPciDevicesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPciDevicesResponseContentWithDefaults() *ListPciDevicesResponseContent {
	this := ListPciDevicesResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *ListPciDevicesResponseContent) GetData() []UsbDeviceSummary {
	if o == nil {
		var ret []UsbDeviceSummary
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListPciDevicesResponseContent) GetDataOk() ([]UsbDeviceSummary, bool) {
	if o == nil {
    return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListPciDevicesResponseContent) SetData(v []UsbDeviceSummary) {
	o.Data = v
}

func (o ListPciDevicesResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListPciDevicesResponseContent struct {
	value *ListPciDevicesResponseContent
	isSet bool
}

func (v NullableListPciDevicesResponseContent) Get() *ListPciDevicesResponseContent {
	return v.value
}

func (v *NullableListPciDevicesResponseContent) Set(val *ListPciDevicesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListPciDevicesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListPciDevicesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPciDevicesResponseContent(val *ListPciDevicesResponseContent) *NullableListPciDevicesResponseContent {
	return &NullableListPciDevicesResponseContent{value: val, isSet: true}
}

func (v NullableListPciDevicesResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPciDevicesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

