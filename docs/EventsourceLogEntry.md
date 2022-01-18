# EventsourceLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | Pointer to **string** |  | [optional] 
**EventSourceName** | Pointer to **string** |  | [optional] 
**EventSourceType** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 

## Methods

### NewEventsourceLogEntry

`func NewEventsourceLogEntry() *EventsourceLogEntry`

NewEventsourceLogEntry instantiates a new EventsourceLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsourceLogEntryWithDefaults

`func NewEventsourceLogEntryWithDefaults() *EventsourceLogEntry`

NewEventsourceLogEntryWithDefaults instantiates a new EventsourceLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *EventsourceLogEntry) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *EventsourceLogEntry) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *EventsourceLogEntry) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *EventsourceLogEntry) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetEventSourceName

`func (o *EventsourceLogEntry) GetEventSourceName() string`

GetEventSourceName returns the EventSourceName field if non-nil, zero value otherwise.

### GetEventSourceNameOk

`func (o *EventsourceLogEntry) GetEventSourceNameOk() (*string, bool)`

GetEventSourceNameOk returns a tuple with the EventSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSourceName

`func (o *EventsourceLogEntry) SetEventSourceName(v string)`

SetEventSourceName sets EventSourceName field to given value.

### HasEventSourceName

`func (o *EventsourceLogEntry) HasEventSourceName() bool`

HasEventSourceName returns a boolean if a field has been set.

### GetEventSourceType

`func (o *EventsourceLogEntry) GetEventSourceType() string`

GetEventSourceType returns the EventSourceType field if non-nil, zero value otherwise.

### GetEventSourceTypeOk

`func (o *EventsourceLogEntry) GetEventSourceTypeOk() (*string, bool)`

GetEventSourceTypeOk returns a tuple with the EventSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSourceType

`func (o *EventsourceLogEntry) SetEventSourceType(v string)`

SetEventSourceType sets EventSourceType field to given value.

### HasEventSourceType

`func (o *EventsourceLogEntry) HasEventSourceType() bool`

HasEventSourceType returns a boolean if a field has been set.

### GetLevel

`func (o *EventsourceLogEntry) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *EventsourceLogEntry) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *EventsourceLogEntry) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *EventsourceLogEntry) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMsg

`func (o *EventsourceLogEntry) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *EventsourceLogEntry) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *EventsourceLogEntry) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *EventsourceLogEntry) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetNamespace

`func (o *EventsourceLogEntry) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EventsourceLogEntry) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EventsourceLogEntry) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *EventsourceLogEntry) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTime

`func (o *EventsourceLogEntry) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EventsourceLogEntry) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EventsourceLogEntry) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *EventsourceLogEntry) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


