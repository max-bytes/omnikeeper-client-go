# Context

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**ExtractConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**TransformConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**LoadConfig** | Pointer to [**ILoadConfig**](ILoadConfig.md) |  | [optional] 

## Methods

### NewContext

`func NewContext() *Context`

NewContext instantiates a new Context object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextWithDefaults

`func NewContextWithDefaults() *Context`

NewContextWithDefaults instantiates a new Context object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Context) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Context) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Context) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Context) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Context) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Context) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetExtractConfig

`func (o *Context) GetExtractConfig() map[string]interface{}`

GetExtractConfig returns the ExtractConfig field if non-nil, zero value otherwise.

### GetExtractConfigOk

`func (o *Context) GetExtractConfigOk() (*map[string]interface{}, bool)`

GetExtractConfigOk returns a tuple with the ExtractConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractConfig

`func (o *Context) SetExtractConfig(v map[string]interface{})`

SetExtractConfig sets ExtractConfig field to given value.

### HasExtractConfig

`func (o *Context) HasExtractConfig() bool`

HasExtractConfig returns a boolean if a field has been set.

### GetTransformConfig

`func (o *Context) GetTransformConfig() map[string]interface{}`

GetTransformConfig returns the TransformConfig field if non-nil, zero value otherwise.

### GetTransformConfigOk

`func (o *Context) GetTransformConfigOk() (*map[string]interface{}, bool)`

GetTransformConfigOk returns a tuple with the TransformConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformConfig

`func (o *Context) SetTransformConfig(v map[string]interface{})`

SetTransformConfig sets TransformConfig field to given value.

### HasTransformConfig

`func (o *Context) HasTransformConfig() bool`

HasTransformConfig returns a boolean if a field has been set.

### GetLoadConfig

`func (o *Context) GetLoadConfig() ILoadConfig`

GetLoadConfig returns the LoadConfig field if non-nil, zero value otherwise.

### GetLoadConfigOk

`func (o *Context) GetLoadConfigOk() (*ILoadConfig, bool)`

GetLoadConfigOk returns a tuple with the LoadConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadConfig

`func (o *Context) SetLoadConfig(v ILoadConfig)`

SetLoadConfig sets LoadConfig field to given value.

### HasLoadConfig

`func (o *Context) HasLoadConfig() bool`

HasLoadConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


