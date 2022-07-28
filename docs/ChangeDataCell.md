# ChangeDataCell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**AttributeValueDTO**](AttributeValueDTO.md) |  | [optional] 
**Changeable** | Pointer to **bool** |  | [optional] 

## Methods

### NewChangeDataCell

`func NewChangeDataCell() *ChangeDataCell`

NewChangeDataCell instantiates a new ChangeDataCell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeDataCellWithDefaults

`func NewChangeDataCellWithDefaults() *ChangeDataCell`

NewChangeDataCellWithDefaults instantiates a new ChangeDataCell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChangeDataCell) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeDataCell) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeDataCell) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChangeDataCell) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *ChangeDataCell) GetValue() AttributeValueDTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ChangeDataCell) GetValueOk() (*AttributeValueDTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ChangeDataCell) SetValue(v AttributeValueDTO)`

SetValue sets Value field to given value.

### HasValue

`func (o *ChangeDataCell) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetChangeable

`func (o *ChangeDataCell) GetChangeable() bool`

GetChangeable returns the Changeable field if non-nil, zero value otherwise.

### GetChangeableOk

`func (o *ChangeDataCell) GetChangeableOk() (*bool, bool)`

GetChangeableOk returns a tuple with the Changeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeable

`func (o *ChangeDataCell) SetChangeable(v bool)`

SetChangeable sets Changeable field to given value.

### HasChangeable

`func (o *ChangeDataCell) HasChangeable() bool`

HasChangeable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


