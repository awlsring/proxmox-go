# CreateUserRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Userid** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**Enable** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **string** |  | [optional] 
**Keys** | Pointer to **string** |  | [optional] 
**Expire** | Pointer to **float32** |  | [optional] 

## Methods

### NewCreateUserRequestContent

`func NewCreateUserRequestContent(userid string, ) *CreateUserRequestContent`

NewCreateUserRequestContent instantiates a new CreateUserRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRequestContentWithDefaults

`func NewCreateUserRequestContentWithDefaults() *CreateUserRequestContent`

NewCreateUserRequestContentWithDefaults instantiates a new CreateUserRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserid

`func (o *CreateUserRequestContent) GetUserid() string`

GetUserid returns the Userid field if non-nil, zero value otherwise.

### GetUseridOk

`func (o *CreateUserRequestContent) GetUseridOk() (*string, bool)`

GetUseridOk returns a tuple with the Userid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserid

`func (o *CreateUserRequestContent) SetUserid(v string)`

SetUserid sets Userid field to given value.


### GetEmail

`func (o *CreateUserRequestContent) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserRequestContent) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserRequestContent) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateUserRequestContent) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstname

`func (o *CreateUserRequestContent) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *CreateUserRequestContent) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *CreateUserRequestContent) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *CreateUserRequestContent) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *CreateUserRequestContent) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *CreateUserRequestContent) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *CreateUserRequestContent) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *CreateUserRequestContent) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetEnable

`func (o *CreateUserRequestContent) GetEnable() float32`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CreateUserRequestContent) GetEnableOk() (*float32, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CreateUserRequestContent) SetEnable(v float32)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *CreateUserRequestContent) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetComment

`func (o *CreateUserRequestContent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateUserRequestContent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateUserRequestContent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateUserRequestContent) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGroups

`func (o *CreateUserRequestContent) GetGroups() string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *CreateUserRequestContent) GetGroupsOk() (*string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *CreateUserRequestContent) SetGroups(v string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *CreateUserRequestContent) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetKeys

`func (o *CreateUserRequestContent) GetKeys() string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *CreateUserRequestContent) GetKeysOk() (*string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *CreateUserRequestContent) SetKeys(v string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *CreateUserRequestContent) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetExpire

`func (o *CreateUserRequestContent) GetExpire() float32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *CreateUserRequestContent) GetExpireOk() (*float32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *CreateUserRequestContent) SetExpire(v float32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *CreateUserRequestContent) HasExpire() bool`

HasExpire returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


