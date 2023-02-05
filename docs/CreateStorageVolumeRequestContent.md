# CreateStorageVolumeRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | The filename of the new file | 
**Size** | **float32** | The new size in byte string | 
**Vmid** | **string** | The id of the virtual machine as a string | 
**Format** | Pointer to [**StorageVolumeType**](StorageVolumeType.md) |  | [optional] 

## Methods

### NewCreateStorageVolumeRequestContent

`func NewCreateStorageVolumeRequestContent(filename string, size float32, vmid string, ) *CreateStorageVolumeRequestContent`

NewCreateStorageVolumeRequestContent instantiates a new CreateStorageVolumeRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStorageVolumeRequestContentWithDefaults

`func NewCreateStorageVolumeRequestContentWithDefaults() *CreateStorageVolumeRequestContent`

NewCreateStorageVolumeRequestContentWithDefaults instantiates a new CreateStorageVolumeRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *CreateStorageVolumeRequestContent) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CreateStorageVolumeRequestContent) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CreateStorageVolumeRequestContent) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetSize

`func (o *CreateStorageVolumeRequestContent) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateStorageVolumeRequestContent) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateStorageVolumeRequestContent) SetSize(v float32)`

SetSize sets Size field to given value.


### GetVmid

`func (o *CreateStorageVolumeRequestContent) GetVmid() string`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *CreateStorageVolumeRequestContent) GetVmidOk() (*string, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *CreateStorageVolumeRequestContent) SetVmid(v string)`

SetVmid sets Vmid field to given value.


### GetFormat

`func (o *CreateStorageVolumeRequestContent) GetFormat() StorageVolumeType`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateStorageVolumeRequestContent) GetFormatOk() (*StorageVolumeType, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateStorageVolumeRequestContent) SetFormat(v StorageVolumeType)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateStorageVolumeRequestContent) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


