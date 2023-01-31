# DiskInformationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to **float32** |  | [optional] 
**Bus** | Pointer to **float32** |  | [optional] 
**PciController** | Pointer to [**PciControllerSummary**](PciControllerSummary.md) |  | [optional] 
**Unit** | Pointer to **float32** |  | [optional] 
**Dev** | Pointer to **string** |  | [optional] 
**BusType** | Pointer to **string** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 

## Methods

### NewDiskInformationSummary

`func NewDiskInformationSummary() *DiskInformationSummary`

NewDiskInformationSummary instantiates a new DiskInformationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskInformationSummaryWithDefaults

`func NewDiskInformationSummaryWithDefaults() *DiskInformationSummary`

NewDiskInformationSummaryWithDefaults instantiates a new DiskInformationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *DiskInformationSummary) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DiskInformationSummary) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DiskInformationSummary) SetTarget(v float32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DiskInformationSummary) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetBus

`func (o *DiskInformationSummary) GetBus() float32`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *DiskInformationSummary) GetBusOk() (*float32, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *DiskInformationSummary) SetBus(v float32)`

SetBus sets Bus field to given value.

### HasBus

`func (o *DiskInformationSummary) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetPciController

`func (o *DiskInformationSummary) GetPciController() PciControllerSummary`

GetPciController returns the PciController field if non-nil, zero value otherwise.

### GetPciControllerOk

`func (o *DiskInformationSummary) GetPciControllerOk() (*PciControllerSummary, bool)`

GetPciControllerOk returns a tuple with the PciController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciController

`func (o *DiskInformationSummary) SetPciController(v PciControllerSummary)`

SetPciController sets PciController field to given value.

### HasPciController

`func (o *DiskInformationSummary) HasPciController() bool`

HasPciController returns a boolean if a field has been set.

### GetUnit

`func (o *DiskInformationSummary) GetUnit() float32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *DiskInformationSummary) GetUnitOk() (*float32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *DiskInformationSummary) SetUnit(v float32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *DiskInformationSummary) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetDev

`func (o *DiskInformationSummary) GetDev() string`

GetDev returns the Dev field if non-nil, zero value otherwise.

### GetDevOk

`func (o *DiskInformationSummary) GetDevOk() (*string, bool)`

GetDevOk returns a tuple with the Dev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev

`func (o *DiskInformationSummary) SetDev(v string)`

SetDev sets Dev field to given value.

### HasDev

`func (o *DiskInformationSummary) HasDev() bool`

HasDev returns a boolean if a field has been set.

### GetBusType

`func (o *DiskInformationSummary) GetBusType() string`

GetBusType returns the BusType field if non-nil, zero value otherwise.

### GetBusTypeOk

`func (o *DiskInformationSummary) GetBusTypeOk() (*string, bool)`

GetBusTypeOk returns a tuple with the BusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusType

`func (o *DiskInformationSummary) SetBusType(v string)`

SetBusType sets BusType field to given value.

### HasBusType

`func (o *DiskInformationSummary) HasBusType() bool`

HasBusType returns a boolean if a field has been set.

### GetSerial

`func (o *DiskInformationSummary) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *DiskInformationSummary) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *DiskInformationSummary) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *DiskInformationSummary) HasSerial() bool`

HasSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


