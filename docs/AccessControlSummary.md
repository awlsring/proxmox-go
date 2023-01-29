# AccessControlSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**Roleid** | **string** | The role id | 
**Ugid** | **string** | The user group id | 
**Type** | [**AccessControlType**](AccessControlType.md) |  | 
**Propagate** | **float32** | Allow permission propegation | 

## Methods

### NewAccessControlSummary

`func NewAccessControlSummary(path string, roleid string, ugid string, type_ AccessControlType, propagate float32, ) *AccessControlSummary`

NewAccessControlSummary instantiates a new AccessControlSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlSummaryWithDefaults

`func NewAccessControlSummaryWithDefaults() *AccessControlSummary`

NewAccessControlSummaryWithDefaults instantiates a new AccessControlSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *AccessControlSummary) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AccessControlSummary) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AccessControlSummary) SetPath(v string)`

SetPath sets Path field to given value.


### GetRoleid

`func (o *AccessControlSummary) GetRoleid() string`

GetRoleid returns the Roleid field if non-nil, zero value otherwise.

### GetRoleidOk

`func (o *AccessControlSummary) GetRoleidOk() (*string, bool)`

GetRoleidOk returns a tuple with the Roleid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleid

`func (o *AccessControlSummary) SetRoleid(v string)`

SetRoleid sets Roleid field to given value.


### GetUgid

`func (o *AccessControlSummary) GetUgid() string`

GetUgid returns the Ugid field if non-nil, zero value otherwise.

### GetUgidOk

`func (o *AccessControlSummary) GetUgidOk() (*string, bool)`

GetUgidOk returns a tuple with the Ugid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUgid

`func (o *AccessControlSummary) SetUgid(v string)`

SetUgid sets Ugid field to given value.


### GetType

`func (o *AccessControlSummary) GetType() AccessControlType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessControlSummary) GetTypeOk() (*AccessControlType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessControlSummary) SetType(v AccessControlType)`

SetType sets Type field to given value.


### GetPropagate

`func (o *AccessControlSummary) GetPropagate() float32`

GetPropagate returns the Propagate field if non-nil, zero value otherwise.

### GetPropagateOk

`func (o *AccessControlSummary) GetPropagateOk() (*float32, bool)`

GetPropagateOk returns a tuple with the Propagate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagate

`func (o *AccessControlSummary) SetPropagate(v float32)`

SetPropagate sets Propagate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


