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

// WhatsappTemplateCreateRequest See [WhatsApp Templates](https://developers.facebook.com/docs/whatsapp/business-management-api/message-templates).
type WhatsappTemplateCreateRequest struct {
	// WhatsApp Business Account ID.
	WabaId string `json:"wabaId"`
	// Name of the template.
	Name string `json:"name"`
	// Language code of the template. See [Supported Languages](https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#supported-languages) for all codes.
	Language string `json:"language"`
	Category WhatsappTemplateCategory `json:"category"`
	Components []WhatsappTemplateComponent `json:"components"`
}

// NewWhatsappTemplateCreateRequest instantiates a new WhatsappTemplateCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappTemplateCreateRequest(wabaId string, name string, language string, category WhatsappTemplateCategory, components []WhatsappTemplateComponent) *WhatsappTemplateCreateRequest {
	this := WhatsappTemplateCreateRequest{}
	this.WabaId = wabaId
	this.Name = name
	this.Language = language
	this.Category = category
	this.Components = components
	return &this
}

// NewWhatsappTemplateCreateRequestWithDefaults instantiates a new WhatsappTemplateCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappTemplateCreateRequestWithDefaults() *WhatsappTemplateCreateRequest {
	this := WhatsappTemplateCreateRequest{}
	return &this
}

// GetWabaId returns the WabaId field value
func (o *WhatsappTemplateCreateRequest) GetWabaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WabaId
}

// GetWabaIdOk returns a tuple with the WabaId field value
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateCreateRequest) GetWabaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WabaId, true
}

// SetWabaId sets field value
func (o *WhatsappTemplateCreateRequest) SetWabaId(v string) {
	o.WabaId = v
}

// GetName returns the Name field value
func (o *WhatsappTemplateCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WhatsappTemplateCreateRequest) SetName(v string) {
	o.Name = v
}

// GetLanguage returns the Language field value
func (o *WhatsappTemplateCreateRequest) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateCreateRequest) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *WhatsappTemplateCreateRequest) SetLanguage(v string) {
	o.Language = v
}

// GetCategory returns the Category field value
func (o *WhatsappTemplateCreateRequest) GetCategory() WhatsappTemplateCategory {
	if o == nil {
		var ret WhatsappTemplateCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateCreateRequest) GetCategoryOk() (*WhatsappTemplateCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *WhatsappTemplateCreateRequest) SetCategory(v WhatsappTemplateCategory) {
	o.Category = v
}

// GetComponents returns the Components field value
func (o *WhatsappTemplateCreateRequest) GetComponents() []WhatsappTemplateComponent {
	if o == nil {
		var ret []WhatsappTemplateComponent
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateCreateRequest) GetComponentsOk() ([]WhatsappTemplateComponent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *WhatsappTemplateCreateRequest) SetComponents(v []WhatsappTemplateComponent) {
	o.Components = v
}

func (o WhatsappTemplateCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["wabaId"] = o.WabaId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["language"] = o.Language
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["components"] = o.Components
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappTemplateCreateRequest struct {
	value *WhatsappTemplateCreateRequest
	isSet bool
}

func (v NullableWhatsappTemplateCreateRequest) Get() *WhatsappTemplateCreateRequest {
	return v.value
}

func (v *NullableWhatsappTemplateCreateRequest) Set(val *WhatsappTemplateCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappTemplateCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappTemplateCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappTemplateCreateRequest(val *WhatsappTemplateCreateRequest) *NullableWhatsappTemplateCreateRequest {
	return &NullableWhatsappTemplateCreateRequest{value: val, isSet: true}
}

func (v NullableWhatsappTemplateCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappTemplateCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


