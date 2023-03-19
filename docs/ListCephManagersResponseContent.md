# ListCephManagersResponseContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CephMonitorSummary**](CephMonitorSummary.md) |  | 

## Methods

### NewListCephManagersResponseContent

`func NewListCephManagersResponseContent(data []CephMonitorSummary, ) *ListCephManagersResponseContent`

NewListCephManagersResponseContent instantiates a new ListCephManagersResponseContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCephManagersResponseContentWithDefaults

`func NewListCephManagersResponseContentWithDefaults() *ListCephManagersResponseContent`

NewListCephManagersResponseContentWithDefaults instantiates a new ListCephManagersResponseContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCephManagersResponseContent) GetData() []CephMonitorSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCephManagersResponseContent) GetDataOk() (*[]CephMonitorSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCephManagersResponseContent) SetData(v []CephMonitorSummary)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


