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

// WhatsappTemplateStatusUpdateEventEnum Used when an event happened on WhatsApp template status updates.
type WhatsappTemplateStatusUpdateEventEnum string

// List of WhatsappTemplateStatusUpdateEventEnum
const (
	WHATSAPPTEMPLATESTATUSUPDATEEVENTENUM_PENDING WhatsappTemplateStatusUpdateEventEnum = "PENDING"
	WHATSAPPTEMPLATESTATUSUPDATEEVENTENUM_APPROVED WhatsappTemplateStatusUpdateEventEnum = "APPROVED"
	WHATSAPPTEMPLATESTATUSUPDATEEVENTENUM_REJECTED WhatsappTemplateStatusUpdateEventEnum = "REJECTED"
	WHATSAPPTEMPLATESTATUSUPDATEEVENTENUM_IN_APPEAL WhatsappTemplateStatusUpdateEventEnum = "IN_APPEAL"
	WHATSAPPTEMPLATESTATUSUPDATEEVENTENUM_PAUSED WhatsappTemplateStatusUpdateEventEnum = "PAUSED"
	WHATSAPPTEMPLATESTATUSUPDATEEVENTENUM_FLAGGED WhatsappTemplateStatusUpdateEventEnum = "FLAGGED"
	WHATSAPPTEMPLATESTATUSUPDATEEVENTENUM_DISABLED WhatsappTemplateStatusUpdateEventEnum = "DISABLED"
	WHATSAPPTEMPLATESTATUSUPDATEEVENTENUM_REINSTATED WhatsappTemplateStatusUpdateEventEnum = "REINSTATED"
	WHATSAPPTEMPLATESTATUSUPDATEEVENTENUM_PENDING_DELETION WhatsappTemplateStatusUpdateEventEnum = "PENDING_DELETION"
)

// All allowed values of WhatsappTemplateStatusUpdateEventEnum enum
var AllowedWhatsappTemplateStatusUpdateEventEnumEnumValues = []WhatsappTemplateStatusUpdateEventEnum{
	"PENDING",
	"APPROVED",
	"REJECTED",
	"IN_APPEAL",
	"PAUSED",
	"FLAGGED",
	"DISABLED",
	"REINSTATED",
	"PENDING_DELETION",
}

func (v *WhatsappTemplateStatusUpdateEventEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WhatsappTemplateStatusUpdateEventEnum(value)
	*v = enumTypeValue
	return nil
}

// NewWhatsappTemplateStatusUpdateEventEnumFromValue returns a pointer to a valid WhatsappTemplateStatusUpdateEventEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWhatsappTemplateStatusUpdateEventEnumFromValue(v string) (*WhatsappTemplateStatusUpdateEventEnum, error) {
	ev := WhatsappTemplateStatusUpdateEventEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WhatsappTemplateStatusUpdateEventEnum: valid values are %v", v, AllowedWhatsappTemplateStatusUpdateEventEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WhatsappTemplateStatusUpdateEventEnum) IsValid() bool {
	for _, existing := range AllowedWhatsappTemplateStatusUpdateEventEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WhatsappTemplateStatusUpdateEventEnum value
func (v WhatsappTemplateStatusUpdateEventEnum) Ptr() *WhatsappTemplateStatusUpdateEventEnum {
	return &v
}

type NullableWhatsappTemplateStatusUpdateEventEnum struct {
	value *WhatsappTemplateStatusUpdateEventEnum
	isSet bool
}

func (v NullableWhatsappTemplateStatusUpdateEventEnum) Get() *WhatsappTemplateStatusUpdateEventEnum {
	return v.value
}

func (v *NullableWhatsappTemplateStatusUpdateEventEnum) Set(val *WhatsappTemplateStatusUpdateEventEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappTemplateStatusUpdateEventEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappTemplateStatusUpdateEventEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappTemplateStatusUpdateEventEnum(val *WhatsappTemplateStatusUpdateEventEnum) *NullableWhatsappTemplateStatusUpdateEventEnum {
	return &NullableWhatsappTemplateStatusUpdateEventEnum{value: val, isSet: true}
}

func (v NullableWhatsappTemplateStatusUpdateEventEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappTemplateStatusUpdateEventEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
