/*
YCloud API

The [YCloud](https://ycloud.com) API is organized around [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). Our API is designed to have predictable, resource-oriented URLs, return [JSON](https://www.json.org) responses, and use standard HTTP response codes and verbs.

API version: v2
Contact: service@ycloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ycloud

import (
	"encoding/json"
)

// Error struct for Error
type Error struct {
	// HTTP status code, [RFC 7231, Section 6](https://datatracker.ietf.org/doc/html/rfc7231#section-6).
	Status int32 `json:"status"`
	// One of a server-defined set of error codes, which could be handled programmatically.
	Code string `json:"code"`
	// A human-readable representation of the error. It is intended as an aid to developers and is not suitable for exposure to end users.
	Message *string `json:"message,omitempty"`
	// The target of the error.
	Target *string `json:"target,omitempty"`
	// A URL to more information about the error.
	DocUrl *string `json:"docUrl,omitempty"`
}

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError(status int32, code string) *Error {
	this := Error{}
	this.Status = status
	this.Code = code
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetStatus returns the Status field value
func (o *Error) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Error) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Error) SetStatus(v int32) {
	o.Status = v
}

// GetCode returns the Code field value
func (o *Error) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Error) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Error) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Error) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Error) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Error) SetMessage(v string) {
	o.Message = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *Error) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *Error) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *Error) SetTarget(v string) {
	o.Target = &v
}

// GetDocUrl returns the DocUrl field value if set, zero value otherwise.
func (o *Error) GetDocUrl() string {
	if o == nil || o.DocUrl == nil {
		var ret string
		return ret
	}
	return *o.DocUrl
}

// GetDocUrlOk returns a tuple with the DocUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetDocUrlOk() (*string, bool) {
	if o == nil || o.DocUrl == nil {
		return nil, false
	}
	return o.DocUrl, true
}

// HasDocUrl returns a boolean if a field has been set.
func (o *Error) HasDocUrl() bool {
	if o != nil && o.DocUrl != nil {
		return true
	}

	return false
}

// SetDocUrl gets a reference to the given string and assigns it to the DocUrl field.
func (o *Error) SetDocUrl(v string) {
	o.DocUrl = &v
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.DocUrl != nil {
		toSerialize["docUrl"] = o.DocUrl
	}
	return json.Marshal(toSerialize)
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


