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

// Voice struct for Voice
type Voice struct {
	// Unique ID for the object.
	Id string `json:"id"`
	// The recipient's phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	To string `json:"to"`
	// The verification code to be sent, 4 to 6 digits.
	VerificationCode *string `json:"verificationCode,omitempty"`
	// [ISO 639 Language Code](https://www.iso.org/iso-639-language-codes.html).
	Language *string `json:"language,omitempty"`
	// [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	RegionCode *string `json:"regionCode,omitempty"`
	// Number of message segments. It's always 1 for voice calls.
	TotalSegments *int32 `json:"totalSegments,omitempty"`
	// Total price of this message.
	TotalPrice *float64 `json:"totalPrice,omitempty"`
	// [ISO 4217 currency code](https://en.wikipedia.org/wiki/ISO_4217)
	Currency *string `json:"currency,omitempty"`
	// Delivery status. One of `accepted`, `sent`, `delivered`, `undelivered`.
	Status *string `json:"status,omitempty"`
	// Error code when the message is undeliverable.
	ErrorCode *string `json:"errorCode,omitempty"`
	// The time at which this message was created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., `2022-03-01T12:00:00.000Z`.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// The time at which the delivery report for this message was updated, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., `2022-03-01T12:00:00.000Z`.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// A unique string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems.
	ExternalId *string `json:"externalId,omitempty"`
	// Delivery report URL. You can provide a URL, and we will push the updated status report to your server in time. e.g., https://httpbin.org/anything?tag=api. Note: We recommend configuring Webhook Endpoints instead.
	CallbackUrl *string `json:"callbackUrl,omitempty"`
}

// NewVoice instantiates a new Voice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoice(id string, to string) *Voice {
	this := Voice{}
	this.Id = id
	this.To = to
	return &this
}

// NewVoiceWithDefaults instantiates a new Voice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceWithDefaults() *Voice {
	this := Voice{}
	return &this
}

// GetId returns the Id field value
func (o *Voice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Voice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Voice) SetId(v string) {
	o.Id = v
}

// GetTo returns the To field value
func (o *Voice) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *Voice) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *Voice) SetTo(v string) {
	o.To = v
}

// GetVerificationCode returns the VerificationCode field value if set, zero value otherwise.
func (o *Voice) GetVerificationCode() string {
	if o == nil || o.VerificationCode == nil {
		var ret string
		return ret
	}
	return *o.VerificationCode
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voice) GetVerificationCodeOk() (*string, bool) {
	if o == nil || o.VerificationCode == nil {
		return nil, false
	}
	return o.VerificationCode, true
}

// HasVerificationCode returns a boolean if a field has been set.
func (o *Voice) HasVerificationCode() bool {
	if o != nil && o.VerificationCode != nil {
		return true
	}

	return false
}

// SetVerificationCode gets a reference to the given string and assigns it to the VerificationCode field.
func (o *Voice) SetVerificationCode(v string) {
	o.VerificationCode = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Voice) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voice) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Voice) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Voice) SetLanguage(v string) {
	o.Language = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *Voice) GetRegionCode() string {
	if o == nil || o.RegionCode == nil {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voice) GetRegionCodeOk() (*string, bool) {
	if o == nil || o.RegionCode == nil {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *Voice) HasRegionCode() bool {
	if o != nil && o.RegionCode != nil {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *Voice) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetTotalSegments returns the TotalSegments field value if set, zero value otherwise.
func (o *Voice) GetTotalSegments() int32 {
	if o == nil || o.TotalSegments == nil {
		var ret int32
		return ret
	}
	return *o.TotalSegments
}

// GetTotalSegmentsOk returns a tuple with the TotalSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voice) GetTotalSegmentsOk() (*int32, bool) {
	if o == nil || o.TotalSegments == nil {
		return nil, false
	}
	return o.TotalSegments, true
}

// HasTotalSegments returns a boolean if a field has been set.
func (o *Voice) HasTotalSegments() bool {
	if o != nil && o.TotalSegments != nil {
		return true
	}

	return false
}

// SetTotalSegments gets a reference to the given int32 and assigns it to the TotalSegments field.
func (o *Voice) SetTotalSegments(v int32) {
	o.TotalSegments = &v
}

// GetTotalPrice returns the TotalPrice field value if set, zero value otherwise.
func (o *Voice) GetTotalPrice() float64 {
	if o == nil || o.TotalPrice == nil {
		var ret float64
		return ret
	}
	return *o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voice) GetTotalPriceOk() (*float64, bool) {
	if o == nil || o.TotalPrice == nil {
		return nil, false
	}
	return o.TotalPrice, true
}

// HasTotalPrice returns a boolean if a field has been set.
func (o *Voice) HasTotalPrice() bool {
	if o != nil && o.TotalPrice != nil {
		return true
	}

	return false
}

// SetTotalPrice gets a reference to the given float64 and assigns it to the TotalPrice field.
func (o *Voice) SetTotalPrice(v float64) {
	o.TotalPrice = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Voice) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voice) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Voice) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Voice) SetCurrency(v string) {
	o.Currency = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Voice) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voice) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Voice) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Voice) SetStatus(v string) {
	o.Status = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *Voice) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voice) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *Voice) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *Voice) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *Voice) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voice) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *Voice) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *Voice) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *Voice) GetUpdateTime() time.Time {
	if o == nil || o.UpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voice) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *Voice) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *Voice) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Voice) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voice) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Voice) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Voice) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *Voice) GetCallbackUrl() string {
	if o == nil || o.CallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Voice) GetCallbackUrlOk() (*string, bool) {
	if o == nil || o.CallbackUrl == nil {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *Voice) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl != nil {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *Voice) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

func (o Voice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["to"] = o.To
	}
	if o.VerificationCode != nil {
		toSerialize["verificationCode"] = o.VerificationCode
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.RegionCode != nil {
		toSerialize["regionCode"] = o.RegionCode
	}
	if o.TotalSegments != nil {
		toSerialize["totalSegments"] = o.TotalSegments
	}
	if o.TotalPrice != nil {
		toSerialize["totalPrice"] = o.TotalPrice
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.UpdateTime != nil {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.CallbackUrl != nil {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	return json.Marshal(toSerialize)
}

type NullableVoice struct {
	value *Voice
	isSet bool
}

func (v NullableVoice) Get() *Voice {
	return v.value
}

func (v *NullableVoice) Set(val *Voice) {
	v.value = val
	v.isSet = true
}

func (v NullableVoice) IsSet() bool {
	return v.isSet
}

func (v *NullableVoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoice(val *Voice) *NullableVoice {
	return &NullableVoice{value: val, isSet: true}
}

func (v NullableVoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


