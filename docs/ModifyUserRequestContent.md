# ModifyUserRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**Append** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Enable** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **string** |  | [optional] 
**Keys** | Pointer to **string** |  | [optional] 
**Expire** | Pointer to **float32** |  | [optional] 

## Methods

### NewModifyUserRequestContent

`func NewModifyUserRequestContent() *ModifyUserRequestContent`

NewModifyUserRequestContent instantiates a new ModifyUserRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyUserRequestContentWithDefaults

`func NewModifyUserRequestContentWithDefaults() *ModifyUserRequestContent`

NewModifyUserRequestContentWithDefaults instantiates a new ModifyUserRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ModifyUserRequestContent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModifyUserRequestContent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModifyUserRequestContent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModifyUserRequestContent) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAppend

`func (o *ModifyUserRequestContent) GetAppend() float32`

GetAppend returns the Append field if non-nil, zero value otherwise.

### GetAppendOk

`func (o *ModifyUserRequestContent) GetAppendOk() (*float32, bool)`

GetAppendOk returns a tuple with the Append field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppend

`func (o *ModifyUserRequestContent) SetAppend(v float32)`

SetAppend sets Append field to given value.

### HasAppend

`func (o *ModifyUserRequestContent) HasAppend() bool`

HasAppend returns a boolean if a field has been set.

### GetEmail

`func (o *ModifyUserRequestContent) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModifyUserRequestContent) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModifyUserRequestContent) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ModifyUserRequestContent) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnable

`func (o *ModifyUserRequestContent) GetEnable() float32`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ModifyUserRequestContent) GetEnableOk() (*float32, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ModifyUserRequestContent) SetEnable(v float32)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ModifyUserRequestContent) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetFirstname

`func (o *ModifyUserRequestContent) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *ModifyUserRequestContent) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *ModifyUserRequestContent) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *ModifyUserRequestContent) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *ModifyUserRequestContent) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *ModifyUserRequestContent) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *ModifyUserRequestContent) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *ModifyUserRequestContent) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetGroups

`func (o *ModifyUserRequestContent) GetGroups() string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ModifyUserRequestContent) GetGroupsOk() (*string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ModifyUserRequestContent) SetGroups(v string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ModifyUserRequestContent) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetKeys

`func (o *ModifyUserRequestContent) GetKeys() string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ModifyUserRequestContent) GetKeysOk() (*string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ModifyUserRequestContent) SetKeys(v string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ModifyUserRequestContent) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetExpire

`func (o *ModifyUserRequestContent) GetExpire() float32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *ModifyUserRequestContent) GetExpireOk() (*float32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *ModifyUserRequestContent) SetExpire(v float32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *ModifyUserRequestContent) HasExpire() bool`

HasExpire returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


