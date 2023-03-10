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

// checks if the VirtualMachineWriteFileRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualMachineWriteFileRequestContent{}

// VirtualMachineWriteFileRequestContent struct for VirtualMachineWriteFileRequestContent
type VirtualMachineWriteFileRequestContent struct {
	File string `json:"file"`
	Content string `json:"content"`
}

// NewVirtualMachineWriteFileRequestContent instantiates a new VirtualMachineWriteFileRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualMachineWriteFileRequestContent(file string, content string) *VirtualMachineWriteFileRequestContent {
	this := VirtualMachineWriteFileRequestContent{}
	this.File = file
	this.Content = content
	return &this
}

// NewVirtualMachineWriteFileRequestContentWithDefaults instantiates a new VirtualMachineWriteFileRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualMachineWriteFileRequestContentWithDefaults() *VirtualMachineWriteFileRequestContent {
	this := VirtualMachineWriteFileRequestContent{}
	return &this
}

// GetFile returns the File field value
func (o *VirtualMachineWriteFileRequestContent) GetFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineWriteFileRequestContent) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *VirtualMachineWriteFileRequestContent) SetFile(v string) {
	o.File = v
}

// GetContent returns the Content field value
func (o *VirtualMachineWriteFileRequestContent) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *VirtualMachineWriteFileRequestContent) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *VirtualMachineWriteFileRequestContent) SetContent(v string) {
	o.Content = v
}

func (o VirtualMachineWriteFileRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualMachineWriteFileRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file"] = o.File
	toSerialize["content"] = o.Content
	return toSerialize, nil
}

type NullableVirtualMachineWriteFileRequestContent struct {
	value *VirtualMachineWriteFileRequestContent
	isSet bool
}

func (v NullableVirtualMachineWriteFileRequestContent) Get() *VirtualMachineWriteFileRequestContent {
	return v.value
}

func (v *NullableVirtualMachineWriteFileRequestContent) Set(val *VirtualMachineWriteFileRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineWriteFileRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineWriteFileRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineWriteFileRequestContent(val *VirtualMachineWriteFileRequestContent) *NullableVirtualMachineWriteFileRequestContent {
	return &NullableVirtualMachineWriteFileRequestContent{value: val, isSet: true}
}

func (v NullableVirtualMachineWriteFileRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineWriteFileRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


