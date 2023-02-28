# VirtualMachineFirewallReferenceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FirewallReferenceType**](FirewallReferenceType.md) |  | 
**Name** | **string** |  | 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewVirtualMachineFirewallReferenceSummary

`func NewVirtualMachineFirewallReferenceSummary(type_ FirewallReferenceType, name string, ) *VirtualMachineFirewallReferenceSummary`

NewVirtualMachineFirewallReferenceSummary instantiates a new VirtualMachineFirewallReferenceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineFirewallReferenceSummaryWithDefaults

`func NewVirtualMachineFirewallReferenceSummaryWithDefaults() *VirtualMachineFirewallReferenceSummary`

NewVirtualMachineFirewallReferenceSummaryWithDefaults instantiates a new VirtualMachineFirewallReferenceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VirtualMachineFirewallReferenceSummary) GetType() FirewallReferenceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualMachineFirewallReferenceSummary) GetTypeOk() (*FirewallReferenceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualMachineFirewallReferenceSummary) SetType(v FirewallReferenceType)`

SetType sets Type field to given value.


### GetName

`func (o *VirtualMachineFirewallReferenceSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualMachineFirewallReferenceSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualMachineFirewallReferenceSummary) SetName(v string)`

SetName sets Name field to given value.


### GetComment

`func (o *VirtualMachineFirewallReferenceSummary) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *VirtualMachineFirewallReferenceSummary) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *VirtualMachineFirewallReferenceSummary) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *VirtualMachineFirewallReferenceSummary) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


