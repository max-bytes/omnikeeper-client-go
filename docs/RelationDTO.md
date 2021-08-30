# RelationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FromCIID** | **string** |  | 
**ToCIID** | **string** |  | 
**PredicateID** | **string** |  | 
**State** | [**RelationState**](RelationState.md) |  | 

## Methods

### NewRelationDTO

`func NewRelationDTO(id string, fromCIID string, toCIID string, predicateID string, state RelationState, ) *RelationDTO`

NewRelationDTO instantiates a new RelationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationDTOWithDefaults

`func NewRelationDTOWithDefaults() *RelationDTO`

NewRelationDTOWithDefaults instantiates a new RelationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelationDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelationDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelationDTO) SetId(v string)`

SetId sets Id field to given value.


### GetFromCIID

`func (o *RelationDTO) GetFromCIID() string`

GetFromCIID returns the FromCIID field if non-nil, zero value otherwise.

### GetFromCIIDOk

`func (o *RelationDTO) GetFromCIIDOk() (*string, bool)`

GetFromCIIDOk returns a tuple with the FromCIID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCIID

`func (o *RelationDTO) SetFromCIID(v string)`

SetFromCIID sets FromCIID field to given value.


### GetToCIID

`func (o *RelationDTO) GetToCIID() string`

GetToCIID returns the ToCIID field if non-nil, zero value otherwise.

### GetToCIIDOk

`func (o *RelationDTO) GetToCIIDOk() (*string, bool)`

GetToCIIDOk returns a tuple with the ToCIID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCIID

`func (o *RelationDTO) SetToCIID(v string)`

SetToCIID sets ToCIID field to given value.


### GetPredicateID

`func (o *RelationDTO) GetPredicateID() string`

GetPredicateID returns the PredicateID field if non-nil, zero value otherwise.

### GetPredicateIDOk

`func (o *RelationDTO) GetPredicateIDOk() (*string, bool)`

GetPredicateIDOk returns a tuple with the PredicateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicateID

`func (o *RelationDTO) SetPredicateID(v string)`

SetPredicateID sets PredicateID field to given value.


### GetState

`func (o *RelationDTO) GetState() RelationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RelationDTO) GetStateOk() (*RelationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RelationDTO) SetState(v RelationState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


