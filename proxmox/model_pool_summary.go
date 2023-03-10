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

// checks if the PoolSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolSummary{}

// PoolSummary struct for PoolSummary
type PoolSummary struct {
	Poolid string `json:"poolid"`
	Comment *string `json:"comment,omitempty"`
}

// NewPoolSummary instantiates a new PoolSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolSummary(poolid string) *PoolSummary {
	this := PoolSummary{}
	this.Poolid = poolid
	return &this
}

// NewPoolSummaryWithDefaults instantiates a new PoolSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolSummaryWithDefaults() *PoolSummary {
	this := PoolSummary{}
	return &this
}

// GetPoolid returns the Poolid field value
func (o *PoolSummary) GetPoolid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Poolid
}

// GetPoolidOk returns a tuple with the Poolid field value
// and a boolean to check if the value has been set.
func (o *PoolSummary) GetPoolidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Poolid, true
}

// SetPoolid sets field value
func (o *PoolSummary) SetPoolid(v string) {
	o.Poolid = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *PoolSummary) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolSummary) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *PoolSummary) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *PoolSummary) SetComment(v string) {
	o.Comment = &v
}

func (o PoolSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["poolid"] = o.Poolid
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullablePoolSummary struct {
	value *PoolSummary
	isSet bool
}

func (v NullablePoolSummary) Get() *PoolSummary {
	return v.value
}

func (v *NullablePoolSummary) Set(val *PoolSummary) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolSummary) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolSummary(val *PoolSummary) *NullablePoolSummary {
	return &NullablePoolSummary{value: val, isSet: true}
}

func (v NullablePoolSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


