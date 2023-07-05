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
	"fmt"
)

// WhatsappConversationType Conversation type. There is a charge when the first business message of this conversation is delivered, initiating the 24-hour conversation session. As such, the conversation type can be `null` before the first message is delivered. `FREE_ENTRY`: Conversations originating from a [free entry point](https://developers.facebook.com/docs/whatsapp/pricing#free-entry-points). `FREE_TIER`: Conversations within the monthly [free tier](https://developers.facebook.com/docs/whatsapp/pricing#free-tier-conversations). `REGULAR`: Any conversations that did not originate from a [free entry point](https://developers.facebook.com/docs/whatsapp/pricing#free-entry-points) or are above the monthly [free tier](https://developers.facebook.com/docs/whatsapp/pricing#free-tier-conversations) allotment.
type WhatsappConversationType string

// List of WhatsappConversationType
const (
	WHATSAPPCONVERSATIONTYPE_FREE_ENTRY WhatsappConversationType = "FREE_ENTRY"
	WHATSAPPCONVERSATIONTYPE_FREE_TIER WhatsappConversationType = "FREE_TIER"
	WHATSAPPCONVERSATIONTYPE_REGULAR WhatsappConversationType = "REGULAR"
)

// All allowed values of WhatsappConversationType enum
var AllowedWhatsappConversationTypeEnumValues = []WhatsappConversationType{
	"FREE_ENTRY",
	"FREE_TIER",
	"REGULAR",
}

func (v *WhatsappConversationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WhatsappConversationType(value)
	*v = enumTypeValue
	return nil
}

// NewWhatsappConversationTypeFromValue returns a pointer to a valid WhatsappConversationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWhatsappConversationTypeFromValue(v string) (*WhatsappConversationType, error) {
	ev := WhatsappConversationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WhatsappConversationType: valid values are %v", v, AllowedWhatsappConversationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WhatsappConversationType) IsValid() bool {
	for _, existing := range AllowedWhatsappConversationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WhatsappConversationType value
func (v WhatsappConversationType) Ptr() *WhatsappConversationType {
	return &v
}

type NullableWhatsappConversationType struct {
	value *WhatsappConversationType
	isSet bool
}

func (v NullableWhatsappConversationType) Get() *WhatsappConversationType {
	return v.value
}

func (v *NullableWhatsappConversationType) Set(val *WhatsappConversationType) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappConversationType) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappConversationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappConversationType(val *WhatsappConversationType) *NullableWhatsappConversationType {
	return &NullableWhatsappConversationType{value: val, isSet: true}
}

func (v NullableWhatsappConversationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappConversationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

