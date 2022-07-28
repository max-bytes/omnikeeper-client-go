# EditContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpeakingName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to [**GridViewConfiguration**](GridViewConfiguration.md) |  | [optional] 

## Methods

### NewEditContextRequest

`func NewEditContextRequest() *EditContextRequest`

NewEditContextRequest instantiates a new EditContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditContextRequestWithDefaults

`func NewEditContextRequestWithDefaults() *EditContextRequest`

NewEditContextRequestWithDefaults instantiates a new EditContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpeakingName

`func (o *EditContextRequest) GetSpeakingName() string`

GetSpeakingName returns the SpeakingName field if non-nil, zero value otherwise.

### GetSpeakingNameOk

`func (o *EditContextRequest) GetSpeakingNameOk() (*string, bool)`

GetSpeakingNameOk returns a tuple with the SpeakingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakingName

`func (o *EditContextRequest) SetSpeakingName(v string)`

SetSpeakingName sets SpeakingName field to given value.

### HasSpeakingName

`func (o *EditContextRequest) HasSpeakingName() bool`

HasSpeakingName returns a boolean if a field has been set.

### GetDescription

`func (o *EditContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfiguration

`func (o *EditContextRequest) GetConfiguration() GridViewConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EditContextRequest) GetConfigurationOk() (*GridViewConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EditContextRequest) SetConfiguration(v GridViewConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EditContextRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


