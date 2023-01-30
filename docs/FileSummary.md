# FileSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Repositories** | Pointer to [**[]FileRepositorySummary**](FileRepositorySummary.md) |  | [optional] 
**Digest** | Pointer to **[]float32** |  | [optional] 
**FileType** | Pointer to **string** |  | [optional] 

## Methods

### NewFileSummary

`func NewFileSummary() *FileSummary`

NewFileSummary instantiates a new FileSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSummaryWithDefaults

`func NewFileSummaryWithDefaults() *FileSummary`

NewFileSummaryWithDefaults instantiates a new FileSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FileSummary) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileSummary) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileSummary) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileSummary) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetContent

`func (o *FileSummary) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FileSummary) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FileSummary) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *FileSummary) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRepositories

`func (o *FileSummary) GetRepositories() []FileRepositorySummary`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *FileSummary) GetRepositoriesOk() (*[]FileRepositorySummary, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *FileSummary) SetRepositories(v []FileRepositorySummary)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *FileSummary) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetDigest

`func (o *FileSummary) GetDigest() []float32`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *FileSummary) GetDigestOk() (*[]float32, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *FileSummary) SetDigest(v []float32)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *FileSummary) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetFileType

`func (o *FileSummary) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *FileSummary) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *FileSummary) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *FileSummary) HasFileType() bool`

HasFileType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


