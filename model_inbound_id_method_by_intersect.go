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

// checks if the InboundIDMethodByIntersect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundIDMethodByIntersect{}

// InboundIDMethodByIntersect struct for InboundIDMethodByIntersect
type InboundIDMethodByIntersect struct {
	AbstractInboundIDMethod
	Inner []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect `json:"inner,omitempty"`
}

type _InboundIDMethodByIntersect InboundIDMethodByIntersect

// NewInboundIDMethodByIntersect instantiates a new InboundIDMethodByIntersect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundIDMethodByIntersect(type_ string) *InboundIDMethodByIntersect {
	this := InboundIDMethodByIntersect{}
	this.Type = type_
	return &this
}

// NewInboundIDMethodByIntersectWithDefaults instantiates a new InboundIDMethodByIntersect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundIDMethodByIntersectWithDefaults() *InboundIDMethodByIntersect {
	this := InboundIDMethodByIntersect{}
	return &this
}

// GetInner returns the Inner field value if set, zero value otherwise.
func (o *InboundIDMethodByIntersect) GetInner() []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect {
	if o == nil || IsNil(o.Inner) {
		var ret []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect
		return ret
	}
	return o.Inner
}

// GetInnerOk returns a tuple with the Inner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundIDMethodByIntersect) GetInnerOk() ([]OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect, bool) {
	if o == nil || IsNil(o.Inner) {
		return nil, false
	}
	return o.Inner, true
}

// HasInner returns a boolean if a field has been set.
func (o *InboundIDMethodByIntersect) HasInner() bool {
	if o != nil && !IsNil(o.Inner) {
		return true
	}

	return false
}

// SetInner gets a reference to the given []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect and assigns it to the Inner field.
func (o *InboundIDMethodByIntersect) SetInner(v []OneOfInboundIDMethodByDataInboundIDMethodByAttributeModifiersInboundIDMethodByAttributeInboundIDMethodByRelatedTempIDInboundIDMethodByTemporaryCIIDInboundIDMethodByByUnionInboundIDMethodByIntersect) {
	o.Inner = v
}

func (o InboundIDMethodByIntersect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InboundIDMethodByIntersect) ToMap() (map[string]interface{}, error) {
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

func (o *InboundIDMethodByIntersect) UnmarshalJSON(data []byte) (err error) {
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

	varInboundIDMethodByIntersect := _InboundIDMethodByIntersect{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInboundIDMethodByIntersect)

	if err != nil {
		return err
	}

	*o = InboundIDMethodByIntersect(varInboundIDMethodByIntersect)

	return err
}

type NullableInboundIDMethodByIntersect struct {
	value *InboundIDMethodByIntersect
	isSet bool
}

func (v NullableInboundIDMethodByIntersect) Get() *InboundIDMethodByIntersect {
	return v.value
}

func (v *NullableInboundIDMethodByIntersect) Set(val *InboundIDMethodByIntersect) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundIDMethodByIntersect) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundIDMethodByIntersect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundIDMethodByIntersect(val *InboundIDMethodByIntersect) *NullableInboundIDMethodByIntersect {
	return &NullableInboundIDMethodByIntersect{value: val, isSet: true}
}

func (v NullableInboundIDMethodByIntersect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundIDMethodByIntersect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


