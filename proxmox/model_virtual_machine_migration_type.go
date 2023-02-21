/*
Proxmox

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2023-01-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package proxmox

import (
	"encoding/json"
	"fmt"
)

// VirtualMachineMigrationType the model 'VirtualMachineMigrationType'
type VirtualMachineMigrationType string

// List of VirtualMachineMigrationType
const (
	VIRTUALMACHINEMIGRATIONTYPE_SECURE VirtualMachineMigrationType = "secure"
	VIRTUALMACHINEMIGRATIONTYPE_INSECURE VirtualMachineMigrationType = "insecure"
)

// All allowed values of VirtualMachineMigrationType enum
var AllowedVirtualMachineMigrationTypeEnumValues = []VirtualMachineMigrationType{
	"secure",
	"insecure",
}

func (v *VirtualMachineMigrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualMachineMigrationType(value)
	for _, existing := range AllowedVirtualMachineMigrationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualMachineMigrationType", value)
}

// NewVirtualMachineMigrationTypeFromValue returns a pointer to a valid VirtualMachineMigrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualMachineMigrationTypeFromValue(v string) (*VirtualMachineMigrationType, error) {
	ev := VirtualMachineMigrationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualMachineMigrationType: valid values are %v", v, AllowedVirtualMachineMigrationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualMachineMigrationType) IsValid() bool {
	for _, existing := range AllowedVirtualMachineMigrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VirtualMachineMigrationType value
func (v VirtualMachineMigrationType) Ptr() *VirtualMachineMigrationType {
	return &v
}

type NullableVirtualMachineMigrationType struct {
	value *VirtualMachineMigrationType
	isSet bool
}

func (v NullableVirtualMachineMigrationType) Get() *VirtualMachineMigrationType {
	return v.value
}

func (v *NullableVirtualMachineMigrationType) Set(val *VirtualMachineMigrationType) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualMachineMigrationType) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualMachineMigrationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualMachineMigrationType(val *VirtualMachineMigrationType) *NullableVirtualMachineMigrationType {
	return &NullableVirtualMachineMigrationType{value: val, isSet: true}
}

func (v NullableVirtualMachineMigrationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualMachineMigrationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

