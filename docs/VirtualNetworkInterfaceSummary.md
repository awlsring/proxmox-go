# VirtualNetworkInterfaceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Statistics** | Pointer to [**NetworkInterfaceStatisticsSummary**](NetworkInterfaceStatisticsSummary.md) |  | [optional] 
**IpAddresses** | Pointer to [**[]NetworkInterfaceIpAddressSummary**](NetworkInterfaceIpAddressSummary.md) |  | [optional] 
**HardwareAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewVirtualNetworkInterfaceSummary

`func NewVirtualNetworkInterfaceSummary() *VirtualNetworkInterfaceSummary`

NewVirtualNetworkInterfaceSummary instantiates a new VirtualNetworkInterfaceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualNetworkInterfaceSummaryWithDefaults

`func NewVirtualNetworkInterfaceSummaryWithDefaults() *VirtualNetworkInterfaceSummary`

NewVirtualNetworkInterfaceSummaryWithDefaults instantiates a new VirtualNetworkInterfaceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualNetworkInterfaceSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualNetworkInterfaceSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualNetworkInterfaceSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualNetworkInterfaceSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatistics

`func (o *VirtualNetworkInterfaceSummary) GetStatistics() NetworkInterfaceStatisticsSummary`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *VirtualNetworkInterfaceSummary) GetStatisticsOk() (*NetworkInterfaceStatisticsSummary, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *VirtualNetworkInterfaceSummary) SetStatistics(v NetworkInterfaceStatisticsSummary)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *VirtualNetworkInterfaceSummary) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetIpAddresses

`func (o *VirtualNetworkInterfaceSummary) GetIpAddresses() []NetworkInterfaceIpAddressSummary`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *VirtualNetworkInterfaceSummary) GetIpAddressesOk() (*[]NetworkInterfaceIpAddressSummary, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *VirtualNetworkInterfaceSummary) SetIpAddresses(v []NetworkInterfaceIpAddressSummary)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *VirtualNetworkInterfaceSummary) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetHardwareAddress

`func (o *VirtualNetworkInterfaceSummary) GetHardwareAddress() string`

GetHardwareAddress returns the HardwareAddress field if non-nil, zero value otherwise.

### GetHardwareAddressOk

`func (o *VirtualNetworkInterfaceSummary) GetHardwareAddressOk() (*string, bool)`

GetHardwareAddressOk returns a tuple with the HardwareAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareAddress

`func (o *VirtualNetworkInterfaceSummary) SetHardwareAddress(v string)`

SetHardwareAddress sets HardwareAddress field to given value.

### HasHardwareAddress

`func (o *VirtualNetworkInterfaceSummary) HasHardwareAddress() bool`

HasHardwareAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


