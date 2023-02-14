# CreateRoleRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roleid** | **string** |  | 
**Privs** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateRoleRequestContent

`func NewCreateRoleRequestContent(roleid string, ) *CreateRoleRequestContent`

NewCreateRoleRequestContent instantiates a new CreateRoleRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleRequestContentWithDefaults

`func NewCreateRoleRequestContentWithDefaults() *CreateRoleRequestContent`

NewCreateRoleRequestContentWithDefaults instantiates a new CreateRoleRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleid

`func (o *CreateRoleRequestContent) GetRoleid() string`

GetRoleid returns the Roleid field if non-nil, zero value otherwise.

### GetRoleidOk

`func (o *CreateRoleRequestContent) GetRoleidOk() (*string, bool)`

GetRoleidOk returns a tuple with the Roleid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleid

`func (o *CreateRoleRequestContent) SetRoleid(v string)`

SetRoleid sets Roleid field to given value.


### GetPrivs

`func (o *CreateRoleRequestContent) GetPrivs() string`

GetPrivs returns the Privs field if non-nil, zero value otherwise.

### GetPrivsOk

`func (o *CreateRoleRequestContent) GetPrivsOk() (*string, bool)`

GetPrivsOk returns a tuple with the Privs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivs

`func (o *CreateRoleRequestContent) SetPrivs(v string)`

SetPrivs sets Privs field to given value.

### HasPrivs

`func (o *CreateRoleRequestContent) HasPrivs() bool`

HasPrivs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


