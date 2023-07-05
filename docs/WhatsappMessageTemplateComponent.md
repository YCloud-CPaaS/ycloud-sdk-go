# WhatsappMessageTemplateComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Describes the component type. One of &#x60;header&#x60;, &#x60;body&#x60;, &#x60;button&#x60;. For text-based templates, we only support the type&#x3D;body. | 
**SubType** | Pointer to **string** | **Required when type is &#x60;button&#x60;.** Type of button to create. &#x60;quick_reply&#x60;: Refers to a previously created quick reply button that allows for the customer to return a predefined message. &#x60;url&#x60;: Refers to a previously created button that allows the customer to visit the URL generated by appending the text parameter to the predefined prefix URL in the template. | [optional] 
**Index** | Pointer to **int32** | **Required when &#x60;type&#x60; &#x3D; &#x60;button&#x60;. Not used for the other types.** Position index of the button. You can have up to 3 buttons using index values of 0 to 2. | [optional] 
**Parameters** | Pointer to [**[]WhatsappMessageTemplateComponentParameter**](WhatsappMessageTemplateComponentParameter.md) | **Required when &#x60;type&#x60; &#x3D; &#x60;button&#x60;, or there are variables in the corresponding template component, or the template &#x60;HEADER&#x60; format is media (&#x60;IMAGE&#x60;, &#x60;VIDEO&#x60;, or &#x60;DOCUMENT&#x60;).** Array of parameter objects with the content of the message. For components of &#x60;type&#x60; &#x3D; &#x60;button&#x60;, see the [button parameter object](https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages#button-parameter-object). | [optional] 

## Methods

### NewWhatsappMessageTemplateComponent

`func NewWhatsappMessageTemplateComponent(type_ string, ) *WhatsappMessageTemplateComponent`

NewWhatsappMessageTemplateComponent instantiates a new WhatsappMessageTemplateComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageTemplateComponentWithDefaults

`func NewWhatsappMessageTemplateComponentWithDefaults() *WhatsappMessageTemplateComponent`

NewWhatsappMessageTemplateComponentWithDefaults instantiates a new WhatsappMessageTemplateComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappMessageTemplateComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageTemplateComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageTemplateComponent) SetType(v string)`

SetType sets Type field to given value.


### GetSubType

`func (o *WhatsappMessageTemplateComponent) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *WhatsappMessageTemplateComponent) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *WhatsappMessageTemplateComponent) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *WhatsappMessageTemplateComponent) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetIndex

`func (o *WhatsappMessageTemplateComponent) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *WhatsappMessageTemplateComponent) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *WhatsappMessageTemplateComponent) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *WhatsappMessageTemplateComponent) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetParameters

`func (o *WhatsappMessageTemplateComponent) GetParameters() []WhatsappMessageTemplateComponentParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *WhatsappMessageTemplateComponent) GetParametersOk() (*[]WhatsappMessageTemplateComponentParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *WhatsappMessageTemplateComponent) SetParameters(v []WhatsappMessageTemplateComponentParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *WhatsappMessageTemplateComponent) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

