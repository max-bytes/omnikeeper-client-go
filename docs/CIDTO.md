# CIDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Attributes** | [**map[string]CIAttributeDTO**](CIAttributeDTO.md) |  | 

## Methods

### NewCIDTO

`func NewCIDTO(id string, attributes map[string]CIAttributeDTO, ) *CIDTO`

NewCIDTO instantiates a new CIDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCIDTOWithDefaults

`func NewCIDTOWithDefaults() *CIDTO`

NewCIDTOWithDefaults instantiates a new CIDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CIDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CIDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CIDTO) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CIDTO) GetAttributes() map[string]CIAttributeDTO`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CIDTO) GetAttributesOk() (*map[string]CIAttributeDTO, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CIDTO) SetAttributes(v map[string]CIAttributeDTO)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


