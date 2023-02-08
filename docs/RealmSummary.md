# RealmSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Realm** | **string** |  | 
**Type** | **string** |  | 
**Comment** | Pointer to **string** |  | [optional] 
**Tfa** | Pointer to [**TFAType**](TFAType.md) |  | [optional] 

## Methods

### NewRealmSummary

`func NewRealmSummary(realm string, type_ string, ) *RealmSummary`

NewRealmSummary instantiates a new RealmSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmSummaryWithDefaults

`func NewRealmSummaryWithDefaults() *RealmSummary`

NewRealmSummaryWithDefaults instantiates a new RealmSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRealm

`func (o *RealmSummary) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *RealmSummary) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *RealmSummary) SetRealm(v string)`

SetRealm sets Realm field to given value.


### GetType

`func (o *RealmSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmSummary) SetType(v string)`

SetType sets Type field to given value.


### GetComment

`func (o *RealmSummary) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RealmSummary) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RealmSummary) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RealmSummary) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetTfa

`func (o *RealmSummary) GetTfa() TFAType`

GetTfa returns the Tfa field if non-nil, zero value otherwise.

### GetTfaOk

`func (o *RealmSummary) GetTfaOk() (*TFAType, bool)`

GetTfaOk returns a tuple with the Tfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfa

`func (o *RealmSummary) SetTfa(v TFAType)`

SetTfa sets Tfa field to given value.

### HasTfa

`func (o *RealmSummary) HasTfa() bool`

HasTfa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


