# CreateRealmRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Realm** | **string** |  | 
**Type** | [**RealmType**](RealmType.md) |  | 
**AcrValues** | Pointer to **string** |  | [optional] 
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
**UsernameClaim** | Pointer to **string** |  | [optional] 
**Verify** | Pointer to **float32** | An integer used to represent a boolean. 0 is false, 1 is true. | [optional] 

## Methods

### NewCreateRealmRequestContent

`func NewCreateRealmRequestContent(realm string, type_ RealmType, ) *CreateRealmRequestContent`

NewCreateRealmRequestContent instantiates a new CreateRealmRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRealmRequestContentWithDefaults

`func NewCreateRealmRequestContentWithDefaults() *CreateRealmRequestContent`

NewCreateRealmRequestContentWithDefaults instantiates a new CreateRealmRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRealm

`func (o *CreateRealmRequestContent) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *CreateRealmRequestContent) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *CreateRealmRequestContent) SetRealm(v string)`

SetRealm sets Realm field to given value.


### GetType

`func (o *CreateRealmRequestContent) GetType() RealmType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRealmRequestContent) GetTypeOk() (*RealmType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRealmRequestContent) SetType(v RealmType)`

SetType sets Type field to given value.


### GetAcrValues

`func (o *CreateRealmRequestContent) GetAcrValues() string`

GetAcrValues returns the AcrValues field if non-nil, zero value otherwise.

### GetAcrValuesOk

`func (o *CreateRealmRequestContent) GetAcrValuesOk() (*string, bool)`

GetAcrValuesOk returns a tuple with the AcrValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrValues

`func (o *CreateRealmRequestContent) SetAcrValues(v string)`

SetAcrValues sets AcrValues field to given value.

### HasAcrValues

`func (o *CreateRealmRequestContent) HasAcrValues() bool`

HasAcrValues returns a boolean if a field has been set.

### GetAutocreate

`func (o *CreateRealmRequestContent) GetAutocreate() float32`

GetAutocreate returns the Autocreate field if non-nil, zero value otherwise.

### GetAutocreateOk

`func (o *CreateRealmRequestContent) GetAutocreateOk() (*float32, bool)`

GetAutocreateOk returns a tuple with the Autocreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocreate

`func (o *CreateRealmRequestContent) SetAutocreate(v float32)`

SetAutocreate sets Autocreate field to given value.

### HasAutocreate

`func (o *CreateRealmRequestContent) HasAutocreate() bool`

HasAutocreate returns a boolean if a field has been set.

### GetBaseDn

`func (o *CreateRealmRequestContent) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *CreateRealmRequestContent) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *CreateRealmRequestContent) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *CreateRealmRequestContent) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetBindDn

`func (o *CreateRealmRequestContent) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *CreateRealmRequestContent) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *CreateRealmRequestContent) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *CreateRealmRequestContent) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetCapath

`func (o *CreateRealmRequestContent) GetCapath() string`

GetCapath returns the Capath field if non-nil, zero value otherwise.

### GetCapathOk

`func (o *CreateRealmRequestContent) GetCapathOk() (*string, bool)`

GetCapathOk returns a tuple with the Capath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapath

`func (o *CreateRealmRequestContent) SetCapath(v string)`

SetCapath sets Capath field to given value.

### HasCapath

`func (o *CreateRealmRequestContent) HasCapath() bool`

HasCapath returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *CreateRealmRequestContent) GetCaseSensitive() float32`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *CreateRealmRequestContent) GetCaseSensitiveOk() (*float32, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *CreateRealmRequestContent) SetCaseSensitive(v float32)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *CreateRealmRequestContent) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetCert

`func (o *CreateRealmRequestContent) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *CreateRealmRequestContent) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *CreateRealmRequestContent) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *CreateRealmRequestContent) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetCertkey

`func (o *CreateRealmRequestContent) GetCertkey() string`

GetCertkey returns the Certkey field if non-nil, zero value otherwise.

### GetCertkeyOk

`func (o *CreateRealmRequestContent) GetCertkeyOk() (*string, bool)`

GetCertkeyOk returns a tuple with the Certkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertkey

`func (o *CreateRealmRequestContent) SetCertkey(v string)`

SetCertkey sets Certkey field to given value.

### HasCertkey

`func (o *CreateRealmRequestContent) HasCertkey() bool`

HasCertkey returns a boolean if a field has been set.

### GetClientId

`func (o *CreateRealmRequestContent) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateRealmRequestContent) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateRealmRequestContent) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CreateRealmRequestContent) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientKey

`func (o *CreateRealmRequestContent) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *CreateRealmRequestContent) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *CreateRealmRequestContent) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *CreateRealmRequestContent) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### GetComment

`func (o *CreateRealmRequestContent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateRealmRequestContent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateRealmRequestContent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateRealmRequestContent) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefault

`func (o *CreateRealmRequestContent) GetDefault() float32`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CreateRealmRequestContent) GetDefaultOk() (*float32, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CreateRealmRequestContent) SetDefault(v float32)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CreateRealmRequestContent) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDomain

`func (o *CreateRealmRequestContent) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CreateRealmRequestContent) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CreateRealmRequestContent) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CreateRealmRequestContent) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFilter

`func (o *CreateRealmRequestContent) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CreateRealmRequestContent) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CreateRealmRequestContent) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CreateRealmRequestContent) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupClasses

`func (o *CreateRealmRequestContent) GetGroupClasses() string`

GetGroupClasses returns the GroupClasses field if non-nil, zero value otherwise.

### GetGroupClassesOk

`func (o *CreateRealmRequestContent) GetGroupClassesOk() (*string, bool)`

GetGroupClassesOk returns a tuple with the GroupClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClasses

`func (o *CreateRealmRequestContent) SetGroupClasses(v string)`

SetGroupClasses sets GroupClasses field to given value.

### HasGroupClasses

`func (o *CreateRealmRequestContent) HasGroupClasses() bool`

HasGroupClasses returns a boolean if a field has been set.

### GetGroupFilter

`func (o *CreateRealmRequestContent) GetGroupFilter() string`

GetGroupFilter returns the GroupFilter field if non-nil, zero value otherwise.

### GetGroupFilterOk

`func (o *CreateRealmRequestContent) GetGroupFilterOk() (*string, bool)`

GetGroupFilterOk returns a tuple with the GroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFilter

`func (o *CreateRealmRequestContent) SetGroupFilter(v string)`

SetGroupFilter sets GroupFilter field to given value.

### HasGroupFilter

`func (o *CreateRealmRequestContent) HasGroupFilter() bool`

HasGroupFilter returns a boolean if a field has been set.

### GetGroupDn

`func (o *CreateRealmRequestContent) GetGroupDn() string`

GetGroupDn returns the GroupDn field if non-nil, zero value otherwise.

### GetGroupDnOk

`func (o *CreateRealmRequestContent) GetGroupDnOk() (*string, bool)`

GetGroupDnOk returns a tuple with the GroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDn

`func (o *CreateRealmRequestContent) SetGroupDn(v string)`

SetGroupDn sets GroupDn field to given value.

### HasGroupDn

`func (o *CreateRealmRequestContent) HasGroupDn() bool`

HasGroupDn returns a boolean if a field has been set.

### GetGroupNameAttr

`func (o *CreateRealmRequestContent) GetGroupNameAttr() string`

GetGroupNameAttr returns the GroupNameAttr field if non-nil, zero value otherwise.

### GetGroupNameAttrOk

`func (o *CreateRealmRequestContent) GetGroupNameAttrOk() (*string, bool)`

GetGroupNameAttrOk returns a tuple with the GroupNameAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNameAttr

`func (o *CreateRealmRequestContent) SetGroupNameAttr(v string)`

SetGroupNameAttr sets GroupNameAttr field to given value.

### HasGroupNameAttr

`func (o *CreateRealmRequestContent) HasGroupNameAttr() bool`

HasGroupNameAttr returns a boolean if a field has been set.

### GetIssuerUrl

`func (o *CreateRealmRequestContent) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *CreateRealmRequestContent) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *CreateRealmRequestContent) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *CreateRealmRequestContent) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### GetMode

`func (o *CreateRealmRequestContent) GetMode() RealmMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateRealmRequestContent) GetModeOk() (*RealmMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateRealmRequestContent) SetMode(v RealmMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreateRealmRequestContent) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPassword

`func (o *CreateRealmRequestContent) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateRealmRequestContent) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateRealmRequestContent) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateRealmRequestContent) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *CreateRealmRequestContent) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateRealmRequestContent) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateRealmRequestContent) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateRealmRequestContent) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrompt

`func (o *CreateRealmRequestContent) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *CreateRealmRequestContent) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *CreateRealmRequestContent) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *CreateRealmRequestContent) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetScopes

`func (o *CreateRealmRequestContent) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateRealmRequestContent) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateRealmRequestContent) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CreateRealmRequestContent) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSecure

`func (o *CreateRealmRequestContent) GetSecure() float32`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *CreateRealmRequestContent) GetSecureOk() (*float32, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *CreateRealmRequestContent) SetSecure(v float32)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *CreateRealmRequestContent) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetServer1

`func (o *CreateRealmRequestContent) GetServer1() string`

GetServer1 returns the Server1 field if non-nil, zero value otherwise.

### GetServer1Ok

`func (o *CreateRealmRequestContent) GetServer1Ok() (*string, bool)`

GetServer1Ok returns a tuple with the Server1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer1

`func (o *CreateRealmRequestContent) SetServer1(v string)`

SetServer1 sets Server1 field to given value.

### HasServer1

`func (o *CreateRealmRequestContent) HasServer1() bool`

HasServer1 returns a boolean if a field has been set.

### GetServer2

`func (o *CreateRealmRequestContent) GetServer2() string`

GetServer2 returns the Server2 field if non-nil, zero value otherwise.

### GetServer2Ok

`func (o *CreateRealmRequestContent) GetServer2Ok() (*string, bool)`

GetServer2Ok returns a tuple with the Server2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer2

`func (o *CreateRealmRequestContent) SetServer2(v string)`

SetServer2 sets Server2 field to given value.

### HasServer2

`func (o *CreateRealmRequestContent) HasServer2() bool`

HasServer2 returns a boolean if a field has been set.

### GetSslversion

`func (o *CreateRealmRequestContent) GetSslversion() RealmSslVersion`

GetSslversion returns the Sslversion field if non-nil, zero value otherwise.

### GetSslversionOk

`func (o *CreateRealmRequestContent) GetSslversionOk() (*RealmSslVersion, bool)`

GetSslversionOk returns a tuple with the Sslversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslversion

`func (o *CreateRealmRequestContent) SetSslversion(v RealmSslVersion)`

SetSslversion sets Sslversion field to given value.

### HasSslversion

`func (o *CreateRealmRequestContent) HasSslversion() bool`

HasSslversion returns a boolean if a field has been set.

### GetSyncDefaultsOptions

`func (o *CreateRealmRequestContent) GetSyncDefaultsOptions() string`

GetSyncDefaultsOptions returns the SyncDefaultsOptions field if non-nil, zero value otherwise.

### GetSyncDefaultsOptionsOk

`func (o *CreateRealmRequestContent) GetSyncDefaultsOptionsOk() (*string, bool)`

GetSyncDefaultsOptionsOk returns a tuple with the SyncDefaultsOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDefaultsOptions

`func (o *CreateRealmRequestContent) SetSyncDefaultsOptions(v string)`

SetSyncDefaultsOptions sets SyncDefaultsOptions field to given value.

### HasSyncDefaultsOptions

`func (o *CreateRealmRequestContent) HasSyncDefaultsOptions() bool`

HasSyncDefaultsOptions returns a boolean if a field has been set.

### GetSyncAttributes

`func (o *CreateRealmRequestContent) GetSyncAttributes() string`

GetSyncAttributes returns the SyncAttributes field if non-nil, zero value otherwise.

### GetSyncAttributesOk

`func (o *CreateRealmRequestContent) GetSyncAttributesOk() (*string, bool)`

GetSyncAttributesOk returns a tuple with the SyncAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncAttributes

`func (o *CreateRealmRequestContent) SetSyncAttributes(v string)`

SetSyncAttributes sets SyncAttributes field to given value.

### HasSyncAttributes

`func (o *CreateRealmRequestContent) HasSyncAttributes() bool`

HasSyncAttributes returns a boolean if a field has been set.

### GetTfa

`func (o *CreateRealmRequestContent) GetTfa() string`

GetTfa returns the Tfa field if non-nil, zero value otherwise.

### GetTfaOk

`func (o *CreateRealmRequestContent) GetTfaOk() (*string, bool)`

GetTfaOk returns a tuple with the Tfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfa

`func (o *CreateRealmRequestContent) SetTfa(v string)`

SetTfa sets Tfa field to given value.

### HasTfa

`func (o *CreateRealmRequestContent) HasTfa() bool`

HasTfa returns a boolean if a field has been set.

### GetUserAttr

`func (o *CreateRealmRequestContent) GetUserAttr() string`

GetUserAttr returns the UserAttr field if non-nil, zero value otherwise.

### GetUserAttrOk

`func (o *CreateRealmRequestContent) GetUserAttrOk() (*string, bool)`

GetUserAttrOk returns a tuple with the UserAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttr

`func (o *CreateRealmRequestContent) SetUserAttr(v string)`

SetUserAttr sets UserAttr field to given value.

### HasUserAttr

`func (o *CreateRealmRequestContent) HasUserAttr() bool`

HasUserAttr returns a boolean if a field has been set.

### GetUserClasses

`func (o *CreateRealmRequestContent) GetUserClasses() string`

GetUserClasses returns the UserClasses field if non-nil, zero value otherwise.

### GetUserClassesOk

`func (o *CreateRealmRequestContent) GetUserClassesOk() (*string, bool)`

GetUserClassesOk returns a tuple with the UserClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClasses

`func (o *CreateRealmRequestContent) SetUserClasses(v string)`

SetUserClasses sets UserClasses field to given value.

### HasUserClasses

`func (o *CreateRealmRequestContent) HasUserClasses() bool`

HasUserClasses returns a boolean if a field has been set.

### GetUsernameClaim

`func (o *CreateRealmRequestContent) GetUsernameClaim() string`

GetUsernameClaim returns the UsernameClaim field if non-nil, zero value otherwise.

### GetUsernameClaimOk

`func (o *CreateRealmRequestContent) GetUsernameClaimOk() (*string, bool)`

GetUsernameClaimOk returns a tuple with the UsernameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameClaim

`func (o *CreateRealmRequestContent) SetUsernameClaim(v string)`

SetUsernameClaim sets UsernameClaim field to given value.

### HasUsernameClaim

`func (o *CreateRealmRequestContent) HasUsernameClaim() bool`

HasUsernameClaim returns a boolean if a field has been set.

### GetVerify

`func (o *CreateRealmRequestContent) GetVerify() float32`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *CreateRealmRequestContent) GetVerifyOk() (*float32, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *CreateRealmRequestContent) SetVerify(v float32)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *CreateRealmRequestContent) HasVerify() bool`

HasVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


