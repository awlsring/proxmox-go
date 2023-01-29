# DiskSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devpath** | **string** |  | 
**Gpt** | **float32** |  | 
**Osdid** | **float32** |  | 
**Size** | **float32** | The size of the disk in bytes | 
**ByIdLink** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**Rpm** | Pointer to **float32** |  | [optional] 
**Wwn** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Wearout** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to [**DiskType**](DiskType.md) |  | [optional] 
**Mounted** | Pointer to **float32** |  | [optional] 

## Methods

### NewDiskSummary

`func NewDiskSummary(devpath string, gpt float32, osdid float32, size float32, ) *DiskSummary`

NewDiskSummary instantiates a new DiskSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskSummaryWithDefaults

`func NewDiskSummaryWithDefaults() *DiskSummary`

NewDiskSummaryWithDefaults instantiates a new DiskSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevpath

`func (o *DiskSummary) GetDevpath() string`

GetDevpath returns the Devpath field if non-nil, zero value otherwise.

### GetDevpathOk

`func (o *DiskSummary) GetDevpathOk() (*string, bool)`

GetDevpathOk returns a tuple with the Devpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevpath

`func (o *DiskSummary) SetDevpath(v string)`

SetDevpath sets Devpath field to given value.


### GetGpt

`func (o *DiskSummary) GetGpt() float32`

GetGpt returns the Gpt field if non-nil, zero value otherwise.

### GetGptOk

`func (o *DiskSummary) GetGptOk() (*float32, bool)`

GetGptOk returns a tuple with the Gpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpt

`func (o *DiskSummary) SetGpt(v float32)`

SetGpt sets Gpt field to given value.


### GetOsdid

`func (o *DiskSummary) GetOsdid() float32`

GetOsdid returns the Osdid field if non-nil, zero value otherwise.

### GetOsdidOk

`func (o *DiskSummary) GetOsdidOk() (*float32, bool)`

GetOsdidOk returns a tuple with the Osdid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsdid

`func (o *DiskSummary) SetOsdid(v float32)`

SetOsdid sets Osdid field to given value.


### GetSize

`func (o *DiskSummary) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DiskSummary) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DiskSummary) SetSize(v float32)`

SetSize sets Size field to given value.


### GetByIdLink

`func (o *DiskSummary) GetByIdLink() string`

GetByIdLink returns the ByIdLink field if non-nil, zero value otherwise.

### GetByIdLinkOk

`func (o *DiskSummary) GetByIdLinkOk() (*string, bool)`

GetByIdLinkOk returns a tuple with the ByIdLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByIdLink

`func (o *DiskSummary) SetByIdLink(v string)`

SetByIdLink sets ByIdLink field to given value.

### HasByIdLink

`func (o *DiskSummary) HasByIdLink() bool`

HasByIdLink returns a boolean if a field has been set.

### GetModel

`func (o *DiskSummary) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DiskSummary) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DiskSummary) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DiskSummary) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetParent

`func (o *DiskSummary) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *DiskSummary) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *DiskSummary) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *DiskSummary) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetHealth

`func (o *DiskSummary) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *DiskSummary) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *DiskSummary) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *DiskSummary) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetSerial

`func (o *DiskSummary) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *DiskSummary) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *DiskSummary) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *DiskSummary) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetRpm

`func (o *DiskSummary) GetRpm() float32`

GetRpm returns the Rpm field if non-nil, zero value otherwise.

### GetRpmOk

`func (o *DiskSummary) GetRpmOk() (*float32, bool)`

GetRpmOk returns a tuple with the Rpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpm

`func (o *DiskSummary) SetRpm(v float32)`

SetRpm sets Rpm field to given value.

### HasRpm

`func (o *DiskSummary) HasRpm() bool`

HasRpm returns a boolean if a field has been set.

### GetWwn

`func (o *DiskSummary) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *DiskSummary) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *DiskSummary) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *DiskSummary) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetVendor

`func (o *DiskSummary) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DiskSummary) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DiskSummary) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DiskSummary) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetWearout

`func (o *DiskSummary) GetWearout() float32`

GetWearout returns the Wearout field if non-nil, zero value otherwise.

### GetWearoutOk

`func (o *DiskSummary) GetWearoutOk() (*float32, bool)`

GetWearoutOk returns a tuple with the Wearout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWearout

`func (o *DiskSummary) SetWearout(v float32)`

SetWearout sets Wearout field to given value.

### HasWearout

`func (o *DiskSummary) HasWearout() bool`

HasWearout returns a boolean if a field has been set.

### GetType

`func (o *DiskSummary) GetType() DiskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiskSummary) GetTypeOk() (*DiskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiskSummary) SetType(v DiskType)`

SetType sets Type field to given value.

### HasType

`func (o *DiskSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMounted

`func (o *DiskSummary) GetMounted() float32`

GetMounted returns the Mounted field if non-nil, zero value otherwise.

### GetMountedOk

`func (o *DiskSummary) GetMountedOk() (*float32, bool)`

GetMountedOk returns a tuple with the Mounted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounted

`func (o *DiskSummary) SetMounted(v float32)`

SetMounted sets Mounted field to given value.

### HasMounted

`func (o *DiskSummary) HasMounted() bool`

HasMounted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


