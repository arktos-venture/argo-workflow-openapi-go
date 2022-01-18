# \WorkflowServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowServiceCreateWorkflow**](WorkflowServiceApi.md#WorkflowServiceCreateWorkflow) | **Post** /workflow/argo/api/v1/workflows/{namespace} | 
[**WorkflowServiceDeleteWorkflow**](WorkflowServiceApi.md#WorkflowServiceDeleteWorkflow) | **Delete** /workflow/argo/api/v1/workflows/{namespace}/{name} | 
[**WorkflowServiceGetWorkflow**](WorkflowServiceApi.md#WorkflowServiceGetWorkflow) | **Get** /workflow/argo/api/v1/workflows/{namespace}/{name} | 
[**WorkflowServiceLintWorkflow**](WorkflowServiceApi.md#WorkflowServiceLintWorkflow) | **Post** /workflow/argo/api/v1/workflows/{namespace}/lint | 
[**WorkflowServiceListWorkflows**](WorkflowServiceApi.md#WorkflowServiceListWorkflows) | **Get** /workflow/argo/api/v1/workflows/{namespace} | 
[**WorkflowServicePodLogs**](WorkflowServiceApi.md#WorkflowServicePodLogs) | **Get** /workflow/argo/api/v1/workflows/{namespace}/{name}/{podName}/log | DEPRECATED: Cannot work via HTTP if podName is an empty string. Use WorkflowLogs.
[**WorkflowServiceResubmitWorkflow**](WorkflowServiceApi.md#WorkflowServiceResubmitWorkflow) | **Put** /workflow/argo/api/v1/workflows/{namespace}/{name}/resubmit | 
[**WorkflowServiceResumeWorkflow**](WorkflowServiceApi.md#WorkflowServiceResumeWorkflow) | **Put** /workflow/argo/api/v1/workflows/{namespace}/{name}/resume | 
[**WorkflowServiceRetryWorkflow**](WorkflowServiceApi.md#WorkflowServiceRetryWorkflow) | **Put** /workflow/argo/api/v1/workflows/{namespace}/{name}/retry | 
[**WorkflowServiceSetWorkflow**](WorkflowServiceApi.md#WorkflowServiceSetWorkflow) | **Put** /workflow/argo/api/v1/workflows/{namespace}/{name}/set | 
[**WorkflowServiceStopWorkflow**](WorkflowServiceApi.md#WorkflowServiceStopWorkflow) | **Put** /workflow/argo/api/v1/workflows/{namespace}/{name}/stop | 
[**WorkflowServiceSubmitWorkflow**](WorkflowServiceApi.md#WorkflowServiceSubmitWorkflow) | **Post** /workflow/argo/api/v1/workflows/{namespace}/submit | 
[**WorkflowServiceSuspendWorkflow**](WorkflowServiceApi.md#WorkflowServiceSuspendWorkflow) | **Put** /workflow/argo/api/v1/workflows/{namespace}/{name}/suspend | 
[**WorkflowServiceTerminateWorkflow**](WorkflowServiceApi.md#WorkflowServiceTerminateWorkflow) | **Put** /workflow/argo/api/v1/workflows/{namespace}/{name}/terminate | 
[**WorkflowServiceWatchEvents**](WorkflowServiceApi.md#WorkflowServiceWatchEvents) | **Get** /workflow/argo/api/v1/stream/events/{namespace} | 
[**WorkflowServiceWatchWorkflows**](WorkflowServiceApi.md#WorkflowServiceWatchWorkflows) | **Get** /workflow/argo/api/v1/workflow-events/{namespace} | 
[**WorkflowServiceWorkflowLogs**](WorkflowServiceApi.md#WorkflowServiceWorkflowLogs) | **Get** /workflow/argo/api/v1/workflows/{namespace}/{name}/log | 



## WorkflowServiceCreateWorkflow

> IoArgoprojWorkflowV1alpha1Workflow WorkflowServiceCreateWorkflow(ctx, namespace).Body(body).Execute()



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
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowCreateRequest() // IoArgoprojWorkflowV1alpha1WorkflowCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceCreateWorkflow(context.Background(), namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceCreateWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceCreateWorkflow`: IoArgoprojWorkflowV1alpha1Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceCreateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceCreateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowCreateRequest**](IoArgoprojWorkflowV1alpha1WorkflowCreateRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1Workflow**](IoArgoprojWorkflowV1alpha1Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceDeleteWorkflow

> map[string]interface{} WorkflowServiceDeleteWorkflow(ctx, namespace, name).DeleteOptionsGracePeriodSeconds(deleteOptionsGracePeriodSeconds).DeleteOptionsPreconditionsUid(deleteOptionsPreconditionsUid).DeleteOptionsPreconditionsResourceVersion(deleteOptionsPreconditionsResourceVersion).DeleteOptionsOrphanDependents(deleteOptionsOrphanDependents).DeleteOptionsPropagationPolicy(deleteOptionsPropagationPolicy).DeleteOptionsDryRun(deleteOptionsDryRun).Execute()



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
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceDeleteWorkflow(context.Background(), namespace, name).DeleteOptionsGracePeriodSeconds(deleteOptionsGracePeriodSeconds).DeleteOptionsPreconditionsUid(deleteOptionsPreconditionsUid).DeleteOptionsPreconditionsResourceVersion(deleteOptionsPreconditionsResourceVersion).DeleteOptionsOrphanDependents(deleteOptionsOrphanDependents).DeleteOptionsPropagationPolicy(deleteOptionsPropagationPolicy).DeleteOptionsDryRun(deleteOptionsDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceDeleteWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceDeleteWorkflow`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceDeleteWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceDeleteWorkflowRequest struct via the builder pattern


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


## WorkflowServiceGetWorkflow

> IoArgoprojWorkflowV1alpha1Workflow WorkflowServiceGetWorkflow(ctx, namespace, name).GetOptionsResourceVersion(getOptionsResourceVersion).Fields(fields).Execute()



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
    fields := "fields_example" // string | Fields to be included or excluded in the response. e.g. \"spec,status.phase\", \"-status.nodes\". (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceGetWorkflow(context.Background(), namespace, name).GetOptionsResourceVersion(getOptionsResourceVersion).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceGetWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceGetWorkflow`: IoArgoprojWorkflowV1alpha1Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceGetWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceGetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getOptionsResourceVersion** | **string** | resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional | 
 **fields** | **string** | Fields to be included or excluded in the response. e.g. \&quot;spec,status.phase\&quot;, \&quot;-status.nodes\&quot;. | 

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


## WorkflowServiceLintWorkflow

> IoArgoprojWorkflowV1alpha1Workflow WorkflowServiceLintWorkflow(ctx, namespace).Body(body).Execute()



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
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowLintRequest() // IoArgoprojWorkflowV1alpha1WorkflowLintRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceLintWorkflow(context.Background(), namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceLintWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceLintWorkflow`: IoArgoprojWorkflowV1alpha1Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceLintWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceLintWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowLintRequest**](IoArgoprojWorkflowV1alpha1WorkflowLintRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1Workflow**](IoArgoprojWorkflowV1alpha1Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceListWorkflows

> IoArgoprojWorkflowV1alpha1WorkflowList WorkflowServiceListWorkflows(ctx, namespace).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Fields(fields).Execute()



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
    fields := "fields_example" // string | Fields to be included or excluded in the response. e.g. \"items.spec,items.status.phase\", \"-items.status.nodes\". (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceListWorkflows(context.Background(), namespace).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceListWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceListWorkflows`: IoArgoprojWorkflowV1alpha1WorkflowList
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceListWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceListWorkflowsRequest struct via the builder pattern


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
 **fields** | **string** | Fields to be included or excluded in the response. e.g. \&quot;items.spec,items.status.phase\&quot;, \&quot;-items.status.nodes\&quot;. | 

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


## WorkflowServicePodLogs

> StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry WorkflowServicePodLogs(ctx, namespace, name, podName).LogOptionsContainer(logOptionsContainer).LogOptionsFollow(logOptionsFollow).LogOptionsPrevious(logOptionsPrevious).LogOptionsSinceSeconds(logOptionsSinceSeconds).LogOptionsSinceTimeSeconds(logOptionsSinceTimeSeconds).LogOptionsSinceTimeNanos(logOptionsSinceTimeNanos).LogOptionsTimestamps(logOptionsTimestamps).LogOptionsTailLines(logOptionsTailLines).LogOptionsLimitBytes(logOptionsLimitBytes).LogOptionsInsecureSkipTLSVerifyBackend(logOptionsInsecureSkipTLSVerifyBackend).Grep(grep).Execute()

DEPRECATED: Cannot work via HTTP if podName is an empty string. Use WorkflowLogs.

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
    podName := "podName_example" // string | 
    logOptionsContainer := "logOptionsContainer_example" // string | The container for which to stream logs. Defaults to only container if there is one container in the pod. +optional. (optional)
    logOptionsFollow := true // bool | Follow the log stream of the pod. Defaults to false. +optional. (optional)
    logOptionsPrevious := true // bool | Return previous terminated container logs. Defaults to false. +optional. (optional)
    logOptionsSinceSeconds := "logOptionsSinceSeconds_example" // string | A relative time in seconds before the current time from which to show logs. If this value precedes the time a pod was started, only logs since the pod start will be returned. If this value is in the future, no logs will be returned. Only one of sinceSeconds or sinceTime may be specified. +optional. (optional)
    logOptionsSinceTimeSeconds := "logOptionsSinceTimeSeconds_example" // string | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. (optional)
    logOptionsSinceTimeNanos := int32(56) // int32 | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. (optional)
    logOptionsTimestamps := true // bool | If true, add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. Defaults to false. +optional. (optional)
    logOptionsTailLines := "logOptionsTailLines_example" // string | If set, the number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or sinceSeconds or sinceTime +optional. (optional)
    logOptionsLimitBytes := "logOptionsLimitBytes_example" // string | If set, the number of bytes to read from the server before terminating the log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit. +optional. (optional)
    logOptionsInsecureSkipTLSVerifyBackend := true // bool | insecureSkipTLSVerifyBackend indicates that the apiserver should not confirm the validity of the serving certificate of the backend it is connecting to.  This will make the HTTPS connection between the apiserver and the backend insecure. This means the apiserver cannot verify the log data it is receiving came from the real kubelet.  If the kubelet is configured to verify the apiserver's TLS credentials, it does not mean the connection to the real kubelet is vulnerable to a man in the middle attack (e.g. an attacker could not intercept the actual log data coming from the real kubelet). +optional. (optional)
    grep := "grep_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServicePodLogs(context.Background(), namespace, name, podName).LogOptionsContainer(logOptionsContainer).LogOptionsFollow(logOptionsFollow).LogOptionsPrevious(logOptionsPrevious).LogOptionsSinceSeconds(logOptionsSinceSeconds).LogOptionsSinceTimeSeconds(logOptionsSinceTimeSeconds).LogOptionsSinceTimeNanos(logOptionsSinceTimeNanos).LogOptionsTimestamps(logOptionsTimestamps).LogOptionsTailLines(logOptionsTailLines).LogOptionsLimitBytes(logOptionsLimitBytes).LogOptionsInsecureSkipTLSVerifyBackend(logOptionsInsecureSkipTLSVerifyBackend).Grep(grep).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServicePodLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServicePodLogs`: StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServicePodLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 
**podName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServicePodLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **logOptionsContainer** | **string** | The container for which to stream logs. Defaults to only container if there is one container in the pod. +optional. | 
 **logOptionsFollow** | **bool** | Follow the log stream of the pod. Defaults to false. +optional. | 
 **logOptionsPrevious** | **bool** | Return previous terminated container logs. Defaults to false. +optional. | 
 **logOptionsSinceSeconds** | **string** | A relative time in seconds before the current time from which to show logs. If this value precedes the time a pod was started, only logs since the pod start will be returned. If this value is in the future, no logs will be returned. Only one of sinceSeconds or sinceTime may be specified. +optional. | 
 **logOptionsSinceTimeSeconds** | **string** | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. | 
 **logOptionsSinceTimeNanos** | **int32** | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. | 
 **logOptionsTimestamps** | **bool** | If true, add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. Defaults to false. +optional. | 
 **logOptionsTailLines** | **string** | If set, the number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or sinceSeconds or sinceTime +optional. | 
 **logOptionsLimitBytes** | **string** | If set, the number of bytes to read from the server before terminating the log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit. +optional. | 
 **logOptionsInsecureSkipTLSVerifyBackend** | **bool** | insecureSkipTLSVerifyBackend indicates that the apiserver should not confirm the validity of the serving certificate of the backend it is connecting to.  This will make the HTTPS connection between the apiserver and the backend insecure. This means the apiserver cannot verify the log data it is receiving came from the real kubelet.  If the kubelet is configured to verify the apiserver&#39;s TLS credentials, it does not mean the connection to the real kubelet is vulnerable to a man in the middle attack (e.g. an attacker could not intercept the actual log data coming from the real kubelet). +optional. | 
 **grep** | **string** |  | 

### Return type

[**StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry**](StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceResubmitWorkflow

> IoArgoprojWorkflowV1alpha1Workflow WorkflowServiceResubmitWorkflow(ctx, namespace, name).Body(body).Execute()



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
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowResubmitRequest() // IoArgoprojWorkflowV1alpha1WorkflowResubmitRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceResubmitWorkflow(context.Background(), namespace, name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceResubmitWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceResubmitWorkflow`: IoArgoprojWorkflowV1alpha1Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceResubmitWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceResubmitWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowResubmitRequest**](IoArgoprojWorkflowV1alpha1WorkflowResubmitRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1Workflow**](IoArgoprojWorkflowV1alpha1Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceResumeWorkflow

> IoArgoprojWorkflowV1alpha1Workflow WorkflowServiceResumeWorkflow(ctx, namespace, name).Body(body).Execute()



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
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowResumeRequest() // IoArgoprojWorkflowV1alpha1WorkflowResumeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceResumeWorkflow(context.Background(), namespace, name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceResumeWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceResumeWorkflow`: IoArgoprojWorkflowV1alpha1Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceResumeWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceResumeWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowResumeRequest**](IoArgoprojWorkflowV1alpha1WorkflowResumeRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1Workflow**](IoArgoprojWorkflowV1alpha1Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceRetryWorkflow

> IoArgoprojWorkflowV1alpha1Workflow WorkflowServiceRetryWorkflow(ctx, namespace, name).Body(body).Execute()



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
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowRetryRequest() // IoArgoprojWorkflowV1alpha1WorkflowRetryRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceRetryWorkflow(context.Background(), namespace, name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceRetryWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceRetryWorkflow`: IoArgoprojWorkflowV1alpha1Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceRetryWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceRetryWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowRetryRequest**](IoArgoprojWorkflowV1alpha1WorkflowRetryRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1Workflow**](IoArgoprojWorkflowV1alpha1Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceSetWorkflow

> IoArgoprojWorkflowV1alpha1Workflow WorkflowServiceSetWorkflow(ctx, namespace, name).Body(body).Execute()



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
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowSetRequest() // IoArgoprojWorkflowV1alpha1WorkflowSetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceSetWorkflow(context.Background(), namespace, name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceSetWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceSetWorkflow`: IoArgoprojWorkflowV1alpha1Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceSetWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceSetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowSetRequest**](IoArgoprojWorkflowV1alpha1WorkflowSetRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1Workflow**](IoArgoprojWorkflowV1alpha1Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceStopWorkflow

> IoArgoprojWorkflowV1alpha1Workflow WorkflowServiceStopWorkflow(ctx, namespace, name).Body(body).Execute()



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
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowStopRequest() // IoArgoprojWorkflowV1alpha1WorkflowStopRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceStopWorkflow(context.Background(), namespace, name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceStopWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceStopWorkflow`: IoArgoprojWorkflowV1alpha1Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceStopWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceStopWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowStopRequest**](IoArgoprojWorkflowV1alpha1WorkflowStopRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1Workflow**](IoArgoprojWorkflowV1alpha1Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceSubmitWorkflow

> IoArgoprojWorkflowV1alpha1Workflow WorkflowServiceSubmitWorkflow(ctx, namespace).Body(body).Execute()



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
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowSubmitRequest() // IoArgoprojWorkflowV1alpha1WorkflowSubmitRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceSubmitWorkflow(context.Background(), namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceSubmitWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceSubmitWorkflow`: IoArgoprojWorkflowV1alpha1Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceSubmitWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceSubmitWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowSubmitRequest**](IoArgoprojWorkflowV1alpha1WorkflowSubmitRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1Workflow**](IoArgoprojWorkflowV1alpha1Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceSuspendWorkflow

> IoArgoprojWorkflowV1alpha1Workflow WorkflowServiceSuspendWorkflow(ctx, namespace, name).Body(body).Execute()



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
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowSuspendRequest() // IoArgoprojWorkflowV1alpha1WorkflowSuspendRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceSuspendWorkflow(context.Background(), namespace, name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceSuspendWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceSuspendWorkflow`: IoArgoprojWorkflowV1alpha1Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceSuspendWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceSuspendWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowSuspendRequest**](IoArgoprojWorkflowV1alpha1WorkflowSuspendRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1Workflow**](IoArgoprojWorkflowV1alpha1Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceTerminateWorkflow

> IoArgoprojWorkflowV1alpha1Workflow WorkflowServiceTerminateWorkflow(ctx, namespace, name).Body(body).Execute()



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
    body := *openapiclient.NewIoArgoprojWorkflowV1alpha1WorkflowTerminateRequest() // IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceTerminateWorkflow(context.Background(), namespace, name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceTerminateWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceTerminateWorkflow`: IoArgoprojWorkflowV1alpha1Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceTerminateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceTerminateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest**](IoArgoprojWorkflowV1alpha1WorkflowTerminateRequest.md) |  | 

### Return type

[**IoArgoprojWorkflowV1alpha1Workflow**](IoArgoprojWorkflowV1alpha1Workflow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceWatchEvents

> StreamResultOfIoK8sApiCoreV1Event WorkflowServiceWatchEvents(ctx, namespace).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Execute()



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
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceWatchEvents(context.Background(), namespace).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceWatchEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceWatchEvents`: StreamResultOfIoK8sApiCoreV1Event
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceWatchEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceWatchEventsRequest struct via the builder pattern


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

[**StreamResultOfIoK8sApiCoreV1Event**](StreamResultOfIoK8sApiCoreV1Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceWatchWorkflows

> StreamResultOfIoArgoprojWorkflowV1alpha1WorkflowWatchEvent WorkflowServiceWatchWorkflows(ctx, namespace).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Fields(fields).Execute()



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
    fields := "fields_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceWatchWorkflows(context.Background(), namespace).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceWatchWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceWatchWorkflows`: StreamResultOfIoArgoprojWorkflowV1alpha1WorkflowWatchEvent
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceWatchWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceWatchWorkflowsRequest struct via the builder pattern


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
 **fields** | **string** |  | 

### Return type

[**StreamResultOfIoArgoprojWorkflowV1alpha1WorkflowWatchEvent**](StreamResultOfIoArgoprojWorkflowV1alpha1WorkflowWatchEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowServiceWorkflowLogs

> StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry WorkflowServiceWorkflowLogs(ctx, namespace, name).PodName(podName).LogOptionsContainer(logOptionsContainer).LogOptionsFollow(logOptionsFollow).LogOptionsPrevious(logOptionsPrevious).LogOptionsSinceSeconds(logOptionsSinceSeconds).LogOptionsSinceTimeSeconds(logOptionsSinceTimeSeconds).LogOptionsSinceTimeNanos(logOptionsSinceTimeNanos).LogOptionsTimestamps(logOptionsTimestamps).LogOptionsTailLines(logOptionsTailLines).LogOptionsLimitBytes(logOptionsLimitBytes).LogOptionsInsecureSkipTLSVerifyBackend(logOptionsInsecureSkipTLSVerifyBackend).Grep(grep).Execute()



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
    podName := "podName_example" // string |  (optional)
    logOptionsContainer := "logOptionsContainer_example" // string | The container for which to stream logs. Defaults to only container if there is one container in the pod. +optional. (optional)
    logOptionsFollow := true // bool | Follow the log stream of the pod. Defaults to false. +optional. (optional)
    logOptionsPrevious := true // bool | Return previous terminated container logs. Defaults to false. +optional. (optional)
    logOptionsSinceSeconds := "logOptionsSinceSeconds_example" // string | A relative time in seconds before the current time from which to show logs. If this value precedes the time a pod was started, only logs since the pod start will be returned. If this value is in the future, no logs will be returned. Only one of sinceSeconds or sinceTime may be specified. +optional. (optional)
    logOptionsSinceTimeSeconds := "logOptionsSinceTimeSeconds_example" // string | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. (optional)
    logOptionsSinceTimeNanos := int32(56) // int32 | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. (optional)
    logOptionsTimestamps := true // bool | If true, add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. Defaults to false. +optional. (optional)
    logOptionsTailLines := "logOptionsTailLines_example" // string | If set, the number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or sinceSeconds or sinceTime +optional. (optional)
    logOptionsLimitBytes := "logOptionsLimitBytes_example" // string | If set, the number of bytes to read from the server before terminating the log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit. +optional. (optional)
    logOptionsInsecureSkipTLSVerifyBackend := true // bool | insecureSkipTLSVerifyBackend indicates that the apiserver should not confirm the validity of the serving certificate of the backend it is connecting to.  This will make the HTTPS connection between the apiserver and the backend insecure. This means the apiserver cannot verify the log data it is receiving came from the real kubelet.  If the kubelet is configured to verify the apiserver's TLS credentials, it does not mean the connection to the real kubelet is vulnerable to a man in the middle attack (e.g. an attacker could not intercept the actual log data coming from the real kubelet). +optional. (optional)
    grep := "grep_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.WorkflowServiceApi.WorkflowServiceWorkflowLogs(context.Background(), namespace, name).PodName(podName).LogOptionsContainer(logOptionsContainer).LogOptionsFollow(logOptionsFollow).LogOptionsPrevious(logOptionsPrevious).LogOptionsSinceSeconds(logOptionsSinceSeconds).LogOptionsSinceTimeSeconds(logOptionsSinceTimeSeconds).LogOptionsSinceTimeNanos(logOptionsSinceTimeNanos).LogOptionsTimestamps(logOptionsTimestamps).LogOptionsTailLines(logOptionsTailLines).LogOptionsLimitBytes(logOptionsLimitBytes).LogOptionsInsecureSkipTLSVerifyBackend(logOptionsInsecureSkipTLSVerifyBackend).Grep(grep).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowServiceApi.WorkflowServiceWorkflowLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkflowServiceWorkflowLogs`: StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry
    fmt.Fprintf(os.Stdout, "Response from `WorkflowServiceApi.WorkflowServiceWorkflowLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowServiceWorkflowLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **podName** | **string** |  | 
 **logOptionsContainer** | **string** | The container for which to stream logs. Defaults to only container if there is one container in the pod. +optional. | 
 **logOptionsFollow** | **bool** | Follow the log stream of the pod. Defaults to false. +optional. | 
 **logOptionsPrevious** | **bool** | Return previous terminated container logs. Defaults to false. +optional. | 
 **logOptionsSinceSeconds** | **string** | A relative time in seconds before the current time from which to show logs. If this value precedes the time a pod was started, only logs since the pod start will be returned. If this value is in the future, no logs will be returned. Only one of sinceSeconds or sinceTime may be specified. +optional. | 
 **logOptionsSinceTimeSeconds** | **string** | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. | 
 **logOptionsSinceTimeNanos** | **int32** | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. | 
 **logOptionsTimestamps** | **bool** | If true, add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. Defaults to false. +optional. | 
 **logOptionsTailLines** | **string** | If set, the number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or sinceSeconds or sinceTime +optional. | 
 **logOptionsLimitBytes** | **string** | If set, the number of bytes to read from the server before terminating the log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit. +optional. | 
 **logOptionsInsecureSkipTLSVerifyBackend** | **bool** | insecureSkipTLSVerifyBackend indicates that the apiserver should not confirm the validity of the serving certificate of the backend it is connecting to.  This will make the HTTPS connection between the apiserver and the backend insecure. This means the apiserver cannot verify the log data it is receiving came from the real kubelet.  If the kubelet is configured to verify the apiserver&#39;s TLS credentials, it does not mean the connection to the real kubelet is vulnerable to a man in the middle attack (e.g. an attacker could not intercept the actual log data coming from the real kubelet). +optional. | 
 **grep** | **string** |  | 

### Return type

[**StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry**](StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

