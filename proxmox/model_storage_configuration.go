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

// checks if the StorageConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageConfiguration{}

// StorageConfiguration struct for StorageConfiguration
type StorageConfiguration struct {
	Storage string `json:"storage"`
	Type StorageType `json:"type"`
}

// NewStorageConfiguration instantiates a new StorageConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageConfiguration(storage string, type_ StorageType) *StorageConfiguration {
	this := StorageConfiguration{}
	this.Storage = storage
	this.Type = type_
	return &this
}

// NewStorageConfigurationWithDefaults instantiates a new StorageConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageConfigurationWithDefaults() *StorageConfiguration {
	this := StorageConfiguration{}
	return &this
}

// GetStorage returns the Storage field value
func (o *StorageConfiguration) GetStorage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *StorageConfiguration) GetStorageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *StorageConfiguration) SetStorage(v string) {
	o.Storage = v
}

// GetType returns the Type field value
func (o *StorageConfiguration) GetType() StorageType {
	if o == nil {
		var ret StorageType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StorageConfiguration) GetTypeOk() (*StorageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StorageConfiguration) SetType(v StorageType) {
	o.Type = v
}

func (o StorageConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["storage"] = o.Storage
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableStorageConfiguration struct {
	value *StorageConfiguration
	isSet bool
}

func (v NullableStorageConfiguration) Get() *StorageConfiguration {
	return v.value
}

func (v *NullableStorageConfiguration) Set(val *StorageConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageConfiguration(val *StorageConfiguration) *NullableStorageConfiguration {
	return &NullableStorageConfiguration{value: val, isSet: true}
}

func (v NullableStorageConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


