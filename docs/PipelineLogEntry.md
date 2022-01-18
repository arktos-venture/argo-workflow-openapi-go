# PipelineLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**PipelineName** | Pointer to **string** |  | [optional] 
**StepName** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 

## Methods

### NewPipelineLogEntry

`func NewPipelineLogEntry() *PipelineLogEntry`

NewPipelineLogEntry instantiates a new PipelineLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineLogEntryWithDefaults

`func NewPipelineLogEntryWithDefaults() *PipelineLogEntry`

NewPipelineLogEntryWithDefaults instantiates a new PipelineLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *PipelineLogEntry) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *PipelineLogEntry) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *PipelineLogEntry) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *PipelineLogEntry) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetNamespace

`func (o *PipelineLogEntry) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PipelineLogEntry) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PipelineLogEntry) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PipelineLogEntry) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPipelineName

`func (o *PipelineLogEntry) GetPipelineName() string`

GetPipelineName returns the PipelineName field if non-nil, zero value otherwise.

### GetPipelineNameOk

`func (o *PipelineLogEntry) GetPipelineNameOk() (*string, bool)`

GetPipelineNameOk returns a tuple with the PipelineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineName

`func (o *PipelineLogEntry) SetPipelineName(v string)`

SetPipelineName sets PipelineName field to given value.

### HasPipelineName

`func (o *PipelineLogEntry) HasPipelineName() bool`

HasPipelineName returns a boolean if a field has been set.

### GetStepName

`func (o *PipelineLogEntry) GetStepName() string`

GetStepName returns the StepName field if non-nil, zero value otherwise.

### GetStepNameOk

`func (o *PipelineLogEntry) GetStepNameOk() (*string, bool)`

GetStepNameOk returns a tuple with the StepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepName

`func (o *PipelineLogEntry) SetStepName(v string)`

SetStepName sets StepName field to given value.

### HasStepName

`func (o *PipelineLogEntry) HasStepName() bool`

HasStepName returns a boolean if a field has been set.

### GetTime

`func (o *PipelineLogEntry) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PipelineLogEntry) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PipelineLogEntry) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *PipelineLogEntry) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


