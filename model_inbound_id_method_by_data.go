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

// checks if the InboundIDMethodByData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundIDMethodByData{}

// InboundIDMethodByData struct for InboundIDMethodByData
type InboundIDMethodByData struct {
	AbstractInboundIDMethod
	Attributes []string `json:"attributes,omitempty"`
}

type _InboundIDMethodByData InboundIDMethodByData

// NewInboundIDMethodByData instantiates a new InboundIDMethodByData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundIDMethodByData(type_ string) *InboundIDMethodByData {
	this := InboundIDMethodByData{}
	this.Type = type_
	return &this
}

// NewInboundIDMethodByDataWithDefaults instantiates a new InboundIDMethodByData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundIDMethodByDataWithDefaults() *InboundIDMethodByData {
	this := InboundIDMethodByData{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InboundIDMethodByData) GetAttributes() []string {
	if o == nil || IsNil(o.Attributes) {
		var ret []string
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundIDMethodByData) GetAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InboundIDMethodByData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []string and assigns it to the Attributes field.
func (o *InboundIDMethodByData) SetAttributes(v []string) {
	o.Attributes = v
}

func (o InboundIDMethodByData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InboundIDMethodByData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAbstractInboundIDMethod, errAbstractInboundIDMethod := json.Marshal(o.AbstractInboundIDMethod)
	if errAbstractInboundIDMethod != nil {
		return map[string]interface{}{}, errAbstractInboundIDMethod
	}
	errAbstractInboundIDMethod = json.Unmarshal([]byte(serializedAbstractInboundIDMethod), &toSerialize)
	if errAbstractInboundIDMethod != nil {
		return map[string]interface{}{}, errAbstractInboundIDMethod
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

func (o *InboundIDMethodByData) UnmarshalJSON(data []byte) (err error) {
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

	varInboundIDMethodByData := _InboundIDMethodByData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInboundIDMethodByData)

	if err != nil {
		return err
	}

	*o = InboundIDMethodByData(varInboundIDMethodByData)

	return err
}

type NullableInboundIDMethodByData struct {
	value *InboundIDMethodByData
	isSet bool
}

func (v NullableInboundIDMethodByData) Get() *InboundIDMethodByData {
	return v.value
}

func (v *NullableInboundIDMethodByData) Set(val *InboundIDMethodByData) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundIDMethodByData) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundIDMethodByData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundIDMethodByData(val *InboundIDMethodByData) *NullableInboundIDMethodByData {
	return &NullableInboundIDMethodByData{value: val, isSet: true}
}

func (v NullableInboundIDMethodByData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundIDMethodByData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


