# GroupDetailsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | **[]string** |  | 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupDetailsSummary

`func NewGroupDetailsSummary(users []string, ) *GroupDetailsSummary`

NewGroupDetailsSummary instantiates a new GroupDetailsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupDetailsSummaryWithDefaults

`func NewGroupDetailsSummaryWithDefaults() *GroupDetailsSummary`

NewGroupDetailsSummaryWithDefaults instantiates a new GroupDetailsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GroupDetailsSummary) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupDetailsSummary) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupDetailsSummary) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetComment

`func (o *GroupDetailsSummary) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GroupDetailsSummary) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GroupDetailsSummary) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GroupDetailsSummary) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


