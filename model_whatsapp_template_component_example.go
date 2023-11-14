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

// WhatsappTemplateComponentExample **Required** when: - `type` is `HEADER`, and `format` is one of `IMAGE`, `VIDEO`, or `DOCUMENT`. Provide a sample media URL in `header_url`. - `type` is `HEADER`, `format` is `TEXT`, and a variable is used in `text`. Provide a sample value for that variable in `header_text`. There can be at most 1 variable in `HEADER` text. - `type` is `BODY`, and variables are used in `text`. Provide sample values for those variables in `body_text`.
type WhatsappTemplateComponentExample struct {
	// Sample values for variables in `text` of a `BODY` component.
	BodyText [][]string `json:"body_text,omitempty"`
	// Sample value for the variable in `text` of a `HEADER` component.
	HeaderText []string `json:"header_text,omitempty"`
	// Sample media URL for a `HEADER` component whose format is one of `IMAGE`, `VIDEO`, or `DOCUMENT`. Supported types: - For `IMAGE`, the URL must end with one of `.jpg`, `.jpeg`, or `.png`, size limit is 5MB. - For `VIDEO`, the URL must end with `.mp4`, size limit is 16MB. - For `DOCUMENT`, the URL must end with `.pdf`, size limit is 100MB.
	HeaderUrl []string `json:"header_url,omitempty"`
}

// NewWhatsappTemplateComponentExample instantiates a new WhatsappTemplateComponentExample object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappTemplateComponentExample() *WhatsappTemplateComponentExample {
	this := WhatsappTemplateComponentExample{}
	return &this
}

// NewWhatsappTemplateComponentExampleWithDefaults instantiates a new WhatsappTemplateComponentExample object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappTemplateComponentExampleWithDefaults() *WhatsappTemplateComponentExample {
	this := WhatsappTemplateComponentExample{}
	return &this
}

// GetBodyText returns the BodyText field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentExample) GetBodyText() [][]string {
	if o == nil || o.BodyText == nil {
		var ret [][]string
		return ret
	}
	return o.BodyText
}

// GetBodyTextOk returns a tuple with the BodyText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentExample) GetBodyTextOk() ([][]string, bool) {
	if o == nil || o.BodyText == nil {
		return nil, false
	}
	return o.BodyText, true
}

// HasBodyText returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentExample) HasBodyText() bool {
	if o != nil && o.BodyText != nil {
		return true
	}

	return false
}

// SetBodyText gets a reference to the given [][]string and assigns it to the BodyText field.
func (o *WhatsappTemplateComponentExample) SetBodyText(v [][]string) {
	o.BodyText = v
}

// GetHeaderText returns the HeaderText field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentExample) GetHeaderText() []string {
	if o == nil || o.HeaderText == nil {
		var ret []string
		return ret
	}
	return o.HeaderText
}

// GetHeaderTextOk returns a tuple with the HeaderText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentExample) GetHeaderTextOk() ([]string, bool) {
	if o == nil || o.HeaderText == nil {
		return nil, false
	}
	return o.HeaderText, true
}

// HasHeaderText returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentExample) HasHeaderText() bool {
	if o != nil && o.HeaderText != nil {
		return true
	}

	return false
}

// SetHeaderText gets a reference to the given []string and assigns it to the HeaderText field.
func (o *WhatsappTemplateComponentExample) SetHeaderText(v []string) {
	o.HeaderText = v
}

// GetHeaderUrl returns the HeaderUrl field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentExample) GetHeaderUrl() []string {
	if o == nil || o.HeaderUrl == nil {
		var ret []string
		return ret
	}
	return o.HeaderUrl
}

// GetHeaderUrlOk returns a tuple with the HeaderUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentExample) GetHeaderUrlOk() ([]string, bool) {
	if o == nil || o.HeaderUrl == nil {
		return nil, false
	}
	return o.HeaderUrl, true
}

// HasHeaderUrl returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentExample) HasHeaderUrl() bool {
	if o != nil && o.HeaderUrl != nil {
		return true
	}

	return false
}

// SetHeaderUrl gets a reference to the given []string and assigns it to the HeaderUrl field.
func (o *WhatsappTemplateComponentExample) SetHeaderUrl(v []string) {
	o.HeaderUrl = v
}

func (o WhatsappTemplateComponentExample) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BodyText != nil {
		toSerialize["body_text"] = o.BodyText
	}
	if o.HeaderText != nil {
		toSerialize["header_text"] = o.HeaderText
	}
	if o.HeaderUrl != nil {
		toSerialize["header_url"] = o.HeaderUrl
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappTemplateComponentExample struct {
	value *WhatsappTemplateComponentExample
	isSet bool
}

func (v NullableWhatsappTemplateComponentExample) Get() *WhatsappTemplateComponentExample {
	return v.value
}

func (v *NullableWhatsappTemplateComponentExample) Set(val *WhatsappTemplateComponentExample) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappTemplateComponentExample) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappTemplateComponentExample) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappTemplateComponentExample(val *WhatsappTemplateComponentExample) *NullableWhatsappTemplateComponentExample {
	return &NullableWhatsappTemplateComponentExample{value: val, isSet: true}
}

func (v NullableWhatsappTemplateComponentExample) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappTemplateComponentExample) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
