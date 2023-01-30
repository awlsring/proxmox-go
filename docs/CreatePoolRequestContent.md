# CreatePoolRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Poolid** | **string** |  | 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewCreatePoolRequestContent

`func NewCreatePoolRequestContent(poolid string, ) *CreatePoolRequestContent`

NewCreatePoolRequestContent instantiates a new CreatePoolRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePoolRequestContentWithDefaults

`func NewCreatePoolRequestContentWithDefaults() *CreatePoolRequestContent`

NewCreatePoolRequestContentWithDefaults instantiates a new CreatePoolRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolid

`func (o *CreatePoolRequestContent) GetPoolid() string`

GetPoolid returns the Poolid field if non-nil, zero value otherwise.

### GetPoolidOk

`func (o *CreatePoolRequestContent) GetPoolidOk() (*string, bool)`

GetPoolidOk returns a tuple with the Poolid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolid

`func (o *CreatePoolRequestContent) SetPoolid(v string)`

SetPoolid sets Poolid field to given value.


### GetComment

`func (o *CreatePoolRequestContent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreatePoolRequestContent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreatePoolRequestContent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreatePoolRequestContent) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


