# NodeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Node** | **string** | The name of the node | 
**Maxmem** | Pointer to **float32** | Max memory allocated in bytes | [optional] 
**Mem** | Pointer to **float32** | Current memory utilization in bytes | [optional] 
**Disk** | Pointer to **float32** | Current disk utilization in bytes | [optional] 
**Maxdisk** | Pointer to **float32** | Max disk space available in bytes | [optional] 
**Maxcpu** | Pointer to **float32** | Amount of CPU cores available on node | [optional] 
**Id** | Pointer to **string** | The node id. This is a string of &#39;node/&lt;name&gt;&#39; | [optional] 
**Level** | Pointer to **string** | The node&#39;s support level | [optional] 
**SslFingerprint** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The system type. Seems to be always node | [optional] 
**Cpu** | Pointer to **float32** | The virtual machines cpu utilization in percent | [optional] 
**Uptime** | Pointer to **float32** | The node&#39;s uptime in seconds | [optional] 
**Status** | Pointer to [**NodeStatus**](NodeStatus.md) |  | [optional] 

## Methods

### NewNodeSummary

`func NewNodeSummary(node string, ) *NodeSummary`

NewNodeSummary instantiates a new NodeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeSummaryWithDefaults

`func NewNodeSummaryWithDefaults() *NodeSummary`

NewNodeSummaryWithDefaults instantiates a new NodeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNode

`func (o *NodeSummary) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *NodeSummary) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *NodeSummary) SetNode(v string)`

SetNode sets Node field to given value.


### GetMaxmem

`func (o *NodeSummary) GetMaxmem() float32`

GetMaxmem returns the Maxmem field if non-nil, zero value otherwise.

### GetMaxmemOk

`func (o *NodeSummary) GetMaxmemOk() (*float32, bool)`

GetMaxmemOk returns a tuple with the Maxmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxmem

`func (o *NodeSummary) SetMaxmem(v float32)`

SetMaxmem sets Maxmem field to given value.

### HasMaxmem

`func (o *NodeSummary) HasMaxmem() bool`

HasMaxmem returns a boolean if a field has been set.

### GetMem

`func (o *NodeSummary) GetMem() float32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *NodeSummary) GetMemOk() (*float32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *NodeSummary) SetMem(v float32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *NodeSummary) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetDisk

`func (o *NodeSummary) GetDisk() float32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *NodeSummary) GetDiskOk() (*float32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *NodeSummary) SetDisk(v float32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *NodeSummary) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetMaxdisk

`func (o *NodeSummary) GetMaxdisk() float32`

GetMaxdisk returns the Maxdisk field if non-nil, zero value otherwise.

### GetMaxdiskOk

`func (o *NodeSummary) GetMaxdiskOk() (*float32, bool)`

GetMaxdiskOk returns a tuple with the Maxdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxdisk

`func (o *NodeSummary) SetMaxdisk(v float32)`

SetMaxdisk sets Maxdisk field to given value.

### HasMaxdisk

`func (o *NodeSummary) HasMaxdisk() bool`

HasMaxdisk returns a boolean if a field has been set.

### GetMaxcpu

`func (o *NodeSummary) GetMaxcpu() float32`

GetMaxcpu returns the Maxcpu field if non-nil, zero value otherwise.

### GetMaxcpuOk

`func (o *NodeSummary) GetMaxcpuOk() (*float32, bool)`

GetMaxcpuOk returns a tuple with the Maxcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxcpu

`func (o *NodeSummary) SetMaxcpu(v float32)`

SetMaxcpu sets Maxcpu field to given value.

### HasMaxcpu

`func (o *NodeSummary) HasMaxcpu() bool`

HasMaxcpu returns a boolean if a field has been set.

### GetId

`func (o *NodeSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NodeSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *NodeSummary) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *NodeSummary) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *NodeSummary) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *NodeSummary) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetSslFingerprint

`func (o *NodeSummary) GetSslFingerprint() string`

GetSslFingerprint returns the SslFingerprint field if non-nil, zero value otherwise.

### GetSslFingerprintOk

`func (o *NodeSummary) GetSslFingerprintOk() (*string, bool)`

GetSslFingerprintOk returns a tuple with the SslFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslFingerprint

`func (o *NodeSummary) SetSslFingerprint(v string)`

SetSslFingerprint sets SslFingerprint field to given value.

### HasSslFingerprint

`func (o *NodeSummary) HasSslFingerprint() bool`

HasSslFingerprint returns a boolean if a field has been set.

### GetType

`func (o *NodeSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NodeSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCpu

`func (o *NodeSummary) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NodeSummary) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NodeSummary) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *NodeSummary) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetUptime

`func (o *NodeSummary) GetUptime() float32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *NodeSummary) GetUptimeOk() (*float32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *NodeSummary) SetUptime(v float32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *NodeSummary) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetStatus

`func (o *NodeSummary) GetStatus() NodeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeSummary) GetStatusOk() (*NodeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeSummary) SetStatus(v NodeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


