/*
Landscape omnikeeper REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okclient

import (
	"encoding/json"
)

// checks if the AddContextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddContextRequest{}

// AddContextRequest struct for AddContextRequest
type AddContextRequest struct {
	Id *string `json:"id,omitempty"`
	SpeakingName *string `json:"speakingName,omitempty"`
	Description *string `json:"description,omitempty"`
	Configuration *GridViewConfiguration `json:"configuration,omitempty"`
}

// NewAddContextRequest instantiates a new AddContextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddContextRequest() *AddContextRequest {
	this := AddContextRequest{}
	return &this
}

// NewAddContextRequestWithDefaults instantiates a new AddContextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddContextRequestWithDefaults() *AddContextRequest {
	this := AddContextRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddContextRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContextRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AddContextRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AddContextRequest) SetId(v string) {
	o.Id = &v
}

// GetSpeakingName returns the SpeakingName field value if set, zero value otherwise.
func (o *AddContextRequest) GetSpeakingName() string {
	if o == nil || IsNil(o.SpeakingName) {
		var ret string
		return ret
	}
	return *o.SpeakingName
}

// GetSpeakingNameOk returns a tuple with the SpeakingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContextRequest) GetSpeakingNameOk() (*string, bool) {
	if o == nil || IsNil(o.SpeakingName) {
		return nil, false
	}
	return o.SpeakingName, true
}

// HasSpeakingName returns a boolean if a field has been set.
func (o *AddContextRequest) HasSpeakingName() bool {
	if o != nil && !IsNil(o.SpeakingName) {
		return true
	}

	return false
}

// SetSpeakingName gets a reference to the given string and assigns it to the SpeakingName field.
func (o *AddContextRequest) SetSpeakingName(v string) {
	o.SpeakingName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddContextRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContextRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddContextRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddContextRequest) SetDescription(v string) {
	o.Description = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *AddContextRequest) GetConfiguration() GridViewConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret GridViewConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddContextRequest) GetConfigurationOk() (*GridViewConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *AddContextRequest) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given GridViewConfiguration and assigns it to the Configuration field.
func (o *AddContextRequest) SetConfiguration(v GridViewConfiguration) {
	o.Configuration = &v
}

func (o AddContextRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddContextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SpeakingName) {
		toSerialize["speakingName"] = o.SpeakingName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	return toSerialize, nil
}

type NullableAddContextRequest struct {
	value *AddContextRequest
	isSet bool
}

func (v NullableAddContextRequest) Get() *AddContextRequest {
	return v.value
}

func (v *NullableAddContextRequest) Set(val *AddContextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddContextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddContextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddContextRequest(val *AddContextRequest) *NullableAddContextRequest {
	return &NullableAddContextRequest{value: val, isSet: true}
}

func (v NullableAddContextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddContextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


