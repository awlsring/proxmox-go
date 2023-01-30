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

// LVMSummary struct for LVMSummary
type LVMSummary struct {
	Children []LVMSummary `json:"children"`
	// Is leaf. This is a boolean intergar, 1 is true, 0 is false
	Leaf float32 `json:"leaf"`
	Name *string `json:"name,omitempty"`
	// The free space on lvm in bytes
	Free *float32 `json:"free,omitempty"`
	// The total size of lvm in bytes
	Size *float32 `json:"size,omitempty"`
	// The number of logical volumes
	Lvcount *float32 `json:"lvcount,omitempty"`
}

// NewLVMSummary instantiates a new LVMSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLVMSummary(children []LVMSummary, leaf float32) *LVMSummary {
	this := LVMSummary{}
	this.Children = children
	this.Leaf = leaf
	return &this
}

// NewLVMSummaryWithDefaults instantiates a new LVMSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLVMSummaryWithDefaults() *LVMSummary {
	this := LVMSummary{}
	return &this
}

// GetChildren returns the Children field value
func (o *LVMSummary) GetChildren() []LVMSummary {
	if o == nil {
		var ret []LVMSummary
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *LVMSummary) GetChildrenOk() ([]LVMSummary, bool) {
	if o == nil {
    return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *LVMSummary) SetChildren(v []LVMSummary) {
	o.Children = v
}

// GetLeaf returns the Leaf field value
func (o *LVMSummary) GetLeaf() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Leaf
}

// GetLeafOk returns a tuple with the Leaf field value
// and a boolean to check if the value has been set.
func (o *LVMSummary) GetLeafOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Leaf, true
}

// SetLeaf sets field value
func (o *LVMSummary) SetLeaf(v float32) {
	o.Leaf = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LVMSummary) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LVMSummary) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LVMSummary) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LVMSummary) SetName(v string) {
	o.Name = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *LVMSummary) GetFree() float32 {
	if o == nil || isNil(o.Free) {
		var ret float32
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LVMSummary) GetFreeOk() (*float32, bool) {
	if o == nil || isNil(o.Free) {
    return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *LVMSummary) HasFree() bool {
	if o != nil && !isNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given float32 and assigns it to the Free field.
func (o *LVMSummary) SetFree(v float32) {
	o.Free = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *LVMSummary) GetSize() float32 {
	if o == nil || isNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LVMSummary) GetSizeOk() (*float32, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *LVMSummary) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *LVMSummary) SetSize(v float32) {
	o.Size = &v
}

// GetLvcount returns the Lvcount field value if set, zero value otherwise.
func (o *LVMSummary) GetLvcount() float32 {
	if o == nil || isNil(o.Lvcount) {
		var ret float32
		return ret
	}
	return *o.Lvcount
}

// GetLvcountOk returns a tuple with the Lvcount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LVMSummary) GetLvcountOk() (*float32, bool) {
	if o == nil || isNil(o.Lvcount) {
    return nil, false
	}
	return o.Lvcount, true
}

// HasLvcount returns a boolean if a field has been set.
func (o *LVMSummary) HasLvcount() bool {
	if o != nil && !isNil(o.Lvcount) {
		return true
	}

	return false
}

// SetLvcount gets a reference to the given float32 and assigns it to the Lvcount field.
func (o *LVMSummary) SetLvcount(v float32) {
	o.Lvcount = &v
}

func (o LVMSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["children"] = o.Children
	}
	if true {
		toSerialize["leaf"] = o.Leaf
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Lvcount) {
		toSerialize["lvcount"] = o.Lvcount
	}
	return json.Marshal(toSerialize)
}

type NullableLVMSummary struct {
	value *LVMSummary
	isSet bool
}

func (v NullableLVMSummary) Get() *LVMSummary {
	return v.value
}

func (v *NullableLVMSummary) Set(val *LVMSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableLVMSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableLVMSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLVMSummary(val *LVMSummary) *NullableLVMSummary {
	return &NullableLVMSummary{value: val, isSet: true}
}

func (v NullableLVMSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLVMSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


