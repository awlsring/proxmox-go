# CreateUserTokenRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privsep** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Expire** | Pointer to **float32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateUserTokenRequestContent

`func NewCreateUserTokenRequestContent() *CreateUserTokenRequestContent`

NewCreateUserTokenRequestContent instantiates a new CreateUserTokenRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserTokenRequestContentWithDefaults

`func NewCreateUserTokenRequestContentWithDefaults() *CreateUserTokenRequestContent`

NewCreateUserTokenRequestContentWithDefaults instantiates a new CreateUserTokenRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivsep

`func (o *CreateUserTokenRequestContent) GetPrivsep() float32`

GetPrivsep returns the Privsep field if non-nil, zero value otherwise.

### GetPrivsepOk

`func (o *CreateUserTokenRequestContent) GetPrivsepOk() (*float32, bool)`

GetPrivsepOk returns a tuple with the Privsep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivsep

`func (o *CreateUserTokenRequestContent) SetPrivsep(v float32)`

SetPrivsep sets Privsep field to given value.

### HasPrivsep

`func (o *CreateUserTokenRequestContent) HasPrivsep() bool`

HasPrivsep returns a boolean if a field has been set.

### GetExpire

`func (o *CreateUserTokenRequestContent) GetExpire() float32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *CreateUserTokenRequestContent) GetExpireOk() (*float32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *CreateUserTokenRequestContent) SetExpire(v float32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *CreateUserTokenRequestContent) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetComment

`func (o *CreateUserTokenRequestContent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateUserTokenRequestContent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateUserTokenRequestContent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateUserTokenRequestContent) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


