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

// WhatsappMessageSendRequest struct for WhatsappMessageSendRequest
type WhatsappMessageSendRequest struct {
	// The sender's phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	From string `json:"from"`
	// The recipient's phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	To string `json:"to"`
	Type WhatsappMessageType `json:"type"`
	Template *WhatsappMessageTemplate `json:"template,omitempty"`
	Text *WhatsappMessageText `json:"text,omitempty"`
	Image *WhatsappMessageMedia `json:"image,omitempty"`
	Video *WhatsappMessageMedia `json:"video,omitempty"`
	Audio *WhatsappMessageMedia `json:"audio,omitempty"`
	Document *WhatsappMessageMedia `json:"document,omitempty"`
	Location *WhatsappMessageLocation `json:"location,omitempty"`
	Interactive *WhatsappMessageInteractive `json:"interactive,omitempty"`
	// Required when `type` is `contacts`.
	Contacts []WhatsappMessageContact `json:"contacts,omitempty"`
	// A unique string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems.
	ExternalId *string `json:"externalId,omitempty"`
}

// NewWhatsappMessageSendRequest instantiates a new WhatsappMessageSendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageSendRequest(from string, to string, type_ WhatsappMessageType) *WhatsappMessageSendRequest {
	this := WhatsappMessageSendRequest{}
	this.From = from
	this.To = to
	this.Type = type_
	return &this
}

// NewWhatsappMessageSendRequestWithDefaults instantiates a new WhatsappMessageSendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageSendRequestWithDefaults() *WhatsappMessageSendRequest {
	this := WhatsappMessageSendRequest{}
	return &this
}

// GetFrom returns the From field value
func (o *WhatsappMessageSendRequest) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *WhatsappMessageSendRequest) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *WhatsappMessageSendRequest) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *WhatsappMessageSendRequest) SetTo(v string) {
	o.To = v
}

// GetType returns the Type field value
func (o *WhatsappMessageSendRequest) GetType() WhatsappMessageType {
	if o == nil {
		var ret WhatsappMessageType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetTypeOk() (*WhatsappMessageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WhatsappMessageSendRequest) SetType(v WhatsappMessageType) {
	o.Type = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *WhatsappMessageSendRequest) GetTemplate() WhatsappMessageTemplate {
	if o == nil || o.Template == nil {
		var ret WhatsappMessageTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetTemplateOk() (*WhatsappMessageTemplate, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *WhatsappMessageSendRequest) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given WhatsappMessageTemplate and assigns it to the Template field.
func (o *WhatsappMessageSendRequest) SetTemplate(v WhatsappMessageTemplate) {
	o.Template = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *WhatsappMessageSendRequest) GetText() WhatsappMessageText {
	if o == nil || o.Text == nil {
		var ret WhatsappMessageText
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetTextOk() (*WhatsappMessageText, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WhatsappMessageSendRequest) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given WhatsappMessageText and assigns it to the Text field.
func (o *WhatsappMessageSendRequest) SetText(v WhatsappMessageText) {
	o.Text = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *WhatsappMessageSendRequest) GetImage() WhatsappMessageMedia {
	if o == nil || o.Image == nil {
		var ret WhatsappMessageMedia
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetImageOk() (*WhatsappMessageMedia, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *WhatsappMessageSendRequest) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given WhatsappMessageMedia and assigns it to the Image field.
func (o *WhatsappMessageSendRequest) SetImage(v WhatsappMessageMedia) {
	o.Image = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *WhatsappMessageSendRequest) GetVideo() WhatsappMessageMedia {
	if o == nil || o.Video == nil {
		var ret WhatsappMessageMedia
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetVideoOk() (*WhatsappMessageMedia, bool) {
	if o == nil || o.Video == nil {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *WhatsappMessageSendRequest) HasVideo() bool {
	if o != nil && o.Video != nil {
		return true
	}

	return false
}

// SetVideo gets a reference to the given WhatsappMessageMedia and assigns it to the Video field.
func (o *WhatsappMessageSendRequest) SetVideo(v WhatsappMessageMedia) {
	o.Video = &v
}

// GetAudio returns the Audio field value if set, zero value otherwise.
func (o *WhatsappMessageSendRequest) GetAudio() WhatsappMessageMedia {
	if o == nil || o.Audio == nil {
		var ret WhatsappMessageMedia
		return ret
	}
	return *o.Audio
}

// GetAudioOk returns a tuple with the Audio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetAudioOk() (*WhatsappMessageMedia, bool) {
	if o == nil || o.Audio == nil {
		return nil, false
	}
	return o.Audio, true
}

// HasAudio returns a boolean if a field has been set.
func (o *WhatsappMessageSendRequest) HasAudio() bool {
	if o != nil && o.Audio != nil {
		return true
	}

	return false
}

// SetAudio gets a reference to the given WhatsappMessageMedia and assigns it to the Audio field.
func (o *WhatsappMessageSendRequest) SetAudio(v WhatsappMessageMedia) {
	o.Audio = &v
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *WhatsappMessageSendRequest) GetDocument() WhatsappMessageMedia {
	if o == nil || o.Document == nil {
		var ret WhatsappMessageMedia
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetDocumentOk() (*WhatsappMessageMedia, bool) {
	if o == nil || o.Document == nil {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *WhatsappMessageSendRequest) HasDocument() bool {
	if o != nil && o.Document != nil {
		return true
	}

	return false
}

// SetDocument gets a reference to the given WhatsappMessageMedia and assigns it to the Document field.
func (o *WhatsappMessageSendRequest) SetDocument(v WhatsappMessageMedia) {
	o.Document = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *WhatsappMessageSendRequest) GetLocation() WhatsappMessageLocation {
	if o == nil || o.Location == nil {
		var ret WhatsappMessageLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetLocationOk() (*WhatsappMessageLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *WhatsappMessageSendRequest) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given WhatsappMessageLocation and assigns it to the Location field.
func (o *WhatsappMessageSendRequest) SetLocation(v WhatsappMessageLocation) {
	o.Location = &v
}

// GetInteractive returns the Interactive field value if set, zero value otherwise.
func (o *WhatsappMessageSendRequest) GetInteractive() WhatsappMessageInteractive {
	if o == nil || o.Interactive == nil {
		var ret WhatsappMessageInteractive
		return ret
	}
	return *o.Interactive
}

// GetInteractiveOk returns a tuple with the Interactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetInteractiveOk() (*WhatsappMessageInteractive, bool) {
	if o == nil || o.Interactive == nil {
		return nil, false
	}
	return o.Interactive, true
}

// HasInteractive returns a boolean if a field has been set.
func (o *WhatsappMessageSendRequest) HasInteractive() bool {
	if o != nil && o.Interactive != nil {
		return true
	}

	return false
}

// SetInteractive gets a reference to the given WhatsappMessageInteractive and assigns it to the Interactive field.
func (o *WhatsappMessageSendRequest) SetInteractive(v WhatsappMessageInteractive) {
	o.Interactive = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *WhatsappMessageSendRequest) GetContacts() []WhatsappMessageContact {
	if o == nil || o.Contacts == nil {
		var ret []WhatsappMessageContact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetContactsOk() ([]WhatsappMessageContact, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *WhatsappMessageSendRequest) HasContacts() bool {
	if o != nil && o.Contacts != nil {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []WhatsappMessageContact and assigns it to the Contacts field.
func (o *WhatsappMessageSendRequest) SetContacts(v []WhatsappMessageContact) {
	o.Contacts = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *WhatsappMessageSendRequest) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageSendRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *WhatsappMessageSendRequest) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *WhatsappMessageSendRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o WhatsappMessageSendRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Video != nil {
		toSerialize["video"] = o.Video
	}
	if o.Audio != nil {
		toSerialize["audio"] = o.Audio
	}
	if o.Document != nil {
		toSerialize["document"] = o.Document
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Interactive != nil {
		toSerialize["interactive"] = o.Interactive
	}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageSendRequest struct {
	value *WhatsappMessageSendRequest
	isSet bool
}

func (v NullableWhatsappMessageSendRequest) Get() *WhatsappMessageSendRequest {
	return v.value
}

func (v *NullableWhatsappMessageSendRequest) Set(val *WhatsappMessageSendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageSendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageSendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageSendRequest(val *WhatsappMessageSendRequest) *NullableWhatsappMessageSendRequest {
	return &NullableWhatsappMessageSendRequest{value: val, isSet: true}
}

func (v NullableWhatsappMessageSendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageSendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

