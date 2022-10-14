# WhatsappMessageInteractiveBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | The body content of the message. Emojis and markdown are supported. Links are supported. | [optional] 

## Methods

### NewWhatsappMessageInteractiveBody

`func NewWhatsappMessageInteractiveBody() *WhatsappMessageInteractiveBody`

NewWhatsappMessageInteractiveBody instantiates a new WhatsappMessageInteractiveBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveBodyWithDefaults

`func NewWhatsappMessageInteractiveBodyWithDefaults() *WhatsappMessageInteractiveBody`

NewWhatsappMessageInteractiveBodyWithDefaults instantiates a new WhatsappMessageInteractiveBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *WhatsappMessageInteractiveBody) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappMessageInteractiveBody) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappMessageInteractiveBody) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappMessageInteractiveBody) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


