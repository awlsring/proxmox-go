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

// checks if the GetVirtualNetworkInterfacesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVirtualNetworkInterfacesResult{}

// GetVirtualNetworkInterfacesResult struct for GetVirtualNetworkInterfacesResult
type GetVirtualNetworkInterfacesResult struct {
	Result []VirtualNetworkInterfaceSummary `json:"result,omitempty"`
}

// NewGetVirtualNetworkInterfacesResult instantiates a new GetVirtualNetworkInterfacesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVirtualNetworkInterfacesResult() *GetVirtualNetworkInterfacesResult {
	this := GetVirtualNetworkInterfacesResult{}
	return &this
}

// NewGetVirtualNetworkInterfacesResultWithDefaults instantiates a new GetVirtualNetworkInterfacesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVirtualNetworkInterfacesResultWithDefaults() *GetVirtualNetworkInterfacesResult {
	this := GetVirtualNetworkInterfacesResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetVirtualNetworkInterfacesResult) GetResult() []VirtualNetworkInterfaceSummary {
	if o == nil || IsNil(o.Result) {
		var ret []VirtualNetworkInterfaceSummary
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVirtualNetworkInterfacesResult) GetResultOk() ([]VirtualNetworkInterfaceSummary, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetVirtualNetworkInterfacesResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []VirtualNetworkInterfaceSummary and assigns it to the Result field.
func (o *GetVirtualNetworkInterfacesResult) SetResult(v []VirtualNetworkInterfaceSummary) {
	o.Result = v
}

func (o GetVirtualNetworkInterfacesResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVirtualNetworkInterfacesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetVirtualNetworkInterfacesResult struct {
	value *GetVirtualNetworkInterfacesResult
	isSet bool
}

func (v NullableGetVirtualNetworkInterfacesResult) Get() *GetVirtualNetworkInterfacesResult {
	return v.value
}

func (v *NullableGetVirtualNetworkInterfacesResult) Set(val *GetVirtualNetworkInterfacesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVirtualNetworkInterfacesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVirtualNetworkInterfacesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVirtualNetworkInterfacesResult(val *GetVirtualNetworkInterfacesResult) *NullableGetVirtualNetworkInterfacesResult {
	return &NullableGetVirtualNetworkInterfacesResult{value: val, isSet: true}
}

func (v NullableGetVirtualNetworkInterfacesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVirtualNetworkInterfacesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


