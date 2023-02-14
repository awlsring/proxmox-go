# LVMChild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]LVMChild**](LVMChild.md) |  | [optional] 
**Leaf** | **float32** | Is leaf. This is a boolean intergar, 1 is true, 0 is false | 
**Name** | Pointer to **string** |  | [optional] 
**Free** | Pointer to **float32** | The free space on lvm in bytes | [optional] 
**Size** | Pointer to **float32** | The total size of lvm in bytes | [optional] 
**Lvcount** | Pointer to **float32** | The number of logical volumes | [optional] 

## Methods

### NewLVMChild

`func NewLVMChild(leaf float32, ) *LVMChild`

NewLVMChild instantiates a new LVMChild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLVMChildWithDefaults

`func NewLVMChildWithDefaults() *LVMChild`

NewLVMChildWithDefaults instantiates a new LVMChild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *LVMChild) GetChildren() []LVMChild`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *LVMChild) GetChildrenOk() (*[]LVMChild, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *LVMChild) SetChildren(v []LVMChild)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *LVMChild) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetLeaf

`func (o *LVMChild) GetLeaf() float32`

GetLeaf returns the Leaf field if non-nil, zero value otherwise.

### GetLeafOk

`func (o *LVMChild) GetLeafOk() (*float32, bool)`

GetLeafOk returns a tuple with the Leaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaf

`func (o *LVMChild) SetLeaf(v float32)`

SetLeaf sets Leaf field to given value.


### GetName

`func (o *LVMChild) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LVMChild) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LVMChild) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LVMChild) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFree

`func (o *LVMChild) GetFree() float32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *LVMChild) GetFreeOk() (*float32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *LVMChild) SetFree(v float32)`

SetFree sets Free field to given value.

### HasFree

`func (o *LVMChild) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetSize

`func (o *LVMChild) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LVMChild) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LVMChild) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *LVMChild) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLvcount

`func (o *LVMChild) GetLvcount() float32`

GetLvcount returns the Lvcount field if non-nil, zero value otherwise.

### GetLvcountOk

`func (o *LVMChild) GetLvcountOk() (*float32, bool)`

GetLvcountOk returns a tuple with the Lvcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvcount

`func (o *LVMChild) SetLvcount(v float32)`

SetLvcount sets Lvcount field to given value.

### HasLvcount

`func (o *LVMChild) HasLvcount() bool`

HasLvcount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


