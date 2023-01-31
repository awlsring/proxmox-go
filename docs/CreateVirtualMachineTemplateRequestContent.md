# CreateVirtualMachineTemplateRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disk** | Pointer to [**VirtualMachineDiskTarget**](VirtualMachineDiskTarget.md) |  | [optional] 

## Methods

### NewCreateVirtualMachineTemplateRequestContent

`func NewCreateVirtualMachineTemplateRequestContent() *CreateVirtualMachineTemplateRequestContent`

NewCreateVirtualMachineTemplateRequestContent instantiates a new CreateVirtualMachineTemplateRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVirtualMachineTemplateRequestContentWithDefaults

`func NewCreateVirtualMachineTemplateRequestContentWithDefaults() *CreateVirtualMachineTemplateRequestContent`

NewCreateVirtualMachineTemplateRequestContentWithDefaults instantiates a new CreateVirtualMachineTemplateRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisk

`func (o *CreateVirtualMachineTemplateRequestContent) GetDisk() VirtualMachineDiskTarget`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *CreateVirtualMachineTemplateRequestContent) GetDiskOk() (*VirtualMachineDiskTarget, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *CreateVirtualMachineTemplateRequestContent) SetDisk(v VirtualMachineDiskTarget)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *CreateVirtualMachineTemplateRequestContent) HasDisk() bool`

HasDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


