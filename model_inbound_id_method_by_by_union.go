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

// InboundIDMethodByByUnion struct for InboundIDMethodByByUnion
type InboundIDMethodByByUnion struct {
	Inner []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect `json:"inner,omitempty"`
}

// NewInboundIDMethodByByUnion instantiates a new InboundIDMethodByByUnion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundIDMethodByByUnion(type_ string) *InboundIDMethodByByUnion {
	this := InboundIDMethodByByUnion{}
	this.Type = type_
	return &this
}

// NewInboundIDMethodByByUnionWithDefaults instantiates a new InboundIDMethodByByUnion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundIDMethodByByUnionWithDefaults() *InboundIDMethodByByUnion {
	this := InboundIDMethodByByUnion{}
	return &this
}

// GetInner returns the Inner field value if set, zero value otherwise.
func (o *InboundIDMethodByByUnion) GetInner() []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect {
	if o == nil || o.Inner == nil {
		var ret []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect
		return ret
	}
	return o.Inner
}

// GetInnerOk returns a tuple with the Inner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundIDMethodByByUnion) GetInnerOk() ([]OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect, bool) {
	if o == nil || o.Inner == nil {
		return nil, false
	}
	return o.Inner, true
}

// HasInner returns a boolean if a field has been set.
func (o *InboundIDMethodByByUnion) HasInner() bool {
	if o != nil && o.Inner != nil {
		return true
	}

	return false
}

// SetInner gets a reference to the given []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect and assigns it to the Inner field.
func (o *InboundIDMethodByByUnion) SetInner(v []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect) {
	o.Inner = v
}

func (o InboundIDMethodByByUnion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Inner != nil {
		toSerialize["inner"] = o.Inner
	}
	return json.Marshal(toSerialize)
}

type NullableInboundIDMethodByByUnion struct {
	value *InboundIDMethodByByUnion
	isSet bool
}

func (v NullableInboundIDMethodByByUnion) Get() *InboundIDMethodByByUnion {
	return v.value
}

func (v *NullableInboundIDMethodByByUnion) Set(val *InboundIDMethodByByUnion) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundIDMethodByByUnion) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundIDMethodByByUnion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundIDMethodByByUnion(val *InboundIDMethodByByUnion) *NullableInboundIDMethodByByUnion {
	return &NullableInboundIDMethodByByUnion{value: val, isSet: true}
}

func (v NullableInboundIDMethodByByUnion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundIDMethodByByUnion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


