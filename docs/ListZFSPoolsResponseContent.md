# ListZFSPoolsResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ZFSPoolSummary**](ZFSPoolSummary.md) |  | 

## Methods

### NewListZFSPoolsResponseContent

`func NewListZFSPoolsResponseContent(data []ZFSPoolSummary, ) *ListZFSPoolsResponseContent`

NewListZFSPoolsResponseContent instantiates a new ListZFSPoolsResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListZFSPoolsResponseContentWithDefaults

`func NewListZFSPoolsResponseContentWithDefaults() *ListZFSPoolsResponseContent`

NewListZFSPoolsResponseContentWithDefaults instantiates a new ListZFSPoolsResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListZFSPoolsResponseContent) GetData() []ZFSPoolSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListZFSPoolsResponseContent) GetDataOk() (*[]ZFSPoolSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListZFSPoolsResponseContent) SetData(v []ZFSPoolSummary)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


