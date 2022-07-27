# GenericInboundCI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TempID** | Pointer to **NullableString** |  | [optional] 
**IdMethod** | Pointer to [**NullableGenericInboundCIIdMethod**](GenericInboundCIIdMethod.md) |  | [optional] 
**SameTempIDHandling** | Pointer to [**SameTempIDHandling**](SameTempIDHandling.md) |  | [optional] 
**SameTargetCIHandling** | Pointer to [**SameTargetCIHandling**](SameTargetCIHandling.md) |  | [optional] 
**NoFoundTargetCIHandling** | Pointer to [**NoFoundTargetCIHandling**](NoFoundTargetCIHandling.md) |  | [optional] 
**Attributes** | Pointer to [**[]GenericInboundAttribute**](GenericInboundAttribute.md) |  | [optional] 

## Methods

### NewGenericInboundCI

`func NewGenericInboundCI() *GenericInboundCI`

NewGenericInboundCI instantiates a new GenericInboundCI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericInboundCIWithDefaults

`func NewGenericInboundCIWithDefaults() *GenericInboundCI`

NewGenericInboundCIWithDefaults instantiates a new GenericInboundCI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTempID

`func (o *GenericInboundCI) GetTempID() string`

GetTempID returns the TempID field if non-nil, zero value otherwise.

### GetTempIDOk

`func (o *GenericInboundCI) GetTempIDOk() (*string, bool)`

GetTempIDOk returns a tuple with the TempID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempID

`func (o *GenericInboundCI) SetTempID(v string)`

SetTempID sets TempID field to given value.

### HasTempID

`func (o *GenericInboundCI) HasTempID() bool`

HasTempID returns a boolean if a field has been set.

### SetTempIDNil

`func (o *GenericInboundCI) SetTempIDNil(b bool)`

 SetTempIDNil sets the value for TempID to be an explicit nil

### UnsetTempID
`func (o *GenericInboundCI) UnsetTempID()`

UnsetTempID ensures that no value is present for TempID, not even an explicit nil
### GetIdMethod

`func (o *GenericInboundCI) GetIdMethod() GenericInboundCIIdMethod`

GetIdMethod returns the IdMethod field if non-nil, zero value otherwise.

### GetIdMethodOk

`func (o *GenericInboundCI) GetIdMethodOk() (*GenericInboundCIIdMethod, bool)`

GetIdMethodOk returns a tuple with the IdMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdMethod

`func (o *GenericInboundCI) SetIdMethod(v GenericInboundCIIdMethod)`

SetIdMethod sets IdMethod field to given value.

### HasIdMethod

`func (o *GenericInboundCI) HasIdMethod() bool`

HasIdMethod returns a boolean if a field has been set.

### SetIdMethodNil

`func (o *GenericInboundCI) SetIdMethodNil(b bool)`

 SetIdMethodNil sets the value for IdMethod to be an explicit nil

### UnsetIdMethod
`func (o *GenericInboundCI) UnsetIdMethod()`

UnsetIdMethod ensures that no value is present for IdMethod, not even an explicit nil
### GetSameTempIDHandling

`func (o *GenericInboundCI) GetSameTempIDHandling() SameTempIDHandling`

GetSameTempIDHandling returns the SameTempIDHandling field if non-nil, zero value otherwise.

### GetSameTempIDHandlingOk

`func (o *GenericInboundCI) GetSameTempIDHandlingOk() (*SameTempIDHandling, bool)`

GetSameTempIDHandlingOk returns a tuple with the SameTempIDHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameTempIDHandling

`func (o *GenericInboundCI) SetSameTempIDHandling(v SameTempIDHandling)`

SetSameTempIDHandling sets SameTempIDHandling field to given value.

### HasSameTempIDHandling

`func (o *GenericInboundCI) HasSameTempIDHandling() bool`

HasSameTempIDHandling returns a boolean if a field has been set.

### GetSameTargetCIHandling

`func (o *GenericInboundCI) GetSameTargetCIHandling() SameTargetCIHandling`

GetSameTargetCIHandling returns the SameTargetCIHandling field if non-nil, zero value otherwise.

### GetSameTargetCIHandlingOk

`func (o *GenericInboundCI) GetSameTargetCIHandlingOk() (*SameTargetCIHandling, bool)`

GetSameTargetCIHandlingOk returns a tuple with the SameTargetCIHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameTargetCIHandling

`func (o *GenericInboundCI) SetSameTargetCIHandling(v SameTargetCIHandling)`

SetSameTargetCIHandling sets SameTargetCIHandling field to given value.

### HasSameTargetCIHandling

`func (o *GenericInboundCI) HasSameTargetCIHandling() bool`

HasSameTargetCIHandling returns a boolean if a field has been set.

### GetNoFoundTargetCIHandling

`func (o *GenericInboundCI) GetNoFoundTargetCIHandling() NoFoundTargetCIHandling`

GetNoFoundTargetCIHandling returns the NoFoundTargetCIHandling field if non-nil, zero value otherwise.

### GetNoFoundTargetCIHandlingOk

`func (o *GenericInboundCI) GetNoFoundTargetCIHandlingOk() (*NoFoundTargetCIHandling, bool)`

GetNoFoundTargetCIHandlingOk returns a tuple with the NoFoundTargetCIHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoFoundTargetCIHandling

`func (o *GenericInboundCI) SetNoFoundTargetCIHandling(v NoFoundTargetCIHandling)`

SetNoFoundTargetCIHandling sets NoFoundTargetCIHandling field to given value.

### HasNoFoundTargetCIHandling

`func (o *GenericInboundCI) HasNoFoundTargetCIHandling() bool`

HasNoFoundTargetCIHandling returns a boolean if a field has been set.

### GetAttributes

`func (o *GenericInboundCI) GetAttributes() []GenericInboundAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GenericInboundCI) GetAttributesOk() (*[]GenericInboundAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GenericInboundCI) SetAttributes(v []GenericInboundAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GenericInboundCI) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *GenericInboundCI) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *GenericInboundCI) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


