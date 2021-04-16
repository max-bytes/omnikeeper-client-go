# ChangeDataCell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
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

### GetName

`func (o *ChangeDataCell) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChangeDataCell) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChangeDataCell) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChangeDataCell) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ChangeDataCell) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ChangeDataCell) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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


