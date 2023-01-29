# CorosyncNodeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodeid** | **string** |  | 
**QuorumVotes** | **string** |  | 
**Ring0Addr** | **string** |  | 
**Ring1Addr** | Pointer to **string** |  | [optional] 
**Ring2Addr** | Pointer to **string** |  | [optional] 
**Ring3Addr** | Pointer to **string** |  | [optional] 
**Ring4Addr** | Pointer to **string** |  | [optional] 
**Ring5Addr** | Pointer to **string** |  | [optional] 
**Ring6Addr** | Pointer to **string** |  | [optional] 
**Ring7Addr** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**PveAddr** | Pointer to **string** |  | [optional] 

## Methods

### NewCorosyncNodeSummary

`func NewCorosyncNodeSummary(nodeid string, quorumVotes string, ring0Addr string, name string, ) *CorosyncNodeSummary`

NewCorosyncNodeSummary instantiates a new CorosyncNodeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorosyncNodeSummaryWithDefaults

`func NewCorosyncNodeSummaryWithDefaults() *CorosyncNodeSummary`

NewCorosyncNodeSummaryWithDefaults instantiates a new CorosyncNodeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeid

`func (o *CorosyncNodeSummary) GetNodeid() string`

GetNodeid returns the Nodeid field if non-nil, zero value otherwise.

### GetNodeidOk

`func (o *CorosyncNodeSummary) GetNodeidOk() (*string, bool)`

GetNodeidOk returns a tuple with the Nodeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeid

`func (o *CorosyncNodeSummary) SetNodeid(v string)`

SetNodeid sets Nodeid field to given value.


### GetQuorumVotes

`func (o *CorosyncNodeSummary) GetQuorumVotes() string`

GetQuorumVotes returns the QuorumVotes field if non-nil, zero value otherwise.

### GetQuorumVotesOk

`func (o *CorosyncNodeSummary) GetQuorumVotesOk() (*string, bool)`

GetQuorumVotesOk returns a tuple with the QuorumVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorumVotes

`func (o *CorosyncNodeSummary) SetQuorumVotes(v string)`

SetQuorumVotes sets QuorumVotes field to given value.


### GetRing0Addr

`func (o *CorosyncNodeSummary) GetRing0Addr() string`

GetRing0Addr returns the Ring0Addr field if non-nil, zero value otherwise.

### GetRing0AddrOk

`func (o *CorosyncNodeSummary) GetRing0AddrOk() (*string, bool)`

GetRing0AddrOk returns a tuple with the Ring0Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRing0Addr

`func (o *CorosyncNodeSummary) SetRing0Addr(v string)`

SetRing0Addr sets Ring0Addr field to given value.


### GetRing1Addr

`func (o *CorosyncNodeSummary) GetRing1Addr() string`

GetRing1Addr returns the Ring1Addr field if non-nil, zero value otherwise.

### GetRing1AddrOk

`func (o *CorosyncNodeSummary) GetRing1AddrOk() (*string, bool)`

GetRing1AddrOk returns a tuple with the Ring1Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRing1Addr

`func (o *CorosyncNodeSummary) SetRing1Addr(v string)`

SetRing1Addr sets Ring1Addr field to given value.

### HasRing1Addr

`func (o *CorosyncNodeSummary) HasRing1Addr() bool`

HasRing1Addr returns a boolean if a field has been set.

### GetRing2Addr

`func (o *CorosyncNodeSummary) GetRing2Addr() string`

GetRing2Addr returns the Ring2Addr field if non-nil, zero value otherwise.

### GetRing2AddrOk

`func (o *CorosyncNodeSummary) GetRing2AddrOk() (*string, bool)`

GetRing2AddrOk returns a tuple with the Ring2Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRing2Addr

`func (o *CorosyncNodeSummary) SetRing2Addr(v string)`

SetRing2Addr sets Ring2Addr field to given value.

### HasRing2Addr

`func (o *CorosyncNodeSummary) HasRing2Addr() bool`

HasRing2Addr returns a boolean if a field has been set.

### GetRing3Addr

`func (o *CorosyncNodeSummary) GetRing3Addr() string`

GetRing3Addr returns the Ring3Addr field if non-nil, zero value otherwise.

### GetRing3AddrOk

`func (o *CorosyncNodeSummary) GetRing3AddrOk() (*string, bool)`

GetRing3AddrOk returns a tuple with the Ring3Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRing3Addr

`func (o *CorosyncNodeSummary) SetRing3Addr(v string)`

SetRing3Addr sets Ring3Addr field to given value.

### HasRing3Addr

`func (o *CorosyncNodeSummary) HasRing3Addr() bool`

HasRing3Addr returns a boolean if a field has been set.

### GetRing4Addr

`func (o *CorosyncNodeSummary) GetRing4Addr() string`

GetRing4Addr returns the Ring4Addr field if non-nil, zero value otherwise.

### GetRing4AddrOk

`func (o *CorosyncNodeSummary) GetRing4AddrOk() (*string, bool)`

GetRing4AddrOk returns a tuple with the Ring4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRing4Addr

`func (o *CorosyncNodeSummary) SetRing4Addr(v string)`

SetRing4Addr sets Ring4Addr field to given value.

### HasRing4Addr

`func (o *CorosyncNodeSummary) HasRing4Addr() bool`

HasRing4Addr returns a boolean if a field has been set.

### GetRing5Addr

`func (o *CorosyncNodeSummary) GetRing5Addr() string`

GetRing5Addr returns the Ring5Addr field if non-nil, zero value otherwise.

### GetRing5AddrOk

`func (o *CorosyncNodeSummary) GetRing5AddrOk() (*string, bool)`

GetRing5AddrOk returns a tuple with the Ring5Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRing5Addr

`func (o *CorosyncNodeSummary) SetRing5Addr(v string)`

SetRing5Addr sets Ring5Addr field to given value.

### HasRing5Addr

`func (o *CorosyncNodeSummary) HasRing5Addr() bool`

HasRing5Addr returns a boolean if a field has been set.

### GetRing6Addr

`func (o *CorosyncNodeSummary) GetRing6Addr() string`

GetRing6Addr returns the Ring6Addr field if non-nil, zero value otherwise.

### GetRing6AddrOk

`func (o *CorosyncNodeSummary) GetRing6AddrOk() (*string, bool)`

GetRing6AddrOk returns a tuple with the Ring6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRing6Addr

`func (o *CorosyncNodeSummary) SetRing6Addr(v string)`

SetRing6Addr sets Ring6Addr field to given value.

### HasRing6Addr

`func (o *CorosyncNodeSummary) HasRing6Addr() bool`

HasRing6Addr returns a boolean if a field has been set.

### GetRing7Addr

`func (o *CorosyncNodeSummary) GetRing7Addr() string`

GetRing7Addr returns the Ring7Addr field if non-nil, zero value otherwise.

### GetRing7AddrOk

`func (o *CorosyncNodeSummary) GetRing7AddrOk() (*string, bool)`

GetRing7AddrOk returns a tuple with the Ring7Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRing7Addr

`func (o *CorosyncNodeSummary) SetRing7Addr(v string)`

SetRing7Addr sets Ring7Addr field to given value.

### HasRing7Addr

`func (o *CorosyncNodeSummary) HasRing7Addr() bool`

HasRing7Addr returns a boolean if a field has been set.

### GetName

`func (o *CorosyncNodeSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorosyncNodeSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorosyncNodeSummary) SetName(v string)`

SetName sets Name field to given value.


### GetPveAddr

`func (o *CorosyncNodeSummary) GetPveAddr() string`

GetPveAddr returns the PveAddr field if non-nil, zero value otherwise.

### GetPveAddrOk

`func (o *CorosyncNodeSummary) GetPveAddrOk() (*string, bool)`

GetPveAddrOk returns a tuple with the PveAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPveAddr

`func (o *CorosyncNodeSummary) SetPveAddr(v string)`

SetPveAddr sets PveAddr field to given value.

### HasPveAddr

`func (o *CorosyncNodeSummary) HasPveAddr() bool`

HasPveAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


