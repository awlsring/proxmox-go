# StartVirtualMachineRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **float32** |  | [optional] 
**Skiplock** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**ForceCpu** | Pointer to **string** | Override QEMU &#x60;-cpu&#x60; argument. | [optional] 
**Machine** | Pointer to **string** | Specify QEMU machine type. | [optional] 
**Migratedfrom** | Pointer to **string** | Specify migrated from node. | [optional] 
**MigrationNetwork** | Pointer to **string** | Specify migration network. | [optional] 
**MigrationType** | Pointer to [**VirtualMachineMigrationType**](VirtualMachineMigrationType.md) |  | [optional] 
**Stateuri** | Pointer to **string** | Save/restore state from the URI. | [optional] 
**Targetstorage** | Pointer to **string** | Specify target storage. | [optional] 

## Methods

### NewStartVirtualMachineRequestContent

`func NewStartVirtualMachineRequestContent() *StartVirtualMachineRequestContent`

NewStartVirtualMachineRequestContent instantiates a new StartVirtualMachineRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartVirtualMachineRequestContentWithDefaults

`func NewStartVirtualMachineRequestContentWithDefaults() *StartVirtualMachineRequestContent`

NewStartVirtualMachineRequestContentWithDefaults instantiates a new StartVirtualMachineRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *StartVirtualMachineRequestContent) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *StartVirtualMachineRequestContent) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *StartVirtualMachineRequestContent) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *StartVirtualMachineRequestContent) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetSkiplock

`func (o *StartVirtualMachineRequestContent) GetSkiplock() float32`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *StartVirtualMachineRequestContent) GetSkiplockOk() (*float32, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *StartVirtualMachineRequestContent) SetSkiplock(v float32)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *StartVirtualMachineRequestContent) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.

### GetForceCpu

`func (o *StartVirtualMachineRequestContent) GetForceCpu() string`

GetForceCpu returns the ForceCpu field if non-nil, zero value otherwise.

### GetForceCpuOk

`func (o *StartVirtualMachineRequestContent) GetForceCpuOk() (*string, bool)`

GetForceCpuOk returns a tuple with the ForceCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceCpu

`func (o *StartVirtualMachineRequestContent) SetForceCpu(v string)`

SetForceCpu sets ForceCpu field to given value.

### HasForceCpu

`func (o *StartVirtualMachineRequestContent) HasForceCpu() bool`

HasForceCpu returns a boolean if a field has been set.

### GetMachine

`func (o *StartVirtualMachineRequestContent) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *StartVirtualMachineRequestContent) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *StartVirtualMachineRequestContent) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *StartVirtualMachineRequestContent) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetMigratedfrom

`func (o *StartVirtualMachineRequestContent) GetMigratedfrom() string`

GetMigratedfrom returns the Migratedfrom field if non-nil, zero value otherwise.

### GetMigratedfromOk

`func (o *StartVirtualMachineRequestContent) GetMigratedfromOk() (*string, bool)`

GetMigratedfromOk returns a tuple with the Migratedfrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedfrom

`func (o *StartVirtualMachineRequestContent) SetMigratedfrom(v string)`

SetMigratedfrom sets Migratedfrom field to given value.

### HasMigratedfrom

`func (o *StartVirtualMachineRequestContent) HasMigratedfrom() bool`

HasMigratedfrom returns a boolean if a field has been set.

### GetMigrationNetwork

`func (o *StartVirtualMachineRequestContent) GetMigrationNetwork() string`

GetMigrationNetwork returns the MigrationNetwork field if non-nil, zero value otherwise.

### GetMigrationNetworkOk

`func (o *StartVirtualMachineRequestContent) GetMigrationNetworkOk() (*string, bool)`

GetMigrationNetworkOk returns a tuple with the MigrationNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationNetwork

`func (o *StartVirtualMachineRequestContent) SetMigrationNetwork(v string)`

SetMigrationNetwork sets MigrationNetwork field to given value.

### HasMigrationNetwork

`func (o *StartVirtualMachineRequestContent) HasMigrationNetwork() bool`

HasMigrationNetwork returns a boolean if a field has been set.

### GetMigrationType

`func (o *StartVirtualMachineRequestContent) GetMigrationType() VirtualMachineMigrationType`

GetMigrationType returns the MigrationType field if non-nil, zero value otherwise.

### GetMigrationTypeOk

`func (o *StartVirtualMachineRequestContent) GetMigrationTypeOk() (*VirtualMachineMigrationType, bool)`

GetMigrationTypeOk returns a tuple with the MigrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationType

`func (o *StartVirtualMachineRequestContent) SetMigrationType(v VirtualMachineMigrationType)`

SetMigrationType sets MigrationType field to given value.

### HasMigrationType

`func (o *StartVirtualMachineRequestContent) HasMigrationType() bool`

HasMigrationType returns a boolean if a field has been set.

### GetStateuri

`func (o *StartVirtualMachineRequestContent) GetStateuri() string`

GetStateuri returns the Stateuri field if non-nil, zero value otherwise.

### GetStateuriOk

`func (o *StartVirtualMachineRequestContent) GetStateuriOk() (*string, bool)`

GetStateuriOk returns a tuple with the Stateuri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateuri

`func (o *StartVirtualMachineRequestContent) SetStateuri(v string)`

SetStateuri sets Stateuri field to given value.

### HasStateuri

`func (o *StartVirtualMachineRequestContent) HasStateuri() bool`

HasStateuri returns a boolean if a field has been set.

### GetTargetstorage

`func (o *StartVirtualMachineRequestContent) GetTargetstorage() string`

GetTargetstorage returns the Targetstorage field if non-nil, zero value otherwise.

### GetTargetstorageOk

`func (o *StartVirtualMachineRequestContent) GetTargetstorageOk() (*string, bool)`

GetTargetstorageOk returns a tuple with the Targetstorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetstorage

`func (o *StartVirtualMachineRequestContent) SetTargetstorage(v string)`

SetTargetstorage sets Targetstorage field to given value.

### HasTargetstorage

`func (o *StartVirtualMachineRequestContent) HasTargetstorage() bool`

HasTargetstorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


