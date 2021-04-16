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
)

// PredicateDTO struct for PredicateDTO
type PredicateDTO struct {
	Id NullableString `json:"id,omitempty"`
	WordingFrom NullableString `json:"wordingFrom,omitempty"`
	WordingTo NullableString `json:"wordingTo,omitempty"`
}

// NewPredicateDTO instantiates a new PredicateDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredicateDTO() *PredicateDTO {
	this := PredicateDTO{}
	return &this
}

// NewPredicateDTOWithDefaults instantiates a new PredicateDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredicateDTOWithDefaults() *PredicateDTO {
	this := PredicateDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PredicateDTO) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PredicateDTO) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *PredicateDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *PredicateDTO) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *PredicateDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *PredicateDTO) UnsetId() {
	o.Id.Unset()
}

// GetWordingFrom returns the WordingFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PredicateDTO) GetWordingFrom() string {
	if o == nil || o.WordingFrom.Get() == nil {
		var ret string
		return ret
	}
	return *o.WordingFrom.Get()
}

// GetWordingFromOk returns a tuple with the WordingFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PredicateDTO) GetWordingFromOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WordingFrom.Get(), o.WordingFrom.IsSet()
}

// HasWordingFrom returns a boolean if a field has been set.
func (o *PredicateDTO) HasWordingFrom() bool {
	if o != nil && o.WordingFrom.IsSet() {
		return true
	}

	return false
}

// SetWordingFrom gets a reference to the given NullableString and assigns it to the WordingFrom field.
func (o *PredicateDTO) SetWordingFrom(v string) {
	o.WordingFrom.Set(&v)
}
// SetWordingFromNil sets the value for WordingFrom to be an explicit nil
func (o *PredicateDTO) SetWordingFromNil() {
	o.WordingFrom.Set(nil)
}

// UnsetWordingFrom ensures that no value is present for WordingFrom, not even an explicit nil
func (o *PredicateDTO) UnsetWordingFrom() {
	o.WordingFrom.Unset()
}

// GetWordingTo returns the WordingTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PredicateDTO) GetWordingTo() string {
	if o == nil || o.WordingTo.Get() == nil {
		var ret string
		return ret
	}
	return *o.WordingTo.Get()
}

// GetWordingToOk returns a tuple with the WordingTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PredicateDTO) GetWordingToOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WordingTo.Get(), o.WordingTo.IsSet()
}

// HasWordingTo returns a boolean if a field has been set.
func (o *PredicateDTO) HasWordingTo() bool {
	if o != nil && o.WordingTo.IsSet() {
		return true
	}

	return false
}

// SetWordingTo gets a reference to the given NullableString and assigns it to the WordingTo field.
func (o *PredicateDTO) SetWordingTo(v string) {
	o.WordingTo.Set(&v)
}
// SetWordingToNil sets the value for WordingTo to be an explicit nil
func (o *PredicateDTO) SetWordingToNil() {
	o.WordingTo.Set(nil)
}

// UnsetWordingTo ensures that no value is present for WordingTo, not even an explicit nil
func (o *PredicateDTO) UnsetWordingTo() {
	o.WordingTo.Unset()
}

func (o PredicateDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.WordingFrom.IsSet() {
		toSerialize["wordingFrom"] = o.WordingFrom.Get()
	}
	if o.WordingTo.IsSet() {
		toSerialize["wordingTo"] = o.WordingTo.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePredicateDTO struct {
	value *PredicateDTO
	isSet bool
}

func (v NullablePredicateDTO) Get() *PredicateDTO {
	return v.value
}

func (v *NullablePredicateDTO) Set(val *PredicateDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePredicateDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePredicateDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredicateDTO(val *PredicateDTO) *NullablePredicateDTO {
	return &NullablePredicateDTO{value: val, isSet: true}
}

func (v NullablePredicateDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredicateDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


