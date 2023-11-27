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

// EventType Type of event.
type EventType string

// List of EventType
const (
	EVENTTYPE_EMAIL_DELIVERY_UPDATED                EventType = "email.delivery.updated"
	EVENTTYPE_SMS_MESSAGE_UPDATED                   EventType = "sms.message.updated"
	EVENTTYPE_SMS_INBOUND_RECEIVED                  EventType = "sms.inbound.received"
	EVENTTYPE_VOICE_MESSAGE_UPDATED                 EventType = "voice.message.updated"
	EVENTTYPE_WHATSAPP_BUSINESS_ACCOUNT_DELETED     EventType = "whatsapp.business_account.deleted"
	EVENTTYPE_WHATSAPP_BUSINESS_ACCOUNT_REVIEWED    EventType = "whatsapp.business_account.reviewed"
	EVENTTYPE_WHATSAPP_BUSINESS_ACCOUNT_UPDATED     EventType = "whatsapp.business_account.updated"
	EVENTTYPE_WHATSAPP_INBOUND_MESSAGE_RECEIVED     EventType = "whatsapp.inbound_message.received"
	EVENTTYPE_WHATSAPP_MESSAGE_UPDATED              EventType = "whatsapp.message.updated"
	EVENTTYPE_WHATSAPP_PHONE_NUMBER_DELETED         EventType = "whatsapp.phone_number.deleted"
	EVENTTYPE_WHATSAPP_PHONE_NUMBER_NAME_UPDATED    EventType = "whatsapp.phone_number.name_updated"
	EVENTTYPE_WHATSAPP_PHONE_NUMBER_QUALITY_UPDATED EventType = "whatsapp.phone_number.quality_updated"
	EVENTTYPE_WHATSAPP_TEMPLATE_CATEGORY_UPDATED    EventType = "whatsapp.template.category_updated"
	EVENTTYPE_WHATSAPP_TEMPLATE_QUALITY_UPDATED     EventType = "whatsapp.template.quality_updated"
	EVENTTYPE_WHATSAPP_TEMPLATE_REVIEWED            EventType = "whatsapp.template.reviewed"
)

// All allowed values of EventType enum
var AllowedEventTypeEnumValues = []EventType{
	"email.delivery.updated",
	"sms.message.updated",
	"sms.inbound.received",
	"voice.message.updated",
	"whatsapp.business_account.deleted",
	"whatsapp.business_account.reviewed",
	"whatsapp.business_account.updated",
	"whatsapp.inbound_message.received",
	"whatsapp.message.updated",
	"whatsapp.phone_number.deleted",
	"whatsapp.phone_number.name_updated",
	"whatsapp.phone_number.quality_updated",
	"whatsapp.template.category_updated",
	"whatsapp.template.quality_updated",
	"whatsapp.template.reviewed",
}

func (v *EventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventType(value)
	*v = enumTypeValue
	return nil
}

// NewEventTypeFromValue returns a pointer to a valid EventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventTypeFromValue(v string) (*EventType, error) {
	ev := EventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventType: valid values are %v", v, AllowedEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventType) IsValid() bool {
	for _, existing := range AllowedEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventType value
func (v EventType) Ptr() *EventType {
	return &v
}

type NullableEventType struct {
	value *EventType
	isSet bool
}

func (v NullableEventType) Get() *EventType {
	return v.value
}

func (v *NullableEventType) Set(val *EventType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventType(val *EventType) *NullableEventType {
	return &NullableEventType{value: val, isSet: true}
}

func (v NullableEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
