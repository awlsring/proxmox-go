# ShutdownVirtualMachineRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **float32** |  | [optional] 
**Skiplock** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**ForeStop** | Pointer to **float32** | Ensure the virtual machine stops. | [optional] 
**KeepActive** | Pointer to **float32** | Keep the storage active after Shutdownping it. | [optional] 

## Methods

### NewShutdownVirtualMachineRequestContent

`func NewShutdownVirtualMachineRequestContent() *ShutdownVirtualMachineRequestContent`

NewShutdownVirtualMachineRequestContent instantiates a new ShutdownVirtualMachineRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShutdownVirtualMachineRequestContentWithDefaults

`func NewShutdownVirtualMachineRequestContentWithDefaults() *ShutdownVirtualMachineRequestContent`

NewShutdownVirtualMachineRequestContentWithDefaults instantiates a new ShutdownVirtualMachineRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *ShutdownVirtualMachineRequestContent) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ShutdownVirtualMachineRequestContent) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ShutdownVirtualMachineRequestContent) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ShutdownVirtualMachineRequestContent) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetSkiplock

`func (o *ShutdownVirtualMachineRequestContent) GetSkiplock() float32`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *ShutdownVirtualMachineRequestContent) GetSkiplockOk() (*float32, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *ShutdownVirtualMachineRequestContent) SetSkiplock(v float32)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *ShutdownVirtualMachineRequestContent) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.

### GetForeStop

`func (o *ShutdownVirtualMachineRequestContent) GetForeStop() float32`

GetForeStop returns the ForeStop field if non-nil, zero value otherwise.

### GetForeStopOk

`func (o *ShutdownVirtualMachineRequestContent) GetForeStopOk() (*float32, bool)`

GetForeStopOk returns a tuple with the ForeStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeStop

`func (o *ShutdownVirtualMachineRequestContent) SetForeStop(v float32)`

SetForeStop sets ForeStop field to given value.

### HasForeStop

`func (o *ShutdownVirtualMachineRequestContent) HasForeStop() bool`

HasForeStop returns a boolean if a field has been set.

### GetKeepActive

`func (o *ShutdownVirtualMachineRequestContent) GetKeepActive() float32`

GetKeepActive returns the KeepActive field if non-nil, zero value otherwise.

### GetKeepActiveOk

`func (o *ShutdownVirtualMachineRequestContent) GetKeepActiveOk() (*float32, bool)`

GetKeepActiveOk returns a tuple with the KeepActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepActive

`func (o *ShutdownVirtualMachineRequestContent) SetKeepActive(v float32)`

SetKeepActive sets KeepActive field to given value.

### HasKeepActive

`func (o *ShutdownVirtualMachineRequestContent) HasKeepActive() bool`

HasKeepActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


