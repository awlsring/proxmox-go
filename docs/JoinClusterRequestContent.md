# JoinClusterRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | **string** | The fingerprint of the cluster certificate | 
**Hostname** | **string** | The hostname or IP address an existing cluster node | 
**Password** | **string** | Password for the node root user | 
**Force** | Pointer to **bool** |  | [optional] 
**Link0** | Pointer to **string** |  | [optional] 
**Link1** | Pointer to **string** |  | [optional] 
**Link2** | Pointer to **string** |  | [optional] 
**Link3** | Pointer to **string** |  | [optional] 
**Link4** | Pointer to **string** |  | [optional] 
**Link5** | Pointer to **string** |  | [optional] 
**Link6** | Pointer to **string** |  | [optional] 
**Link7** | Pointer to **string** |  | [optional] 
**Nodeid** | Pointer to **float32** | The node ID to use for the joining node | [optional] 
**Votes** | Pointer to **float32** | Votes for the node | [optional] 

## Methods

### NewJoinClusterRequestContent

`func NewJoinClusterRequestContent(fingerprint string, hostname string, password string, ) *JoinClusterRequestContent`

NewJoinClusterRequestContent instantiates a new JoinClusterRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinClusterRequestContentWithDefaults

`func NewJoinClusterRequestContentWithDefaults() *JoinClusterRequestContent`

NewJoinClusterRequestContentWithDefaults instantiates a new JoinClusterRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *JoinClusterRequestContent) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *JoinClusterRequestContent) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *JoinClusterRequestContent) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetHostname

`func (o *JoinClusterRequestContent) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *JoinClusterRequestContent) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *JoinClusterRequestContent) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPassword

`func (o *JoinClusterRequestContent) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *JoinClusterRequestContent) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *JoinClusterRequestContent) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetForce

`func (o *JoinClusterRequestContent) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *JoinClusterRequestContent) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *JoinClusterRequestContent) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *JoinClusterRequestContent) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetLink0

`func (o *JoinClusterRequestContent) GetLink0() string`

GetLink0 returns the Link0 field if non-nil, zero value otherwise.

### GetLink0Ok

`func (o *JoinClusterRequestContent) GetLink0Ok() (*string, bool)`

GetLink0Ok returns a tuple with the Link0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink0

`func (o *JoinClusterRequestContent) SetLink0(v string)`

SetLink0 sets Link0 field to given value.

### HasLink0

`func (o *JoinClusterRequestContent) HasLink0() bool`

HasLink0 returns a boolean if a field has been set.

### GetLink1

`func (o *JoinClusterRequestContent) GetLink1() string`

GetLink1 returns the Link1 field if non-nil, zero value otherwise.

### GetLink1Ok

`func (o *JoinClusterRequestContent) GetLink1Ok() (*string, bool)`

GetLink1Ok returns a tuple with the Link1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink1

`func (o *JoinClusterRequestContent) SetLink1(v string)`

SetLink1 sets Link1 field to given value.

### HasLink1

`func (o *JoinClusterRequestContent) HasLink1() bool`

HasLink1 returns a boolean if a field has been set.

### GetLink2

`func (o *JoinClusterRequestContent) GetLink2() string`

GetLink2 returns the Link2 field if non-nil, zero value otherwise.

### GetLink2Ok

`func (o *JoinClusterRequestContent) GetLink2Ok() (*string, bool)`

GetLink2Ok returns a tuple with the Link2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink2

`func (o *JoinClusterRequestContent) SetLink2(v string)`

SetLink2 sets Link2 field to given value.

### HasLink2

`func (o *JoinClusterRequestContent) HasLink2() bool`

HasLink2 returns a boolean if a field has been set.

### GetLink3

`func (o *JoinClusterRequestContent) GetLink3() string`

GetLink3 returns the Link3 field if non-nil, zero value otherwise.

### GetLink3Ok

`func (o *JoinClusterRequestContent) GetLink3Ok() (*string, bool)`

GetLink3Ok returns a tuple with the Link3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink3

`func (o *JoinClusterRequestContent) SetLink3(v string)`

SetLink3 sets Link3 field to given value.

### HasLink3

`func (o *JoinClusterRequestContent) HasLink3() bool`

HasLink3 returns a boolean if a field has been set.

### GetLink4

`func (o *JoinClusterRequestContent) GetLink4() string`

GetLink4 returns the Link4 field if non-nil, zero value otherwise.

### GetLink4Ok

`func (o *JoinClusterRequestContent) GetLink4Ok() (*string, bool)`

GetLink4Ok returns a tuple with the Link4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink4

`func (o *JoinClusterRequestContent) SetLink4(v string)`

SetLink4 sets Link4 field to given value.

### HasLink4

`func (o *JoinClusterRequestContent) HasLink4() bool`

HasLink4 returns a boolean if a field has been set.

### GetLink5

`func (o *JoinClusterRequestContent) GetLink5() string`

GetLink5 returns the Link5 field if non-nil, zero value otherwise.

### GetLink5Ok

`func (o *JoinClusterRequestContent) GetLink5Ok() (*string, bool)`

GetLink5Ok returns a tuple with the Link5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink5

`func (o *JoinClusterRequestContent) SetLink5(v string)`

SetLink5 sets Link5 field to given value.

### HasLink5

`func (o *JoinClusterRequestContent) HasLink5() bool`

HasLink5 returns a boolean if a field has been set.

### GetLink6

`func (o *JoinClusterRequestContent) GetLink6() string`

GetLink6 returns the Link6 field if non-nil, zero value otherwise.

### GetLink6Ok

`func (o *JoinClusterRequestContent) GetLink6Ok() (*string, bool)`

GetLink6Ok returns a tuple with the Link6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink6

`func (o *JoinClusterRequestContent) SetLink6(v string)`

SetLink6 sets Link6 field to given value.

### HasLink6

`func (o *JoinClusterRequestContent) HasLink6() bool`

HasLink6 returns a boolean if a field has been set.

### GetLink7

`func (o *JoinClusterRequestContent) GetLink7() string`

GetLink7 returns the Link7 field if non-nil, zero value otherwise.

### GetLink7Ok

`func (o *JoinClusterRequestContent) GetLink7Ok() (*string, bool)`

GetLink7Ok returns a tuple with the Link7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink7

`func (o *JoinClusterRequestContent) SetLink7(v string)`

SetLink7 sets Link7 field to given value.

### HasLink7

`func (o *JoinClusterRequestContent) HasLink7() bool`

HasLink7 returns a boolean if a field has been set.

### GetNodeid

`func (o *JoinClusterRequestContent) GetNodeid() float32`

GetNodeid returns the Nodeid field if non-nil, zero value otherwise.

### GetNodeidOk

`func (o *JoinClusterRequestContent) GetNodeidOk() (*float32, bool)`

GetNodeidOk returns a tuple with the Nodeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeid

`func (o *JoinClusterRequestContent) SetNodeid(v float32)`

SetNodeid sets Nodeid field to given value.

### HasNodeid

`func (o *JoinClusterRequestContent) HasNodeid() bool`

HasNodeid returns a boolean if a field has been set.

### GetVotes

`func (o *JoinClusterRequestContent) GetVotes() float32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *JoinClusterRequestContent) GetVotesOk() (*float32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *JoinClusterRequestContent) SetVotes(v float32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *JoinClusterRequestContent) HasVotes() bool`

HasVotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


