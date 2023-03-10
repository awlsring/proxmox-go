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

// checks if the PciControllerSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PciControllerSummary{}

// PciControllerSummary struct for PciControllerSummary
type PciControllerSummary struct {
	Bus *float32 `json:"bus,omitempty"`
	Domain *float32 `json:"domain,omitempty"`
	Function *float32 `json:"function,omitempty"`
	Slot *float32 `json:"slot,omitempty"`
}

// NewPciControllerSummary instantiates a new PciControllerSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPciControllerSummary() *PciControllerSummary {
	this := PciControllerSummary{}
	return &this
}

// NewPciControllerSummaryWithDefaults instantiates a new PciControllerSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPciControllerSummaryWithDefaults() *PciControllerSummary {
	this := PciControllerSummary{}
	return &this
}

// GetBus returns the Bus field value if set, zero value otherwise.
func (o *PciControllerSummary) GetBus() float32 {
	if o == nil || IsNil(o.Bus) {
		var ret float32
		return ret
	}
	return *o.Bus
}

// GetBusOk returns a tuple with the Bus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciControllerSummary) GetBusOk() (*float32, bool) {
	if o == nil || IsNil(o.Bus) {
		return nil, false
	}
	return o.Bus, true
}

// HasBus returns a boolean if a field has been set.
func (o *PciControllerSummary) HasBus() bool {
	if o != nil && !IsNil(o.Bus) {
		return true
	}

	return false
}

// SetBus gets a reference to the given float32 and assigns it to the Bus field.
func (o *PciControllerSummary) SetBus(v float32) {
	o.Bus = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *PciControllerSummary) GetDomain() float32 {
	if o == nil || IsNil(o.Domain) {
		var ret float32
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciControllerSummary) GetDomainOk() (*float32, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *PciControllerSummary) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given float32 and assigns it to the Domain field.
func (o *PciControllerSummary) SetDomain(v float32) {
	o.Domain = &v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *PciControllerSummary) GetFunction() float32 {
	if o == nil || IsNil(o.Function) {
		var ret float32
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciControllerSummary) GetFunctionOk() (*float32, bool) {
	if o == nil || IsNil(o.Function) {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *PciControllerSummary) HasFunction() bool {
	if o != nil && !IsNil(o.Function) {
		return true
	}

	return false
}

// SetFunction gets a reference to the given float32 and assigns it to the Function field.
func (o *PciControllerSummary) SetFunction(v float32) {
	o.Function = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *PciControllerSummary) GetSlot() float32 {
	if o == nil || IsNil(o.Slot) {
		var ret float32
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciControllerSummary) GetSlotOk() (*float32, bool) {
	if o == nil || IsNil(o.Slot) {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *PciControllerSummary) HasSlot() bool {
	if o != nil && !IsNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given float32 and assigns it to the Slot field.
func (o *PciControllerSummary) SetSlot(v float32) {
	o.Slot = &v
}

func (o PciControllerSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PciControllerSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bus) {
		toSerialize["bus"] = o.Bus
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Function) {
		toSerialize["function"] = o.Function
	}
	if !IsNil(o.Slot) {
		toSerialize["slot"] = o.Slot
	}
	return toSerialize, nil
}

type NullablePciControllerSummary struct {
	value *PciControllerSummary
	isSet bool
}

func (v NullablePciControllerSummary) Get() *PciControllerSummary {
	return v.value
}

func (v *NullablePciControllerSummary) Set(val *PciControllerSummary) {
	v.value = val
	v.isSet = true
}

func (v NullablePciControllerSummary) IsSet() bool {
	return v.isSet
}

func (v *NullablePciControllerSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePciControllerSummary(val *PciControllerSummary) *NullablePciControllerSummary {
	return &NullablePciControllerSummary{value: val, isSet: true}
}

func (v NullablePciControllerSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePciControllerSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


