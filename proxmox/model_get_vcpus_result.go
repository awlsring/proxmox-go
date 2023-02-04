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

// checks if the GetVcpusResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVcpusResult{}

// GetVcpusResult struct for GetVcpusResult
type GetVcpusResult struct {
	Result []VcpuSummary `json:"result,omitempty"`
}

// NewGetVcpusResult instantiates a new GetVcpusResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVcpusResult() *GetVcpusResult {
	this := GetVcpusResult{}
	return &this
}

// NewGetVcpusResultWithDefaults instantiates a new GetVcpusResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVcpusResultWithDefaults() *GetVcpusResult {
	this := GetVcpusResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetVcpusResult) GetResult() []VcpuSummary {
	if o == nil || isNil(o.Result) {
		var ret []VcpuSummary
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVcpusResult) GetResultOk() ([]VcpuSummary, bool) {
	if o == nil || isNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetVcpusResult) HasResult() bool {
	if o != nil && !isNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []VcpuSummary and assigns it to the Result field.
func (o *GetVcpusResult) SetResult(v []VcpuSummary) {
	o.Result = v
}

func (o GetVcpusResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVcpusResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetVcpusResult struct {
	value *GetVcpusResult
	isSet bool
}

func (v NullableGetVcpusResult) Get() *GetVcpusResult {
	return v.value
}

func (v *NullableGetVcpusResult) Set(val *GetVcpusResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVcpusResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVcpusResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVcpusResult(val *GetVcpusResult) *NullableGetVcpusResult {
	return &NullableGetVcpusResult{value: val, isSet: true}
}

func (v NullableGetVcpusResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVcpusResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

