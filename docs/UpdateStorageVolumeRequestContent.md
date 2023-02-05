# UpdateStorageVolumeRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to **string** | New notes for the volume | [optional] 
**Protected** | **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | 

## Methods

### NewUpdateStorageVolumeRequestContent

`func NewUpdateStorageVolumeRequestContent(protected float32, ) *UpdateStorageVolumeRequestContent`

NewUpdateStorageVolumeRequestContent instantiates a new UpdateStorageVolumeRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageVolumeRequestContentWithDefaults

`func NewUpdateStorageVolumeRequestContentWithDefaults() *UpdateStorageVolumeRequestContent`

NewUpdateStorageVolumeRequestContentWithDefaults instantiates a new UpdateStorageVolumeRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *UpdateStorageVolumeRequestContent) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateStorageVolumeRequestContent) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateStorageVolumeRequestContent) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateStorageVolumeRequestContent) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetProtected

`func (o *UpdateStorageVolumeRequestContent) GetProtected() float32`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *UpdateStorageVolumeRequestContent) GetProtectedOk() (*float32, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *UpdateStorageVolumeRequestContent) SetProtected(v float32)`

SetProtected sets Protected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


