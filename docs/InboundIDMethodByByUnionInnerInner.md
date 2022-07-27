# InboundIDMethodByByUnionInnerInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** |  | [optional] [readonly] 
**Attributes** | Pointer to **[]string** |  | [optional] 
**CaseInsensitive** | Pointer to **bool** |  | [optional] 
**Attribute** | Pointer to [**GenericInboundAttribute**](GenericInboundAttribute.md) |  | [optional] 
**Modifiers** | Pointer to [**InboundIDMethodByAttributeModifiers**](InboundIDMethodByAttributeModifiers.md) |  | [optional] 
**TempID** | Pointer to **NullableString** |  | [optional] 
**OutgoingRelation** | Pointer to **bool** |  | [optional] 
**PredicateID** | Pointer to **NullableString** |  | [optional] 
**Inner** | Pointer to [**[]InboundIDMethodByByUnionInnerInner**](InboundIDMethodByByUnionInnerInner.md) |  | [optional] 

## Methods

### NewInboundIDMethodByByUnionInnerInner

`func NewInboundIDMethodByByUnionInnerInner() *InboundIDMethodByByUnionInnerInner`

NewInboundIDMethodByByUnionInnerInner instantiates a new InboundIDMethodByByUnionInnerInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundIDMethodByByUnionInnerInnerWithDefaults

`func NewInboundIDMethodByByUnionInnerInnerWithDefaults() *InboundIDMethodByByUnionInnerInner`

NewInboundIDMethodByByUnionInnerInnerWithDefaults instantiates a new InboundIDMethodByByUnionInnerInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InboundIDMethodByByUnionInnerInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InboundIDMethodByByUnionInnerInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InboundIDMethodByByUnionInnerInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InboundIDMethodByByUnionInnerInner) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *InboundIDMethodByByUnionInnerInner) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InboundIDMethodByByUnionInnerInner) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *InboundIDMethodByByUnionInnerInner) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InboundIDMethodByByUnionInnerInner) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InboundIDMethodByByUnionInnerInner) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *InboundIDMethodByByUnionInnerInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *InboundIDMethodByByUnionInnerInner) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *InboundIDMethodByByUnionInnerInner) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetCaseInsensitive

`func (o *InboundIDMethodByByUnionInnerInner) GetCaseInsensitive() bool`

GetCaseInsensitive returns the CaseInsensitive field if non-nil, zero value otherwise.

### GetCaseInsensitiveOk

`func (o *InboundIDMethodByByUnionInnerInner) GetCaseInsensitiveOk() (*bool, bool)`

GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitive

`func (o *InboundIDMethodByByUnionInnerInner) SetCaseInsensitive(v bool)`

SetCaseInsensitive sets CaseInsensitive field to given value.

### HasCaseInsensitive

`func (o *InboundIDMethodByByUnionInnerInner) HasCaseInsensitive() bool`

HasCaseInsensitive returns a boolean if a field has been set.

### GetAttribute

`func (o *InboundIDMethodByByUnionInnerInner) GetAttribute() GenericInboundAttribute`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *InboundIDMethodByByUnionInnerInner) GetAttributeOk() (*GenericInboundAttribute, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *InboundIDMethodByByUnionInnerInner) SetAttribute(v GenericInboundAttribute)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *InboundIDMethodByByUnionInnerInner) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetModifiers

`func (o *InboundIDMethodByByUnionInnerInner) GetModifiers() InboundIDMethodByAttributeModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *InboundIDMethodByByUnionInnerInner) GetModifiersOk() (*InboundIDMethodByAttributeModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *InboundIDMethodByByUnionInnerInner) SetModifiers(v InboundIDMethodByAttributeModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *InboundIDMethodByByUnionInnerInner) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetTempID

`func (o *InboundIDMethodByByUnionInnerInner) GetTempID() string`

GetTempID returns the TempID field if non-nil, zero value otherwise.

### GetTempIDOk

`func (o *InboundIDMethodByByUnionInnerInner) GetTempIDOk() (*string, bool)`

GetTempIDOk returns a tuple with the TempID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempID

`func (o *InboundIDMethodByByUnionInnerInner) SetTempID(v string)`

SetTempID sets TempID field to given value.

### HasTempID

`func (o *InboundIDMethodByByUnionInnerInner) HasTempID() bool`

HasTempID returns a boolean if a field has been set.

### SetTempIDNil

`func (o *InboundIDMethodByByUnionInnerInner) SetTempIDNil(b bool)`

 SetTempIDNil sets the value for TempID to be an explicit nil

### UnsetTempID
`func (o *InboundIDMethodByByUnionInnerInner) UnsetTempID()`

UnsetTempID ensures that no value is present for TempID, not even an explicit nil
### GetOutgoingRelation

`func (o *InboundIDMethodByByUnionInnerInner) GetOutgoingRelation() bool`

GetOutgoingRelation returns the OutgoingRelation field if non-nil, zero value otherwise.

### GetOutgoingRelationOk

`func (o *InboundIDMethodByByUnionInnerInner) GetOutgoingRelationOk() (*bool, bool)`

GetOutgoingRelationOk returns a tuple with the OutgoingRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingRelation

`func (o *InboundIDMethodByByUnionInnerInner) SetOutgoingRelation(v bool)`

SetOutgoingRelation sets OutgoingRelation field to given value.

### HasOutgoingRelation

`func (o *InboundIDMethodByByUnionInnerInner) HasOutgoingRelation() bool`

HasOutgoingRelation returns a boolean if a field has been set.

### GetPredicateID

`func (o *InboundIDMethodByByUnionInnerInner) GetPredicateID() string`

GetPredicateID returns the PredicateID field if non-nil, zero value otherwise.

### GetPredicateIDOk

`func (o *InboundIDMethodByByUnionInnerInner) GetPredicateIDOk() (*string, bool)`

GetPredicateIDOk returns a tuple with the PredicateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicateID

`func (o *InboundIDMethodByByUnionInnerInner) SetPredicateID(v string)`

SetPredicateID sets PredicateID field to given value.

### HasPredicateID

`func (o *InboundIDMethodByByUnionInnerInner) HasPredicateID() bool`

HasPredicateID returns a boolean if a field has been set.

### SetPredicateIDNil

`func (o *InboundIDMethodByByUnionInnerInner) SetPredicateIDNil(b bool)`

 SetPredicateIDNil sets the value for PredicateID to be an explicit nil

### UnsetPredicateID
`func (o *InboundIDMethodByByUnionInnerInner) UnsetPredicateID()`

UnsetPredicateID ensures that no value is present for PredicateID, not even an explicit nil
### GetInner

`func (o *InboundIDMethodByByUnionInnerInner) GetInner() []InboundIDMethodByByUnionInnerInner`

GetInner returns the Inner field if non-nil, zero value otherwise.

### GetInnerOk

`func (o *InboundIDMethodByByUnionInnerInner) GetInnerOk() (*[]InboundIDMethodByByUnionInnerInner, bool)`

GetInnerOk returns a tuple with the Inner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInner

`func (o *InboundIDMethodByByUnionInnerInner) SetInner(v []InboundIDMethodByByUnionInnerInner)`

SetInner sets Inner field to given value.

### HasInner

`func (o *InboundIDMethodByByUnionInnerInner) HasInner() bool`

HasInner returns a boolean if a field has been set.

### SetInnerNil

`func (o *InboundIDMethodByByUnionInnerInner) SetInnerNil(b bool)`

 SetInnerNil sets the value for Inner to be an explicit nil

### UnsetInner
`func (o *InboundIDMethodByByUnionInnerInner) UnsetInner()`

UnsetInner ensures that no value is present for Inner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


