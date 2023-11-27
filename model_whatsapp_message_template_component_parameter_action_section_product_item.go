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

// WhatsappMessageTemplateComponentParameterActionSectionProductItem struct for WhatsappMessageTemplateComponentParameterActionSectionProductItem
type WhatsappMessageTemplateComponentParameterActionSectionProductItem struct {
	// SKU number of the item you want to appear in the section. SKU numbers are labeled as **Content ID** in the [Commerce Manager](https://business.facebook.com/commerce/).
	ProductRetailerId *string `json:"product_retailer_id,omitempty"`
}

// NewWhatsappMessageTemplateComponentParameterActionSectionProductItem instantiates a new WhatsappMessageTemplateComponentParameterActionSectionProductItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageTemplateComponentParameterActionSectionProductItem() *WhatsappMessageTemplateComponentParameterActionSectionProductItem {
	this := WhatsappMessageTemplateComponentParameterActionSectionProductItem{}
	return &this
}

// NewWhatsappMessageTemplateComponentParameterActionSectionProductItemWithDefaults instantiates a new WhatsappMessageTemplateComponentParameterActionSectionProductItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageTemplateComponentParameterActionSectionProductItemWithDefaults() *WhatsappMessageTemplateComponentParameterActionSectionProductItem {
	this := WhatsappMessageTemplateComponentParameterActionSectionProductItem{}
	return &this
}

// GetProductRetailerId returns the ProductRetailerId field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameterActionSectionProductItem) GetProductRetailerId() string {
	if o == nil || o.ProductRetailerId == nil {
		var ret string
		return ret
	}
	return *o.ProductRetailerId
}

// GetProductRetailerIdOk returns a tuple with the ProductRetailerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameterActionSectionProductItem) GetProductRetailerIdOk() (*string, bool) {
	if o == nil || o.ProductRetailerId == nil {
		return nil, false
	}
	return o.ProductRetailerId, true
}

// HasProductRetailerId returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameterActionSectionProductItem) HasProductRetailerId() bool {
	if o != nil && o.ProductRetailerId != nil {
		return true
	}

	return false
}

// SetProductRetailerId gets a reference to the given string and assigns it to the ProductRetailerId field.
func (o *WhatsappMessageTemplateComponentParameterActionSectionProductItem) SetProductRetailerId(v string) {
	o.ProductRetailerId = &v
}

func (o WhatsappMessageTemplateComponentParameterActionSectionProductItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProductRetailerId != nil {
		toSerialize["product_retailer_id"] = o.ProductRetailerId
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageTemplateComponentParameterActionSectionProductItem struct {
	value *WhatsappMessageTemplateComponentParameterActionSectionProductItem
	isSet bool
}

func (v NullableWhatsappMessageTemplateComponentParameterActionSectionProductItem) Get() *WhatsappMessageTemplateComponentParameterActionSectionProductItem {
	return v.value
}

func (v *NullableWhatsappMessageTemplateComponentParameterActionSectionProductItem) Set(val *WhatsappMessageTemplateComponentParameterActionSectionProductItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageTemplateComponentParameterActionSectionProductItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageTemplateComponentParameterActionSectionProductItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageTemplateComponentParameterActionSectionProductItem(val *WhatsappMessageTemplateComponentParameterActionSectionProductItem) *NullableWhatsappMessageTemplateComponentParameterActionSectionProductItem {
	return &NullableWhatsappMessageTemplateComponentParameterActionSectionProductItem{value: val, isSet: true}
}

func (v NullableWhatsappMessageTemplateComponentParameterActionSectionProductItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageTemplateComponentParameterActionSectionProductItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
