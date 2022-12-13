# WhatsappInboundMessageReferral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceUrl** | Pointer to **string** | Specifies the URL that leads to the ad or post clicked by the user. Opening this URL takes you to the ad viewed by your user. | [optional] 
**SourceType** | Pointer to **string** | Specifies the type of the ad&#39;s source. Supported values are \&quot;ad\&quot; or \&quot;post\&quot;. | [optional] 
**SourceId** | Pointer to **string** | Specifies the Meta ID for an ad or post. | [optional] 
**Headline** | Pointer to **string** | Specifies the headline used in the ad or post that generated the message. | [optional] 
**Body** | Pointer to **string** | The description, or body, from the ad or post that generated the message. | [optional] 
**MediaType** | Pointer to **string** | Media present in the ad or post the user clicked. Supported values are \&quot;image\&quot; or \&quot;video\&quot;. | [optional] 
**ImageUrl** | Pointer to **string** | **Added if media_type is \&quot;image\&quot;.**  Contains a URL to the raw image. | [optional] 
**VideoUrl** | Pointer to **string** | **Added if media_type is \&quot;video\&quot;.**  Contains a URL to the video. | [optional] 
**ThumbnailUrl** | Pointer to **string** | **Added if media_type is \&quot;video\&quot;.**  Contains a URL to the thumbnail image of the clicked video. | [optional] 

## Methods

### NewWhatsappInboundMessageReferral

`func NewWhatsappInboundMessageReferral() *WhatsappInboundMessageReferral`

NewWhatsappInboundMessageReferral instantiates a new WhatsappInboundMessageReferral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappInboundMessageReferralWithDefaults

`func NewWhatsappInboundMessageReferralWithDefaults() *WhatsappInboundMessageReferral`

NewWhatsappInboundMessageReferralWithDefaults instantiates a new WhatsappInboundMessageReferral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceUrl

`func (o *WhatsappInboundMessageReferral) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *WhatsappInboundMessageReferral) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *WhatsappInboundMessageReferral) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *WhatsappInboundMessageReferral) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetSourceType

`func (o *WhatsappInboundMessageReferral) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *WhatsappInboundMessageReferral) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *WhatsappInboundMessageReferral) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *WhatsappInboundMessageReferral) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSourceId

`func (o *WhatsappInboundMessageReferral) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *WhatsappInboundMessageReferral) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *WhatsappInboundMessageReferral) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *WhatsappInboundMessageReferral) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetHeadline

`func (o *WhatsappInboundMessageReferral) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *WhatsappInboundMessageReferral) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *WhatsappInboundMessageReferral) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *WhatsappInboundMessageReferral) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetBody

`func (o *WhatsappInboundMessageReferral) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WhatsappInboundMessageReferral) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WhatsappInboundMessageReferral) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *WhatsappInboundMessageReferral) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetMediaType

`func (o *WhatsappInboundMessageReferral) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *WhatsappInboundMessageReferral) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *WhatsappInboundMessageReferral) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *WhatsappInboundMessageReferral) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetImageUrl

`func (o *WhatsappInboundMessageReferral) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *WhatsappInboundMessageReferral) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *WhatsappInboundMessageReferral) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *WhatsappInboundMessageReferral) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetVideoUrl

`func (o *WhatsappInboundMessageReferral) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *WhatsappInboundMessageReferral) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *WhatsappInboundMessageReferral) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *WhatsappInboundMessageReferral) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *WhatsappInboundMessageReferral) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *WhatsappInboundMessageReferral) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *WhatsappInboundMessageReferral) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *WhatsappInboundMessageReferral) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

