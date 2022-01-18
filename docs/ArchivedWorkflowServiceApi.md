# \ArchivedWorkflowServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchivedWorkflowServiceDeleteArchivedWorkflow**](ArchivedWorkflowServiceApi.md#ArchivedWorkflowServiceDeleteArchivedWorkflow) | **Delete** /workflow/argo/api/v1/archived-workflows/{uid} | 
[**ArchivedWorkflowServiceGetArchivedWorkflow**](ArchivedWorkflowServiceApi.md#ArchivedWorkflowServiceGetArchivedWorkflow) | **Get** /workflow/argo/api/v1/archived-workflows/{uid} | 
[**ArchivedWorkflowServiceListArchivedWorkflowLabelKeys**](ArchivedWorkflowServiceApi.md#ArchivedWorkflowServiceListArchivedWorkflowLabelKeys) | **Get** /workflow/argo/api/v1/archived-workflows-label-keys | 
[**ArchivedWorkflowServiceListArchivedWorkflowLabelValues**](ArchivedWorkflowServiceApi.md#ArchivedWorkflowServiceListArchivedWorkflowLabelValues) | **Get** /workflow/argo/api/v1/archived-workflows-label-values | 
[**ArchivedWorkflowServiceListArchivedWorkflows**](ArchivedWorkflowServiceApi.md#ArchivedWorkflowServiceListArchivedWorkflows) | **Get** /workflow/argo/api/v1/archived-workflows | 



## ArchivedWorkflowServiceDeleteArchivedWorkflow

> map[string]interface{} ArchivedWorkflowServiceDeleteArchivedWorkflow(ctx, uid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uid := "uid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivedWorkflowServiceApi.ArchivedWorkflowServiceDeleteArchivedWorkflow(context.Background(), uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivedWorkflowServiceApi.ArchivedWorkflowServiceDeleteArchivedWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchivedWorkflowServiceDeleteArchivedWorkflow`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ArchivedWorkflowServiceApi.ArchivedWorkflowServiceDeleteArchivedWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchivedWorkflowServiceDeleteArchivedWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchivedWorkflowServiceGetArchivedWorkflow

> IoArgoprojWorkflowV1alpha1Workflow ArchivedWorkflowServiceGetArchivedWorkflow(ctx, uid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uid := "uid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivedWorkflowServiceApi.ArchivedWorkflowServiceGetArchivedWorkflow(context.Background(), uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivedWorkflowServiceApi.ArchivedWorkflowServiceGetArchivedWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchivedWorkflowServiceGetArchivedWorkflow`: IoArgoprojWorkflowV1alpha1Workflow
    fmt.Fprintf(os.Stdout, "Response from `ArchivedWorkflowServiceApi.ArchivedWorkflowServiceGetArchivedWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchivedWorkflowServiceGetArchivedWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IoArgoprojWorkflowV1alpha1Workflow**](IoArgoprojWorkflowV1alpha1Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchivedWorkflowServiceListArchivedWorkflowLabelKeys

> IoArgoprojWorkflowV1alpha1LabelKeys ArchivedWorkflowServiceListArchivedWorkflowLabelKeys(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivedWorkflowServiceApi.ArchivedWorkflowServiceListArchivedWorkflowLabelKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivedWorkflowServiceApi.ArchivedWorkflowServiceListArchivedWorkflowLabelKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchivedWorkflowServiceListArchivedWorkflowLabelKeys`: IoArgoprojWorkflowV1alpha1LabelKeys
    fmt.Fprintf(os.Stdout, "Response from `ArchivedWorkflowServiceApi.ArchivedWorkflowServiceListArchivedWorkflowLabelKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiArchivedWorkflowServiceListArchivedWorkflowLabelKeysRequest struct via the builder pattern


### Return type

[**IoArgoprojWorkflowV1alpha1LabelKeys**](IoArgoprojWorkflowV1alpha1LabelKeys.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchivedWorkflowServiceListArchivedWorkflowLabelValues

> IoArgoprojWorkflowV1alpha1LabelValues ArchivedWorkflowServiceListArchivedWorkflowLabelValues(ctx).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    listOptionsLabelSelector := "listOptionsLabelSelector_example" // string | A selector to restrict the list of returned objects by their labels. Defaults to everything. +optional. (optional)
    listOptionsFieldSelector := "listOptionsFieldSelector_example" // string | A selector to restrict the list of returned objects by their fields. Defaults to everything. +optional. (optional)
    listOptionsWatch := true // bool | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. +optional. (optional)
    listOptionsAllowWatchBookmarks := true // bool | allowWatchBookmarks requests watch events with type \"BOOKMARK\". Servers that do not implement bookmarks may ignore this flag and bookmarks are sent at the server's discretion. Clients should not assume bookmarks are returned at any specific interval, nor may they assume the server will send any BOOKMARK event during a session. If this is not a watch, this field is ignored. If the feature gate WatchBookmarks is not enabled in apiserver, this field is ignored. +optional. (optional)
    listOptionsResourceVersion := "listOptionsResourceVersion_example" // string | resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional (optional)
    listOptionsResourceVersionMatch := "listOptionsResourceVersionMatch_example" // string | resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional (optional)
    listOptionsTimeoutSeconds := "listOptionsTimeoutSeconds_example" // string | Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. +optional. (optional)
    listOptionsLimit := "listOptionsLimit_example" // string | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. (optional)
    listOptionsContinue := "listOptionsContinue_example" // string | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".  This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivedWorkflowServiceApi.ArchivedWorkflowServiceListArchivedWorkflowLabelValues(context.Background()).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivedWorkflowServiceApi.ArchivedWorkflowServiceListArchivedWorkflowLabelValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchivedWorkflowServiceListArchivedWorkflowLabelValues`: IoArgoprojWorkflowV1alpha1LabelValues
    fmt.Fprintf(os.Stdout, "Response from `ArchivedWorkflowServiceApi.ArchivedWorkflowServiceListArchivedWorkflowLabelValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchivedWorkflowServiceListArchivedWorkflowLabelValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listOptionsLabelSelector** | **string** | A selector to restrict the list of returned objects by their labels. Defaults to everything. +optional. | 
 **listOptionsFieldSelector** | **string** | A selector to restrict the list of returned objects by their fields. Defaults to everything. +optional. | 
 **listOptionsWatch** | **bool** | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. +optional. | 
 **listOptionsAllowWatchBookmarks** | **bool** | allowWatchBookmarks requests watch events with type \&quot;BOOKMARK\&quot;. Servers that do not implement bookmarks may ignore this flag and bookmarks are sent at the server&#39;s discretion. Clients should not assume bookmarks are returned at any specific interval, nor may they assume the server will send any BOOKMARK event during a session. If this is not a watch, this field is ignored. If the feature gate WatchBookmarks is not enabled in apiserver, this field is ignored. +optional. | 
 **listOptionsResourceVersion** | **string** | resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional | 
 **listOptionsResourceVersionMatch** | **string** | resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional | 
 **listOptionsTimeoutSeconds** | **string** | Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. +optional. | 
 **listOptionsLimit** | **string** | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | 
 **listOptionsContinue** | **string** | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \&quot;next key\&quot;.  This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | 

### Return type

[**IoArgoprojWorkflowV1alpha1LabelValues**](IoArgoprojWorkflowV1alpha1LabelValues.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchivedWorkflowServiceListArchivedWorkflows

> IoArgoprojWorkflowV1alpha1WorkflowList ArchivedWorkflowServiceListArchivedWorkflows(ctx).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).NamePrefix(namePrefix).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    listOptionsLabelSelector := "listOptionsLabelSelector_example" // string | A selector to restrict the list of returned objects by their labels. Defaults to everything. +optional. (optional)
    listOptionsFieldSelector := "listOptionsFieldSelector_example" // string | A selector to restrict the list of returned objects by their fields. Defaults to everything. +optional. (optional)
    listOptionsWatch := true // bool | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. +optional. (optional)
    listOptionsAllowWatchBookmarks := true // bool | allowWatchBookmarks requests watch events with type \"BOOKMARK\". Servers that do not implement bookmarks may ignore this flag and bookmarks are sent at the server's discretion. Clients should not assume bookmarks are returned at any specific interval, nor may they assume the server will send any BOOKMARK event during a session. If this is not a watch, this field is ignored. If the feature gate WatchBookmarks is not enabled in apiserver, this field is ignored. +optional. (optional)
    listOptionsResourceVersion := "listOptionsResourceVersion_example" // string | resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional (optional)
    listOptionsResourceVersionMatch := "listOptionsResourceVersionMatch_example" // string | resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional (optional)
    listOptionsTimeoutSeconds := "listOptionsTimeoutSeconds_example" // string | Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. +optional. (optional)
    listOptionsLimit := "listOptionsLimit_example" // string | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. (optional)
    listOptionsContinue := "listOptionsContinue_example" // string | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".  This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. (optional)
    namePrefix := "namePrefix_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchivedWorkflowServiceApi.ArchivedWorkflowServiceListArchivedWorkflows(context.Background()).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).NamePrefix(namePrefix).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchivedWorkflowServiceApi.ArchivedWorkflowServiceListArchivedWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchivedWorkflowServiceListArchivedWorkflows`: IoArgoprojWorkflowV1alpha1WorkflowList
    fmt.Fprintf(os.Stdout, "Response from `ArchivedWorkflowServiceApi.ArchivedWorkflowServiceListArchivedWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchivedWorkflowServiceListArchivedWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listOptionsLabelSelector** | **string** | A selector to restrict the list of returned objects by their labels. Defaults to everything. +optional. | 
 **listOptionsFieldSelector** | **string** | A selector to restrict the list of returned objects by their fields. Defaults to everything. +optional. | 
 **listOptionsWatch** | **bool** | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. +optional. | 
 **listOptionsAllowWatchBookmarks** | **bool** | allowWatchBookmarks requests watch events with type \&quot;BOOKMARK\&quot;. Servers that do not implement bookmarks may ignore this flag and bookmarks are sent at the server&#39;s discretion. Clients should not assume bookmarks are returned at any specific interval, nor may they assume the server will send any BOOKMARK event during a session. If this is not a watch, this field is ignored. If the feature gate WatchBookmarks is not enabled in apiserver, this field is ignored. +optional. | 
 **listOptionsResourceVersion** | **string** | resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional | 
 **listOptionsResourceVersionMatch** | **string** | resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional | 
 **listOptionsTimeoutSeconds** | **string** | Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. +optional. | 
 **listOptionsLimit** | **string** | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | 
 **listOptionsContinue** | **string** | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \&quot;next key\&quot;.  This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | 
 **namePrefix** | **string** |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1WorkflowList**](IoArgoprojWorkflowV1alpha1WorkflowList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

