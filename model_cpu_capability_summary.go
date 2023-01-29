/*
Proxmox

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2023-01-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CpuCapabilitySummary struct for CpuCapabilitySummary
type CpuCapabilitySummary struct {
	Custom float32 `json:"custom"`
	Vendor string `json:"vendor"`
	Name string `json:"name"`
}

// NewCpuCapabilitySummary instantiates a new CpuCapabilitySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpuCapabilitySummary(custom float32, vendor string, name string) *CpuCapabilitySummary {
	this := CpuCapabilitySummary{}
	this.Custom = custom
	this.Vendor = vendor
	this.Name = name
	return &this
}

// NewCpuCapabilitySummaryWithDefaults instantiates a new CpuCapabilitySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpuCapabilitySummaryWithDefaults() *CpuCapabilitySummary {
	this := CpuCapabilitySummary{}
	return &this
}

// GetCustom returns the Custom field value
func (o *CpuCapabilitySummary) GetCustom() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value
// and a boolean to check if the value has been set.
func (o *CpuCapabilitySummary) GetCustomOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Custom, true
}

// SetCustom sets field value
func (o *CpuCapabilitySummary) SetCustom(v float32) {
	o.Custom = v
}

// GetVendor returns the Vendor field value
func (o *CpuCapabilitySummary) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *CpuCapabilitySummary) GetVendorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *CpuCapabilitySummary) SetVendor(v string) {
	o.Vendor = v
}

// GetName returns the Name field value
func (o *CpuCapabilitySummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CpuCapabilitySummary) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CpuCapabilitySummary) SetName(v string) {
	o.Name = v
}

func (o CpuCapabilitySummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["custom"] = o.Custom
	}
	if true {
		toSerialize["vendor"] = o.Vendor
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCpuCapabilitySummary struct {
	value *CpuCapabilitySummary
	isSet bool
}

func (v NullableCpuCapabilitySummary) Get() *CpuCapabilitySummary {
	return v.value
}

func (v *NullableCpuCapabilitySummary) Set(val *CpuCapabilitySummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCpuCapabilitySummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCpuCapabilitySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpuCapabilitySummary(val *CpuCapabilitySummary) *NullableCpuCapabilitySummary {
	return &NullableCpuCapabilitySummary{value: val, isSet: true}
}

func (v NullableCpuCapabilitySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpuCapabilitySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


