# CreateLVMThinRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | **string** | The device to create the lvm thinpool on. | 
**Name** | **string** | The storage identifier. | 
**AddStorage** | Pointer to **float32** | Configure storage using the lvm thinpool. Takes a boolean integer value (0 false, 1 true). | [optional] 

## Methods

### NewCreateLVMThinRequestContent

`func NewCreateLVMThinRequestContent(device string, name string, ) *CreateLVMThinRequestContent`

NewCreateLVMThinRequestContent instantiates a new CreateLVMThinRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLVMThinRequestContentWithDefaults

`func NewCreateLVMThinRequestContentWithDefaults() *CreateLVMThinRequestContent`

NewCreateLVMThinRequestContentWithDefaults instantiates a new CreateLVMThinRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *CreateLVMThinRequestContent) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CreateLVMThinRequestContent) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CreateLVMThinRequestContent) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetName

`func (o *CreateLVMThinRequestContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLVMThinRequestContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLVMThinRequestContent) SetName(v string)`

SetName sets Name field to given value.


### GetAddStorage

`func (o *CreateLVMThinRequestContent) GetAddStorage() float32`

GetAddStorage returns the AddStorage field if non-nil, zero value otherwise.

### GetAddStorageOk

`func (o *CreateLVMThinRequestContent) GetAddStorageOk() (*float32, bool)`

GetAddStorageOk returns a tuple with the AddStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddStorage

`func (o *CreateLVMThinRequestContent) SetAddStorage(v float32)`

SetAddStorage sets AddStorage field to given value.

### HasAddStorage

`func (o *CreateLVMThinRequestContent) HasAddStorage() bool`

HasAddStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


