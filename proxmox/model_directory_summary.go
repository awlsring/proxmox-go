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

// checks if the DirectorySummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectorySummary{}

// DirectorySummary struct for DirectorySummary
type DirectorySummary struct {
	// The mounted device
	Device string `json:"device"`
	// The mount options
	Options string `json:"options"`
	// The mount path
	Path string `json:"path"`
	// The mount type
	Type string `json:"type"`
	// The path of the mount unit
	Unitfile string `json:"unitfile"`
}

// NewDirectorySummary instantiates a new DirectorySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectorySummary(device string, options string, path string, type_ string, unitfile string) *DirectorySummary {
	this := DirectorySummary{}
	this.Device = device
	this.Options = options
	this.Path = path
	this.Type = type_
	this.Unitfile = unitfile
	return &this
}

// NewDirectorySummaryWithDefaults instantiates a new DirectorySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectorySummaryWithDefaults() *DirectorySummary {
	this := DirectorySummary{}
	return &this
}

// GetDevice returns the Device field value
func (o *DirectorySummary) GetDevice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *DirectorySummary) GetDeviceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *DirectorySummary) SetDevice(v string) {
	o.Device = v
}

// GetOptions returns the Options field value
func (o *DirectorySummary) GetOptions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *DirectorySummary) GetOptionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *DirectorySummary) SetOptions(v string) {
	o.Options = v
}

// GetPath returns the Path field value
func (o *DirectorySummary) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DirectorySummary) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DirectorySummary) SetPath(v string) {
	o.Path = v
}

// GetType returns the Type field value
func (o *DirectorySummary) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DirectorySummary) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DirectorySummary) SetType(v string) {
	o.Type = v
}

// GetUnitfile returns the Unitfile field value
func (o *DirectorySummary) GetUnitfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unitfile
}

// GetUnitfileOk returns a tuple with the Unitfile field value
// and a boolean to check if the value has been set.
func (o *DirectorySummary) GetUnitfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unitfile, true
}

// SetUnitfile sets field value
func (o *DirectorySummary) SetUnitfile(v string) {
	o.Unitfile = v
}

func (o DirectorySummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectorySummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	toSerialize["options"] = o.Options
	toSerialize["path"] = o.Path
	toSerialize["type"] = o.Type
	toSerialize["unitfile"] = o.Unitfile
	return toSerialize, nil
}

type NullableDirectorySummary struct {
	value *DirectorySummary
	isSet bool
}

func (v NullableDirectorySummary) Get() *DirectorySummary {
	return v.value
}

func (v *NullableDirectorySummary) Set(val *DirectorySummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectorySummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectorySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectorySummary(val *DirectorySummary) *NullableDirectorySummary {
	return &NullableDirectorySummary{value: val, isSet: true}
}

func (v NullableDirectorySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectorySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


