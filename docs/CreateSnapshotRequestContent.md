# CreateSnapshotRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapname** | **string** | The snapshot name. | 
**Description** | Pointer to **string** |  | [optional] 
**Vmstate** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 

## Methods

### NewCreateSnapshotRequestContent

`func NewCreateSnapshotRequestContent(snapname string, ) *CreateSnapshotRequestContent`

NewCreateSnapshotRequestContent instantiates a new CreateSnapshotRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotRequestContentWithDefaults

`func NewCreateSnapshotRequestContentWithDefaults() *CreateSnapshotRequestContent`

NewCreateSnapshotRequestContentWithDefaults instantiates a new CreateSnapshotRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapname

`func (o *CreateSnapshotRequestContent) GetSnapname() string`

GetSnapname returns the Snapname field if non-nil, zero value otherwise.

### GetSnapnameOk

`func (o *CreateSnapshotRequestContent) GetSnapnameOk() (*string, bool)`

GetSnapnameOk returns a tuple with the Snapname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapname

`func (o *CreateSnapshotRequestContent) SetSnapname(v string)`

SetSnapname sets Snapname field to given value.


### GetDescription

`func (o *CreateSnapshotRequestContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSnapshotRequestContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSnapshotRequestContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSnapshotRequestContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVmstate

`func (o *CreateSnapshotRequestContent) GetVmstate() float32`

GetVmstate returns the Vmstate field if non-nil, zero value otherwise.

### GetVmstateOk

`func (o *CreateSnapshotRequestContent) GetVmstateOk() (*float32, bool)`

GetVmstateOk returns a tuple with the Vmstate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstate

`func (o *CreateSnapshotRequestContent) SetVmstate(v float32)`

SetVmstate sets Vmstate field to given value.

### HasVmstate

`func (o *CreateSnapshotRequestContent) HasVmstate() bool`

HasVmstate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


