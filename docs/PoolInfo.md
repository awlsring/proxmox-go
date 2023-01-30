# PoolInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Poolid** | **string** |  | 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewPoolInfo

`func NewPoolInfo(poolid string, ) *PoolInfo`

NewPoolInfo instantiates a new PoolInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolInfoWithDefaults

`func NewPoolInfoWithDefaults() *PoolInfo`

NewPoolInfoWithDefaults instantiates a new PoolInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolid

`func (o *PoolInfo) GetPoolid() string`

GetPoolid returns the Poolid field if non-nil, zero value otherwise.

### GetPoolidOk

`func (o *PoolInfo) GetPoolidOk() (*string, bool)`

GetPoolidOk returns a tuple with the Poolid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolid

`func (o *PoolInfo) SetPoolid(v string)`

SetPoolid sets Poolid field to given value.


### GetComment

`func (o *PoolInfo) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PoolInfo) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PoolInfo) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PoolInfo) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


