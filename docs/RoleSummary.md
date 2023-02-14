# RoleSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roleid** | **string** |  | 
**Privs** | Pointer to **string** |  | [optional] 
**Special** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 

## Methods

### NewRoleSummary

`func NewRoleSummary(roleid string, ) *RoleSummary`

NewRoleSummary instantiates a new RoleSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleSummaryWithDefaults

`func NewRoleSummaryWithDefaults() *RoleSummary`

NewRoleSummaryWithDefaults instantiates a new RoleSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleid

`func (o *RoleSummary) GetRoleid() string`

GetRoleid returns the Roleid field if non-nil, zero value otherwise.

### GetRoleidOk

`func (o *RoleSummary) GetRoleidOk() (*string, bool)`

GetRoleidOk returns a tuple with the Roleid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleid

`func (o *RoleSummary) SetRoleid(v string)`

SetRoleid sets Roleid field to given value.


### GetPrivs

`func (o *RoleSummary) GetPrivs() string`

GetPrivs returns the Privs field if non-nil, zero value otherwise.

### GetPrivsOk

`func (o *RoleSummary) GetPrivsOk() (*string, bool)`

GetPrivsOk returns a tuple with the Privs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivs

`func (o *RoleSummary) SetPrivs(v string)`

SetPrivs sets Privs field to given value.

### HasPrivs

`func (o *RoleSummary) HasPrivs() bool`

HasPrivs returns a boolean if a field has been set.

### GetSpecial

`func (o *RoleSummary) GetSpecial() float32`

GetSpecial returns the Special field if non-nil, zero value otherwise.

### GetSpecialOk

`func (o *RoleSummary) GetSpecialOk() (*float32, bool)`

GetSpecialOk returns a tuple with the Special field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecial

`func (o *RoleSummary) SetSpecial(v float32)`

SetSpecial sets Special field to given value.

### HasSpecial

`func (o *RoleSummary) HasSpecial() bool`

HasSpecial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


