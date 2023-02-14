# ListLVMsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leaf** | **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | 
**Children** | [**[]LVMChild**](LVMChild.md) |  | 

## Methods

### NewListLVMsData

`func NewListLVMsData(leaf float32, children []LVMChild, ) *ListLVMsData`

NewListLVMsData instantiates a new ListLVMsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLVMsDataWithDefaults

`func NewListLVMsDataWithDefaults() *ListLVMsData`

NewListLVMsDataWithDefaults instantiates a new ListLVMsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaf

`func (o *ListLVMsData) GetLeaf() float32`

GetLeaf returns the Leaf field if non-nil, zero value otherwise.

### GetLeafOk

`func (o *ListLVMsData) GetLeafOk() (*float32, bool)`

GetLeafOk returns a tuple with the Leaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaf

`func (o *ListLVMsData) SetLeaf(v float32)`

SetLeaf sets Leaf field to given value.


### GetChildren

`func (o *ListLVMsData) GetChildren() []LVMChild`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ListLVMsData) GetChildrenOk() (*[]LVMChild, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ListLVMsData) SetChildren(v []LVMChild)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


