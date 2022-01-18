# IoArgoprojEventsV1alpha1EventDependencyFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**IoArgoprojEventsV1alpha1EventContext**](IoArgoprojEventsV1alpha1EventContext.md) |  | [optional] 
**Data** | Pointer to [**[]IoArgoprojEventsV1alpha1DataFilter**](IoArgoprojEventsV1alpha1DataFilter.md) |  | [optional] 
**Exprs** | Pointer to [**[]IoArgoprojEventsV1alpha1ExprFilter**](IoArgoprojEventsV1alpha1ExprFilter.md) | Exprs contains the list of expressions evaluated against the event payload. | [optional] 
**Time** | Pointer to [**IoArgoprojEventsV1alpha1TimeFilter**](IoArgoprojEventsV1alpha1TimeFilter.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1EventDependencyFilter

`func NewIoArgoprojEventsV1alpha1EventDependencyFilter() *IoArgoprojEventsV1alpha1EventDependencyFilter`

NewIoArgoprojEventsV1alpha1EventDependencyFilter instantiates a new IoArgoprojEventsV1alpha1EventDependencyFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1EventDependencyFilterWithDefaults

`func NewIoArgoprojEventsV1alpha1EventDependencyFilterWithDefaults() *IoArgoprojEventsV1alpha1EventDependencyFilter`

NewIoArgoprojEventsV1alpha1EventDependencyFilterWithDefaults instantiates a new IoArgoprojEventsV1alpha1EventDependencyFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) GetContext() IoArgoprojEventsV1alpha1EventContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) GetContextOk() (*IoArgoprojEventsV1alpha1EventContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) SetContext(v IoArgoprojEventsV1alpha1EventContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetData

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) GetData() []IoArgoprojEventsV1alpha1DataFilter`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) GetDataOk() (*[]IoArgoprojEventsV1alpha1DataFilter, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) SetData(v []IoArgoprojEventsV1alpha1DataFilter)`

SetData sets Data field to given value.

### HasData

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) HasData() bool`

HasData returns a boolean if a field has been set.

### GetExprs

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) GetExprs() []IoArgoprojEventsV1alpha1ExprFilter`

GetExprs returns the Exprs field if non-nil, zero value otherwise.

### GetExprsOk

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) GetExprsOk() (*[]IoArgoprojEventsV1alpha1ExprFilter, bool)`

GetExprsOk returns a tuple with the Exprs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExprs

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) SetExprs(v []IoArgoprojEventsV1alpha1ExprFilter)`

SetExprs sets Exprs field to given value.

### HasExprs

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) HasExprs() bool`

HasExprs returns a boolean if a field has been set.

### GetTime

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) GetTime() IoArgoprojEventsV1alpha1TimeFilter`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) GetTimeOk() (*IoArgoprojEventsV1alpha1TimeFilter, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) SetTime(v IoArgoprojEventsV1alpha1TimeFilter)`

SetTime sets Time field to given value.

### HasTime

`func (o *IoArgoprojEventsV1alpha1EventDependencyFilter) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


