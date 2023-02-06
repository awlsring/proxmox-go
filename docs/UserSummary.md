# UserSummary

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
**Expire** | Pointer to **float32** |  | [optional] 
**RealmType** | Pointer to **string** |  | [optional] 
**Tokens** | Pointer to [**[]UserTokenSummary**](UserTokenSummary.md) |  | [optional] 

## Methods

### NewUserSummary

`func NewUserSummary(userid string, ) *UserSummary`

NewUserSummary instantiates a new UserSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSummaryWithDefaults

`func NewUserSummaryWithDefaults() *UserSummary`

NewUserSummaryWithDefaults instantiates a new UserSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserid

`func (o *UserSummary) GetUserid() string`

GetUserid returns the Userid field if non-nil, zero value otherwise.

### GetUseridOk

`func (o *UserSummary) GetUseridOk() (*string, bool)`

GetUseridOk returns a tuple with the Userid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserid

`func (o *UserSummary) SetUserid(v string)`

SetUserid sets Userid field to given value.


### GetEmail

`func (o *UserSummary) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSummary) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSummary) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserSummary) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstname

`func (o *UserSummary) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserSummary) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserSummary) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserSummary) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *UserSummary) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserSummary) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserSummary) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserSummary) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetEnable

`func (o *UserSummary) GetEnable() float32`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *UserSummary) GetEnableOk() (*float32, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *UserSummary) SetEnable(v float32)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *UserSummary) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetComment

`func (o *UserSummary) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UserSummary) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UserSummary) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UserSummary) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGroups

`func (o *UserSummary) GetGroups() string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserSummary) GetGroupsOk() (*string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserSummary) SetGroups(v string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserSummary) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetExpire

`func (o *UserSummary) GetExpire() float32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *UserSummary) GetExpireOk() (*float32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *UserSummary) SetExpire(v float32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *UserSummary) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetRealmType

`func (o *UserSummary) GetRealmType() string`

GetRealmType returns the RealmType field if non-nil, zero value otherwise.

### GetRealmTypeOk

`func (o *UserSummary) GetRealmTypeOk() (*string, bool)`

GetRealmTypeOk returns a tuple with the RealmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmType

`func (o *UserSummary) SetRealmType(v string)`

SetRealmType sets RealmType field to given value.

### HasRealmType

`func (o *UserSummary) HasRealmType() bool`

HasRealmType returns a boolean if a field has been set.

### GetTokens

`func (o *UserSummary) GetTokens() []UserTokenSummary`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *UserSummary) GetTokensOk() (*[]UserTokenSummary, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *UserSummary) SetTokens(v []UserTokenSummary)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *UserSummary) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


