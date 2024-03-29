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

// WhatsappTemplateQualityRating Quality rating of WhatsApp template. One of `GREEN`, `YELLOW`, `RED`, or `UNKNOWN`. See also [Template Quality Rating](https://developers.facebook.com/docs/whatsapp/message-templates/guidelines/#quality-rating). - `GREEN`: High quality. - `YELLOW`: Medium quality. - `RED`: Low quality. - `UNKNOWN`: Unknown quality.
type WhatsappTemplateQualityRating string

// List of WhatsappTemplateQualityRating
const (
	WHATSAPPTEMPLATEQUALITYRATING_GREEN WhatsappTemplateQualityRating = "GREEN"
	WHATSAPPTEMPLATEQUALITYRATING_YELLOW WhatsappTemplateQualityRating = "YELLOW"
	WHATSAPPTEMPLATEQUALITYRATING_RED WhatsappTemplateQualityRating = "RED"
	WHATSAPPTEMPLATEQUALITYRATING_UNKNOWN WhatsappTemplateQualityRating = "UNKNOWN"
)

// All allowed values of WhatsappTemplateQualityRating enum
var AllowedWhatsappTemplateQualityRatingEnumValues = []WhatsappTemplateQualityRating{
	"GREEN",
	"YELLOW",
	"RED",
	"UNKNOWN",
}

func (v *WhatsappTemplateQualityRating) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WhatsappTemplateQualityRating(value)
	*v = enumTypeValue
	return nil
}

// NewWhatsappTemplateQualityRatingFromValue returns a pointer to a valid WhatsappTemplateQualityRating
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWhatsappTemplateQualityRatingFromValue(v string) (*WhatsappTemplateQualityRating, error) {
	ev := WhatsappTemplateQualityRating(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WhatsappTemplateQualityRating: valid values are %v", v, AllowedWhatsappTemplateQualityRatingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WhatsappTemplateQualityRating) IsValid() bool {
	for _, existing := range AllowedWhatsappTemplateQualityRatingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WhatsappTemplateQualityRating value
func (v WhatsappTemplateQualityRating) Ptr() *WhatsappTemplateQualityRating {
	return &v
}

type NullableWhatsappTemplateQualityRating struct {
	value *WhatsappTemplateQualityRating
	isSet bool
}

func (v NullableWhatsappTemplateQualityRating) Get() *WhatsappTemplateQualityRating {
	return v.value
}

func (v *NullableWhatsappTemplateQualityRating) Set(val *WhatsappTemplateQualityRating) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappTemplateQualityRating) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappTemplateQualityRating) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappTemplateQualityRating(val *WhatsappTemplateQualityRating) *NullableWhatsappTemplateQualityRating {
	return &NullableWhatsappTemplateQualityRating{value: val, isSet: true}
}

func (v NullableWhatsappTemplateQualityRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappTemplateQualityRating) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

