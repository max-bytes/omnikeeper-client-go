# AddContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SpeakingName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
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

### GetId

`func (o *AddContextRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddContextRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddContextRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddContextRequest) HasId() bool`

HasId returns a boolean if a field has been set.

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


