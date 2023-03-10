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

// StoragePreallocation the model 'StoragePreallocation'
type StoragePreallocation string

// List of StoragePreallocation
const (
	STORAGEPREALLOCATION_OFF StoragePreallocation = "off"
	STORAGEPREALLOCATION_FALLOC StoragePreallocation = "falloc"
	STORAGEPREALLOCATION_FULL StoragePreallocation = "full"
	STORAGEPREALLOCATION_METADATA StoragePreallocation = "metadata"
)

// All allowed values of StoragePreallocation enum
var AllowedStoragePreallocationEnumValues = []StoragePreallocation{
	"off",
	"falloc",
	"full",
	"metadata",
}

func (v *StoragePreallocation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StoragePreallocation(value)
	for _, existing := range AllowedStoragePreallocationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StoragePreallocation", value)
}

// NewStoragePreallocationFromValue returns a pointer to a valid StoragePreallocation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStoragePreallocationFromValue(v string) (*StoragePreallocation, error) {
	ev := StoragePreallocation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StoragePreallocation: valid values are %v", v, AllowedStoragePreallocationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StoragePreallocation) IsValid() bool {
	for _, existing := range AllowedStoragePreallocationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StoragePreallocation value
func (v StoragePreallocation) Ptr() *StoragePreallocation {
	return &v
}

type NullableStoragePreallocation struct {
	value *StoragePreallocation
	isSet bool
}

func (v NullableStoragePreallocation) Get() *StoragePreallocation {
	return v.value
}

func (v *NullableStoragePreallocation) Set(val *StoragePreallocation) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePreallocation) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePreallocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePreallocation(val *StoragePreallocation) *NullableStoragePreallocation {
	return &NullableStoragePreallocation{value: val, isSet: true}
}

func (v NullableStoragePreallocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePreallocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

