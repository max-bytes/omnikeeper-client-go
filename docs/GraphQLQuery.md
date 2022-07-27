# GraphQLQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationName** | Pointer to **NullableString** |  | [optional] 
**Query** | Pointer to **NullableString** |  | [optional] 
**Variables** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGraphQLQuery

`func NewGraphQLQuery() *GraphQLQuery`

NewGraphQLQuery instantiates a new GraphQLQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLQueryWithDefaults

`func NewGraphQLQueryWithDefaults() *GraphQLQuery`

NewGraphQLQueryWithDefaults instantiates a new GraphQLQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationName

`func (o *GraphQLQuery) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *GraphQLQuery) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *GraphQLQuery) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.

### HasOperationName

`func (o *GraphQLQuery) HasOperationName() bool`

HasOperationName returns a boolean if a field has been set.

### SetOperationNameNil

`func (o *GraphQLQuery) SetOperationNameNil(b bool)`

 SetOperationNameNil sets the value for OperationName to be an explicit nil

### UnsetOperationName
`func (o *GraphQLQuery) UnsetOperationName()`

UnsetOperationName ensures that no value is present for OperationName, not even an explicit nil
### GetQuery

`func (o *GraphQLQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GraphQLQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GraphQLQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GraphQLQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *GraphQLQuery) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *GraphQLQuery) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetVariables

`func (o *GraphQLQuery) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *GraphQLQuery) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *GraphQLQuery) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *GraphQLQuery) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *GraphQLQuery) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *GraphQLQuery) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


