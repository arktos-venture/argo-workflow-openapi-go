# \ArtifactServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArtifactServiceGetInputArtifact**](ArtifactServiceApi.md#ArtifactServiceGetInputArtifact) | **Get** /input-artifacts/{namespace}/{name}/{podName}/{artifactName} | Get an input artifact.
[**ArtifactServiceGetInputArtifactByUID**](ArtifactServiceApi.md#ArtifactServiceGetInputArtifactByUID) | **Get** /input-artifacts-by-uid/{uid}/{podName}/{artifactName} | Get an input artifact by UID.
[**ArtifactServiceGetOutputArtifact**](ArtifactServiceApi.md#ArtifactServiceGetOutputArtifact) | **Get** /artifacts/{namespace}/{name}/{podName}/{artifactName} | Get an output artifact.
[**ArtifactServiceGetOutputArtifactByUID**](ArtifactServiceApi.md#ArtifactServiceGetOutputArtifactByUID) | **Get** /artifacts-by-uid/{uid}/{podName}/{artifactName} | Get an output artifact by UID.



## ArtifactServiceGetInputArtifact

> ArtifactServiceGetInputArtifact(ctx, namespace, name, podName, artifactName).Execute()

Get an input artifact.

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
    artifactName := "artifactName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactServiceApi.ArtifactServiceGetInputArtifact(context.Background(), namespace, name, podName, artifactName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactServiceApi.ArtifactServiceGetInputArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 
**podName** | **string** |  | 
**artifactName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactServiceGetInputArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArtifactServiceGetInputArtifactByUID

> ArtifactServiceGetInputArtifactByUID(ctx, namespace, name, podName, artifactName).Execute()

Get an input artifact by UID.

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
    artifactName := "artifactName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactServiceApi.ArtifactServiceGetInputArtifactByUID(context.Background(), namespace, name, podName, artifactName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactServiceApi.ArtifactServiceGetInputArtifactByUID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 
**podName** | **string** |  | 
**artifactName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactServiceGetInputArtifactByUIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArtifactServiceGetOutputArtifact

> ArtifactServiceGetOutputArtifact(ctx, namespace, name, podName, artifactName).Execute()

Get an output artifact.

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
    artifactName := "artifactName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactServiceApi.ArtifactServiceGetOutputArtifact(context.Background(), namespace, name, podName, artifactName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactServiceApi.ArtifactServiceGetOutputArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**name** | **string** |  | 
**podName** | **string** |  | 
**artifactName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactServiceGetOutputArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArtifactServiceGetOutputArtifactByUID

> ArtifactServiceGetOutputArtifactByUID(ctx, uid, podName, artifactName).Execute()

Get an output artifact by UID.

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
    podName := "podName_example" // string | 
    artifactName := "artifactName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArtifactServiceApi.ArtifactServiceGetOutputArtifactByUID(context.Background(), uid, podName, artifactName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactServiceApi.ArtifactServiceGetOutputArtifactByUID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 
**podName** | **string** |  | 
**artifactName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactServiceGetOutputArtifactByUIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

