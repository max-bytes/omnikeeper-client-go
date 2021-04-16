# AddContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**SpeakingName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Configuration** | Pointer to [**GridViewConfiguration**](GridViewConfiguration.md) |  | [optional] 

## Methods

### NewAddContextRequest

`func NewAddContextRequest() *AddContextRequest`

NewAddContextRequest instantiates a new AddContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddContextRequestWithDefaults

`func NewAddContextRequestWithDefaults() *AddContextRequest`

NewAddContextRequestWithDefaults instantiates a new AddContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddContextRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddContextRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AddContextRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AddContextRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSpeakingName

`func (o *AddContextRequest) GetSpeakingName() string`

GetSpeakingName returns the SpeakingName field if non-nil, zero value otherwise.

### GetSpeakingNameOk

`func (o *AddContextRequest) GetSpeakingNameOk() (*string, bool)`

GetSpeakingNameOk returns a tuple with the SpeakingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakingName

`func (o *AddContextRequest) SetSpeakingName(v string)`

SetSpeakingName sets SpeakingName field to given value.

### HasSpeakingName

`func (o *AddContextRequest) HasSpeakingName() bool`

HasSpeakingName returns a boolean if a field has been set.

### SetSpeakingNameNil

`func (o *AddContextRequest) SetSpeakingNameNil(b bool)`

 SetSpeakingNameNil sets the value for SpeakingName to be an explicit nil

### UnsetSpeakingName
`func (o *AddContextRequest) UnsetSpeakingName()`

UnsetSpeakingName ensures that no value is present for SpeakingName, not even an explicit nil
### GetDescription

`func (o *AddContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddContextRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddContextRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetConfiguration

`func (o *AddContextRequest) GetConfiguration() GridViewConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AddContextRequest) GetConfigurationOk() (*GridViewConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AddContextRequest) SetConfiguration(v GridViewConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *AddContextRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


