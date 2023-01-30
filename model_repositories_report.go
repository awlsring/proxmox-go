/*
Proxmox

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2023-01-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// RepositoriesReport struct for RepositoriesReport
type RepositoriesReport struct {
	Digest string `json:"digest"`
	StandardRepos []RepositorySummary `json:"standard-repos"`
	Errors []string `json:"errors"`
	Files []FileSummary `json:"files"`
	Infos []FileInfoSummary `json:"infos"`
}

// NewRepositoriesReport instantiates a new RepositoriesReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositoriesReport(digest string, standardRepos []RepositorySummary, errors []string, files []FileSummary, infos []FileInfoSummary) *RepositoriesReport {
	this := RepositoriesReport{}
	this.Digest = digest
	this.StandardRepos = standardRepos
	this.Errors = errors
	this.Files = files
	this.Infos = infos
	return &this
}

// NewRepositoriesReportWithDefaults instantiates a new RepositoriesReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoriesReportWithDefaults() *RepositoriesReport {
	this := RepositoriesReport{}
	return &this
}

// GetDigest returns the Digest field value
func (o *RepositoriesReport) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *RepositoriesReport) GetDigestOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *RepositoriesReport) SetDigest(v string) {
	o.Digest = v
}

// GetStandardRepos returns the StandardRepos field value
func (o *RepositoriesReport) GetStandardRepos() []RepositorySummary {
	if o == nil {
		var ret []RepositorySummary
		return ret
	}

	return o.StandardRepos
}

// GetStandardReposOk returns a tuple with the StandardRepos field value
// and a boolean to check if the value has been set.
func (o *RepositoriesReport) GetStandardReposOk() ([]RepositorySummary, bool) {
	if o == nil {
    return nil, false
	}
	return o.StandardRepos, true
}

// SetStandardRepos sets field value
func (o *RepositoriesReport) SetStandardRepos(v []RepositorySummary) {
	o.StandardRepos = v
}

// GetErrors returns the Errors field value
func (o *RepositoriesReport) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *RepositoriesReport) GetErrorsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *RepositoriesReport) SetErrors(v []string) {
	o.Errors = v
}

// GetFiles returns the Files field value
func (o *RepositoriesReport) GetFiles() []FileSummary {
	if o == nil {
		var ret []FileSummary
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *RepositoriesReport) GetFilesOk() ([]FileSummary, bool) {
	if o == nil {
    return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *RepositoriesReport) SetFiles(v []FileSummary) {
	o.Files = v
}

// GetInfos returns the Infos field value
func (o *RepositoriesReport) GetInfos() []FileInfoSummary {
	if o == nil {
		var ret []FileInfoSummary
		return ret
	}

	return o.Infos
}

// GetInfosOk returns a tuple with the Infos field value
// and a boolean to check if the value has been set.
func (o *RepositoriesReport) GetInfosOk() ([]FileInfoSummary, bool) {
	if o == nil {
    return nil, false
	}
	return o.Infos, true
}

// SetInfos sets field value
func (o *RepositoriesReport) SetInfos(v []FileInfoSummary) {
	o.Infos = v
}

func (o RepositoriesReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["digest"] = o.Digest
	}
	if true {
		toSerialize["standard-repos"] = o.StandardRepos
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	if true {
		toSerialize["files"] = o.Files
	}
	if true {
		toSerialize["infos"] = o.Infos
	}
	return json.Marshal(toSerialize)
}

type NullableRepositoriesReport struct {
	value *RepositoriesReport
	isSet bool
}

func (v NullableRepositoriesReport) Get() *RepositoriesReport {
	return v.value
}

func (v *NullableRepositoriesReport) Set(val *RepositoriesReport) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositoriesReport) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositoriesReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositoriesReport(val *RepositoriesReport) *NullableRepositoriesReport {
	return &NullableRepositoriesReport{value: val, isSet: true}
}

func (v NullableRepositoriesReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositoriesReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


