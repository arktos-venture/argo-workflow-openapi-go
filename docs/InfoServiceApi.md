# \InfoServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InfoServiceGetInfo**](InfoServiceApi.md#InfoServiceGetInfo) | **Get** /workflow/argo/api/v1/info | 
[**InfoServiceGetUserInfo**](InfoServiceApi.md#InfoServiceGetUserInfo) | **Get** /workflow/argo/api/v1/userinfo | 
[**InfoServiceGetVersion**](InfoServiceApi.md#InfoServiceGetVersion) | **Get** /workflow/argo/api/v1/version | 



## InfoServiceGetInfo

> IoArgoprojWorkflowV1alpha1InfoResponse InfoServiceGetInfo(ctx).Execute()



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
    resp, r, err := api_client.InfoServiceApi.InfoServiceGetInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoServiceApi.InfoServiceGetInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InfoServiceGetInfo`: IoArgoprojWorkflowV1alpha1InfoResponse
    fmt.Fprintf(os.Stdout, "Response from `InfoServiceApi.InfoServiceGetInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInfoServiceGetInfoRequest struct via the builder pattern


### Return type

[**IoArgoprojWorkflowV1alpha1InfoResponse**](IoArgoprojWorkflowV1alpha1InfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InfoServiceGetUserInfo

> IoArgoprojWorkflowV1alpha1GetUserInfoResponse InfoServiceGetUserInfo(ctx).Execute()



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
    resp, r, err := api_client.InfoServiceApi.InfoServiceGetUserInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoServiceApi.InfoServiceGetUserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InfoServiceGetUserInfo`: IoArgoprojWorkflowV1alpha1GetUserInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `InfoServiceApi.InfoServiceGetUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInfoServiceGetUserInfoRequest struct via the builder pattern


### Return type

[**IoArgoprojWorkflowV1alpha1GetUserInfoResponse**](IoArgoprojWorkflowV1alpha1GetUserInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InfoServiceGetVersion

> IoArgoprojWorkflowV1alpha1Version InfoServiceGetVersion(ctx).Execute()



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
    resp, r, err := api_client.InfoServiceApi.InfoServiceGetVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InfoServiceApi.InfoServiceGetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InfoServiceGetVersion`: IoArgoprojWorkflowV1alpha1Version
    fmt.Fprintf(os.Stdout, "Response from `InfoServiceApi.InfoServiceGetVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInfoServiceGetVersionRequest struct via the builder pattern


### Return type

[**IoArgoprojWorkflowV1alpha1Version**](IoArgoprojWorkflowV1alpha1Version.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

