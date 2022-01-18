# StreamResultOfSensorLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**GrpcGatewayRuntimeStreamError**](GrpcGatewayRuntimeStreamError.md) |  | [optional] 
**Result** | Pointer to [**SensorLogEntry**](SensorLogEntry.md) |  | [optional] 

## Methods

### NewStreamResultOfSensorLogEntry

`func NewStreamResultOfSensorLogEntry() *StreamResultOfSensorLogEntry`

NewStreamResultOfSensorLogEntry instantiates a new StreamResultOfSensorLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfSensorLogEntryWithDefaults

`func NewStreamResultOfSensorLogEntryWithDefaults() *StreamResultOfSensorLogEntry`

NewStreamResultOfSensorLogEntryWithDefaults instantiates a new StreamResultOfSensorLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StreamResultOfSensorLogEntry) GetError() GrpcGatewayRuntimeStreamError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfSensorLogEntry) GetErrorOk() (*GrpcGatewayRuntimeStreamError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfSensorLogEntry) SetError(v GrpcGatewayRuntimeStreamError)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfSensorLogEntry) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResult

`func (o *StreamResultOfSensorLogEntry) GetResult() SensorLogEntry`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfSensorLogEntry) GetResultOk() (*SensorLogEntry, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfSensorLogEntry) SetResult(v SensorLogEntry)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfSensorLogEntry) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


