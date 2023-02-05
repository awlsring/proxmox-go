# NodeStorageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Storage** | **string** |  | 
**Type** | [**StorageType**](StorageType.md) |  | 
**Content** | **string** | The allowed storage content types. | 
**Active** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Enabled** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Shared** | Pointer to **float32** | Whether the storage is shared across all nodes. | [optional] 
**Total** | Pointer to **float32** |  | [optional] 
**Used** | Pointer to **float32** |  | [optional] 
**UsedFraction** | Pointer to **float32** |  | [optional] 
**Avail** | Pointer to **float32** |  | [optional] 

## Methods

### NewNodeStorageSummary

`func NewNodeStorageSummary(storage string, type_ StorageType, content string, ) *NodeStorageSummary`

NewNodeStorageSummary instantiates a new NodeStorageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeStorageSummaryWithDefaults

`func NewNodeStorageSummaryWithDefaults() *NodeStorageSummary`

NewNodeStorageSummaryWithDefaults instantiates a new NodeStorageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorage

`func (o *NodeStorageSummary) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *NodeStorageSummary) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *NodeStorageSummary) SetStorage(v string)`

SetStorage sets Storage field to given value.


### GetType

`func (o *NodeStorageSummary) GetType() StorageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeStorageSummary) GetTypeOk() (*StorageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeStorageSummary) SetType(v StorageType)`

SetType sets Type field to given value.


### GetContent

`func (o *NodeStorageSummary) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NodeStorageSummary) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NodeStorageSummary) SetContent(v string)`

SetContent sets Content field to given value.


### GetActive

`func (o *NodeStorageSummary) GetActive() float32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NodeStorageSummary) GetActiveOk() (*float32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NodeStorageSummary) SetActive(v float32)`

SetActive sets Active field to given value.

### HasActive

`func (o *NodeStorageSummary) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEnabled

`func (o *NodeStorageSummary) GetEnabled() float32`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NodeStorageSummary) GetEnabledOk() (*float32, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NodeStorageSummary) SetEnabled(v float32)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NodeStorageSummary) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetShared

`func (o *NodeStorageSummary) GetShared() float32`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *NodeStorageSummary) GetSharedOk() (*float32, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *NodeStorageSummary) SetShared(v float32)`

SetShared sets Shared field to given value.

### HasShared

`func (o *NodeStorageSummary) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetTotal

`func (o *NodeStorageSummary) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NodeStorageSummary) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NodeStorageSummary) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *NodeStorageSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *NodeStorageSummary) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *NodeStorageSummary) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *NodeStorageSummary) SetUsed(v float32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *NodeStorageSummary) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUsedFraction

`func (o *NodeStorageSummary) GetUsedFraction() float32`

GetUsedFraction returns the UsedFraction field if non-nil, zero value otherwise.

### GetUsedFractionOk

`func (o *NodeStorageSummary) GetUsedFractionOk() (*float32, bool)`

GetUsedFractionOk returns a tuple with the UsedFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedFraction

`func (o *NodeStorageSummary) SetUsedFraction(v float32)`

SetUsedFraction sets UsedFraction field to given value.

### HasUsedFraction

`func (o *NodeStorageSummary) HasUsedFraction() bool`

HasUsedFraction returns a boolean if a field has been set.

### GetAvail

`func (o *NodeStorageSummary) GetAvail() float32`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *NodeStorageSummary) GetAvailOk() (*float32, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *NodeStorageSummary) SetAvail(v float32)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *NodeStorageSummary) HasAvail() bool`

HasAvail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


