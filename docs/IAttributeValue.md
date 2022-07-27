# IAttributeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AttributeValueType**](AttributeValueType.md) |  | [optional] 
**IsArray** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewIAttributeValue

`func NewIAttributeValue() *IAttributeValue`

NewIAttributeValue instantiates a new IAttributeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIAttributeValueWithDefaults

`func NewIAttributeValueWithDefaults() *IAttributeValue`

NewIAttributeValueWithDefaults instantiates a new IAttributeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IAttributeValue) GetType() AttributeValueType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IAttributeValue) GetTypeOk() (*AttributeValueType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IAttributeValue) SetType(v AttributeValueType)`

SetType sets Type field to given value.

### HasType

`func (o *IAttributeValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsArray

`func (o *IAttributeValue) GetIsArray() bool`

GetIsArray returns the IsArray field if non-nil, zero value otherwise.

### GetIsArrayOk

`func (o *IAttributeValue) GetIsArrayOk() (*bool, bool)`

GetIsArrayOk returns a tuple with the IsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArray

`func (o *IAttributeValue) SetIsArray(v bool)`

SetIsArray sets IsArray field to given value.

### HasIsArray

`func (o *IAttributeValue) HasIsArray() bool`

HasIsArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


