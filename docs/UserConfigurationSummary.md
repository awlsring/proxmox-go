# UserConfigurationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**Enable** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Keys** | Pointer to **string** |  | [optional] 
**Expire** | Pointer to **float32** |  | [optional] 
**Tokens** | Pointer to [**map[string]UserConfigurationTokenSummary**](UserConfigurationTokenSummary.md) |  | [optional] 

## Methods

### NewUserConfigurationSummary

`func NewUserConfigurationSummary() *UserConfigurationSummary`

NewUserConfigurationSummary instantiates a new UserConfigurationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConfigurationSummaryWithDefaults

`func NewUserConfigurationSummaryWithDefaults() *UserConfigurationSummary`

NewUserConfigurationSummaryWithDefaults instantiates a new UserConfigurationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserConfigurationSummary) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserConfigurationSummary) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserConfigurationSummary) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserConfigurationSummary) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstname

`func (o *UserConfigurationSummary) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserConfigurationSummary) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserConfigurationSummary) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserConfigurationSummary) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *UserConfigurationSummary) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserConfigurationSummary) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserConfigurationSummary) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserConfigurationSummary) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetEnable

`func (o *UserConfigurationSummary) GetEnable() float32`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *UserConfigurationSummary) GetEnableOk() (*float32, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *UserConfigurationSummary) SetEnable(v float32)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *UserConfigurationSummary) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetComment

`func (o *UserConfigurationSummary) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UserConfigurationSummary) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UserConfigurationSummary) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UserConfigurationSummary) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGroups

`func (o *UserConfigurationSummary) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserConfigurationSummary) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserConfigurationSummary) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserConfigurationSummary) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetKeys

`func (o *UserConfigurationSummary) GetKeys() string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *UserConfigurationSummary) GetKeysOk() (*string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *UserConfigurationSummary) SetKeys(v string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *UserConfigurationSummary) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetExpire

`func (o *UserConfigurationSummary) GetExpire() float32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *UserConfigurationSummary) GetExpireOk() (*float32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *UserConfigurationSummary) SetExpire(v float32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *UserConfigurationSummary) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetTokens

`func (o *UserConfigurationSummary) GetTokens() map[string]UserConfigurationTokenSummary`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *UserConfigurationSummary) GetTokensOk() (*map[string]UserConfigurationTokenSummary, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *UserConfigurationSummary) SetTokens(v map[string]UserConfigurationTokenSummary)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *UserConfigurationSummary) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


