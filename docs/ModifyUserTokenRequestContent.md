# ModifyUserTokenRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privsep** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Expire** | Pointer to **float32** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewModifyUserTokenRequestContent

`func NewModifyUserTokenRequestContent() *ModifyUserTokenRequestContent`

NewModifyUserTokenRequestContent instantiates a new ModifyUserTokenRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyUserTokenRequestContentWithDefaults

`func NewModifyUserTokenRequestContentWithDefaults() *ModifyUserTokenRequestContent`

NewModifyUserTokenRequestContentWithDefaults instantiates a new ModifyUserTokenRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivsep

`func (o *ModifyUserTokenRequestContent) GetPrivsep() float32`

GetPrivsep returns the Privsep field if non-nil, zero value otherwise.

### GetPrivsepOk

`func (o *ModifyUserTokenRequestContent) GetPrivsepOk() (*float32, bool)`

GetPrivsepOk returns a tuple with the Privsep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivsep

`func (o *ModifyUserTokenRequestContent) SetPrivsep(v float32)`

SetPrivsep sets Privsep field to given value.

### HasPrivsep

`func (o *ModifyUserTokenRequestContent) HasPrivsep() bool`

HasPrivsep returns a boolean if a field has been set.

### GetExpire

`func (o *ModifyUserTokenRequestContent) GetExpire() float32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *ModifyUserTokenRequestContent) GetExpireOk() (*float32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *ModifyUserTokenRequestContent) SetExpire(v float32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *ModifyUserTokenRequestContent) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetComment

`func (o *ModifyUserTokenRequestContent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModifyUserTokenRequestContent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModifyUserTokenRequestContent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModifyUserTokenRequestContent) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


