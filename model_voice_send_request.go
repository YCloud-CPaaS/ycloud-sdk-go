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

// VoiceSendRequest struct for VoiceSendRequest
type VoiceSendRequest struct {
	// The recipient's phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	To string `json:"to"`
	// The verification code to be sent, 4 to 6 digits.
	VerificationCode string `json:"verificationCode"`
	// [ISO 639 Language Code](https://www.iso.org/iso-639-language-codes.html). If not specified, language will be set as `en` by default. Notably, in certain countries or regions, language will be automatically set as the local language due to the regional restrictions. Other applicable languages: `en`: English `zh`: Chinese `id`: Indonesian `vi`: Vietnamese `tr`: Turkish `ru`: Russian `de`: German `fr`: French `it`: Italian `pt`: Portuguese
	Language *string `json:"language,omitempty"`
	// A unique string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems.
	ExternalId *string `json:"externalId,omitempty"`
	// Delivery report URL. You can provide a URL, and we will push the updated status report to your server in time. e.g., https://httpbin.org/anything?tag=api. Note: We recommend configuring Webhook Endpoints instead.
	CallbackUrl *string `json:"callbackUrl,omitempty"`
}

// NewVoiceSendRequest instantiates a new VoiceSendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoiceSendRequest(to string, verificationCode string) *VoiceSendRequest {
	this := VoiceSendRequest{}
	this.To = to
	this.VerificationCode = verificationCode
	return &this
}

// NewVoiceSendRequestWithDefaults instantiates a new VoiceSendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceSendRequestWithDefaults() *VoiceSendRequest {
	this := VoiceSendRequest{}
	return &this
}

// GetTo returns the To field value
func (o *VoiceSendRequest) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *VoiceSendRequest) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *VoiceSendRequest) SetTo(v string) {
	o.To = v
}

// GetVerificationCode returns the VerificationCode field value
func (o *VoiceSendRequest) GetVerificationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationCode
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value
// and a boolean to check if the value has been set.
func (o *VoiceSendRequest) GetVerificationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationCode, true
}

// SetVerificationCode sets field value
func (o *VoiceSendRequest) SetVerificationCode(v string) {
	o.VerificationCode = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *VoiceSendRequest) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceSendRequest) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *VoiceSendRequest) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *VoiceSendRequest) SetLanguage(v string) {
	o.Language = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *VoiceSendRequest) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceSendRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *VoiceSendRequest) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *VoiceSendRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *VoiceSendRequest) GetCallbackUrl() string {
	if o == nil || o.CallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceSendRequest) GetCallbackUrlOk() (*string, bool) {
	if o == nil || o.CallbackUrl == nil {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *VoiceSendRequest) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl != nil {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *VoiceSendRequest) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

func (o VoiceSendRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["to"] = o.To
	}
	if true {
		toSerialize["verificationCode"] = o.VerificationCode
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.CallbackUrl != nil {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	return json.Marshal(toSerialize)
}

type NullableVoiceSendRequest struct {
	value *VoiceSendRequest
	isSet bool
}

func (v NullableVoiceSendRequest) Get() *VoiceSendRequest {
	return v.value
}

func (v *NullableVoiceSendRequest) Set(val *VoiceSendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceSendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceSendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceSendRequest(val *VoiceSendRequest) *NullableVoiceSendRequest {
	return &NullableVoiceSendRequest{value: val, isSet: true}
}

func (v NullableVoiceSendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceSendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

