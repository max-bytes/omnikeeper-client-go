# CIAttributeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Value** | [**AttributeValueDTO**](AttributeValueDTO.md) |  | 
**Ciid** | **string** |  | 
**State** | [**AttributeState**](AttributeState.md) |  | 

## Methods

### NewCIAttributeDTO

`func NewCIAttributeDTO(id string, name string, value AttributeValueDTO, ciid string, state AttributeState, ) *CIAttributeDTO`

NewCIAttributeDTO instantiates a new CIAttributeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCIAttributeDTOWithDefaults

`func NewCIAttributeDTOWithDefaults() *CIAttributeDTO`

NewCIAttributeDTOWithDefaults instantiates a new CIAttributeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CIAttributeDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CIAttributeDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CIAttributeDTO) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CIAttributeDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CIAttributeDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CIAttributeDTO) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *CIAttributeDTO) GetValue() AttributeValueDTO`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CIAttributeDTO) GetValueOk() (*AttributeValueDTO, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CIAttributeDTO) SetValue(v AttributeValueDTO)`

SetValue sets Value field to given value.


### GetCiid

`func (o *CIAttributeDTO) GetCiid() string`

GetCiid returns the Ciid field if non-nil, zero value otherwise.

### GetCiidOk

`func (o *CIAttributeDTO) GetCiidOk() (*string, bool)`

GetCiidOk returns a tuple with the Ciid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiid

`func (o *CIAttributeDTO) SetCiid(v string)`

SetCiid sets Ciid field to given value.


### GetState

`func (o *CIAttributeDTO) GetState() AttributeState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CIAttributeDTO) GetStateOk() (*AttributeState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CIAttributeDTO) SetState(v AttributeState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


