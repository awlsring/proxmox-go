# UserConfigurationTokenSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privsep** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Expire** | Pointer to **float32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewUserConfigurationTokenSummary

`func NewUserConfigurationTokenSummary() *UserConfigurationTokenSummary`

NewUserConfigurationTokenSummary instantiates a new UserConfigurationTokenSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConfigurationTokenSummaryWithDefaults

`func NewUserConfigurationTokenSummaryWithDefaults() *UserConfigurationTokenSummary`

NewUserConfigurationTokenSummaryWithDefaults instantiates a new UserConfigurationTokenSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivsep

`func (o *UserConfigurationTokenSummary) GetPrivsep() float32`

GetPrivsep returns the Privsep field if non-nil, zero value otherwise.

### GetPrivsepOk

`func (o *UserConfigurationTokenSummary) GetPrivsepOk() (*float32, bool)`

GetPrivsepOk returns a tuple with the Privsep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivsep

`func (o *UserConfigurationTokenSummary) SetPrivsep(v float32)`

SetPrivsep sets Privsep field to given value.

### HasPrivsep

`func (o *UserConfigurationTokenSummary) HasPrivsep() bool`

HasPrivsep returns a boolean if a field has been set.

### GetExpire

`func (o *UserConfigurationTokenSummary) GetExpire() float32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *UserConfigurationTokenSummary) GetExpireOk() (*float32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *UserConfigurationTokenSummary) SetExpire(v float32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *UserConfigurationTokenSummary) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetComment

`func (o *UserConfigurationTokenSummary) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UserConfigurationTokenSummary) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UserConfigurationTokenSummary) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UserConfigurationTokenSummary) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


