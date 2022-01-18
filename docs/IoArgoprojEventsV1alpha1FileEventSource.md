# IoArgoprojEventsV1alpha1FileEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Polling** | Pointer to **bool** |  | [optional] 
**WatchPathConfig** | Pointer to [**IoArgoprojEventsV1alpha1WatchPathConfig**](IoArgoprojEventsV1alpha1WatchPathConfig.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1FileEventSource

`func NewIoArgoprojEventsV1alpha1FileEventSource() *IoArgoprojEventsV1alpha1FileEventSource`

NewIoArgoprojEventsV1alpha1FileEventSource instantiates a new IoArgoprojEventsV1alpha1FileEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1FileEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1FileEventSourceWithDefaults() *IoArgoprojEventsV1alpha1FileEventSource`

NewIoArgoprojEventsV1alpha1FileEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1FileEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *IoArgoprojEventsV1alpha1FileEventSource) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *IoArgoprojEventsV1alpha1FileEventSource) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *IoArgoprojEventsV1alpha1FileEventSource) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *IoArgoprojEventsV1alpha1FileEventSource) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1FileEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1FileEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1FileEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1FileEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPolling

`func (o *IoArgoprojEventsV1alpha1FileEventSource) GetPolling() bool`

GetPolling returns the Polling field if non-nil, zero value otherwise.

### GetPollingOk

`func (o *IoArgoprojEventsV1alpha1FileEventSource) GetPollingOk() (*bool, bool)`

GetPollingOk returns a tuple with the Polling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolling

`func (o *IoArgoprojEventsV1alpha1FileEventSource) SetPolling(v bool)`

SetPolling sets Polling field to given value.

### HasPolling

`func (o *IoArgoprojEventsV1alpha1FileEventSource) HasPolling() bool`

HasPolling returns a boolean if a field has been set.

### GetWatchPathConfig

`func (o *IoArgoprojEventsV1alpha1FileEventSource) GetWatchPathConfig() IoArgoprojEventsV1alpha1WatchPathConfig`

GetWatchPathConfig returns the WatchPathConfig field if non-nil, zero value otherwise.

### GetWatchPathConfigOk

`func (o *IoArgoprojEventsV1alpha1FileEventSource) GetWatchPathConfigOk() (*IoArgoprojEventsV1alpha1WatchPathConfig, bool)`

GetWatchPathConfigOk returns a tuple with the WatchPathConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchPathConfig

`func (o *IoArgoprojEventsV1alpha1FileEventSource) SetWatchPathConfig(v IoArgoprojEventsV1alpha1WatchPathConfig)`

SetWatchPathConfig sets WatchPathConfig field to given value.

### HasWatchPathConfig

`func (o *IoArgoprojEventsV1alpha1FileEventSource) HasWatchPathConfig() bool`

HasWatchPathConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


