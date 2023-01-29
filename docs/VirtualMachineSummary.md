# VirtualMachineSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vmid** | **float32** | The ID of the virtual machine. Unique across cluster | 
**Mem** | Pointer to **float32** | Current memory utilization in bytes | [optional] 
**Maxmem** | Pointer to **float32** | Max memory allocated in bytes | [optional] 
**Maxdisk** | Pointer to **float32** | Max root disk size in bytes | [optional] 
**Name** | Pointer to **string** | The name of the virtual machine | [optional] 
**Status** | Pointer to [**VirtualMachineStatus**](VirtualMachineStatus.md) |  | [optional] 
**Uptime** | Pointer to **float32** | The uptime of the virtual machine in seconds | [optional] 
**Netin** | Pointer to **float32** | Current incoming network traffic in bytes | [optional] 
**Netout** | Pointer to **float32** | Current outgoing network traffic in bytes | [optional] 
**Diskread** | Pointer to **float32** | Current disk read in bytes | [optional] 
**Diskwrite** | Pointer to **float32** | Current disk write in bytes | [optional] 
**Cpu** | Pointer to **float32** | The virtual machines cpu utilization in percent | [optional] 
**Cpus** | Pointer to **float32** | Amount of CPU cores allocated to the virtual machine | [optional] 
**Pid** | Pointer to **float32** | PID of the qemu process on the host node | [optional] 
**Tags** | Pointer to **string** | Tags assigned to the virtual machine. Comma separated list of tags returned as a string. | [optional] 
**Lock** | Pointer to **string** | The current configuration lock | [optional] 
**Qmpstatus** | Pointer to **string** | Qemu QMP agent status | [optional] 

## Methods

### NewVirtualMachineSummary

`func NewVirtualMachineSummary(vmid float32, ) *VirtualMachineSummary`

NewVirtualMachineSummary instantiates a new VirtualMachineSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineSummaryWithDefaults

`func NewVirtualMachineSummaryWithDefaults() *VirtualMachineSummary`

NewVirtualMachineSummaryWithDefaults instantiates a new VirtualMachineSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmid

`func (o *VirtualMachineSummary) GetVmid() float32`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *VirtualMachineSummary) GetVmidOk() (*float32, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *VirtualMachineSummary) SetVmid(v float32)`

SetVmid sets Vmid field to given value.


### GetMem

`func (o *VirtualMachineSummary) GetMem() float32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *VirtualMachineSummary) GetMemOk() (*float32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *VirtualMachineSummary) SetMem(v float32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *VirtualMachineSummary) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetMaxmem

`func (o *VirtualMachineSummary) GetMaxmem() float32`

GetMaxmem returns the Maxmem field if non-nil, zero value otherwise.

### GetMaxmemOk

`func (o *VirtualMachineSummary) GetMaxmemOk() (*float32, bool)`

GetMaxmemOk returns a tuple with the Maxmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxmem

`func (o *VirtualMachineSummary) SetMaxmem(v float32)`

SetMaxmem sets Maxmem field to given value.

### HasMaxmem

`func (o *VirtualMachineSummary) HasMaxmem() bool`

HasMaxmem returns a boolean if a field has been set.

### GetMaxdisk

`func (o *VirtualMachineSummary) GetMaxdisk() float32`

GetMaxdisk returns the Maxdisk field if non-nil, zero value otherwise.

### GetMaxdiskOk

`func (o *VirtualMachineSummary) GetMaxdiskOk() (*float32, bool)`

GetMaxdiskOk returns a tuple with the Maxdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxdisk

`func (o *VirtualMachineSummary) SetMaxdisk(v float32)`

SetMaxdisk sets Maxdisk field to given value.

### HasMaxdisk

`func (o *VirtualMachineSummary) HasMaxdisk() bool`

HasMaxdisk returns a boolean if a field has been set.

### GetName

`func (o *VirtualMachineSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualMachineSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualMachineSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualMachineSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualMachineSummary) GetStatus() VirtualMachineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualMachineSummary) GetStatusOk() (*VirtualMachineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualMachineSummary) SetStatus(v VirtualMachineStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualMachineSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUptime

`func (o *VirtualMachineSummary) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *VirtualMachineSummary) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *VirtualMachineSummary) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *VirtualMachineSummary) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetNetin

`func (o *VirtualMachineSummary) GetNetin() float32`

GetNetin returns the Netin field if non-nil, zero value otherwise.

### GetNetinOk

`func (o *VirtualMachineSummary) GetNetinOk() (*float32, bool)`

GetNetinOk returns a tuple with the Netin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetin

`func (o *VirtualMachineSummary) SetNetin(v float32)`

SetNetin sets Netin field to given value.

### HasNetin

`func (o *VirtualMachineSummary) HasNetin() bool`

HasNetin returns a boolean if a field has been set.

### GetNetout

`func (o *VirtualMachineSummary) GetNetout() float32`

GetNetout returns the Netout field if non-nil, zero value otherwise.

### GetNetoutOk

`func (o *VirtualMachineSummary) GetNetoutOk() (*float32, bool)`

GetNetoutOk returns a tuple with the Netout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetout

`func (o *VirtualMachineSummary) SetNetout(v float32)`

SetNetout sets Netout field to given value.

### HasNetout

`func (o *VirtualMachineSummary) HasNetout() bool`

HasNetout returns a boolean if a field has been set.

### GetDiskread

`func (o *VirtualMachineSummary) GetDiskread() float32`

GetDiskread returns the Diskread field if non-nil, zero value otherwise.

### GetDiskreadOk

`func (o *VirtualMachineSummary) GetDiskreadOk() (*float32, bool)`

GetDiskreadOk returns a tuple with the Diskread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskread

`func (o *VirtualMachineSummary) SetDiskread(v float32)`

SetDiskread sets Diskread field to given value.

### HasDiskread

`func (o *VirtualMachineSummary) HasDiskread() bool`

HasDiskread returns a boolean if a field has been set.

### GetDiskwrite

`func (o *VirtualMachineSummary) GetDiskwrite() float32`

GetDiskwrite returns the Diskwrite field if non-nil, zero value otherwise.

### GetDiskwriteOk

`func (o *VirtualMachineSummary) GetDiskwriteOk() (*float32, bool)`

GetDiskwriteOk returns a tuple with the Diskwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskwrite

`func (o *VirtualMachineSummary) SetDiskwrite(v float32)`

SetDiskwrite sets Diskwrite field to given value.

### HasDiskwrite

`func (o *VirtualMachineSummary) HasDiskwrite() bool`

HasDiskwrite returns a boolean if a field has been set.

### GetCpu

`func (o *VirtualMachineSummary) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *VirtualMachineSummary) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *VirtualMachineSummary) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *VirtualMachineSummary) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCpus

`func (o *VirtualMachineSummary) GetCpus() float32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *VirtualMachineSummary) GetCpusOk() (*float32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *VirtualMachineSummary) SetCpus(v float32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *VirtualMachineSummary) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetPid

`func (o *VirtualMachineSummary) GetPid() float32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *VirtualMachineSummary) GetPidOk() (*float32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *VirtualMachineSummary) SetPid(v float32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *VirtualMachineSummary) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetTags

`func (o *VirtualMachineSummary) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualMachineSummary) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualMachineSummary) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualMachineSummary) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLock

`func (o *VirtualMachineSummary) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *VirtualMachineSummary) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *VirtualMachineSummary) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *VirtualMachineSummary) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetQmpstatus

`func (o *VirtualMachineSummary) GetQmpstatus() string`

GetQmpstatus returns the Qmpstatus field if non-nil, zero value otherwise.

### GetQmpstatusOk

`func (o *VirtualMachineSummary) GetQmpstatusOk() (*string, bool)`

GetQmpstatusOk returns a tuple with the Qmpstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQmpstatus

`func (o *VirtualMachineSummary) SetQmpstatus(v string)`

SetQmpstatus sets Qmpstatus field to given value.

### HasQmpstatus

`func (o *VirtualMachineSummary) HasQmpstatus() bool`

HasQmpstatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


