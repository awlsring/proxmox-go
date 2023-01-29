# Ticket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Cap** | Pointer to **map[string]map[string]float32** |  | [optional] 
**CSRFPreventionToken** | Pointer to **string** |  | [optional] 
**Ticket** | Pointer to **string** |  | [optional] 
**Clustername** | Pointer to **string** |  | [optional] 

## Methods

### NewTicket

`func NewTicket(username string, ) *Ticket`

NewTicket instantiates a new Ticket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketWithDefaults

`func NewTicketWithDefaults() *Ticket`

NewTicketWithDefaults instantiates a new Ticket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *Ticket) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Ticket) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Ticket) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetCap

`func (o *Ticket) GetCap() map[string]map[string]float32`

GetCap returns the Cap field if non-nil, zero value otherwise.

### GetCapOk

`func (o *Ticket) GetCapOk() (*map[string]map[string]float32, bool)`

GetCapOk returns a tuple with the Cap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCap

`func (o *Ticket) SetCap(v map[string]map[string]float32)`

SetCap sets Cap field to given value.

### HasCap

`func (o *Ticket) HasCap() bool`

HasCap returns a boolean if a field has been set.

### GetCSRFPreventionToken

`func (o *Ticket) GetCSRFPreventionToken() string`

GetCSRFPreventionToken returns the CSRFPreventionToken field if non-nil, zero value otherwise.

### GetCSRFPreventionTokenOk

`func (o *Ticket) GetCSRFPreventionTokenOk() (*string, bool)`

GetCSRFPreventionTokenOk returns a tuple with the CSRFPreventionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSRFPreventionToken

`func (o *Ticket) SetCSRFPreventionToken(v string)`

SetCSRFPreventionToken sets CSRFPreventionToken field to given value.

### HasCSRFPreventionToken

`func (o *Ticket) HasCSRFPreventionToken() bool`

HasCSRFPreventionToken returns a boolean if a field has been set.

### GetTicket

`func (o *Ticket) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *Ticket) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *Ticket) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *Ticket) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetClustername

`func (o *Ticket) GetClustername() string`

GetClustername returns the Clustername field if non-nil, zero value otherwise.

### GetClusternameOk

`func (o *Ticket) GetClusternameOk() (*string, bool)`

GetClusternameOk returns a tuple with the Clustername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClustername

`func (o *Ticket) SetClustername(v string)`

SetClustername sets Clustername field to given value.

### HasClustername

`func (o *Ticket) HasClustername() bool`

HasClustername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


