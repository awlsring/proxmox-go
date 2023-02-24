# SnapshotSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The snapshot name. | 
**Description** | **string** | The snapshot description. | 
**Parent** | Pointer to **string** | The snapshot name. | [optional] 
**Snaptime** | Pointer to **float32** | The snapshot time. | [optional] 
**VmState** | Pointer to **bool** | If the snapshot includes RAM. | [optional] 

## Methods

### NewSnapshotSummary

`func NewSnapshotSummary(name string, description string, ) *SnapshotSummary`

NewSnapshotSummary instantiates a new SnapshotSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotSummaryWithDefaults

`func NewSnapshotSummaryWithDefaults() *SnapshotSummary`

NewSnapshotSummaryWithDefaults instantiates a new SnapshotSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SnapshotSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnapshotSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnapshotSummary) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SnapshotSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotSummary) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetParent

`func (o *SnapshotSummary) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *SnapshotSummary) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *SnapshotSummary) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *SnapshotSummary) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetSnaptime

`func (o *SnapshotSummary) GetSnaptime() float32`

GetSnaptime returns the Snaptime field if non-nil, zero value otherwise.

### GetSnaptimeOk

`func (o *SnapshotSummary) GetSnaptimeOk() (*float32, bool)`

GetSnaptimeOk returns a tuple with the Snaptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnaptime

`func (o *SnapshotSummary) SetSnaptime(v float32)`

SetSnaptime sets Snaptime field to given value.

### HasSnaptime

`func (o *SnapshotSummary) HasSnaptime() bool`

HasSnaptime returns a boolean if a field has been set.

### GetVmState

`func (o *SnapshotSummary) GetVmState() bool`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *SnapshotSummary) GetVmStateOk() (*bool, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *SnapshotSummary) SetVmState(v bool)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *SnapshotSummary) HasVmState() bool`

HasVmState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


