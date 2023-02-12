# GroupSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groupid** | **string** |  | 
**Users** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupSummary

`func NewGroupSummary(groupid string, ) *GroupSummary`

NewGroupSummary instantiates a new GroupSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSummaryWithDefaults

`func NewGroupSummaryWithDefaults() *GroupSummary`

NewGroupSummaryWithDefaults instantiates a new GroupSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupid

`func (o *GroupSummary) GetGroupid() string`

GetGroupid returns the Groupid field if non-nil, zero value otherwise.

### GetGroupidOk

`func (o *GroupSummary) GetGroupidOk() (*string, bool)`

GetGroupidOk returns a tuple with the Groupid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupid

`func (o *GroupSummary) SetGroupid(v string)`

SetGroupid sets Groupid field to given value.


### GetUsers

`func (o *GroupSummary) GetUsers() string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupSummary) GetUsersOk() (*string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupSummary) SetUsers(v string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GroupSummary) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetComment

`func (o *GroupSummary) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GroupSummary) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GroupSummary) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GroupSummary) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


