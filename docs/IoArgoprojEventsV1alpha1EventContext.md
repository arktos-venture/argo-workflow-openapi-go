# IoArgoprojEventsV1alpha1EventContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datacontenttype** | Pointer to **string** | DataContentType - A MIME (RFC2046) string describing the media type of &#x60;data&#x60;. | [optional] 
**Id** | Pointer to **string** | ID of the event; must be non-empty and unique within the scope of the producer. | [optional] 
**Source** | Pointer to **string** | Source - A URI describing the event producer. | [optional] 
**Specversion** | Pointer to **string** | SpecVersion - The version of the CloudEvents specification used by the io.argoproj.workflow.v1alpha1. | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Type** | Pointer to **string** | Type - The type of the occurrence which has happened. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1EventContext

`func NewIoArgoprojEventsV1alpha1EventContext() *IoArgoprojEventsV1alpha1EventContext`

NewIoArgoprojEventsV1alpha1EventContext instantiates a new IoArgoprojEventsV1alpha1EventContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1EventContextWithDefaults

`func NewIoArgoprojEventsV1alpha1EventContextWithDefaults() *IoArgoprojEventsV1alpha1EventContext`

NewIoArgoprojEventsV1alpha1EventContextWithDefaults instantiates a new IoArgoprojEventsV1alpha1EventContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacontenttype

`func (o *IoArgoprojEventsV1alpha1EventContext) GetDatacontenttype() string`

GetDatacontenttype returns the Datacontenttype field if non-nil, zero value otherwise.

### GetDatacontenttypeOk

`func (o *IoArgoprojEventsV1alpha1EventContext) GetDatacontenttypeOk() (*string, bool)`

GetDatacontenttypeOk returns a tuple with the Datacontenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacontenttype

`func (o *IoArgoprojEventsV1alpha1EventContext) SetDatacontenttype(v string)`

SetDatacontenttype sets Datacontenttype field to given value.

### HasDatacontenttype

`func (o *IoArgoprojEventsV1alpha1EventContext) HasDatacontenttype() bool`

HasDatacontenttype returns a boolean if a field has been set.

### GetId

`func (o *IoArgoprojEventsV1alpha1EventContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IoArgoprojEventsV1alpha1EventContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IoArgoprojEventsV1alpha1EventContext) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IoArgoprojEventsV1alpha1EventContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSource

`func (o *IoArgoprojEventsV1alpha1EventContext) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IoArgoprojEventsV1alpha1EventContext) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IoArgoprojEventsV1alpha1EventContext) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IoArgoprojEventsV1alpha1EventContext) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpecversion

`func (o *IoArgoprojEventsV1alpha1EventContext) GetSpecversion() string`

GetSpecversion returns the Specversion field if non-nil, zero value otherwise.

### GetSpecversionOk

`func (o *IoArgoprojEventsV1alpha1EventContext) GetSpecversionOk() (*string, bool)`

GetSpecversionOk returns a tuple with the Specversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecversion

`func (o *IoArgoprojEventsV1alpha1EventContext) SetSpecversion(v string)`

SetSpecversion sets Specversion field to given value.

### HasSpecversion

`func (o *IoArgoprojEventsV1alpha1EventContext) HasSpecversion() bool`

HasSpecversion returns a boolean if a field has been set.

### GetSubject

`func (o *IoArgoprojEventsV1alpha1EventContext) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *IoArgoprojEventsV1alpha1EventContext) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *IoArgoprojEventsV1alpha1EventContext) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *IoArgoprojEventsV1alpha1EventContext) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTime

`func (o *IoArgoprojEventsV1alpha1EventContext) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *IoArgoprojEventsV1alpha1EventContext) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *IoArgoprojEventsV1alpha1EventContext) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *IoArgoprojEventsV1alpha1EventContext) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetType

`func (o *IoArgoprojEventsV1alpha1EventContext) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoArgoprojEventsV1alpha1EventContext) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoArgoprojEventsV1alpha1EventContext) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoArgoprojEventsV1alpha1EventContext) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


