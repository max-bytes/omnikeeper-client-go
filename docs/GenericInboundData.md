# GenericInboundData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cis** | Pointer to [**[]GenericInboundCI**](GenericInboundCI.md) |  | [optional] 
**Relations** | Pointer to [**[]GenericInboundRelation**](GenericInboundRelation.md) |  | [optional] 

## Methods

### NewGenericInboundData

`func NewGenericInboundData() *GenericInboundData`

NewGenericInboundData instantiates a new GenericInboundData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericInboundDataWithDefaults

`func NewGenericInboundDataWithDefaults() *GenericInboundData`

NewGenericInboundDataWithDefaults instantiates a new GenericInboundData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCis

`func (o *GenericInboundData) GetCis() []GenericInboundCI`

GetCis returns the Cis field if non-nil, zero value otherwise.

### GetCisOk

`func (o *GenericInboundData) GetCisOk() (*[]GenericInboundCI, bool)`

GetCisOk returns a tuple with the Cis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCis

`func (o *GenericInboundData) SetCis(v []GenericInboundCI)`

SetCis sets Cis field to given value.

### HasCis

`func (o *GenericInboundData) HasCis() bool`

HasCis returns a boolean if a field has been set.

### GetRelations

`func (o *GenericInboundData) GetRelations() []GenericInboundRelation`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *GenericInboundData) GetRelationsOk() (*[]GenericInboundRelation, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *GenericInboundData) SetRelations(v []GenericInboundRelation)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *GenericInboundData) HasRelations() bool`

HasRelations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


