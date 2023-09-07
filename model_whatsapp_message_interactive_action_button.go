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

// WhatsappMessageInteractiveActionButton A button object in `interactive` messages.
type WhatsappMessageInteractiveActionButton struct {
	// Only supported type is `reply` (for Reply Button).
	Type *string `json:"type,omitempty"`
	Reply *WhatsappMessageInteractiveActionButtonReply `json:"reply,omitempty"`
}

// NewWhatsappMessageInteractiveActionButton instantiates a new WhatsappMessageInteractiveActionButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageInteractiveActionButton() *WhatsappMessageInteractiveActionButton {
	this := WhatsappMessageInteractiveActionButton{}
	return &this
}

// NewWhatsappMessageInteractiveActionButtonWithDefaults instantiates a new WhatsappMessageInteractiveActionButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageInteractiveActionButtonWithDefaults() *WhatsappMessageInteractiveActionButton {
	this := WhatsappMessageInteractiveActionButton{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WhatsappMessageInteractiveActionButton) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageInteractiveActionButton) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WhatsappMessageInteractiveActionButton) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WhatsappMessageInteractiveActionButton) SetType(v string) {
	o.Type = &v
}

// GetReply returns the Reply field value if set, zero value otherwise.
func (o *WhatsappMessageInteractiveActionButton) GetReply() WhatsappMessageInteractiveActionButtonReply {
	if o == nil || o.Reply == nil {
		var ret WhatsappMessageInteractiveActionButtonReply
		return ret
	}
	return *o.Reply
}

// GetReplyOk returns a tuple with the Reply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageInteractiveActionButton) GetReplyOk() (*WhatsappMessageInteractiveActionButtonReply, bool) {
	if o == nil || o.Reply == nil {
		return nil, false
	}
	return o.Reply, true
}

// HasReply returns a boolean if a field has been set.
func (o *WhatsappMessageInteractiveActionButton) HasReply() bool {
	if o != nil && o.Reply != nil {
		return true
	}

	return false
}

// SetReply gets a reference to the given WhatsappMessageInteractiveActionButtonReply and assigns it to the Reply field.
func (o *WhatsappMessageInteractiveActionButton) SetReply(v WhatsappMessageInteractiveActionButtonReply) {
	o.Reply = &v
}

func (o WhatsappMessageInteractiveActionButton) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Reply != nil {
		toSerialize["reply"] = o.Reply
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageInteractiveActionButton struct {
	value *WhatsappMessageInteractiveActionButton
	isSet bool
}

func (v NullableWhatsappMessageInteractiveActionButton) Get() *WhatsappMessageInteractiveActionButton {
	return v.value
}

func (v *NullableWhatsappMessageInteractiveActionButton) Set(val *WhatsappMessageInteractiveActionButton) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageInteractiveActionButton) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageInteractiveActionButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageInteractiveActionButton(val *WhatsappMessageInteractiveActionButton) *NullableWhatsappMessageInteractiveActionButton {
	return &NullableWhatsappMessageInteractiveActionButton{value: val, isSet: true}
}

func (v NullableWhatsappMessageInteractiveActionButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageInteractiveActionButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


