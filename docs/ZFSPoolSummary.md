# ZFSPoolSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alloc** | **float32** | The allocated space on the pool in bytes | 
**Dedup** | **float32** |  | 
**Frag** | **float32** |  | 
**Health** | **string** | The health of the pool | 
**Free** | **float32** | The free space on the pool in bytes | 
**Size** | **float32** | The total size of the pool in bytes | 
**Name** | **string** | Name of the pool | 

## Methods

### NewZFSPoolSummary

`func NewZFSPoolSummary(alloc float32, dedup float32, frag float32, health string, free float32, size float32, name string, ) *ZFSPoolSummary`

NewZFSPoolSummary instantiates a new ZFSPoolSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZFSPoolSummaryWithDefaults

`func NewZFSPoolSummaryWithDefaults() *ZFSPoolSummary`

NewZFSPoolSummaryWithDefaults instantiates a new ZFSPoolSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlloc

`func (o *ZFSPoolSummary) GetAlloc() float32`

GetAlloc returns the Alloc field if non-nil, zero value otherwise.

### GetAllocOk

`func (o *ZFSPoolSummary) GetAllocOk() (*float32, bool)`

GetAllocOk returns a tuple with the Alloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlloc

`func (o *ZFSPoolSummary) SetAlloc(v float32)`

SetAlloc sets Alloc field to given value.


### GetDedup

`func (o *ZFSPoolSummary) GetDedup() float32`

GetDedup returns the Dedup field if non-nil, zero value otherwise.

### GetDedupOk

`func (o *ZFSPoolSummary) GetDedupOk() (*float32, bool)`

GetDedupOk returns a tuple with the Dedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedup

`func (o *ZFSPoolSummary) SetDedup(v float32)`

SetDedup sets Dedup field to given value.


### GetFrag

`func (o *ZFSPoolSummary) GetFrag() float32`

GetFrag returns the Frag field if non-nil, zero value otherwise.

### GetFragOk

`func (o *ZFSPoolSummary) GetFragOk() (*float32, bool)`

GetFragOk returns a tuple with the Frag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrag

`func (o *ZFSPoolSummary) SetFrag(v float32)`

SetFrag sets Frag field to given value.


### GetHealth

`func (o *ZFSPoolSummary) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ZFSPoolSummary) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ZFSPoolSummary) SetHealth(v string)`

SetHealth sets Health field to given value.


### GetFree

`func (o *ZFSPoolSummary) GetFree() float32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *ZFSPoolSummary) GetFreeOk() (*float32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *ZFSPoolSummary) SetFree(v float32)`

SetFree sets Free field to given value.


### GetSize

`func (o *ZFSPoolSummary) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ZFSPoolSummary) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ZFSPoolSummary) SetSize(v float32)`

SetSize sets Size field to given value.


### GetName

`func (o *ZFSPoolSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZFSPoolSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZFSPoolSummary) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


