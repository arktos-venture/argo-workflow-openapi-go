# \WorkflowTemplateServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowTemplateServiceCreateWorkflowTemplate**](WorkflowTemplateServiceApi.md#WorkflowTemplateServiceCreateWorkflowTemplate) | **Post** /workflow/argo/api/v1/workflow-templates/{namespace} | 
[**WorkflowTemplateServiceDeleteWorkflowTemplate**](WorkflowTemplateServiceApi.md#WorkflowTemplateServiceDeleteWorkflowTemplate) | **Delete** /workflow/argo/api/v1/workflow-templates/{namespace}/{name} | 
[**WorkflowTemplateServiceGetWorkflowTemplate**](WorkflowTemplateServiceApi.md#WorkflowTemplateServiceGetWorkflowTemplate) | **Get** /workflow/argo/api/v1/workflow-templates/{namespace}/{name} | 
[**WorkflowTemplateServiceLintWorkflowTemplate**](WorkflowTemplateServiceApi.md#WorkflowTemplateServiceLintWorkflowTemplate) | **Post** /workflow/argo/api/v1/workflow-templates/{namespace}/lint | 
[**WorkflowTemplateServiceListWorkflowTemplates**](WorkflowTemplateServiceApi.md#WorkflowTemplateServiceListWorkflowTemplates) | **Get** /workflow/argo/api/v1/workflow-templates/{namespace} | 
[**WorkflowTemplateServiceUpdateWorkflowTemplate**](WorkflowTemplateServiceApi.md#WorkflowTemplateServiceUpdateWorkflowTemplate) | **Put** /workflow/argo/api/v1/workflow-templates/{namespace}/{name} | 



## WorkflowTemplateServiceCreateWorkflowTemplate

> IoArgoprojWorkflowV1alpha1WorkflowTemplate WorkflowTemplateServiceCreateWorkflowTemplate(ctx, namespace).Body(body).Execute()



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
    namespace := "namespace_example" // string | 
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest() // IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowTemplateServiceApi.WorkflowTemplateServiceCreateWorkflowTemplate(context.Background(), namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTemplateServiceApi.WorkflowTemplateServiceCreateWorkflowTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowTemplateServiceCreateWorkflowTemplate`: IoArgoprojWorkflowV1alpha1WorkflowTemplate
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTemplateServiceApi.WorkflowTemplateServiceCreateWorkflowTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowTemplateServiceCreateWorkflowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest**](IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1WorkflowTemplate**](IoArgoprojWorkflowV1alpha1WorkflowTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowTemplateServiceDeleteWorkflowTemplate

> map[string]interface{} WorkflowTemplateServiceDeleteWorkflowTemplate(ctx, namespace, name).DeleteOptionsGracePeriodSeconds(deleteOptionsGracePeriodSeconds).DeleteOptionsPreconditionsUid(deleteOptionsPreconditionsUid).DeleteOptionsPreconditionsResourceVersion(deleteOptionsPreconditionsResourceVersion).DeleteOptionsOrphanDependents(deleteOptionsOrphanDependents).DeleteOptionsPropagationPolicy(deleteOptionsPropagationPolicy).DeleteOptionsDryRun(deleteOptionsDryRun).Execute()



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
    namespace := "namespace_example" // string | 
    name := "name_example" // string | 
    deleteOptionsGracePeriodSeconds := "deleteOptionsGracePeriodSeconds_example" // string | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. +optional. (optional)
    deleteOptionsPreconditionsUid := "deleteOptionsPreconditionsUid_example" // string | Specifies the target UID. +optional. (optional)
    deleteOptionsPreconditionsResourceVersion := "deleteOptionsPreconditionsResourceVersion_example" // string | Specifies the target ResourceVersion +optional. (optional)
    deleteOptionsOrphanDependents := true // bool | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both. +optional. (optional)
    deleteOptionsPropagationPolicy := "deleteOptionsPropagationPolicy_example" // string | Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground. +optional. (optional)
    deleteOptionsDryRun := []string{"Inner_example"} // []string | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed +optional. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowTemplateServiceApi.WorkflowTemplateServiceDeleteWorkflowTemplate(context.Background(), namespace, name).DeleteOptionsGracePeriodSeconds(deleteOptionsGracePeriodSeconds).DeleteOptionsPreconditionsUid(deleteOptionsPreconditionsUid).DeleteOptionsPreconditionsResourceVersion(deleteOptionsPreconditionsResourceVersion).DeleteOptionsOrphanDependents(deleteOptionsOrphanDependents).DeleteOptionsPropagationPolicy(deleteOptionsPropagationPolicy).DeleteOptionsDryRun(deleteOptionsDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTemplateServiceApi.WorkflowTemplateServiceDeleteWorkflowTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowTemplateServiceDeleteWorkflowTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTemplateServiceApi.WorkflowTemplateServiceDeleteWorkflowTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowTemplateServiceDeleteWorkflowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteOptionsGracePeriodSeconds** | **string** | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. +optional. | 
 **deleteOptionsPreconditionsUid** | **string** | Specifies the target UID. +optional. | 
 **deleteOptionsPreconditionsResourceVersion** | **string** | Specifies the target ResourceVersion +optional. | 
 **deleteOptionsOrphanDependents** | **bool** | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. +optional. | 
 **deleteOptionsPropagationPolicy** | **string** | Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: &#39;Orphan&#39; - orphan the dependents; &#39;Background&#39; - allow the garbage collector to delete the dependents in the background; &#39;Foreground&#39; - a cascading policy that deletes all dependents in the foreground. +optional. | 
 **deleteOptionsDryRun** | **[]string** | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed +optional. | 

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


## WorkflowTemplateServiceGetWorkflowTemplate

> IoArgoprojWorkflowV1alpha1WorkflowTemplate WorkflowTemplateServiceGetWorkflowTemplate(ctx, namespace, name).GetOptionsResourceVersion(getOptionsResourceVersion).Execute()



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
    namespace := "namespace_example" // string | 
    name := "name_example" // string | 
    getOptionsResourceVersion := "getOptionsResourceVersion_example" // string | resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowTemplateServiceApi.WorkflowTemplateServiceGetWorkflowTemplate(context.Background(), namespace, name).GetOptionsResourceVersion(getOptionsResourceVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTemplateServiceApi.WorkflowTemplateServiceGetWorkflowTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowTemplateServiceGetWorkflowTemplate`: IoArgoprojWorkflowV1alpha1WorkflowTemplate
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTemplateServiceApi.WorkflowTemplateServiceGetWorkflowTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowTemplateServiceGetWorkflowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getOptionsResourceVersion** | **string** | resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional | 

### Return type

[**IoArgoprojWorkflowV1alpha1WorkflowTemplate**](IoArgoprojWorkflowV1alpha1WorkflowTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowTemplateServiceLintWorkflowTemplate

> IoArgoprojWorkflowV1alpha1WorkflowTemplate WorkflowTemplateServiceLintWorkflowTemplate(ctx, namespace).Body(body).Execute()



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
    namespace := "namespace_example" // string | 
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowTemplateLintRequest() // IoArgoprojWorkflowV1alpha1WorkflowTemplateLintRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowTemplateServiceApi.WorkflowTemplateServiceLintWorkflowTemplate(context.Background(), namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTemplateServiceApi.WorkflowTemplateServiceLintWorkflowTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowTemplateServiceLintWorkflowTemplate`: IoArgoprojWorkflowV1alpha1WorkflowTemplate
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTemplateServiceApi.WorkflowTemplateServiceLintWorkflowTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowTemplateServiceLintWorkflowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowTemplateLintRequest**](IoArgoprojWorkflowV1alpha1WorkflowTemplateLintRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1WorkflowTemplate**](IoArgoprojWorkflowV1alpha1WorkflowTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowTemplateServiceListWorkflowTemplates

> IoArgoprojWorkflowV1alpha1WorkflowTemplateList WorkflowTemplateServiceListWorkflowTemplates(ctx, namespace).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Execute()



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
    namespace := "namespace_example" // string | 
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
    resp, r, err := api_client.WorkflowTemplateServiceApi.WorkflowTemplateServiceListWorkflowTemplates(context.Background(), namespace).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTemplateServiceApi.WorkflowTemplateServiceListWorkflowTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowTemplateServiceListWorkflowTemplates`: IoArgoprojWorkflowV1alpha1WorkflowTemplateList
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTemplateServiceApi.WorkflowTemplateServiceListWorkflowTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowTemplateServiceListWorkflowTemplatesRequest struct via the builder pattern


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

[**IoArgoprojWorkflowV1alpha1WorkflowTemplateList**](IoArgoprojWorkflowV1alpha1WorkflowTemplateList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowTemplateServiceUpdateWorkflowTemplate

> IoArgoprojWorkflowV1alpha1WorkflowTemplate WorkflowTemplateServiceUpdateWorkflowTemplate(ctx, namespace, name).Body(body).Execute()



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
    namespace := "namespace_example" // string | 
    name := "name_example" // string | DEPRECATED: This field is ignored.
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest() // IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowTemplateServiceApi.WorkflowTemplateServiceUpdateWorkflowTemplate(context.Background(), namespace, name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTemplateServiceApi.WorkflowTemplateServiceUpdateWorkflowTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowTemplateServiceUpdateWorkflowTemplate`: IoArgoprojWorkflowV1alpha1WorkflowTemplate
    fmt.Fprintf(os.Stdout, "Response from `WorkflowTemplateServiceApi.WorkflowTemplateServiceUpdateWorkflowTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** | DEPRECATED: This field is ignored. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowTemplateServiceUpdateWorkflowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest**](IoArgoprojWorkflowV1alpha1WorkflowTemplateUpdateRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1WorkflowTemplate**](IoArgoprojWorkflowV1alpha1WorkflowTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

