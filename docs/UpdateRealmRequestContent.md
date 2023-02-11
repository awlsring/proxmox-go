# UpdateRealmRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autocreate** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**BaseDn** | Pointer to **string** |  | [optional] 
**BindDn** | Pointer to **string** |  | [optional] 
**Capath** | Pointer to **string** |  | [optional] 
**CaseSensitive** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Cert** | Pointer to **string** |  | [optional] 
**Certkey** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientKey** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Delete** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Digest** | Pointer to **string** |  | [optional] 
**Default** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to **string** |  | [optional] 
**GroupClasses** | Pointer to **string** |  | [optional] 
**GroupFilter** | Pointer to **string** |  | [optional] 
**GroupDn** | Pointer to **string** |  | [optional] 
**GroupNameAttr** | Pointer to **string** |  | [optional] 
**IssuerUrl** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**RealmMode**](RealmMode.md) |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **float32** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **string** |  | [optional] 
**Secure** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 
**Server1** | Pointer to **string** |  | [optional] 
**Server2** | Pointer to **string** |  | [optional] 
**Sslversion** | Pointer to [**RealmSslVersion**](RealmSslVersion.md) |  | [optional] 
**SyncDefaultsOptions** | Pointer to **string** |  | [optional] 
**SyncAttributes** | Pointer to **string** |  | [optional] 
**Tfa** | Pointer to **string** |  | [optional] 
**UserAttr** | Pointer to **string** |  | [optional] 
**UserClasses** | Pointer to **string** |  | [optional] 
**Verify** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 

## Methods

### NewUpdateRealmRequestContent

`func NewUpdateRealmRequestContent() *UpdateRealmRequestContent`

NewUpdateRealmRequestContent instantiates a new UpdateRealmRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRealmRequestContentWithDefaults

`func NewUpdateRealmRequestContentWithDefaults() *UpdateRealmRequestContent`

NewUpdateRealmRequestContentWithDefaults instantiates a new UpdateRealmRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutocreate

`func (o *UpdateRealmRequestContent) GetAutocreate() float32`

GetAutocreate returns the Autocreate field if non-nil, zero value otherwise.

### GetAutocreateOk

`func (o *UpdateRealmRequestContent) GetAutocreateOk() (*float32, bool)`

GetAutocreateOk returns a tuple with the Autocreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocreate

`func (o *UpdateRealmRequestContent) SetAutocreate(v float32)`

SetAutocreate sets Autocreate field to given value.

### HasAutocreate

`func (o *UpdateRealmRequestContent) HasAutocreate() bool`

HasAutocreate returns a boolean if a field has been set.

### GetBaseDn

`func (o *UpdateRealmRequestContent) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *UpdateRealmRequestContent) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *UpdateRealmRequestContent) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *UpdateRealmRequestContent) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetBindDn

`func (o *UpdateRealmRequestContent) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *UpdateRealmRequestContent) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *UpdateRealmRequestContent) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *UpdateRealmRequestContent) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetCapath

`func (o *UpdateRealmRequestContent) GetCapath() string`

GetCapath returns the Capath field if non-nil, zero value otherwise.

### GetCapathOk

`func (o *UpdateRealmRequestContent) GetCapathOk() (*string, bool)`

GetCapathOk returns a tuple with the Capath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapath

`func (o *UpdateRealmRequestContent) SetCapath(v string)`

SetCapath sets Capath field to given value.

### HasCapath

`func (o *UpdateRealmRequestContent) HasCapath() bool`

HasCapath returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *UpdateRealmRequestContent) GetCaseSensitive() float32`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *UpdateRealmRequestContent) GetCaseSensitiveOk() (*float32, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *UpdateRealmRequestContent) SetCaseSensitive(v float32)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *UpdateRealmRequestContent) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetCert

`func (o *UpdateRealmRequestContent) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *UpdateRealmRequestContent) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *UpdateRealmRequestContent) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *UpdateRealmRequestContent) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetCertkey

`func (o *UpdateRealmRequestContent) GetCertkey() string`

GetCertkey returns the Certkey field if non-nil, zero value otherwise.

### GetCertkeyOk

`func (o *UpdateRealmRequestContent) GetCertkeyOk() (*string, bool)`

GetCertkeyOk returns a tuple with the Certkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertkey

`func (o *UpdateRealmRequestContent) SetCertkey(v string)`

SetCertkey sets Certkey field to given value.

### HasCertkey

`func (o *UpdateRealmRequestContent) HasCertkey() bool`

HasCertkey returns a boolean if a field has been set.

### GetClientId

`func (o *UpdateRealmRequestContent) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UpdateRealmRequestContent) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UpdateRealmRequestContent) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UpdateRealmRequestContent) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientKey

`func (o *UpdateRealmRequestContent) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *UpdateRealmRequestContent) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *UpdateRealmRequestContent) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *UpdateRealmRequestContent) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### GetComment

`func (o *UpdateRealmRequestContent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateRealmRequestContent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateRealmRequestContent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateRealmRequestContent) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDelete

`func (o *UpdateRealmRequestContent) GetDelete() float32`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *UpdateRealmRequestContent) GetDeleteOk() (*float32, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *UpdateRealmRequestContent) SetDelete(v float32)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *UpdateRealmRequestContent) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDigest

`func (o *UpdateRealmRequestContent) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *UpdateRealmRequestContent) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *UpdateRealmRequestContent) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *UpdateRealmRequestContent) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetDefault

`func (o *UpdateRealmRequestContent) GetDefault() float32`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *UpdateRealmRequestContent) GetDefaultOk() (*float32, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *UpdateRealmRequestContent) SetDefault(v float32)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *UpdateRealmRequestContent) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDomain

`func (o *UpdateRealmRequestContent) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UpdateRealmRequestContent) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UpdateRealmRequestContent) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UpdateRealmRequestContent) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFilter

`func (o *UpdateRealmRequestContent) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UpdateRealmRequestContent) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UpdateRealmRequestContent) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UpdateRealmRequestContent) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupClasses

`func (o *UpdateRealmRequestContent) GetGroupClasses() string`

GetGroupClasses returns the GroupClasses field if non-nil, zero value otherwise.

### GetGroupClassesOk

`func (o *UpdateRealmRequestContent) GetGroupClassesOk() (*string, bool)`

GetGroupClassesOk returns a tuple with the GroupClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClasses

`func (o *UpdateRealmRequestContent) SetGroupClasses(v string)`

SetGroupClasses sets GroupClasses field to given value.

### HasGroupClasses

`func (o *UpdateRealmRequestContent) HasGroupClasses() bool`

HasGroupClasses returns a boolean if a field has been set.

### GetGroupFilter

`func (o *UpdateRealmRequestContent) GetGroupFilter() string`

GetGroupFilter returns the GroupFilter field if non-nil, zero value otherwise.

### GetGroupFilterOk

`func (o *UpdateRealmRequestContent) GetGroupFilterOk() (*string, bool)`

GetGroupFilterOk returns a tuple with the GroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFilter

`func (o *UpdateRealmRequestContent) SetGroupFilter(v string)`

SetGroupFilter sets GroupFilter field to given value.

### HasGroupFilter

`func (o *UpdateRealmRequestContent) HasGroupFilter() bool`

HasGroupFilter returns a boolean if a field has been set.

### GetGroupDn

`func (o *UpdateRealmRequestContent) GetGroupDn() string`

GetGroupDn returns the GroupDn field if non-nil, zero value otherwise.

### GetGroupDnOk

`func (o *UpdateRealmRequestContent) GetGroupDnOk() (*string, bool)`

GetGroupDnOk returns a tuple with the GroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDn

`func (o *UpdateRealmRequestContent) SetGroupDn(v string)`

SetGroupDn sets GroupDn field to given value.

### HasGroupDn

`func (o *UpdateRealmRequestContent) HasGroupDn() bool`

HasGroupDn returns a boolean if a field has been set.

### GetGroupNameAttr

`func (o *UpdateRealmRequestContent) GetGroupNameAttr() string`

GetGroupNameAttr returns the GroupNameAttr field if non-nil, zero value otherwise.

### GetGroupNameAttrOk

`func (o *UpdateRealmRequestContent) GetGroupNameAttrOk() (*string, bool)`

GetGroupNameAttrOk returns a tuple with the GroupNameAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNameAttr

`func (o *UpdateRealmRequestContent) SetGroupNameAttr(v string)`

SetGroupNameAttr sets GroupNameAttr field to given value.

### HasGroupNameAttr

`func (o *UpdateRealmRequestContent) HasGroupNameAttr() bool`

HasGroupNameAttr returns a boolean if a field has been set.

### GetIssuerUrl

`func (o *UpdateRealmRequestContent) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *UpdateRealmRequestContent) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *UpdateRealmRequestContent) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *UpdateRealmRequestContent) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### GetMode

`func (o *UpdateRealmRequestContent) GetMode() RealmMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UpdateRealmRequestContent) GetModeOk() (*RealmMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UpdateRealmRequestContent) SetMode(v RealmMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UpdateRealmRequestContent) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateRealmRequestContent) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateRealmRequestContent) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateRealmRequestContent) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateRealmRequestContent) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *UpdateRealmRequestContent) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateRealmRequestContent) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateRealmRequestContent) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateRealmRequestContent) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrompt

`func (o *UpdateRealmRequestContent) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *UpdateRealmRequestContent) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *UpdateRealmRequestContent) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *UpdateRealmRequestContent) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetScopes

`func (o *UpdateRealmRequestContent) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UpdateRealmRequestContent) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UpdateRealmRequestContent) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UpdateRealmRequestContent) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSecure

`func (o *UpdateRealmRequestContent) GetSecure() float32`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *UpdateRealmRequestContent) GetSecureOk() (*float32, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *UpdateRealmRequestContent) SetSecure(v float32)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *UpdateRealmRequestContent) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetServer1

`func (o *UpdateRealmRequestContent) GetServer1() string`

GetServer1 returns the Server1 field if non-nil, zero value otherwise.

### GetServer1Ok

`func (o *UpdateRealmRequestContent) GetServer1Ok() (*string, bool)`

GetServer1Ok returns a tuple with the Server1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer1

`func (o *UpdateRealmRequestContent) SetServer1(v string)`

SetServer1 sets Server1 field to given value.

### HasServer1

`func (o *UpdateRealmRequestContent) HasServer1() bool`

HasServer1 returns a boolean if a field has been set.

### GetServer2

`func (o *UpdateRealmRequestContent) GetServer2() string`

GetServer2 returns the Server2 field if non-nil, zero value otherwise.

### GetServer2Ok

`func (o *UpdateRealmRequestContent) GetServer2Ok() (*string, bool)`

GetServer2Ok returns a tuple with the Server2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer2

`func (o *UpdateRealmRequestContent) SetServer2(v string)`

SetServer2 sets Server2 field to given value.

### HasServer2

`func (o *UpdateRealmRequestContent) HasServer2() bool`

HasServer2 returns a boolean if a field has been set.

### GetSslversion

`func (o *UpdateRealmRequestContent) GetSslversion() RealmSslVersion`

GetSslversion returns the Sslversion field if non-nil, zero value otherwise.

### GetSslversionOk

`func (o *UpdateRealmRequestContent) GetSslversionOk() (*RealmSslVersion, bool)`

GetSslversionOk returns a tuple with the Sslversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslversion

`func (o *UpdateRealmRequestContent) SetSslversion(v RealmSslVersion)`

SetSslversion sets Sslversion field to given value.

### HasSslversion

`func (o *UpdateRealmRequestContent) HasSslversion() bool`

HasSslversion returns a boolean if a field has been set.

### GetSyncDefaultsOptions

`func (o *UpdateRealmRequestContent) GetSyncDefaultsOptions() string`

GetSyncDefaultsOptions returns the SyncDefaultsOptions field if non-nil, zero value otherwise.

### GetSyncDefaultsOptionsOk

`func (o *UpdateRealmRequestContent) GetSyncDefaultsOptionsOk() (*string, bool)`

GetSyncDefaultsOptionsOk returns a tuple with the SyncDefaultsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDefaultsOptions

`func (o *UpdateRealmRequestContent) SetSyncDefaultsOptions(v string)`

SetSyncDefaultsOptions sets SyncDefaultsOptions field to given value.

### HasSyncDefaultsOptions

`func (o *UpdateRealmRequestContent) HasSyncDefaultsOptions() bool`

HasSyncDefaultsOptions returns a boolean if a field has been set.

### GetSyncAttributes

`func (o *UpdateRealmRequestContent) GetSyncAttributes() string`

GetSyncAttributes returns the SyncAttributes field if non-nil, zero value otherwise.

### GetSyncAttributesOk

`func (o *UpdateRealmRequestContent) GetSyncAttributesOk() (*string, bool)`

GetSyncAttributesOk returns a tuple with the SyncAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncAttributes

`func (o *UpdateRealmRequestContent) SetSyncAttributes(v string)`

SetSyncAttributes sets SyncAttributes field to given value.

### HasSyncAttributes

`func (o *UpdateRealmRequestContent) HasSyncAttributes() bool`

HasSyncAttributes returns a boolean if a field has been set.

### GetTfa

`func (o *UpdateRealmRequestContent) GetTfa() string`

GetTfa returns the Tfa field if non-nil, zero value otherwise.

### GetTfaOk

`func (o *UpdateRealmRequestContent) GetTfaOk() (*string, bool)`

GetTfaOk returns a tuple with the Tfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfa

`func (o *UpdateRealmRequestContent) SetTfa(v string)`

SetTfa sets Tfa field to given value.

### HasTfa

`func (o *UpdateRealmRequestContent) HasTfa() bool`

HasTfa returns a boolean if a field has been set.

### GetUserAttr

`func (o *UpdateRealmRequestContent) GetUserAttr() string`

GetUserAttr returns the UserAttr field if non-nil, zero value otherwise.

### GetUserAttrOk

`func (o *UpdateRealmRequestContent) GetUserAttrOk() (*string, bool)`

GetUserAttrOk returns a tuple with the UserAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttr

`func (o *UpdateRealmRequestContent) SetUserAttr(v string)`

SetUserAttr sets UserAttr field to given value.

### HasUserAttr

`func (o *UpdateRealmRequestContent) HasUserAttr() bool`

HasUserAttr returns a boolean if a field has been set.

### GetUserClasses

`func (o *UpdateRealmRequestContent) GetUserClasses() string`

GetUserClasses returns the UserClasses field if non-nil, zero value otherwise.

### GetUserClassesOk

`func (o *UpdateRealmRequestContent) GetUserClassesOk() (*string, bool)`

GetUserClassesOk returns a tuple with the UserClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClasses

`func (o *UpdateRealmRequestContent) SetUserClasses(v string)`

SetUserClasses sets UserClasses field to given value.

### HasUserClasses

`func (o *UpdateRealmRequestContent) HasUserClasses() bool`

HasUserClasses returns a boolean if a field has been set.

### GetVerify

`func (o *UpdateRealmRequestContent) GetVerify() float32`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *UpdateRealmRequestContent) GetVerifyOk() (*float32, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *UpdateRealmRequestContent) SetVerify(v float32)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *UpdateRealmRequestContent) HasVerify() bool`

HasVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


