# IoArgoprojEventsV1alpha1CalendarEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExclusionDates** | Pointer to **[]string** |  | [optional] 
**Interval** | Pointer to **string** | Interval is a string that describes an interval duration, e.g. 1s, 30m, 2h... | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Persistence** | Pointer to [**IoArgoprojEventsV1alpha1EventPersistence**](IoArgoprojEventsV1alpha1EventPersistence.md) |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**UserPayload** | Pointer to **string** | UserPayload will be sent to sensor as extra data once the event is triggered +optional Deprecated: will be removed in v1.5. Please use Metadata instead. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1CalendarEventSource

`func NewIoArgoprojEventsV1alpha1CalendarEventSource() *IoArgoprojEventsV1alpha1CalendarEventSource`

NewIoArgoprojEventsV1alpha1CalendarEventSource instantiates a new IoArgoprojEventsV1alpha1CalendarEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1CalendarEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1CalendarEventSourceWithDefaults() *IoArgoprojEventsV1alpha1CalendarEventSource`

NewIoArgoprojEventsV1alpha1CalendarEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1CalendarEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclusionDates

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetExclusionDates() []string`

GetExclusionDates returns the ExclusionDates field if non-nil, zero value otherwise.

### GetExclusionDatesOk

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetExclusionDatesOk() (*[]string, bool)`

GetExclusionDatesOk returns a tuple with the ExclusionDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionDates

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) SetExclusionDates(v []string)`

SetExclusionDates sets ExclusionDates field to given value.

### HasExclusionDates

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) HasExclusionDates() bool`

HasExclusionDates returns a boolean if a field has been set.

### GetInterval

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPersistence

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetPersistence() IoArgoprojEventsV1alpha1EventPersistence`

GetPersistence returns the Persistence field if non-nil, zero value otherwise.

### GetPersistenceOk

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetPersistenceOk() (*IoArgoprojEventsV1alpha1EventPersistence, bool)`

GetPersistenceOk returns a tuple with the Persistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistence

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) SetPersistence(v IoArgoprojEventsV1alpha1EventPersistence)`

SetPersistence sets Persistence field to given value.

### HasPersistence

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) HasPersistence() bool`

HasPersistence returns a boolean if a field has been set.

### GetSchedule

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTimezone

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUserPayload

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetUserPayload() string`

GetUserPayload returns the UserPayload field if non-nil, zero value otherwise.

### GetUserPayloadOk

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) GetUserPayloadOk() (*string, bool)`

GetUserPayloadOk returns a tuple with the UserPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPayload

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) SetUserPayload(v string)`

SetUserPayload sets UserPayload field to given value.

### HasUserPayload

`func (o *IoArgoprojEventsV1alpha1CalendarEventSource) HasUserPayload() bool`

HasUserPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


