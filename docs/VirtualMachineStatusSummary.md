# VirtualMachineStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**VirtualMachineStatus**](VirtualMachineStatus.md) |  | 
**Ha** | [**VirtualMachineHighAvailabilityStatus**](VirtualMachineHighAvailabilityStatus.md) |  | 
**Vmid** | **float32** | The id of the virtual machine as an integer | 
**Agent** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Cpus** | Pointer to **float32** |  | [optional] 
**Lock** | Pointer to **string** |  | [optional] 
**Maxdisk** | Pointer to **float32** |  | [optional] 
**Maxmem** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Pid** | Pointer to **float32** |  | [optional] 
**Qmpstatus** | Pointer to **string** |  | [optional] 
**RunningMachine** | Pointer to **string** |  | [optional] 
**RunningQemu** | Pointer to **string** |  | [optional] 
**Spice** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**Uptime** | Pointer to **float32** |  | [optional] 
**Cpu** | Pointer to **float32** |  | [optional] 
**Mem** | Pointer to **float32** |  | [optional] 
**Balloon** | Pointer to **float32** |  | [optional] 
**Disk** | Pointer to **float32** |  | [optional] 
**Netin** | Pointer to **float32** |  | [optional] 
**Netout** | Pointer to **float32** |  | [optional] 
**Diskwrite** | Pointer to **float32** |  | [optional] 
**Diskread** | Pointer to **float32** |  | [optional] 
**Freemem** | Pointer to **float32** |  | [optional] 
**Nics** | Pointer to [**map[string]VirtualMachineNicStatus**](VirtualMachineNicStatus.md) |  | [optional] 
**Ballooninfo** | Pointer to [**VirtualMachineBalloonSummary**](VirtualMachineBalloonSummary.md) |  | [optional] 
**Blockstat** | Pointer to [**map[string]VirtualMachineNicBlockStatistics**](VirtualMachineNicBlockStatistics.md) |  | [optional] 
**ProxmoxSupport** | Pointer to [**ProxmoxSupportSummary**](ProxmoxSupportSummary.md) |  | [optional] 

## Methods

### NewVirtualMachineStatusSummary

`func NewVirtualMachineStatusSummary(status VirtualMachineStatus, ha VirtualMachineHighAvailabilityStatus, vmid float32, ) *VirtualMachineStatusSummary`

NewVirtualMachineStatusSummary instantiates a new VirtualMachineStatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineStatusSummaryWithDefaults

`func NewVirtualMachineStatusSummaryWithDefaults() *VirtualMachineStatusSummary`

NewVirtualMachineStatusSummaryWithDefaults instantiates a new VirtualMachineStatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VirtualMachineStatusSummary) GetStatus() VirtualMachineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualMachineStatusSummary) GetStatusOk() (*VirtualMachineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualMachineStatusSummary) SetStatus(v VirtualMachineStatus)`

SetStatus sets Status field to given value.


### GetHa

`func (o *VirtualMachineStatusSummary) GetHa() VirtualMachineHighAvailabilityStatus`

GetHa returns the Ha field if non-nil, zero value otherwise.

### GetHaOk

`func (o *VirtualMachineStatusSummary) GetHaOk() (*VirtualMachineHighAvailabilityStatus, bool)`

GetHaOk returns a tuple with the Ha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHa

`func (o *VirtualMachineStatusSummary) SetHa(v VirtualMachineHighAvailabilityStatus)`

SetHa sets Ha field to given value.


### GetVmid

`func (o *VirtualMachineStatusSummary) GetVmid() float32`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *VirtualMachineStatusSummary) GetVmidOk() (*float32, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *VirtualMachineStatusSummary) SetVmid(v float32)`

SetVmid sets Vmid field to given value.


### GetAgent

`func (o *VirtualMachineStatusSummary) GetAgent() float32`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *VirtualMachineStatusSummary) GetAgentOk() (*float32, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *VirtualMachineStatusSummary) SetAgent(v float32)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *VirtualMachineStatusSummary) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetCpus

`func (o *VirtualMachineStatusSummary) GetCpus() float32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *VirtualMachineStatusSummary) GetCpusOk() (*float32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *VirtualMachineStatusSummary) SetCpus(v float32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *VirtualMachineStatusSummary) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetLock

`func (o *VirtualMachineStatusSummary) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *VirtualMachineStatusSummary) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *VirtualMachineStatusSummary) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *VirtualMachineStatusSummary) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMaxdisk

`func (o *VirtualMachineStatusSummary) GetMaxdisk() float32`

GetMaxdisk returns the Maxdisk field if non-nil, zero value otherwise.

### GetMaxdiskOk

`func (o *VirtualMachineStatusSummary) GetMaxdiskOk() (*float32, bool)`

GetMaxdiskOk returns a tuple with the Maxdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxdisk

`func (o *VirtualMachineStatusSummary) SetMaxdisk(v float32)`

SetMaxdisk sets Maxdisk field to given value.

### HasMaxdisk

`func (o *VirtualMachineStatusSummary) HasMaxdisk() bool`

HasMaxdisk returns a boolean if a field has been set.

### GetMaxmem

`func (o *VirtualMachineStatusSummary) GetMaxmem() float32`

GetMaxmem returns the Maxmem field if non-nil, zero value otherwise.

### GetMaxmemOk

`func (o *VirtualMachineStatusSummary) GetMaxmemOk() (*float32, bool)`

GetMaxmemOk returns a tuple with the Maxmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxmem

`func (o *VirtualMachineStatusSummary) SetMaxmem(v float32)`

SetMaxmem sets Maxmem field to given value.

### HasMaxmem

`func (o *VirtualMachineStatusSummary) HasMaxmem() bool`

HasMaxmem returns a boolean if a field has been set.

### GetName

`func (o *VirtualMachineStatusSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualMachineStatusSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualMachineStatusSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualMachineStatusSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPid

`func (o *VirtualMachineStatusSummary) GetPid() float32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *VirtualMachineStatusSummary) GetPidOk() (*float32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *VirtualMachineStatusSummary) SetPid(v float32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *VirtualMachineStatusSummary) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetQmpstatus

`func (o *VirtualMachineStatusSummary) GetQmpstatus() string`

GetQmpstatus returns the Qmpstatus field if non-nil, zero value otherwise.

### GetQmpstatusOk

`func (o *VirtualMachineStatusSummary) GetQmpstatusOk() (*string, bool)`

GetQmpstatusOk returns a tuple with the Qmpstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQmpstatus

`func (o *VirtualMachineStatusSummary) SetQmpstatus(v string)`

SetQmpstatus sets Qmpstatus field to given value.

### HasQmpstatus

`func (o *VirtualMachineStatusSummary) HasQmpstatus() bool`

HasQmpstatus returns a boolean if a field has been set.

### GetRunningMachine

`func (o *VirtualMachineStatusSummary) GetRunningMachine() string`

GetRunningMachine returns the RunningMachine field if non-nil, zero value otherwise.

### GetRunningMachineOk

`func (o *VirtualMachineStatusSummary) GetRunningMachineOk() (*string, bool)`

GetRunningMachineOk returns a tuple with the RunningMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningMachine

`func (o *VirtualMachineStatusSummary) SetRunningMachine(v string)`

SetRunningMachine sets RunningMachine field to given value.

### HasRunningMachine

`func (o *VirtualMachineStatusSummary) HasRunningMachine() bool`

HasRunningMachine returns a boolean if a field has been set.

### GetRunningQemu

`func (o *VirtualMachineStatusSummary) GetRunningQemu() string`

GetRunningQemu returns the RunningQemu field if non-nil, zero value otherwise.

### GetRunningQemuOk

`func (o *VirtualMachineStatusSummary) GetRunningQemuOk() (*string, bool)`

GetRunningQemuOk returns a tuple with the RunningQemu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningQemu

`func (o *VirtualMachineStatusSummary) SetRunningQemu(v string)`

SetRunningQemu sets RunningQemu field to given value.

### HasRunningQemu

`func (o *VirtualMachineStatusSummary) HasRunningQemu() bool`

HasRunningQemu returns a boolean if a field has been set.

### GetSpice

`func (o *VirtualMachineStatusSummary) GetSpice() float32`

GetSpice returns the Spice field if non-nil, zero value otherwise.

### GetSpiceOk

`func (o *VirtualMachineStatusSummary) GetSpiceOk() (*float32, bool)`

GetSpiceOk returns a tuple with the Spice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpice

`func (o *VirtualMachineStatusSummary) SetSpice(v float32)`

SetSpice sets Spice field to given value.

### HasSpice

`func (o *VirtualMachineStatusSummary) HasSpice() bool`

HasSpice returns a boolean if a field has been set.

### GetTags

`func (o *VirtualMachineStatusSummary) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualMachineStatusSummary) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualMachineStatusSummary) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualMachineStatusSummary) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUptime

`func (o *VirtualMachineStatusSummary) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *VirtualMachineStatusSummary) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *VirtualMachineStatusSummary) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *VirtualMachineStatusSummary) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetCpu

`func (o *VirtualMachineStatusSummary) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *VirtualMachineStatusSummary) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *VirtualMachineStatusSummary) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *VirtualMachineStatusSummary) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMem

`func (o *VirtualMachineStatusSummary) GetMem() float32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *VirtualMachineStatusSummary) GetMemOk() (*float32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *VirtualMachineStatusSummary) SetMem(v float32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *VirtualMachineStatusSummary) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetBalloon

`func (o *VirtualMachineStatusSummary) GetBalloon() float32`

GetBalloon returns the Balloon field if non-nil, zero value otherwise.

### GetBalloonOk

`func (o *VirtualMachineStatusSummary) GetBalloonOk() (*float32, bool)`

GetBalloonOk returns a tuple with the Balloon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalloon

`func (o *VirtualMachineStatusSummary) SetBalloon(v float32)`

SetBalloon sets Balloon field to given value.

### HasBalloon

`func (o *VirtualMachineStatusSummary) HasBalloon() bool`

HasBalloon returns a boolean if a field has been set.

### GetDisk

`func (o *VirtualMachineStatusSummary) GetDisk() float32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *VirtualMachineStatusSummary) GetDiskOk() (*float32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *VirtualMachineStatusSummary) SetDisk(v float32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *VirtualMachineStatusSummary) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetNetin

`func (o *VirtualMachineStatusSummary) GetNetin() float32`

GetNetin returns the Netin field if non-nil, zero value otherwise.

### GetNetinOk

`func (o *VirtualMachineStatusSummary) GetNetinOk() (*float32, bool)`

GetNetinOk returns a tuple with the Netin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetin

`func (o *VirtualMachineStatusSummary) SetNetin(v float32)`

SetNetin sets Netin field to given value.

### HasNetin

`func (o *VirtualMachineStatusSummary) HasNetin() bool`

HasNetin returns a boolean if a field has been set.

### GetNetout

`func (o *VirtualMachineStatusSummary) GetNetout() float32`

GetNetout returns the Netout field if non-nil, zero value otherwise.

### GetNetoutOk

`func (o *VirtualMachineStatusSummary) GetNetoutOk() (*float32, bool)`

GetNetoutOk returns a tuple with the Netout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetout

`func (o *VirtualMachineStatusSummary) SetNetout(v float32)`

SetNetout sets Netout field to given value.

### HasNetout

`func (o *VirtualMachineStatusSummary) HasNetout() bool`

HasNetout returns a boolean if a field has been set.

### GetDiskwrite

`func (o *VirtualMachineStatusSummary) GetDiskwrite() float32`

GetDiskwrite returns the Diskwrite field if non-nil, zero value otherwise.

### GetDiskwriteOk

`func (o *VirtualMachineStatusSummary) GetDiskwriteOk() (*float32, bool)`

GetDiskwriteOk returns a tuple with the Diskwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskwrite

`func (o *VirtualMachineStatusSummary) SetDiskwrite(v float32)`

SetDiskwrite sets Diskwrite field to given value.

### HasDiskwrite

`func (o *VirtualMachineStatusSummary) HasDiskwrite() bool`

HasDiskwrite returns a boolean if a field has been set.

### GetDiskread

`func (o *VirtualMachineStatusSummary) GetDiskread() float32`

GetDiskread returns the Diskread field if non-nil, zero value otherwise.

### GetDiskreadOk

`func (o *VirtualMachineStatusSummary) GetDiskreadOk() (*float32, bool)`

GetDiskreadOk returns a tuple with the Diskread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskread

`func (o *VirtualMachineStatusSummary) SetDiskread(v float32)`

SetDiskread sets Diskread field to given value.

### HasDiskread

`func (o *VirtualMachineStatusSummary) HasDiskread() bool`

HasDiskread returns a boolean if a field has been set.

### GetFreemem

`func (o *VirtualMachineStatusSummary) GetFreemem() float32`

GetFreemem returns the Freemem field if non-nil, zero value otherwise.

### GetFreememOk

`func (o *VirtualMachineStatusSummary) GetFreememOk() (*float32, bool)`

GetFreememOk returns a tuple with the Freemem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreemem

`func (o *VirtualMachineStatusSummary) SetFreemem(v float32)`

SetFreemem sets Freemem field to given value.

### HasFreemem

`func (o *VirtualMachineStatusSummary) HasFreemem() bool`

HasFreemem returns a boolean if a field has been set.

### GetNics

`func (o *VirtualMachineStatusSummary) GetNics() map[string]VirtualMachineNicStatus`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *VirtualMachineStatusSummary) GetNicsOk() (*map[string]VirtualMachineNicStatus, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *VirtualMachineStatusSummary) SetNics(v map[string]VirtualMachineNicStatus)`

SetNics sets Nics field to given value.

### HasNics

`func (o *VirtualMachineStatusSummary) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetBallooninfo

`func (o *VirtualMachineStatusSummary) GetBallooninfo() VirtualMachineBalloonSummary`

GetBallooninfo returns the Ballooninfo field if non-nil, zero value otherwise.

### GetBallooninfoOk

`func (o *VirtualMachineStatusSummary) GetBallooninfoOk() (*VirtualMachineBalloonSummary, bool)`

GetBallooninfoOk returns a tuple with the Ballooninfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBallooninfo

`func (o *VirtualMachineStatusSummary) SetBallooninfo(v VirtualMachineBalloonSummary)`

SetBallooninfo sets Ballooninfo field to given value.

### HasBallooninfo

`func (o *VirtualMachineStatusSummary) HasBallooninfo() bool`

HasBallooninfo returns a boolean if a field has been set.

### GetBlockstat

`func (o *VirtualMachineStatusSummary) GetBlockstat() map[string]VirtualMachineNicBlockStatistics`

GetBlockstat returns the Blockstat field if non-nil, zero value otherwise.

### GetBlockstatOk

`func (o *VirtualMachineStatusSummary) GetBlockstatOk() (*map[string]VirtualMachineNicBlockStatistics, bool)`

GetBlockstatOk returns a tuple with the Blockstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockstat

`func (o *VirtualMachineStatusSummary) SetBlockstat(v map[string]VirtualMachineNicBlockStatistics)`

SetBlockstat sets Blockstat field to given value.

### HasBlockstat

`func (o *VirtualMachineStatusSummary) HasBlockstat() bool`

HasBlockstat returns a boolean if a field has been set.

### GetProxmoxSupport

`func (o *VirtualMachineStatusSummary) GetProxmoxSupport() ProxmoxSupportSummary`

GetProxmoxSupport returns the ProxmoxSupport field if non-nil, zero value otherwise.

### GetProxmoxSupportOk

`func (o *VirtualMachineStatusSummary) GetProxmoxSupportOk() (*ProxmoxSupportSummary, bool)`

GetProxmoxSupportOk returns a tuple with the ProxmoxSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxmoxSupport

`func (o *VirtualMachineStatusSummary) SetProxmoxSupport(v ProxmoxSupportSummary)`

SetProxmoxSupport sets ProxmoxSupport field to given value.

### HasProxmoxSupport

`func (o *VirtualMachineStatusSummary) HasProxmoxSupport() bool`

HasProxmoxSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


