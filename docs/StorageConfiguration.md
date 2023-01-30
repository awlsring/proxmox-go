# StorageConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Storage** | **string** |  | 
**Type** | [**StorageType**](StorageType.md) |  | 

## Methods

### NewStorageConfiguration

`func NewStorageConfiguration(storage string, type_ StorageType, ) *StorageConfiguration`

NewStorageConfiguration instantiates a new StorageConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageConfigurationWithDefaults

`func NewStorageConfigurationWithDefaults() *StorageConfiguration`

NewStorageConfigurationWithDefaults instantiates a new StorageConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorage

`func (o *StorageConfiguration) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *StorageConfiguration) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *StorageConfiguration) SetStorage(v string)`

SetStorage sets Storage field to given value.


### GetType

`func (o *StorageConfiguration) GetType() StorageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageConfiguration) GetTypeOk() (*StorageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageConfiguration) SetType(v StorageType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


