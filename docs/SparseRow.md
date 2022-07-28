# SparseRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciid** | Pointer to **string** |  | [optional] 
**Cells** | Pointer to [**[]ChangeDataCell**](ChangeDataCell.md) |  | [optional] 

## Methods

### NewSparseRow

`func NewSparseRow() *SparseRow`

NewSparseRow instantiates a new SparseRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSparseRowWithDefaults

`func NewSparseRowWithDefaults() *SparseRow`

NewSparseRowWithDefaults instantiates a new SparseRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiid

`func (o *SparseRow) GetCiid() string`

GetCiid returns the Ciid field if non-nil, zero value otherwise.

### GetCiidOk

`func (o *SparseRow) GetCiidOk() (*string, bool)`

GetCiidOk returns a tuple with the Ciid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiid

`func (o *SparseRow) SetCiid(v string)`

SetCiid sets Ciid field to given value.

### HasCiid

`func (o *SparseRow) HasCiid() bool`

HasCiid returns a boolean if a field has been set.

### GetCells

`func (o *SparseRow) GetCells() []ChangeDataCell`

GetCells returns the Cells field if non-nil, zero value otherwise.

### GetCellsOk

`func (o *SparseRow) GetCellsOk() (*[]ChangeDataCell, bool)`

GetCellsOk returns a tuple with the Cells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCells

`func (o *SparseRow) SetCells(v []ChangeDataCell)`

SetCells sets Cells field to given value.

### HasCells

`func (o *SparseRow) HasCells() bool`

HasCells returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


