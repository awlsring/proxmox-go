# RolePermissionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreAllocate** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**DatastoreAllocateSpace** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**DatastoreAudit** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**DatastoreAllocateTemplate** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**GroupAllocate** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**PermissionModify** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**PoolAllocate** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**PoolAudit** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**RealmAllocate** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**RealmAllocateUser** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**SDNAllocate** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**SDNAudit** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**SysAudit** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**SysModify** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**SysConsole** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**SysIncoming** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**SysPowerMgmt** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**SysSyslog** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**UserModify** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMAllocate** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMAudit** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMBackup** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMClone** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMConfigCDROM** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMConfigCPU** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMConfigCloudinit** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMConfigDisk** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMConfigHWType** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMConfigMemory** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMConfigNetwork** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMConfigOptions** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMConsole** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMMigrate** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMMonitor** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMPowerMgmt** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMSnapshot** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**VMSnapshotRollback** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 

## Methods

### NewRolePermissionSummary

`func NewRolePermissionSummary() *RolePermissionSummary`

NewRolePermissionSummary instantiates a new RolePermissionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionSummaryWithDefaults

`func NewRolePermissionSummaryWithDefaults() *RolePermissionSummary`

NewRolePermissionSummaryWithDefaults instantiates a new RolePermissionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreAllocate

`func (o *RolePermissionSummary) GetDatastoreAllocate() float32`

GetDatastoreAllocate returns the DatastoreAllocate field if non-nil, zero value otherwise.

### GetDatastoreAllocateOk

`func (o *RolePermissionSummary) GetDatastoreAllocateOk() (*float32, bool)`

GetDatastoreAllocateOk returns a tuple with the DatastoreAllocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreAllocate

`func (o *RolePermissionSummary) SetDatastoreAllocate(v float32)`

SetDatastoreAllocate sets DatastoreAllocate field to given value.

### HasDatastoreAllocate

`func (o *RolePermissionSummary) HasDatastoreAllocate() bool`

HasDatastoreAllocate returns a boolean if a field has been set.

### GetDatastoreAllocateSpace

`func (o *RolePermissionSummary) GetDatastoreAllocateSpace() float32`

GetDatastoreAllocateSpace returns the DatastoreAllocateSpace field if non-nil, zero value otherwise.

### GetDatastoreAllocateSpaceOk

`func (o *RolePermissionSummary) GetDatastoreAllocateSpaceOk() (*float32, bool)`

GetDatastoreAllocateSpaceOk returns a tuple with the DatastoreAllocateSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreAllocateSpace

`func (o *RolePermissionSummary) SetDatastoreAllocateSpace(v float32)`

SetDatastoreAllocateSpace sets DatastoreAllocateSpace field to given value.

### HasDatastoreAllocateSpace

`func (o *RolePermissionSummary) HasDatastoreAllocateSpace() bool`

HasDatastoreAllocateSpace returns a boolean if a field has been set.

### GetDatastoreAudit

`func (o *RolePermissionSummary) GetDatastoreAudit() float32`

GetDatastoreAudit returns the DatastoreAudit field if non-nil, zero value otherwise.

### GetDatastoreAuditOk

`func (o *RolePermissionSummary) GetDatastoreAuditOk() (*float32, bool)`

GetDatastoreAuditOk returns a tuple with the DatastoreAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreAudit

`func (o *RolePermissionSummary) SetDatastoreAudit(v float32)`

SetDatastoreAudit sets DatastoreAudit field to given value.

### HasDatastoreAudit

`func (o *RolePermissionSummary) HasDatastoreAudit() bool`

HasDatastoreAudit returns a boolean if a field has been set.

### GetDatastoreAllocateTemplate

`func (o *RolePermissionSummary) GetDatastoreAllocateTemplate() float32`

GetDatastoreAllocateTemplate returns the DatastoreAllocateTemplate field if non-nil, zero value otherwise.

### GetDatastoreAllocateTemplateOk

`func (o *RolePermissionSummary) GetDatastoreAllocateTemplateOk() (*float32, bool)`

GetDatastoreAllocateTemplateOk returns a tuple with the DatastoreAllocateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreAllocateTemplate

`func (o *RolePermissionSummary) SetDatastoreAllocateTemplate(v float32)`

SetDatastoreAllocateTemplate sets DatastoreAllocateTemplate field to given value.

### HasDatastoreAllocateTemplate

`func (o *RolePermissionSummary) HasDatastoreAllocateTemplate() bool`

HasDatastoreAllocateTemplate returns a boolean if a field has been set.

### GetGroupAllocate

`func (o *RolePermissionSummary) GetGroupAllocate() float32`

GetGroupAllocate returns the GroupAllocate field if non-nil, zero value otherwise.

### GetGroupAllocateOk

`func (o *RolePermissionSummary) GetGroupAllocateOk() (*float32, bool)`

GetGroupAllocateOk returns a tuple with the GroupAllocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAllocate

`func (o *RolePermissionSummary) SetGroupAllocate(v float32)`

SetGroupAllocate sets GroupAllocate field to given value.

### HasGroupAllocate

`func (o *RolePermissionSummary) HasGroupAllocate() bool`

HasGroupAllocate returns a boolean if a field has been set.

### GetPermissionModify

`func (o *RolePermissionSummary) GetPermissionModify() float32`

GetPermissionModify returns the PermissionModify field if non-nil, zero value otherwise.

### GetPermissionModifyOk

`func (o *RolePermissionSummary) GetPermissionModifyOk() (*float32, bool)`

GetPermissionModifyOk returns a tuple with the PermissionModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionModify

`func (o *RolePermissionSummary) SetPermissionModify(v float32)`

SetPermissionModify sets PermissionModify field to given value.

### HasPermissionModify

`func (o *RolePermissionSummary) HasPermissionModify() bool`

HasPermissionModify returns a boolean if a field has been set.

### GetPoolAllocate

`func (o *RolePermissionSummary) GetPoolAllocate() float32`

GetPoolAllocate returns the PoolAllocate field if non-nil, zero value otherwise.

### GetPoolAllocateOk

`func (o *RolePermissionSummary) GetPoolAllocateOk() (*float32, bool)`

GetPoolAllocateOk returns a tuple with the PoolAllocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAllocate

`func (o *RolePermissionSummary) SetPoolAllocate(v float32)`

SetPoolAllocate sets PoolAllocate field to given value.

### HasPoolAllocate

`func (o *RolePermissionSummary) HasPoolAllocate() bool`

HasPoolAllocate returns a boolean if a field has been set.

### GetPoolAudit

`func (o *RolePermissionSummary) GetPoolAudit() float32`

GetPoolAudit returns the PoolAudit field if non-nil, zero value otherwise.

### GetPoolAuditOk

`func (o *RolePermissionSummary) GetPoolAuditOk() (*float32, bool)`

GetPoolAuditOk returns a tuple with the PoolAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAudit

`func (o *RolePermissionSummary) SetPoolAudit(v float32)`

SetPoolAudit sets PoolAudit field to given value.

### HasPoolAudit

`func (o *RolePermissionSummary) HasPoolAudit() bool`

HasPoolAudit returns a boolean if a field has been set.

### GetRealmAllocate

`func (o *RolePermissionSummary) GetRealmAllocate() float32`

GetRealmAllocate returns the RealmAllocate field if non-nil, zero value otherwise.

### GetRealmAllocateOk

`func (o *RolePermissionSummary) GetRealmAllocateOk() (*float32, bool)`

GetRealmAllocateOk returns a tuple with the RealmAllocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmAllocate

`func (o *RolePermissionSummary) SetRealmAllocate(v float32)`

SetRealmAllocate sets RealmAllocate field to given value.

### HasRealmAllocate

`func (o *RolePermissionSummary) HasRealmAllocate() bool`

HasRealmAllocate returns a boolean if a field has been set.

### GetRealmAllocateUser

`func (o *RolePermissionSummary) GetRealmAllocateUser() float32`

GetRealmAllocateUser returns the RealmAllocateUser field if non-nil, zero value otherwise.

### GetRealmAllocateUserOk

`func (o *RolePermissionSummary) GetRealmAllocateUserOk() (*float32, bool)`

GetRealmAllocateUserOk returns a tuple with the RealmAllocateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmAllocateUser

`func (o *RolePermissionSummary) SetRealmAllocateUser(v float32)`

SetRealmAllocateUser sets RealmAllocateUser field to given value.

### HasRealmAllocateUser

`func (o *RolePermissionSummary) HasRealmAllocateUser() bool`

HasRealmAllocateUser returns a boolean if a field has been set.

### GetSDNAllocate

`func (o *RolePermissionSummary) GetSDNAllocate() float32`

GetSDNAllocate returns the SDNAllocate field if non-nil, zero value otherwise.

### GetSDNAllocateOk

`func (o *RolePermissionSummary) GetSDNAllocateOk() (*float32, bool)`

GetSDNAllocateOk returns a tuple with the SDNAllocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDNAllocate

`func (o *RolePermissionSummary) SetSDNAllocate(v float32)`

SetSDNAllocate sets SDNAllocate field to given value.

### HasSDNAllocate

`func (o *RolePermissionSummary) HasSDNAllocate() bool`

HasSDNAllocate returns a boolean if a field has been set.

### GetSDNAudit

`func (o *RolePermissionSummary) GetSDNAudit() float32`

GetSDNAudit returns the SDNAudit field if non-nil, zero value otherwise.

### GetSDNAuditOk

`func (o *RolePermissionSummary) GetSDNAuditOk() (*float32, bool)`

GetSDNAuditOk returns a tuple with the SDNAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDNAudit

`func (o *RolePermissionSummary) SetSDNAudit(v float32)`

SetSDNAudit sets SDNAudit field to given value.

### HasSDNAudit

`func (o *RolePermissionSummary) HasSDNAudit() bool`

HasSDNAudit returns a boolean if a field has been set.

### GetSysAudit

`func (o *RolePermissionSummary) GetSysAudit() float32`

GetSysAudit returns the SysAudit field if non-nil, zero value otherwise.

### GetSysAuditOk

`func (o *RolePermissionSummary) GetSysAuditOk() (*float32, bool)`

GetSysAuditOk returns a tuple with the SysAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysAudit

`func (o *RolePermissionSummary) SetSysAudit(v float32)`

SetSysAudit sets SysAudit field to given value.

### HasSysAudit

`func (o *RolePermissionSummary) HasSysAudit() bool`

HasSysAudit returns a boolean if a field has been set.

### GetSysModify

`func (o *RolePermissionSummary) GetSysModify() float32`

GetSysModify returns the SysModify field if non-nil, zero value otherwise.

### GetSysModifyOk

`func (o *RolePermissionSummary) GetSysModifyOk() (*float32, bool)`

GetSysModifyOk returns a tuple with the SysModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysModify

`func (o *RolePermissionSummary) SetSysModify(v float32)`

SetSysModify sets SysModify field to given value.

### HasSysModify

`func (o *RolePermissionSummary) HasSysModify() bool`

HasSysModify returns a boolean if a field has been set.

### GetSysConsole

`func (o *RolePermissionSummary) GetSysConsole() float32`

GetSysConsole returns the SysConsole field if non-nil, zero value otherwise.

### GetSysConsoleOk

`func (o *RolePermissionSummary) GetSysConsoleOk() (*float32, bool)`

GetSysConsoleOk returns a tuple with the SysConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysConsole

`func (o *RolePermissionSummary) SetSysConsole(v float32)`

SetSysConsole sets SysConsole field to given value.

### HasSysConsole

`func (o *RolePermissionSummary) HasSysConsole() bool`

HasSysConsole returns a boolean if a field has been set.

### GetSysIncoming

`func (o *RolePermissionSummary) GetSysIncoming() float32`

GetSysIncoming returns the SysIncoming field if non-nil, zero value otherwise.

### GetSysIncomingOk

`func (o *RolePermissionSummary) GetSysIncomingOk() (*float32, bool)`

GetSysIncomingOk returns a tuple with the SysIncoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysIncoming

`func (o *RolePermissionSummary) SetSysIncoming(v float32)`

SetSysIncoming sets SysIncoming field to given value.

### HasSysIncoming

`func (o *RolePermissionSummary) HasSysIncoming() bool`

HasSysIncoming returns a boolean if a field has been set.

### GetSysPowerMgmt

`func (o *RolePermissionSummary) GetSysPowerMgmt() float32`

GetSysPowerMgmt returns the SysPowerMgmt field if non-nil, zero value otherwise.

### GetSysPowerMgmtOk

`func (o *RolePermissionSummary) GetSysPowerMgmtOk() (*float32, bool)`

GetSysPowerMgmtOk returns a tuple with the SysPowerMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysPowerMgmt

`func (o *RolePermissionSummary) SetSysPowerMgmt(v float32)`

SetSysPowerMgmt sets SysPowerMgmt field to given value.

### HasSysPowerMgmt

`func (o *RolePermissionSummary) HasSysPowerMgmt() bool`

HasSysPowerMgmt returns a boolean if a field has been set.

### GetSysSyslog

`func (o *RolePermissionSummary) GetSysSyslog() float32`

GetSysSyslog returns the SysSyslog field if non-nil, zero value otherwise.

### GetSysSyslogOk

`func (o *RolePermissionSummary) GetSysSyslogOk() (*float32, bool)`

GetSysSyslogOk returns a tuple with the SysSyslog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysSyslog

`func (o *RolePermissionSummary) SetSysSyslog(v float32)`

SetSysSyslog sets SysSyslog field to given value.

### HasSysSyslog

`func (o *RolePermissionSummary) HasSysSyslog() bool`

HasSysSyslog returns a boolean if a field has been set.

### GetUserModify

`func (o *RolePermissionSummary) GetUserModify() float32`

GetUserModify returns the UserModify field if non-nil, zero value otherwise.

### GetUserModifyOk

`func (o *RolePermissionSummary) GetUserModifyOk() (*float32, bool)`

GetUserModifyOk returns a tuple with the UserModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserModify

`func (o *RolePermissionSummary) SetUserModify(v float32)`

SetUserModify sets UserModify field to given value.

### HasUserModify

`func (o *RolePermissionSummary) HasUserModify() bool`

HasUserModify returns a boolean if a field has been set.

### GetVMAllocate

`func (o *RolePermissionSummary) GetVMAllocate() float32`

GetVMAllocate returns the VMAllocate field if non-nil, zero value otherwise.

### GetVMAllocateOk

`func (o *RolePermissionSummary) GetVMAllocateOk() (*float32, bool)`

GetVMAllocateOk returns a tuple with the VMAllocate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMAllocate

`func (o *RolePermissionSummary) SetVMAllocate(v float32)`

SetVMAllocate sets VMAllocate field to given value.

### HasVMAllocate

`func (o *RolePermissionSummary) HasVMAllocate() bool`

HasVMAllocate returns a boolean if a field has been set.

### GetVMAudit

`func (o *RolePermissionSummary) GetVMAudit() float32`

GetVMAudit returns the VMAudit field if non-nil, zero value otherwise.

### GetVMAuditOk

`func (o *RolePermissionSummary) GetVMAuditOk() (*float32, bool)`

GetVMAuditOk returns a tuple with the VMAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMAudit

`func (o *RolePermissionSummary) SetVMAudit(v float32)`

SetVMAudit sets VMAudit field to given value.

### HasVMAudit

`func (o *RolePermissionSummary) HasVMAudit() bool`

HasVMAudit returns a boolean if a field has been set.

### GetVMBackup

`func (o *RolePermissionSummary) GetVMBackup() float32`

GetVMBackup returns the VMBackup field if non-nil, zero value otherwise.

### GetVMBackupOk

`func (o *RolePermissionSummary) GetVMBackupOk() (*float32, bool)`

GetVMBackupOk returns a tuple with the VMBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMBackup

`func (o *RolePermissionSummary) SetVMBackup(v float32)`

SetVMBackup sets VMBackup field to given value.

### HasVMBackup

`func (o *RolePermissionSummary) HasVMBackup() bool`

HasVMBackup returns a boolean if a field has been set.

### GetVMClone

`func (o *RolePermissionSummary) GetVMClone() float32`

GetVMClone returns the VMClone field if non-nil, zero value otherwise.

### GetVMCloneOk

`func (o *RolePermissionSummary) GetVMCloneOk() (*float32, bool)`

GetVMCloneOk returns a tuple with the VMClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMClone

`func (o *RolePermissionSummary) SetVMClone(v float32)`

SetVMClone sets VMClone field to given value.

### HasVMClone

`func (o *RolePermissionSummary) HasVMClone() bool`

HasVMClone returns a boolean if a field has been set.

### GetVMConfigCDROM

`func (o *RolePermissionSummary) GetVMConfigCDROM() float32`

GetVMConfigCDROM returns the VMConfigCDROM field if non-nil, zero value otherwise.

### GetVMConfigCDROMOk

`func (o *RolePermissionSummary) GetVMConfigCDROMOk() (*float32, bool)`

GetVMConfigCDROMOk returns a tuple with the VMConfigCDROM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMConfigCDROM

`func (o *RolePermissionSummary) SetVMConfigCDROM(v float32)`

SetVMConfigCDROM sets VMConfigCDROM field to given value.

### HasVMConfigCDROM

`func (o *RolePermissionSummary) HasVMConfigCDROM() bool`

HasVMConfigCDROM returns a boolean if a field has been set.

### GetVMConfigCPU

`func (o *RolePermissionSummary) GetVMConfigCPU() float32`

GetVMConfigCPU returns the VMConfigCPU field if non-nil, zero value otherwise.

### GetVMConfigCPUOk

`func (o *RolePermissionSummary) GetVMConfigCPUOk() (*float32, bool)`

GetVMConfigCPUOk returns a tuple with the VMConfigCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMConfigCPU

`func (o *RolePermissionSummary) SetVMConfigCPU(v float32)`

SetVMConfigCPU sets VMConfigCPU field to given value.

### HasVMConfigCPU

`func (o *RolePermissionSummary) HasVMConfigCPU() bool`

HasVMConfigCPU returns a boolean if a field has been set.

### GetVMConfigCloudinit

`func (o *RolePermissionSummary) GetVMConfigCloudinit() float32`

GetVMConfigCloudinit returns the VMConfigCloudinit field if non-nil, zero value otherwise.

### GetVMConfigCloudinitOk

`func (o *RolePermissionSummary) GetVMConfigCloudinitOk() (*float32, bool)`

GetVMConfigCloudinitOk returns a tuple with the VMConfigCloudinit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMConfigCloudinit

`func (o *RolePermissionSummary) SetVMConfigCloudinit(v float32)`

SetVMConfigCloudinit sets VMConfigCloudinit field to given value.

### HasVMConfigCloudinit

`func (o *RolePermissionSummary) HasVMConfigCloudinit() bool`

HasVMConfigCloudinit returns a boolean if a field has been set.

### GetVMConfigDisk

`func (o *RolePermissionSummary) GetVMConfigDisk() float32`

GetVMConfigDisk returns the VMConfigDisk field if non-nil, zero value otherwise.

### GetVMConfigDiskOk

`func (o *RolePermissionSummary) GetVMConfigDiskOk() (*float32, bool)`

GetVMConfigDiskOk returns a tuple with the VMConfigDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMConfigDisk

`func (o *RolePermissionSummary) SetVMConfigDisk(v float32)`

SetVMConfigDisk sets VMConfigDisk field to given value.

### HasVMConfigDisk

`func (o *RolePermissionSummary) HasVMConfigDisk() bool`

HasVMConfigDisk returns a boolean if a field has been set.

### GetVMConfigHWType

`func (o *RolePermissionSummary) GetVMConfigHWType() float32`

GetVMConfigHWType returns the VMConfigHWType field if non-nil, zero value otherwise.

### GetVMConfigHWTypeOk

`func (o *RolePermissionSummary) GetVMConfigHWTypeOk() (*float32, bool)`

GetVMConfigHWTypeOk returns a tuple with the VMConfigHWType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMConfigHWType

`func (o *RolePermissionSummary) SetVMConfigHWType(v float32)`

SetVMConfigHWType sets VMConfigHWType field to given value.

### HasVMConfigHWType

`func (o *RolePermissionSummary) HasVMConfigHWType() bool`

HasVMConfigHWType returns a boolean if a field has been set.

### GetVMConfigMemory

`func (o *RolePermissionSummary) GetVMConfigMemory() float32`

GetVMConfigMemory returns the VMConfigMemory field if non-nil, zero value otherwise.

### GetVMConfigMemoryOk

`func (o *RolePermissionSummary) GetVMConfigMemoryOk() (*float32, bool)`

GetVMConfigMemoryOk returns a tuple with the VMConfigMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMConfigMemory

`func (o *RolePermissionSummary) SetVMConfigMemory(v float32)`

SetVMConfigMemory sets VMConfigMemory field to given value.

### HasVMConfigMemory

`func (o *RolePermissionSummary) HasVMConfigMemory() bool`

HasVMConfigMemory returns a boolean if a field has been set.

### GetVMConfigNetwork

`func (o *RolePermissionSummary) GetVMConfigNetwork() float32`

GetVMConfigNetwork returns the VMConfigNetwork field if non-nil, zero value otherwise.

### GetVMConfigNetworkOk

`func (o *RolePermissionSummary) GetVMConfigNetworkOk() (*float32, bool)`

GetVMConfigNetworkOk returns a tuple with the VMConfigNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMConfigNetwork

`func (o *RolePermissionSummary) SetVMConfigNetwork(v float32)`

SetVMConfigNetwork sets VMConfigNetwork field to given value.

### HasVMConfigNetwork

`func (o *RolePermissionSummary) HasVMConfigNetwork() bool`

HasVMConfigNetwork returns a boolean if a field has been set.

### GetVMConfigOptions

`func (o *RolePermissionSummary) GetVMConfigOptions() float32`

GetVMConfigOptions returns the VMConfigOptions field if non-nil, zero value otherwise.

### GetVMConfigOptionsOk

`func (o *RolePermissionSummary) GetVMConfigOptionsOk() (*float32, bool)`

GetVMConfigOptionsOk returns a tuple with the VMConfigOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMConfigOptions

`func (o *RolePermissionSummary) SetVMConfigOptions(v float32)`

SetVMConfigOptions sets VMConfigOptions field to given value.

### HasVMConfigOptions

`func (o *RolePermissionSummary) HasVMConfigOptions() bool`

HasVMConfigOptions returns a boolean if a field has been set.

### GetVMConsole

`func (o *RolePermissionSummary) GetVMConsole() float32`

GetVMConsole returns the VMConsole field if non-nil, zero value otherwise.

### GetVMConsoleOk

`func (o *RolePermissionSummary) GetVMConsoleOk() (*float32, bool)`

GetVMConsoleOk returns a tuple with the VMConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMConsole

`func (o *RolePermissionSummary) SetVMConsole(v float32)`

SetVMConsole sets VMConsole field to given value.

### HasVMConsole

`func (o *RolePermissionSummary) HasVMConsole() bool`

HasVMConsole returns a boolean if a field has been set.

### GetVMMigrate

`func (o *RolePermissionSummary) GetVMMigrate() float32`

GetVMMigrate returns the VMMigrate field if non-nil, zero value otherwise.

### GetVMMigrateOk

`func (o *RolePermissionSummary) GetVMMigrateOk() (*float32, bool)`

GetVMMigrateOk returns a tuple with the VMMigrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMMigrate

`func (o *RolePermissionSummary) SetVMMigrate(v float32)`

SetVMMigrate sets VMMigrate field to given value.

### HasVMMigrate

`func (o *RolePermissionSummary) HasVMMigrate() bool`

HasVMMigrate returns a boolean if a field has been set.

### GetVMMonitor

`func (o *RolePermissionSummary) GetVMMonitor() float32`

GetVMMonitor returns the VMMonitor field if non-nil, zero value otherwise.

### GetVMMonitorOk

`func (o *RolePermissionSummary) GetVMMonitorOk() (*float32, bool)`

GetVMMonitorOk returns a tuple with the VMMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMMonitor

`func (o *RolePermissionSummary) SetVMMonitor(v float32)`

SetVMMonitor sets VMMonitor field to given value.

### HasVMMonitor

`func (o *RolePermissionSummary) HasVMMonitor() bool`

HasVMMonitor returns a boolean if a field has been set.

### GetVMPowerMgmt

`func (o *RolePermissionSummary) GetVMPowerMgmt() float32`

GetVMPowerMgmt returns the VMPowerMgmt field if non-nil, zero value otherwise.

### GetVMPowerMgmtOk

`func (o *RolePermissionSummary) GetVMPowerMgmtOk() (*float32, bool)`

GetVMPowerMgmtOk returns a tuple with the VMPowerMgmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMPowerMgmt

`func (o *RolePermissionSummary) SetVMPowerMgmt(v float32)`

SetVMPowerMgmt sets VMPowerMgmt field to given value.

### HasVMPowerMgmt

`func (o *RolePermissionSummary) HasVMPowerMgmt() bool`

HasVMPowerMgmt returns a boolean if a field has been set.

### GetVMSnapshot

`func (o *RolePermissionSummary) GetVMSnapshot() float32`

GetVMSnapshot returns the VMSnapshot field if non-nil, zero value otherwise.

### GetVMSnapshotOk

`func (o *RolePermissionSummary) GetVMSnapshotOk() (*float32, bool)`

GetVMSnapshotOk returns a tuple with the VMSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMSnapshot

`func (o *RolePermissionSummary) SetVMSnapshot(v float32)`

SetVMSnapshot sets VMSnapshot field to given value.

### HasVMSnapshot

`func (o *RolePermissionSummary) HasVMSnapshot() bool`

HasVMSnapshot returns a boolean if a field has been set.

### GetVMSnapshotRollback

`func (o *RolePermissionSummary) GetVMSnapshotRollback() float32`

GetVMSnapshotRollback returns the VMSnapshotRollback field if non-nil, zero value otherwise.

### GetVMSnapshotRollbackOk

`func (o *RolePermissionSummary) GetVMSnapshotRollbackOk() (*float32, bool)`

GetVMSnapshotRollbackOk returns a tuple with the VMSnapshotRollback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMSnapshotRollback

`func (o *RolePermissionSummary) SetVMSnapshotRollback(v float32)`

SetVMSnapshotRollback sets VMSnapshotRollback field to given value.

### HasVMSnapshotRollback

`func (o *RolePermissionSummary) HasVMSnapshotRollback() bool`

HasVMSnapshotRollback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


