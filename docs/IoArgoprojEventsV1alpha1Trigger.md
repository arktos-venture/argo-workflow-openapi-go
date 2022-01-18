# IoArgoprojEventsV1alpha1Trigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) |  | [optional] 
**Policy** | Pointer to [**IoArgoprojEventsV1alpha1TriggerPolicy**](IoArgoprojEventsV1alpha1TriggerPolicy.md) |  | [optional] 
**RetryStrategy** | Pointer to [**IoArgoprojEventsV1alpha1Backoff**](IoArgoprojEventsV1alpha1Backoff.md) |  | [optional] 
**Template** | Pointer to [**IoArgoprojEventsV1alpha1TriggerTemplate**](IoArgoprojEventsV1alpha1TriggerTemplate.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1Trigger

`func NewIoArgoprojEventsV1alpha1Trigger() *IoArgoprojEventsV1alpha1Trigger`

NewIoArgoprojEventsV1alpha1Trigger instantiates a new IoArgoprojEventsV1alpha1Trigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1TriggerWithDefaults

`func NewIoArgoprojEventsV1alpha1TriggerWithDefaults() *IoArgoprojEventsV1alpha1Trigger`

NewIoArgoprojEventsV1alpha1TriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1Trigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *IoArgoprojEventsV1alpha1Trigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojEventsV1alpha1Trigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojEventsV1alpha1Trigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojEventsV1alpha1Trigger) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPolicy

`func (o *IoArgoprojEventsV1alpha1Trigger) GetPolicy() IoArgoprojEventsV1alpha1TriggerPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *IoArgoprojEventsV1alpha1Trigger) GetPolicyOk() (*IoArgoprojEventsV1alpha1TriggerPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *IoArgoprojEventsV1alpha1Trigger) SetPolicy(v IoArgoprojEventsV1alpha1TriggerPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *IoArgoprojEventsV1alpha1Trigger) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRetryStrategy

`func (o *IoArgoprojEventsV1alpha1Trigger) GetRetryStrategy() IoArgoprojEventsV1alpha1Backoff`

GetRetryStrategy returns the RetryStrategy field if non-nil, zero value otherwise.

### GetRetryStrategyOk

`func (o *IoArgoprojEventsV1alpha1Trigger) GetRetryStrategyOk() (*IoArgoprojEventsV1alpha1Backoff, bool)`

GetRetryStrategyOk returns a tuple with the RetryStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryStrategy

`func (o *IoArgoprojEventsV1alpha1Trigger) SetRetryStrategy(v IoArgoprojEventsV1alpha1Backoff)`

SetRetryStrategy sets RetryStrategy field to given value.

### HasRetryStrategy

`func (o *IoArgoprojEventsV1alpha1Trigger) HasRetryStrategy() bool`

HasRetryStrategy returns a boolean if a field has been set.

### GetTemplate

`func (o *IoArgoprojEventsV1alpha1Trigger) GetTemplate() IoArgoprojEventsV1alpha1TriggerTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoArgoprojEventsV1alpha1Trigger) GetTemplateOk() (*IoArgoprojEventsV1alpha1TriggerTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoArgoprojEventsV1alpha1Trigger) SetTemplate(v IoArgoprojEventsV1alpha1TriggerTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IoArgoprojEventsV1alpha1Trigger) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


