# IoArgoprojWorkflowV1alpha1Backoff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **string** | Duration is the amount to back off. Default unit is seconds, but could also be a duration (e.g. \&quot;2m\&quot;, \&quot;1h\&quot;) | [optional] 
**Factor** | Pointer to **string** |  | [optional] 
**MaxDuration** | Pointer to **string** | MaxDuration is the maximum amount of time allowed for the backoff strategy | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1Backoff

`func NewIoArgoprojWorkflowV1alpha1Backoff() *IoArgoprojWorkflowV1alpha1Backoff`

NewIoArgoprojWorkflowV1alpha1Backoff instantiates a new IoArgoprojWorkflowV1alpha1Backoff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1BackoffWithDefaults

`func NewIoArgoprojWorkflowV1alpha1BackoffWithDefaults() *IoArgoprojWorkflowV1alpha1Backoff`

NewIoArgoprojWorkflowV1alpha1BackoffWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Backoff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *IoArgoprojWorkflowV1alpha1Backoff) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *IoArgoprojWorkflowV1alpha1Backoff) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *IoArgoprojWorkflowV1alpha1Backoff) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *IoArgoprojWorkflowV1alpha1Backoff) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFactor

`func (o *IoArgoprojWorkflowV1alpha1Backoff) GetFactor() string`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *IoArgoprojWorkflowV1alpha1Backoff) GetFactorOk() (*string, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *IoArgoprojWorkflowV1alpha1Backoff) SetFactor(v string)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *IoArgoprojWorkflowV1alpha1Backoff) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetMaxDuration

`func (o *IoArgoprojWorkflowV1alpha1Backoff) GetMaxDuration() string`

GetMaxDuration returns the MaxDuration field if non-nil, zero value otherwise.

### GetMaxDurationOk

`func (o *IoArgoprojWorkflowV1alpha1Backoff) GetMaxDurationOk() (*string, bool)`

GetMaxDurationOk returns a tuple with the MaxDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDuration

`func (o *IoArgoprojWorkflowV1alpha1Backoff) SetMaxDuration(v string)`

SetMaxDuration sets MaxDuration field to given value.

### HasMaxDuration

`func (o *IoArgoprojWorkflowV1alpha1Backoff) HasMaxDuration() bool`

HasMaxDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


