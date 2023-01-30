# ZFSPoolStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]ZFSPoolStatusChild**](ZFSPoolStatusChild.md) | The pool configuration including vdevs for each section. May be nested. | 
**Errors** | **string** | Errors on the pool | 
**Name** | **string** | The pool name | 
**State** | **string** | The state of the pool | 
**Action** | Pointer to **string** | Recommended action to address the pool state | [optional] 
**Scan** | Pointer to **string** | Information on the last or current scrub. | [optional] 
**Status** | Pointer to **string** | Information on the state of the pool | [optional] 

## Methods

### NewZFSPoolStatusSummary

`func NewZFSPoolStatusSummary(children []ZFSPoolStatusChild, errors string, name string, state string, ) *ZFSPoolStatusSummary`

NewZFSPoolStatusSummary instantiates a new ZFSPoolStatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZFSPoolStatusSummaryWithDefaults

`func NewZFSPoolStatusSummaryWithDefaults() *ZFSPoolStatusSummary`

NewZFSPoolStatusSummaryWithDefaults instantiates a new ZFSPoolStatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *ZFSPoolStatusSummary) GetChildren() []ZFSPoolStatusChild`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ZFSPoolStatusSummary) GetChildrenOk() (*[]ZFSPoolStatusChild, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ZFSPoolStatusSummary) SetChildren(v []ZFSPoolStatusChild)`

SetChildren sets Children field to given value.


### GetErrors

`func (o *ZFSPoolStatusSummary) GetErrors() string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ZFSPoolStatusSummary) GetErrorsOk() (*string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ZFSPoolStatusSummary) SetErrors(v string)`

SetErrors sets Errors field to given value.


### GetName

`func (o *ZFSPoolStatusSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZFSPoolStatusSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZFSPoolStatusSummary) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *ZFSPoolStatusSummary) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ZFSPoolStatusSummary) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ZFSPoolStatusSummary) SetState(v string)`

SetState sets State field to given value.


### GetAction

`func (o *ZFSPoolStatusSummary) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ZFSPoolStatusSummary) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ZFSPoolStatusSummary) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ZFSPoolStatusSummary) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetScan

`func (o *ZFSPoolStatusSummary) GetScan() string`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *ZFSPoolStatusSummary) GetScanOk() (*string, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *ZFSPoolStatusSummary) SetScan(v string)`

SetScan sets Scan field to given value.

### HasScan

`func (o *ZFSPoolStatusSummary) HasScan() bool`

HasScan returns a boolean if a field has been set.

### GetStatus

`func (o *ZFSPoolStatusSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ZFSPoolStatusSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ZFSPoolStatusSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ZFSPoolStatusSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


