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

// checks if the UpdateSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSummary{}

// UpdateSummary struct for UpdateSummary
type UpdateSummary struct {
	Title string `json:"Title"`
	ChangeLogUrl string `json:"ChangeLogUrl"`
	Priority string `json:"Priority"`
	OldVersion string `json:"OldVersion"`
	Description string `json:"Description"`
	Arch string `json:"Arch"`
	Package string `json:"Package"`
	Section string `json:"Section"`
	Version string `json:"Version"`
	Origin string `json:"Origin"`
}

// NewUpdateSummary instantiates a new UpdateSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSummary(title string, changeLogUrl string, priority string, oldVersion string, description string, arch string, package_ string, section string, version string, origin string) *UpdateSummary {
	this := UpdateSummary{}
	this.Title = title
	this.ChangeLogUrl = changeLogUrl
	this.Priority = priority
	this.OldVersion = oldVersion
	this.Description = description
	this.Arch = arch
	this.Package = package_
	this.Section = section
	this.Version = version
	this.Origin = origin
	return &this
}

// NewUpdateSummaryWithDefaults instantiates a new UpdateSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSummaryWithDefaults() *UpdateSummary {
	this := UpdateSummary{}
	return &this
}

// GetTitle returns the Title field value
func (o *UpdateSummary) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *UpdateSummary) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *UpdateSummary) SetTitle(v string) {
	o.Title = v
}

// GetChangeLogUrl returns the ChangeLogUrl field value
func (o *UpdateSummary) GetChangeLogUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChangeLogUrl
}

// GetChangeLogUrlOk returns a tuple with the ChangeLogUrl field value
// and a boolean to check if the value has been set.
func (o *UpdateSummary) GetChangeLogUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeLogUrl, true
}

// SetChangeLogUrl sets field value
func (o *UpdateSummary) SetChangeLogUrl(v string) {
	o.ChangeLogUrl = v
}

// GetPriority returns the Priority field value
func (o *UpdateSummary) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *UpdateSummary) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *UpdateSummary) SetPriority(v string) {
	o.Priority = v
}

// GetOldVersion returns the OldVersion field value
func (o *UpdateSummary) GetOldVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OldVersion
}

// GetOldVersionOk returns a tuple with the OldVersion field value
// and a boolean to check if the value has been set.
func (o *UpdateSummary) GetOldVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldVersion, true
}

// SetOldVersion sets field value
func (o *UpdateSummary) SetOldVersion(v string) {
	o.OldVersion = v
}

// GetDescription returns the Description field value
func (o *UpdateSummary) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UpdateSummary) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UpdateSummary) SetDescription(v string) {
	o.Description = v
}

// GetArch returns the Arch field value
func (o *UpdateSummary) GetArch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arch
}

// GetArchOk returns a tuple with the Arch field value
// and a boolean to check if the value has been set.
func (o *UpdateSummary) GetArchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arch, true
}

// SetArch sets field value
func (o *UpdateSummary) SetArch(v string) {
	o.Arch = v
}

// GetPackage returns the Package field value
func (o *UpdateSummary) GetPackage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Package
}

// GetPackageOk returns a tuple with the Package field value
// and a boolean to check if the value has been set.
func (o *UpdateSummary) GetPackageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Package, true
}

// SetPackage sets field value
func (o *UpdateSummary) SetPackage(v string) {
	o.Package = v
}

// GetSection returns the Section field value
func (o *UpdateSummary) GetSection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Section
}

// GetSectionOk returns a tuple with the Section field value
// and a boolean to check if the value has been set.
func (o *UpdateSummary) GetSectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Section, true
}

// SetSection sets field value
func (o *UpdateSummary) SetSection(v string) {
	o.Section = v
}

// GetVersion returns the Version field value
func (o *UpdateSummary) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UpdateSummary) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *UpdateSummary) SetVersion(v string) {
	o.Version = v
}

// GetOrigin returns the Origin field value
func (o *UpdateSummary) GetOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *UpdateSummary) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *UpdateSummary) SetOrigin(v string) {
	o.Origin = v
}

func (o UpdateSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Title"] = o.Title
	toSerialize["ChangeLogUrl"] = o.ChangeLogUrl
	toSerialize["Priority"] = o.Priority
	toSerialize["OldVersion"] = o.OldVersion
	toSerialize["Description"] = o.Description
	toSerialize["Arch"] = o.Arch
	toSerialize["Package"] = o.Package
	toSerialize["Section"] = o.Section
	toSerialize["Version"] = o.Version
	toSerialize["Origin"] = o.Origin
	return toSerialize, nil
}

type NullableUpdateSummary struct {
	value *UpdateSummary
	isSet bool
}

func (v NullableUpdateSummary) Get() *UpdateSummary {
	return v.value
}

func (v *NullableUpdateSummary) Set(val *UpdateSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSummary(val *UpdateSummary) *NullableUpdateSummary {
	return &NullableUpdateSummary{value: val, isSet: true}
}

func (v NullableUpdateSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


