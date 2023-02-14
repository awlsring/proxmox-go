# ModifyRoleRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privs** | Pointer to **string** |  | [optional] 
**Append** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 

## Methods

### NewModifyRoleRequestContent

`func NewModifyRoleRequestContent() *ModifyRoleRequestContent`

NewModifyRoleRequestContent instantiates a new ModifyRoleRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyRoleRequestContentWithDefaults

`func NewModifyRoleRequestContentWithDefaults() *ModifyRoleRequestContent`

NewModifyRoleRequestContentWithDefaults instantiates a new ModifyRoleRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivs

`func (o *ModifyRoleRequestContent) GetPrivs() string`

GetPrivs returns the Privs field if non-nil, zero value otherwise.

### GetPrivsOk

`func (o *ModifyRoleRequestContent) GetPrivsOk() (*string, bool)`

GetPrivsOk returns a tuple with the Privs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivs

`func (o *ModifyRoleRequestContent) SetPrivs(v string)`

SetPrivs sets Privs field to given value.

### HasPrivs

`func (o *ModifyRoleRequestContent) HasPrivs() bool`

HasPrivs returns a boolean if a field has been set.

### GetAppend

`func (o *ModifyRoleRequestContent) GetAppend() float32`

GetAppend returns the Append field if non-nil, zero value otherwise.

### GetAppendOk

`func (o *ModifyRoleRequestContent) GetAppendOk() (*float32, bool)`

GetAppendOk returns a tuple with the Append field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppend

`func (o *ModifyRoleRequestContent) SetAppend(v float32)`

SetAppend sets Append field to given value.

### HasAppend

`func (o *ModifyRoleRequestContent) HasAppend() bool`

HasAppend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


