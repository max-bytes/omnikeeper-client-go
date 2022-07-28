# AttributeValueDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AttributeValueType**](AttributeValueType.md) |  | [optional] 
**IsArray** | Pointer to **bool** |  | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAttributeValueDTO

`func NewAttributeValueDTO() *AttributeValueDTO`

NewAttributeValueDTO instantiates a new AttributeValueDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeValueDTOWithDefaults

`func NewAttributeValueDTOWithDefaults() *AttributeValueDTO`

NewAttributeValueDTOWithDefaults instantiates a new AttributeValueDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AttributeValueDTO) GetType() AttributeValueType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttributeValueDTO) GetTypeOk() (*AttributeValueType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttributeValueDTO) SetType(v AttributeValueType)`

SetType sets Type field to given value.

### HasType

`func (o *AttributeValueDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsArray

`func (o *AttributeValueDTO) GetIsArray() bool`

GetIsArray returns the IsArray field if non-nil, zero value otherwise.

### GetIsArrayOk

`func (o *AttributeValueDTO) GetIsArrayOk() (*bool, bool)`

GetIsArrayOk returns a tuple with the IsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArray

`func (o *AttributeValueDTO) SetIsArray(v bool)`

SetIsArray sets IsArray field to given value.

### HasIsArray

`func (o *AttributeValueDTO) HasIsArray() bool`

HasIsArray returns a boolean if a field has been set.

### GetValues

`func (o *AttributeValueDTO) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AttributeValueDTO) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AttributeValueDTO) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *AttributeValueDTO) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


