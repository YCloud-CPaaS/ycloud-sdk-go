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

// WhatsappMessageContactEmailsInner Contact email address(es) formatted as an emails object.
type WhatsappMessageContactEmailsInner struct {
	// Email address.
	Email *string `json:"email,omitempty"`
	// Standard values are `HOME` and `WORK`.
	Type *string `json:"type,omitempty"`
}

// NewWhatsappMessageContactEmailsInner instantiates a new WhatsappMessageContactEmailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageContactEmailsInner() *WhatsappMessageContactEmailsInner {
	this := WhatsappMessageContactEmailsInner{}
	return &this
}

// NewWhatsappMessageContactEmailsInnerWithDefaults instantiates a new WhatsappMessageContactEmailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageContactEmailsInnerWithDefaults() *WhatsappMessageContactEmailsInner {
	this := WhatsappMessageContactEmailsInner{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *WhatsappMessageContactEmailsInner) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageContactEmailsInner) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *WhatsappMessageContactEmailsInner) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *WhatsappMessageContactEmailsInner) SetEmail(v string) {
	o.Email = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WhatsappMessageContactEmailsInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageContactEmailsInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WhatsappMessageContactEmailsInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WhatsappMessageContactEmailsInner) SetType(v string) {
	o.Type = &v
}

func (o WhatsappMessageContactEmailsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageContactEmailsInner struct {
	value *WhatsappMessageContactEmailsInner
	isSet bool
}

func (v NullableWhatsappMessageContactEmailsInner) Get() *WhatsappMessageContactEmailsInner {
	return v.value
}

func (v *NullableWhatsappMessageContactEmailsInner) Set(val *WhatsappMessageContactEmailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageContactEmailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageContactEmailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageContactEmailsInner(val *WhatsappMessageContactEmailsInner) *NullableWhatsappMessageContactEmailsInner {
	return &NullableWhatsappMessageContactEmailsInner{value: val, isSet: true}
}

func (v NullableWhatsappMessageContactEmailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageContactEmailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

