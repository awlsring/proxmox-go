# VirtualMachineBalloonSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMem** | Pointer to **float32** |  | [optional] 
**FreeMem** | Pointer to **float32** |  | [optional] 
**MemSwappedOut** | Pointer to **float32** |  | [optional] 
**MemSwappedIn** | Pointer to **float32** |  | [optional] 
**TotalMem** | Pointer to **float32** |  | [optional] 
**Actual** | Pointer to **float32** |  | [optional] 
**LastUpdate** | Pointer to **float32** |  | [optional] 
**MinorPageFaults** | Pointer to **float32** |  | [optional] 
**MajorPageFaults** | Pointer to **float32** |  | [optional] 

## Methods

### NewVirtualMachineBalloonSummary

`func NewVirtualMachineBalloonSummary() *VirtualMachineBalloonSummary`

NewVirtualMachineBalloonSummary instantiates a new VirtualMachineBalloonSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineBalloonSummaryWithDefaults

`func NewVirtualMachineBalloonSummaryWithDefaults() *VirtualMachineBalloonSummary`

NewVirtualMachineBalloonSummaryWithDefaults instantiates a new VirtualMachineBalloonSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMem

`func (o *VirtualMachineBalloonSummary) GetMaxMem() float32`

GetMaxMem returns the MaxMem field if non-nil, zero value otherwise.

### GetMaxMemOk

`func (o *VirtualMachineBalloonSummary) GetMaxMemOk() (*float32, bool)`

GetMaxMemOk returns a tuple with the MaxMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMem

`func (o *VirtualMachineBalloonSummary) SetMaxMem(v float32)`

SetMaxMem sets MaxMem field to given value.

### HasMaxMem

`func (o *VirtualMachineBalloonSummary) HasMaxMem() bool`

HasMaxMem returns a boolean if a field has been set.

### GetFreeMem

`func (o *VirtualMachineBalloonSummary) GetFreeMem() float32`

GetFreeMem returns the FreeMem field if non-nil, zero value otherwise.

### GetFreeMemOk

`func (o *VirtualMachineBalloonSummary) GetFreeMemOk() (*float32, bool)`

GetFreeMemOk returns a tuple with the FreeMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMem

`func (o *VirtualMachineBalloonSummary) SetFreeMem(v float32)`

SetFreeMem sets FreeMem field to given value.

### HasFreeMem

`func (o *VirtualMachineBalloonSummary) HasFreeMem() bool`

HasFreeMem returns a boolean if a field has been set.

### GetMemSwappedOut

`func (o *VirtualMachineBalloonSummary) GetMemSwappedOut() float32`

GetMemSwappedOut returns the MemSwappedOut field if non-nil, zero value otherwise.

### GetMemSwappedOutOk

`func (o *VirtualMachineBalloonSummary) GetMemSwappedOutOk() (*float32, bool)`

GetMemSwappedOutOk returns a tuple with the MemSwappedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemSwappedOut

`func (o *VirtualMachineBalloonSummary) SetMemSwappedOut(v float32)`

SetMemSwappedOut sets MemSwappedOut field to given value.

### HasMemSwappedOut

`func (o *VirtualMachineBalloonSummary) HasMemSwappedOut() bool`

HasMemSwappedOut returns a boolean if a field has been set.

### GetMemSwappedIn

`func (o *VirtualMachineBalloonSummary) GetMemSwappedIn() float32`

GetMemSwappedIn returns the MemSwappedIn field if non-nil, zero value otherwise.

### GetMemSwappedInOk

`func (o *VirtualMachineBalloonSummary) GetMemSwappedInOk() (*float32, bool)`

GetMemSwappedInOk returns a tuple with the MemSwappedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemSwappedIn

`func (o *VirtualMachineBalloonSummary) SetMemSwappedIn(v float32)`

SetMemSwappedIn sets MemSwappedIn field to given value.

### HasMemSwappedIn

`func (o *VirtualMachineBalloonSummary) HasMemSwappedIn() bool`

HasMemSwappedIn returns a boolean if a field has been set.

### GetTotalMem

`func (o *VirtualMachineBalloonSummary) GetTotalMem() float32`

GetTotalMem returns the TotalMem field if non-nil, zero value otherwise.

### GetTotalMemOk

`func (o *VirtualMachineBalloonSummary) GetTotalMemOk() (*float32, bool)`

GetTotalMemOk returns a tuple with the TotalMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMem

`func (o *VirtualMachineBalloonSummary) SetTotalMem(v float32)`

SetTotalMem sets TotalMem field to given value.

### HasTotalMem

`func (o *VirtualMachineBalloonSummary) HasTotalMem() bool`

HasTotalMem returns a boolean if a field has been set.

### GetActual

`func (o *VirtualMachineBalloonSummary) GetActual() float32`

GetActual returns the Actual field if non-nil, zero value otherwise.

### GetActualOk

`func (o *VirtualMachineBalloonSummary) GetActualOk() (*float32, bool)`

GetActualOk returns a tuple with the Actual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActual

`func (o *VirtualMachineBalloonSummary) SetActual(v float32)`

SetActual sets Actual field to given value.

### HasActual

`func (o *VirtualMachineBalloonSummary) HasActual() bool`

HasActual returns a boolean if a field has been set.

### GetLastUpdate

`func (o *VirtualMachineBalloonSummary) GetLastUpdate() float32`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *VirtualMachineBalloonSummary) GetLastUpdateOk() (*float32, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *VirtualMachineBalloonSummary) SetLastUpdate(v float32)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *VirtualMachineBalloonSummary) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetMinorPageFaults

`func (o *VirtualMachineBalloonSummary) GetMinorPageFaults() float32`

GetMinorPageFaults returns the MinorPageFaults field if non-nil, zero value otherwise.

### GetMinorPageFaultsOk

`func (o *VirtualMachineBalloonSummary) GetMinorPageFaultsOk() (*float32, bool)`

GetMinorPageFaultsOk returns a tuple with the MinorPageFaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorPageFaults

`func (o *VirtualMachineBalloonSummary) SetMinorPageFaults(v float32)`

SetMinorPageFaults sets MinorPageFaults field to given value.

### HasMinorPageFaults

`func (o *VirtualMachineBalloonSummary) HasMinorPageFaults() bool`

HasMinorPageFaults returns a boolean if a field has been set.

### GetMajorPageFaults

`func (o *VirtualMachineBalloonSummary) GetMajorPageFaults() float32`

GetMajorPageFaults returns the MajorPageFaults field if non-nil, zero value otherwise.

### GetMajorPageFaultsOk

`func (o *VirtualMachineBalloonSummary) GetMajorPageFaultsOk() (*float32, bool)`

GetMajorPageFaultsOk returns a tuple with the MajorPageFaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorPageFaults

`func (o *VirtualMachineBalloonSummary) SetMajorPageFaults(v float32)`

SetMajorPageFaults sets MajorPageFaults field to given value.

### HasMajorPageFaults

`func (o *VirtualMachineBalloonSummary) HasMajorPageFaults() bool`

HasMajorPageFaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


