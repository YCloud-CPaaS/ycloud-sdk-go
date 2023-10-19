# WhatsappTemplateComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | **Required.** Template component type. | [optional] 
**Format** | Pointer to **string** | **Required for type &#x60;HEADER&#x60;.** | [optional] 
**Text** | Pointer to **string** | **Required for type &#x60;BODY&#x60;, &#x60;FOOTER&#x60;, and format &#x60;TEXT&#x60;.** | [optional] 
**Buttons** | Pointer to [**[]WhatsappTemplateComponentButton**](WhatsappTemplateComponentButton.md) | **Required for type &#x60;BUTTONS&#x60;.** Buttons are optional interactive components that perform specific actions when tapped. Templates can have a mixture of up to 10 button components total, although there are limits to individual buttons of the same type as well as combination limits. If a template has more than three buttons, two buttons will appear in the delivered message and the remaining buttons will be replaced with a **See all options** button. Tapping the **See all options** button reveals the remaining buttons. | [optional] 
**AddSecurityRecommendation** | Pointer to **bool** | **Optional. Only applicable in the &#x60;BODY&#x60; component of an AUTHENTICATION template.** Set to &#x60;true&#x60; if you want the template to include the string, *For your security, do not share this code.* Set to &#x60;false&#x60; to exclude the string. | [optional] 
**CodeExpirationMinutes** | Pointer to **int32** | **Optional. Only applicable in the &#x60;FOOTER&#x60; component of an AUTHENTICATION template.** Indicates number of minutes the password or code is valid. If omitted, the code expiration warning will not be displayed in the delivered message. Minimum 1, maximum 90. | [optional] 
**Example** | Pointer to [**WhatsappTemplateComponentExample**](WhatsappTemplateComponentExample.md) |  | [optional] 

## Methods

### NewWhatsappTemplateComponent

`func NewWhatsappTemplateComponent() *WhatsappTemplateComponent`

NewWhatsappTemplateComponent instantiates a new WhatsappTemplateComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateComponentWithDefaults

`func NewWhatsappTemplateComponentWithDefaults() *WhatsappTemplateComponent`

NewWhatsappTemplateComponentWithDefaults instantiates a new WhatsappTemplateComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappTemplateComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappTemplateComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappTemplateComponent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappTemplateComponent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFormat

`func (o *WhatsappTemplateComponent) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WhatsappTemplateComponent) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WhatsappTemplateComponent) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *WhatsappTemplateComponent) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetText

`func (o *WhatsappTemplateComponent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappTemplateComponent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappTemplateComponent) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappTemplateComponent) HasText() bool`

HasText returns a boolean if a field has been set.

### GetButtons

`func (o *WhatsappTemplateComponent) GetButtons() []WhatsappTemplateComponentButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *WhatsappTemplateComponent) GetButtonsOk() (*[]WhatsappTemplateComponentButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *WhatsappTemplateComponent) SetButtons(v []WhatsappTemplateComponentButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *WhatsappTemplateComponent) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetAddSecurityRecommendation

`func (o *WhatsappTemplateComponent) GetAddSecurityRecommendation() bool`

GetAddSecurityRecommendation returns the AddSecurityRecommendation field if non-nil, zero value otherwise.

### GetAddSecurityRecommendationOk

`func (o *WhatsappTemplateComponent) GetAddSecurityRecommendationOk() (*bool, bool)`

GetAddSecurityRecommendationOk returns a tuple with the AddSecurityRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSecurityRecommendation

`func (o *WhatsappTemplateComponent) SetAddSecurityRecommendation(v bool)`

SetAddSecurityRecommendation sets AddSecurityRecommendation field to given value.

### HasAddSecurityRecommendation

`func (o *WhatsappTemplateComponent) HasAddSecurityRecommendation() bool`

HasAddSecurityRecommendation returns a boolean if a field has been set.

### GetCodeExpirationMinutes

`func (o *WhatsappTemplateComponent) GetCodeExpirationMinutes() int32`

GetCodeExpirationMinutes returns the CodeExpirationMinutes field if non-nil, zero value otherwise.

### GetCodeExpirationMinutesOk

`func (o *WhatsappTemplateComponent) GetCodeExpirationMinutesOk() (*int32, bool)`

GetCodeExpirationMinutesOk returns a tuple with the CodeExpirationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeExpirationMinutes

`func (o *WhatsappTemplateComponent) SetCodeExpirationMinutes(v int32)`

SetCodeExpirationMinutes sets CodeExpirationMinutes field to given value.

### HasCodeExpirationMinutes

`func (o *WhatsappTemplateComponent) HasCodeExpirationMinutes() bool`

HasCodeExpirationMinutes returns a boolean if a field has been set.

### GetExample

`func (o *WhatsappTemplateComponent) GetExample() WhatsappTemplateComponentExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *WhatsappTemplateComponent) GetExampleOk() (*WhatsappTemplateComponentExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *WhatsappTemplateComponent) SetExample(v WhatsappTemplateComponentExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *WhatsappTemplateComponent) HasExample() bool`

HasExample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


