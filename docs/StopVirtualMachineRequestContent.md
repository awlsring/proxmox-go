# StopVirtualMachineRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **float32** |  | [optional] 
**Skiplock** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Migratedfrom** | Pointer to **string** | Specify migrated from node. | [optional] 
**KeepActive** | Pointer to **float32** | Keep the storage active after stopping it. | [optional] 

## Methods

### NewStopVirtualMachineRequestContent

`func NewStopVirtualMachineRequestContent() *StopVirtualMachineRequestContent`

NewStopVirtualMachineRequestContent instantiates a new StopVirtualMachineRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopVirtualMachineRequestContentWithDefaults

`func NewStopVirtualMachineRequestContentWithDefaults() *StopVirtualMachineRequestContent`

NewStopVirtualMachineRequestContentWithDefaults instantiates a new StopVirtualMachineRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *StopVirtualMachineRequestContent) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *StopVirtualMachineRequestContent) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *StopVirtualMachineRequestContent) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *StopVirtualMachineRequestContent) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetSkiplock

`func (o *StopVirtualMachineRequestContent) GetSkiplock() float32`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *StopVirtualMachineRequestContent) GetSkiplockOk() (*float32, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *StopVirtualMachineRequestContent) SetSkiplock(v float32)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *StopVirtualMachineRequestContent) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.

### GetMigratedfrom

`func (o *StopVirtualMachineRequestContent) GetMigratedfrom() string`

GetMigratedfrom returns the Migratedfrom field if non-nil, zero value otherwise.

### GetMigratedfromOk

`func (o *StopVirtualMachineRequestContent) GetMigratedfromOk() (*string, bool)`

GetMigratedfromOk returns a tuple with the Migratedfrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedfrom

`func (o *StopVirtualMachineRequestContent) SetMigratedfrom(v string)`

SetMigratedfrom sets Migratedfrom field to given value.

### HasMigratedfrom

`func (o *StopVirtualMachineRequestContent) HasMigratedfrom() bool`

HasMigratedfrom returns a boolean if a field has been set.

### GetKeepActive

`func (o *StopVirtualMachineRequestContent) GetKeepActive() float32`

GetKeepActive returns the KeepActive field if non-nil, zero value otherwise.

### GetKeepActiveOk

`func (o *StopVirtualMachineRequestContent) GetKeepActiveOk() (*float32, bool)`

GetKeepActiveOk returns a tuple with the KeepActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepActive

`func (o *StopVirtualMachineRequestContent) SetKeepActive(v float32)`

SetKeepActive sets KeepActive field to given value.

### HasKeepActive

`func (o *StopVirtualMachineRequestContent) HasKeepActive() bool`

HasKeepActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


