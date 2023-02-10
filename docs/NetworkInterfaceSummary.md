# NetworkInterfaceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iface** | **string** |  | 
**Type** | [**NetworkInterfaceType**](NetworkInterfaceType.md) |  | 
**Method** | Pointer to [**NetworkInterfaceMethod**](NetworkInterfaceMethod.md) |  | [optional] 
**Method6** | Pointer to [**NetworkInterfaceMethod**](NetworkInterfaceMethod.md) |  | [optional] 
**Priority** | Pointer to **float32** |  | [optional] 
**Families** | Pointer to **[]string** |  | [optional] 
**BondPrimary** | Pointer to **string** |  | [optional] 
**BondMode** | Pointer to [**NetworkInterfaceBondMode**](NetworkInterfaceBondMode.md) |  | [optional] 
**BondXmitHashPolicy** | Pointer to [**NetworkInterfaceBondHashPolicy**](NetworkInterfaceBondHashPolicy.md) |  | [optional] 
**BondMiimon** | Pointer to **string** |  | [optional] 
**Slaves** | Pointer to **string** |  | [optional] 
**Autostart** | Pointer to **float32** |  | [optional] 
**Active** | Pointer to **float32** |  | [optional] 
**Exists** | Pointer to **float32** |  | [optional] 
**BridgeVids** | Pointer to **string** |  | [optional] 
**BridgePorts** | Pointer to **string** |  | [optional] 
**BridgeFd** | Pointer to **string** |  | [optional] 
**BridgeStp** | Pointer to **string** |  | [optional] 
**BridgeVlanAware** | Pointer to **float32** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Gateway6** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Netmask6** | Pointer to **string** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**Cidr6** | Pointer to **string** |  | [optional] 
**Address6** | Pointer to **string** |  | [optional] 

## Methods

### NewNetworkInterfaceSummary

`func NewNetworkInterfaceSummary(iface string, type_ NetworkInterfaceType, ) *NetworkInterfaceSummary`

NewNetworkInterfaceSummary instantiates a new NetworkInterfaceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceSummaryWithDefaults

`func NewNetworkInterfaceSummaryWithDefaults() *NetworkInterfaceSummary`

NewNetworkInterfaceSummaryWithDefaults instantiates a new NetworkInterfaceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIface

`func (o *NetworkInterfaceSummary) GetIface() string`

GetIface returns the Iface field if non-nil, zero value otherwise.

### GetIfaceOk

`func (o *NetworkInterfaceSummary) GetIfaceOk() (*string, bool)`

GetIfaceOk returns a tuple with the Iface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIface

`func (o *NetworkInterfaceSummary) SetIface(v string)`

SetIface sets Iface field to given value.


### GetType

`func (o *NetworkInterfaceSummary) GetType() NetworkInterfaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkInterfaceSummary) GetTypeOk() (*NetworkInterfaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkInterfaceSummary) SetType(v NetworkInterfaceType)`

SetType sets Type field to given value.


### GetMethod

`func (o *NetworkInterfaceSummary) GetMethod() NetworkInterfaceMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *NetworkInterfaceSummary) GetMethodOk() (*NetworkInterfaceMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *NetworkInterfaceSummary) SetMethod(v NetworkInterfaceMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *NetworkInterfaceSummary) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMethod6

`func (o *NetworkInterfaceSummary) GetMethod6() NetworkInterfaceMethod`

GetMethod6 returns the Method6 field if non-nil, zero value otherwise.

### GetMethod6Ok

`func (o *NetworkInterfaceSummary) GetMethod6Ok() (*NetworkInterfaceMethod, bool)`

GetMethod6Ok returns a tuple with the Method6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod6

`func (o *NetworkInterfaceSummary) SetMethod6(v NetworkInterfaceMethod)`

SetMethod6 sets Method6 field to given value.

### HasMethod6

`func (o *NetworkInterfaceSummary) HasMethod6() bool`

HasMethod6 returns a boolean if a field has been set.

### GetPriority

`func (o *NetworkInterfaceSummary) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NetworkInterfaceSummary) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NetworkInterfaceSummary) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NetworkInterfaceSummary) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetFamilies

`func (o *NetworkInterfaceSummary) GetFamilies() []string`

GetFamilies returns the Families field if non-nil, zero value otherwise.

### GetFamiliesOk

`func (o *NetworkInterfaceSummary) GetFamiliesOk() (*[]string, bool)`

GetFamiliesOk returns a tuple with the Families field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilies

`func (o *NetworkInterfaceSummary) SetFamilies(v []string)`

SetFamilies sets Families field to given value.

### HasFamilies

`func (o *NetworkInterfaceSummary) HasFamilies() bool`

HasFamilies returns a boolean if a field has been set.

### GetBondPrimary

`func (o *NetworkInterfaceSummary) GetBondPrimary() string`

GetBondPrimary returns the BondPrimary field if non-nil, zero value otherwise.

### GetBondPrimaryOk

`func (o *NetworkInterfaceSummary) GetBondPrimaryOk() (*string, bool)`

GetBondPrimaryOk returns a tuple with the BondPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondPrimary

`func (o *NetworkInterfaceSummary) SetBondPrimary(v string)`

SetBondPrimary sets BondPrimary field to given value.

### HasBondPrimary

`func (o *NetworkInterfaceSummary) HasBondPrimary() bool`

HasBondPrimary returns a boolean if a field has been set.

### GetBondMode

`func (o *NetworkInterfaceSummary) GetBondMode() NetworkInterfaceBondMode`

GetBondMode returns the BondMode field if non-nil, zero value otherwise.

### GetBondModeOk

`func (o *NetworkInterfaceSummary) GetBondModeOk() (*NetworkInterfaceBondMode, bool)`

GetBondModeOk returns a tuple with the BondMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondMode

`func (o *NetworkInterfaceSummary) SetBondMode(v NetworkInterfaceBondMode)`

SetBondMode sets BondMode field to given value.

### HasBondMode

`func (o *NetworkInterfaceSummary) HasBondMode() bool`

HasBondMode returns a boolean if a field has been set.

### GetBondXmitHashPolicy

`func (o *NetworkInterfaceSummary) GetBondXmitHashPolicy() NetworkInterfaceBondHashPolicy`

GetBondXmitHashPolicy returns the BondXmitHashPolicy field if non-nil, zero value otherwise.

### GetBondXmitHashPolicyOk

`func (o *NetworkInterfaceSummary) GetBondXmitHashPolicyOk() (*NetworkInterfaceBondHashPolicy, bool)`

GetBondXmitHashPolicyOk returns a tuple with the BondXmitHashPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondXmitHashPolicy

`func (o *NetworkInterfaceSummary) SetBondXmitHashPolicy(v NetworkInterfaceBondHashPolicy)`

SetBondXmitHashPolicy sets BondXmitHashPolicy field to given value.

### HasBondXmitHashPolicy

`func (o *NetworkInterfaceSummary) HasBondXmitHashPolicy() bool`

HasBondXmitHashPolicy returns a boolean if a field has been set.

### GetBondMiimon

`func (o *NetworkInterfaceSummary) GetBondMiimon() string`

GetBondMiimon returns the BondMiimon field if non-nil, zero value otherwise.

### GetBondMiimonOk

`func (o *NetworkInterfaceSummary) GetBondMiimonOk() (*string, bool)`

GetBondMiimonOk returns a tuple with the BondMiimon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondMiimon

`func (o *NetworkInterfaceSummary) SetBondMiimon(v string)`

SetBondMiimon sets BondMiimon field to given value.

### HasBondMiimon

`func (o *NetworkInterfaceSummary) HasBondMiimon() bool`

HasBondMiimon returns a boolean if a field has been set.

### GetSlaves

`func (o *NetworkInterfaceSummary) GetSlaves() string`

GetSlaves returns the Slaves field if non-nil, zero value otherwise.

### GetSlavesOk

`func (o *NetworkInterfaceSummary) GetSlavesOk() (*string, bool)`

GetSlavesOk returns a tuple with the Slaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaves

`func (o *NetworkInterfaceSummary) SetSlaves(v string)`

SetSlaves sets Slaves field to given value.

### HasSlaves

`func (o *NetworkInterfaceSummary) HasSlaves() bool`

HasSlaves returns a boolean if a field has been set.

### GetAutostart

`func (o *NetworkInterfaceSummary) GetAutostart() float32`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *NetworkInterfaceSummary) GetAutostartOk() (*float32, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *NetworkInterfaceSummary) SetAutostart(v float32)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *NetworkInterfaceSummary) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetActive

`func (o *NetworkInterfaceSummary) GetActive() float32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NetworkInterfaceSummary) GetActiveOk() (*float32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NetworkInterfaceSummary) SetActive(v float32)`

SetActive sets Active field to given value.

### HasActive

`func (o *NetworkInterfaceSummary) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetExists

`func (o *NetworkInterfaceSummary) GetExists() float32`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *NetworkInterfaceSummary) GetExistsOk() (*float32, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *NetworkInterfaceSummary) SetExists(v float32)`

SetExists sets Exists field to given value.

### HasExists

`func (o *NetworkInterfaceSummary) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetBridgeVids

`func (o *NetworkInterfaceSummary) GetBridgeVids() string`

GetBridgeVids returns the BridgeVids field if non-nil, zero value otherwise.

### GetBridgeVidsOk

`func (o *NetworkInterfaceSummary) GetBridgeVidsOk() (*string, bool)`

GetBridgeVidsOk returns a tuple with the BridgeVids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVids

`func (o *NetworkInterfaceSummary) SetBridgeVids(v string)`

SetBridgeVids sets BridgeVids field to given value.

### HasBridgeVids

`func (o *NetworkInterfaceSummary) HasBridgeVids() bool`

HasBridgeVids returns a boolean if a field has been set.

### GetBridgePorts

`func (o *NetworkInterfaceSummary) GetBridgePorts() string`

GetBridgePorts returns the BridgePorts field if non-nil, zero value otherwise.

### GetBridgePortsOk

`func (o *NetworkInterfaceSummary) GetBridgePortsOk() (*string, bool)`

GetBridgePortsOk returns a tuple with the BridgePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgePorts

`func (o *NetworkInterfaceSummary) SetBridgePorts(v string)`

SetBridgePorts sets BridgePorts field to given value.

### HasBridgePorts

`func (o *NetworkInterfaceSummary) HasBridgePorts() bool`

HasBridgePorts returns a boolean if a field has been set.

### GetBridgeFd

`func (o *NetworkInterfaceSummary) GetBridgeFd() string`

GetBridgeFd returns the BridgeFd field if non-nil, zero value otherwise.

### GetBridgeFdOk

`func (o *NetworkInterfaceSummary) GetBridgeFdOk() (*string, bool)`

GetBridgeFdOk returns a tuple with the BridgeFd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeFd

`func (o *NetworkInterfaceSummary) SetBridgeFd(v string)`

SetBridgeFd sets BridgeFd field to given value.

### HasBridgeFd

`func (o *NetworkInterfaceSummary) HasBridgeFd() bool`

HasBridgeFd returns a boolean if a field has been set.

### GetBridgeStp

`func (o *NetworkInterfaceSummary) GetBridgeStp() string`

GetBridgeStp returns the BridgeStp field if non-nil, zero value otherwise.

### GetBridgeStpOk

`func (o *NetworkInterfaceSummary) GetBridgeStpOk() (*string, bool)`

GetBridgeStpOk returns a tuple with the BridgeStp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeStp

`func (o *NetworkInterfaceSummary) SetBridgeStp(v string)`

SetBridgeStp sets BridgeStp field to given value.

### HasBridgeStp

`func (o *NetworkInterfaceSummary) HasBridgeStp() bool`

HasBridgeStp returns a boolean if a field has been set.

### GetBridgeVlanAware

`func (o *NetworkInterfaceSummary) GetBridgeVlanAware() float32`

GetBridgeVlanAware returns the BridgeVlanAware field if non-nil, zero value otherwise.

### GetBridgeVlanAwareOk

`func (o *NetworkInterfaceSummary) GetBridgeVlanAwareOk() (*float32, bool)`

GetBridgeVlanAwareOk returns a tuple with the BridgeVlanAware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVlanAware

`func (o *NetworkInterfaceSummary) SetBridgeVlanAware(v float32)`

SetBridgeVlanAware sets BridgeVlanAware field to given value.

### HasBridgeVlanAware

`func (o *NetworkInterfaceSummary) HasBridgeVlanAware() bool`

HasBridgeVlanAware returns a boolean if a field has been set.

### GetAddress

`func (o *NetworkInterfaceSummary) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NetworkInterfaceSummary) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NetworkInterfaceSummary) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NetworkInterfaceSummary) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *NetworkInterfaceSummary) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworkInterfaceSummary) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworkInterfaceSummary) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NetworkInterfaceSummary) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGateway6

`func (o *NetworkInterfaceSummary) GetGateway6() string`

GetGateway6 returns the Gateway6 field if non-nil, zero value otherwise.

### GetGateway6Ok

`func (o *NetworkInterfaceSummary) GetGateway6Ok() (*string, bool)`

GetGateway6Ok returns a tuple with the Gateway6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway6

`func (o *NetworkInterfaceSummary) SetGateway6(v string)`

SetGateway6 sets Gateway6 field to given value.

### HasGateway6

`func (o *NetworkInterfaceSummary) HasGateway6() bool`

HasGateway6 returns a boolean if a field has been set.

### GetNetmask

`func (o *NetworkInterfaceSummary) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *NetworkInterfaceSummary) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *NetworkInterfaceSummary) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *NetworkInterfaceSummary) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetmask6

`func (o *NetworkInterfaceSummary) GetNetmask6() string`

GetNetmask6 returns the Netmask6 field if non-nil, zero value otherwise.

### GetNetmask6Ok

`func (o *NetworkInterfaceSummary) GetNetmask6Ok() (*string, bool)`

GetNetmask6Ok returns a tuple with the Netmask6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask6

`func (o *NetworkInterfaceSummary) SetNetmask6(v string)`

SetNetmask6 sets Netmask6 field to given value.

### HasNetmask6

`func (o *NetworkInterfaceSummary) HasNetmask6() bool`

HasNetmask6 returns a boolean if a field has been set.

### GetCidr

`func (o *NetworkInterfaceSummary) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *NetworkInterfaceSummary) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *NetworkInterfaceSummary) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *NetworkInterfaceSummary) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCidr6

`func (o *NetworkInterfaceSummary) GetCidr6() string`

GetCidr6 returns the Cidr6 field if non-nil, zero value otherwise.

### GetCidr6Ok

`func (o *NetworkInterfaceSummary) GetCidr6Ok() (*string, bool)`

GetCidr6Ok returns a tuple with the Cidr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr6

`func (o *NetworkInterfaceSummary) SetCidr6(v string)`

SetCidr6 sets Cidr6 field to given value.

### HasCidr6

`func (o *NetworkInterfaceSummary) HasCidr6() bool`

HasCidr6 returns a boolean if a field has been set.

### GetAddress6

`func (o *NetworkInterfaceSummary) GetAddress6() string`

GetAddress6 returns the Address6 field if non-nil, zero value otherwise.

### GetAddress6Ok

`func (o *NetworkInterfaceSummary) GetAddress6Ok() (*string, bool)`

GetAddress6Ok returns a tuple with the Address6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6

`func (o *NetworkInterfaceSummary) SetAddress6(v string)`

SetAddress6 sets Address6 field to given value.

### HasAddress6

`func (o *NetworkInterfaceSummary) HasAddress6() bool`

HasAddress6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


