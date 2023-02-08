# ModifyPoolRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**Delete** | Pointer to **bool** | Remove storage and vms in request rather than adding. | [optional] 
**Storage** | Pointer to **string** | List of storage identifiers to add to the pool. | [optional] 
**Vms** | Pointer to **string** | List of VM identifiers to add to the pool. | [optional] 

## Methods

### NewModifyPoolRequestContent

`func NewModifyPoolRequestContent() *ModifyPoolRequestContent`

NewModifyPoolRequestContent instantiates a new ModifyPoolRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyPoolRequestContentWithDefaults

`func NewModifyPoolRequestContentWithDefaults() *ModifyPoolRequestContent`

NewModifyPoolRequestContentWithDefaults instantiates a new ModifyPoolRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ModifyPoolRequestContent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModifyPoolRequestContent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModifyPoolRequestContent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModifyPoolRequestContent) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDelete

`func (o *ModifyPoolRequestContent) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ModifyPoolRequestContent) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ModifyPoolRequestContent) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ModifyPoolRequestContent) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetStorage

`func (o *ModifyPoolRequestContent) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ModifyPoolRequestContent) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ModifyPoolRequestContent) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ModifyPoolRequestContent) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetVms

`func (o *ModifyPoolRequestContent) GetVms() string`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *ModifyPoolRequestContent) GetVmsOk() (*string, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *ModifyPoolRequestContent) SetVms(v string)`

SetVms sets Vms field to given value.

### HasVms

`func (o *ModifyPoolRequestContent) HasVms() bool`

HasVms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


