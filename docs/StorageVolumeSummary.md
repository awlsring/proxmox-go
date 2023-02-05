# StorageVolumeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** |  | 
**Size** | **float32** | The size of the storage content in bytes. | 
**Volid** | **string** |  | 
**Ctime** | Pointer to **float32** | The creation time of the storage content in seconds since the epoch. | [optional] 
**Encrypted** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**Protected** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Used** | Pointer to **float32** | The amount of storage content used in bytes. | [optional] 
**Vmid** | Pointer to **string** | The id of the virtual machine as a string | [optional] 

## Methods

### NewStorageVolumeSummary

`func NewStorageVolumeSummary(format string, size float32, volid string, ) *StorageVolumeSummary`

NewStorageVolumeSummary instantiates a new StorageVolumeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVolumeSummaryWithDefaults

`func NewStorageVolumeSummaryWithDefaults() *StorageVolumeSummary`

NewStorageVolumeSummaryWithDefaults instantiates a new StorageVolumeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *StorageVolumeSummary) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *StorageVolumeSummary) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *StorageVolumeSummary) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetSize

`func (o *StorageVolumeSummary) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageVolumeSummary) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageVolumeSummary) SetSize(v float32)`

SetSize sets Size field to given value.


### GetVolid

`func (o *StorageVolumeSummary) GetVolid() string`

GetVolid returns the Volid field if non-nil, zero value otherwise.

### GetVolidOk

`func (o *StorageVolumeSummary) GetVolidOk() (*string, bool)`

GetVolidOk returns a tuple with the Volid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolid

`func (o *StorageVolumeSummary) SetVolid(v string)`

SetVolid sets Volid field to given value.


### GetCtime

`func (o *StorageVolumeSummary) GetCtime() float32`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *StorageVolumeSummary) GetCtimeOk() (*float32, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *StorageVolumeSummary) SetCtime(v float32)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *StorageVolumeSummary) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetEncrypted

`func (o *StorageVolumeSummary) GetEncrypted() string`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *StorageVolumeSummary) GetEncryptedOk() (*string, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *StorageVolumeSummary) SetEncrypted(v string)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *StorageVolumeSummary) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetNotes

`func (o *StorageVolumeSummary) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *StorageVolumeSummary) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *StorageVolumeSummary) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *StorageVolumeSummary) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetParent

`func (o *StorageVolumeSummary) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StorageVolumeSummary) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StorageVolumeSummary) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StorageVolumeSummary) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtected

`func (o *StorageVolumeSummary) GetProtected() float32`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *StorageVolumeSummary) GetProtectedOk() (*float32, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *StorageVolumeSummary) SetProtected(v float32)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *StorageVolumeSummary) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetUsed

`func (o *StorageVolumeSummary) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *StorageVolumeSummary) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *StorageVolumeSummary) SetUsed(v float32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *StorageVolumeSummary) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetVmid

`func (o *StorageVolumeSummary) GetVmid() string`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *StorageVolumeSummary) GetVmidOk() (*string, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *StorageVolumeSummary) SetVmid(v string)`

SetVmid sets Vmid field to given value.

### HasVmid

`func (o *StorageVolumeSummary) HasVmid() bool`

HasVmid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


