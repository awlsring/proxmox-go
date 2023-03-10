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

// checks if the CreateZFSPoolRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateZFSPoolRequestContent{}

// CreateZFSPoolRequestContent struct for CreateZFSPoolRequestContent
type CreateZFSPoolRequestContent struct {
	// The devices to create the zfs pool on. This is a comma seperated list sent as a string.
	Devices string `json:"devices"`
	// The storage identifier.
	Name string `json:"name"`
	Raidlevel ZFSRaidLevel `json:"raidlevel"`
	// Configure storage using the directory. Takes a boolean integer value (0 false, 1 true).
	AddStorage *float32 `json:"add_storage,omitempty"`
	// The pool vector size exponent.
	Ashift *float32 `json:"ashift,omitempty"`
	Compression *ZFSCompression `json:"compression,omitempty"`
	// Draid config. Set as string like 'data=<int>,spares=<int>
	DraidConfig *string `json:"draid-config,omitempty"`
}

// NewCreateZFSPoolRequestContent instantiates a new CreateZFSPoolRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateZFSPoolRequestContent(devices string, name string, raidlevel ZFSRaidLevel) *CreateZFSPoolRequestContent {
	this := CreateZFSPoolRequestContent{}
	this.Devices = devices
	this.Name = name
	this.Raidlevel = raidlevel
	return &this
}

// NewCreateZFSPoolRequestContentWithDefaults instantiates a new CreateZFSPoolRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateZFSPoolRequestContentWithDefaults() *CreateZFSPoolRequestContent {
	this := CreateZFSPoolRequestContent{}
	return &this
}

// GetDevices returns the Devices field value
func (o *CreateZFSPoolRequestContent) GetDevices() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *CreateZFSPoolRequestContent) GetDevicesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Devices, true
}

// SetDevices sets field value
func (o *CreateZFSPoolRequestContent) SetDevices(v string) {
	o.Devices = v
}

// GetName returns the Name field value
func (o *CreateZFSPoolRequestContent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateZFSPoolRequestContent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateZFSPoolRequestContent) SetName(v string) {
	o.Name = v
}

// GetRaidlevel returns the Raidlevel field value
func (o *CreateZFSPoolRequestContent) GetRaidlevel() ZFSRaidLevel {
	if o == nil {
		var ret ZFSRaidLevel
		return ret
	}

	return o.Raidlevel
}

// GetRaidlevelOk returns a tuple with the Raidlevel field value
// and a boolean to check if the value has been set.
func (o *CreateZFSPoolRequestContent) GetRaidlevelOk() (*ZFSRaidLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Raidlevel, true
}

// SetRaidlevel sets field value
func (o *CreateZFSPoolRequestContent) SetRaidlevel(v ZFSRaidLevel) {
	o.Raidlevel = v
}

// GetAddStorage returns the AddStorage field value if set, zero value otherwise.
func (o *CreateZFSPoolRequestContent) GetAddStorage() float32 {
	if o == nil || IsNil(o.AddStorage) {
		var ret float32
		return ret
	}
	return *o.AddStorage
}

// GetAddStorageOk returns a tuple with the AddStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZFSPoolRequestContent) GetAddStorageOk() (*float32, bool) {
	if o == nil || IsNil(o.AddStorage) {
		return nil, false
	}
	return o.AddStorage, true
}

// HasAddStorage returns a boolean if a field has been set.
func (o *CreateZFSPoolRequestContent) HasAddStorage() bool {
	if o != nil && !IsNil(o.AddStorage) {
		return true
	}

	return false
}

// SetAddStorage gets a reference to the given float32 and assigns it to the AddStorage field.
func (o *CreateZFSPoolRequestContent) SetAddStorage(v float32) {
	o.AddStorage = &v
}

// GetAshift returns the Ashift field value if set, zero value otherwise.
func (o *CreateZFSPoolRequestContent) GetAshift() float32 {
	if o == nil || IsNil(o.Ashift) {
		var ret float32
		return ret
	}
	return *o.Ashift
}

// GetAshiftOk returns a tuple with the Ashift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZFSPoolRequestContent) GetAshiftOk() (*float32, bool) {
	if o == nil || IsNil(o.Ashift) {
		return nil, false
	}
	return o.Ashift, true
}

// HasAshift returns a boolean if a field has been set.
func (o *CreateZFSPoolRequestContent) HasAshift() bool {
	if o != nil && !IsNil(o.Ashift) {
		return true
	}

	return false
}

// SetAshift gets a reference to the given float32 and assigns it to the Ashift field.
func (o *CreateZFSPoolRequestContent) SetAshift(v float32) {
	o.Ashift = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *CreateZFSPoolRequestContent) GetCompression() ZFSCompression {
	if o == nil || IsNil(o.Compression) {
		var ret ZFSCompression
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZFSPoolRequestContent) GetCompressionOk() (*ZFSCompression, bool) {
	if o == nil || IsNil(o.Compression) {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *CreateZFSPoolRequestContent) HasCompression() bool {
	if o != nil && !IsNil(o.Compression) {
		return true
	}

	return false
}

// SetCompression gets a reference to the given ZFSCompression and assigns it to the Compression field.
func (o *CreateZFSPoolRequestContent) SetCompression(v ZFSCompression) {
	o.Compression = &v
}

// GetDraidConfig returns the DraidConfig field value if set, zero value otherwise.
func (o *CreateZFSPoolRequestContent) GetDraidConfig() string {
	if o == nil || IsNil(o.DraidConfig) {
		var ret string
		return ret
	}
	return *o.DraidConfig
}

// GetDraidConfigOk returns a tuple with the DraidConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZFSPoolRequestContent) GetDraidConfigOk() (*string, bool) {
	if o == nil || IsNil(o.DraidConfig) {
		return nil, false
	}
	return o.DraidConfig, true
}

// HasDraidConfig returns a boolean if a field has been set.
func (o *CreateZFSPoolRequestContent) HasDraidConfig() bool {
	if o != nil && !IsNil(o.DraidConfig) {
		return true
	}

	return false
}

// SetDraidConfig gets a reference to the given string and assigns it to the DraidConfig field.
func (o *CreateZFSPoolRequestContent) SetDraidConfig(v string) {
	o.DraidConfig = &v
}

func (o CreateZFSPoolRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateZFSPoolRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["devices"] = o.Devices
	toSerialize["name"] = o.Name
	toSerialize["raidlevel"] = o.Raidlevel
	if !IsNil(o.AddStorage) {
		toSerialize["add_storage"] = o.AddStorage
	}
	if !IsNil(o.Ashift) {
		toSerialize["ashift"] = o.Ashift
	}
	if !IsNil(o.Compression) {
		toSerialize["compression"] = o.Compression
	}
	if !IsNil(o.DraidConfig) {
		toSerialize["draid-config"] = o.DraidConfig
	}
	return toSerialize, nil
}

type NullableCreateZFSPoolRequestContent struct {
	value *CreateZFSPoolRequestContent
	isSet bool
}

func (v NullableCreateZFSPoolRequestContent) Get() *CreateZFSPoolRequestContent {
	return v.value
}

func (v *NullableCreateZFSPoolRequestContent) Set(val *CreateZFSPoolRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateZFSPoolRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateZFSPoolRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateZFSPoolRequestContent(val *CreateZFSPoolRequestContent) *NullableCreateZFSPoolRequestContent {
	return &NullableCreateZFSPoolRequestContent{value: val, isSet: true}
}

func (v NullableCreateZFSPoolRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateZFSPoolRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


