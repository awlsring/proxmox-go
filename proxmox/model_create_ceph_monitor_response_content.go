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

// checks if the CreateCephMonitorResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCephMonitorResponseContent{}

// CreateCephMonitorResponseContent struct for CreateCephMonitorResponseContent
type CreateCephMonitorResponseContent struct {
	Data string `json:"data"`
}

// NewCreateCephMonitorResponseContent instantiates a new CreateCephMonitorResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCephMonitorResponseContent(data string) *CreateCephMonitorResponseContent {
	this := CreateCephMonitorResponseContent{}
	this.Data = data
	return &this
}

// NewCreateCephMonitorResponseContentWithDefaults instantiates a new CreateCephMonitorResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCephMonitorResponseContentWithDefaults() *CreateCephMonitorResponseContent {
	this := CreateCephMonitorResponseContent{}
	return &this
}

// GetData returns the Data field value
func (o *CreateCephMonitorResponseContent) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateCephMonitorResponseContent) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateCephMonitorResponseContent) SetData(v string) {
	o.Data = v
}

func (o CreateCephMonitorResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCephMonitorResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCreateCephMonitorResponseContent struct {
	value *CreateCephMonitorResponseContent
	isSet bool
}

func (v NullableCreateCephMonitorResponseContent) Get() *CreateCephMonitorResponseContent {
	return v.value
}

func (v *NullableCreateCephMonitorResponseContent) Set(val *CreateCephMonitorResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCephMonitorResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCephMonitorResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCephMonitorResponseContent(val *CreateCephMonitorResponseContent) *NullableCreateCephMonitorResponseContent {
	return &NullableCreateCephMonitorResponseContent{value: val, isSet: true}
}

func (v NullableCreateCephMonitorResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCephMonitorResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


