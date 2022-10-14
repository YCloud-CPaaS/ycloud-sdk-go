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

// VerificationChannel Supports several independent channels for verification: `sms`: Sends an SMS message with a verification code. `voice`: Makes a voice call with a verification code. `email_code`: Sends an email with a verification code. `whatsapp`: Sends a WhatsApp message with a verification code.
type VerificationChannel string

// List of VerificationChannel
const (
	VERIFICATIONCHANNEL_SMS VerificationChannel = "sms"
	VERIFICATIONCHANNEL_VOICE VerificationChannel = "voice"
	VERIFICATIONCHANNEL_EMAIL_CODE VerificationChannel = "email_code"
	VERIFICATIONCHANNEL_WHATSAPP VerificationChannel = "whatsapp"
)

// All allowed values of VerificationChannel enum
var AllowedVerificationChannelEnumValues = []VerificationChannel{
	"sms",
	"voice",
	"email_code",
	"whatsapp",
}

func (v *VerificationChannel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VerificationChannel(value)
	*v = enumTypeValue
	return nil
}

// NewVerificationChannelFromValue returns a pointer to a valid VerificationChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVerificationChannelFromValue(v string) (*VerificationChannel, error) {
	ev := VerificationChannel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VerificationChannel: valid values are %v", v, AllowedVerificationChannelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VerificationChannel) IsValid() bool {
	for _, existing := range AllowedVerificationChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VerificationChannel value
func (v VerificationChannel) Ptr() *VerificationChannel {
	return &v
}

type NullableVerificationChannel struct {
	value *VerificationChannel
	isSet bool
}

func (v NullableVerificationChannel) Get() *VerificationChannel {
	return v.value
}

func (v *NullableVerificationChannel) Set(val *VerificationChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationChannel(val *VerificationChannel) *NullableVerificationChannel {
	return &NullableVerificationChannel{value: val, isSet: true}
}

func (v NullableVerificationChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
