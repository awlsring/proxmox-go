# VirtualMachineNicBlockStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlushTotalTimeNs** | Pointer to **float32** |  | [optional] 
**RdBytes** | Pointer to **float32** |  | [optional] 
**TimedStats** | Pointer to **[]string** |  | [optional] 
**WrHighestOffset** | Pointer to **float32** |  | [optional] 
**RdTotalTimeNs** | Pointer to **float32** |  | [optional] 
**FlushOperations** | Pointer to **float32** |  | [optional] 
**WrOperations** | Pointer to **float32** |  | [optional] 
**IdleTimeNs** | Pointer to **float32** |  | [optional] 
**WrMerged** | Pointer to **float32** |  | [optional] 
**InvalidRdOperations** | Pointer to **float32** |  | [optional] 
**FailedFlushOperations** | Pointer to **float32** |  | [optional] 
**UnmapBytes** | Pointer to **float32** |  | [optional] 
**FailedRdOperations** | Pointer to **float32** |  | [optional] 
**WrBytes** | Pointer to **float32** |  | [optional] 
**InvalidFlushOperations** | Pointer to **float32** |  | [optional] 
**UnmapOperations** | Pointer to **float32** |  | [optional] 
**WrTotalTimeNs** | Pointer to **float32** |  | [optional] 
**FailedWrOperations** | Pointer to **float32** |  | [optional] 
**UnmapMerged** | Pointer to **float32** |  | [optional] 
**InvalidWrOperations** | Pointer to **float32** |  | [optional] 
**RdOperations** | Pointer to **float32** |  | [optional] 
**UnmapTotalTimeNs** | Pointer to **float32** |  | [optional] 
**InvalidUnmapOperations** | Pointer to **float32** |  | [optional] 
**AccountFailed** | Pointer to **bool** |  | [optional] 
**AccountInvalid** | Pointer to **bool** |  | [optional] 
**RdMerged** | Pointer to **float32** |  | [optional] 
**FailedUnmapOperations** | Pointer to **float32** |  | [optional] 

## Methods

### NewVirtualMachineNicBlockStatistics

`func NewVirtualMachineNicBlockStatistics() *VirtualMachineNicBlockStatistics`

NewVirtualMachineNicBlockStatistics instantiates a new VirtualMachineNicBlockStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineNicBlockStatisticsWithDefaults

`func NewVirtualMachineNicBlockStatisticsWithDefaults() *VirtualMachineNicBlockStatistics`

NewVirtualMachineNicBlockStatisticsWithDefaults instantiates a new VirtualMachineNicBlockStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlushTotalTimeNs

`func (o *VirtualMachineNicBlockStatistics) GetFlushTotalTimeNs() float32`

GetFlushTotalTimeNs returns the FlushTotalTimeNs field if non-nil, zero value otherwise.

### GetFlushTotalTimeNsOk

`func (o *VirtualMachineNicBlockStatistics) GetFlushTotalTimeNsOk() (*float32, bool)`

GetFlushTotalTimeNsOk returns a tuple with the FlushTotalTimeNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushTotalTimeNs

`func (o *VirtualMachineNicBlockStatistics) SetFlushTotalTimeNs(v float32)`

SetFlushTotalTimeNs sets FlushTotalTimeNs field to given value.

### HasFlushTotalTimeNs

`func (o *VirtualMachineNicBlockStatistics) HasFlushTotalTimeNs() bool`

HasFlushTotalTimeNs returns a boolean if a field has been set.

### GetRdBytes

`func (o *VirtualMachineNicBlockStatistics) GetRdBytes() float32`

GetRdBytes returns the RdBytes field if non-nil, zero value otherwise.

### GetRdBytesOk

`func (o *VirtualMachineNicBlockStatistics) GetRdBytesOk() (*float32, bool)`

GetRdBytesOk returns a tuple with the RdBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdBytes

`func (o *VirtualMachineNicBlockStatistics) SetRdBytes(v float32)`

SetRdBytes sets RdBytes field to given value.

### HasRdBytes

`func (o *VirtualMachineNicBlockStatistics) HasRdBytes() bool`

HasRdBytes returns a boolean if a field has been set.

### GetTimedStats

`func (o *VirtualMachineNicBlockStatistics) GetTimedStats() []string`

GetTimedStats returns the TimedStats field if non-nil, zero value otherwise.

### GetTimedStatsOk

`func (o *VirtualMachineNicBlockStatistics) GetTimedStatsOk() (*[]string, bool)`

GetTimedStatsOk returns a tuple with the TimedStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedStats

`func (o *VirtualMachineNicBlockStatistics) SetTimedStats(v []string)`

SetTimedStats sets TimedStats field to given value.

### HasTimedStats

`func (o *VirtualMachineNicBlockStatistics) HasTimedStats() bool`

HasTimedStats returns a boolean if a field has been set.

### GetWrHighestOffset

`func (o *VirtualMachineNicBlockStatistics) GetWrHighestOffset() float32`

GetWrHighestOffset returns the WrHighestOffset field if non-nil, zero value otherwise.

### GetWrHighestOffsetOk

`func (o *VirtualMachineNicBlockStatistics) GetWrHighestOffsetOk() (*float32, bool)`

GetWrHighestOffsetOk returns a tuple with the WrHighestOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrHighestOffset

`func (o *VirtualMachineNicBlockStatistics) SetWrHighestOffset(v float32)`

SetWrHighestOffset sets WrHighestOffset field to given value.

### HasWrHighestOffset

`func (o *VirtualMachineNicBlockStatistics) HasWrHighestOffset() bool`

HasWrHighestOffset returns a boolean if a field has been set.

### GetRdTotalTimeNs

`func (o *VirtualMachineNicBlockStatistics) GetRdTotalTimeNs() float32`

GetRdTotalTimeNs returns the RdTotalTimeNs field if non-nil, zero value otherwise.

### GetRdTotalTimeNsOk

`func (o *VirtualMachineNicBlockStatistics) GetRdTotalTimeNsOk() (*float32, bool)`

GetRdTotalTimeNsOk returns a tuple with the RdTotalTimeNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdTotalTimeNs

`func (o *VirtualMachineNicBlockStatistics) SetRdTotalTimeNs(v float32)`

SetRdTotalTimeNs sets RdTotalTimeNs field to given value.

### HasRdTotalTimeNs

`func (o *VirtualMachineNicBlockStatistics) HasRdTotalTimeNs() bool`

HasRdTotalTimeNs returns a boolean if a field has been set.

### GetFlushOperations

`func (o *VirtualMachineNicBlockStatistics) GetFlushOperations() float32`

GetFlushOperations returns the FlushOperations field if non-nil, zero value otherwise.

### GetFlushOperationsOk

`func (o *VirtualMachineNicBlockStatistics) GetFlushOperationsOk() (*float32, bool)`

GetFlushOperationsOk returns a tuple with the FlushOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushOperations

`func (o *VirtualMachineNicBlockStatistics) SetFlushOperations(v float32)`

SetFlushOperations sets FlushOperations field to given value.

### HasFlushOperations

`func (o *VirtualMachineNicBlockStatistics) HasFlushOperations() bool`

HasFlushOperations returns a boolean if a field has been set.

### GetWrOperations

`func (o *VirtualMachineNicBlockStatistics) GetWrOperations() float32`

GetWrOperations returns the WrOperations field if non-nil, zero value otherwise.

### GetWrOperationsOk

`func (o *VirtualMachineNicBlockStatistics) GetWrOperationsOk() (*float32, bool)`

GetWrOperationsOk returns a tuple with the WrOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrOperations

`func (o *VirtualMachineNicBlockStatistics) SetWrOperations(v float32)`

SetWrOperations sets WrOperations field to given value.

### HasWrOperations

`func (o *VirtualMachineNicBlockStatistics) HasWrOperations() bool`

HasWrOperations returns a boolean if a field has been set.

### GetIdleTimeNs

`func (o *VirtualMachineNicBlockStatistics) GetIdleTimeNs() float32`

GetIdleTimeNs returns the IdleTimeNs field if non-nil, zero value otherwise.

### GetIdleTimeNsOk

`func (o *VirtualMachineNicBlockStatistics) GetIdleTimeNsOk() (*float32, bool)`

GetIdleTimeNsOk returns a tuple with the IdleTimeNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeNs

`func (o *VirtualMachineNicBlockStatistics) SetIdleTimeNs(v float32)`

SetIdleTimeNs sets IdleTimeNs field to given value.

### HasIdleTimeNs

`func (o *VirtualMachineNicBlockStatistics) HasIdleTimeNs() bool`

HasIdleTimeNs returns a boolean if a field has been set.

### GetWrMerged

`func (o *VirtualMachineNicBlockStatistics) GetWrMerged() float32`

GetWrMerged returns the WrMerged field if non-nil, zero value otherwise.

### GetWrMergedOk

`func (o *VirtualMachineNicBlockStatistics) GetWrMergedOk() (*float32, bool)`

GetWrMergedOk returns a tuple with the WrMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrMerged

`func (o *VirtualMachineNicBlockStatistics) SetWrMerged(v float32)`

SetWrMerged sets WrMerged field to given value.

### HasWrMerged

`func (o *VirtualMachineNicBlockStatistics) HasWrMerged() bool`

HasWrMerged returns a boolean if a field has been set.

### GetInvalidRdOperations

`func (o *VirtualMachineNicBlockStatistics) GetInvalidRdOperations() float32`

GetInvalidRdOperations returns the InvalidRdOperations field if non-nil, zero value otherwise.

### GetInvalidRdOperationsOk

`func (o *VirtualMachineNicBlockStatistics) GetInvalidRdOperationsOk() (*float32, bool)`

GetInvalidRdOperationsOk returns a tuple with the InvalidRdOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidRdOperations

`func (o *VirtualMachineNicBlockStatistics) SetInvalidRdOperations(v float32)`

SetInvalidRdOperations sets InvalidRdOperations field to given value.

### HasInvalidRdOperations

`func (o *VirtualMachineNicBlockStatistics) HasInvalidRdOperations() bool`

HasInvalidRdOperations returns a boolean if a field has been set.

### GetFailedFlushOperations

`func (o *VirtualMachineNicBlockStatistics) GetFailedFlushOperations() float32`

GetFailedFlushOperations returns the FailedFlushOperations field if non-nil, zero value otherwise.

### GetFailedFlushOperationsOk

`func (o *VirtualMachineNicBlockStatistics) GetFailedFlushOperationsOk() (*float32, bool)`

GetFailedFlushOperationsOk returns a tuple with the FailedFlushOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedFlushOperations

`func (o *VirtualMachineNicBlockStatistics) SetFailedFlushOperations(v float32)`

SetFailedFlushOperations sets FailedFlushOperations field to given value.

### HasFailedFlushOperations

`func (o *VirtualMachineNicBlockStatistics) HasFailedFlushOperations() bool`

HasFailedFlushOperations returns a boolean if a field has been set.

### GetUnmapBytes

`func (o *VirtualMachineNicBlockStatistics) GetUnmapBytes() float32`

GetUnmapBytes returns the UnmapBytes field if non-nil, zero value otherwise.

### GetUnmapBytesOk

`func (o *VirtualMachineNicBlockStatistics) GetUnmapBytesOk() (*float32, bool)`

GetUnmapBytesOk returns a tuple with the UnmapBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmapBytes

`func (o *VirtualMachineNicBlockStatistics) SetUnmapBytes(v float32)`

SetUnmapBytes sets UnmapBytes field to given value.

### HasUnmapBytes

`func (o *VirtualMachineNicBlockStatistics) HasUnmapBytes() bool`

HasUnmapBytes returns a boolean if a field has been set.

### GetFailedRdOperations

`func (o *VirtualMachineNicBlockStatistics) GetFailedRdOperations() float32`

GetFailedRdOperations returns the FailedRdOperations field if non-nil, zero value otherwise.

### GetFailedRdOperationsOk

`func (o *VirtualMachineNicBlockStatistics) GetFailedRdOperationsOk() (*float32, bool)`

GetFailedRdOperationsOk returns a tuple with the FailedRdOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRdOperations

`func (o *VirtualMachineNicBlockStatistics) SetFailedRdOperations(v float32)`

SetFailedRdOperations sets FailedRdOperations field to given value.

### HasFailedRdOperations

`func (o *VirtualMachineNicBlockStatistics) HasFailedRdOperations() bool`

HasFailedRdOperations returns a boolean if a field has been set.

### GetWrBytes

`func (o *VirtualMachineNicBlockStatistics) GetWrBytes() float32`

GetWrBytes returns the WrBytes field if non-nil, zero value otherwise.

### GetWrBytesOk

`func (o *VirtualMachineNicBlockStatistics) GetWrBytesOk() (*float32, bool)`

GetWrBytesOk returns a tuple with the WrBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrBytes

`func (o *VirtualMachineNicBlockStatistics) SetWrBytes(v float32)`

SetWrBytes sets WrBytes field to given value.

### HasWrBytes

`func (o *VirtualMachineNicBlockStatistics) HasWrBytes() bool`

HasWrBytes returns a boolean if a field has been set.

### GetInvalidFlushOperations

`func (o *VirtualMachineNicBlockStatistics) GetInvalidFlushOperations() float32`

GetInvalidFlushOperations returns the InvalidFlushOperations field if non-nil, zero value otherwise.

### GetInvalidFlushOperationsOk

`func (o *VirtualMachineNicBlockStatistics) GetInvalidFlushOperationsOk() (*float32, bool)`

GetInvalidFlushOperationsOk returns a tuple with the InvalidFlushOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidFlushOperations

`func (o *VirtualMachineNicBlockStatistics) SetInvalidFlushOperations(v float32)`

SetInvalidFlushOperations sets InvalidFlushOperations field to given value.

### HasInvalidFlushOperations

`func (o *VirtualMachineNicBlockStatistics) HasInvalidFlushOperations() bool`

HasInvalidFlushOperations returns a boolean if a field has been set.

### GetUnmapOperations

`func (o *VirtualMachineNicBlockStatistics) GetUnmapOperations() float32`

GetUnmapOperations returns the UnmapOperations field if non-nil, zero value otherwise.

### GetUnmapOperationsOk

`func (o *VirtualMachineNicBlockStatistics) GetUnmapOperationsOk() (*float32, bool)`

GetUnmapOperationsOk returns a tuple with the UnmapOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmapOperations

`func (o *VirtualMachineNicBlockStatistics) SetUnmapOperations(v float32)`

SetUnmapOperations sets UnmapOperations field to given value.

### HasUnmapOperations

`func (o *VirtualMachineNicBlockStatistics) HasUnmapOperations() bool`

HasUnmapOperations returns a boolean if a field has been set.

### GetWrTotalTimeNs

`func (o *VirtualMachineNicBlockStatistics) GetWrTotalTimeNs() float32`

GetWrTotalTimeNs returns the WrTotalTimeNs field if non-nil, zero value otherwise.

### GetWrTotalTimeNsOk

`func (o *VirtualMachineNicBlockStatistics) GetWrTotalTimeNsOk() (*float32, bool)`

GetWrTotalTimeNsOk returns a tuple with the WrTotalTimeNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrTotalTimeNs

`func (o *VirtualMachineNicBlockStatistics) SetWrTotalTimeNs(v float32)`

SetWrTotalTimeNs sets WrTotalTimeNs field to given value.

### HasWrTotalTimeNs

`func (o *VirtualMachineNicBlockStatistics) HasWrTotalTimeNs() bool`

HasWrTotalTimeNs returns a boolean if a field has been set.

### GetFailedWrOperations

`func (o *VirtualMachineNicBlockStatistics) GetFailedWrOperations() float32`

GetFailedWrOperations returns the FailedWrOperations field if non-nil, zero value otherwise.

### GetFailedWrOperationsOk

`func (o *VirtualMachineNicBlockStatistics) GetFailedWrOperationsOk() (*float32, bool)`

GetFailedWrOperationsOk returns a tuple with the FailedWrOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedWrOperations

`func (o *VirtualMachineNicBlockStatistics) SetFailedWrOperations(v float32)`

SetFailedWrOperations sets FailedWrOperations field to given value.

### HasFailedWrOperations

`func (o *VirtualMachineNicBlockStatistics) HasFailedWrOperations() bool`

HasFailedWrOperations returns a boolean if a field has been set.

### GetUnmapMerged

`func (o *VirtualMachineNicBlockStatistics) GetUnmapMerged() float32`

GetUnmapMerged returns the UnmapMerged field if non-nil, zero value otherwise.

### GetUnmapMergedOk

`func (o *VirtualMachineNicBlockStatistics) GetUnmapMergedOk() (*float32, bool)`

GetUnmapMergedOk returns a tuple with the UnmapMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmapMerged

`func (o *VirtualMachineNicBlockStatistics) SetUnmapMerged(v float32)`

SetUnmapMerged sets UnmapMerged field to given value.

### HasUnmapMerged

`func (o *VirtualMachineNicBlockStatistics) HasUnmapMerged() bool`

HasUnmapMerged returns a boolean if a field has been set.

### GetInvalidWrOperations

`func (o *VirtualMachineNicBlockStatistics) GetInvalidWrOperations() float32`

GetInvalidWrOperations returns the InvalidWrOperations field if non-nil, zero value otherwise.

### GetInvalidWrOperationsOk

`func (o *VirtualMachineNicBlockStatistics) GetInvalidWrOperationsOk() (*float32, bool)`

GetInvalidWrOperationsOk returns a tuple with the InvalidWrOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidWrOperations

`func (o *VirtualMachineNicBlockStatistics) SetInvalidWrOperations(v float32)`

SetInvalidWrOperations sets InvalidWrOperations field to given value.

### HasInvalidWrOperations

`func (o *VirtualMachineNicBlockStatistics) HasInvalidWrOperations() bool`

HasInvalidWrOperations returns a boolean if a field has been set.

### GetRdOperations

`func (o *VirtualMachineNicBlockStatistics) GetRdOperations() float32`

GetRdOperations returns the RdOperations field if non-nil, zero value otherwise.

### GetRdOperationsOk

`func (o *VirtualMachineNicBlockStatistics) GetRdOperationsOk() (*float32, bool)`

GetRdOperationsOk returns a tuple with the RdOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdOperations

`func (o *VirtualMachineNicBlockStatistics) SetRdOperations(v float32)`

SetRdOperations sets RdOperations field to given value.

### HasRdOperations

`func (o *VirtualMachineNicBlockStatistics) HasRdOperations() bool`

HasRdOperations returns a boolean if a field has been set.

### GetUnmapTotalTimeNs

`func (o *VirtualMachineNicBlockStatistics) GetUnmapTotalTimeNs() float32`

GetUnmapTotalTimeNs returns the UnmapTotalTimeNs field if non-nil, zero value otherwise.

### GetUnmapTotalTimeNsOk

`func (o *VirtualMachineNicBlockStatistics) GetUnmapTotalTimeNsOk() (*float32, bool)`

GetUnmapTotalTimeNsOk returns a tuple with the UnmapTotalTimeNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmapTotalTimeNs

`func (o *VirtualMachineNicBlockStatistics) SetUnmapTotalTimeNs(v float32)`

SetUnmapTotalTimeNs sets UnmapTotalTimeNs field to given value.

### HasUnmapTotalTimeNs

`func (o *VirtualMachineNicBlockStatistics) HasUnmapTotalTimeNs() bool`

HasUnmapTotalTimeNs returns a boolean if a field has been set.

### GetInvalidUnmapOperations

`func (o *VirtualMachineNicBlockStatistics) GetInvalidUnmapOperations() float32`

GetInvalidUnmapOperations returns the InvalidUnmapOperations field if non-nil, zero value otherwise.

### GetInvalidUnmapOperationsOk

`func (o *VirtualMachineNicBlockStatistics) GetInvalidUnmapOperationsOk() (*float32, bool)`

GetInvalidUnmapOperationsOk returns a tuple with the InvalidUnmapOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidUnmapOperations

`func (o *VirtualMachineNicBlockStatistics) SetInvalidUnmapOperations(v float32)`

SetInvalidUnmapOperations sets InvalidUnmapOperations field to given value.

### HasInvalidUnmapOperations

`func (o *VirtualMachineNicBlockStatistics) HasInvalidUnmapOperations() bool`

HasInvalidUnmapOperations returns a boolean if a field has been set.

### GetAccountFailed

`func (o *VirtualMachineNicBlockStatistics) GetAccountFailed() bool`

GetAccountFailed returns the AccountFailed field if non-nil, zero value otherwise.

### GetAccountFailedOk

`func (o *VirtualMachineNicBlockStatistics) GetAccountFailedOk() (*bool, bool)`

GetAccountFailedOk returns a tuple with the AccountFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFailed

`func (o *VirtualMachineNicBlockStatistics) SetAccountFailed(v bool)`

SetAccountFailed sets AccountFailed field to given value.

### HasAccountFailed

`func (o *VirtualMachineNicBlockStatistics) HasAccountFailed() bool`

HasAccountFailed returns a boolean if a field has been set.

### GetAccountInvalid

`func (o *VirtualMachineNicBlockStatistics) GetAccountInvalid() bool`

GetAccountInvalid returns the AccountInvalid field if non-nil, zero value otherwise.

### GetAccountInvalidOk

`func (o *VirtualMachineNicBlockStatistics) GetAccountInvalidOk() (*bool, bool)`

GetAccountInvalidOk returns a tuple with the AccountInvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInvalid

`func (o *VirtualMachineNicBlockStatistics) SetAccountInvalid(v bool)`

SetAccountInvalid sets AccountInvalid field to given value.

### HasAccountInvalid

`func (o *VirtualMachineNicBlockStatistics) HasAccountInvalid() bool`

HasAccountInvalid returns a boolean if a field has been set.

### GetRdMerged

`func (o *VirtualMachineNicBlockStatistics) GetRdMerged() float32`

GetRdMerged returns the RdMerged field if non-nil, zero value otherwise.

### GetRdMergedOk

`func (o *VirtualMachineNicBlockStatistics) GetRdMergedOk() (*float32, bool)`

GetRdMergedOk returns a tuple with the RdMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdMerged

`func (o *VirtualMachineNicBlockStatistics) SetRdMerged(v float32)`

SetRdMerged sets RdMerged field to given value.

### HasRdMerged

`func (o *VirtualMachineNicBlockStatistics) HasRdMerged() bool`

HasRdMerged returns a boolean if a field has been set.

### GetFailedUnmapOperations

`func (o *VirtualMachineNicBlockStatistics) GetFailedUnmapOperations() float32`

GetFailedUnmapOperations returns the FailedUnmapOperations field if non-nil, zero value otherwise.

### GetFailedUnmapOperationsOk

`func (o *VirtualMachineNicBlockStatistics) GetFailedUnmapOperationsOk() (*float32, bool)`

GetFailedUnmapOperationsOk returns a tuple with the FailedUnmapOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedUnmapOperations

`func (o *VirtualMachineNicBlockStatistics) SetFailedUnmapOperations(v float32)`

SetFailedUnmapOperations sets FailedUnmapOperations field to given value.

### HasFailedUnmapOperations

`func (o *VirtualMachineNicBlockStatistics) HasFailedUnmapOperations() bool`

HasFailedUnmapOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


