# UpdateNetworkInterfaceRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Delete** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Gateway6** | Pointer to **string** |  | [optional] 
**Mtu** | Pointer to **float32** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Netmask6** | Pointer to **string** |  | [optional] 
**OvsBonds** | Pointer to **string** |  | [optional] 
**OvsOptions** | Pointer to **string** |  | [optional] 
**OvsPorts** | Pointer to **string** |  | [optional] 
**OvsTag** | Pointer to **float32** |  | [optional] 
**OvsBridge** | Pointer to **string** |  | [optional] 
**Slaves** | Pointer to **string** |  | [optional] 
**VlanId** | Pointer to **float32** |  | [optional] 
**VlanRawDevice** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateNetworkInterfaceRequestContent

`func NewUpdateNetworkInterfaceRequestContent(type_ NetworkInterfaceType, ) *UpdateNetworkInterfaceRequestContent`

NewUpdateNetworkInterfaceRequestContent instantiates a new UpdateNetworkInterfaceRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkInterfaceRequestContentWithDefaults

`func NewUpdateNetworkInterfaceRequestContentWithDefaults() *UpdateNetworkInterfaceRequestContent`

NewUpdateNetworkInterfaceRequestContentWithDefaults instantiates a new UpdateNetworkInterfaceRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateNetworkInterfaceRequestContent) GetType() NetworkInterfaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateNetworkInterfaceRequestContent) GetTypeOk() (*NetworkInterfaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateNetworkInterfaceRequestContent) SetType(v NetworkInterfaceType)`

SetType sets Type field to given value.


### GetAddress

`func (o *UpdateNetworkInterfaceRequestContent) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateNetworkInterfaceRequestContent) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateNetworkInterfaceRequestContent) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateNetworkInterfaceRequestContent) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddress6

`func (o *UpdateNetworkInterfaceRequestContent) GetAddress6() string`

GetAddress6 returns the Address6 field if non-nil, zero value otherwise.

### GetAddress6Ok

`func (o *UpdateNetworkInterfaceRequestContent) GetAddress6Ok() (*string, bool)`

GetAddress6Ok returns a tuple with the Address6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6

`func (o *UpdateNetworkInterfaceRequestContent) SetAddress6(v string)`

SetAddress6 sets Address6 field to given value.

### HasAddress6

`func (o *UpdateNetworkInterfaceRequestContent) HasAddress6() bool`

HasAddress6 returns a boolean if a field has been set.

### GetAutostart

`func (o *UpdateNetworkInterfaceRequestContent) GetAutostart() bool`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *UpdateNetworkInterfaceRequestContent) GetAutostartOk() (*bool, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *UpdateNetworkInterfaceRequestContent) SetAutostart(v bool)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *UpdateNetworkInterfaceRequestContent) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetBondPrimary

`func (o *UpdateNetworkInterfaceRequestContent) GetBondPrimary() string`

GetBondPrimary returns the BondPrimary field if non-nil, zero value otherwise.

### GetBondPrimaryOk

`func (o *UpdateNetworkInterfaceRequestContent) GetBondPrimaryOk() (*string, bool)`

GetBondPrimaryOk returns a tuple with the BondPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondPrimary

`func (o *UpdateNetworkInterfaceRequestContent) SetBondPrimary(v string)`

SetBondPrimary sets BondPrimary field to given value.

### HasBondPrimary

`func (o *UpdateNetworkInterfaceRequestContent) HasBondPrimary() bool`

HasBondPrimary returns a boolean if a field has been set.

### GetBondMode

`func (o *UpdateNetworkInterfaceRequestContent) GetBondMode() NetworkInterfaceBondMode`

GetBondMode returns the BondMode field if non-nil, zero value otherwise.

### GetBondModeOk

`func (o *UpdateNetworkInterfaceRequestContent) GetBondModeOk() (*NetworkInterfaceBondMode, bool)`

GetBondModeOk returns a tuple with the BondMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondMode

`func (o *UpdateNetworkInterfaceRequestContent) SetBondMode(v NetworkInterfaceBondMode)`

SetBondMode sets BondMode field to given value.

### HasBondMode

`func (o *UpdateNetworkInterfaceRequestContent) HasBondMode() bool`

HasBondMode returns a boolean if a field has been set.

### GetBondXmitHashPolicy

`func (o *UpdateNetworkInterfaceRequestContent) GetBondXmitHashPolicy() NetworkInterfaceBondHashPolicy`

GetBondXmitHashPolicy returns the BondXmitHashPolicy field if non-nil, zero value otherwise.

### GetBondXmitHashPolicyOk

`func (o *UpdateNetworkInterfaceRequestContent) GetBondXmitHashPolicyOk() (*NetworkInterfaceBondHashPolicy, bool)`

GetBondXmitHashPolicyOk returns a tuple with the BondXmitHashPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondXmitHashPolicy

`func (o *UpdateNetworkInterfaceRequestContent) SetBondXmitHashPolicy(v NetworkInterfaceBondHashPolicy)`

SetBondXmitHashPolicy sets BondXmitHashPolicy field to given value.

### HasBondXmitHashPolicy

`func (o *UpdateNetworkInterfaceRequestContent) HasBondXmitHashPolicy() bool`

HasBondXmitHashPolicy returns a boolean if a field has been set.

### GetBridgePorts

`func (o *UpdateNetworkInterfaceRequestContent) GetBridgePorts() string`

GetBridgePorts returns the BridgePorts field if non-nil, zero value otherwise.

### GetBridgePortsOk

`func (o *UpdateNetworkInterfaceRequestContent) GetBridgePortsOk() (*string, bool)`

GetBridgePortsOk returns a tuple with the BridgePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgePorts

`func (o *UpdateNetworkInterfaceRequestContent) SetBridgePorts(v string)`

SetBridgePorts sets BridgePorts field to given value.

### HasBridgePorts

`func (o *UpdateNetworkInterfaceRequestContent) HasBridgePorts() bool`

HasBridgePorts returns a boolean if a field has been set.

### GetBridgeVlanAware

`func (o *UpdateNetworkInterfaceRequestContent) GetBridgeVlanAware() bool`

GetBridgeVlanAware returns the BridgeVlanAware field if non-nil, zero value otherwise.

### GetBridgeVlanAwareOk

`func (o *UpdateNetworkInterfaceRequestContent) GetBridgeVlanAwareOk() (*bool, bool)`

GetBridgeVlanAwareOk returns a tuple with the BridgeVlanAware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVlanAware

`func (o *UpdateNetworkInterfaceRequestContent) SetBridgeVlanAware(v bool)`

SetBridgeVlanAware sets BridgeVlanAware field to given value.

### HasBridgeVlanAware

`func (o *UpdateNetworkInterfaceRequestContent) HasBridgeVlanAware() bool`

HasBridgeVlanAware returns a boolean if a field has been set.

### GetCidr

`func (o *UpdateNetworkInterfaceRequestContent) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *UpdateNetworkInterfaceRequestContent) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *UpdateNetworkInterfaceRequestContent) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *UpdateNetworkInterfaceRequestContent) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCidr6

`func (o *UpdateNetworkInterfaceRequestContent) GetCidr6() string`

GetCidr6 returns the Cidr6 field if non-nil, zero value otherwise.

### GetCidr6Ok

`func (o *UpdateNetworkInterfaceRequestContent) GetCidr6Ok() (*string, bool)`

GetCidr6Ok returns a tuple with the Cidr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr6

`func (o *UpdateNetworkInterfaceRequestContent) SetCidr6(v string)`

SetCidr6 sets Cidr6 field to given value.

### HasCidr6

`func (o *UpdateNetworkInterfaceRequestContent) HasCidr6() bool`

HasCidr6 returns a boolean if a field has been set.

### GetComments

`func (o *UpdateNetworkInterfaceRequestContent) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *UpdateNetworkInterfaceRequestContent) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *UpdateNetworkInterfaceRequestContent) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *UpdateNetworkInterfaceRequestContent) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetComments6

`func (o *UpdateNetworkInterfaceRequestContent) GetComments6() string`

GetComments6 returns the Comments6 field if non-nil, zero value otherwise.

### GetComments6Ok

`func (o *UpdateNetworkInterfaceRequestContent) GetComments6Ok() (*string, bool)`

GetComments6Ok returns a tuple with the Comments6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments6

`func (o *UpdateNetworkInterfaceRequestContent) SetComments6(v string)`

SetComments6 sets Comments6 field to given value.

### HasComments6

`func (o *UpdateNetworkInterfaceRequestContent) HasComments6() bool`

HasComments6 returns a boolean if a field has been set.

### GetDelete

`func (o *UpdateNetworkInterfaceRequestContent) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *UpdateNetworkInterfaceRequestContent) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *UpdateNetworkInterfaceRequestContent) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *UpdateNetworkInterfaceRequestContent) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetGateway

`func (o *UpdateNetworkInterfaceRequestContent) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UpdateNetworkInterfaceRequestContent) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UpdateNetworkInterfaceRequestContent) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UpdateNetworkInterfaceRequestContent) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGateway6

`func (o *UpdateNetworkInterfaceRequestContent) GetGateway6() string`

GetGateway6 returns the Gateway6 field if non-nil, zero value otherwise.

### GetGateway6Ok

`func (o *UpdateNetworkInterfaceRequestContent) GetGateway6Ok() (*string, bool)`

GetGateway6Ok returns a tuple with the Gateway6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway6

`func (o *UpdateNetworkInterfaceRequestContent) SetGateway6(v string)`

SetGateway6 sets Gateway6 field to given value.

### HasGateway6

`func (o *UpdateNetworkInterfaceRequestContent) HasGateway6() bool`

HasGateway6 returns a boolean if a field has been set.

### GetMtu

`func (o *UpdateNetworkInterfaceRequestContent) GetMtu() float32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *UpdateNetworkInterfaceRequestContent) GetMtuOk() (*float32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *UpdateNetworkInterfaceRequestContent) SetMtu(v float32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *UpdateNetworkInterfaceRequestContent) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetNetmask

`func (o *UpdateNetworkInterfaceRequestContent) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *UpdateNetworkInterfaceRequestContent) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *UpdateNetworkInterfaceRequestContent) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *UpdateNetworkInterfaceRequestContent) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetmask6

`func (o *UpdateNetworkInterfaceRequestContent) GetNetmask6() string`

GetNetmask6 returns the Netmask6 field if non-nil, zero value otherwise.

### GetNetmask6Ok

`func (o *UpdateNetworkInterfaceRequestContent) GetNetmask6Ok() (*string, bool)`

GetNetmask6Ok returns a tuple with the Netmask6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask6

`func (o *UpdateNetworkInterfaceRequestContent) SetNetmask6(v string)`

SetNetmask6 sets Netmask6 field to given value.

### HasNetmask6

`func (o *UpdateNetworkInterfaceRequestContent) HasNetmask6() bool`

HasNetmask6 returns a boolean if a field has been set.

### GetOvsBonds

`func (o *UpdateNetworkInterfaceRequestContent) GetOvsBonds() string`

GetOvsBonds returns the OvsBonds field if non-nil, zero value otherwise.

### GetOvsBondsOk

`func (o *UpdateNetworkInterfaceRequestContent) GetOvsBondsOk() (*string, bool)`

GetOvsBondsOk returns a tuple with the OvsBonds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvsBonds

`func (o *UpdateNetworkInterfaceRequestContent) SetOvsBonds(v string)`

SetOvsBonds sets OvsBonds field to given value.

### HasOvsBonds

`func (o *UpdateNetworkInterfaceRequestContent) HasOvsBonds() bool`

HasOvsBonds returns a boolean if a field has been set.

### GetOvsOptions

`func (o *UpdateNetworkInterfaceRequestContent) GetOvsOptions() string`

GetOvsOptions returns the OvsOptions field if non-nil, zero value otherwise.

### GetOvsOptionsOk

`func (o *UpdateNetworkInterfaceRequestContent) GetOvsOptionsOk() (*string, bool)`

GetOvsOptionsOk returns a tuple with the OvsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvsOptions

`func (o *UpdateNetworkInterfaceRequestContent) SetOvsOptions(v string)`

SetOvsOptions sets OvsOptions field to given value.

### HasOvsOptions

`func (o *UpdateNetworkInterfaceRequestContent) HasOvsOptions() bool`

HasOvsOptions returns a boolean if a field has been set.

### GetOvsPorts

`func (o *UpdateNetworkInterfaceRequestContent) GetOvsPorts() string`

GetOvsPorts returns the OvsPorts field if non-nil, zero value otherwise.

### GetOvsPortsOk

`func (o *UpdateNetworkInterfaceRequestContent) GetOvsPortsOk() (*string, bool)`

GetOvsPortsOk returns a tuple with the OvsPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvsPorts

`func (o *UpdateNetworkInterfaceRequestContent) SetOvsPorts(v string)`

SetOvsPorts sets OvsPorts field to given value.

### HasOvsPorts

`func (o *UpdateNetworkInterfaceRequestContent) HasOvsPorts() bool`

HasOvsPorts returns a boolean if a field has been set.

### GetOvsTag

`func (o *UpdateNetworkInterfaceRequestContent) GetOvsTag() float32`

GetOvsTag returns the OvsTag field if non-nil, zero value otherwise.

### GetOvsTagOk

`func (o *UpdateNetworkInterfaceRequestContent) GetOvsTagOk() (*float32, bool)`

GetOvsTagOk returns a tuple with the OvsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvsTag

`func (o *UpdateNetworkInterfaceRequestContent) SetOvsTag(v float32)`

SetOvsTag sets OvsTag field to given value.

### HasOvsTag

`func (o *UpdateNetworkInterfaceRequestContent) HasOvsTag() bool`

HasOvsTag returns a boolean if a field has been set.

### GetOvsBridge

`func (o *UpdateNetworkInterfaceRequestContent) GetOvsBridge() string`

GetOvsBridge returns the OvsBridge field if non-nil, zero value otherwise.

### GetOvsBridgeOk

`func (o *UpdateNetworkInterfaceRequestContent) GetOvsBridgeOk() (*string, bool)`

GetOvsBridgeOk returns a tuple with the OvsBridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvsBridge

`func (o *UpdateNetworkInterfaceRequestContent) SetOvsBridge(v string)`

SetOvsBridge sets OvsBridge field to given value.

### HasOvsBridge

`func (o *UpdateNetworkInterfaceRequestContent) HasOvsBridge() bool`

HasOvsBridge returns a boolean if a field has been set.

### GetSlaves

`func (o *UpdateNetworkInterfaceRequestContent) GetSlaves() string`

GetSlaves returns the Slaves field if non-nil, zero value otherwise.

### GetSlavesOk

`func (o *UpdateNetworkInterfaceRequestContent) GetSlavesOk() (*string, bool)`

GetSlavesOk returns a tuple with the Slaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaves

`func (o *UpdateNetworkInterfaceRequestContent) SetSlaves(v string)`

SetSlaves sets Slaves field to given value.

### HasSlaves

`func (o *UpdateNetworkInterfaceRequestContent) HasSlaves() bool`

HasSlaves returns a boolean if a field has been set.

### GetVlanId

`func (o *UpdateNetworkInterfaceRequestContent) GetVlanId() float32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *UpdateNetworkInterfaceRequestContent) GetVlanIdOk() (*float32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *UpdateNetworkInterfaceRequestContent) SetVlanId(v float32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *UpdateNetworkInterfaceRequestContent) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanRawDevice

`func (o *UpdateNetworkInterfaceRequestContent) GetVlanRawDevice() string`

GetVlanRawDevice returns the VlanRawDevice field if non-nil, zero value otherwise.

### GetVlanRawDeviceOk

`func (o *UpdateNetworkInterfaceRequestContent) GetVlanRawDeviceOk() (*string, bool)`

GetVlanRawDeviceOk returns a tuple with the VlanRawDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanRawDevice

`func (o *UpdateNetworkInterfaceRequestContent) SetVlanRawDevice(v string)`

SetVlanRawDevice sets VlanRawDevice field to given value.

### HasVlanRawDevice

`func (o *UpdateNetworkInterfaceRequestContent) HasVlanRawDevice() bool`

HasVlanRawDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


