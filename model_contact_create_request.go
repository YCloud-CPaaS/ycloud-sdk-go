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

// ContactCreateRequest Contains the properties of the contact to be created.
type ContactCreateRequest struct {
	// Contact's nickname. Maximum length: 250 characters.
	Nickname *string `json:"nickname,omitempty"`
	// Unique Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	PhoneNumber string `json:"phoneNumber"`
	// Two-letter country abbreviation. See [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	CountryCode *string `json:"countryCode,omitempty"`
	// Contact's email address. If present, the email address must be unique.
	Email *string `json:"email,omitempty"`
	// Contact's tags. Max items: 50. Max characters per tag: 50.
	Tags []string `json:"tags,omitempty"`
	// Contact's custom attributes.
	CustomAttributes []ContactCustomAttribute `json:"customAttributes,omitempty"`
}

// NewContactCreateRequest instantiates a new ContactCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactCreateRequest(phoneNumber string) *ContactCreateRequest {
	this := ContactCreateRequest{}
	this.PhoneNumber = phoneNumber
	return &this
}

// NewContactCreateRequestWithDefaults instantiates a new ContactCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactCreateRequestWithDefaults() *ContactCreateRequest {
	this := ContactCreateRequest{}
	return &this
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *ContactCreateRequest) GetNickname() string {
	if o == nil || o.Nickname == nil {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactCreateRequest) GetNicknameOk() (*string, bool) {
	if o == nil || o.Nickname == nil {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *ContactCreateRequest) HasNickname() bool {
	if o != nil && o.Nickname != nil {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *ContactCreateRequest) SetNickname(v string) {
	o.Nickname = &v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *ContactCreateRequest) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *ContactCreateRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *ContactCreateRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *ContactCreateRequest) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactCreateRequest) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *ContactCreateRequest) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *ContactCreateRequest) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ContactCreateRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactCreateRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ContactCreateRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ContactCreateRequest) SetEmail(v string) {
	o.Email = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ContactCreateRequest) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactCreateRequest) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ContactCreateRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ContactCreateRequest) SetTags(v []string) {
	o.Tags = v
}

// GetCustomAttributes returns the CustomAttributes field value if set, zero value otherwise.
func (o *ContactCreateRequest) GetCustomAttributes() []ContactCustomAttribute {
	if o == nil || o.CustomAttributes == nil {
		var ret []ContactCustomAttribute
		return ret
	}
	return o.CustomAttributes
}

// GetCustomAttributesOk returns a tuple with the CustomAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactCreateRequest) GetCustomAttributesOk() ([]ContactCustomAttribute, bool) {
	if o == nil || o.CustomAttributes == nil {
		return nil, false
	}
	return o.CustomAttributes, true
}

// HasCustomAttributes returns a boolean if a field has been set.
func (o *ContactCreateRequest) HasCustomAttributes() bool {
	if o != nil && o.CustomAttributes != nil {
		return true
	}

	return false
}

// SetCustomAttributes gets a reference to the given []ContactCustomAttribute and assigns it to the CustomAttributes field.
func (o *ContactCreateRequest) SetCustomAttributes(v []ContactCustomAttribute) {
	o.CustomAttributes = v
}

func (o ContactCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nickname != nil {
		toSerialize["nickname"] = o.Nickname
	}
	if true {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CustomAttributes != nil {
		toSerialize["customAttributes"] = o.CustomAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableContactCreateRequest struct {
	value *ContactCreateRequest
	isSet bool
}

func (v NullableContactCreateRequest) Get() *ContactCreateRequest {
	return v.value
}

func (v *NullableContactCreateRequest) Set(val *ContactCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContactCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContactCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactCreateRequest(val *ContactCreateRequest) *NullableContactCreateRequest {
	return &NullableContactCreateRequest{value: val, isSet: true}
}

func (v NullableContactCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


