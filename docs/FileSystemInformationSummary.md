# FileSystemInformationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disk** | Pointer to [**[]DiskInformationSummary**](DiskInformationSummary.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Mountpoint** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UsedBytes** | Pointer to **float32** |  | [optional] 
**TotalBytes** | Pointer to **float32** |  | [optional] 

## Methods

### NewFileSystemInformationSummary

`func NewFileSystemInformationSummary() *FileSystemInformationSummary`

NewFileSystemInformationSummary instantiates a new FileSystemInformationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSystemInformationSummaryWithDefaults

`func NewFileSystemInformationSummaryWithDefaults() *FileSystemInformationSummary`

NewFileSystemInformationSummaryWithDefaults instantiates a new FileSystemInformationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisk

`func (o *FileSystemInformationSummary) GetDisk() []DiskInformationSummary`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *FileSystemInformationSummary) GetDiskOk() (*[]DiskInformationSummary, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *FileSystemInformationSummary) SetDisk(v []DiskInformationSummary)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *FileSystemInformationSummary) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetName

`func (o *FileSystemInformationSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileSystemInformationSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileSystemInformationSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileSystemInformationSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMountpoint

`func (o *FileSystemInformationSummary) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *FileSystemInformationSummary) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *FileSystemInformationSummary) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *FileSystemInformationSummary) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetType

`func (o *FileSystemInformationSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileSystemInformationSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileSystemInformationSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FileSystemInformationSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsedBytes

`func (o *FileSystemInformationSummary) GetUsedBytes() float32`

GetUsedBytes returns the UsedBytes field if non-nil, zero value otherwise.

### GetUsedBytesOk

`func (o *FileSystemInformationSummary) GetUsedBytesOk() (*float32, bool)`

GetUsedBytesOk returns a tuple with the UsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBytes

`func (o *FileSystemInformationSummary) SetUsedBytes(v float32)`

SetUsedBytes sets UsedBytes field to given value.

### HasUsedBytes

`func (o *FileSystemInformationSummary) HasUsedBytes() bool`

HasUsedBytes returns a boolean if a field has been set.

### GetTotalBytes

`func (o *FileSystemInformationSummary) GetTotalBytes() float32`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *FileSystemInformationSummary) GetTotalBytesOk() (*float32, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *FileSystemInformationSummary) SetTotalBytes(v float32)`

SetTotalBytes sets TotalBytes field to given value.

### HasTotalBytes

`func (o *FileSystemInformationSummary) HasTotalBytes() bool`

HasTotalBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


