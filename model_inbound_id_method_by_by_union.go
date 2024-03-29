/*
Landscape omnikeeper REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the InboundIDMethodByByUnion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundIDMethodByByUnion{}

// InboundIDMethodByByUnion struct for InboundIDMethodByByUnion
type InboundIDMethodByByUnion struct {
	AbstractInboundIDMethod
	Inner []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect `json:"inner,omitempty"`
}

type _InboundIDMethodByByUnion InboundIDMethodByByUnion

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
	if o == nil || IsNil(o.Inner) {
		var ret []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect
		return ret
	}
	return o.Inner
}

// GetInnerOk returns a tuple with the Inner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundIDMethodByByUnion) GetInnerOk() ([]OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect, bool) {
	if o == nil || IsNil(o.Inner) {
		return nil, false
	}
	return o.Inner, true
}

// HasInner returns a boolean if a field has been set.
func (o *InboundIDMethodByByUnion) HasInner() bool {
	if o != nil && !IsNil(o.Inner) {
		return true
	}

	return false
}

// SetInner gets a reference to the given []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect and assigns it to the Inner field.
func (o *InboundIDMethodByByUnion) SetInner(v []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect) {
	o.Inner = v
}

func (o InboundIDMethodByByUnion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InboundIDMethodByByUnion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAbstractInboundIDMethod, errAbstractInboundIDMethod := json.Marshal(o.AbstractInboundIDMethod)
	if errAbstractInboundIDMethod != nil {
		return map[string]interface{}{}, errAbstractInboundIDMethod
	}
	errAbstractInboundIDMethod = json.Unmarshal([]byte(serializedAbstractInboundIDMethod), &toSerialize)
	if errAbstractInboundIDMethod != nil {
		return map[string]interface{}{}, errAbstractInboundIDMethod
	}
	if !IsNil(o.Inner) {
		toSerialize["inner"] = o.Inner
	}
	return toSerialize, nil
}

func (o *InboundIDMethodByByUnion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInboundIDMethodByByUnion := _InboundIDMethodByByUnion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInboundIDMethodByByUnion)

	if err != nil {
		return err
	}

	*o = InboundIDMethodByByUnion(varInboundIDMethodByByUnion)

	return err
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


