# PackageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Priority** | **string** |  | 
**OldVersion** | **string** |  | 
**Description** | **string** |  | 
**Arch** | **string** |  | 
**Package** | **string** |  | 
**Section** | **string** |  | 
**Version** | **string** |  | 
**Origin** | **string** |  | 
**CurrentState** | **string** |  | 
**RunningKernel** | Pointer to **string** |  | [optional] 

## Methods

### NewPackageSummary

`func NewPackageSummary(title string, priority string, oldVersion string, description string, arch string, package_ string, section string, version string, origin string, currentState string, ) *PackageSummary`

NewPackageSummary instantiates a new PackageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageSummaryWithDefaults

`func NewPackageSummaryWithDefaults() *PackageSummary`

NewPackageSummaryWithDefaults instantiates a new PackageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PackageSummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PackageSummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PackageSummary) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPriority

`func (o *PackageSummary) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PackageSummary) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PackageSummary) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetOldVersion

`func (o *PackageSummary) GetOldVersion() string`

GetOldVersion returns the OldVersion field if non-nil, zero value otherwise.

### GetOldVersionOk

`func (o *PackageSummary) GetOldVersionOk() (*string, bool)`

GetOldVersionOk returns a tuple with the OldVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldVersion

`func (o *PackageSummary) SetOldVersion(v string)`

SetOldVersion sets OldVersion field to given value.


### GetDescription

`func (o *PackageSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackageSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackageSummary) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetArch

`func (o *PackageSummary) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *PackageSummary) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *PackageSummary) SetArch(v string)`

SetArch sets Arch field to given value.


### GetPackage

`func (o *PackageSummary) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *PackageSummary) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *PackageSummary) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetSection

`func (o *PackageSummary) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *PackageSummary) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *PackageSummary) SetSection(v string)`

SetSection sets Section field to given value.


### GetVersion

`func (o *PackageSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PackageSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PackageSummary) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetOrigin

`func (o *PackageSummary) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *PackageSummary) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *PackageSummary) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetCurrentState

`func (o *PackageSummary) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *PackageSummary) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *PackageSummary) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.


### GetRunningKernel

`func (o *PackageSummary) GetRunningKernel() string`

GetRunningKernel returns the RunningKernel field if non-nil, zero value otherwise.

### GetRunningKernelOk

`func (o *PackageSummary) GetRunningKernelOk() (*string, bool)`

GetRunningKernelOk returns a tuple with the RunningKernel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningKernel

`func (o *PackageSummary) SetRunningKernel(v string)`

SetRunningKernel sets RunningKernel field to given value.

### HasRunningKernel

`func (o *PackageSummary) HasRunningKernel() bool`

HasRunningKernel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


