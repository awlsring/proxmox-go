# StorageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Storage** | **string** | The storage class name | 
**Content** | Pointer to **string** | Comma seperated list of content types in storage. Returned as a string | [optional] 
**Digest** | Pointer to **string** |  | [optional] 
**Thinpool** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**StorageType**](StorageType.md) |  | [optional] 
**Vgname** | Pointer to **string** | The volume group name | [optional] 
**Pool** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** | The storage path | [optional] 
**Mountpoint** | Pointer to **string** | The storage mountpoint | [optional] 
**Nodes** | Pointer to **string** | The nodes that have access to this storage | [optional] 

## Methods

### NewStorageSummary

`func NewStorageSummary(storage string, ) *StorageSummary`

NewStorageSummary instantiates a new StorageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSummaryWithDefaults

`func NewStorageSummaryWithDefaults() *StorageSummary`

NewStorageSummaryWithDefaults instantiates a new StorageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorage

`func (o *StorageSummary) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *StorageSummary) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *StorageSummary) SetStorage(v string)`

SetStorage sets Storage field to given value.


### GetContent

`func (o *StorageSummary) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *StorageSummary) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *StorageSummary) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *StorageSummary) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDigest

`func (o *StorageSummary) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *StorageSummary) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *StorageSummary) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *StorageSummary) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetThinpool

`func (o *StorageSummary) GetThinpool() string`

GetThinpool returns the Thinpool field if non-nil, zero value otherwise.

### GetThinpoolOk

`func (o *StorageSummary) GetThinpoolOk() (*string, bool)`

GetThinpoolOk returns a tuple with the Thinpool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinpool

`func (o *StorageSummary) SetThinpool(v string)`

SetThinpool sets Thinpool field to given value.

### HasThinpool

`func (o *StorageSummary) HasThinpool() bool`

HasThinpool returns a boolean if a field has been set.

### GetType

`func (o *StorageSummary) GetType() StorageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageSummary) GetTypeOk() (*StorageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageSummary) SetType(v StorageType)`

SetType sets Type field to given value.

### HasType

`func (o *StorageSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVgname

`func (o *StorageSummary) GetVgname() string`

GetVgname returns the Vgname field if non-nil, zero value otherwise.

### GetVgnameOk

`func (o *StorageSummary) GetVgnameOk() (*string, bool)`

GetVgnameOk returns a tuple with the Vgname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgname

`func (o *StorageSummary) SetVgname(v string)`

SetVgname sets Vgname field to given value.

### HasVgname

`func (o *StorageSummary) HasVgname() bool`

HasVgname returns a boolean if a field has been set.

### GetPool

`func (o *StorageSummary) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *StorageSummary) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *StorageSummary) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *StorageSummary) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPath

`func (o *StorageSummary) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StorageSummary) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StorageSummary) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StorageSummary) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMountpoint

`func (o *StorageSummary) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *StorageSummary) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *StorageSummary) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *StorageSummary) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetNodes

`func (o *StorageSummary) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *StorageSummary) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *StorageSummary) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *StorageSummary) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


