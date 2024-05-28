# VoicesApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](VoicesApi.md#List) | **Get** /voices | List voice records
[**Send**](VoicesApi.md#Send) | **Post** /voices | Send a voice code



## List

> VoicePage List(ctx).Page(page).Limit(limit).IncludeTotal(includeTotal).FilterCreateTimeGte(filterCreateTimeGte).FilterCreateTimeLte(filterCreateTimeLte).FilterId(filterId).Execute()

List voice records



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    page := int32(56) // int32 | Page number of the results to be returned, 1-based. (optional) (default to 1)
    limit := int32(56) // int32 | A limit on the number of results to be returned, or number of results per page, between 1 and 100, defaults to 10. (optional) (default to 10)
    includeTotal := true // bool | Return results inside an object that contains the total result count or not. (optional) (default to false)
    filterCreateTimeGte := time.Now() // time.Time | Return results where the `createTime` field is greater than or equal to this value. Default: One day ago from now. (optional)
    filterCreateTimeLte := time.Now() // time.Time | Return results where the `createTime` field is less than or equal to this value. (optional)
    filterId := "filterId_example" // string | Unique object ID on our side. Other filter parameters will be ignored if this parameter is present. (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.VoicesApi.List(context.Background()).Page(page).Limit(limit).IncludeTotal(includeTotal).FilterCreateTimeGte(filterCreateTimeGte).FilterCreateTimeLte(filterCreateTimeLte).FilterId(filterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoicesApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: VoicePage
    fmt.Fprintf(os.Stdout, "Response from `VoicesApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of the results to be returned, 1-based. | [default to 1]
 **limit** | **int32** | A limit on the number of results to be returned, or number of results per page, between 1 and 100, defaults to 10. | [default to 10]
 **includeTotal** | **bool** | Return results inside an object that contains the total result count or not. | [default to false]
 **filterCreateTimeGte** | **time.Time** | Return results where the &#x60;createTime&#x60; field is greater than or equal to this value. Default: One day ago from now. | 
 **filterCreateTimeLte** | **time.Time** | Return results where the &#x60;createTime&#x60; field is less than or equal to this value. | 
 **filterId** | **string** | Unique object ID on our side. Other filter parameters will be ignored if this parameter is present. | 

### Return type

[**VoicePage**](VoicePage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Send

> Voice Send(ctx).VoiceSendRequest(voiceSendRequest).Execute()

Send a voice code



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    voiceSendRequest := *ycloud.NewVoiceSendRequest("+16315551111", "123456") // VoiceSendRequest | Voice call request that needs to be sent.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.VoicesApi.Send(context.Background()).VoiceSendRequest(voiceSendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoicesApi.Send``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Send`: Voice
    fmt.Fprintf(os.Stdout, "Response from `VoicesApi.Send`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceSendRequest** | [**VoiceSendRequest**](VoiceSendRequest.md) | Voice call request that needs to be sent. | 

### Return type

[**Voice**](Voice.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

