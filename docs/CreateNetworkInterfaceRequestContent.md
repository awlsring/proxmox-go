# CreateNetworkInterfaceRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iface** | **string** |  | 
**Type** | [**NetworkInterfaceType**](NetworkInterfaceType.md) |  | 
**Address** | Pointer to **string** |  | [optional] 
**Address6** | Pointer to **string** |  | [optional] 
**Autostart** | Pointer to **bool** |  | [optional] 
**BondPrimary** | Pointer to **string** |  | [optional] 
**BondMode** | Pointer to [**NetworkInterfaceBondMode**](NetworkInterfaceBondMode.md) |  | [optional] 
**BondXmitHashPolicy** | Pointer to [**NetworkInterfaceBondHashPolicy**](NetworkInterfaceBondHashPolicy.md) |  | [optional] 
**BridgePorts** | Pointer to **string** |  | [optional] 
**BridgeVlanAware** | Pointer to **bool** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**Cidr6** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Comments6** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Gateway6** | Pointer to **string** |  | [optional] 
**Mtu** | Pointer to **float32** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Netmask6** | Pointer to **float32** |  | [optional] 
**OvsBonds** | Pointer to **string** |  | [optional] 
**OvsOptions** | Pointer to **string** |  | [optional] 
**OvsPorts** | Pointer to **string** |  | [optional] 
**OvsTag** | Pointer to **float32** |  | [optional] 
**OvsBridge** | Pointer to **string** |  | [optional] 
**Slaves** | Pointer to **string** |  | [optional] 
**VlanId** | Pointer to **float32** |  | [optional] 
**VlanRawDevice** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateNetworkInterfaceRequestContent

`func NewCreateNetworkInterfaceRequestContent(iface string, type_ NetworkInterfaceType, ) *CreateNetworkInterfaceRequestContent`

NewCreateNetworkInterfaceRequestContent instantiates a new CreateNetworkInterfaceRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkInterfaceRequestContentWithDefaults

`func NewCreateNetworkInterfaceRequestContentWithDefaults() *CreateNetworkInterfaceRequestContent`

NewCreateNetworkInterfaceRequestContentWithDefaults instantiates a new CreateNetworkInterfaceRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIface

`func (o *CreateNetworkInterfaceRequestContent) GetIface() string`

GetIface returns the Iface field if non-nil, zero value otherwise.

### GetIfaceOk

`func (o *CreateNetworkInterfaceRequestContent) GetIfaceOk() (*string, bool)`

GetIfaceOk returns a tuple with the Iface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIface

`func (o *CreateNetworkInterfaceRequestContent) SetIface(v string)`

SetIface sets Iface field to given value.


### GetType

`func (o *CreateNetworkInterfaceRequestContent) GetType() NetworkInterfaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNetworkInterfaceRequestContent) GetTypeOk() (*NetworkInterfaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNetworkInterfaceRequestContent) SetType(v NetworkInterfaceType)`

SetType sets Type field to given value.


### GetAddress

`func (o *CreateNetworkInterfaceRequestContent) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateNetworkInterfaceRequestContent) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateNetworkInterfaceRequestContent) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CreateNetworkInterfaceRequestContent) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddress6

`func (o *CreateNetworkInterfaceRequestContent) GetAddress6() string`

GetAddress6 returns the Address6 field if non-nil, zero value otherwise.

### GetAddress6Ok

`func (o *CreateNetworkInterfaceRequestContent) GetAddress6Ok() (*string, bool)`

GetAddress6Ok returns a tuple with the Address6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6

`func (o *CreateNetworkInterfaceRequestContent) SetAddress6(v string)`

SetAddress6 sets Address6 field to given value.

### HasAddress6

`func (o *CreateNetworkInterfaceRequestContent) HasAddress6() bool`

HasAddress6 returns a boolean if a field has been set.

### GetAutostart

`func (o *CreateNetworkInterfaceRequestContent) GetAutostart() bool`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *CreateNetworkInterfaceRequestContent) GetAutostartOk() (*bool, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *CreateNetworkInterfaceRequestContent) SetAutostart(v bool)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *CreateNetworkInterfaceRequestContent) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetBondPrimary

`func (o *CreateNetworkInterfaceRequestContent) GetBondPrimary() string`

GetBondPrimary returns the BondPrimary field if non-nil, zero value otherwise.

### GetBondPrimaryOk

`func (o *CreateNetworkInterfaceRequestContent) GetBondPrimaryOk() (*string, bool)`

GetBondPrimaryOk returns a tuple with the BondPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondPrimary

`func (o *CreateNetworkInterfaceRequestContent) SetBondPrimary(v string)`

SetBondPrimary sets BondPrimary field to given value.

### HasBondPrimary

`func (o *CreateNetworkInterfaceRequestContent) HasBondPrimary() bool`

HasBondPrimary returns a boolean if a field has been set.

### GetBondMode

`func (o *CreateNetworkInterfaceRequestContent) GetBondMode() NetworkInterfaceBondMode`

GetBondMode returns the BondMode field if non-nil, zero value otherwise.

### GetBondModeOk

`func (o *CreateNetworkInterfaceRequestContent) GetBondModeOk() (*NetworkInterfaceBondMode, bool)`

GetBondModeOk returns a tuple with the BondMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondMode

`func (o *CreateNetworkInterfaceRequestContent) SetBondMode(v NetworkInterfaceBondMode)`

SetBondMode sets BondMode field to given value.

### HasBondMode

`func (o *CreateNetworkInterfaceRequestContent) HasBondMode() bool`

HasBondMode returns a boolean if a field has been set.

### GetBondXmitHashPolicy

`func (o *CreateNetworkInterfaceRequestContent) GetBondXmitHashPolicy() NetworkInterfaceBondHashPolicy`

GetBondXmitHashPolicy returns the BondXmitHashPolicy field if non-nil, zero value otherwise.

### GetBondXmitHashPolicyOk

`func (o *CreateNetworkInterfaceRequestContent) GetBondXmitHashPolicyOk() (*NetworkInterfaceBondHashPolicy, bool)`

GetBondXmitHashPolicyOk returns a tuple with the BondXmitHashPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondXmitHashPolicy

`func (o *CreateNetworkInterfaceRequestContent) SetBondXmitHashPolicy(v NetworkInterfaceBondHashPolicy)`

SetBondXmitHashPolicy sets BondXmitHashPolicy field to given value.

### HasBondXmitHashPolicy

`func (o *CreateNetworkInterfaceRequestContent) HasBondXmitHashPolicy() bool`

HasBondXmitHashPolicy returns a boolean if a field has been set.

### GetBridgePorts

`func (o *CreateNetworkInterfaceRequestContent) GetBridgePorts() string`

GetBridgePorts returns the BridgePorts field if non-nil, zero value otherwise.

### GetBridgePortsOk

`func (o *CreateNetworkInterfaceRequestContent) GetBridgePortsOk() (*string, bool)`

GetBridgePortsOk returns a tuple with the BridgePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgePorts

`func (o *CreateNetworkInterfaceRequestContent) SetBridgePorts(v string)`

SetBridgePorts sets BridgePorts field to given value.

### HasBridgePorts

`func (o *CreateNetworkInterfaceRequestContent) HasBridgePorts() bool`

HasBridgePorts returns a boolean if a field has been set.

### GetBridgeVlanAware

`func (o *CreateNetworkInterfaceRequestContent) GetBridgeVlanAware() bool`

GetBridgeVlanAware returns the BridgeVlanAware field if non-nil, zero value otherwise.

### GetBridgeVlanAwareOk

`func (o *CreateNetworkInterfaceRequestContent) GetBridgeVlanAwareOk() (*bool, bool)`

GetBridgeVlanAwareOk returns a tuple with the BridgeVlanAware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVlanAware

`func (o *CreateNetworkInterfaceRequestContent) SetBridgeVlanAware(v bool)`

SetBridgeVlanAware sets BridgeVlanAware field to given value.

### HasBridgeVlanAware

`func (o *CreateNetworkInterfaceRequestContent) HasBridgeVlanAware() bool`

HasBridgeVlanAware returns a boolean if a field has been set.

### GetCidr

`func (o *CreateNetworkInterfaceRequestContent) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateNetworkInterfaceRequestContent) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateNetworkInterfaceRequestContent) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CreateNetworkInterfaceRequestContent) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCidr6

`func (o *CreateNetworkInterfaceRequestContent) GetCidr6() string`

GetCidr6 returns the Cidr6 field if non-nil, zero value otherwise.

### GetCidr6Ok

`func (o *CreateNetworkInterfaceRequestContent) GetCidr6Ok() (*string, bool)`

GetCidr6Ok returns a tuple with the Cidr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr6

`func (o *CreateNetworkInterfaceRequestContent) SetCidr6(v string)`

SetCidr6 sets Cidr6 field to given value.

### HasCidr6

`func (o *CreateNetworkInterfaceRequestContent) HasCidr6() bool`

HasCidr6 returns a boolean if a field has been set.

### GetComments

`func (o *CreateNetworkInterfaceRequestContent) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CreateNetworkInterfaceRequestContent) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CreateNetworkInterfaceRequestContent) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CreateNetworkInterfaceRequestContent) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetComments6

`func (o *CreateNetworkInterfaceRequestContent) GetComments6() string`

GetComments6 returns the Comments6 field if non-nil, zero value otherwise.

### GetComments6Ok

`func (o *CreateNetworkInterfaceRequestContent) GetComments6Ok() (*string, bool)`

GetComments6Ok returns a tuple with the Comments6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments6

`func (o *CreateNetworkInterfaceRequestContent) SetComments6(v string)`

SetComments6 sets Comments6 field to given value.

### HasComments6

`func (o *CreateNetworkInterfaceRequestContent) HasComments6() bool`

HasComments6 returns a boolean if a field has been set.

### GetGateway

`func (o *CreateNetworkInterfaceRequestContent) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CreateNetworkInterfaceRequestContent) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CreateNetworkInterfaceRequestContent) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CreateNetworkInterfaceRequestContent) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGateway6

`func (o *CreateNetworkInterfaceRequestContent) GetGateway6() string`

GetGateway6 returns the Gateway6 field if non-nil, zero value otherwise.

### GetGateway6Ok

`func (o *CreateNetworkInterfaceRequestContent) GetGateway6Ok() (*string, bool)`

GetGateway6Ok returns a tuple with the Gateway6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway6

`func (o *CreateNetworkInterfaceRequestContent) SetGateway6(v string)`

SetGateway6 sets Gateway6 field to given value.

### HasGateway6

`func (o *CreateNetworkInterfaceRequestContent) HasGateway6() bool`

HasGateway6 returns a boolean if a field has been set.

### GetMtu

`func (o *CreateNetworkInterfaceRequestContent) GetMtu() float32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *CreateNetworkInterfaceRequestContent) GetMtuOk() (*float32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *CreateNetworkInterfaceRequestContent) SetMtu(v float32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *CreateNetworkInterfaceRequestContent) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNetmask

`func (o *CreateNetworkInterfaceRequestContent) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *CreateNetworkInterfaceRequestContent) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *CreateNetworkInterfaceRequestContent) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *CreateNetworkInterfaceRequestContent) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetmask6

`func (o *CreateNetworkInterfaceRequestContent) GetNetmask6() float32`

GetNetmask6 returns the Netmask6 field if non-nil, zero value otherwise.

### GetNetmask6Ok

`func (o *CreateNetworkInterfaceRequestContent) GetNetmask6Ok() (*float32, bool)`

GetNetmask6Ok returns a tuple with the Netmask6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask6

`func (o *CreateNetworkInterfaceRequestContent) SetNetmask6(v float32)`

SetNetmask6 sets Netmask6 field to given value.

### HasNetmask6

`func (o *CreateNetworkInterfaceRequestContent) HasNetmask6() bool`

HasNetmask6 returns a boolean if a field has been set.

### GetOvsBonds

`func (o *CreateNetworkInterfaceRequestContent) GetOvsBonds() string`

GetOvsBonds returns the OvsBonds field if non-nil, zero value otherwise.

### GetOvsBondsOk

`func (o *CreateNetworkInterfaceRequestContent) GetOvsBondsOk() (*string, bool)`

GetOvsBondsOk returns a tuple with the OvsBonds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvsBonds

`func (o *CreateNetworkInterfaceRequestContent) SetOvsBonds(v string)`

SetOvsBonds sets OvsBonds field to given value.

### HasOvsBonds

`func (o *CreateNetworkInterfaceRequestContent) HasOvsBonds() bool`

HasOvsBonds returns a boolean if a field has been set.

### GetOvsOptions

`func (o *CreateNetworkInterfaceRequestContent) GetOvsOptions() string`

GetOvsOptions returns the OvsOptions field if non-nil, zero value otherwise.

### GetOvsOptionsOk

`func (o *CreateNetworkInterfaceRequestContent) GetOvsOptionsOk() (*string, bool)`

GetOvsOptionsOk returns a tuple with the OvsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvsOptions

`func (o *CreateNetworkInterfaceRequestContent) SetOvsOptions(v string)`

SetOvsOptions sets OvsOptions field to given value.

### HasOvsOptions

`func (o *CreateNetworkInterfaceRequestContent) HasOvsOptions() bool`

HasOvsOptions returns a boolean if a field has been set.

### GetOvsPorts

`func (o *CreateNetworkInterfaceRequestContent) GetOvsPorts() string`

GetOvsPorts returns the OvsPorts field if non-nil, zero value otherwise.

### GetOvsPortsOk

`func (o *CreateNetworkInterfaceRequestContent) GetOvsPortsOk() (*string, bool)`

GetOvsPortsOk returns a tuple with the OvsPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvsPorts

`func (o *CreateNetworkInterfaceRequestContent) SetOvsPorts(v string)`

SetOvsPorts sets OvsPorts field to given value.

### HasOvsPorts

`func (o *CreateNetworkInterfaceRequestContent) HasOvsPorts() bool`

HasOvsPorts returns a boolean if a field has been set.

### GetOvsTag

`func (o *CreateNetworkInterfaceRequestContent) GetOvsTag() float32`

GetOvsTag returns the OvsTag field if non-nil, zero value otherwise.

### GetOvsTagOk

`func (o *CreateNetworkInterfaceRequestContent) GetOvsTagOk() (*float32, bool)`

GetOvsTagOk returns a tuple with the OvsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvsTag

`func (o *CreateNetworkInterfaceRequestContent) SetOvsTag(v float32)`

SetOvsTag sets OvsTag field to given value.

### HasOvsTag

`func (o *CreateNetworkInterfaceRequestContent) HasOvsTag() bool`

HasOvsTag returns a boolean if a field has been set.

### GetOvsBridge

`func (o *CreateNetworkInterfaceRequestContent) GetOvsBridge() string`

GetOvsBridge returns the OvsBridge field if non-nil, zero value otherwise.

### GetOvsBridgeOk

`func (o *CreateNetworkInterfaceRequestContent) GetOvsBridgeOk() (*string, bool)`

GetOvsBridgeOk returns a tuple with the OvsBridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvsBridge

`func (o *CreateNetworkInterfaceRequestContent) SetOvsBridge(v string)`

SetOvsBridge sets OvsBridge field to given value.

### HasOvsBridge

`func (o *CreateNetworkInterfaceRequestContent) HasOvsBridge() bool`

HasOvsBridge returns a boolean if a field has been set.

### GetSlaves

`func (o *CreateNetworkInterfaceRequestContent) GetSlaves() string`

GetSlaves returns the Slaves field if non-nil, zero value otherwise.

### GetSlavesOk

`func (o *CreateNetworkInterfaceRequestContent) GetSlavesOk() (*string, bool)`

GetSlavesOk returns a tuple with the Slaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaves

`func (o *CreateNetworkInterfaceRequestContent) SetSlaves(v string)`

SetSlaves sets Slaves field to given value.

### HasSlaves

`func (o *CreateNetworkInterfaceRequestContent) HasSlaves() bool`

HasSlaves returns a boolean if a field has been set.

### GetVlanId

`func (o *CreateNetworkInterfaceRequestContent) GetVlanId() float32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *CreateNetworkInterfaceRequestContent) GetVlanIdOk() (*float32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *CreateNetworkInterfaceRequestContent) SetVlanId(v float32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *CreateNetworkInterfaceRequestContent) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanRawDevice

`func (o *CreateNetworkInterfaceRequestContent) GetVlanRawDevice() string`

GetVlanRawDevice returns the VlanRawDevice field if non-nil, zero value otherwise.

### GetVlanRawDeviceOk

`func (o *CreateNetworkInterfaceRequestContent) GetVlanRawDeviceOk() (*string, bool)`

GetVlanRawDeviceOk returns a tuple with the VlanRawDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRawDevice

`func (o *CreateNetworkInterfaceRequestContent) SetVlanRawDevice(v string)`

SetVlanRawDevice sets VlanRawDevice field to given value.

### HasVlanRawDevice

`func (o *CreateNetworkInterfaceRequestContent) HasVlanRawDevice() bool`

HasVlanRawDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


