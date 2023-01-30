# AddCustomNodeCertificateRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | **string** | PEM encoded certificate. | 
**Force** | Pointer to **float32** | Overwrite existing custom or ACME certificate. | [optional] 
**Key** | Pointer to **string** | PEM encoded private key. | [optional] 
**Restart** | Pointer to **float32** | Restart the PVE proxy. | [optional] 

## Methods

### NewAddCustomNodeCertificateRequestContent

`func NewAddCustomNodeCertificateRequestContent(certificates string, ) *AddCustomNodeCertificateRequestContent`

NewAddCustomNodeCertificateRequestContent instantiates a new AddCustomNodeCertificateRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCustomNodeCertificateRequestContentWithDefaults

`func NewAddCustomNodeCertificateRequestContentWithDefaults() *AddCustomNodeCertificateRequestContent`

NewAddCustomNodeCertificateRequestContentWithDefaults instantiates a new AddCustomNodeCertificateRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *AddCustomNodeCertificateRequestContent) GetCertificates() string`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *AddCustomNodeCertificateRequestContent) GetCertificatesOk() (*string, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *AddCustomNodeCertificateRequestContent) SetCertificates(v string)`

SetCertificates sets Certificates field to given value.


### GetForce

`func (o *AddCustomNodeCertificateRequestContent) GetForce() float32`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *AddCustomNodeCertificateRequestContent) GetForceOk() (*float32, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *AddCustomNodeCertificateRequestContent) SetForce(v float32)`

SetForce sets Force field to given value.

### HasForce

`func (o *AddCustomNodeCertificateRequestContent) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetKey

`func (o *AddCustomNodeCertificateRequestContent) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AddCustomNodeCertificateRequestContent) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AddCustomNodeCertificateRequestContent) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AddCustomNodeCertificateRequestContent) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRestart

`func (o *AddCustomNodeCertificateRequestContent) GetRestart() float32`

GetRestart returns the Restart field if non-nil, zero value otherwise.

### GetRestartOk

`func (o *AddCustomNodeCertificateRequestContent) GetRestartOk() (*float32, bool)`

GetRestartOk returns a tuple with the Restart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestart

`func (o *AddCustomNodeCertificateRequestContent) SetRestart(v float32)`

SetRestart sets Restart field to given value.

### HasRestart

`func (o *AddCustomNodeCertificateRequestContent) HasRestart() bool`

HasRestart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


