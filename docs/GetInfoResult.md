# GetInfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**AgentInfoSummary**](AgentInfoSummary.md) |  | [optional] 

## Methods

### NewGetInfoResult

`func NewGetInfoResult() *GetInfoResult`

NewGetInfoResult instantiates a new GetInfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInfoResultWithDefaults

`func NewGetInfoResultWithDefaults() *GetInfoResult`

NewGetInfoResultWithDefaults instantiates a new GetInfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetInfoResult) GetResult() AgentInfoSummary`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetInfoResult) GetResultOk() (*AgentInfoSummary, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetInfoResult) SetResult(v AgentInfoSummary)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetInfoResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


