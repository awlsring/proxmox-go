# InitCephClusterRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterNetwork** | Pointer to **string** |  | [optional] 
**DisableCephx** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**MinSize** | Pointer to **float32** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**PgBits** | Pointer to **float32** |  | [optional] 
**Size** | Pointer to **float32** |  | [optional] 

## Methods

### NewInitCephClusterRequestContent

`func NewInitCephClusterRequestContent() *InitCephClusterRequestContent`

NewInitCephClusterRequestContent instantiates a new InitCephClusterRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitCephClusterRequestContentWithDefaults

`func NewInitCephClusterRequestContentWithDefaults() *InitCephClusterRequestContent`

NewInitCephClusterRequestContentWithDefaults instantiates a new InitCephClusterRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterNetwork

`func (o *InitCephClusterRequestContent) GetClusterNetwork() string`

GetClusterNetwork returns the ClusterNetwork field if non-nil, zero value otherwise.

### GetClusterNetworkOk

`func (o *InitCephClusterRequestContent) GetClusterNetworkOk() (*string, bool)`

GetClusterNetworkOk returns a tuple with the ClusterNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNetwork

`func (o *InitCephClusterRequestContent) SetClusterNetwork(v string)`

SetClusterNetwork sets ClusterNetwork field to given value.

### HasClusterNetwork

`func (o *InitCephClusterRequestContent) HasClusterNetwork() bool`

HasClusterNetwork returns a boolean if a field has been set.

### GetDisableCephx

`func (o *InitCephClusterRequestContent) GetDisableCephx() float32`

GetDisableCephx returns the DisableCephx field if non-nil, zero value otherwise.

### GetDisableCephxOk

`func (o *InitCephClusterRequestContent) GetDisableCephxOk() (*float32, bool)`

GetDisableCephxOk returns a tuple with the DisableCephx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCephx

`func (o *InitCephClusterRequestContent) SetDisableCephx(v float32)`

SetDisableCephx sets DisableCephx field to given value.

### HasDisableCephx

`func (o *InitCephClusterRequestContent) HasDisableCephx() bool`

HasDisableCephx returns a boolean if a field has been set.

### GetMinSize

`func (o *InitCephClusterRequestContent) GetMinSize() float32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *InitCephClusterRequestContent) GetMinSizeOk() (*float32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *InitCephClusterRequestContent) SetMinSize(v float32)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *InitCephClusterRequestContent) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.

### GetNetwork

`func (o *InitCephClusterRequestContent) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InitCephClusterRequestContent) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InitCephClusterRequestContent) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InitCephClusterRequestContent) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPgBits

`func (o *InitCephClusterRequestContent) GetPgBits() float32`

GetPgBits returns the PgBits field if non-nil, zero value otherwise.

### GetPgBitsOk

`func (o *InitCephClusterRequestContent) GetPgBitsOk() (*float32, bool)`

GetPgBitsOk returns a tuple with the PgBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgBits

`func (o *InitCephClusterRequestContent) SetPgBits(v float32)`

SetPgBits sets PgBits field to given value.

### HasPgBits

`func (o *InitCephClusterRequestContent) HasPgBits() bool`

HasPgBits returns a boolean if a field has been set.

### GetSize

`func (o *InitCephClusterRequestContent) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *InitCephClusterRequestContent) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *InitCephClusterRequestContent) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *InitCephClusterRequestContent) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


