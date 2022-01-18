# IoArgoprojEventsV1alpha1Backoff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to [**IoArgoprojEventsV1alpha1Int64OrString**](IoArgoprojEventsV1alpha1Int64OrString.md) |  | [optional] 
**Factor** | Pointer to [**IoArgoprojEventsV1alpha1Amount**](IoArgoprojEventsV1alpha1Amount.md) |  | [optional] 
**Jitter** | Pointer to [**IoArgoprojEventsV1alpha1Amount**](IoArgoprojEventsV1alpha1Amount.md) |  | [optional] 
**Steps** | Pointer to **int32** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1Backoff

`func NewIoArgoprojEventsV1alpha1Backoff() *IoArgoprojEventsV1alpha1Backoff`

NewIoArgoprojEventsV1alpha1Backoff instantiates a new IoArgoprojEventsV1alpha1Backoff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1BackoffWithDefaults

`func NewIoArgoprojEventsV1alpha1BackoffWithDefaults() *IoArgoprojEventsV1alpha1Backoff`

NewIoArgoprojEventsV1alpha1BackoffWithDefaults instantiates a new IoArgoprojEventsV1alpha1Backoff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *IoArgoprojEventsV1alpha1Backoff) GetDuration() IoArgoprojEventsV1alpha1Int64OrString`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *IoArgoprojEventsV1alpha1Backoff) GetDurationOk() (*IoArgoprojEventsV1alpha1Int64OrString, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *IoArgoprojEventsV1alpha1Backoff) SetDuration(v IoArgoprojEventsV1alpha1Int64OrString)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *IoArgoprojEventsV1alpha1Backoff) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFactor

`func (o *IoArgoprojEventsV1alpha1Backoff) GetFactor() IoArgoprojEventsV1alpha1Amount`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *IoArgoprojEventsV1alpha1Backoff) GetFactorOk() (*IoArgoprojEventsV1alpha1Amount, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *IoArgoprojEventsV1alpha1Backoff) SetFactor(v IoArgoprojEventsV1alpha1Amount)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *IoArgoprojEventsV1alpha1Backoff) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetJitter

`func (o *IoArgoprojEventsV1alpha1Backoff) GetJitter() IoArgoprojEventsV1alpha1Amount`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *IoArgoprojEventsV1alpha1Backoff) GetJitterOk() (*IoArgoprojEventsV1alpha1Amount, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *IoArgoprojEventsV1alpha1Backoff) SetJitter(v IoArgoprojEventsV1alpha1Amount)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *IoArgoprojEventsV1alpha1Backoff) HasJitter() bool`

HasJitter returns a boolean if a field has been set.

### GetSteps

`func (o *IoArgoprojEventsV1alpha1Backoff) GetSteps() int32`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *IoArgoprojEventsV1alpha1Backoff) GetStepsOk() (*int32, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *IoArgoprojEventsV1alpha1Backoff) SetSteps(v int32)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *IoArgoprojEventsV1alpha1Backoff) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


