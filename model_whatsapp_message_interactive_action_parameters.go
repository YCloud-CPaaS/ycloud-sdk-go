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

// WhatsappMessageInteractiveActionParameters Action parameters. Required for Call-To-Action (CTA) URL Button Messages.
type WhatsappMessageInteractiveActionParameters struct {
	// Text of the CTA URL button. Maximum length: 20 bytes.
	DisplayText *string `json:"display_text,omitempty"`
	// URL of the CTA URL button.
	Url *string `json:"url,omitempty"`
	// Item SKU number. Labeled as **Content ID** in the [Commerce Manager](https://business.facebook.com/commerce/). The thumbnail of this item will be used as the message's header image.
	ThumbnailProductRetailerId *string `json:"thumbnail_product_retailer_id,omitempty"`
}

// NewWhatsappMessageInteractiveActionParameters instantiates a new WhatsappMessageInteractiveActionParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageInteractiveActionParameters() *WhatsappMessageInteractiveActionParameters {
	this := WhatsappMessageInteractiveActionParameters{}
	return &this
}

// NewWhatsappMessageInteractiveActionParametersWithDefaults instantiates a new WhatsappMessageInteractiveActionParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageInteractiveActionParametersWithDefaults() *WhatsappMessageInteractiveActionParameters {
	this := WhatsappMessageInteractiveActionParameters{}
	return &this
}

// GetDisplayText returns the DisplayText field value if set, zero value otherwise.
func (o *WhatsappMessageInteractiveActionParameters) GetDisplayText() string {
	if o == nil || o.DisplayText == nil {
		var ret string
		return ret
	}
	return *o.DisplayText
}

// GetDisplayTextOk returns a tuple with the DisplayText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageInteractiveActionParameters) GetDisplayTextOk() (*string, bool) {
	if o == nil || o.DisplayText == nil {
		return nil, false
	}
	return o.DisplayText, true
}

// HasDisplayText returns a boolean if a field has been set.
func (o *WhatsappMessageInteractiveActionParameters) HasDisplayText() bool {
	if o != nil && o.DisplayText != nil {
		return true
	}

	return false
}

// SetDisplayText gets a reference to the given string and assigns it to the DisplayText field.
func (o *WhatsappMessageInteractiveActionParameters) SetDisplayText(v string) {
	o.DisplayText = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WhatsappMessageInteractiveActionParameters) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageInteractiveActionParameters) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WhatsappMessageInteractiveActionParameters) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WhatsappMessageInteractiveActionParameters) SetUrl(v string) {
	o.Url = &v
}

// GetThumbnailProductRetailerId returns the ThumbnailProductRetailerId field value if set, zero value otherwise.
func (o *WhatsappMessageInteractiveActionParameters) GetThumbnailProductRetailerId() string {
	if o == nil || o.ThumbnailProductRetailerId == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailProductRetailerId
}

// GetThumbnailProductRetailerIdOk returns a tuple with the ThumbnailProductRetailerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageInteractiveActionParameters) GetThumbnailProductRetailerIdOk() (*string, bool) {
	if o == nil || o.ThumbnailProductRetailerId == nil {
		return nil, false
	}
	return o.ThumbnailProductRetailerId, true
}

// HasThumbnailProductRetailerId returns a boolean if a field has been set.
func (o *WhatsappMessageInteractiveActionParameters) HasThumbnailProductRetailerId() bool {
	if o != nil && o.ThumbnailProductRetailerId != nil {
		return true
	}

	return false
}

// SetThumbnailProductRetailerId gets a reference to the given string and assigns it to the ThumbnailProductRetailerId field.
func (o *WhatsappMessageInteractiveActionParameters) SetThumbnailProductRetailerId(v string) {
	o.ThumbnailProductRetailerId = &v
}

func (o WhatsappMessageInteractiveActionParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayText != nil {
		toSerialize["display_text"] = o.DisplayText
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.ThumbnailProductRetailerId != nil {
		toSerialize["thumbnail_product_retailer_id"] = o.ThumbnailProductRetailerId
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageInteractiveActionParameters struct {
	value *WhatsappMessageInteractiveActionParameters
	isSet bool
}

func (v NullableWhatsappMessageInteractiveActionParameters) Get() *WhatsappMessageInteractiveActionParameters {
	return v.value
}

func (v *NullableWhatsappMessageInteractiveActionParameters) Set(val *WhatsappMessageInteractiveActionParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageInteractiveActionParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageInteractiveActionParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageInteractiveActionParameters(val *WhatsappMessageInteractiveActionParameters) *NullableWhatsappMessageInteractiveActionParameters {
	return &NullableWhatsappMessageInteractiveActionParameters{value: val, isSet: true}
}

func (v NullableWhatsappMessageInteractiveActionParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageInteractiveActionParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
