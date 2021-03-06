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

// Context struct for Context
type Context struct {
	Id NullableString `json:"id,omitempty"`
	ExtractConfig map[string]interface{} `json:"extractConfig,omitempty"`
	TransformConfig map[string]interface{} `json:"transformConfig,omitempty"`
	LoadConfig *ILoadConfig `json:"loadConfig,omitempty"`
}

// NewContext instantiates a new Context object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContext() *Context {
	this := Context{}
	return &this
}

// NewContextWithDefaults instantiates a new Context object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextWithDefaults() *Context {
	this := Context{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Context) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Context) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Context) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Context) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Context) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Context) UnsetId() {
	o.Id.Unset()
}

// GetExtractConfig returns the ExtractConfig field value if set, zero value otherwise.
func (o *Context) GetExtractConfig() map[string]interface{} {
	if o == nil || o.ExtractConfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtractConfig
}

// GetExtractConfigOk returns a tuple with the ExtractConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetExtractConfigOk() (map[string]interface{}, bool) {
	if o == nil || o.ExtractConfig == nil {
		return nil, false
	}
	return o.ExtractConfig, true
}

// HasExtractConfig returns a boolean if a field has been set.
func (o *Context) HasExtractConfig() bool {
	if o != nil && o.ExtractConfig != nil {
		return true
	}

	return false
}

// SetExtractConfig gets a reference to the given map[string]interface{} and assigns it to the ExtractConfig field.
func (o *Context) SetExtractConfig(v map[string]interface{}) {
	o.ExtractConfig = v
}

// GetTransformConfig returns the TransformConfig field value if set, zero value otherwise.
func (o *Context) GetTransformConfig() map[string]interface{} {
	if o == nil || o.TransformConfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TransformConfig
}

// GetTransformConfigOk returns a tuple with the TransformConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetTransformConfigOk() (map[string]interface{}, bool) {
	if o == nil || o.TransformConfig == nil {
		return nil, false
	}
	return o.TransformConfig, true
}

// HasTransformConfig returns a boolean if a field has been set.
func (o *Context) HasTransformConfig() bool {
	if o != nil && o.TransformConfig != nil {
		return true
	}

	return false
}

// SetTransformConfig gets a reference to the given map[string]interface{} and assigns it to the TransformConfig field.
func (o *Context) SetTransformConfig(v map[string]interface{}) {
	o.TransformConfig = v
}

// GetLoadConfig returns the LoadConfig field value if set, zero value otherwise.
func (o *Context) GetLoadConfig() ILoadConfig {
	if o == nil || o.LoadConfig == nil {
		var ret ILoadConfig
		return ret
	}
	return *o.LoadConfig
}

// GetLoadConfigOk returns a tuple with the LoadConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetLoadConfigOk() (*ILoadConfig, bool) {
	if o == nil || o.LoadConfig == nil {
		return nil, false
	}
	return o.LoadConfig, true
}

// HasLoadConfig returns a boolean if a field has been set.
func (o *Context) HasLoadConfig() bool {
	if o != nil && o.LoadConfig != nil {
		return true
	}

	return false
}

// SetLoadConfig gets a reference to the given ILoadConfig and assigns it to the LoadConfig field.
func (o *Context) SetLoadConfig(v ILoadConfig) {
	o.LoadConfig = &v
}

func (o Context) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ExtractConfig != nil {
		toSerialize["extractConfig"] = o.ExtractConfig
	}
	if o.TransformConfig != nil {
		toSerialize["transformConfig"] = o.TransformConfig
	}
	if o.LoadConfig != nil {
		toSerialize["loadConfig"] = o.LoadConfig
	}
	return json.Marshal(toSerialize)
}

type NullableContext struct {
	value *Context
	isSet bool
}

func (v NullableContext) Get() *Context {
	return v.value
}

func (v *NullableContext) Set(val *Context) {
	v.value = val
	v.isSet = true
}

func (v NullableContext) IsSet() bool {
	return v.isSet
}

func (v *NullableContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContext(val *Context) *NullableContext {
	return &NullableContext{value: val, isSet: true}
}

func (v NullableContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


