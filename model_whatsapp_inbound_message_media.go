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

// WhatsappInboundMessageMedia When a message with media (`image` | `document` | `audio` | `video` | `sticker`) is received, the WhatsApp Business API client will download the media. Once the media is downloaded, a notification is sent to your Webhook. This message contains information that identifies the media object and enables you to find and download the object.
type WhatsappInboundMessageMedia struct {
	// ID of the media. Can be used to delete the media if stored locally on the client.
	Id *string `json:"id,omitempty"`
	// The url to download the media file. Note that This link can be directly accessed in a few minutes for the convenience of the consumer, but you should always include an `X-API-Key` header to download this file within a month.
	Link *string `json:"link,omitempty"`
	// The provided caption for the media. Only present if specified.
	Caption *string `json:"caption,omitempty"`
	// Filename on the sender's device. This will only be present in `document` media messages.
	Filename *string `json:"filename,omitempty"`
	// Metadata pertaining to `sticker` media.
	Metadata map[string]map[string]interface{} `json:"metadata,omitempty"`
	// Mime type of the media.
	MimeType *string `json:"mime_type,omitempty"`
	// Checksum.
	Sha256 *string `json:"sha256,omitempty"`
}

// NewWhatsappInboundMessageMedia instantiates a new WhatsappInboundMessageMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappInboundMessageMedia() *WhatsappInboundMessageMedia {
	this := WhatsappInboundMessageMedia{}
	return &this
}

// NewWhatsappInboundMessageMediaWithDefaults instantiates a new WhatsappInboundMessageMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappInboundMessageMediaWithDefaults() *WhatsappInboundMessageMedia {
	this := WhatsappInboundMessageMedia{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WhatsappInboundMessageMedia) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageMedia) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WhatsappInboundMessageMedia) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WhatsappInboundMessageMedia) SetId(v string) {
	o.Id = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *WhatsappInboundMessageMedia) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageMedia) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *WhatsappInboundMessageMedia) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *WhatsappInboundMessageMedia) SetLink(v string) {
	o.Link = &v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *WhatsappInboundMessageMedia) GetCaption() string {
	if o == nil || o.Caption == nil {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageMedia) GetCaptionOk() (*string, bool) {
	if o == nil || o.Caption == nil {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *WhatsappInboundMessageMedia) HasCaption() bool {
	if o != nil && o.Caption != nil {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *WhatsappInboundMessageMedia) SetCaption(v string) {
	o.Caption = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *WhatsappInboundMessageMedia) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageMedia) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *WhatsappInboundMessageMedia) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *WhatsappInboundMessageMedia) SetFilename(v string) {
	o.Filename = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *WhatsappInboundMessageMedia) GetMetadata() map[string]map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageMedia) GetMetadataOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WhatsappInboundMessageMedia) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]map[string]interface{} and assigns it to the Metadata field.
func (o *WhatsappInboundMessageMedia) SetMetadata(v map[string]map[string]interface{}) {
	o.Metadata = v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *WhatsappInboundMessageMedia) GetMimeType() string {
	if o == nil || o.MimeType == nil {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageMedia) GetMimeTypeOk() (*string, bool) {
	if o == nil || o.MimeType == nil {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *WhatsappInboundMessageMedia) HasMimeType() bool {
	if o != nil && o.MimeType != nil {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *WhatsappInboundMessageMedia) SetMimeType(v string) {
	o.MimeType = &v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *WhatsappInboundMessageMedia) GetSha256() string {
	if o == nil || o.Sha256 == nil {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageMedia) GetSha256Ok() (*string, bool) {
	if o == nil || o.Sha256 == nil {
		return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *WhatsappInboundMessageMedia) HasSha256() bool {
	if o != nil && o.Sha256 != nil {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *WhatsappInboundMessageMedia) SetSha256(v string) {
	o.Sha256 = &v
}

func (o WhatsappInboundMessageMedia) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.Caption != nil {
		toSerialize["caption"] = o.Caption
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.MimeType != nil {
		toSerialize["mime_type"] = o.MimeType
	}
	if o.Sha256 != nil {
		toSerialize["sha256"] = o.Sha256
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappInboundMessageMedia struct {
	value *WhatsappInboundMessageMedia
	isSet bool
}

func (v NullableWhatsappInboundMessageMedia) Get() *WhatsappInboundMessageMedia {
	return v.value
}

func (v *NullableWhatsappInboundMessageMedia) Set(val *WhatsappInboundMessageMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappInboundMessageMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappInboundMessageMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappInboundMessageMedia(val *WhatsappInboundMessageMedia) *NullableWhatsappInboundMessageMedia {
	return &NullableWhatsappInboundMessageMedia{value: val, isSet: true}
}

func (v NullableWhatsappInboundMessageMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappInboundMessageMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


