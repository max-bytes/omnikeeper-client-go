# GenericInboundAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to [**IAttributeValue**](IAttributeValue.md) |  | [optional] 

## Methods

### NewGenericInboundAttribute

`func NewGenericInboundAttribute() *GenericInboundAttribute`

NewGenericInboundAttribute instantiates a new GenericInboundAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericInboundAttributeWithDefaults

`func NewGenericInboundAttributeWithDefaults() *GenericInboundAttribute`

NewGenericInboundAttributeWithDefaults instantiates a new GenericInboundAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GenericInboundAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericInboundAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericInboundAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenericInboundAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GenericInboundAttribute) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GenericInboundAttribute) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *GenericInboundAttribute) GetValue() IAttributeValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GenericInboundAttribute) GetValueOk() (*IAttributeValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GenericInboundAttribute) SetValue(v IAttributeValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *GenericInboundAttribute) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


