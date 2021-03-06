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

// GenericInboundCI struct for GenericInboundCI
type GenericInboundCI struct {
	TempID *string `json:"tempID,omitempty"`
	IdMethod *GenericInboundCIIdMethod `json:"idMethod,omitempty"`
	SameTempIDHandling *SameTempIDHandling `json:"sameTempIDHandling,omitempty"`
	SameTargetCIHandling *SameTargetCIHandling `json:"sameTargetCIHandling,omitempty"`
	NoFoundTargetCIHandling *NoFoundTargetCIHandling `json:"noFoundTargetCIHandling,omitempty"`
	Attributes []GenericInboundAttribute `json:"attributes,omitempty"`
}

// NewGenericInboundCI instantiates a new GenericInboundCI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericInboundCI() *GenericInboundCI {
	this := GenericInboundCI{}
	return &this
}

// NewGenericInboundCIWithDefaults instantiates a new GenericInboundCI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericInboundCIWithDefaults() *GenericInboundCI {
	this := GenericInboundCI{}
	return &this
}

// GetTempID returns the TempID field value if set, zero value otherwise.
func (o *GenericInboundCI) GetTempID() string {
	if o == nil || o.TempID == nil {
		var ret string
		return ret
	}
	return *o.TempID
}

// GetTempIDOk returns a tuple with the TempID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericInboundCI) GetTempIDOk() (*string, bool) {
	if o == nil || o.TempID == nil {
		return nil, false
	}
	return o.TempID, true
}

// HasTempID returns a boolean if a field has been set.
func (o *GenericInboundCI) HasTempID() bool {
	if o != nil && o.TempID != nil {
		return true
	}

	return false
}

// SetTempID gets a reference to the given string and assigns it to the TempID field.
func (o *GenericInboundCI) SetTempID(v string) {
	o.TempID = &v
}

// GetIdMethod returns the IdMethod field value if set, zero value otherwise.
func (o *GenericInboundCI) GetIdMethod() GenericInboundCIIdMethod {
	if o == nil || o.IdMethod == nil {
		var ret GenericInboundCIIdMethod
		return ret
	}
	return *o.IdMethod
}

// GetIdMethodOk returns a tuple with the IdMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericInboundCI) GetIdMethodOk() (*GenericInboundCIIdMethod, bool) {
	if o == nil || o.IdMethod == nil {
		return nil, false
	}
	return o.IdMethod, true
}

// HasIdMethod returns a boolean if a field has been set.
func (o *GenericInboundCI) HasIdMethod() bool {
	if o != nil && o.IdMethod != nil {
		return true
	}

	return false
}

// SetIdMethod gets a reference to the given GenericInboundCIIdMethod and assigns it to the IdMethod field.
func (o *GenericInboundCI) SetIdMethod(v GenericInboundCIIdMethod) {
	o.IdMethod = &v
}

// GetSameTempIDHandling returns the SameTempIDHandling field value if set, zero value otherwise.
func (o *GenericInboundCI) GetSameTempIDHandling() SameTempIDHandling {
	if o == nil || o.SameTempIDHandling == nil {
		var ret SameTempIDHandling
		return ret
	}
	return *o.SameTempIDHandling
}

// GetSameTempIDHandlingOk returns a tuple with the SameTempIDHandling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericInboundCI) GetSameTempIDHandlingOk() (*SameTempIDHandling, bool) {
	if o == nil || o.SameTempIDHandling == nil {
		return nil, false
	}
	return o.SameTempIDHandling, true
}

// HasSameTempIDHandling returns a boolean if a field has been set.
func (o *GenericInboundCI) HasSameTempIDHandling() bool {
	if o != nil && o.SameTempIDHandling != nil {
		return true
	}

	return false
}

// SetSameTempIDHandling gets a reference to the given SameTempIDHandling and assigns it to the SameTempIDHandling field.
func (o *GenericInboundCI) SetSameTempIDHandling(v SameTempIDHandling) {
	o.SameTempIDHandling = &v
}

// GetSameTargetCIHandling returns the SameTargetCIHandling field value if set, zero value otherwise.
func (o *GenericInboundCI) GetSameTargetCIHandling() SameTargetCIHandling {
	if o == nil || o.SameTargetCIHandling == nil {
		var ret SameTargetCIHandling
		return ret
	}
	return *o.SameTargetCIHandling
}

// GetSameTargetCIHandlingOk returns a tuple with the SameTargetCIHandling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericInboundCI) GetSameTargetCIHandlingOk() (*SameTargetCIHandling, bool) {
	if o == nil || o.SameTargetCIHandling == nil {
		return nil, false
	}
	return o.SameTargetCIHandling, true
}

// HasSameTargetCIHandling returns a boolean if a field has been set.
func (o *GenericInboundCI) HasSameTargetCIHandling() bool {
	if o != nil && o.SameTargetCIHandling != nil {
		return true
	}

	return false
}

// SetSameTargetCIHandling gets a reference to the given SameTargetCIHandling and assigns it to the SameTargetCIHandling field.
func (o *GenericInboundCI) SetSameTargetCIHandling(v SameTargetCIHandling) {
	o.SameTargetCIHandling = &v
}

// GetNoFoundTargetCIHandling returns the NoFoundTargetCIHandling field value if set, zero value otherwise.
func (o *GenericInboundCI) GetNoFoundTargetCIHandling() NoFoundTargetCIHandling {
	if o == nil || o.NoFoundTargetCIHandling == nil {
		var ret NoFoundTargetCIHandling
		return ret
	}
	return *o.NoFoundTargetCIHandling
}

// GetNoFoundTargetCIHandlingOk returns a tuple with the NoFoundTargetCIHandling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericInboundCI) GetNoFoundTargetCIHandlingOk() (*NoFoundTargetCIHandling, bool) {
	if o == nil || o.NoFoundTargetCIHandling == nil {
		return nil, false
	}
	return o.NoFoundTargetCIHandling, true
}

// HasNoFoundTargetCIHandling returns a boolean if a field has been set.
func (o *GenericInboundCI) HasNoFoundTargetCIHandling() bool {
	if o != nil && o.NoFoundTargetCIHandling != nil {
		return true
	}

	return false
}

// SetNoFoundTargetCIHandling gets a reference to the given NoFoundTargetCIHandling and assigns it to the NoFoundTargetCIHandling field.
func (o *GenericInboundCI) SetNoFoundTargetCIHandling(v NoFoundTargetCIHandling) {
	o.NoFoundTargetCIHandling = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GenericInboundCI) GetAttributes() []GenericInboundAttribute {
	if o == nil || o.Attributes == nil {
		var ret []GenericInboundAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericInboundCI) GetAttributesOk() ([]GenericInboundAttribute, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GenericInboundCI) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []GenericInboundAttribute and assigns it to the Attributes field.
func (o *GenericInboundCI) SetAttributes(v []GenericInboundAttribute) {
	o.Attributes = v
}

func (o GenericInboundCI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TempID != nil {
		toSerialize["tempID"] = o.TempID
	}
	if o.IdMethod != nil {
		toSerialize["idMethod"] = o.IdMethod
	}
	if o.SameTempIDHandling != nil {
		toSerialize["sameTempIDHandling"] = o.SameTempIDHandling
	}
	if o.SameTargetCIHandling != nil {
		toSerialize["sameTargetCIHandling"] = o.SameTargetCIHandling
	}
	if o.NoFoundTargetCIHandling != nil {
		toSerialize["noFoundTargetCIHandling"] = o.NoFoundTargetCIHandling
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableGenericInboundCI struct {
	value *GenericInboundCI
	isSet bool
}

func (v NullableGenericInboundCI) Get() *GenericInboundCI {
	return v.value
}

func (v *NullableGenericInboundCI) Set(val *GenericInboundCI) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericInboundCI) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericInboundCI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericInboundCI(val *GenericInboundCI) *NullableGenericInboundCI {
	return &NullableGenericInboundCI{value: val, isSet: true}
}

func (v NullableGenericInboundCI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericInboundCI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


