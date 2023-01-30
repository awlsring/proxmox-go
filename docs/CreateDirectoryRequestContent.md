# CreateDirectoryRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | **string** | The device to create the directory on. | 
**Name** | **string** | The storage identifier. | 
**AddStorage** | Pointer to **float32** | Configure storage using the directory. Takes a boolean integer value (0 false, 1 true). | [optional] 
**Filesystem** | Pointer to [**DirectoryFileSystem**](DirectoryFileSystem.md) |  | [optional] 

## Methods

### NewCreateDirectoryRequestContent

`func NewCreateDirectoryRequestContent(device string, name string, ) *CreateDirectoryRequestContent`

NewCreateDirectoryRequestContent instantiates a new CreateDirectoryRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDirectoryRequestContentWithDefaults

`func NewCreateDirectoryRequestContentWithDefaults() *CreateDirectoryRequestContent`

NewCreateDirectoryRequestContentWithDefaults instantiates a new CreateDirectoryRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *CreateDirectoryRequestContent) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CreateDirectoryRequestContent) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CreateDirectoryRequestContent) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetName

`func (o *CreateDirectoryRequestContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDirectoryRequestContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDirectoryRequestContent) SetName(v string)`

SetName sets Name field to given value.


### GetAddStorage

`func (o *CreateDirectoryRequestContent) GetAddStorage() float32`

GetAddStorage returns the AddStorage field if non-nil, zero value otherwise.

### GetAddStorageOk

`func (o *CreateDirectoryRequestContent) GetAddStorageOk() (*float32, bool)`

GetAddStorageOk returns a tuple with the AddStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddStorage

`func (o *CreateDirectoryRequestContent) SetAddStorage(v float32)`

SetAddStorage sets AddStorage field to given value.

### HasAddStorage

`func (o *CreateDirectoryRequestContent) HasAddStorage() bool`

HasAddStorage returns a boolean if a field has been set.

### GetFilesystem

`func (o *CreateDirectoryRequestContent) GetFilesystem() DirectoryFileSystem`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *CreateDirectoryRequestContent) GetFilesystemOk() (*DirectoryFileSystem, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *CreateDirectoryRequestContent) SetFilesystem(v DirectoryFileSystem)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *CreateDirectoryRequestContent) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


