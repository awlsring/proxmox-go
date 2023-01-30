# CreateZFSPoolRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | **string** | The devices to create the zfs pool on. This is a comma seperated list sent as a string. | 
**Name** | **string** | The storage identifier. | 
**Raidlevel** | [**ZFSRaidLevel**](ZFSRaidLevel.md) |  | 
**AddStorage** | Pointer to **float32** | Configure storage using the directory. Takes a boolean integer value (0 false, 1 true). | [optional] 
**Ashift** | Pointer to **float32** | The pool vector size exponent. | [optional] 
**Compression** | Pointer to [**ZFSCompression**](ZFSCompression.md) |  | [optional] 
**DraidConfig** | Pointer to **string** | Draid config. Set as string like &#39;data&#x3D;&lt;int&gt;,spares&#x3D;&lt;int&gt; | [optional] 

## Methods

### NewCreateZFSPoolRequestContent

`func NewCreateZFSPoolRequestContent(devices string, name string, raidlevel ZFSRaidLevel, ) *CreateZFSPoolRequestContent`

NewCreateZFSPoolRequestContent instantiates a new CreateZFSPoolRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateZFSPoolRequestContentWithDefaults

`func NewCreateZFSPoolRequestContentWithDefaults() *CreateZFSPoolRequestContent`

NewCreateZFSPoolRequestContentWithDefaults instantiates a new CreateZFSPoolRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *CreateZFSPoolRequestContent) GetDevices() string`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *CreateZFSPoolRequestContent) GetDevicesOk() (*string, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *CreateZFSPoolRequestContent) SetDevices(v string)`

SetDevices sets Devices field to given value.


### GetName

`func (o *CreateZFSPoolRequestContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateZFSPoolRequestContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateZFSPoolRequestContent) SetName(v string)`

SetName sets Name field to given value.


### GetRaidlevel

`func (o *CreateZFSPoolRequestContent) GetRaidlevel() ZFSRaidLevel`

GetRaidlevel returns the Raidlevel field if non-nil, zero value otherwise.

### GetRaidlevelOk

`func (o *CreateZFSPoolRequestContent) GetRaidlevelOk() (*ZFSRaidLevel, bool)`

GetRaidlevelOk returns a tuple with the Raidlevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidlevel

`func (o *CreateZFSPoolRequestContent) SetRaidlevel(v ZFSRaidLevel)`

SetRaidlevel sets Raidlevel field to given value.


### GetAddStorage

`func (o *CreateZFSPoolRequestContent) GetAddStorage() float32`

GetAddStorage returns the AddStorage field if non-nil, zero value otherwise.

### GetAddStorageOk

`func (o *CreateZFSPoolRequestContent) GetAddStorageOk() (*float32, bool)`

GetAddStorageOk returns a tuple with the AddStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddStorage

`func (o *CreateZFSPoolRequestContent) SetAddStorage(v float32)`

SetAddStorage sets AddStorage field to given value.

### HasAddStorage

`func (o *CreateZFSPoolRequestContent) HasAddStorage() bool`

HasAddStorage returns a boolean if a field has been set.

### GetAshift

`func (o *CreateZFSPoolRequestContent) GetAshift() float32`

GetAshift returns the Ashift field if non-nil, zero value otherwise.

### GetAshiftOk

`func (o *CreateZFSPoolRequestContent) GetAshiftOk() (*float32, bool)`

GetAshiftOk returns a tuple with the Ashift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAshift

`func (o *CreateZFSPoolRequestContent) SetAshift(v float32)`

SetAshift sets Ashift field to given value.

### HasAshift

`func (o *CreateZFSPoolRequestContent) HasAshift() bool`

HasAshift returns a boolean if a field has been set.

### GetCompression

`func (o *CreateZFSPoolRequestContent) GetCompression() ZFSCompression`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *CreateZFSPoolRequestContent) GetCompressionOk() (*ZFSCompression, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *CreateZFSPoolRequestContent) SetCompression(v ZFSCompression)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *CreateZFSPoolRequestContent) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetDraidConfig

`func (o *CreateZFSPoolRequestContent) GetDraidConfig() string`

GetDraidConfig returns the DraidConfig field if non-nil, zero value otherwise.

### GetDraidConfigOk

`func (o *CreateZFSPoolRequestContent) GetDraidConfigOk() (*string, bool)`

GetDraidConfigOk returns a tuple with the DraidConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraidConfig

`func (o *CreateZFSPoolRequestContent) SetDraidConfig(v string)`

SetDraidConfig sets DraidConfig field to given value.

### HasDraidConfig

`func (o *CreateZFSPoolRequestContent) HasDraidConfig() bool`

HasDraidConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


