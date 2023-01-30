# NodeCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** | The certificate&#39;s filename. | [optional] 
**Fingerprint** | Pointer to **string** | The certificate&#39;s fingerprint. | [optional] 
**Issuer** | Pointer to **string** | The certificate&#39;s issuer. | [optional] 
**NotAfter** | Pointer to **float32** | The certificate&#39;s notAfter timestamp. | [optional] 
**NotBefore** | Pointer to **float32** | The certificate&#39;s notAfter timestamp. | [optional] 
**Pem** | Pointer to **string** |  | [optional] 
**PublicKeyBits** | Pointer to **string** | The certificate&#39;s key size. | [optional] 
**PublicKeyType** | Pointer to **string** | The certificate&#39;s key algorithm. | [optional] 
**Sans** | Pointer to **[]string** | The certificate&#39;s sans | [optional] 
**Subject** | Pointer to **string** | The certificate&#39;s subject. | [optional] 

## Methods

### NewNodeCertificate

`func NewNodeCertificate() *NodeCertificate`

NewNodeCertificate instantiates a new NodeCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeCertificateWithDefaults

`func NewNodeCertificateWithDefaults() *NodeCertificate`

NewNodeCertificateWithDefaults instantiates a new NodeCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *NodeCertificate) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *NodeCertificate) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *NodeCertificate) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *NodeCertificate) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFingerprint

`func (o *NodeCertificate) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *NodeCertificate) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *NodeCertificate) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *NodeCertificate) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetIssuer

`func (o *NodeCertificate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *NodeCertificate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *NodeCertificate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *NodeCertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetNotAfter

`func (o *NodeCertificate) GetNotAfter() float32`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *NodeCertificate) GetNotAfterOk() (*float32, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *NodeCertificate) SetNotAfter(v float32)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *NodeCertificate) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *NodeCertificate) GetNotBefore() float32`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *NodeCertificate) GetNotBeforeOk() (*float32, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *NodeCertificate) SetNotBefore(v float32)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *NodeCertificate) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetPem

`func (o *NodeCertificate) GetPem() string`

GetPem returns the Pem field if non-nil, zero value otherwise.

### GetPemOk

`func (o *NodeCertificate) GetPemOk() (*string, bool)`

GetPemOk returns a tuple with the Pem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPem

`func (o *NodeCertificate) SetPem(v string)`

SetPem sets Pem field to given value.

### HasPem

`func (o *NodeCertificate) HasPem() bool`

HasPem returns a boolean if a field has been set.

### GetPublicKeyBits

`func (o *NodeCertificate) GetPublicKeyBits() string`

GetPublicKeyBits returns the PublicKeyBits field if non-nil, zero value otherwise.

### GetPublicKeyBitsOk

`func (o *NodeCertificate) GetPublicKeyBitsOk() (*string, bool)`

GetPublicKeyBitsOk returns a tuple with the PublicKeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyBits

`func (o *NodeCertificate) SetPublicKeyBits(v string)`

SetPublicKeyBits sets PublicKeyBits field to given value.

### HasPublicKeyBits

`func (o *NodeCertificate) HasPublicKeyBits() bool`

HasPublicKeyBits returns a boolean if a field has been set.

### GetPublicKeyType

`func (o *NodeCertificate) GetPublicKeyType() string`

GetPublicKeyType returns the PublicKeyType field if non-nil, zero value otherwise.

### GetPublicKeyTypeOk

`func (o *NodeCertificate) GetPublicKeyTypeOk() (*string, bool)`

GetPublicKeyTypeOk returns a tuple with the PublicKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyType

`func (o *NodeCertificate) SetPublicKeyType(v string)`

SetPublicKeyType sets PublicKeyType field to given value.

### HasPublicKeyType

`func (o *NodeCertificate) HasPublicKeyType() bool`

HasPublicKeyType returns a boolean if a field has been set.

### GetSans

`func (o *NodeCertificate) GetSans() []string`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *NodeCertificate) GetSansOk() (*[]string, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *NodeCertificate) SetSans(v []string)`

SetSans sets Sans field to given value.

### HasSans

`func (o *NodeCertificate) HasSans() bool`

HasSans returns a boolean if a field has been set.

### GetSubject

`func (o *NodeCertificate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NodeCertificate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NodeCertificate) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *NodeCertificate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


