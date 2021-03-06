/*
Argo Server API

You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`

API version: VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ArchivedWorkflowServiceApiService ArchivedWorkflowServiceApi service
type ArchivedWorkflowServiceApiService service

type ApiArchivedWorkflowServiceDeleteArchivedWorkflowRequest struct {
	ctx _context.Context
	ApiService *ArchivedWorkflowServiceApiService
	uid string
}


func (r ApiArchivedWorkflowServiceDeleteArchivedWorkflowRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.ArchivedWorkflowServiceDeleteArchivedWorkflowExecute(r)
}

/*
ArchivedWorkflowServiceDeleteArchivedWorkflow Method for ArchivedWorkflowServiceDeleteArchivedWorkflow

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uid
 @return ApiArchivedWorkflowServiceDeleteArchivedWorkflowRequest
*/
func (a *ArchivedWorkflowServiceApiService) ArchivedWorkflowServiceDeleteArchivedWorkflow(ctx _context.Context, uid string) ApiArchivedWorkflowServiceDeleteArchivedWorkflowRequest {
	return ApiArchivedWorkflowServiceDeleteArchivedWorkflowRequest{
		ApiService: a,
		ctx: ctx,
		uid: uid,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ArchivedWorkflowServiceApiService) ArchivedWorkflowServiceDeleteArchivedWorkflowExecute(r ApiArchivedWorkflowServiceDeleteArchivedWorkflowRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivedWorkflowServiceApiService.ArchivedWorkflowServiceDeleteArchivedWorkflow")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflow/argo/api/v1/archived-workflows/{uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", _neturl.PathEscape(parameterToString(r.uid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GrpcGatewayRuntimeError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiArchivedWorkflowServiceGetArchivedWorkflowRequest struct {
	ctx _context.Context
	ApiService *ArchivedWorkflowServiceApiService
	uid string
}


func (r ApiArchivedWorkflowServiceGetArchivedWorkflowRequest) Execute() (IoArgoprojWorkflowV1alpha1Workflow, *_nethttp.Response, error) {
	return r.ApiService.ArchivedWorkflowServiceGetArchivedWorkflowExecute(r)
}

/*
ArchivedWorkflowServiceGetArchivedWorkflow Method for ArchivedWorkflowServiceGetArchivedWorkflow

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uid
 @return ApiArchivedWorkflowServiceGetArchivedWorkflowRequest
*/
func (a *ArchivedWorkflowServiceApiService) ArchivedWorkflowServiceGetArchivedWorkflow(ctx _context.Context, uid string) ApiArchivedWorkflowServiceGetArchivedWorkflowRequest {
	return ApiArchivedWorkflowServiceGetArchivedWorkflowRequest{
		ApiService: a,
		ctx: ctx,
		uid: uid,
	}
}

// Execute executes the request
//  @return IoArgoprojWorkflowV1alpha1Workflow
func (a *ArchivedWorkflowServiceApiService) ArchivedWorkflowServiceGetArchivedWorkflowExecute(r ApiArchivedWorkflowServiceGetArchivedWorkflowRequest) (IoArgoprojWorkflowV1alpha1Workflow, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IoArgoprojWorkflowV1alpha1Workflow
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivedWorkflowServiceApiService.ArchivedWorkflowServiceGetArchivedWorkflow")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflow/argo/api/v1/archived-workflows/{uid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uid"+"}", _neturl.PathEscape(parameterToString(r.uid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GrpcGatewayRuntimeError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiArchivedWorkflowServiceListArchivedWorkflowLabelKeysRequest struct {
	ctx _context.Context
	ApiService *ArchivedWorkflowServiceApiService
}


func (r ApiArchivedWorkflowServiceListArchivedWorkflowLabelKeysRequest) Execute() (IoArgoprojWorkflowV1alpha1LabelKeys, *_nethttp.Response, error) {
	return r.ApiService.ArchivedWorkflowServiceListArchivedWorkflowLabelKeysExecute(r)
}

/*
ArchivedWorkflowServiceListArchivedWorkflowLabelKeys Method for ArchivedWorkflowServiceListArchivedWorkflowLabelKeys

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiArchivedWorkflowServiceListArchivedWorkflowLabelKeysRequest
*/
func (a *ArchivedWorkflowServiceApiService) ArchivedWorkflowServiceListArchivedWorkflowLabelKeys(ctx _context.Context) ApiArchivedWorkflowServiceListArchivedWorkflowLabelKeysRequest {
	return ApiArchivedWorkflowServiceListArchivedWorkflowLabelKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IoArgoprojWorkflowV1alpha1LabelKeys
func (a *ArchivedWorkflowServiceApiService) ArchivedWorkflowServiceListArchivedWorkflowLabelKeysExecute(r ApiArchivedWorkflowServiceListArchivedWorkflowLabelKeysRequest) (IoArgoprojWorkflowV1alpha1LabelKeys, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IoArgoprojWorkflowV1alpha1LabelKeys
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivedWorkflowServiceApiService.ArchivedWorkflowServiceListArchivedWorkflowLabelKeys")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflow/argo/api/v1/archived-workflows-label-keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GrpcGatewayRuntimeError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest struct {
	ctx _context.Context
	ApiService *ArchivedWorkflowServiceApiService
	listOptionsLabelSelector *string
	listOptionsFieldSelector *string
	listOptionsWatch *bool
	listOptionsAllowWatchBookmarks *bool
	listOptionsResourceVersion *string
	listOptionsResourceVersionMatch *string
	listOptionsTimeoutSeconds *string
	listOptionsLimit *string
	listOptionsContinue *string
}

// A selector to restrict the list of returned objects by their labels. Defaults to everything. +optional.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest) ListOptionsLabelSelector(listOptionsLabelSelector string) ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest {
	r.listOptionsLabelSelector = &listOptionsLabelSelector
	return r
}
// A selector to restrict the list of returned objects by their fields. Defaults to everything. +optional.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest) ListOptionsFieldSelector(listOptionsFieldSelector string) ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest {
	r.listOptionsFieldSelector = &listOptionsFieldSelector
	return r
}
// Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. +optional.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest) ListOptionsWatch(listOptionsWatch bool) ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest {
	r.listOptionsWatch = &listOptionsWatch
	return r
}
// allowWatchBookmarks requests watch events with type \&quot;BOOKMARK\&quot;. Servers that do not implement bookmarks may ignore this flag and bookmarks are sent at the server&#39;s discretion. Clients should not assume bookmarks are returned at any specific interval, nor may they assume the server will send any BOOKMARK event during a session. If this is not a watch, this field is ignored. If the feature gate WatchBookmarks is not enabled in apiserver, this field is ignored. +optional.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest) ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks bool) ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest {
	r.listOptionsAllowWatchBookmarks = &listOptionsAllowWatchBookmarks
	return r
}
// resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional
func (r ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest) ListOptionsResourceVersion(listOptionsResourceVersion string) ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest {
	r.listOptionsResourceVersion = &listOptionsResourceVersion
	return r
}
// resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional
func (r ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest) ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch string) ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest {
	r.listOptionsResourceVersionMatch = &listOptionsResourceVersionMatch
	return r
}
// Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. +optional.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest) ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds string) ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest {
	r.listOptionsTimeoutSeconds = &listOptionsTimeoutSeconds
	return r
}
// limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest) ListOptionsLimit(listOptionsLimit string) ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest {
	r.listOptionsLimit = &listOptionsLimit
	return r
}
// The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \&quot;next key\&quot;.  This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest) ListOptionsContinue(listOptionsContinue string) ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest {
	r.listOptionsContinue = &listOptionsContinue
	return r
}

func (r ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest) Execute() (IoArgoprojWorkflowV1alpha1LabelValues, *_nethttp.Response, error) {
	return r.ApiService.ArchivedWorkflowServiceListArchivedWorkflowLabelValuesExecute(r)
}

/*
ArchivedWorkflowServiceListArchivedWorkflowLabelValues Method for ArchivedWorkflowServiceListArchivedWorkflowLabelValues

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest
*/
func (a *ArchivedWorkflowServiceApiService) ArchivedWorkflowServiceListArchivedWorkflowLabelValues(ctx _context.Context) ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest {
	return ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IoArgoprojWorkflowV1alpha1LabelValues
func (a *ArchivedWorkflowServiceApiService) ArchivedWorkflowServiceListArchivedWorkflowLabelValuesExecute(r ApiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest) (IoArgoprojWorkflowV1alpha1LabelValues, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IoArgoprojWorkflowV1alpha1LabelValues
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivedWorkflowServiceApiService.ArchivedWorkflowServiceListArchivedWorkflowLabelValues")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflow/argo/api/v1/archived-workflows-label-values"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.listOptionsLabelSelector != nil {
		localVarQueryParams.Add("listOptions.labelSelector", parameterToString(*r.listOptionsLabelSelector, ""))
	}
	if r.listOptionsFieldSelector != nil {
		localVarQueryParams.Add("listOptions.fieldSelector", parameterToString(*r.listOptionsFieldSelector, ""))
	}
	if r.listOptionsWatch != nil {
		localVarQueryParams.Add("listOptions.watch", parameterToString(*r.listOptionsWatch, ""))
	}
	if r.listOptionsAllowWatchBookmarks != nil {
		localVarQueryParams.Add("listOptions.allowWatchBookmarks", parameterToString(*r.listOptionsAllowWatchBookmarks, ""))
	}
	if r.listOptionsResourceVersion != nil {
		localVarQueryParams.Add("listOptions.resourceVersion", parameterToString(*r.listOptionsResourceVersion, ""))
	}
	if r.listOptionsResourceVersionMatch != nil {
		localVarQueryParams.Add("listOptions.resourceVersionMatch", parameterToString(*r.listOptionsResourceVersionMatch, ""))
	}
	if r.listOptionsTimeoutSeconds != nil {
		localVarQueryParams.Add("listOptions.timeoutSeconds", parameterToString(*r.listOptionsTimeoutSeconds, ""))
	}
	if r.listOptionsLimit != nil {
		localVarQueryParams.Add("listOptions.limit", parameterToString(*r.listOptionsLimit, ""))
	}
	if r.listOptionsContinue != nil {
		localVarQueryParams.Add("listOptions.continue", parameterToString(*r.listOptionsContinue, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GrpcGatewayRuntimeError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiArchivedWorkflowServiceListArchivedWorkflowsRequest struct {
	ctx _context.Context
	ApiService *ArchivedWorkflowServiceApiService
	listOptionsLabelSelector *string
	listOptionsFieldSelector *string
	listOptionsWatch *bool
	listOptionsAllowWatchBookmarks *bool
	listOptionsResourceVersion *string
	listOptionsResourceVersionMatch *string
	listOptionsTimeoutSeconds *string
	listOptionsLimit *string
	listOptionsContinue *string
	namePrefix *string
}

// A selector to restrict the list of returned objects by their labels. Defaults to everything. +optional.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowsRequest) ListOptionsLabelSelector(listOptionsLabelSelector string) ApiArchivedWorkflowServiceListArchivedWorkflowsRequest {
	r.listOptionsLabelSelector = &listOptionsLabelSelector
	return r
}
// A selector to restrict the list of returned objects by their fields. Defaults to everything. +optional.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowsRequest) ListOptionsFieldSelector(listOptionsFieldSelector string) ApiArchivedWorkflowServiceListArchivedWorkflowsRequest {
	r.listOptionsFieldSelector = &listOptionsFieldSelector
	return r
}
// Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. +optional.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowsRequest) ListOptionsWatch(listOptionsWatch bool) ApiArchivedWorkflowServiceListArchivedWorkflowsRequest {
	r.listOptionsWatch = &listOptionsWatch
	return r
}
// allowWatchBookmarks requests watch events with type \&quot;BOOKMARK\&quot;. Servers that do not implement bookmarks may ignore this flag and bookmarks are sent at the server&#39;s discretion. Clients should not assume bookmarks are returned at any specific interval, nor may they assume the server will send any BOOKMARK event during a session. If this is not a watch, this field is ignored. If the feature gate WatchBookmarks is not enabled in apiserver, this field is ignored. +optional.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowsRequest) ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks bool) ApiArchivedWorkflowServiceListArchivedWorkflowsRequest {
	r.listOptionsAllowWatchBookmarks = &listOptionsAllowWatchBookmarks
	return r
}
// resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional
func (r ApiArchivedWorkflowServiceListArchivedWorkflowsRequest) ListOptionsResourceVersion(listOptionsResourceVersion string) ApiArchivedWorkflowServiceListArchivedWorkflowsRequest {
	r.listOptionsResourceVersion = &listOptionsResourceVersion
	return r
}
// resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional
func (r ApiArchivedWorkflowServiceListArchivedWorkflowsRequest) ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch string) ApiArchivedWorkflowServiceListArchivedWorkflowsRequest {
	r.listOptionsResourceVersionMatch = &listOptionsResourceVersionMatch
	return r
}
// Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. +optional.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowsRequest) ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds string) ApiArchivedWorkflowServiceListArchivedWorkflowsRequest {
	r.listOptionsTimeoutSeconds = &listOptionsTimeoutSeconds
	return r
}
// limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowsRequest) ListOptionsLimit(listOptionsLimit string) ApiArchivedWorkflowServiceListArchivedWorkflowsRequest {
	r.listOptionsLimit = &listOptionsLimit
	return r
}
// The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \&quot;next key\&quot;.  This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.
func (r ApiArchivedWorkflowServiceListArchivedWorkflowsRequest) ListOptionsContinue(listOptionsContinue string) ApiArchivedWorkflowServiceListArchivedWorkflowsRequest {
	r.listOptionsContinue = &listOptionsContinue
	return r
}
func (r ApiArchivedWorkflowServiceListArchivedWorkflowsRequest) NamePrefix(namePrefix string) ApiArchivedWorkflowServiceListArchivedWorkflowsRequest {
	r.namePrefix = &namePrefix
	return r
}

func (r ApiArchivedWorkflowServiceListArchivedWorkflowsRequest) Execute() (IoArgoprojWorkflowV1alpha1WorkflowList, *_nethttp.Response, error) {
	return r.ApiService.ArchivedWorkflowServiceListArchivedWorkflowsExecute(r)
}

/*
ArchivedWorkflowServiceListArchivedWorkflows Method for ArchivedWorkflowServiceListArchivedWorkflows

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiArchivedWorkflowServiceListArchivedWorkflowsRequest
*/
func (a *ArchivedWorkflowServiceApiService) ArchivedWorkflowServiceListArchivedWorkflows(ctx _context.Context) ApiArchivedWorkflowServiceListArchivedWorkflowsRequest {
	return ApiArchivedWorkflowServiceListArchivedWorkflowsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IoArgoprojWorkflowV1alpha1WorkflowList
func (a *ArchivedWorkflowServiceApiService) ArchivedWorkflowServiceListArchivedWorkflowsExecute(r ApiArchivedWorkflowServiceListArchivedWorkflowsRequest) (IoArgoprojWorkflowV1alpha1WorkflowList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IoArgoprojWorkflowV1alpha1WorkflowList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArchivedWorkflowServiceApiService.ArchivedWorkflowServiceListArchivedWorkflows")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/workflow/argo/api/v1/archived-workflows"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.listOptionsLabelSelector != nil {
		localVarQueryParams.Add("listOptions.labelSelector", parameterToString(*r.listOptionsLabelSelector, ""))
	}
	if r.listOptionsFieldSelector != nil {
		localVarQueryParams.Add("listOptions.fieldSelector", parameterToString(*r.listOptionsFieldSelector, ""))
	}
	if r.listOptionsWatch != nil {
		localVarQueryParams.Add("listOptions.watch", parameterToString(*r.listOptionsWatch, ""))
	}
	if r.listOptionsAllowWatchBookmarks != nil {
		localVarQueryParams.Add("listOptions.allowWatchBookmarks", parameterToString(*r.listOptionsAllowWatchBookmarks, ""))
	}
	if r.listOptionsResourceVersion != nil {
		localVarQueryParams.Add("listOptions.resourceVersion", parameterToString(*r.listOptionsResourceVersion, ""))
	}
	if r.listOptionsResourceVersionMatch != nil {
		localVarQueryParams.Add("listOptions.resourceVersionMatch", parameterToString(*r.listOptionsResourceVersionMatch, ""))
	}
	if r.listOptionsTimeoutSeconds != nil {
		localVarQueryParams.Add("listOptions.timeoutSeconds", parameterToString(*r.listOptionsTimeoutSeconds, ""))
	}
	if r.listOptionsLimit != nil {
		localVarQueryParams.Add("listOptions.limit", parameterToString(*r.listOptionsLimit, ""))
	}
	if r.listOptionsContinue != nil {
		localVarQueryParams.Add("listOptions.continue", parameterToString(*r.listOptionsContinue, ""))
	}
	if r.namePrefix != nil {
		localVarQueryParams.Add("namePrefix", parameterToString(*r.namePrefix, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v GrpcGatewayRuntimeError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
