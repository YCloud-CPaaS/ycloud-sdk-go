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

// WhatsappMessage WhatsApp outbound message object.
type WhatsappMessage struct {
	// Unique ID for the object.
	Id string `json:"id"`
	// The original message ID on WhatsApp's platform.
	Wamid *string `json:"wamid,omitempty"`
	// WhatsApp Business Account ID.
	WabaId string `json:"wabaId"`
	// The sender's phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	From string `json:"from"`
	// The recipient's phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	To string `json:"to"`
	Conversation *WhatsappConversation `json:"conversation,omitempty"`
	Type WhatsappMessageType `json:"type"`
	Template *WhatsappMessageTemplate `json:"template,omitempty"`
	Text *WhatsappMessageText `json:"text,omitempty"`
	Image *WhatsappMessageMedia `json:"image,omitempty"`
	Video *WhatsappMessageMedia `json:"video,omitempty"`
	Audio *WhatsappMessageMedia `json:"audio,omitempty"`
	Document *WhatsappMessageMedia `json:"document,omitempty"`
	Sticker *WhatsappMessageMedia `json:"sticker,omitempty"`
	Location *WhatsappMessageLocation `json:"location,omitempty"`
	Interactive *WhatsappMessageInteractive `json:"interactive,omitempty"`
	Contacts []WhatsappMessageContact `json:"contacts,omitempty"`
	Reaction *WhatsappMessageReaction `json:"reaction,omitempty"`
	Context *WhatsappMessageContext `json:"context,omitempty"`
	// A unique string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems.
	ExternalId *string `json:"externalId,omitempty"`
	Status *WhatsappMessageStatus `json:"status,omitempty"`
	// Error code when the message status is `failed`.
	ErrorCode *string `json:"errorCode,omitempty"`
	// Error message when the message status is `failed`.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The time at which this message is created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., `2022-06-01T12:00:00.000Z`.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// The time at which this message is updated, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., `2022-06-01T12:00:00.000Z`.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// Total price of this message.
	TotalPrice *float64 `json:"totalPrice,omitempty"`
	// Price currency. [ISO 4217 currency code](https://en.wikipedia.org/wiki/ISO_4217).
	Currency *string `json:"currency,omitempty"`
	WhatsappApiError *WhatsappApiError `json:"whatsappApiError,omitempty"`
	// This can be either empty or one of `whatsapp`, or `verify`. Defaults to `whatsapp`. - `whatsapp`: Indicates that the message is sent via [WhatsApp](https://www.ycloud.com/whatsapp) product. - `verify`: Indicates that the message is sent via [Verify](https://www.ycloud.com/verify) product.
	BizType *string `json:"bizType,omitempty"`
	// The verification ID. Included only when `bizType` is `verify`.
	VerificationId *string `json:"verificationId,omitempty"`
}

// NewWhatsappMessage instantiates a new WhatsappMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessage(id string, wabaId string, from string, to string, type_ WhatsappMessageType) *WhatsappMessage {
	this := WhatsappMessage{}
	this.Id = id
	this.WabaId = wabaId
	this.From = from
	this.To = to
	this.Type = type_
	return &this
}

// NewWhatsappMessageWithDefaults instantiates a new WhatsappMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageWithDefaults() *WhatsappMessage {
	this := WhatsappMessage{}
	return &this
}

// GetId returns the Id field value
func (o *WhatsappMessage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WhatsappMessage) SetId(v string) {
	o.Id = v
}

// GetWamid returns the Wamid field value if set, zero value otherwise.
func (o *WhatsappMessage) GetWamid() string {
	if o == nil || o.Wamid == nil {
		var ret string
		return ret
	}
	return *o.Wamid
}

// GetWamidOk returns a tuple with the Wamid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetWamidOk() (*string, bool) {
	if o == nil || o.Wamid == nil {
		return nil, false
	}
	return o.Wamid, true
}

// HasWamid returns a boolean if a field has been set.
func (o *WhatsappMessage) HasWamid() bool {
	if o != nil && o.Wamid != nil {
		return true
	}

	return false
}

// SetWamid gets a reference to the given string and assigns it to the Wamid field.
func (o *WhatsappMessage) SetWamid(v string) {
	o.Wamid = &v
}

// GetWabaId returns the WabaId field value
func (o *WhatsappMessage) GetWabaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WabaId
}

// GetWabaIdOk returns a tuple with the WabaId field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetWabaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WabaId, true
}

// SetWabaId sets field value
func (o *WhatsappMessage) SetWabaId(v string) {
	o.WabaId = v
}

// GetFrom returns the From field value
func (o *WhatsappMessage) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *WhatsappMessage) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *WhatsappMessage) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *WhatsappMessage) SetTo(v string) {
	o.To = v
}

// GetConversation returns the Conversation field value if set, zero value otherwise.
func (o *WhatsappMessage) GetConversation() WhatsappConversation {
	if o == nil || o.Conversation == nil {
		var ret WhatsappConversation
		return ret
	}
	return *o.Conversation
}

// GetConversationOk returns a tuple with the Conversation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetConversationOk() (*WhatsappConversation, bool) {
	if o == nil || o.Conversation == nil {
		return nil, false
	}
	return o.Conversation, true
}

// HasConversation returns a boolean if a field has been set.
func (o *WhatsappMessage) HasConversation() bool {
	if o != nil && o.Conversation != nil {
		return true
	}

	return false
}

// SetConversation gets a reference to the given WhatsappConversation and assigns it to the Conversation field.
func (o *WhatsappMessage) SetConversation(v WhatsappConversation) {
	o.Conversation = &v
}

// GetType returns the Type field value
func (o *WhatsappMessage) GetType() WhatsappMessageType {
	if o == nil {
		var ret WhatsappMessageType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetTypeOk() (*WhatsappMessageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WhatsappMessage) SetType(v WhatsappMessageType) {
	o.Type = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *WhatsappMessage) GetTemplate() WhatsappMessageTemplate {
	if o == nil || o.Template == nil {
		var ret WhatsappMessageTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetTemplateOk() (*WhatsappMessageTemplate, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *WhatsappMessage) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given WhatsappMessageTemplate and assigns it to the Template field.
func (o *WhatsappMessage) SetTemplate(v WhatsappMessageTemplate) {
	o.Template = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *WhatsappMessage) GetText() WhatsappMessageText {
	if o == nil || o.Text == nil {
		var ret WhatsappMessageText
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetTextOk() (*WhatsappMessageText, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WhatsappMessage) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given WhatsappMessageText and assigns it to the Text field.
func (o *WhatsappMessage) SetText(v WhatsappMessageText) {
	o.Text = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *WhatsappMessage) GetImage() WhatsappMessageMedia {
	if o == nil || o.Image == nil {
		var ret WhatsappMessageMedia
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetImageOk() (*WhatsappMessageMedia, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *WhatsappMessage) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given WhatsappMessageMedia and assigns it to the Image field.
func (o *WhatsappMessage) SetImage(v WhatsappMessageMedia) {
	o.Image = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *WhatsappMessage) GetVideo() WhatsappMessageMedia {
	if o == nil || o.Video == nil {
		var ret WhatsappMessageMedia
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetVideoOk() (*WhatsappMessageMedia, bool) {
	if o == nil || o.Video == nil {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *WhatsappMessage) HasVideo() bool {
	if o != nil && o.Video != nil {
		return true
	}

	return false
}

// SetVideo gets a reference to the given WhatsappMessageMedia and assigns it to the Video field.
func (o *WhatsappMessage) SetVideo(v WhatsappMessageMedia) {
	o.Video = &v
}

// GetAudio returns the Audio field value if set, zero value otherwise.
func (o *WhatsappMessage) GetAudio() WhatsappMessageMedia {
	if o == nil || o.Audio == nil {
		var ret WhatsappMessageMedia
		return ret
	}
	return *o.Audio
}

// GetAudioOk returns a tuple with the Audio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetAudioOk() (*WhatsappMessageMedia, bool) {
	if o == nil || o.Audio == nil {
		return nil, false
	}
	return o.Audio, true
}

// HasAudio returns a boolean if a field has been set.
func (o *WhatsappMessage) HasAudio() bool {
	if o != nil && o.Audio != nil {
		return true
	}

	return false
}

// SetAudio gets a reference to the given WhatsappMessageMedia and assigns it to the Audio field.
func (o *WhatsappMessage) SetAudio(v WhatsappMessageMedia) {
	o.Audio = &v
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *WhatsappMessage) GetDocument() WhatsappMessageMedia {
	if o == nil || o.Document == nil {
		var ret WhatsappMessageMedia
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetDocumentOk() (*WhatsappMessageMedia, bool) {
	if o == nil || o.Document == nil {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *WhatsappMessage) HasDocument() bool {
	if o != nil && o.Document != nil {
		return true
	}

	return false
}

// SetDocument gets a reference to the given WhatsappMessageMedia and assigns it to the Document field.
func (o *WhatsappMessage) SetDocument(v WhatsappMessageMedia) {
	o.Document = &v
}

// GetSticker returns the Sticker field value if set, zero value otherwise.
func (o *WhatsappMessage) GetSticker() WhatsappMessageMedia {
	if o == nil || o.Sticker == nil {
		var ret WhatsappMessageMedia
		return ret
	}
	return *o.Sticker
}

// GetStickerOk returns a tuple with the Sticker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetStickerOk() (*WhatsappMessageMedia, bool) {
	if o == nil || o.Sticker == nil {
		return nil, false
	}
	return o.Sticker, true
}

// HasSticker returns a boolean if a field has been set.
func (o *WhatsappMessage) HasSticker() bool {
	if o != nil && o.Sticker != nil {
		return true
	}

	return false
}

// SetSticker gets a reference to the given WhatsappMessageMedia and assigns it to the Sticker field.
func (o *WhatsappMessage) SetSticker(v WhatsappMessageMedia) {
	o.Sticker = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *WhatsappMessage) GetLocation() WhatsappMessageLocation {
	if o == nil || o.Location == nil {
		var ret WhatsappMessageLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetLocationOk() (*WhatsappMessageLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *WhatsappMessage) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given WhatsappMessageLocation and assigns it to the Location field.
func (o *WhatsappMessage) SetLocation(v WhatsappMessageLocation) {
	o.Location = &v
}

// GetInteractive returns the Interactive field value if set, zero value otherwise.
func (o *WhatsappMessage) GetInteractive() WhatsappMessageInteractive {
	if o == nil || o.Interactive == nil {
		var ret WhatsappMessageInteractive
		return ret
	}
	return *o.Interactive
}

// GetInteractiveOk returns a tuple with the Interactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetInteractiveOk() (*WhatsappMessageInteractive, bool) {
	if o == nil || o.Interactive == nil {
		return nil, false
	}
	return o.Interactive, true
}

// HasInteractive returns a boolean if a field has been set.
func (o *WhatsappMessage) HasInteractive() bool {
	if o != nil && o.Interactive != nil {
		return true
	}

	return false
}

// SetInteractive gets a reference to the given WhatsappMessageInteractive and assigns it to the Interactive field.
func (o *WhatsappMessage) SetInteractive(v WhatsappMessageInteractive) {
	o.Interactive = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *WhatsappMessage) GetContacts() []WhatsappMessageContact {
	if o == nil || o.Contacts == nil {
		var ret []WhatsappMessageContact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetContactsOk() ([]WhatsappMessageContact, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *WhatsappMessage) HasContacts() bool {
	if o != nil && o.Contacts != nil {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []WhatsappMessageContact and assigns it to the Contacts field.
func (o *WhatsappMessage) SetContacts(v []WhatsappMessageContact) {
	o.Contacts = v
}

// GetReaction returns the Reaction field value if set, zero value otherwise.
func (o *WhatsappMessage) GetReaction() WhatsappMessageReaction {
	if o == nil || o.Reaction == nil {
		var ret WhatsappMessageReaction
		return ret
	}
	return *o.Reaction
}

// GetReactionOk returns a tuple with the Reaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetReactionOk() (*WhatsappMessageReaction, bool) {
	if o == nil || o.Reaction == nil {
		return nil, false
	}
	return o.Reaction, true
}

// HasReaction returns a boolean if a field has been set.
func (o *WhatsappMessage) HasReaction() bool {
	if o != nil && o.Reaction != nil {
		return true
	}

	return false
}

// SetReaction gets a reference to the given WhatsappMessageReaction and assigns it to the Reaction field.
func (o *WhatsappMessage) SetReaction(v WhatsappMessageReaction) {
	o.Reaction = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *WhatsappMessage) GetContext() WhatsappMessageContext {
	if o == nil || o.Context == nil {
		var ret WhatsappMessageContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetContextOk() (*WhatsappMessageContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *WhatsappMessage) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given WhatsappMessageContext and assigns it to the Context field.
func (o *WhatsappMessage) SetContext(v WhatsappMessageContext) {
	o.Context = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *WhatsappMessage) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *WhatsappMessage) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *WhatsappMessage) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WhatsappMessage) GetStatus() WhatsappMessageStatus {
	if o == nil || o.Status == nil {
		var ret WhatsappMessageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetStatusOk() (*WhatsappMessageStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WhatsappMessage) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WhatsappMessageStatus and assigns it to the Status field.
func (o *WhatsappMessage) SetStatus(v WhatsappMessageStatus) {
	o.Status = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *WhatsappMessage) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *WhatsappMessage) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *WhatsappMessage) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *WhatsappMessage) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *WhatsappMessage) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *WhatsappMessage) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *WhatsappMessage) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *WhatsappMessage) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *WhatsappMessage) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *WhatsappMessage) GetUpdateTime() time.Time {
	if o == nil || o.UpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *WhatsappMessage) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *WhatsappMessage) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetTotalPrice returns the TotalPrice field value if set, zero value otherwise.
func (o *WhatsappMessage) GetTotalPrice() float64 {
	if o == nil || o.TotalPrice == nil {
		var ret float64
		return ret
	}
	return *o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetTotalPriceOk() (*float64, bool) {
	if o == nil || o.TotalPrice == nil {
		return nil, false
	}
	return o.TotalPrice, true
}

// HasTotalPrice returns a boolean if a field has been set.
func (o *WhatsappMessage) HasTotalPrice() bool {
	if o != nil && o.TotalPrice != nil {
		return true
	}

	return false
}

// SetTotalPrice gets a reference to the given float64 and assigns it to the TotalPrice field.
func (o *WhatsappMessage) SetTotalPrice(v float64) {
	o.TotalPrice = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *WhatsappMessage) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *WhatsappMessage) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *WhatsappMessage) SetCurrency(v string) {
	o.Currency = &v
}

// GetWhatsappApiError returns the WhatsappApiError field value if set, zero value otherwise.
func (o *WhatsappMessage) GetWhatsappApiError() WhatsappApiError {
	if o == nil || o.WhatsappApiError == nil {
		var ret WhatsappApiError
		return ret
	}
	return *o.WhatsappApiError
}

// GetWhatsappApiErrorOk returns a tuple with the WhatsappApiError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetWhatsappApiErrorOk() (*WhatsappApiError, bool) {
	if o == nil || o.WhatsappApiError == nil {
		return nil, false
	}
	return o.WhatsappApiError, true
}

// HasWhatsappApiError returns a boolean if a field has been set.
func (o *WhatsappMessage) HasWhatsappApiError() bool {
	if o != nil && o.WhatsappApiError != nil {
		return true
	}

	return false
}

// SetWhatsappApiError gets a reference to the given WhatsappApiError and assigns it to the WhatsappApiError field.
func (o *WhatsappMessage) SetWhatsappApiError(v WhatsappApiError) {
	o.WhatsappApiError = &v
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *WhatsappMessage) GetBizType() string {
	if o == nil || o.BizType == nil {
		var ret string
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetBizTypeOk() (*string, bool) {
	if o == nil || o.BizType == nil {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *WhatsappMessage) HasBizType() bool {
	if o != nil && o.BizType != nil {
		return true
	}

	return false
}

// SetBizType gets a reference to the given string and assigns it to the BizType field.
func (o *WhatsappMessage) SetBizType(v string) {
	o.BizType = &v
}

// GetVerificationId returns the VerificationId field value if set, zero value otherwise.
func (o *WhatsappMessage) GetVerificationId() string {
	if o == nil || o.VerificationId == nil {
		var ret string
		return ret
	}
	return *o.VerificationId
}

// GetVerificationIdOk returns a tuple with the VerificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessage) GetVerificationIdOk() (*string, bool) {
	if o == nil || o.VerificationId == nil {
		return nil, false
	}
	return o.VerificationId, true
}

// HasVerificationId returns a boolean if a field has been set.
func (o *WhatsappMessage) HasVerificationId() bool {
	if o != nil && o.VerificationId != nil {
		return true
	}

	return false
}

// SetVerificationId gets a reference to the given string and assigns it to the VerificationId field.
func (o *WhatsappMessage) SetVerificationId(v string) {
	o.VerificationId = &v
}

func (o WhatsappMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Wamid != nil {
		toSerialize["wamid"] = o.Wamid
	}
	if true {
		toSerialize["wabaId"] = o.WabaId
	}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	if o.Conversation != nil {
		toSerialize["conversation"] = o.Conversation
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
	if o.Sticker != nil {
		toSerialize["sticker"] = o.Sticker
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
	if o.Reaction != nil {
		toSerialize["reaction"] = o.Reaction
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.UpdateTime != nil {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if o.TotalPrice != nil {
		toSerialize["totalPrice"] = o.TotalPrice
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.WhatsappApiError != nil {
		toSerialize["whatsappApiError"] = o.WhatsappApiError
	}
	if o.BizType != nil {
		toSerialize["bizType"] = o.BizType
	}
	if o.VerificationId != nil {
		toSerialize["verificationId"] = o.VerificationId
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessage struct {
	value *WhatsappMessage
	isSet bool
}

func (v NullableWhatsappMessage) Get() *WhatsappMessage {
	return v.value
}

func (v *NullableWhatsappMessage) Set(val *WhatsappMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessage(val *WhatsappMessage) *NullableWhatsappMessage {
	return &NullableWhatsappMessage{value: val, isSet: true}
}

func (v NullableWhatsappMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


