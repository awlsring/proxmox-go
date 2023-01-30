# UsbDeviceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Busnum** | **float32** |  | 
**Devnum** | **float32** |  | 
**Level** | **float32** |  | 
**Port** | **float32** |  | 
**Prodid** | **string** |  | 
**Vendid** | **string** |  | 
**Speed** | **string** |  | 
**Manufacturer** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**Usbpath** | Pointer to **string** |  | [optional] 

## Methods

### NewUsbDeviceSummary

`func NewUsbDeviceSummary(busnum float32, devnum float32, level float32, port float32, prodid string, vendid string, speed string, ) *UsbDeviceSummary`

NewUsbDeviceSummary instantiates a new UsbDeviceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsbDeviceSummaryWithDefaults

`func NewUsbDeviceSummaryWithDefaults() *UsbDeviceSummary`

NewUsbDeviceSummaryWithDefaults instantiates a new UsbDeviceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusnum

`func (o *UsbDeviceSummary) GetBusnum() float32`

GetBusnum returns the Busnum field if non-nil, zero value otherwise.

### GetBusnumOk

`func (o *UsbDeviceSummary) GetBusnumOk() (*float32, bool)`

GetBusnumOk returns a tuple with the Busnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusnum

`func (o *UsbDeviceSummary) SetBusnum(v float32)`

SetBusnum sets Busnum field to given value.


### GetDevnum

`func (o *UsbDeviceSummary) GetDevnum() float32`

GetDevnum returns the Devnum field if non-nil, zero value otherwise.

### GetDevnumOk

`func (o *UsbDeviceSummary) GetDevnumOk() (*float32, bool)`

GetDevnumOk returns a tuple with the Devnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevnum

`func (o *UsbDeviceSummary) SetDevnum(v float32)`

SetDevnum sets Devnum field to given value.


### GetLevel

`func (o *UsbDeviceSummary) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *UsbDeviceSummary) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *UsbDeviceSummary) SetLevel(v float32)`

SetLevel sets Level field to given value.


### GetPort

`func (o *UsbDeviceSummary) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UsbDeviceSummary) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UsbDeviceSummary) SetPort(v float32)`

SetPort sets Port field to given value.


### GetProdid

`func (o *UsbDeviceSummary) GetProdid() string`

GetProdid returns the Prodid field if non-nil, zero value otherwise.

### GetProdidOk

`func (o *UsbDeviceSummary) GetProdidOk() (*string, bool)`

GetProdidOk returns a tuple with the Prodid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProdid

`func (o *UsbDeviceSummary) SetProdid(v string)`

SetProdid sets Prodid field to given value.


### GetVendid

`func (o *UsbDeviceSummary) GetVendid() string`

GetVendid returns the Vendid field if non-nil, zero value otherwise.

### GetVendidOk

`func (o *UsbDeviceSummary) GetVendidOk() (*string, bool)`

GetVendidOk returns a tuple with the Vendid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendid

`func (o *UsbDeviceSummary) SetVendid(v string)`

SetVendid sets Vendid field to given value.


### GetSpeed

`func (o *UsbDeviceSummary) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *UsbDeviceSummary) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *UsbDeviceSummary) SetSpeed(v string)`

SetSpeed sets Speed field to given value.


### GetManufacturer

`func (o *UsbDeviceSummary) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *UsbDeviceSummary) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *UsbDeviceSummary) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *UsbDeviceSummary) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetProduct

`func (o *UsbDeviceSummary) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *UsbDeviceSummary) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *UsbDeviceSummary) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *UsbDeviceSummary) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSerial

`func (o *UsbDeviceSummary) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *UsbDeviceSummary) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *UsbDeviceSummary) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *UsbDeviceSummary) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetUsbpath

`func (o *UsbDeviceSummary) GetUsbpath() string`

GetUsbpath returns the Usbpath field if non-nil, zero value otherwise.

### GetUsbpathOk

`func (o *UsbDeviceSummary) GetUsbpathOk() (*string, bool)`

GetUsbpathOk returns a tuple with the Usbpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbpath

`func (o *UsbDeviceSummary) SetUsbpath(v string)`

SetUsbpath sets Usbpath field to given value.

### HasUsbpath

`func (o *UsbDeviceSummary) HasUsbpath() bool`

HasUsbpath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


