# SensorLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependencyName** | Pointer to **string** |  | [optional] 
**EventContext** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**SensorName** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**TriggerName** | Pointer to **string** |  | [optional] 

## Methods

### NewSensorLogEntry

`func NewSensorLogEntry() *SensorLogEntry`

NewSensorLogEntry instantiates a new SensorLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorLogEntryWithDefaults

`func NewSensorLogEntryWithDefaults() *SensorLogEntry`

NewSensorLogEntryWithDefaults instantiates a new SensorLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencyName

`func (o *SensorLogEntry) GetDependencyName() string`

GetDependencyName returns the DependencyName field if non-nil, zero value otherwise.

### GetDependencyNameOk

`func (o *SensorLogEntry) GetDependencyNameOk() (*string, bool)`

GetDependencyNameOk returns a tuple with the DependencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyName

`func (o *SensorLogEntry) SetDependencyName(v string)`

SetDependencyName sets DependencyName field to given value.

### HasDependencyName

`func (o *SensorLogEntry) HasDependencyName() bool`

HasDependencyName returns a boolean if a field has been set.

### GetEventContext

`func (o *SensorLogEntry) GetEventContext() string`

GetEventContext returns the EventContext field if non-nil, zero value otherwise.

### GetEventContextOk

`func (o *SensorLogEntry) GetEventContextOk() (*string, bool)`

GetEventContextOk returns a tuple with the EventContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventContext

`func (o *SensorLogEntry) SetEventContext(v string)`

SetEventContext sets EventContext field to given value.

### HasEventContext

`func (o *SensorLogEntry) HasEventContext() bool`

HasEventContext returns a boolean if a field has been set.

### GetLevel

`func (o *SensorLogEntry) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SensorLogEntry) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SensorLogEntry) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *SensorLogEntry) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMsg

`func (o *SensorLogEntry) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SensorLogEntry) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SensorLogEntry) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *SensorLogEntry) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetNamespace

`func (o *SensorLogEntry) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SensorLogEntry) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SensorLogEntry) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SensorLogEntry) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSensorName

`func (o *SensorLogEntry) GetSensorName() string`

GetSensorName returns the SensorName field if non-nil, zero value otherwise.

### GetSensorNameOk

`func (o *SensorLogEntry) GetSensorNameOk() (*string, bool)`

GetSensorNameOk returns a tuple with the SensorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorName

`func (o *SensorLogEntry) SetSensorName(v string)`

SetSensorName sets SensorName field to given value.

### HasSensorName

`func (o *SensorLogEntry) HasSensorName() bool`

HasSensorName returns a boolean if a field has been set.

### GetTime

`func (o *SensorLogEntry) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SensorLogEntry) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SensorLogEntry) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *SensorLogEntry) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTriggerName

`func (o *SensorLogEntry) GetTriggerName() string`

GetTriggerName returns the TriggerName field if non-nil, zero value otherwise.

### GetTriggerNameOk

`func (o *SensorLogEntry) GetTriggerNameOk() (*string, bool)`

GetTriggerNameOk returns a tuple with the TriggerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerName

`func (o *SensorLogEntry) SetTriggerName(v string)`

SetTriggerName sets TriggerName field to given value.

### HasTriggerName

`func (o *SensorLogEntry) HasTriggerName() bool`

HasTriggerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


