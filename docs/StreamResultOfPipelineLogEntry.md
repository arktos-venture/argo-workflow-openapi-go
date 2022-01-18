# StreamResultOfPipelineLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**GrpcGatewayRuntimeStreamError**](GrpcGatewayRuntimeStreamError.md) |  | [optional] 
**Result** | Pointer to [**PipelineLogEntry**](PipelineLogEntry.md) |  | [optional] 

## Methods

### NewStreamResultOfPipelineLogEntry

`func NewStreamResultOfPipelineLogEntry() *StreamResultOfPipelineLogEntry`

NewStreamResultOfPipelineLogEntry instantiates a new StreamResultOfPipelineLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfPipelineLogEntryWithDefaults

`func NewStreamResultOfPipelineLogEntryWithDefaults() *StreamResultOfPipelineLogEntry`

NewStreamResultOfPipelineLogEntryWithDefaults instantiates a new StreamResultOfPipelineLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StreamResultOfPipelineLogEntry) GetError() GrpcGatewayRuntimeStreamError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfPipelineLogEntry) GetErrorOk() (*GrpcGatewayRuntimeStreamError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfPipelineLogEntry) SetError(v GrpcGatewayRuntimeStreamError)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfPipelineLogEntry) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResult

`func (o *StreamResultOfPipelineLogEntry) GetResult() PipelineLogEntry`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfPipelineLogEntry) GetResultOk() (*PipelineLogEntry, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfPipelineLogEntry) SetResult(v PipelineLogEntry)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfPipelineLogEntry) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


