# CloneVirtualMachineRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Newid** | **float32** | The id of the virtual machine as an integer | 
**Bwlimit** | Pointer to **float32** | Override deafult bandwidth limit in KiB/s | [optional] 
**Description** | Pointer to **string** | Set a description for the new VM | [optional] 
**Full** | Pointer to **bool** | Create a full clone with all copy of disks. | [optional] 
**Name** | Pointer to **string** | Set a name for the new VM | [optional] 
**Format** | Pointer to [**CloneVirtualMachineDiskFormat**](CloneVirtualMachineDiskFormat.md) |  | [optional] 
**Storage** | Pointer to **string** | Set the storage for the new VM. Only valid for full clones. | [optional] 
**Target** | Pointer to **string** | Set the target node for the new VM. Only valid if the original VM is on shared storage. | [optional] 
**Snapname** | Pointer to **string** | Set the snapshot name for the new VM. | [optional] 
**Pool** | Pointer to **string** | Add the VM to a resource pool. | [optional] 

## Methods

### NewCloneVirtualMachineRequestContent

`func NewCloneVirtualMachineRequestContent(newid float32, ) *CloneVirtualMachineRequestContent`

NewCloneVirtualMachineRequestContent instantiates a new CloneVirtualMachineRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneVirtualMachineRequestContentWithDefaults

`func NewCloneVirtualMachineRequestContentWithDefaults() *CloneVirtualMachineRequestContent`

NewCloneVirtualMachineRequestContentWithDefaults instantiates a new CloneVirtualMachineRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewid

`func (o *CloneVirtualMachineRequestContent) GetNewid() float32`

GetNewid returns the Newid field if non-nil, zero value otherwise.

### GetNewidOk

`func (o *CloneVirtualMachineRequestContent) GetNewidOk() (*float32, bool)`

GetNewidOk returns a tuple with the Newid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewid

`func (o *CloneVirtualMachineRequestContent) SetNewid(v float32)`

SetNewid sets Newid field to given value.


### GetBwlimit

`func (o *CloneVirtualMachineRequestContent) GetBwlimit() float32`

GetBwlimit returns the Bwlimit field if non-nil, zero value otherwise.

### GetBwlimitOk

`func (o *CloneVirtualMachineRequestContent) GetBwlimitOk() (*float32, bool)`

GetBwlimitOk returns a tuple with the Bwlimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwlimit

`func (o *CloneVirtualMachineRequestContent) SetBwlimit(v float32)`

SetBwlimit sets Bwlimit field to given value.

### HasBwlimit

`func (o *CloneVirtualMachineRequestContent) HasBwlimit() bool`

HasBwlimit returns a boolean if a field has been set.

### GetDescription

`func (o *CloneVirtualMachineRequestContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloneVirtualMachineRequestContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloneVirtualMachineRequestContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloneVirtualMachineRequestContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFull

`func (o *CloneVirtualMachineRequestContent) GetFull() bool`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *CloneVirtualMachineRequestContent) GetFullOk() (*bool, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFull

`func (o *CloneVirtualMachineRequestContent) SetFull(v bool)`

SetFull sets Full field to given value.

### HasFull

`func (o *CloneVirtualMachineRequestContent) HasFull() bool`

HasFull returns a boolean if a field has been set.

### GetName

`func (o *CloneVirtualMachineRequestContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloneVirtualMachineRequestContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloneVirtualMachineRequestContent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloneVirtualMachineRequestContent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormat

`func (o *CloneVirtualMachineRequestContent) GetFormat() CloneVirtualMachineDiskFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CloneVirtualMachineRequestContent) GetFormatOk() (*CloneVirtualMachineDiskFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CloneVirtualMachineRequestContent) SetFormat(v CloneVirtualMachineDiskFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CloneVirtualMachineRequestContent) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetStorage

`func (o *CloneVirtualMachineRequestContent) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CloneVirtualMachineRequestContent) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CloneVirtualMachineRequestContent) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CloneVirtualMachineRequestContent) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetTarget

`func (o *CloneVirtualMachineRequestContent) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloneVirtualMachineRequestContent) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloneVirtualMachineRequestContent) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloneVirtualMachineRequestContent) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetSnapname

`func (o *CloneVirtualMachineRequestContent) GetSnapname() string`

GetSnapname returns the Snapname field if non-nil, zero value otherwise.

### GetSnapnameOk

`func (o *CloneVirtualMachineRequestContent) GetSnapnameOk() (*string, bool)`

GetSnapnameOk returns a tuple with the Snapname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapname

`func (o *CloneVirtualMachineRequestContent) SetSnapname(v string)`

SetSnapname sets Snapname field to given value.

### HasSnapname

`func (o *CloneVirtualMachineRequestContent) HasSnapname() bool`

HasSnapname returns a boolean if a field has been set.

### GetPool

`func (o *CloneVirtualMachineRequestContent) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *CloneVirtualMachineRequestContent) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *CloneVirtualMachineRequestContent) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *CloneVirtualMachineRequestContent) HasPool() bool`

HasPool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


