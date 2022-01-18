# StreamResultOfIoK8sApiCoreV1Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**GrpcGatewayRuntimeStreamError**](GrpcGatewayRuntimeStreamError.md) |  | [optional] 
**Result** | Pointer to [**IoK8sApiCoreV1Event**](IoK8sApiCoreV1Event.md) |  | [optional] 

## Methods

### NewStreamResultOfIoK8sApiCoreV1Event

`func NewStreamResultOfIoK8sApiCoreV1Event() *StreamResultOfIoK8sApiCoreV1Event`

NewStreamResultOfIoK8sApiCoreV1Event instantiates a new StreamResultOfIoK8sApiCoreV1Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfIoK8sApiCoreV1EventWithDefaults

`func NewStreamResultOfIoK8sApiCoreV1EventWithDefaults() *StreamResultOfIoK8sApiCoreV1Event`

NewStreamResultOfIoK8sApiCoreV1EventWithDefaults instantiates a new StreamResultOfIoK8sApiCoreV1Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StreamResultOfIoK8sApiCoreV1Event) GetError() GrpcGatewayRuntimeStreamError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfIoK8sApiCoreV1Event) GetErrorOk() (*GrpcGatewayRuntimeStreamError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfIoK8sApiCoreV1Event) SetError(v GrpcGatewayRuntimeStreamError)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfIoK8sApiCoreV1Event) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResult

`func (o *StreamResultOfIoK8sApiCoreV1Event) GetResult() IoK8sApiCoreV1Event`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfIoK8sApiCoreV1Event) GetResultOk() (*IoK8sApiCoreV1Event, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfIoK8sApiCoreV1Event) SetResult(v IoK8sApiCoreV1Event)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfIoK8sApiCoreV1Event) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


