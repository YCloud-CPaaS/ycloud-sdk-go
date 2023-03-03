/*
YCloud API

The [YCloud](https://ycloud.com) API is organized around [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). Our API is designed to have predictable, resource-oriented URLs, return [JSON](https://www.json.org) responses, and use standard HTTP response codes and verbs.

API version: v2
Contact: service@ycloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ycloud

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// WhatsappTemplatesApiService WhatsappTemplatesApi service
type WhatsappTemplatesApiService service

type WhatsappTemplatesApiCreateRequest struct {
	ctx context.Context
	ApiService *WhatsappTemplatesApiService
	whatsappTemplateCreateRequest *WhatsappTemplateCreateRequest
}

func (r WhatsappTemplatesApiCreateRequest) WhatsappTemplateCreateRequest(whatsappTemplateCreateRequest WhatsappTemplateCreateRequest) WhatsappTemplatesApiCreateRequest {
	r.whatsappTemplateCreateRequest = &whatsappTemplateCreateRequest
	return r
}

func (r WhatsappTemplatesApiCreateRequest) Execute() (*WhatsappTemplate, *http.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create a WhatsApp template

Creates a WhatsApp template.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return WhatsappTemplatesApiCreateRequest
*/
func (a *WhatsappTemplatesApiService) Create(ctx context.Context) WhatsappTemplatesApiCreateRequest {
	return WhatsappTemplatesApiCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WhatsappTemplate
func (a *WhatsappTemplatesApiService) CreateExecute(r WhatsappTemplatesApiCreateRequest) (*WhatsappTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WhatsappTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappTemplatesApiService.Create")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.whatsappTemplateCreateRequest == nil {
		return localVarReturnValue, nil, reportError("whatsappTemplateCreateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.whatsappTemplateCreateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type WhatsappTemplatesApiDeleteByNameRequest struct {
	ctx context.Context
	ApiService *WhatsappTemplatesApiService
	wabaId string
	name string
}

func (r WhatsappTemplatesApiDeleteByNameRequest) Execute() ([]WhatsappTemplate, *http.Response, error) {
	return r.ApiService.DeleteByNameExecute(r)
}

/*
DeleteByName Delete WhatsApp templates by name

Deletes WhatsApp templates by name. If that template name exists in multiple languages, all languages will be deleted.
HTTP status `404` is returned if no templates are found for the specific name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wabaId WhatsApp Business Account ID.
 @param name Name of the template.
 @return WhatsappTemplatesApiDeleteByNameRequest
*/
func (a *WhatsappTemplatesApiService) DeleteByName(ctx context.Context, wabaId string, name string) WhatsappTemplatesApiDeleteByNameRequest {
	return WhatsappTemplatesApiDeleteByNameRequest{
		ApiService: a,
		ctx: ctx,
		wabaId: wabaId,
		name: name,
	}
}

// Execute executes the request
//  @return []WhatsappTemplate
func (a *WhatsappTemplatesApiService) DeleteByNameExecute(r WhatsappTemplatesApiDeleteByNameRequest) ([]WhatsappTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []WhatsappTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappTemplatesApiService.DeleteByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/templates/{wabaId}/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"wabaId"+"}", url.PathEscape(parameterToString(r.wabaId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type WhatsappTemplatesApiEditByNameAndLanguageRequest struct {
	ctx context.Context
	ApiService *WhatsappTemplatesApiService
	wabaId string
	name string
	language string
	whatsappTemplateEditRequest *WhatsappTemplateEditRequest
}

func (r WhatsappTemplatesApiEditByNameAndLanguageRequest) WhatsappTemplateEditRequest(whatsappTemplateEditRequest WhatsappTemplateEditRequest) WhatsappTemplatesApiEditByNameAndLanguageRequest {
	r.whatsappTemplateEditRequest = &whatsappTemplateEditRequest
	return r
}

func (r WhatsappTemplatesApiEditByNameAndLanguageRequest) Execute() (*WhatsappTemplate, *http.Response, error) {
	return r.ApiService.EditByNameAndLanguageExecute(r)
}

/*
EditByNameAndLanguage Edit a WhatsApp template

Edits a WhatsApp template by name and language.
Editing a template replaces its old contents entirely, so include any components you wish to preserve as well as components you wish to update using the components parameter.
See also [Edit a Message Template](https://developers.facebook.com/docs/whatsapp/business-management-api/message-templates#edit-a-message-template).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wabaId WhatsApp Business Account ID.
 @param name Name of the template.
 @param language Language code of the template. See [Supported Languages](https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#supported-languages-) for all codes.
 @return WhatsappTemplatesApiEditByNameAndLanguageRequest
*/
func (a *WhatsappTemplatesApiService) EditByNameAndLanguage(ctx context.Context, wabaId string, name string, language string) WhatsappTemplatesApiEditByNameAndLanguageRequest {
	return WhatsappTemplatesApiEditByNameAndLanguageRequest{
		ApiService: a,
		ctx: ctx,
		wabaId: wabaId,
		name: name,
		language: language,
	}
}

// Execute executes the request
//  @return WhatsappTemplate
func (a *WhatsappTemplatesApiService) EditByNameAndLanguageExecute(r WhatsappTemplatesApiEditByNameAndLanguageRequest) (*WhatsappTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WhatsappTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappTemplatesApiService.EditByNameAndLanguage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/templates/{wabaId}/{name}/{language}"
	localVarPath = strings.Replace(localVarPath, "{"+"wabaId"+"}", url.PathEscape(parameterToString(r.wabaId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterToString(r.language, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.whatsappTemplateEditRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type WhatsappTemplatesApiListRequest struct {
	ctx context.Context
	ApiService *WhatsappTemplatesApiService
	page *int32
	limit *int32
	includeTotal *bool
	filterWabaId *string
	filterName *string
	filterLanguage *string
}

// Page number of the results to be returned, 1-based.
func (r WhatsappTemplatesApiListRequest) Page(page int32) WhatsappTemplatesApiListRequest {
	r.page = &page
	return r
}

// A limit on the number of results to be returned, or number of results per page, between 1 and 100, defaults to 10.
func (r WhatsappTemplatesApiListRequest) Limit(limit int32) WhatsappTemplatesApiListRequest {
	r.limit = &limit
	return r
}

// Return results inside an object that contains the total result count or not.
func (r WhatsappTemplatesApiListRequest) IncludeTotal(includeTotal bool) WhatsappTemplatesApiListRequest {
	r.includeTotal = &includeTotal
	return r
}

// WhatsApp Business Account ID.
func (r WhatsappTemplatesApiListRequest) FilterWabaId(filterWabaId string) WhatsappTemplatesApiListRequest {
	r.filterWabaId = &filterWabaId
	return r
}

// Name of the template.
func (r WhatsappTemplatesApiListRequest) FilterName(filterName string) WhatsappTemplatesApiListRequest {
	r.filterName = &filterName
	return r
}

// Language of the template.
func (r WhatsappTemplatesApiListRequest) FilterLanguage(filterLanguage string) WhatsappTemplatesApiListRequest {
	r.filterLanguage = &filterLanguage
	return r
}

func (r WhatsappTemplatesApiListRequest) Execute() (*WhatsappTemplatePage, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List List WhatsApp templates

Returns a paginated list of WhatsApp templates you've previously created.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return WhatsappTemplatesApiListRequest
*/
func (a *WhatsappTemplatesApiService) List(ctx context.Context) WhatsappTemplatesApiListRequest {
	return WhatsappTemplatesApiListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WhatsappTemplatePage
func (a *WhatsappTemplatesApiService) ListExecute(r WhatsappTemplatesApiListRequest) (*WhatsappTemplatePage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WhatsappTemplatePage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappTemplatesApiService.List")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.includeTotal != nil {
		localVarQueryParams.Add("includeTotal", parameterToString(*r.includeTotal, ""))
	}
	if r.filterWabaId != nil {
		localVarQueryParams.Add("filter.wabaId", parameterToString(*r.filterWabaId, ""))
	}
	if r.filterName != nil {
		localVarQueryParams.Add("filter.name", parameterToString(*r.filterName, ""))
	}
	if r.filterLanguage != nil {
		localVarQueryParams.Add("filter.language", parameterToString(*r.filterLanguage, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type WhatsappTemplatesApiRetrieveByNameAndLanguageRequest struct {
	ctx context.Context
	ApiService *WhatsappTemplatesApiService
	wabaId string
	name string
	language string
}

func (r WhatsappTemplatesApiRetrieveByNameAndLanguageRequest) Execute() (*WhatsappTemplate, *http.Response, error) {
	return r.ApiService.RetrieveByNameAndLanguageExecute(r)
}

/*
RetrieveByNameAndLanguage Retrieve a WhatsApp template

Retrieves a WhatsApp template by name and language.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wabaId WhatsApp Business Account ID.
 @param name Name of the template.
 @param language Language code of the template. See [Supported Languages](https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#supported-languages-) for all codes.
 @return WhatsappTemplatesApiRetrieveByNameAndLanguageRequest
*/
func (a *WhatsappTemplatesApiService) RetrieveByNameAndLanguage(ctx context.Context, wabaId string, name string, language string) WhatsappTemplatesApiRetrieveByNameAndLanguageRequest {
	return WhatsappTemplatesApiRetrieveByNameAndLanguageRequest{
		ApiService: a,
		ctx: ctx,
		wabaId: wabaId,
		name: name,
		language: language,
	}
}

// Execute executes the request
//  @return WhatsappTemplate
func (a *WhatsappTemplatesApiService) RetrieveByNameAndLanguageExecute(r WhatsappTemplatesApiRetrieveByNameAndLanguageRequest) (*WhatsappTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WhatsappTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappTemplatesApiService.RetrieveByNameAndLanguage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/templates/{wabaId}/{name}/{language}"
	localVarPath = strings.Replace(localVarPath, "{"+"wabaId"+"}", url.PathEscape(parameterToString(r.wabaId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"language"+"}", url.PathEscape(parameterToString(r.language, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
