# StreamResultOfEventsourceLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**GrpcGatewayRuntimeStreamError**](GrpcGatewayRuntimeStreamError.md) |  | [optional] 
**Result** | Pointer to [**EventsourceLogEntry**](EventsourceLogEntry.md) |  | [optional] 

## Methods

### NewStreamResultOfEventsourceLogEntry

`func NewStreamResultOfEventsourceLogEntry() *StreamResultOfEventsourceLogEntry`

NewStreamResultOfEventsourceLogEntry instantiates a new StreamResultOfEventsourceLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamResultOfEventsourceLogEntryWithDefaults

`func NewStreamResultOfEventsourceLogEntryWithDefaults() *StreamResultOfEventsourceLogEntry`

NewStreamResultOfEventsourceLogEntryWithDefaults instantiates a new StreamResultOfEventsourceLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StreamResultOfEventsourceLogEntry) GetError() GrpcGatewayRuntimeStreamError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamResultOfEventsourceLogEntry) GetErrorOk() (*GrpcGatewayRuntimeStreamError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamResultOfEventsourceLogEntry) SetError(v GrpcGatewayRuntimeStreamError)`

SetError sets Error field to given value.

### HasError

`func (o *StreamResultOfEventsourceLogEntry) HasError() bool`

HasError returns a boolean if a field has been set.

### GetResult

`func (o *StreamResultOfEventsourceLogEntry) GetResult() EventsourceLogEntry`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *StreamResultOfEventsourceLogEntry) GetResultOk() (*EventsourceLogEntry, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *StreamResultOfEventsourceLogEntry) SetResult(v EventsourceLogEntry)`

SetResult sets Result field to given value.

### HasResult

`func (o *StreamResultOfEventsourceLogEntry) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


