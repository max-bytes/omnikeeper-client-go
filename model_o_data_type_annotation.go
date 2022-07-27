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

// ODataTypeAnnotation struct for ODataTypeAnnotation
type ODataTypeAnnotation struct {
	TypeName NullableString `json:"typeName,omitempty"`
}

// NewODataTypeAnnotation instantiates a new ODataTypeAnnotation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewODataTypeAnnotation() *ODataTypeAnnotation {
	this := ODataTypeAnnotation{}
	return &this
}

// NewODataTypeAnnotationWithDefaults instantiates a new ODataTypeAnnotation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewODataTypeAnnotationWithDefaults() *ODataTypeAnnotation {
	this := ODataTypeAnnotation{}
	return &this
}

// GetTypeName returns the TypeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataTypeAnnotation) GetTypeName() string {
	if o == nil || o.TypeName.Get() == nil {
		var ret string
		return ret
	}
	return *o.TypeName.Get()
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataTypeAnnotation) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TypeName.Get(), o.TypeName.IsSet()
}

// HasTypeName returns a boolean if a field has been set.
func (o *ODataTypeAnnotation) HasTypeName() bool {
	if o != nil && o.TypeName.IsSet() {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given NullableString and assigns it to the TypeName field.
func (o *ODataTypeAnnotation) SetTypeName(v string) {
	o.TypeName.Set(&v)
}
// SetTypeNameNil sets the value for TypeName to be an explicit nil
func (o *ODataTypeAnnotation) SetTypeNameNil() {
	o.TypeName.Set(nil)
}

// UnsetTypeName ensures that no value is present for TypeName, not even an explicit nil
func (o *ODataTypeAnnotation) UnsetTypeName() {
	o.TypeName.Unset()
}

func (o ODataTypeAnnotation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TypeName.IsSet() {
		toSerialize["typeName"] = o.TypeName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableODataTypeAnnotation struct {
	value *ODataTypeAnnotation
	isSet bool
}

func (v NullableODataTypeAnnotation) Get() *ODataTypeAnnotation {
	return v.value
}

func (v *NullableODataTypeAnnotation) Set(val *ODataTypeAnnotation) {
	v.value = val
	v.isSet = true
}

func (v NullableODataTypeAnnotation) IsSet() bool {
	return v.isSet
}

func (v *NullableODataTypeAnnotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableODataTypeAnnotation(val *ODataTypeAnnotation) *NullableODataTypeAnnotation {
	return &NullableODataTypeAnnotation{value: val, isSet: true}
}

func (v NullableODataTypeAnnotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableODataTypeAnnotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

