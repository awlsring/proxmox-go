# AddCorosyncNodeRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apiversion** | Pointer to **float32** | The api version on the new node | [optional] 
**Force** | Pointer to **bool** | Do not throw an error if the node is already in the cluster | [optional] 
**Link0** | Pointer to **string** | Address and priority of the link. Input as string with format address&#x3D;&lt;ip&gt;,priority&#x3D;&lt;int&gt; | [optional] 
**Link1** | Pointer to **string** | Address and priority of the link. Input as string with format address&#x3D;&lt;ip&gt;,priority&#x3D;&lt;int&gt; | [optional] 
**Link2** | Pointer to **string** | Address and priority of the link. Input as string with format address&#x3D;&lt;ip&gt;,priority&#x3D;&lt;int&gt; | [optional] 
**Link3** | Pointer to **string** | Address and priority of the link. Input as string with format address&#x3D;&lt;ip&gt;,priority&#x3D;&lt;int&gt; | [optional] 
**Link4** | Pointer to **string** | Address and priority of the link. Input as string with format address&#x3D;&lt;ip&gt;,priority&#x3D;&lt;int&gt; | [optional] 
**Link5** | Pointer to **string** | Address and priority of the link. Input as string with format address&#x3D;&lt;ip&gt;,priority&#x3D;&lt;int&gt; | [optional] 
**Link6** | Pointer to **string** | Address and priority of the link. Input as string with format address&#x3D;&lt;ip&gt;,priority&#x3D;&lt;int&gt; | [optional] 
**Link7** | Pointer to **string** | Address and priority of the link. Input as string with format address&#x3D;&lt;ip&gt;,priority&#x3D;&lt;int&gt; | [optional] 
**NewNodeIp** | Pointer to **string** | The IP address of the node to add | [optional] 
**Nodeid** | Pointer to **float32** | NodeID of the node to add | [optional] 
**Votes** | Pointer to **float32** | Votes thew new node should have | [optional] 

## Methods

### NewAddCorosyncNodeRequestContent

`func NewAddCorosyncNodeRequestContent() *AddCorosyncNodeRequestContent`

NewAddCorosyncNodeRequestContent instantiates a new AddCorosyncNodeRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCorosyncNodeRequestContentWithDefaults

`func NewAddCorosyncNodeRequestContentWithDefaults() *AddCorosyncNodeRequestContent`

NewAddCorosyncNodeRequestContentWithDefaults instantiates a new AddCorosyncNodeRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiversion

`func (o *AddCorosyncNodeRequestContent) GetApiversion() float32`

GetApiversion returns the Apiversion field if non-nil, zero value otherwise.

### GetApiversionOk

`func (o *AddCorosyncNodeRequestContent) GetApiversionOk() (*float32, bool)`

GetApiversionOk returns a tuple with the Apiversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiversion

`func (o *AddCorosyncNodeRequestContent) SetApiversion(v float32)`

SetApiversion sets Apiversion field to given value.

### HasApiversion

`func (o *AddCorosyncNodeRequestContent) HasApiversion() bool`

HasApiversion returns a boolean if a field has been set.

### GetForce

`func (o *AddCorosyncNodeRequestContent) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *AddCorosyncNodeRequestContent) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *AddCorosyncNodeRequestContent) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *AddCorosyncNodeRequestContent) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetLink0

`func (o *AddCorosyncNodeRequestContent) GetLink0() string`

GetLink0 returns the Link0 field if non-nil, zero value otherwise.

### GetLink0Ok

`func (o *AddCorosyncNodeRequestContent) GetLink0Ok() (*string, bool)`

GetLink0Ok returns a tuple with the Link0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink0

`func (o *AddCorosyncNodeRequestContent) SetLink0(v string)`

SetLink0 sets Link0 field to given value.

### HasLink0

`func (o *AddCorosyncNodeRequestContent) HasLink0() bool`

HasLink0 returns a boolean if a field has been set.

### GetLink1

`func (o *AddCorosyncNodeRequestContent) GetLink1() string`

GetLink1 returns the Link1 field if non-nil, zero value otherwise.

### GetLink1Ok

`func (o *AddCorosyncNodeRequestContent) GetLink1Ok() (*string, bool)`

GetLink1Ok returns a tuple with the Link1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink1

`func (o *AddCorosyncNodeRequestContent) SetLink1(v string)`

SetLink1 sets Link1 field to given value.

### HasLink1

`func (o *AddCorosyncNodeRequestContent) HasLink1() bool`

HasLink1 returns a boolean if a field has been set.

### GetLink2

`func (o *AddCorosyncNodeRequestContent) GetLink2() string`

GetLink2 returns the Link2 field if non-nil, zero value otherwise.

### GetLink2Ok

`func (o *AddCorosyncNodeRequestContent) GetLink2Ok() (*string, bool)`

GetLink2Ok returns a tuple with the Link2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink2

`func (o *AddCorosyncNodeRequestContent) SetLink2(v string)`

SetLink2 sets Link2 field to given value.

### HasLink2

`func (o *AddCorosyncNodeRequestContent) HasLink2() bool`

HasLink2 returns a boolean if a field has been set.

### GetLink3

`func (o *AddCorosyncNodeRequestContent) GetLink3() string`

GetLink3 returns the Link3 field if non-nil, zero value otherwise.

### GetLink3Ok

`func (o *AddCorosyncNodeRequestContent) GetLink3Ok() (*string, bool)`

GetLink3Ok returns a tuple with the Link3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink3

`func (o *AddCorosyncNodeRequestContent) SetLink3(v string)`

SetLink3 sets Link3 field to given value.

### HasLink3

`func (o *AddCorosyncNodeRequestContent) HasLink3() bool`

HasLink3 returns a boolean if a field has been set.

### GetLink4

`func (o *AddCorosyncNodeRequestContent) GetLink4() string`

GetLink4 returns the Link4 field if non-nil, zero value otherwise.

### GetLink4Ok

`func (o *AddCorosyncNodeRequestContent) GetLink4Ok() (*string, bool)`

GetLink4Ok returns a tuple with the Link4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink4

`func (o *AddCorosyncNodeRequestContent) SetLink4(v string)`

SetLink4 sets Link4 field to given value.

### HasLink4

`func (o *AddCorosyncNodeRequestContent) HasLink4() bool`

HasLink4 returns a boolean if a field has been set.

### GetLink5

`func (o *AddCorosyncNodeRequestContent) GetLink5() string`

GetLink5 returns the Link5 field if non-nil, zero value otherwise.

### GetLink5Ok

`func (o *AddCorosyncNodeRequestContent) GetLink5Ok() (*string, bool)`

GetLink5Ok returns a tuple with the Link5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink5

`func (o *AddCorosyncNodeRequestContent) SetLink5(v string)`

SetLink5 sets Link5 field to given value.

### HasLink5

`func (o *AddCorosyncNodeRequestContent) HasLink5() bool`

HasLink5 returns a boolean if a field has been set.

### GetLink6

`func (o *AddCorosyncNodeRequestContent) GetLink6() string`

GetLink6 returns the Link6 field if non-nil, zero value otherwise.

### GetLink6Ok

`func (o *AddCorosyncNodeRequestContent) GetLink6Ok() (*string, bool)`

GetLink6Ok returns a tuple with the Link6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink6

`func (o *AddCorosyncNodeRequestContent) SetLink6(v string)`

SetLink6 sets Link6 field to given value.

### HasLink6

`func (o *AddCorosyncNodeRequestContent) HasLink6() bool`

HasLink6 returns a boolean if a field has been set.

### GetLink7

`func (o *AddCorosyncNodeRequestContent) GetLink7() string`

GetLink7 returns the Link7 field if non-nil, zero value otherwise.

### GetLink7Ok

`func (o *AddCorosyncNodeRequestContent) GetLink7Ok() (*string, bool)`

GetLink7Ok returns a tuple with the Link7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink7

`func (o *AddCorosyncNodeRequestContent) SetLink7(v string)`

SetLink7 sets Link7 field to given value.

### HasLink7

`func (o *AddCorosyncNodeRequestContent) HasLink7() bool`

HasLink7 returns a boolean if a field has been set.

### GetNewNodeIp

`func (o *AddCorosyncNodeRequestContent) GetNewNodeIp() string`

GetNewNodeIp returns the NewNodeIp field if non-nil, zero value otherwise.

### GetNewNodeIpOk

`func (o *AddCorosyncNodeRequestContent) GetNewNodeIpOk() (*string, bool)`

GetNewNodeIpOk returns a tuple with the NewNodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNodeIp

`func (o *AddCorosyncNodeRequestContent) SetNewNodeIp(v string)`

SetNewNodeIp sets NewNodeIp field to given value.

### HasNewNodeIp

`func (o *AddCorosyncNodeRequestContent) HasNewNodeIp() bool`

HasNewNodeIp returns a boolean if a field has been set.

### GetNodeid

`func (o *AddCorosyncNodeRequestContent) GetNodeid() float32`

GetNodeid returns the Nodeid field if non-nil, zero value otherwise.

### GetNodeidOk

`func (o *AddCorosyncNodeRequestContent) GetNodeidOk() (*float32, bool)`

GetNodeidOk returns a tuple with the Nodeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeid

`func (o *AddCorosyncNodeRequestContent) SetNodeid(v float32)`

SetNodeid sets Nodeid field to given value.

### HasNodeid

`func (o *AddCorosyncNodeRequestContent) HasNodeid() bool`

HasNodeid returns a boolean if a field has been set.

### GetVotes

`func (o *AddCorosyncNodeRequestContent) GetVotes() float32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *AddCorosyncNodeRequestContent) GetVotesOk() (*float32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *AddCorosyncNodeRequestContent) SetVotes(v float32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *AddCorosyncNodeRequestContent) HasVotes() bool`

HasVotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


