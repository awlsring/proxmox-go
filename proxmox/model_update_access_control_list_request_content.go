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

// checks if the UpdateAccessControlListRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAccessControlListRequestContent{}

// UpdateAccessControlListRequestContent struct for UpdateAccessControlListRequestContent
type UpdateAccessControlListRequestContent struct {
	Path string `json:"path"`
	Roles string `json:"roles"`
	// Removed permissions specified in request
	Delete *bool `json:"delete,omitempty"`
	// Comma separated list of groups
	Groups *string `json:"groups,omitempty"`
	// Allow propegation of permissions.
	Propagate *bool `json:"propagate,omitempty"`
	// Comma separated list of tokens
	Tokens *string `json:"tokens,omitempty"`
	// Comma separated list of users
	Users *string `json:"users,omitempty"`
}

// NewUpdateAccessControlListRequestContent instantiates a new UpdateAccessControlListRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccessControlListRequestContent(path string, roles string) *UpdateAccessControlListRequestContent {
	this := UpdateAccessControlListRequestContent{}
	this.Path = path
	this.Roles = roles
	return &this
}

// NewUpdateAccessControlListRequestContentWithDefaults instantiates a new UpdateAccessControlListRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccessControlListRequestContentWithDefaults() *UpdateAccessControlListRequestContent {
	this := UpdateAccessControlListRequestContent{}
	return &this
}

// GetPath returns the Path field value
func (o *UpdateAccessControlListRequestContent) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *UpdateAccessControlListRequestContent) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *UpdateAccessControlListRequestContent) SetPath(v string) {
	o.Path = v
}

// GetRoles returns the Roles field value
func (o *UpdateAccessControlListRequestContent) GetRoles() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UpdateAccessControlListRequestContent) GetRolesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *UpdateAccessControlListRequestContent) SetRoles(v string) {
	o.Roles = v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *UpdateAccessControlListRequestContent) GetDelete() bool {
	if o == nil || IsNil(o.Delete) {
		var ret bool
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccessControlListRequestContent) GetDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *UpdateAccessControlListRequestContent) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given bool and assigns it to the Delete field.
func (o *UpdateAccessControlListRequestContent) SetDelete(v bool) {
	o.Delete = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *UpdateAccessControlListRequestContent) GetGroups() string {
	if o == nil || IsNil(o.Groups) {
		var ret string
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccessControlListRequestContent) GetGroupsOk() (*string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *UpdateAccessControlListRequestContent) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given string and assigns it to the Groups field.
func (o *UpdateAccessControlListRequestContent) SetGroups(v string) {
	o.Groups = &v
}

// GetPropagate returns the Propagate field value if set, zero value otherwise.
func (o *UpdateAccessControlListRequestContent) GetPropagate() bool {
	if o == nil || IsNil(o.Propagate) {
		var ret bool
		return ret
	}
	return *o.Propagate
}

// GetPropagateOk returns a tuple with the Propagate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccessControlListRequestContent) GetPropagateOk() (*bool, bool) {
	if o == nil || IsNil(o.Propagate) {
		return nil, false
	}
	return o.Propagate, true
}

// HasPropagate returns a boolean if a field has been set.
func (o *UpdateAccessControlListRequestContent) HasPropagate() bool {
	if o != nil && !IsNil(o.Propagate) {
		return true
	}

	return false
}

// SetPropagate gets a reference to the given bool and assigns it to the Propagate field.
func (o *UpdateAccessControlListRequestContent) SetPropagate(v bool) {
	o.Propagate = &v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *UpdateAccessControlListRequestContent) GetTokens() string {
	if o == nil || IsNil(o.Tokens) {
		var ret string
		return ret
	}
	return *o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccessControlListRequestContent) GetTokensOk() (*string, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *UpdateAccessControlListRequestContent) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given string and assigns it to the Tokens field.
func (o *UpdateAccessControlListRequestContent) SetTokens(v string) {
	o.Tokens = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *UpdateAccessControlListRequestContent) GetUsers() string {
	if o == nil || IsNil(o.Users) {
		var ret string
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccessControlListRequestContent) GetUsersOk() (*string, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UpdateAccessControlListRequestContent) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given string and assigns it to the Users field.
func (o *UpdateAccessControlListRequestContent) SetUsers(v string) {
	o.Users = &v
}

func (o UpdateAccessControlListRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAccessControlListRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["roles"] = o.Roles
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Propagate) {
		toSerialize["propagate"] = o.Propagate
	}
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableUpdateAccessControlListRequestContent struct {
	value *UpdateAccessControlListRequestContent
	isSet bool
}

func (v NullableUpdateAccessControlListRequestContent) Get() *UpdateAccessControlListRequestContent {
	return v.value
}

func (v *NullableUpdateAccessControlListRequestContent) Set(val *UpdateAccessControlListRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccessControlListRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccessControlListRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccessControlListRequestContent(val *UpdateAccessControlListRequestContent) *NullableUpdateAccessControlListRequestContent {
	return &NullableUpdateAccessControlListRequestContent{value: val, isSet: true}
}

func (v NullableUpdateAccessControlListRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccessControlListRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


