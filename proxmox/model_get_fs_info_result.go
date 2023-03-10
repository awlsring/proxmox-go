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

// checks if the GetFsInfoResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFsInfoResult{}

// GetFsInfoResult struct for GetFsInfoResult
type GetFsInfoResult struct {
	Result []FileSystemInformationSummary `json:"result,omitempty"`
}

// NewGetFsInfoResult instantiates a new GetFsInfoResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFsInfoResult() *GetFsInfoResult {
	this := GetFsInfoResult{}
	return &this
}

// NewGetFsInfoResultWithDefaults instantiates a new GetFsInfoResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFsInfoResultWithDefaults() *GetFsInfoResult {
	this := GetFsInfoResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetFsInfoResult) GetResult() []FileSystemInformationSummary {
	if o == nil || IsNil(o.Result) {
		var ret []FileSystemInformationSummary
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFsInfoResult) GetResultOk() ([]FileSystemInformationSummary, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetFsInfoResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []FileSystemInformationSummary and assigns it to the Result field.
func (o *GetFsInfoResult) SetResult(v []FileSystemInformationSummary) {
	o.Result = v
}

func (o GetFsInfoResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFsInfoResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetFsInfoResult struct {
	value *GetFsInfoResult
	isSet bool
}

func (v NullableGetFsInfoResult) Get() *GetFsInfoResult {
	return v.value
}

func (v *NullableGetFsInfoResult) Set(val *GetFsInfoResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFsInfoResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFsInfoResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFsInfoResult(val *GetFsInfoResult) *NullableGetFsInfoResult {
	return &NullableGetFsInfoResult{value: val, isSet: true}
}

func (v NullableGetFsInfoResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFsInfoResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


