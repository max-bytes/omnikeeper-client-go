# GridViewColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceAttributeName** | Pointer to **NullableString** |  | [optional] 
**ColumnDescription** | Pointer to **NullableString** |  | [optional] 
**ValueType** | Pointer to [**AttributeValueType**](AttributeValueType.md) |  | [optional] 
**WriteLayer** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewGridViewColumn

`func NewGridViewColumn() *GridViewColumn`

NewGridViewColumn instantiates a new GridViewColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridViewColumnWithDefaults

`func NewGridViewColumnWithDefaults() *GridViewColumn`

NewGridViewColumnWithDefaults instantiates a new GridViewColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceAttributeName

`func (o *GridViewColumn) GetSourceAttributeName() string`

GetSourceAttributeName returns the SourceAttributeName field if non-nil, zero value otherwise.

### GetSourceAttributeNameOk

`func (o *GridViewColumn) GetSourceAttributeNameOk() (*string, bool)`

GetSourceAttributeNameOk returns a tuple with the SourceAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttributeName

`func (o *GridViewColumn) SetSourceAttributeName(v string)`

SetSourceAttributeName sets SourceAttributeName field to given value.

### HasSourceAttributeName

`func (o *GridViewColumn) HasSourceAttributeName() bool`

HasSourceAttributeName returns a boolean if a field has been set.

### SetSourceAttributeNameNil

`func (o *GridViewColumn) SetSourceAttributeNameNil(b bool)`

 SetSourceAttributeNameNil sets the value for SourceAttributeName to be an explicit nil

### UnsetSourceAttributeName
`func (o *GridViewColumn) UnsetSourceAttributeName()`

UnsetSourceAttributeName ensures that no value is present for SourceAttributeName, not even an explicit nil
### GetColumnDescription

`func (o *GridViewColumn) GetColumnDescription() string`

GetColumnDescription returns the ColumnDescription field if non-nil, zero value otherwise.

### GetColumnDescriptionOk

`func (o *GridViewColumn) GetColumnDescriptionOk() (*string, bool)`

GetColumnDescriptionOk returns a tuple with the ColumnDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnDescription

`func (o *GridViewColumn) SetColumnDescription(v string)`

SetColumnDescription sets ColumnDescription field to given value.

### HasColumnDescription

`func (o *GridViewColumn) HasColumnDescription() bool`

HasColumnDescription returns a boolean if a field has been set.

### SetColumnDescriptionNil

`func (o *GridViewColumn) SetColumnDescriptionNil(b bool)`

 SetColumnDescriptionNil sets the value for ColumnDescription to be an explicit nil

### UnsetColumnDescription
`func (o *GridViewColumn) UnsetColumnDescription()`

UnsetColumnDescription ensures that no value is present for ColumnDescription, not even an explicit nil
### GetValueType

`func (o *GridViewColumn) GetValueType() AttributeValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *GridViewColumn) GetValueTypeOk() (*AttributeValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *GridViewColumn) SetValueType(v AttributeValueType)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *GridViewColumn) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetWriteLayer

`func (o *GridViewColumn) GetWriteLayer() int64`

GetWriteLayer returns the WriteLayer field if non-nil, zero value otherwise.

### GetWriteLayerOk

`func (o *GridViewColumn) GetWriteLayerOk() (*int64, bool)`

GetWriteLayerOk returns a tuple with the WriteLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLayer

`func (o *GridViewColumn) SetWriteLayer(v int64)`

SetWriteLayer sets WriteLayer field to given value.

### HasWriteLayer

`func (o *GridViewColumn) HasWriteLayer() bool`

HasWriteLayer returns a boolean if a field has been set.

### SetWriteLayerNil

`func (o *GridViewColumn) SetWriteLayerNil(b bool)`

 SetWriteLayerNil sets the value for WriteLayer to be an explicit nil

### UnsetWriteLayer
`func (o *GridViewColumn) UnsetWriteLayer()`

UnsetWriteLayer ensures that no value is present for WriteLayer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


