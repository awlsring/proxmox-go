# LVMThinSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lv** | [**[]LVMThinSummary**](LVMThinSummary.md) |  | 
**LvSize** | **float32** | The size of the lvm thin pool in bytes | 
**Used** | **float32** | The used size of the lvm thin pool in bytes | 
**MetadataSize** | **float32** | The size of the metadata lv in bytes | 
**MetadataUsed** | **float32** | The used size of the metadata lv in bytes | 
**Vg** | **string** | The associated volume group. | 

## Methods

### NewLVMThinSummary

`func NewLVMThinSummary(lv []LVMThinSummary, lvSize float32, used float32, metadataSize float32, metadataUsed float32, vg string, ) *LVMThinSummary`

NewLVMThinSummary instantiates a new LVMThinSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLVMThinSummaryWithDefaults

`func NewLVMThinSummaryWithDefaults() *LVMThinSummary`

NewLVMThinSummaryWithDefaults instantiates a new LVMThinSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLv

`func (o *LVMThinSummary) GetLv() []LVMThinSummary`

GetLv returns the Lv field if non-nil, zero value otherwise.

### GetLvOk

`func (o *LVMThinSummary) GetLvOk() (*[]LVMThinSummary, bool)`

GetLvOk returns a tuple with the Lv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLv

`func (o *LVMThinSummary) SetLv(v []LVMThinSummary)`

SetLv sets Lv field to given value.


### GetLvSize

`func (o *LVMThinSummary) GetLvSize() float32`

GetLvSize returns the LvSize field if non-nil, zero value otherwise.

### GetLvSizeOk

`func (o *LVMThinSummary) GetLvSizeOk() (*float32, bool)`

GetLvSizeOk returns a tuple with the LvSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvSize

`func (o *LVMThinSummary) SetLvSize(v float32)`

SetLvSize sets LvSize field to given value.


### GetUsed

`func (o *LVMThinSummary) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *LVMThinSummary) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *LVMThinSummary) SetUsed(v float32)`

SetUsed sets Used field to given value.


### GetMetadataSize

`func (o *LVMThinSummary) GetMetadataSize() float32`

GetMetadataSize returns the MetadataSize field if non-nil, zero value otherwise.

### GetMetadataSizeOk

`func (o *LVMThinSummary) GetMetadataSizeOk() (*float32, bool)`

GetMetadataSizeOk returns a tuple with the MetadataSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSize

`func (o *LVMThinSummary) SetMetadataSize(v float32)`

SetMetadataSize sets MetadataSize field to given value.


### GetMetadataUsed

`func (o *LVMThinSummary) GetMetadataUsed() float32`

GetMetadataUsed returns the MetadataUsed field if non-nil, zero value otherwise.

### GetMetadataUsedOk

`func (o *LVMThinSummary) GetMetadataUsedOk() (*float32, bool)`

GetMetadataUsedOk returns a tuple with the MetadataUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUsed

`func (o *LVMThinSummary) SetMetadataUsed(v float32)`

SetMetadataUsed sets MetadataUsed field to given value.


### GetVg

`func (o *LVMThinSummary) GetVg() string`

GetVg returns the Vg field if non-nil, zero value otherwise.

### GetVgOk

`func (o *LVMThinSummary) GetVgOk() (*string, bool)`

GetVgOk returns a tuple with the Vg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVg

`func (o *LVMThinSummary) SetVg(v string)`

SetVg sets Vg field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


