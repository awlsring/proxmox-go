# CommandStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exited** | **bool** |  | 
**ErrData** | Pointer to **string** |  | [optional] 
**ErrTruncated** | Pointer to **bool** |  | [optional] 
**Exitcode** | Pointer to **float32** |  | [optional] 
**OutData** | Pointer to **string** |  | [optional] 
**OutTruncated** | Pointer to **bool** |  | [optional] 
**Signal** | Pointer to **float32** |  | [optional] 

## Methods

### NewCommandStatusSummary

`func NewCommandStatusSummary(exited bool, ) *CommandStatusSummary`

NewCommandStatusSummary instantiates a new CommandStatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandStatusSummaryWithDefaults

`func NewCommandStatusSummaryWithDefaults() *CommandStatusSummary`

NewCommandStatusSummaryWithDefaults instantiates a new CommandStatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExited

`func (o *CommandStatusSummary) GetExited() bool`

GetExited returns the Exited field if non-nil, zero value otherwise.

### GetExitedOk

`func (o *CommandStatusSummary) GetExitedOk() (*bool, bool)`

GetExitedOk returns a tuple with the Exited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExited

`func (o *CommandStatusSummary) SetExited(v bool)`

SetExited sets Exited field to given value.


### GetErrData

`func (o *CommandStatusSummary) GetErrData() string`

GetErrData returns the ErrData field if non-nil, zero value otherwise.

### GetErrDataOk

`func (o *CommandStatusSummary) GetErrDataOk() (*string, bool)`

GetErrDataOk returns a tuple with the ErrData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrData

`func (o *CommandStatusSummary) SetErrData(v string)`

SetErrData sets ErrData field to given value.

### HasErrData

`func (o *CommandStatusSummary) HasErrData() bool`

HasErrData returns a boolean if a field has been set.

### GetErrTruncated

`func (o *CommandStatusSummary) GetErrTruncated() bool`

GetErrTruncated returns the ErrTruncated field if non-nil, zero value otherwise.

### GetErrTruncatedOk

`func (o *CommandStatusSummary) GetErrTruncatedOk() (*bool, bool)`

GetErrTruncatedOk returns a tuple with the ErrTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrTruncated

`func (o *CommandStatusSummary) SetErrTruncated(v bool)`

SetErrTruncated sets ErrTruncated field to given value.

### HasErrTruncated

`func (o *CommandStatusSummary) HasErrTruncated() bool`

HasErrTruncated returns a boolean if a field has been set.

### GetExitcode

`func (o *CommandStatusSummary) GetExitcode() float32`

GetExitcode returns the Exitcode field if non-nil, zero value otherwise.

### GetExitcodeOk

`func (o *CommandStatusSummary) GetExitcodeOk() (*float32, bool)`

GetExitcodeOk returns a tuple with the Exitcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitcode

`func (o *CommandStatusSummary) SetExitcode(v float32)`

SetExitcode sets Exitcode field to given value.

### HasExitcode

`func (o *CommandStatusSummary) HasExitcode() bool`

HasExitcode returns a boolean if a field has been set.

### GetOutData

`func (o *CommandStatusSummary) GetOutData() string`

GetOutData returns the OutData field if non-nil, zero value otherwise.

### GetOutDataOk

`func (o *CommandStatusSummary) GetOutDataOk() (*string, bool)`

GetOutDataOk returns a tuple with the OutData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutData

`func (o *CommandStatusSummary) SetOutData(v string)`

SetOutData sets OutData field to given value.

### HasOutData

`func (o *CommandStatusSummary) HasOutData() bool`

HasOutData returns a boolean if a field has been set.

### GetOutTruncated

`func (o *CommandStatusSummary) GetOutTruncated() bool`

GetOutTruncated returns the OutTruncated field if non-nil, zero value otherwise.

### GetOutTruncatedOk

`func (o *CommandStatusSummary) GetOutTruncatedOk() (*bool, bool)`

GetOutTruncatedOk returns a tuple with the OutTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutTruncated

`func (o *CommandStatusSummary) SetOutTruncated(v bool)`

SetOutTruncated sets OutTruncated field to given value.

### HasOutTruncated

`func (o *CommandStatusSummary) HasOutTruncated() bool`

HasOutTruncated returns a boolean if a field has been set.

### GetSignal

`func (o *CommandStatusSummary) GetSignal() float32`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *CommandStatusSummary) GetSignalOk() (*float32, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *CommandStatusSummary) SetSignal(v float32)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *CommandStatusSummary) HasSignal() bool`

HasSignal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


