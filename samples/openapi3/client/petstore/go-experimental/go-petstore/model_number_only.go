/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"bytes"
	"encoding/json"
)

// NumberOnly struct for NumberOnly
type NumberOnly struct {
	JustNumber *float32 `json:"JustNumber,omitempty"`
}

// GetJustNumber returns the JustNumber field value if set, zero value otherwise.
func (o *NumberOnly) GetJustNumber() float32 {
	if o == nil || o.JustNumber == nil {
		var ret float32
		return ret
	}
	return *o.JustNumber
}

// GetJustNumberOk returns a tuple with the JustNumber field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NumberOnly) GetJustNumberOk() (float32, bool) {
	if o == nil || o.JustNumber == nil {
		var ret float32
		return ret, false
	}
	return *o.JustNumber, true
}

// HasJustNumber returns a boolean if a field has been set.
func (o *NumberOnly) HasJustNumber() bool {
	if o != nil && o.JustNumber != nil {
		return true
	}

	return false
}

// SetJustNumber gets a reference to the given float32 and assigns it to the JustNumber field.
func (o *NumberOnly) SetJustNumber(v float32) {
	o.JustNumber = &v
}

type NullableNumberOnly struct {
	Value NumberOnly
	ExplicitNull bool
}

func (v NullableNumberOnly) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNumberOnly) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}