/*
Landscape omnikeeper REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okclient

import (
	"encoding/json"
	"fmt"
)

// checks if the InboundIDMethodByRelatedTempID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundIDMethodByRelatedTempID{}

// InboundIDMethodByRelatedTempID struct for InboundIDMethodByRelatedTempID
type InboundIDMethodByRelatedTempID struct {
	AbstractInboundIDMethod
	TempID *string `json:"tempID,omitempty"`
	OutgoingRelation *bool `json:"outgoingRelation,omitempty"`
	PredicateID *string `json:"predicateID,omitempty"`
}

type _InboundIDMethodByRelatedTempID InboundIDMethodByRelatedTempID

// NewInboundIDMethodByRelatedTempID instantiates a new InboundIDMethodByRelatedTempID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundIDMethodByRelatedTempID(type_ string) *InboundIDMethodByRelatedTempID {
	this := InboundIDMethodByRelatedTempID{}
	this.Type = type_
	return &this
}

// NewInboundIDMethodByRelatedTempIDWithDefaults instantiates a new InboundIDMethodByRelatedTempID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundIDMethodByRelatedTempIDWithDefaults() *InboundIDMethodByRelatedTempID {
	this := InboundIDMethodByRelatedTempID{}
	return &this
}

// GetTempID returns the TempID field value if set, zero value otherwise.
func (o *InboundIDMethodByRelatedTempID) GetTempID() string {
	if o == nil || IsNil(o.TempID) {
		var ret string
		return ret
	}
	return *o.TempID
}

// GetTempIDOk returns a tuple with the TempID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundIDMethodByRelatedTempID) GetTempIDOk() (*string, bool) {
	if o == nil || IsNil(o.TempID) {
		return nil, false
	}
	return o.TempID, true
}

// HasTempID returns a boolean if a field has been set.
func (o *InboundIDMethodByRelatedTempID) HasTempID() bool {
	if o != nil && !IsNil(o.TempID) {
		return true
	}

	return false
}

// SetTempID gets a reference to the given string and assigns it to the TempID field.
func (o *InboundIDMethodByRelatedTempID) SetTempID(v string) {
	o.TempID = &v
}

// GetOutgoingRelation returns the OutgoingRelation field value if set, zero value otherwise.
func (o *InboundIDMethodByRelatedTempID) GetOutgoingRelation() bool {
	if o == nil || IsNil(o.OutgoingRelation) {
		var ret bool
		return ret
	}
	return *o.OutgoingRelation
}

// GetOutgoingRelationOk returns a tuple with the OutgoingRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundIDMethodByRelatedTempID) GetOutgoingRelationOk() (*bool, bool) {
	if o == nil || IsNil(o.OutgoingRelation) {
		return nil, false
	}
	return o.OutgoingRelation, true
}

// HasOutgoingRelation returns a boolean if a field has been set.
func (o *InboundIDMethodByRelatedTempID) HasOutgoingRelation() bool {
	if o != nil && !IsNil(o.OutgoingRelation) {
		return true
	}

	return false
}

// SetOutgoingRelation gets a reference to the given bool and assigns it to the OutgoingRelation field.
func (o *InboundIDMethodByRelatedTempID) SetOutgoingRelation(v bool) {
	o.OutgoingRelation = &v
}

// GetPredicateID returns the PredicateID field value if set, zero value otherwise.
func (o *InboundIDMethodByRelatedTempID) GetPredicateID() string {
	if o == nil || IsNil(o.PredicateID) {
		var ret string
		return ret
	}
	return *o.PredicateID
}

// GetPredicateIDOk returns a tuple with the PredicateID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundIDMethodByRelatedTempID) GetPredicateIDOk() (*string, bool) {
	if o == nil || IsNil(o.PredicateID) {
		return nil, false
	}
	return o.PredicateID, true
}

// HasPredicateID returns a boolean if a field has been set.
func (o *InboundIDMethodByRelatedTempID) HasPredicateID() bool {
	if o != nil && !IsNil(o.PredicateID) {
		return true
	}

	return false
}

// SetPredicateID gets a reference to the given string and assigns it to the PredicateID field.
func (o *InboundIDMethodByRelatedTempID) SetPredicateID(v string) {
	o.PredicateID = &v
}

func (o InboundIDMethodByRelatedTempID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InboundIDMethodByRelatedTempID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAbstractInboundIDMethod, errAbstractInboundIDMethod := json.Marshal(o.AbstractInboundIDMethod)
	if errAbstractInboundIDMethod != nil {
		return map[string]interface{}{}, errAbstractInboundIDMethod
	}
	errAbstractInboundIDMethod = json.Unmarshal([]byte(serializedAbstractInboundIDMethod), &toSerialize)
	if errAbstractInboundIDMethod != nil {
		return map[string]interface{}{}, errAbstractInboundIDMethod
	}
	if !IsNil(o.TempID) {
		toSerialize["tempID"] = o.TempID
	}
	if !IsNil(o.OutgoingRelation) {
		toSerialize["outgoingRelation"] = o.OutgoingRelation
	}
	if !IsNil(o.PredicateID) {
		toSerialize["predicateID"] = o.PredicateID
	}
	return toSerialize, nil
}

func (o *InboundIDMethodByRelatedTempID) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInboundIDMethodByRelatedTempID := _InboundIDMethodByRelatedTempID{}

	err = json.Unmarshal(bytes, &varInboundIDMethodByRelatedTempID)

	if err != nil {
		return err
	}

	*o = InboundIDMethodByRelatedTempID(varInboundIDMethodByRelatedTempID)

	return err
}

type NullableInboundIDMethodByRelatedTempID struct {
	value *InboundIDMethodByRelatedTempID
	isSet bool
}

func (v NullableInboundIDMethodByRelatedTempID) Get() *InboundIDMethodByRelatedTempID {
	return v.value
}

func (v *NullableInboundIDMethodByRelatedTempID) Set(val *InboundIDMethodByRelatedTempID) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundIDMethodByRelatedTempID) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundIDMethodByRelatedTempID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundIDMethodByRelatedTempID(val *InboundIDMethodByRelatedTempID) *NullableInboundIDMethodByRelatedTempID {
	return &NullableInboundIDMethodByRelatedTempID{value: val, isSet: true}
}

func (v NullableInboundIDMethodByRelatedTempID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundIDMethodByRelatedTempID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


