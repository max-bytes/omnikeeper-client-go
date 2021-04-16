/*
 * Landscape omnikeeper REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okclient

import (
	"encoding/json"
	"fmt"
)

// AttributeState the model 'AttributeState'
type AttributeState string

// List of AttributeState
const (
	ATTRIBUTESTATE_NEW AttributeState = "New"
	ATTRIBUTESTATE_CHANGED AttributeState = "Changed"
	ATTRIBUTESTATE_REMOVED AttributeState = "Removed"
	ATTRIBUTESTATE_RENEWED AttributeState = "Renewed"
)

var allowedAttributeStateEnumValues = []AttributeState{
	"New",
	"Changed",
	"Removed",
	"Renewed",
}

func (v *AttributeState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AttributeState(value)
	for _, existing := range allowedAttributeStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AttributeState", value)
}

// NewAttributeStateFromValue returns a pointer to a valid AttributeState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAttributeStateFromValue(v string) (*AttributeState, error) {
	ev := AttributeState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AttributeState: valid values are %v", v, allowedAttributeStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AttributeState) IsValid() bool {
	for _, existing := range allowedAttributeStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AttributeState value
func (v AttributeState) Ptr() *AttributeState {
	return &v
}

type NullableAttributeState struct {
	value *AttributeState
	isSet bool
}

func (v NullableAttributeState) Get() *AttributeState {
	return v.value
}

func (v *NullableAttributeState) Set(val *AttributeState) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeState) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeState(val *AttributeState) *NullableAttributeState {
	return &NullableAttributeState{value: val, isSet: true}
}

func (v NullableAttributeState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

