# GenericInboundCIIdMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseInsensitive** | Pointer to **bool** |  | [optional] 
**Attribute** | Pointer to [**GenericInboundAttribute**](GenericInboundAttribute.md) |  | [optional] 
**Modifiers** | Pointer to [**InboundIDMethodByAttributeModifiers**](InboundIDMethodByAttributeModifiers.md) |  | [optional] 
**Type** | **string** |  | 
**TempID** | Pointer to **string** |  | [optional] 
**OutgoingRelation** | Pointer to **bool** |  | [optional] 
**PredicateID** | Pointer to **string** |  | [optional] 
**Inner** | Pointer to [**[]OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect**](OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect.md) |  | [optional] 

## Methods

### NewGenericInboundCIIdMethod

`func NewGenericInboundCIIdMethod(type_ string, ) *GenericInboundCIIdMethod`

NewGenericInboundCIIdMethod instantiates a new GenericInboundCIIdMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericInboundCIIdMethodWithDefaults

`func NewGenericInboundCIIdMethodWithDefaults() *GenericInboundCIIdMethod`

NewGenericInboundCIIdMethodWithDefaults instantiates a new GenericInboundCIIdMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseInsensitive

`func (o *GenericInboundCIIdMethod) GetCaseInsensitive() bool`

GetCaseInsensitive returns the CaseInsensitive field if non-nil, zero value otherwise.

### GetCaseInsensitiveOk

`func (o *GenericInboundCIIdMethod) GetCaseInsensitiveOk() (*bool, bool)`

GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitive

`func (o *GenericInboundCIIdMethod) SetCaseInsensitive(v bool)`

SetCaseInsensitive sets CaseInsensitive field to given value.

### HasCaseInsensitive

`func (o *GenericInboundCIIdMethod) HasCaseInsensitive() bool`

HasCaseInsensitive returns a boolean if a field has been set.

### GetAttribute

`func (o *GenericInboundCIIdMethod) GetAttribute() GenericInboundAttribute`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *GenericInboundCIIdMethod) GetAttributeOk() (*GenericInboundAttribute, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *GenericInboundCIIdMethod) SetAttribute(v GenericInboundAttribute)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *GenericInboundCIIdMethod) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetModifiers

`func (o *GenericInboundCIIdMethod) GetModifiers() InboundIDMethodByAttributeModifiers`

GetModifiers returns the Modifiers field if non-nil, zero value otherwise.

### GetModifiersOk

`func (o *GenericInboundCIIdMethod) GetModifiersOk() (*InboundIDMethodByAttributeModifiers, bool)`

GetModifiersOk returns a tuple with the Modifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiers

`func (o *GenericInboundCIIdMethod) SetModifiers(v InboundIDMethodByAttributeModifiers)`

SetModifiers sets Modifiers field to given value.

### HasModifiers

`func (o *GenericInboundCIIdMethod) HasModifiers() bool`

HasModifiers returns a boolean if a field has been set.

### GetType

`func (o *GenericInboundCIIdMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericInboundCIIdMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericInboundCIIdMethod) SetType(v string)`

SetType sets Type field to given value.


### GetTempID

`func (o *GenericInboundCIIdMethod) GetTempID() string`

GetTempID returns the TempID field if non-nil, zero value otherwise.

### GetTempIDOk

`func (o *GenericInboundCIIdMethod) GetTempIDOk() (*string, bool)`

GetTempIDOk returns a tuple with the TempID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempID

`func (o *GenericInboundCIIdMethod) SetTempID(v string)`

SetTempID sets TempID field to given value.

### HasTempID

`func (o *GenericInboundCIIdMethod) HasTempID() bool`

HasTempID returns a boolean if a field has been set.

### GetOutgoingRelation

`func (o *GenericInboundCIIdMethod) GetOutgoingRelation() bool`

GetOutgoingRelation returns the OutgoingRelation field if non-nil, zero value otherwise.

### GetOutgoingRelationOk

`func (o *GenericInboundCIIdMethod) GetOutgoingRelationOk() (*bool, bool)`

GetOutgoingRelationOk returns a tuple with the OutgoingRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingRelation

`func (o *GenericInboundCIIdMethod) SetOutgoingRelation(v bool)`

SetOutgoingRelation sets OutgoingRelation field to given value.

### HasOutgoingRelation

`func (o *GenericInboundCIIdMethod) HasOutgoingRelation() bool`

HasOutgoingRelation returns a boolean if a field has been set.

### GetPredicateID

`func (o *GenericInboundCIIdMethod) GetPredicateID() string`

GetPredicateID returns the PredicateID field if non-nil, zero value otherwise.

### GetPredicateIDOk

`func (o *GenericInboundCIIdMethod) GetPredicateIDOk() (*string, bool)`

GetPredicateIDOk returns a tuple with the PredicateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicateID

`func (o *GenericInboundCIIdMethod) SetPredicateID(v string)`

SetPredicateID sets PredicateID field to given value.

### HasPredicateID

`func (o *GenericInboundCIIdMethod) HasPredicateID() bool`

HasPredicateID returns a boolean if a field has been set.

### GetInner

`func (o *GenericInboundCIIdMethod) GetInner() []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect`

GetInner returns the Inner field if non-nil, zero value otherwise.

### GetInnerOk

`func (o *GenericInboundCIIdMethod) GetInnerOk() (*[]OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect, bool)`

GetInnerOk returns a tuple with the Inner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInner

`func (o *GenericInboundCIIdMethod) SetInner(v []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect)`

SetInner sets Inner field to given value.

### HasInner

`func (o *GenericInboundCIIdMethod) HasInner() bool`

HasInner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


