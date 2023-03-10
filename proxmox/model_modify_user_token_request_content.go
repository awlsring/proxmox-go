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

// checks if the ModifyUserTokenRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyUserTokenRequestContent{}

// ModifyUserTokenRequestContent struct for ModifyUserTokenRequestContent
type ModifyUserTokenRequestContent struct {
	// An integer used to represent a boolean. 0 is false, 1 is true.
	Privsep *float32 `json:"privsep,omitempty"`
	Expire *float32 `json:"expire,omitempty"`
	Comment *string `json:"comment,omitempty"`
}

// NewModifyUserTokenRequestContent instantiates a new ModifyUserTokenRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyUserTokenRequestContent() *ModifyUserTokenRequestContent {
	this := ModifyUserTokenRequestContent{}
	return &this
}

// NewModifyUserTokenRequestContentWithDefaults instantiates a new ModifyUserTokenRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyUserTokenRequestContentWithDefaults() *ModifyUserTokenRequestContent {
	this := ModifyUserTokenRequestContent{}
	return &this
}

// GetPrivsep returns the Privsep field value if set, zero value otherwise.
func (o *ModifyUserTokenRequestContent) GetPrivsep() float32 {
	if o == nil || IsNil(o.Privsep) {
		var ret float32
		return ret
	}
	return *o.Privsep
}

// GetPrivsepOk returns a tuple with the Privsep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserTokenRequestContent) GetPrivsepOk() (*float32, bool) {
	if o == nil || IsNil(o.Privsep) {
		return nil, false
	}
	return o.Privsep, true
}

// HasPrivsep returns a boolean if a field has been set.
func (o *ModifyUserTokenRequestContent) HasPrivsep() bool {
	if o != nil && !IsNil(o.Privsep) {
		return true
	}

	return false
}

// SetPrivsep gets a reference to the given float32 and assigns it to the Privsep field.
func (o *ModifyUserTokenRequestContent) SetPrivsep(v float32) {
	o.Privsep = &v
}

// GetExpire returns the Expire field value if set, zero value otherwise.
func (o *ModifyUserTokenRequestContent) GetExpire() float32 {
	if o == nil || IsNil(o.Expire) {
		var ret float32
		return ret
	}
	return *o.Expire
}

// GetExpireOk returns a tuple with the Expire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserTokenRequestContent) GetExpireOk() (*float32, bool) {
	if o == nil || IsNil(o.Expire) {
		return nil, false
	}
	return o.Expire, true
}

// HasExpire returns a boolean if a field has been set.
func (o *ModifyUserTokenRequestContent) HasExpire() bool {
	if o != nil && !IsNil(o.Expire) {
		return true
	}

	return false
}

// SetExpire gets a reference to the given float32 and assigns it to the Expire field.
func (o *ModifyUserTokenRequestContent) SetExpire(v float32) {
	o.Expire = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ModifyUserTokenRequestContent) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyUserTokenRequestContent) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ModifyUserTokenRequestContent) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ModifyUserTokenRequestContent) SetComment(v string) {
	o.Comment = &v
}

func (o ModifyUserTokenRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyUserTokenRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Privsep) {
		toSerialize["privsep"] = o.Privsep
	}
	if !IsNil(o.Expire) {
		toSerialize["expire"] = o.Expire
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableModifyUserTokenRequestContent struct {
	value *ModifyUserTokenRequestContent
	isSet bool
}

func (v NullableModifyUserTokenRequestContent) Get() *ModifyUserTokenRequestContent {
	return v.value
}

func (v *NullableModifyUserTokenRequestContent) Set(val *ModifyUserTokenRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyUserTokenRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyUserTokenRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyUserTokenRequestContent(val *ModifyUserTokenRequestContent) *NullableModifyUserTokenRequestContent {
	return &NullableModifyUserTokenRequestContent{value: val, isSet: true}
}

func (v NullableModifyUserTokenRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyUserTokenRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


