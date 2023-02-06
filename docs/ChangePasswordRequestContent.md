# ChangePasswordRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Userid** | **string** |  | 
**Password** | **string** |  | 

## Methods

### NewChangePasswordRequestContent

`func NewChangePasswordRequestContent(userid string, password string, ) *ChangePasswordRequestContent`

NewChangePasswordRequestContent instantiates a new ChangePasswordRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePasswordRequestContentWithDefaults

`func NewChangePasswordRequestContentWithDefaults() *ChangePasswordRequestContent`

NewChangePasswordRequestContentWithDefaults instantiates a new ChangePasswordRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserid

`func (o *ChangePasswordRequestContent) GetUserid() string`

GetUserid returns the Userid field if non-nil, zero value otherwise.

### GetUseridOk

`func (o *ChangePasswordRequestContent) GetUseridOk() (*string, bool)`

GetUseridOk returns a tuple with the Userid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserid

`func (o *ChangePasswordRequestContent) SetUserid(v string)`

SetUserid sets Userid field to given value.


### GetPassword

`func (o *ChangePasswordRequestContent) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ChangePasswordRequestContent) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ChangePasswordRequestContent) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


