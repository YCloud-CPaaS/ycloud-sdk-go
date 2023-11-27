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

// WhatsappMessageTemplateComponentParameterAction Required if template uses catalog or MPM (multi-product message) buttons.
type WhatsappMessageTemplateComponentParameterAction struct {
	// **Optional.** Use for catalog and MPM template messages. Item SKU number. Labeled as Content ID in the Commerce Manager. The thumbnail of this item will be used as the message's header image. If the `parameters` object is omitted, the product image of the first item in your catalog will be used.
	ThumbnailProductRetailerId *string `json:"thumbnail_product_retailer_id,omitempty"`
	// Use for MPM templates. Product sections. You can define up to 10 sections.
	Sections []WhatsappMessageTemplateComponentParameterActionSection `json:"sections,omitempty"`
}

// NewWhatsappMessageTemplateComponentParameterAction instantiates a new WhatsappMessageTemplateComponentParameterAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageTemplateComponentParameterAction() *WhatsappMessageTemplateComponentParameterAction {
	this := WhatsappMessageTemplateComponentParameterAction{}
	return &this
}

// NewWhatsappMessageTemplateComponentParameterActionWithDefaults instantiates a new WhatsappMessageTemplateComponentParameterAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageTemplateComponentParameterActionWithDefaults() *WhatsappMessageTemplateComponentParameterAction {
	this := WhatsappMessageTemplateComponentParameterAction{}
	return &this
}

// GetThumbnailProductRetailerId returns the ThumbnailProductRetailerId field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameterAction) GetThumbnailProductRetailerId() string {
	if o == nil || o.ThumbnailProductRetailerId == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailProductRetailerId
}

// GetThumbnailProductRetailerIdOk returns a tuple with the ThumbnailProductRetailerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameterAction) GetThumbnailProductRetailerIdOk() (*string, bool) {
	if o == nil || o.ThumbnailProductRetailerId == nil {
		return nil, false
	}
	return o.ThumbnailProductRetailerId, true
}

// HasThumbnailProductRetailerId returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameterAction) HasThumbnailProductRetailerId() bool {
	if o != nil && o.ThumbnailProductRetailerId != nil {
		return true
	}

	return false
}

// SetThumbnailProductRetailerId gets a reference to the given string and assigns it to the ThumbnailProductRetailerId field.
func (o *WhatsappMessageTemplateComponentParameterAction) SetThumbnailProductRetailerId(v string) {
	o.ThumbnailProductRetailerId = &v
}

// GetSections returns the Sections field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameterAction) GetSections() []WhatsappMessageTemplateComponentParameterActionSection {
	if o == nil || o.Sections == nil {
		var ret []WhatsappMessageTemplateComponentParameterActionSection
		return ret
	}
	return o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameterAction) GetSectionsOk() ([]WhatsappMessageTemplateComponentParameterActionSection, bool) {
	if o == nil || o.Sections == nil {
		return nil, false
	}
	return o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameterAction) HasSections() bool {
	if o != nil && o.Sections != nil {
		return true
	}

	return false
}

// SetSections gets a reference to the given []WhatsappMessageTemplateComponentParameterActionSection and assigns it to the Sections field.
func (o *WhatsappMessageTemplateComponentParameterAction) SetSections(v []WhatsappMessageTemplateComponentParameterActionSection) {
	o.Sections = v
}

func (o WhatsappMessageTemplateComponentParameterAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ThumbnailProductRetailerId != nil {
		toSerialize["thumbnail_product_retailer_id"] = o.ThumbnailProductRetailerId
	}
	if o.Sections != nil {
		toSerialize["sections"] = o.Sections
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageTemplateComponentParameterAction struct {
	value *WhatsappMessageTemplateComponentParameterAction
	isSet bool
}

func (v NullableWhatsappMessageTemplateComponentParameterAction) Get() *WhatsappMessageTemplateComponentParameterAction {
	return v.value
}

func (v *NullableWhatsappMessageTemplateComponentParameterAction) Set(val *WhatsappMessageTemplateComponentParameterAction) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageTemplateComponentParameterAction) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageTemplateComponentParameterAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageTemplateComponentParameterAction(val *WhatsappMessageTemplateComponentParameterAction) *NullableWhatsappMessageTemplateComponentParameterAction {
	return &NullableWhatsappMessageTemplateComponentParameterAction{value: val, isSet: true}
}

func (v NullableWhatsappMessageTemplateComponentParameterAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageTemplateComponentParameterAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
