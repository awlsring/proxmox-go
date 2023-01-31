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

// CreatePoolRequestContent struct for CreatePoolRequestContent
type CreatePoolRequestContent struct {
	Poolid string `json:"poolid"`
	Comment *string `json:"comment,omitempty"`
}

// NewCreatePoolRequestContent instantiates a new CreatePoolRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePoolRequestContent(poolid string) *CreatePoolRequestContent {
	this := CreatePoolRequestContent{}
	this.Poolid = poolid
	return &this
}

// NewCreatePoolRequestContentWithDefaults instantiates a new CreatePoolRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePoolRequestContentWithDefaults() *CreatePoolRequestContent {
	this := CreatePoolRequestContent{}
	return &this
}

// GetPoolid returns the Poolid field value
func (o *CreatePoolRequestContent) GetPoolid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Poolid
}

// GetPoolidOk returns a tuple with the Poolid field value
// and a boolean to check if the value has been set.
func (o *CreatePoolRequestContent) GetPoolidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Poolid, true
}

// SetPoolid sets field value
func (o *CreatePoolRequestContent) SetPoolid(v string) {
	o.Poolid = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreatePoolRequestContent) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePoolRequestContent) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreatePoolRequestContent) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreatePoolRequestContent) SetComment(v string) {
	o.Comment = &v
}

func (o CreatePoolRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["poolid"] = o.Poolid
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePoolRequestContent struct {
	value *CreatePoolRequestContent
	isSet bool
}

func (v NullableCreatePoolRequestContent) Get() *CreatePoolRequestContent {
	return v.value
}

func (v *NullableCreatePoolRequestContent) Set(val *CreatePoolRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePoolRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePoolRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePoolRequestContent(val *CreatePoolRequestContent) *NullableCreatePoolRequestContent {
	return &NullableCreatePoolRequestContent{value: val, isSet: true}
}

func (v NullableCreatePoolRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePoolRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

