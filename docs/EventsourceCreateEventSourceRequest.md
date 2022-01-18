# EventsourceCreateEventSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSource** | Pointer to [**IoArgoprojEventsV1alpha1EventSource**](IoArgoprojEventsV1alpha1EventSource.md) |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewEventsourceCreateEventSourceRequest

`func NewEventsourceCreateEventSourceRequest() *EventsourceCreateEventSourceRequest`

NewEventsourceCreateEventSourceRequest instantiates a new EventsourceCreateEventSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsourceCreateEventSourceRequestWithDefaults

`func NewEventsourceCreateEventSourceRequestWithDefaults() *EventsourceCreateEventSourceRequest`

NewEventsourceCreateEventSourceRequestWithDefaults instantiates a new EventsourceCreateEventSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSource

`func (o *EventsourceCreateEventSourceRequest) GetEventSource() IoArgoprojEventsV1alpha1EventSource`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *EventsourceCreateEventSourceRequest) GetEventSourceOk() (*IoArgoprojEventsV1alpha1EventSource, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *EventsourceCreateEventSourceRequest) SetEventSource(v IoArgoprojEventsV1alpha1EventSource)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *EventsourceCreateEventSourceRequest) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetNamespace

`func (o *EventsourceCreateEventSourceRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EventsourceCreateEventSourceRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EventsourceCreateEventSourceRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *EventsourceCreateEventSourceRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


