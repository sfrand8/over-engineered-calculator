/*
Over Engineered Calculator

This is a sample API for an over engineered calculator for a coding test

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CalculateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalculateRequest{}

// CalculateRequest struct for CalculateRequest
type CalculateRequest struct {
	Expression *string `json:"expression,omitempty"`
}

// NewCalculateRequest instantiates a new CalculateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalculateRequest() *CalculateRequest {
	this := CalculateRequest{}
	return &this
}

// NewCalculateRequestWithDefaults instantiates a new CalculateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalculateRequestWithDefaults() *CalculateRequest {
	this := CalculateRequest{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *CalculateRequest) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalculateRequest) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *CalculateRequest) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *CalculateRequest) SetExpression(v string) {
	o.Expression = &v
}

func (o CalculateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalculateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	return toSerialize, nil
}

type NullableCalculateRequest struct {
	value *CalculateRequest
	isSet bool
}

func (v NullableCalculateRequest) Get() *CalculateRequest {
	return v.value
}

func (v *NullableCalculateRequest) Set(val *CalculateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCalculateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCalculateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalculateRequest(val *CalculateRequest) *NullableCalculateRequest {
	return &NullableCalculateRequest{value: val, isSet: true}
}

func (v NullableCalculateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalculateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


