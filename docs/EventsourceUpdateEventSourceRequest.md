# EventsourceUpdateEventSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSource** | Pointer to [**IoArgoprojEventsV1alpha1EventSource**](IoArgoprojEventsV1alpha1EventSource.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewEventsourceUpdateEventSourceRequest

`func NewEventsourceUpdateEventSourceRequest() *EventsourceUpdateEventSourceRequest`

NewEventsourceUpdateEventSourceRequest instantiates a new EventsourceUpdateEventSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsourceUpdateEventSourceRequestWithDefaults

`func NewEventsourceUpdateEventSourceRequestWithDefaults() *EventsourceUpdateEventSourceRequest`

NewEventsourceUpdateEventSourceRequestWithDefaults instantiates a new EventsourceUpdateEventSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSource

`func (o *EventsourceUpdateEventSourceRequest) GetEventSource() IoArgoprojEventsV1alpha1EventSource`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *EventsourceUpdateEventSourceRequest) GetEventSourceOk() (*IoArgoprojEventsV1alpha1EventSource, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *EventsourceUpdateEventSourceRequest) SetEventSource(v IoArgoprojEventsV1alpha1EventSource)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *EventsourceUpdateEventSourceRequest) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetName

`func (o *EventsourceUpdateEventSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventsourceUpdateEventSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventsourceUpdateEventSourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventsourceUpdateEventSourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *EventsourceUpdateEventSourceRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EventsourceUpdateEventSourceRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EventsourceUpdateEventSourceRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *EventsourceUpdateEventSourceRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


