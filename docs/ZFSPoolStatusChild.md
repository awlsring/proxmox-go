# ZFSPoolStatusChild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cksum** | **float32** |  | 
**Read** | **float32** |  | 
**Write** | **float32** |  | 
**Name** | **string** |  | 
**State** | **string** |  | 
**Msg** | Pointer to **string** |  | [optional] 
**Children** | Pointer to [**[]ZFSPoolStatusChild**](ZFSPoolStatusChild.md) |  | [optional] 

## Methods

### NewZFSPoolStatusChild

`func NewZFSPoolStatusChild(cksum float32, read float32, write float32, name string, state string, ) *ZFSPoolStatusChild`

NewZFSPoolStatusChild instantiates a new ZFSPoolStatusChild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZFSPoolStatusChildWithDefaults

`func NewZFSPoolStatusChildWithDefaults() *ZFSPoolStatusChild`

NewZFSPoolStatusChildWithDefaults instantiates a new ZFSPoolStatusChild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCksum

`func (o *ZFSPoolStatusChild) GetCksum() float32`

GetCksum returns the Cksum field if non-nil, zero value otherwise.

### GetCksumOk

`func (o *ZFSPoolStatusChild) GetCksumOk() (*float32, bool)`

GetCksumOk returns a tuple with the Cksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCksum

`func (o *ZFSPoolStatusChild) SetCksum(v float32)`

SetCksum sets Cksum field to given value.


### GetRead

`func (o *ZFSPoolStatusChild) GetRead() float32`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *ZFSPoolStatusChild) GetReadOk() (*float32, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *ZFSPoolStatusChild) SetRead(v float32)`

SetRead sets Read field to given value.


### GetWrite

`func (o *ZFSPoolStatusChild) GetWrite() float32`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *ZFSPoolStatusChild) GetWriteOk() (*float32, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *ZFSPoolStatusChild) SetWrite(v float32)`

SetWrite sets Write field to given value.


### GetName

`func (o *ZFSPoolStatusChild) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZFSPoolStatusChild) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZFSPoolStatusChild) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *ZFSPoolStatusChild) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ZFSPoolStatusChild) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ZFSPoolStatusChild) SetState(v string)`

SetState sets State field to given value.


### GetMsg

`func (o *ZFSPoolStatusChild) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ZFSPoolStatusChild) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ZFSPoolStatusChild) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ZFSPoolStatusChild) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetChildren

`func (o *ZFSPoolStatusChild) GetChildren() []ZFSPoolStatusChild`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ZFSPoolStatusChild) GetChildrenOk() (*[]ZFSPoolStatusChild, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ZFSPoolStatusChild) SetChildren(v []ZFSPoolStatusChild)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ZFSPoolStatusChild) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


