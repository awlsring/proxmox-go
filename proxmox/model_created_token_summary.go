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

// checks if the CreatedTokenSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatedTokenSummary{}

// CreatedTokenSummary struct for CreatedTokenSummary
type CreatedTokenSummary struct {
	FullTokenid string `json:"full-tokenid"`
	Value string `json:"value"`
	Info UserConfigurationTokenSummary `json:"info"`
}

// NewCreatedTokenSummary instantiates a new CreatedTokenSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedTokenSummary(fullTokenid string, value string, info UserConfigurationTokenSummary) *CreatedTokenSummary {
	this := CreatedTokenSummary{}
	this.FullTokenid = fullTokenid
	this.Value = value
	this.Info = info
	return &this
}

// NewCreatedTokenSummaryWithDefaults instantiates a new CreatedTokenSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedTokenSummaryWithDefaults() *CreatedTokenSummary {
	this := CreatedTokenSummary{}
	return &this
}

// GetFullTokenid returns the FullTokenid field value
func (o *CreatedTokenSummary) GetFullTokenid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullTokenid
}

// GetFullTokenidOk returns a tuple with the FullTokenid field value
// and a boolean to check if the value has been set.
func (o *CreatedTokenSummary) GetFullTokenidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullTokenid, true
}

// SetFullTokenid sets field value
func (o *CreatedTokenSummary) SetFullTokenid(v string) {
	o.FullTokenid = v
}

// GetValue returns the Value field value
func (o *CreatedTokenSummary) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CreatedTokenSummary) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CreatedTokenSummary) SetValue(v string) {
	o.Value = v
}

// GetInfo returns the Info field value
func (o *CreatedTokenSummary) GetInfo() UserConfigurationTokenSummary {
	if o == nil {
		var ret UserConfigurationTokenSummary
		return ret
	}

	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
func (o *CreatedTokenSummary) GetInfoOk() (*UserConfigurationTokenSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Info, true
}

// SetInfo sets field value
func (o *CreatedTokenSummary) SetInfo(v UserConfigurationTokenSummary) {
	o.Info = v
}

func (o CreatedTokenSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatedTokenSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["full-tokenid"] = o.FullTokenid
	toSerialize["value"] = o.Value
	toSerialize["info"] = o.Info
	return toSerialize, nil
}

type NullableCreatedTokenSummary struct {
	value *CreatedTokenSummary
	isSet bool
}

func (v NullableCreatedTokenSummary) Get() *CreatedTokenSummary {
	return v.value
}

func (v *NullableCreatedTokenSummary) Set(val *CreatedTokenSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedTokenSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedTokenSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedTokenSummary(val *CreatedTokenSummary) *NullableCreatedTokenSummary {
	return &NullableCreatedTokenSummary{value: val, isSet: true}
}

func (v NullableCreatedTokenSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedTokenSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

