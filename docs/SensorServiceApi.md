# \SensorServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SensorServiceCreateSensor**](SensorServiceApi.md#SensorServiceCreateSensor) | **Post** /workflow/argo/api/v1/sensors/{namespace} | 
[**SensorServiceDeleteSensor**](SensorServiceApi.md#SensorServiceDeleteSensor) | **Delete** /workflow/argo/api/v1/sensors/{namespace}/{name} | 
[**SensorServiceGetSensor**](SensorServiceApi.md#SensorServiceGetSensor) | **Get** /workflow/argo/api/v1/sensors/{namespace}/{name} | 
[**SensorServiceListSensors**](SensorServiceApi.md#SensorServiceListSensors) | **Get** /workflow/argo/api/v1/sensors/{namespace} | 
[**SensorServiceSensorsLogs**](SensorServiceApi.md#SensorServiceSensorsLogs) | **Get** /workflow/argo/api/v1/stream/sensors/{namespace}/logs | 
[**SensorServiceUpdateSensor**](SensorServiceApi.md#SensorServiceUpdateSensor) | **Put** /workflow/argo/api/v1/sensors/{namespace}/{name} | 
[**SensorServiceWatchSensors**](SensorServiceApi.md#SensorServiceWatchSensors) | **Get** /workflow/argo/api/v1/stream/sensors/{namespace} | 



## SensorServiceCreateSensor

> IoArgoprojEventsV1alpha1Sensor SensorServiceCreateSensor(ctx, namespace).Body(body).Execute()



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
    body := *openapiclient.NewSensorCreateSensorRequest() // SensorCreateSensorRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensorServiceApi.SensorServiceCreateSensor(context.Background(), namespace).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensorServiceApi.SensorServiceCreateSensor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensorServiceCreateSensor`: IoArgoprojEventsV1alpha1Sensor
    fmt.Fprintf(os.Stdout, "Response from `SensorServiceApi.SensorServiceCreateSensor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensorServiceCreateSensorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SensorCreateSensorRequest**](SensorCreateSensorRequest.md) |  | 

### Return type

[**IoArgoprojEventsV1alpha1Sensor**](IoArgoprojEventsV1alpha1Sensor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensorServiceDeleteSensor

> map[string]interface{} SensorServiceDeleteSensor(ctx, namespace, name).DeleteOptionsGracePeriodSeconds(deleteOptionsGracePeriodSeconds).DeleteOptionsPreconditionsUid(deleteOptionsPreconditionsUid).DeleteOptionsPreconditionsResourceVersion(deleteOptionsPreconditionsResourceVersion).DeleteOptionsOrphanDependents(deleteOptionsOrphanDependents).DeleteOptionsPropagationPolicy(deleteOptionsPropagationPolicy).DeleteOptionsDryRun(deleteOptionsDryRun).Execute()



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
    resp, r, err := api_client.SensorServiceApi.SensorServiceDeleteSensor(context.Background(), namespace, name).DeleteOptionsGracePeriodSeconds(deleteOptionsGracePeriodSeconds).DeleteOptionsPreconditionsUid(deleteOptionsPreconditionsUid).DeleteOptionsPreconditionsResourceVersion(deleteOptionsPreconditionsResourceVersion).DeleteOptionsOrphanDependents(deleteOptionsOrphanDependents).DeleteOptionsPropagationPolicy(deleteOptionsPropagationPolicy).DeleteOptionsDryRun(deleteOptionsDryRun).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensorServiceApi.SensorServiceDeleteSensor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensorServiceDeleteSensor`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SensorServiceApi.SensorServiceDeleteSensor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensorServiceDeleteSensorRequest struct via the builder pattern


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


## SensorServiceGetSensor

> IoArgoprojEventsV1alpha1Sensor SensorServiceGetSensor(ctx, namespace, name).GetOptionsResourceVersion(getOptionsResourceVersion).Execute()



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
    resp, r, err := api_client.SensorServiceApi.SensorServiceGetSensor(context.Background(), namespace, name).GetOptionsResourceVersion(getOptionsResourceVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensorServiceApi.SensorServiceGetSensor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensorServiceGetSensor`: IoArgoprojEventsV1alpha1Sensor
    fmt.Fprintf(os.Stdout, "Response from `SensorServiceApi.SensorServiceGetSensor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensorServiceGetSensorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getOptionsResourceVersion** | **string** | resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.  Defaults to unset +optional | 

### Return type

[**IoArgoprojEventsV1alpha1Sensor**](IoArgoprojEventsV1alpha1Sensor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensorServiceListSensors

> IoArgoprojEventsV1alpha1SensorList SensorServiceListSensors(ctx, namespace).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Execute()



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
    resp, r, err := api_client.SensorServiceApi.SensorServiceListSensors(context.Background(), namespace).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensorServiceApi.SensorServiceListSensors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensorServiceListSensors`: IoArgoprojEventsV1alpha1SensorList
    fmt.Fprintf(os.Stdout, "Response from `SensorServiceApi.SensorServiceListSensors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensorServiceListSensorsRequest struct via the builder pattern


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

[**IoArgoprojEventsV1alpha1SensorList**](IoArgoprojEventsV1alpha1SensorList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensorServiceSensorsLogs

> StreamResultOfSensorLogEntry SensorServiceSensorsLogs(ctx, namespace).Name(name).TriggerName(triggerName).Grep(grep).PodLogOptionsContainer(podLogOptionsContainer).PodLogOptionsFollow(podLogOptionsFollow).PodLogOptionsPrevious(podLogOptionsPrevious).PodLogOptionsSinceSeconds(podLogOptionsSinceSeconds).PodLogOptionsSinceTimeSeconds(podLogOptionsSinceTimeSeconds).PodLogOptionsSinceTimeNanos(podLogOptionsSinceTimeNanos).PodLogOptionsTimestamps(podLogOptionsTimestamps).PodLogOptionsTailLines(podLogOptionsTailLines).PodLogOptionsLimitBytes(podLogOptionsLimitBytes).PodLogOptionsInsecureSkipTLSVerifyBackend(podLogOptionsInsecureSkipTLSVerifyBackend).Execute()



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
    name := "name_example" // string | optional - only return entries for this sensor name. (optional)
    triggerName := "triggerName_example" // string | optional - only return entries for this trigger. (optional)
    grep := "grep_example" // string | option - only return entries where `msg` contains this regular expressions. (optional)
    podLogOptionsContainer := "podLogOptionsContainer_example" // string | The container for which to stream logs. Defaults to only container if there is one container in the pod. +optional. (optional)
    podLogOptionsFollow := true // bool | Follow the log stream of the pod. Defaults to false. +optional. (optional)
    podLogOptionsPrevious := true // bool | Return previous terminated container logs. Defaults to false. +optional. (optional)
    podLogOptionsSinceSeconds := "podLogOptionsSinceSeconds_example" // string | A relative time in seconds before the current time from which to show logs. If this value precedes the time a pod was started, only logs since the pod start will be returned. If this value is in the future, no logs will be returned. Only one of sinceSeconds or sinceTime may be specified. +optional. (optional)
    podLogOptionsSinceTimeSeconds := "podLogOptionsSinceTimeSeconds_example" // string | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. (optional)
    podLogOptionsSinceTimeNanos := int32(56) // int32 | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. (optional)
    podLogOptionsTimestamps := true // bool | If true, add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. Defaults to false. +optional. (optional)
    podLogOptionsTailLines := "podLogOptionsTailLines_example" // string | If set, the number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or sinceSeconds or sinceTime +optional. (optional)
    podLogOptionsLimitBytes := "podLogOptionsLimitBytes_example" // string | If set, the number of bytes to read from the server before terminating the log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit. +optional. (optional)
    podLogOptionsInsecureSkipTLSVerifyBackend := true // bool | insecureSkipTLSVerifyBackend indicates that the apiserver should not confirm the validity of the serving certificate of the backend it is connecting to.  This will make the HTTPS connection between the apiserver and the backend insecure. This means the apiserver cannot verify the log data it is receiving came from the real kubelet.  If the kubelet is configured to verify the apiserver's TLS credentials, it does not mean the connection to the real kubelet is vulnerable to a man in the middle attack (e.g. an attacker could not intercept the actual log data coming from the real kubelet). +optional. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensorServiceApi.SensorServiceSensorsLogs(context.Background(), namespace).Name(name).TriggerName(triggerName).Grep(grep).PodLogOptionsContainer(podLogOptionsContainer).PodLogOptionsFollow(podLogOptionsFollow).PodLogOptionsPrevious(podLogOptionsPrevious).PodLogOptionsSinceSeconds(podLogOptionsSinceSeconds).PodLogOptionsSinceTimeSeconds(podLogOptionsSinceTimeSeconds).PodLogOptionsSinceTimeNanos(podLogOptionsSinceTimeNanos).PodLogOptionsTimestamps(podLogOptionsTimestamps).PodLogOptionsTailLines(podLogOptionsTailLines).PodLogOptionsLimitBytes(podLogOptionsLimitBytes).PodLogOptionsInsecureSkipTLSVerifyBackend(podLogOptionsInsecureSkipTLSVerifyBackend).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensorServiceApi.SensorServiceSensorsLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensorServiceSensorsLogs`: StreamResultOfSensorLogEntry
    fmt.Fprintf(os.Stdout, "Response from `SensorServiceApi.SensorServiceSensorsLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensorServiceSensorsLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | optional - only return entries for this sensor name. | 
 **triggerName** | **string** | optional - only return entries for this trigger. | 
 **grep** | **string** | option - only return entries where &#x60;msg&#x60; contains this regular expressions. | 
 **podLogOptionsContainer** | **string** | The container for which to stream logs. Defaults to only container if there is one container in the pod. +optional. | 
 **podLogOptionsFollow** | **bool** | Follow the log stream of the pod. Defaults to false. +optional. | 
 **podLogOptionsPrevious** | **bool** | Return previous terminated container logs. Defaults to false. +optional. | 
 **podLogOptionsSinceSeconds** | **string** | A relative time in seconds before the current time from which to show logs. If this value precedes the time a pod was started, only logs since the pod start will be returned. If this value is in the future, no logs will be returned. Only one of sinceSeconds or sinceTime may be specified. +optional. | 
 **podLogOptionsSinceTimeSeconds** | **string** | Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. | 
 **podLogOptionsSinceTimeNanos** | **int32** | Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. | 
 **podLogOptionsTimestamps** | **bool** | If true, add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. Defaults to false. +optional. | 
 **podLogOptionsTailLines** | **string** | If set, the number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or sinceSeconds or sinceTime +optional. | 
 **podLogOptionsLimitBytes** | **string** | If set, the number of bytes to read from the server before terminating the log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit. +optional. | 
 **podLogOptionsInsecureSkipTLSVerifyBackend** | **bool** | insecureSkipTLSVerifyBackend indicates that the apiserver should not confirm the validity of the serving certificate of the backend it is connecting to.  This will make the HTTPS connection between the apiserver and the backend insecure. This means the apiserver cannot verify the log data it is receiving came from the real kubelet.  If the kubelet is configured to verify the apiserver&#39;s TLS credentials, it does not mean the connection to the real kubelet is vulnerable to a man in the middle attack (e.g. an attacker could not intercept the actual log data coming from the real kubelet). +optional. | 

### Return type

[**StreamResultOfSensorLogEntry**](StreamResultOfSensorLogEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensorServiceUpdateSensor

> IoArgoprojEventsV1alpha1Sensor SensorServiceUpdateSensor(ctx, namespace, name).Body(body).Execute()



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
    body := *openapiclient.NewSensorUpdateSensorRequest() // SensorUpdateSensorRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SensorServiceApi.SensorServiceUpdateSensor(context.Background(), namespace, name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensorServiceApi.SensorServiceUpdateSensor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensorServiceUpdateSensor`: IoArgoprojEventsV1alpha1Sensor
    fmt.Fprintf(os.Stdout, "Response from `SensorServiceApi.SensorServiceUpdateSensor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensorServiceUpdateSensorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**SensorUpdateSensorRequest**](SensorUpdateSensorRequest.md) |  | 

### Return type

[**IoArgoprojEventsV1alpha1Sensor**](IoArgoprojEventsV1alpha1Sensor.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SensorServiceWatchSensors

> StreamResultOfSensorSensorWatchEvent SensorServiceWatchSensors(ctx, namespace).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Execute()



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
    resp, r, err := api_client.SensorServiceApi.SensorServiceWatchSensors(context.Background(), namespace).ListOptionsLabelSelector(listOptionsLabelSelector).ListOptionsFieldSelector(listOptionsFieldSelector).ListOptionsWatch(listOptionsWatch).ListOptionsAllowWatchBookmarks(listOptionsAllowWatchBookmarks).ListOptionsResourceVersion(listOptionsResourceVersion).ListOptionsResourceVersionMatch(listOptionsResourceVersionMatch).ListOptionsTimeoutSeconds(listOptionsTimeoutSeconds).ListOptionsLimit(listOptionsLimit).ListOptionsContinue(listOptionsContinue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SensorServiceApi.SensorServiceWatchSensors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SensorServiceWatchSensors`: StreamResultOfSensorSensorWatchEvent
    fmt.Fprintf(os.Stdout, "Response from `SensorServiceApi.SensorServiceWatchSensors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSensorServiceWatchSensorsRequest struct via the builder pattern


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

[**StreamResultOfSensorSensorWatchEvent**](StreamResultOfSensorSensorWatchEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

