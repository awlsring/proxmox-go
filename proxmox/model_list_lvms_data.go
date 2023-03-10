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

// checks if the ListLVMsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListLVMsData{}

// ListLVMsData struct for ListLVMsData
type ListLVMsData struct {
	// An integer used to represent a boolean. 0 is false, 1 is true.
	Leaf float32 `json:"leaf"`
	Children []LVMChild `json:"children"`
}

// NewListLVMsData instantiates a new ListLVMsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLVMsData(leaf float32, children []LVMChild) *ListLVMsData {
	this := ListLVMsData{}
	this.Leaf = leaf
	this.Children = children
	return &this
}

// NewListLVMsDataWithDefaults instantiates a new ListLVMsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLVMsDataWithDefaults() *ListLVMsData {
	this := ListLVMsData{}
	return &this
}

// GetLeaf returns the Leaf field value
func (o *ListLVMsData) GetLeaf() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Leaf
}

// GetLeafOk returns a tuple with the Leaf field value
// and a boolean to check if the value has been set.
func (o *ListLVMsData) GetLeafOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Leaf, true
}

// SetLeaf sets field value
func (o *ListLVMsData) SetLeaf(v float32) {
	o.Leaf = v
}

// GetChildren returns the Children field value
func (o *ListLVMsData) GetChildren() []LVMChild {
	if o == nil {
		var ret []LVMChild
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *ListLVMsData) GetChildrenOk() ([]LVMChild, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *ListLVMsData) SetChildren(v []LVMChild) {
	o.Children = v
}

func (o ListLVMsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListLVMsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["leaf"] = o.Leaf
	toSerialize["children"] = o.Children
	return toSerialize, nil
}

type NullableListLVMsData struct {
	value *ListLVMsData
	isSet bool
}

func (v NullableListLVMsData) Get() *ListLVMsData {
	return v.value
}

func (v *NullableListLVMsData) Set(val *ListLVMsData) {
	v.value = val
	v.isSet = true
}

func (v NullableListLVMsData) IsSet() bool {
	return v.isSet
}

func (v *NullableListLVMsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLVMsData(val *ListLVMsData) *NullableListLVMsData {
	return &NullableListLVMsData{value: val, isSet: true}
}

func (v NullableListLVMsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLVMsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


