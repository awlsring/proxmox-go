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

// StorageTransport the model 'StorageTransport'
type StorageTransport string

// List of StorageTransport
const (
	STORAGETRANSPORT_TCP StorageTransport = "tcp"
	STORAGETRANSPORT_RDMA StorageTransport = "rdma"
	STORAGETRANSPORT_UNIX StorageTransport = "unix"
)

// All allowed values of StorageTransport enum
var AllowedStorageTransportEnumValues = []StorageTransport{
	"tcp",
	"rdma",
	"unix",
}

func (v *StorageTransport) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StorageTransport(value)
	for _, existing := range AllowedStorageTransportEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StorageTransport", value)
}

// NewStorageTransportFromValue returns a pointer to a valid StorageTransport
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageTransportFromValue(v string) (*StorageTransport, error) {
	ev := StorageTransport(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StorageTransport: valid values are %v", v, AllowedStorageTransportEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageTransport) IsValid() bool {
	for _, existing := range AllowedStorageTransportEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageTransport value
func (v StorageTransport) Ptr() *StorageTransport {
	return &v
}

type NullableStorageTransport struct {
	value *StorageTransport
	isSet bool
}

func (v NullableStorageTransport) Get() *StorageTransport {
	return v.value
}

func (v *NullableStorageTransport) Set(val *StorageTransport) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageTransport) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageTransport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageTransport(val *StorageTransport) *NullableStorageTransport {
	return &NullableStorageTransport{value: val, isSet: true}
}

func (v NullableStorageTransport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageTransport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

