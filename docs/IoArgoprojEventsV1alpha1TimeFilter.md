# IoArgoprojEventsV1alpha1TimeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** | Start is the beginning of a time window in UTC. Before this time, events for this dependency are ignored. Format is hh:mm:ss. | [optional] 
**Stop** | Pointer to **string** | Stop is the end of a time window in UTC. After or equal to this time, events for this dependency are ignored and Format is hh:mm:ss. If it is smaller than Start, it is treated as next day of Start (e.g.: 22:00:00-01:00:00 means 22:00:00-25:00:00). | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1TimeFilter

`func NewIoArgoprojEventsV1alpha1TimeFilter() *IoArgoprojEventsV1alpha1TimeFilter`

NewIoArgoprojEventsV1alpha1TimeFilter instantiates a new IoArgoprojEventsV1alpha1TimeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1TimeFilterWithDefaults

`func NewIoArgoprojEventsV1alpha1TimeFilterWithDefaults() *IoArgoprojEventsV1alpha1TimeFilter`

NewIoArgoprojEventsV1alpha1TimeFilterWithDefaults instantiates a new IoArgoprojEventsV1alpha1TimeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *IoArgoprojEventsV1alpha1TimeFilter) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *IoArgoprojEventsV1alpha1TimeFilter) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *IoArgoprojEventsV1alpha1TimeFilter) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *IoArgoprojEventsV1alpha1TimeFilter) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStop

`func (o *IoArgoprojEventsV1alpha1TimeFilter) GetStop() string`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *IoArgoprojEventsV1alpha1TimeFilter) GetStopOk() (*string, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *IoArgoprojEventsV1alpha1TimeFilter) SetStop(v string)`

SetStop sets Stop field to given value.

### HasStop

`func (o *IoArgoprojEventsV1alpha1TimeFilter) HasStop() bool`

HasStop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


