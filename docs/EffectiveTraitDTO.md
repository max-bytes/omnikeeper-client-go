# EffectiveTraitDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraitAttributes** | [**map[string]CIAttributeDTO**](CIAttributeDTO.md) |  | 
**TraitRelations** | [**map[string][]RelatedCIDTO**](array.md) |  | 

## Methods

### NewEffectiveTraitDTO

`func NewEffectiveTraitDTO(traitAttributes map[string]CIAttributeDTO, traitRelations map[string][]RelatedCIDTO, ) *EffectiveTraitDTO`

NewEffectiveTraitDTO instantiates a new EffectiveTraitDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectiveTraitDTOWithDefaults

`func NewEffectiveTraitDTOWithDefaults() *EffectiveTraitDTO`

NewEffectiveTraitDTOWithDefaults instantiates a new EffectiveTraitDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraitAttributes

`func (o *EffectiveTraitDTO) GetTraitAttributes() map[string]CIAttributeDTO`

GetTraitAttributes returns the TraitAttributes field if non-nil, zero value otherwise.

### GetTraitAttributesOk

`func (o *EffectiveTraitDTO) GetTraitAttributesOk() (*map[string]CIAttributeDTO, bool)`

GetTraitAttributesOk returns a tuple with the TraitAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitAttributes

`func (o *EffectiveTraitDTO) SetTraitAttributes(v map[string]CIAttributeDTO)`

SetTraitAttributes sets TraitAttributes field to given value.


### GetTraitRelations

`func (o *EffectiveTraitDTO) GetTraitRelations() map[string][]RelatedCIDTO`

GetTraitRelations returns the TraitRelations field if non-nil, zero value otherwise.

### GetTraitRelationsOk

`func (o *EffectiveTraitDTO) GetTraitRelationsOk() (*map[string][]RelatedCIDTO, bool)`

GetTraitRelationsOk returns a tuple with the TraitRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraitRelations

`func (o *EffectiveTraitDTO) SetTraitRelations(v map[string][]RelatedCIDTO)`

SetTraitRelations sets TraitRelations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


