/*
Proxmox

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2023-01-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// VirtualMachineScsiControllerType the model 'VirtualMachineScsiControllerType'
type VirtualMachineScsiControllerType string

// List of VirtualMachineScsiControllerType
const (
	VIRTUALMACHINESCSICONTROLLERTYPE_LSI VirtualMachineScsiControllerType = "lsi"
	VIRTUALMACHINESCSICONTROLLERTYPE_LSI53C810 VirtualMachineScsiControllerType = "lsi53c810"
	VIRTUALMACHINESCSICONTROLLERTYPE_VIRTIO_SCSI_PCI VirtualMachineScsiControllerType = "virtio-scsi-pci"
	VIRTUALMACHINESCSICONTROLLERTYPE_MEGASAS VirtualMachineScsiControllerType = "megasas"
	VIRTUALMACHINESCSICONTROLLERTYPE_PVSCSI VirtualMachineScsiControllerType = "pvscsi"
)

// All allowed values of VirtualMachineScsiControllerType enum
var AllowedVirtualMachineScsiControllerTypeEnumValues = []VirtualMachineScsiControllerType{
	"lsi",
	"lsi53c810",
	"virtio-scsi-pci",
	"megasas",
	"pvscsi",
}

func (v *VirtualMachineScsiControllerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualMachineScsiControllerType(value)
	for _, existing := range AllowedVirtualMachineScsiControllerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualMachineScsiControllerType", value)
}

// NewVirtualMachineScsiControllerTypeFromValue returns a pointer to a valid VirtualMachineScsiControllerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualMachineScsiControllerTypeFromValue(v string) (*VirtualMachineScsiControllerType, error) {
	ev := VirtualMachineScsiControllerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualMachineScsiControllerType: valid values are %v", v, AllowedVirtualMachineScsiControllerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualMachineScsiControllerType) IsValid() bool {
	for _, existing := range AllowedVirtualMachineScsiControllerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VirtualMachineScsiControllerType value
func (v VirtualMachineScsiControllerType) Ptr() *VirtualMachineScsiControllerType {
	return &v
}

type NullableVirtualMachineScsiControllerType struct {
	value *VirtualMachineScsiControllerType
	isSet bool
}

func (v NullableVirtualMachineScsiControllerType) Get() *VirtualMachineScsiControllerType {
	return v.value
}

func (v *NullableVirtualMachineScsiControllerType) Set(val *VirtualMachineScsiControllerType) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineScsiControllerType) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineScsiControllerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineScsiControllerType(val *VirtualMachineScsiControllerType) *NullableVirtualMachineScsiControllerType {
	return &NullableVirtualMachineScsiControllerType{value: val, isSet: true}
}

func (v NullableVirtualMachineScsiControllerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineScsiControllerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

