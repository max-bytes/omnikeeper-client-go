# GridViewConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowCIIDColumn** | Pointer to **bool** |  | [optional] 
**WriteLayer** | Pointer to **string** |  | [optional] 
**ReadLayerset** | Pointer to **[]string** |  | [optional] 
**Columns** | Pointer to [**[]GridViewColumn**](GridViewColumn.md) |  | [optional] 
**Trait** | Pointer to **string** |  | [optional] 

## Methods

### NewGridViewConfiguration

`func NewGridViewConfiguration() *GridViewConfiguration`

NewGridViewConfiguration instantiates a new GridViewConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridViewConfigurationWithDefaults

`func NewGridViewConfigurationWithDefaults() *GridViewConfiguration`

NewGridViewConfigurationWithDefaults instantiates a new GridViewConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowCIIDColumn

`func (o *GridViewConfiguration) GetShowCIIDColumn() bool`

GetShowCIIDColumn returns the ShowCIIDColumn field if non-nil, zero value otherwise.

### GetShowCIIDColumnOk

`func (o *GridViewConfiguration) GetShowCIIDColumnOk() (*bool, bool)`

GetShowCIIDColumnOk returns a tuple with the ShowCIIDColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCIIDColumn

`func (o *GridViewConfiguration) SetShowCIIDColumn(v bool)`

SetShowCIIDColumn sets ShowCIIDColumn field to given value.

### HasShowCIIDColumn

`func (o *GridViewConfiguration) HasShowCIIDColumn() bool`

HasShowCIIDColumn returns a boolean if a field has been set.

### GetWriteLayer

`func (o *GridViewConfiguration) GetWriteLayer() string`

GetWriteLayer returns the WriteLayer field if non-nil, zero value otherwise.

### GetWriteLayerOk

`func (o *GridViewConfiguration) GetWriteLayerOk() (*string, bool)`

GetWriteLayerOk returns a tuple with the WriteLayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteLayer

`func (o *GridViewConfiguration) SetWriteLayer(v string)`

SetWriteLayer sets WriteLayer field to given value.

### HasWriteLayer

`func (o *GridViewConfiguration) HasWriteLayer() bool`

HasWriteLayer returns a boolean if a field has been set.

### GetReadLayerset

`func (o *GridViewConfiguration) GetReadLayerset() []string`

GetReadLayerset returns the ReadLayerset field if non-nil, zero value otherwise.

### GetReadLayersetOk

`func (o *GridViewConfiguration) GetReadLayersetOk() (*[]string, bool)`

GetReadLayersetOk returns a tuple with the ReadLayerset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadLayerset

`func (o *GridViewConfiguration) SetReadLayerset(v []string)`

SetReadLayerset sets ReadLayerset field to given value.

### HasReadLayerset

`func (o *GridViewConfiguration) HasReadLayerset() bool`

HasReadLayerset returns a boolean if a field has been set.

### GetColumns

`func (o *GridViewConfiguration) GetColumns() []GridViewColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *GridViewConfiguration) GetColumnsOk() (*[]GridViewColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *GridViewConfiguration) SetColumns(v []GridViewColumn)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *GridViewConfiguration) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetTrait

`func (o *GridViewConfiguration) GetTrait() string`

GetTrait returns the Trait field if non-nil, zero value otherwise.

### GetTraitOk

`func (o *GridViewConfiguration) GetTraitOk() (*string, bool)`

GetTraitOk returns a tuple with the Trait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrait

`func (o *GridViewConfiguration) SetTrait(v string)`

SetTrait sets Trait field to given value.

### HasTrait

`func (o *GridViewConfiguration) HasTrait() bool`

HasTrait returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


