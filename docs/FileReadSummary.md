# FileReadSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**Truncated** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 

## Methods

### NewFileReadSummary

`func NewFileReadSummary(content string, ) *FileReadSummary`

NewFileReadSummary instantiates a new FileReadSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileReadSummaryWithDefaults

`func NewFileReadSummaryWithDefaults() *FileReadSummary`

NewFileReadSummaryWithDefaults instantiates a new FileReadSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *FileReadSummary) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FileReadSummary) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FileReadSummary) SetContent(v string)`

SetContent sets Content field to given value.


### GetTruncated

`func (o *FileReadSummary) GetTruncated() float32`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *FileReadSummary) GetTruncatedOk() (*float32, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *FileReadSummary) SetTruncated(v float32)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *FileReadSummary) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


