# RepositoriesReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**StandardRepos** | [**[]RepositorySummary**](RepositorySummary.md) |  | 
**Errors** | **[]string** |  | 
**Files** | [**[]FileSummary**](FileSummary.md) |  | 
**Infos** | [**[]FileInfoSummary**](FileInfoSummary.md) |  | 

## Methods

### NewRepositoriesReport

`func NewRepositoriesReport(digest string, standardRepos []RepositorySummary, errors []string, files []FileSummary, infos []FileInfoSummary, ) *RepositoriesReport`

NewRepositoriesReport instantiates a new RepositoriesReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoriesReportWithDefaults

`func NewRepositoriesReportWithDefaults() *RepositoriesReport`

NewRepositoriesReportWithDefaults instantiates a new RepositoriesReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *RepositoriesReport) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *RepositoriesReport) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *RepositoriesReport) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetStandardRepos

`func (o *RepositoriesReport) GetStandardRepos() []RepositorySummary`

GetStandardRepos returns the StandardRepos field if non-nil, zero value otherwise.

### GetStandardReposOk

`func (o *RepositoriesReport) GetStandardReposOk() (*[]RepositorySummary, bool)`

GetStandardReposOk returns a tuple with the StandardRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardRepos

`func (o *RepositoriesReport) SetStandardRepos(v []RepositorySummary)`

SetStandardRepos sets StandardRepos field to given value.


### GetErrors

`func (o *RepositoriesReport) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RepositoriesReport) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RepositoriesReport) SetErrors(v []string)`

SetErrors sets Errors field to given value.


### GetFiles

`func (o *RepositoriesReport) GetFiles() []FileSummary`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *RepositoriesReport) GetFilesOk() (*[]FileSummary, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *RepositoriesReport) SetFiles(v []FileSummary)`

SetFiles sets Files field to given value.


### GetInfos

`func (o *RepositoriesReport) GetInfos() []FileInfoSummary`

GetInfos returns the Infos field if non-nil, zero value otherwise.

### GetInfosOk

`func (o *RepositoriesReport) GetInfosOk() (*[]FileInfoSummary, bool)`

GetInfosOk returns a tuple with the Infos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfos

`func (o *RepositoriesReport) SetInfos(v []FileInfoSummary)`

SetInfos sets Infos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


