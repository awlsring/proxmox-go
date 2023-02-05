# DownloadFromUrlToStorageRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**UploadContentType**](UploadContentType.md) |  | 
**Filename** | **string** | The name of the file to create. | 
**Url** | **string** | The URL to download from. | 
**Checksum** | Pointer to **string** | The expceted checksum of the file. | [optional] 
**ChecksumAlgorithm** | Pointer to [**ChecksumAlgorithm**](ChecksumAlgorithm.md) |  | [optional] 
**Tmpfilename** | Pointer to **string** | The source filename. | [optional] 

## Methods

### NewDownloadFromUrlToStorageRequestContent

`func NewDownloadFromUrlToStorageRequestContent(content UploadContentType, filename string, url string, ) *DownloadFromUrlToStorageRequestContent`

NewDownloadFromUrlToStorageRequestContent instantiates a new DownloadFromUrlToStorageRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadFromUrlToStorageRequestContentWithDefaults

`func NewDownloadFromUrlToStorageRequestContentWithDefaults() *DownloadFromUrlToStorageRequestContent`

NewDownloadFromUrlToStorageRequestContentWithDefaults instantiates a new DownloadFromUrlToStorageRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *DownloadFromUrlToStorageRequestContent) GetContent() UploadContentType`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DownloadFromUrlToStorageRequestContent) GetContentOk() (*UploadContentType, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DownloadFromUrlToStorageRequestContent) SetContent(v UploadContentType)`

SetContent sets Content field to given value.


### GetFilename

`func (o *DownloadFromUrlToStorageRequestContent) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *DownloadFromUrlToStorageRequestContent) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *DownloadFromUrlToStorageRequestContent) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetUrl

`func (o *DownloadFromUrlToStorageRequestContent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DownloadFromUrlToStorageRequestContent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DownloadFromUrlToStorageRequestContent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetChecksum

`func (o *DownloadFromUrlToStorageRequestContent) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *DownloadFromUrlToStorageRequestContent) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *DownloadFromUrlToStorageRequestContent) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *DownloadFromUrlToStorageRequestContent) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetChecksumAlgorithm

`func (o *DownloadFromUrlToStorageRequestContent) GetChecksumAlgorithm() ChecksumAlgorithm`

GetChecksumAlgorithm returns the ChecksumAlgorithm field if non-nil, zero value otherwise.

### GetChecksumAlgorithmOk

`func (o *DownloadFromUrlToStorageRequestContent) GetChecksumAlgorithmOk() (*ChecksumAlgorithm, bool)`

GetChecksumAlgorithmOk returns a tuple with the ChecksumAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksumAlgorithm

`func (o *DownloadFromUrlToStorageRequestContent) SetChecksumAlgorithm(v ChecksumAlgorithm)`

SetChecksumAlgorithm sets ChecksumAlgorithm field to given value.

### HasChecksumAlgorithm

`func (o *DownloadFromUrlToStorageRequestContent) HasChecksumAlgorithm() bool`

HasChecksumAlgorithm returns a boolean if a field has been set.

### GetTmpfilename

`func (o *DownloadFromUrlToStorageRequestContent) GetTmpfilename() string`

GetTmpfilename returns the Tmpfilename field if non-nil, zero value otherwise.

### GetTmpfilenameOk

`func (o *DownloadFromUrlToStorageRequestContent) GetTmpfilenameOk() (*string, bool)`

GetTmpfilenameOk returns a tuple with the Tmpfilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmpfilename

`func (o *DownloadFromUrlToStorageRequestContent) SetTmpfilename(v string)`

SetTmpfilename sets Tmpfilename field to given value.

### HasTmpfilename

`func (o *DownloadFromUrlToStorageRequestContent) HasTmpfilename() bool`

HasTmpfilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


