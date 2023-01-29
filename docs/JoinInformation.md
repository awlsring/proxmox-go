# JoinInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Totem** | [**TotemSummary**](TotemSummary.md) |  | 
**PreferredNode** | **string** |  | 
**ConfigDigest** | **string** |  | 
**Nodelist** | [**[]CorosyncNodeSummary**](CorosyncNodeSummary.md) |  | 

## Methods

### NewJoinInformation

`func NewJoinInformation(totem TotemSummary, preferredNode string, configDigest string, nodelist []CorosyncNodeSummary, ) *JoinInformation`

NewJoinInformation instantiates a new JoinInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinInformationWithDefaults

`func NewJoinInformationWithDefaults() *JoinInformation`

NewJoinInformationWithDefaults instantiates a new JoinInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotem

`func (o *JoinInformation) GetTotem() TotemSummary`

GetTotem returns the Totem field if non-nil, zero value otherwise.

### GetTotemOk

`func (o *JoinInformation) GetTotemOk() (*TotemSummary, bool)`

GetTotemOk returns a tuple with the Totem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotem

`func (o *JoinInformation) SetTotem(v TotemSummary)`

SetTotem sets Totem field to given value.


### GetPreferredNode

`func (o *JoinInformation) GetPreferredNode() string`

GetPreferredNode returns the PreferredNode field if non-nil, zero value otherwise.

### GetPreferredNodeOk

`func (o *JoinInformation) GetPreferredNodeOk() (*string, bool)`

GetPreferredNodeOk returns a tuple with the PreferredNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredNode

`func (o *JoinInformation) SetPreferredNode(v string)`

SetPreferredNode sets PreferredNode field to given value.


### GetConfigDigest

`func (o *JoinInformation) GetConfigDigest() string`

GetConfigDigest returns the ConfigDigest field if non-nil, zero value otherwise.

### GetConfigDigestOk

`func (o *JoinInformation) GetConfigDigestOk() (*string, bool)`

GetConfigDigestOk returns a tuple with the ConfigDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigDigest

`func (o *JoinInformation) SetConfigDigest(v string)`

SetConfigDigest sets ConfigDigest field to given value.


### GetNodelist

`func (o *JoinInformation) GetNodelist() []CorosyncNodeSummary`

GetNodelist returns the Nodelist field if non-nil, zero value otherwise.

### GetNodelistOk

`func (o *JoinInformation) GetNodelistOk() (*[]CorosyncNodeSummary, bool)`

GetNodelistOk returns a tuple with the Nodelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodelist

`func (o *JoinInformation) SetNodelist(v []CorosyncNodeSummary)`

SetNodelist sets Nodelist field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


