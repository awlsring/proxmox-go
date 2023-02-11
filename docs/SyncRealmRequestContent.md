# SyncRealmRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**EnableNew** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Full** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Purge** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**RemoveVanished** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to [**RealmSyncScope**](RealmSyncScope.md) |  | [optional] 

## Methods

### NewSyncRealmRequestContent

`func NewSyncRealmRequestContent() *SyncRealmRequestContent`

NewSyncRealmRequestContent instantiates a new SyncRealmRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncRealmRequestContentWithDefaults

`func NewSyncRealmRequestContentWithDefaults() *SyncRealmRequestContent`

NewSyncRealmRequestContentWithDefaults instantiates a new SyncRealmRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *SyncRealmRequestContent) GetDryRun() float32`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *SyncRealmRequestContent) GetDryRunOk() (*float32, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *SyncRealmRequestContent) SetDryRun(v float32)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *SyncRealmRequestContent) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetEnableNew

`func (o *SyncRealmRequestContent) GetEnableNew() float32`

GetEnableNew returns the EnableNew field if non-nil, zero value otherwise.

### GetEnableNewOk

`func (o *SyncRealmRequestContent) GetEnableNewOk() (*float32, bool)`

GetEnableNewOk returns a tuple with the EnableNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNew

`func (o *SyncRealmRequestContent) SetEnableNew(v float32)`

SetEnableNew sets EnableNew field to given value.

### HasEnableNew

`func (o *SyncRealmRequestContent) HasEnableNew() bool`

HasEnableNew returns a boolean if a field has been set.

### GetFull

`func (o *SyncRealmRequestContent) GetFull() float32`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *SyncRealmRequestContent) GetFullOk() (*float32, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFull

`func (o *SyncRealmRequestContent) SetFull(v float32)`

SetFull sets Full field to given value.

### HasFull

`func (o *SyncRealmRequestContent) HasFull() bool`

HasFull returns a boolean if a field has been set.

### GetPurge

`func (o *SyncRealmRequestContent) GetPurge() float32`

GetPurge returns the Purge field if non-nil, zero value otherwise.

### GetPurgeOk

`func (o *SyncRealmRequestContent) GetPurgeOk() (*float32, bool)`

GetPurgeOk returns a tuple with the Purge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurge

`func (o *SyncRealmRequestContent) SetPurge(v float32)`

SetPurge sets Purge field to given value.

### HasPurge

`func (o *SyncRealmRequestContent) HasPurge() bool`

HasPurge returns a boolean if a field has been set.

### GetRemoveVanished

`func (o *SyncRealmRequestContent) GetRemoveVanished() string`

GetRemoveVanished returns the RemoveVanished field if non-nil, zero value otherwise.

### GetRemoveVanishedOk

`func (o *SyncRealmRequestContent) GetRemoveVanishedOk() (*string, bool)`

GetRemoveVanishedOk returns a tuple with the RemoveVanished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveVanished

`func (o *SyncRealmRequestContent) SetRemoveVanished(v string)`

SetRemoveVanished sets RemoveVanished field to given value.

### HasRemoveVanished

`func (o *SyncRealmRequestContent) HasRemoveVanished() bool`

HasRemoveVanished returns a boolean if a field has been set.

### GetScope

`func (o *SyncRealmRequestContent) GetScope() RealmSyncScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SyncRealmRequestContent) GetScopeOk() (*RealmSyncScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SyncRealmRequestContent) SetScope(v RealmSyncScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SyncRealmRequestContent) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


