# AnsibleInventoryScanDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SetupFacts** | **map[string]string** |  | 
**YumInstalled** | **map[string]string** |  | 
**YumRepos** | **map[string]string** |  | 
**YumUpdates** | **map[string]string** |  | 

## Methods

### NewAnsibleInventoryScanDTO

`func NewAnsibleInventoryScanDTO(setupFacts map[string]string, yumInstalled map[string]string, yumRepos map[string]string, yumUpdates map[string]string, ) *AnsibleInventoryScanDTO`

NewAnsibleInventoryScanDTO instantiates a new AnsibleInventoryScanDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleInventoryScanDTOWithDefaults

`func NewAnsibleInventoryScanDTOWithDefaults() *AnsibleInventoryScanDTO`

NewAnsibleInventoryScanDTOWithDefaults instantiates a new AnsibleInventoryScanDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetupFacts

`func (o *AnsibleInventoryScanDTO) GetSetupFacts() map[string]string`

GetSetupFacts returns the SetupFacts field if non-nil, zero value otherwise.

### GetSetupFactsOk

`func (o *AnsibleInventoryScanDTO) GetSetupFactsOk() (*map[string]string, bool)`

GetSetupFactsOk returns a tuple with the SetupFacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupFacts

`func (o *AnsibleInventoryScanDTO) SetSetupFacts(v map[string]string)`

SetSetupFacts sets SetupFacts field to given value.


### GetYumInstalled

`func (o *AnsibleInventoryScanDTO) GetYumInstalled() map[string]string`

GetYumInstalled returns the YumInstalled field if non-nil, zero value otherwise.

### GetYumInstalledOk

`func (o *AnsibleInventoryScanDTO) GetYumInstalledOk() (*map[string]string, bool)`

GetYumInstalledOk returns a tuple with the YumInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumInstalled

`func (o *AnsibleInventoryScanDTO) SetYumInstalled(v map[string]string)`

SetYumInstalled sets YumInstalled field to given value.


### GetYumRepos

`func (o *AnsibleInventoryScanDTO) GetYumRepos() map[string]string`

GetYumRepos returns the YumRepos field if non-nil, zero value otherwise.

### GetYumReposOk

`func (o *AnsibleInventoryScanDTO) GetYumReposOk() (*map[string]string, bool)`

GetYumReposOk returns a tuple with the YumRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumRepos

`func (o *AnsibleInventoryScanDTO) SetYumRepos(v map[string]string)`

SetYumRepos sets YumRepos field to given value.


### GetYumUpdates

`func (o *AnsibleInventoryScanDTO) GetYumUpdates() map[string]string`

GetYumUpdates returns the YumUpdates field if non-nil, zero value otherwise.

### GetYumUpdatesOk

`func (o *AnsibleInventoryScanDTO) GetYumUpdatesOk() (*map[string]string, bool)`

GetYumUpdatesOk returns a tuple with the YumUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumUpdates

`func (o *AnsibleInventoryScanDTO) SetYumUpdates(v map[string]string)`

SetYumUpdates sets YumUpdates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


