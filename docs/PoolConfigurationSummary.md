# PoolConfigurationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]PoolMemberSummary**](PoolMemberSummary.md) |  | 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewPoolConfigurationSummary

`func NewPoolConfigurationSummary(members []PoolMemberSummary, ) *PoolConfigurationSummary`

NewPoolConfigurationSummary instantiates a new PoolConfigurationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolConfigurationSummaryWithDefaults

`func NewPoolConfigurationSummaryWithDefaults() *PoolConfigurationSummary`

NewPoolConfigurationSummaryWithDefaults instantiates a new PoolConfigurationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *PoolConfigurationSummary) GetMembers() []PoolMemberSummary`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *PoolConfigurationSummary) GetMembersOk() (*[]PoolMemberSummary, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *PoolConfigurationSummary) SetMembers(v []PoolMemberSummary)`

SetMembers sets Members field to given value.


### GetComment

`func (o *PoolConfigurationSummary) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PoolConfigurationSummary) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PoolConfigurationSummary) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PoolConfigurationSummary) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


