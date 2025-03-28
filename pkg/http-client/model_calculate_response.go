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

// checks if the CalculateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalculateResponse{}

// CalculateResponse struct for CalculateResponse
type CalculateResponse struct {
	// @Description The result of the mathematical expression evaluation @Example \"42.0\" // Optional: Provide an example result
	Result *string `json:"result,omitempty"`
}

// NewCalculateResponse instantiates a new CalculateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalculateResponse() *CalculateResponse {
	this := CalculateResponse{}
	return &this
}

// NewCalculateResponseWithDefaults instantiates a new CalculateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalculateResponseWithDefaults() *CalculateResponse {
	this := CalculateResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CalculateResponse) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalculateResponse) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CalculateResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *CalculateResponse) SetResult(v string) {
	o.Result = &v
}

func (o CalculateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalculateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableCalculateResponse struct {
	value *CalculateResponse
	isSet bool
}

func (v NullableCalculateResponse) Get() *CalculateResponse {
	return v.value
}

func (v *NullableCalculateResponse) Set(val *CalculateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCalculateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCalculateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalculateResponse(val *CalculateResponse) *NullableCalculateResponse {
	return &NullableCalculateResponse{value: val, isSet: true}
}

func (v NullableCalculateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalculateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


