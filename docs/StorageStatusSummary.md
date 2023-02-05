# StorageStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**StorageType**](StorageType.md) |  | [optional] 
**Shared** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Active** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Enabled** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Total** | Pointer to **float32** |  | [optional] 
**Used** | Pointer to **float32** |  | [optional] 
**Avail** | Pointer to **float32** |  | [optional] 

## Methods

### NewStorageStatusSummary

`func NewStorageStatusSummary() *StorageStatusSummary`

NewStorageStatusSummary instantiates a new StorageStatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageStatusSummaryWithDefaults

`func NewStorageStatusSummaryWithDefaults() *StorageStatusSummary`

NewStorageStatusSummaryWithDefaults instantiates a new StorageStatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *StorageStatusSummary) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *StorageStatusSummary) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *StorageStatusSummary) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *StorageStatusSummary) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetType

`func (o *StorageStatusSummary) GetType() StorageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageStatusSummary) GetTypeOk() (*StorageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageStatusSummary) SetType(v StorageType)`

SetType sets Type field to given value.

### HasType

`func (o *StorageStatusSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetShared

`func (o *StorageStatusSummary) GetShared() float32`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *StorageStatusSummary) GetSharedOk() (*float32, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *StorageStatusSummary) SetShared(v float32)`

SetShared sets Shared field to given value.

### HasShared

`func (o *StorageStatusSummary) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetActive

`func (o *StorageStatusSummary) GetActive() float32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StorageStatusSummary) GetActiveOk() (*float32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StorageStatusSummary) SetActive(v float32)`

SetActive sets Active field to given value.

### HasActive

`func (o *StorageStatusSummary) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEnabled

`func (o *StorageStatusSummary) GetEnabled() float32`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageStatusSummary) GetEnabledOk() (*float32, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageStatusSummary) SetEnabled(v float32)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageStatusSummary) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTotal

`func (o *StorageStatusSummary) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StorageStatusSummary) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StorageStatusSummary) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *StorageStatusSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *StorageStatusSummary) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *StorageStatusSummary) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *StorageStatusSummary) SetUsed(v float32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *StorageStatusSummary) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetAvail

`func (o *StorageStatusSummary) GetAvail() float32`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *StorageStatusSummary) GetAvailOk() (*float32, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *StorageStatusSummary) SetAvail(v float32)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *StorageStatusSummary) HasAvail() bool`

HasAvail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


