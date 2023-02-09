# PoolMemberSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**PoolMemberType**](PoolMemberType.md) |  | 
**Disk** | Pointer to **float32** |  | [optional] 
**Uptime** | Pointer to **float32** |  | [optional] 
**Maxmem** | Pointer to **float32** |  | [optional] 
**Netin** | Pointer to **float32** |  | [optional] 
**Netout** | Pointer to **float32** |  | [optional] 
**Diskwrite** | Pointer to **float32** |  | [optional] 
**Diskread** | Pointer to **float32** |  | [optional] 
**Maxdisk** | Pointer to **float32** |  | [optional] 
**Cpu** | Pointer to **float32** |  | [optional] 
**Maxcpu** | Pointer to **float32** |  | [optional] 
**Vmid** | Pointer to **float32** |  | [optional] 
**Mem** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **float32** | Whether the VM is a template. This is a boolean integer, where 1 is true and 0 is false. | [optional] 
**Node** | Pointer to **string** |  | [optional] 
**Storage** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Plugintype** | Pointer to **string** |  | [optional] 
**Shared** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewPoolMemberSummary

`func NewPoolMemberSummary(id string, type_ PoolMemberType, ) *PoolMemberSummary`

NewPoolMemberSummary instantiates a new PoolMemberSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolMemberSummaryWithDefaults

`func NewPoolMemberSummaryWithDefaults() *PoolMemberSummary`

NewPoolMemberSummaryWithDefaults instantiates a new PoolMemberSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoolMemberSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoolMemberSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoolMemberSummary) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PoolMemberSummary) GetType() PoolMemberType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PoolMemberSummary) GetTypeOk() (*PoolMemberType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PoolMemberSummary) SetType(v PoolMemberType)`

SetType sets Type field to given value.


### GetDisk

`func (o *PoolMemberSummary) GetDisk() float32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *PoolMemberSummary) GetDiskOk() (*float32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *PoolMemberSummary) SetDisk(v float32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *PoolMemberSummary) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetUptime

`func (o *PoolMemberSummary) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *PoolMemberSummary) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *PoolMemberSummary) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *PoolMemberSummary) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetMaxmem

`func (o *PoolMemberSummary) GetMaxmem() float32`

GetMaxmem returns the Maxmem field if non-nil, zero value otherwise.

### GetMaxmemOk

`func (o *PoolMemberSummary) GetMaxmemOk() (*float32, bool)`

GetMaxmemOk returns a tuple with the Maxmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxmem

`func (o *PoolMemberSummary) SetMaxmem(v float32)`

SetMaxmem sets Maxmem field to given value.

### HasMaxmem

`func (o *PoolMemberSummary) HasMaxmem() bool`

HasMaxmem returns a boolean if a field has been set.

### GetNetin

`func (o *PoolMemberSummary) GetNetin() float32`

GetNetin returns the Netin field if non-nil, zero value otherwise.

### GetNetinOk

`func (o *PoolMemberSummary) GetNetinOk() (*float32, bool)`

GetNetinOk returns a tuple with the Netin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetin

`func (o *PoolMemberSummary) SetNetin(v float32)`

SetNetin sets Netin field to given value.

### HasNetin

`func (o *PoolMemberSummary) HasNetin() bool`

HasNetin returns a boolean if a field has been set.

### GetNetout

`func (o *PoolMemberSummary) GetNetout() float32`

GetNetout returns the Netout field if non-nil, zero value otherwise.

### GetNetoutOk

`func (o *PoolMemberSummary) GetNetoutOk() (*float32, bool)`

GetNetoutOk returns a tuple with the Netout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetout

`func (o *PoolMemberSummary) SetNetout(v float32)`

SetNetout sets Netout field to given value.

### HasNetout

`func (o *PoolMemberSummary) HasNetout() bool`

HasNetout returns a boolean if a field has been set.

### GetDiskwrite

`func (o *PoolMemberSummary) GetDiskwrite() float32`

GetDiskwrite returns the Diskwrite field if non-nil, zero value otherwise.

### GetDiskwriteOk

`func (o *PoolMemberSummary) GetDiskwriteOk() (*float32, bool)`

GetDiskwriteOk returns a tuple with the Diskwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskwrite

`func (o *PoolMemberSummary) SetDiskwrite(v float32)`

SetDiskwrite sets Diskwrite field to given value.

### HasDiskwrite

`func (o *PoolMemberSummary) HasDiskwrite() bool`

HasDiskwrite returns a boolean if a field has been set.

### GetDiskread

`func (o *PoolMemberSummary) GetDiskread() float32`

GetDiskread returns the Diskread field if non-nil, zero value otherwise.

### GetDiskreadOk

`func (o *PoolMemberSummary) GetDiskreadOk() (*float32, bool)`

GetDiskreadOk returns a tuple with the Diskread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskread

`func (o *PoolMemberSummary) SetDiskread(v float32)`

SetDiskread sets Diskread field to given value.

### HasDiskread

`func (o *PoolMemberSummary) HasDiskread() bool`

HasDiskread returns a boolean if a field has been set.

### GetMaxdisk

`func (o *PoolMemberSummary) GetMaxdisk() float32`

GetMaxdisk returns the Maxdisk field if non-nil, zero value otherwise.

### GetMaxdiskOk

`func (o *PoolMemberSummary) GetMaxdiskOk() (*float32, bool)`

GetMaxdiskOk returns a tuple with the Maxdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxdisk

`func (o *PoolMemberSummary) SetMaxdisk(v float32)`

SetMaxdisk sets Maxdisk field to given value.

### HasMaxdisk

`func (o *PoolMemberSummary) HasMaxdisk() bool`

HasMaxdisk returns a boolean if a field has been set.

### GetCpu

`func (o *PoolMemberSummary) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *PoolMemberSummary) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *PoolMemberSummary) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *PoolMemberSummary) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMaxcpu

`func (o *PoolMemberSummary) GetMaxcpu() float32`

GetMaxcpu returns the Maxcpu field if non-nil, zero value otherwise.

### GetMaxcpuOk

`func (o *PoolMemberSummary) GetMaxcpuOk() (*float32, bool)`

GetMaxcpuOk returns a tuple with the Maxcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxcpu

`func (o *PoolMemberSummary) SetMaxcpu(v float32)`

SetMaxcpu sets Maxcpu field to given value.

### HasMaxcpu

`func (o *PoolMemberSummary) HasMaxcpu() bool`

HasMaxcpu returns a boolean if a field has been set.

### GetVmid

`func (o *PoolMemberSummary) GetVmid() float32`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *PoolMemberSummary) GetVmidOk() (*float32, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *PoolMemberSummary) SetVmid(v float32)`

SetVmid sets Vmid field to given value.

### HasVmid

`func (o *PoolMemberSummary) HasVmid() bool`

HasVmid returns a boolean if a field has been set.

### GetMem

`func (o *PoolMemberSummary) GetMem() float32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *PoolMemberSummary) GetMemOk() (*float32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *PoolMemberSummary) SetMem(v float32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *PoolMemberSummary) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetName

`func (o *PoolMemberSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolMemberSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolMemberSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolMemberSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplate

`func (o *PoolMemberSummary) GetTemplate() float32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PoolMemberSummary) GetTemplateOk() (*float32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PoolMemberSummary) SetTemplate(v float32)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PoolMemberSummary) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetNode

`func (o *PoolMemberSummary) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *PoolMemberSummary) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *PoolMemberSummary) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *PoolMemberSummary) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetStorage

`func (o *PoolMemberSummary) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *PoolMemberSummary) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *PoolMemberSummary) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *PoolMemberSummary) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetStatus

`func (o *PoolMemberSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoolMemberSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoolMemberSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PoolMemberSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPlugintype

`func (o *PoolMemberSummary) GetPlugintype() string`

GetPlugintype returns the Plugintype field if non-nil, zero value otherwise.

### GetPlugintypeOk

`func (o *PoolMemberSummary) GetPlugintypeOk() (*string, bool)`

GetPlugintypeOk returns a tuple with the Plugintype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugintype

`func (o *PoolMemberSummary) SetPlugintype(v string)`

SetPlugintype sets Plugintype field to given value.

### HasPlugintype

`func (o *PoolMemberSummary) HasPlugintype() bool`

HasPlugintype returns a boolean if a field has been set.

### GetShared

`func (o *PoolMemberSummary) GetShared() float32`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *PoolMemberSummary) GetSharedOk() (*float32, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *PoolMemberSummary) SetShared(v float32)`

SetShared sets Shared field to given value.

### HasShared

`func (o *PoolMemberSummary) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetContent

`func (o *PoolMemberSummary) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PoolMemberSummary) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PoolMemberSummary) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *PoolMemberSummary) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


