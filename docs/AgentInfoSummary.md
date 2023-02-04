# AgentInfoSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedCommands** | [**[]CommandSummary**](CommandSummary.md) |  | 
**Version** | **string** |  | 

## Methods

### NewAgentInfoSummary

`func NewAgentInfoSummary(supportedCommands []CommandSummary, version string, ) *AgentInfoSummary`

NewAgentInfoSummary instantiates a new AgentInfoSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentInfoSummaryWithDefaults

`func NewAgentInfoSummaryWithDefaults() *AgentInfoSummary`

NewAgentInfoSummaryWithDefaults instantiates a new AgentInfoSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedCommands

`func (o *AgentInfoSummary) GetSupportedCommands() []CommandSummary`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *AgentInfoSummary) GetSupportedCommandsOk() (*[]CommandSummary, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *AgentInfoSummary) SetSupportedCommands(v []CommandSummary)`

SetSupportedCommands sets SupportedCommands field to given value.


### GetVersion

`func (o *AgentInfoSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AgentInfoSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AgentInfoSummary) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


