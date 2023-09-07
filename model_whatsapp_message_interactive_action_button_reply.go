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

// WhatsappMessageInteractiveActionButtonReply struct for WhatsappMessageInteractiveActionButtonReply
type WhatsappMessageInteractiveActionButtonReply struct {
	// Button title. It cannot be an empty string and must be unique within the message. Emojis are supported, markdown is not. Maximum length: 20 characters.
	Title *string `json:"title,omitempty"`
	// Unique identifier for your button. This ID is returned in the webhook when the button is clicked by the user. Maximum length: 256 characters. You cannot have leading or trailing spaces when setting the ID.
	Id *string `json:"id,omitempty"`
}

// NewWhatsappMessageInteractiveActionButtonReply instantiates a new WhatsappMessageInteractiveActionButtonReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageInteractiveActionButtonReply() *WhatsappMessageInteractiveActionButtonReply {
	this := WhatsappMessageInteractiveActionButtonReply{}
	return &this
}

// NewWhatsappMessageInteractiveActionButtonReplyWithDefaults instantiates a new WhatsappMessageInteractiveActionButtonReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageInteractiveActionButtonReplyWithDefaults() *WhatsappMessageInteractiveActionButtonReply {
	this := WhatsappMessageInteractiveActionButtonReply{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WhatsappMessageInteractiveActionButtonReply) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageInteractiveActionButtonReply) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WhatsappMessageInteractiveActionButtonReply) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WhatsappMessageInteractiveActionButtonReply) SetTitle(v string) {
	o.Title = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WhatsappMessageInteractiveActionButtonReply) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageInteractiveActionButtonReply) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WhatsappMessageInteractiveActionButtonReply) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WhatsappMessageInteractiveActionButtonReply) SetId(v string) {
	o.Id = &v
}

func (o WhatsappMessageInteractiveActionButtonReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageInteractiveActionButtonReply struct {
	value *WhatsappMessageInteractiveActionButtonReply
	isSet bool
}

func (v NullableWhatsappMessageInteractiveActionButtonReply) Get() *WhatsappMessageInteractiveActionButtonReply {
	return v.value
}

func (v *NullableWhatsappMessageInteractiveActionButtonReply) Set(val *WhatsappMessageInteractiveActionButtonReply) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageInteractiveActionButtonReply) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageInteractiveActionButtonReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageInteractiveActionButtonReply(val *WhatsappMessageInteractiveActionButtonReply) *NullableWhatsappMessageInteractiveActionButtonReply {
	return &NullableWhatsappMessageInteractiveActionButtonReply{value: val, isSet: true}
}

func (v NullableWhatsappMessageInteractiveActionButtonReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageInteractiveActionButtonReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


