# CreateGroupRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groupid** | **string** |  | 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateGroupRequestContent

`func NewCreateGroupRequestContent(groupid string, ) *CreateGroupRequestContent`

NewCreateGroupRequestContent instantiates a new CreateGroupRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupRequestContentWithDefaults

`func NewCreateGroupRequestContentWithDefaults() *CreateGroupRequestContent`

NewCreateGroupRequestContentWithDefaults instantiates a new CreateGroupRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupid

`func (o *CreateGroupRequestContent) GetGroupid() string`

GetGroupid returns the Groupid field if non-nil, zero value otherwise.

### GetGroupidOk

`func (o *CreateGroupRequestContent) GetGroupidOk() (*string, bool)`

GetGroupidOk returns a tuple with the Groupid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupid

`func (o *CreateGroupRequestContent) SetGroupid(v string)`

SetGroupid sets Groupid field to given value.


### GetComment

`func (o *CreateGroupRequestContent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateGroupRequestContent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateGroupRequestContent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateGroupRequestContent) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


