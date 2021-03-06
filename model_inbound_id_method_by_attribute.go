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

// InboundIDMethodByAttribute struct for InboundIDMethodByAttribute
type InboundIDMethodByAttribute struct {
	Attribute *GenericInboundAttribute `json:"attribute,omitempty"`
	Modifiers *InboundIDMethodByAttributeModifiers `json:"modifiers,omitempty"`
}

// NewInboundIDMethodByAttribute instantiates a new InboundIDMethodByAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundIDMethodByAttribute(type_ string) *InboundIDMethodByAttribute {
	this := InboundIDMethodByAttribute{}
	this.Type = type_
	return &this
}

// NewInboundIDMethodByAttributeWithDefaults instantiates a new InboundIDMethodByAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundIDMethodByAttributeWithDefaults() *InboundIDMethodByAttribute {
	this := InboundIDMethodByAttribute{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *InboundIDMethodByAttribute) GetAttribute() GenericInboundAttribute {
	if o == nil || o.Attribute == nil {
		var ret GenericInboundAttribute
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundIDMethodByAttribute) GetAttributeOk() (*GenericInboundAttribute, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *InboundIDMethodByAttribute) HasAttribute() bool {
	if o != nil && o.Attribute != nil {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given GenericInboundAttribute and assigns it to the Attribute field.
func (o *InboundIDMethodByAttribute) SetAttribute(v GenericInboundAttribute) {
	o.Attribute = &v
}

// GetModifiers returns the Modifiers field value if set, zero value otherwise.
func (o *InboundIDMethodByAttribute) GetModifiers() InboundIDMethodByAttributeModifiers {
	if o == nil || o.Modifiers == nil {
		var ret InboundIDMethodByAttributeModifiers
		return ret
	}
	return *o.Modifiers
}

// GetModifiersOk returns a tuple with the Modifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundIDMethodByAttribute) GetModifiersOk() (*InboundIDMethodByAttributeModifiers, bool) {
	if o == nil || o.Modifiers == nil {
		return nil, false
	}
	return o.Modifiers, true
}

// HasModifiers returns a boolean if a field has been set.
func (o *InboundIDMethodByAttribute) HasModifiers() bool {
	if o != nil && o.Modifiers != nil {
		return true
	}

	return false
}

// SetModifiers gets a reference to the given InboundIDMethodByAttributeModifiers and assigns it to the Modifiers field.
func (o *InboundIDMethodByAttribute) SetModifiers(v InboundIDMethodByAttributeModifiers) {
	o.Modifiers = &v
}

func (o InboundIDMethodByAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attribute != nil {
		toSerialize["attribute"] = o.Attribute
	}
	if o.Modifiers != nil {
		toSerialize["modifiers"] = o.Modifiers
	}
	return json.Marshal(toSerialize)
}

type NullableInboundIDMethodByAttribute struct {
	value *InboundIDMethodByAttribute
	isSet bool
}

func (v NullableInboundIDMethodByAttribute) Get() *InboundIDMethodByAttribute {
	return v.value
}

func (v *NullableInboundIDMethodByAttribute) Set(val *InboundIDMethodByAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundIDMethodByAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundIDMethodByAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundIDMethodByAttribute(val *InboundIDMethodByAttribute) *NullableInboundIDMethodByAttribute {
	return &NullableInboundIDMethodByAttribute{value: val, isSet: true}
}

func (v NullableInboundIDMethodByAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundIDMethodByAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


