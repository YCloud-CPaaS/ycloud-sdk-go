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
	"time"
)

// Event Represents a webhook event payload. Every event contains certain common properties: `id`, `type`, `apiVersion`, `createTime`. Each event may also contain some properties unique to the event. For example, `sms` is returned when `type` is `sms.message.updated`.
type Event struct {
	// Unique ID for the event.
	Id   string    `json:"id"`
	Type EventType `json:"type"`
	// The API version used to render this event.
	ApiVersion string `json:"apiVersion"`
	// The time at which this event was created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., `2022-06-01T12:00:00.000Z`.
	CreateTime              time.Time                `json:"createTime"`
	EmailDelivery           *EmailDelivery           `json:"emailDelivery,omitempty"`
	Sms                     *Sms                     `json:"sms,omitempty"`
	SmsInbound              *SmsInbound              `json:"smsInbound,omitempty"`
	Voice                   *Voice                   `json:"voice,omitempty"`
	WhatsappBusinessAccount *WhatsappBusinessAccount `json:"whatsappBusinessAccount,omitempty"`
	WhatsappInboundMessage  *WhatsappInboundMessage  `json:"whatsappInboundMessage,omitempty"`
	WhatsappMessage         *WhatsappMessage         `json:"whatsappMessage,omitempty"`
	WhatsappPhoneNumber     *WhatsappPhoneNumber     `json:"whatsappPhoneNumber,omitempty"`
	WhatsappTemplate        *WhatsappTemplate        `json:"whatsappTemplate,omitempty"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(id string, type_ EventType, apiVersion string, createTime time.Time) *Event {
	this := Event{}
	this.Id = id
	this.Type = type_
	this.ApiVersion = apiVersion
	this.CreateTime = createTime
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetId returns the Id field value
func (o *Event) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Event) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Event) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *Event) GetType() EventType {
	if o == nil {
		var ret EventType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Event) GetTypeOk() (*EventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Event) SetType(v EventType) {
	o.Type = v
}

// GetApiVersion returns the ApiVersion field value
func (o *Event) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *Event) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *Event) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetCreateTime returns the CreateTime field value
func (o *Event) GetCreateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value
// and a boolean to check if the value has been set.
func (o *Event) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateTime, true
}

// SetCreateTime sets field value
func (o *Event) SetCreateTime(v time.Time) {
	o.CreateTime = v
}

// GetEmailDelivery returns the EmailDelivery field value if set, zero value otherwise.
func (o *Event) GetEmailDelivery() EmailDelivery {
	if o == nil || o.EmailDelivery == nil {
		var ret EmailDelivery
		return ret
	}
	return *o.EmailDelivery
}

// GetEmailDeliveryOk returns a tuple with the EmailDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEmailDeliveryOk() (*EmailDelivery, bool) {
	if o == nil || o.EmailDelivery == nil {
		return nil, false
	}
	return o.EmailDelivery, true
}

// HasEmailDelivery returns a boolean if a field has been set.
func (o *Event) HasEmailDelivery() bool {
	if o != nil && o.EmailDelivery != nil {
		return true
	}

	return false
}

// SetEmailDelivery gets a reference to the given EmailDelivery and assigns it to the EmailDelivery field.
func (o *Event) SetEmailDelivery(v EmailDelivery) {
	o.EmailDelivery = &v
}

// GetSms returns the Sms field value if set, zero value otherwise.
func (o *Event) GetSms() Sms {
	if o == nil || o.Sms == nil {
		var ret Sms
		return ret
	}
	return *o.Sms
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSmsOk() (*Sms, bool) {
	if o == nil || o.Sms == nil {
		return nil, false
	}
	return o.Sms, true
}

// HasSms returns a boolean if a field has been set.
func (o *Event) HasSms() bool {
	if o != nil && o.Sms != nil {
		return true
	}

	return false
}

// SetSms gets a reference to the given Sms and assigns it to the Sms field.
func (o *Event) SetSms(v Sms) {
	o.Sms = &v
}

// GetSmsInbound returns the SmsInbound field value if set, zero value otherwise.
func (o *Event) GetSmsInbound() SmsInbound {
	if o == nil || o.SmsInbound == nil {
		var ret SmsInbound
		return ret
	}
	return *o.SmsInbound
}

// GetSmsInboundOk returns a tuple with the SmsInbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSmsInboundOk() (*SmsInbound, bool) {
	if o == nil || o.SmsInbound == nil {
		return nil, false
	}
	return o.SmsInbound, true
}

// HasSmsInbound returns a boolean if a field has been set.
func (o *Event) HasSmsInbound() bool {
	if o != nil && o.SmsInbound != nil {
		return true
	}

	return false
}

// SetSmsInbound gets a reference to the given SmsInbound and assigns it to the SmsInbound field.
func (o *Event) SetSmsInbound(v SmsInbound) {
	o.SmsInbound = &v
}

// GetVoice returns the Voice field value if set, zero value otherwise.
func (o *Event) GetVoice() Voice {
	if o == nil || o.Voice == nil {
		var ret Voice
		return ret
	}
	return *o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetVoiceOk() (*Voice, bool) {
	if o == nil || o.Voice == nil {
		return nil, false
	}
	return o.Voice, true
}

// HasVoice returns a boolean if a field has been set.
func (o *Event) HasVoice() bool {
	if o != nil && o.Voice != nil {
		return true
	}

	return false
}

// SetVoice gets a reference to the given Voice and assigns it to the Voice field.
func (o *Event) SetVoice(v Voice) {
	o.Voice = &v
}

// GetWhatsappBusinessAccount returns the WhatsappBusinessAccount field value if set, zero value otherwise.
func (o *Event) GetWhatsappBusinessAccount() WhatsappBusinessAccount {
	if o == nil || o.WhatsappBusinessAccount == nil {
		var ret WhatsappBusinessAccount
		return ret
	}
	return *o.WhatsappBusinessAccount
}

// GetWhatsappBusinessAccountOk returns a tuple with the WhatsappBusinessAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetWhatsappBusinessAccountOk() (*WhatsappBusinessAccount, bool) {
	if o == nil || o.WhatsappBusinessAccount == nil {
		return nil, false
	}
	return o.WhatsappBusinessAccount, true
}

// HasWhatsappBusinessAccount returns a boolean if a field has been set.
func (o *Event) HasWhatsappBusinessAccount() bool {
	if o != nil && o.WhatsappBusinessAccount != nil {
		return true
	}

	return false
}

// SetWhatsappBusinessAccount gets a reference to the given WhatsappBusinessAccount and assigns it to the WhatsappBusinessAccount field.
func (o *Event) SetWhatsappBusinessAccount(v WhatsappBusinessAccount) {
	o.WhatsappBusinessAccount = &v
}

// GetWhatsappInboundMessage returns the WhatsappInboundMessage field value if set, zero value otherwise.
func (o *Event) GetWhatsappInboundMessage() WhatsappInboundMessage {
	if o == nil || o.WhatsappInboundMessage == nil {
		var ret WhatsappInboundMessage
		return ret
	}
	return *o.WhatsappInboundMessage
}

// GetWhatsappInboundMessageOk returns a tuple with the WhatsappInboundMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetWhatsappInboundMessageOk() (*WhatsappInboundMessage, bool) {
	if o == nil || o.WhatsappInboundMessage == nil {
		return nil, false
	}
	return o.WhatsappInboundMessage, true
}

// HasWhatsappInboundMessage returns a boolean if a field has been set.
func (o *Event) HasWhatsappInboundMessage() bool {
	if o != nil && o.WhatsappInboundMessage != nil {
		return true
	}

	return false
}

// SetWhatsappInboundMessage gets a reference to the given WhatsappInboundMessage and assigns it to the WhatsappInboundMessage field.
func (o *Event) SetWhatsappInboundMessage(v WhatsappInboundMessage) {
	o.WhatsappInboundMessage = &v
}

// GetWhatsappMessage returns the WhatsappMessage field value if set, zero value otherwise.
func (o *Event) GetWhatsappMessage() WhatsappMessage {
	if o == nil || o.WhatsappMessage == nil {
		var ret WhatsappMessage
		return ret
	}
	return *o.WhatsappMessage
}

// GetWhatsappMessageOk returns a tuple with the WhatsappMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetWhatsappMessageOk() (*WhatsappMessage, bool) {
	if o == nil || o.WhatsappMessage == nil {
		return nil, false
	}
	return o.WhatsappMessage, true
}

// HasWhatsappMessage returns a boolean if a field has been set.
func (o *Event) HasWhatsappMessage() bool {
	if o != nil && o.WhatsappMessage != nil {
		return true
	}

	return false
}

// SetWhatsappMessage gets a reference to the given WhatsappMessage and assigns it to the WhatsappMessage field.
func (o *Event) SetWhatsappMessage(v WhatsappMessage) {
	o.WhatsappMessage = &v
}

// GetWhatsappPhoneNumber returns the WhatsappPhoneNumber field value if set, zero value otherwise.
func (o *Event) GetWhatsappPhoneNumber() WhatsappPhoneNumber {
	if o == nil || o.WhatsappPhoneNumber == nil {
		var ret WhatsappPhoneNumber
		return ret
	}
	return *o.WhatsappPhoneNumber
}

// GetWhatsappPhoneNumberOk returns a tuple with the WhatsappPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetWhatsappPhoneNumberOk() (*WhatsappPhoneNumber, bool) {
	if o == nil || o.WhatsappPhoneNumber == nil {
		return nil, false
	}
	return o.WhatsappPhoneNumber, true
}

// HasWhatsappPhoneNumber returns a boolean if a field has been set.
func (o *Event) HasWhatsappPhoneNumber() bool {
	if o != nil && o.WhatsappPhoneNumber != nil {
		return true
	}

	return false
}

// SetWhatsappPhoneNumber gets a reference to the given WhatsappPhoneNumber and assigns it to the WhatsappPhoneNumber field.
func (o *Event) SetWhatsappPhoneNumber(v WhatsappPhoneNumber) {
	o.WhatsappPhoneNumber = &v
}

// GetWhatsappTemplate returns the WhatsappTemplate field value if set, zero value otherwise.
func (o *Event) GetWhatsappTemplate() WhatsappTemplate {
	if o == nil || o.WhatsappTemplate == nil {
		var ret WhatsappTemplate
		return ret
	}
	return *o.WhatsappTemplate
}

// GetWhatsappTemplateOk returns a tuple with the WhatsappTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetWhatsappTemplateOk() (*WhatsappTemplate, bool) {
	if o == nil || o.WhatsappTemplate == nil {
		return nil, false
	}
	return o.WhatsappTemplate, true
}

// HasWhatsappTemplate returns a boolean if a field has been set.
func (o *Event) HasWhatsappTemplate() bool {
	if o != nil && o.WhatsappTemplate != nil {
		return true
	}

	return false
}

// SetWhatsappTemplate gets a reference to the given WhatsappTemplate and assigns it to the WhatsappTemplate field.
func (o *Event) SetWhatsappTemplate(v WhatsappTemplate) {
	o.WhatsappTemplate = &v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if true {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.EmailDelivery != nil {
		toSerialize["emailDelivery"] = o.EmailDelivery
	}
	if o.Sms != nil {
		toSerialize["sms"] = o.Sms
	}
	if o.SmsInbound != nil {
		toSerialize["smsInbound"] = o.SmsInbound
	}
	if o.Voice != nil {
		toSerialize["voice"] = o.Voice
	}
	if o.WhatsappBusinessAccount != nil {
		toSerialize["whatsappBusinessAccount"] = o.WhatsappBusinessAccount
	}
	if o.WhatsappInboundMessage != nil {
		toSerialize["whatsappInboundMessage"] = o.WhatsappInboundMessage
	}
	if o.WhatsappMessage != nil {
		toSerialize["whatsappMessage"] = o.WhatsappMessage
	}
	if o.WhatsappPhoneNumber != nil {
		toSerialize["whatsappPhoneNumber"] = o.WhatsappPhoneNumber
	}
	if o.WhatsappTemplate != nil {
		toSerialize["whatsappTemplate"] = o.WhatsappTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
