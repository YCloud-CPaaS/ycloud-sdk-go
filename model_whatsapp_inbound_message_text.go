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

// WhatsappInboundMessageText When the notification describes a text message, the text object provides the body of the text message.
type WhatsappInboundMessageText struct {
	// Message text.
	Body *string `json:"body,omitempty"`
}

// NewWhatsappInboundMessageText instantiates a new WhatsappInboundMessageText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappInboundMessageText() *WhatsappInboundMessageText {
	this := WhatsappInboundMessageText{}
	return &this
}

// NewWhatsappInboundMessageTextWithDefaults instantiates a new WhatsappInboundMessageText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappInboundMessageTextWithDefaults() *WhatsappInboundMessageText {
	this := WhatsappInboundMessageText{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *WhatsappInboundMessageText) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageText) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *WhatsappInboundMessageText) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *WhatsappInboundMessageText) SetBody(v string) {
	o.Body = &v
}

func (o WhatsappInboundMessageText) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappInboundMessageText struct {
	value *WhatsappInboundMessageText
	isSet bool
}

func (v NullableWhatsappInboundMessageText) Get() *WhatsappInboundMessageText {
	return v.value
}

func (v *NullableWhatsappInboundMessageText) Set(val *WhatsappInboundMessageText) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappInboundMessageText) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappInboundMessageText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappInboundMessageText(val *WhatsappInboundMessageText) *NullableWhatsappInboundMessageText {
	return &NullableWhatsappInboundMessageText{value: val, isSet: true}
}

func (v NullableWhatsappInboundMessageText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappInboundMessageText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


