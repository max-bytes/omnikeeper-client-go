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

// ChangeDataRequest struct for ChangeDataRequest
type ChangeDataRequest struct {
	SparseRows []SparseRow `json:"sparseRows,omitempty"`
}

// NewChangeDataRequest instantiates a new ChangeDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeDataRequest() *ChangeDataRequest {
	this := ChangeDataRequest{}
	return &this
}

// NewChangeDataRequestWithDefaults instantiates a new ChangeDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeDataRequestWithDefaults() *ChangeDataRequest {
	this := ChangeDataRequest{}
	return &this
}

// GetSparseRows returns the SparseRows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeDataRequest) GetSparseRows() []SparseRow {
	if o == nil  {
		var ret []SparseRow
		return ret
	}
	return o.SparseRows
}

// GetSparseRowsOk returns a tuple with the SparseRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeDataRequest) GetSparseRowsOk() (*[]SparseRow, bool) {
	if o == nil || o.SparseRows == nil {
		return nil, false
	}
	return &o.SparseRows, true
}

// HasSparseRows returns a boolean if a field has been set.
func (o *ChangeDataRequest) HasSparseRows() bool {
	if o != nil && o.SparseRows != nil {
		return true
	}

	return false
}

// SetSparseRows gets a reference to the given []SparseRow and assigns it to the SparseRows field.
func (o *ChangeDataRequest) SetSparseRows(v []SparseRow) {
	o.SparseRows = v
}

func (o ChangeDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SparseRows != nil {
		toSerialize["sparseRows"] = o.SparseRows
	}
	return json.Marshal(toSerialize)
}

type NullableChangeDataRequest struct {
	value *ChangeDataRequest
	isSet bool
}

func (v NullableChangeDataRequest) Get() *ChangeDataRequest {
	return v.value
}

func (v *NullableChangeDataRequest) Set(val *ChangeDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeDataRequest(val *ChangeDataRequest) *NullableChangeDataRequest {
	return &NullableChangeDataRequest{value: val, isSet: true}
}

func (v NullableChangeDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


