# CephMonitorSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Addr** | Pointer to **float32** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 

## Methods

### NewCephMonitorSummary

`func NewCephMonitorSummary(name string, ) *CephMonitorSummary`

NewCephMonitorSummary instantiates a new CephMonitorSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCephMonitorSummaryWithDefaults

`func NewCephMonitorSummaryWithDefaults() *CephMonitorSummary`

NewCephMonitorSummaryWithDefaults instantiates a new CephMonitorSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CephMonitorSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CephMonitorSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CephMonitorSummary) SetName(v string)`

SetName sets Name field to given value.


### GetAddr

`func (o *CephMonitorSummary) GetAddr() float32`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *CephMonitorSummary) GetAddrOk() (*float32, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *CephMonitorSummary) SetAddr(v float32)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *CephMonitorSummary) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetHost

`func (o *CephMonitorSummary) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CephMonitorSummary) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CephMonitorSummary) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CephMonitorSummary) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


