# UpdateAccessControlListRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**Roles** | **string** |  | 
**Delete** | Pointer to **bool** | Removed permissions specified in request | [optional] 
**Groups** | Pointer to **string** | Comma separated list of groups | [optional] 
**Propagate** | Pointer to **bool** | Allow propegation of permissions. | [optional] 
**Tokens** | Pointer to **string** | Comma separated list of tokens | [optional] 
**Users** | Pointer to **string** | Comma separated list of users | [optional] 

## Methods

### NewUpdateAccessControlListRequestContent

`func NewUpdateAccessControlListRequestContent(path string, roles string, ) *UpdateAccessControlListRequestContent`

NewUpdateAccessControlListRequestContent instantiates a new UpdateAccessControlListRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccessControlListRequestContentWithDefaults

`func NewUpdateAccessControlListRequestContentWithDefaults() *UpdateAccessControlListRequestContent`

NewUpdateAccessControlListRequestContentWithDefaults instantiates a new UpdateAccessControlListRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *UpdateAccessControlListRequestContent) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *UpdateAccessControlListRequestContent) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *UpdateAccessControlListRequestContent) SetPath(v string)`

SetPath sets Path field to given value.


### GetRoles

`func (o *UpdateAccessControlListRequestContent) GetRoles() string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateAccessControlListRequestContent) GetRolesOk() (*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateAccessControlListRequestContent) SetRoles(v string)`

SetRoles sets Roles field to given value.


### GetDelete

`func (o *UpdateAccessControlListRequestContent) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *UpdateAccessControlListRequestContent) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *UpdateAccessControlListRequestContent) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *UpdateAccessControlListRequestContent) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetGroups

`func (o *UpdateAccessControlListRequestContent) GetGroups() string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UpdateAccessControlListRequestContent) GetGroupsOk() (*string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UpdateAccessControlListRequestContent) SetGroups(v string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UpdateAccessControlListRequestContent) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetPropagate

`func (o *UpdateAccessControlListRequestContent) GetPropagate() bool`

GetPropagate returns the Propagate field if non-nil, zero value otherwise.

### GetPropagateOk

`func (o *UpdateAccessControlListRequestContent) GetPropagateOk() (*bool, bool)`

GetPropagateOk returns a tuple with the Propagate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagate

`func (o *UpdateAccessControlListRequestContent) SetPropagate(v bool)`

SetPropagate sets Propagate field to given value.

### HasPropagate

`func (o *UpdateAccessControlListRequestContent) HasPropagate() bool`

HasPropagate returns a boolean if a field has been set.

### GetTokens

`func (o *UpdateAccessControlListRequestContent) GetTokens() string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *UpdateAccessControlListRequestContent) GetTokensOk() (*string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *UpdateAccessControlListRequestContent) SetTokens(v string)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *UpdateAccessControlListRequestContent) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetUsers

`func (o *UpdateAccessControlListRequestContent) GetUsers() string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateAccessControlListRequestContent) GetUsersOk() (*string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateAccessControlListRequestContent) SetUsers(v string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UpdateAccessControlListRequestContent) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


