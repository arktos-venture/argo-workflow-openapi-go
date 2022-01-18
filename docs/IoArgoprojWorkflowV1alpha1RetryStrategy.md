# IoArgoprojWorkflowV1alpha1RetryStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affinity** | Pointer to [**IoArgoprojWorkflowV1alpha1RetryAffinity**](IoArgoprojWorkflowV1alpha1RetryAffinity.md) |  | [optional] 
**Backoff** | Pointer to [**IoArgoprojWorkflowV1alpha1Backoff**](IoArgoprojWorkflowV1alpha1Backoff.md) |  | [optional] 
**Expression** | Pointer to **string** | Expression is a condition expression for when a node will be retried. If it evaluates to false, the node will not be retried and the retry strategy will be ignored/ | [optional] 
**Limit** | Pointer to **string** |  | [optional] 
**RetryPolicy** | Pointer to **string** | RetryPolicy is a policy of NodePhase statuses that will be retried | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1RetryStrategy

`func NewIoArgoprojWorkflowV1alpha1RetryStrategy() *IoArgoprojWorkflowV1alpha1RetryStrategy`

NewIoArgoprojWorkflowV1alpha1RetryStrategy instantiates a new IoArgoprojWorkflowV1alpha1RetryStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1RetryStrategyWithDefaults

`func NewIoArgoprojWorkflowV1alpha1RetryStrategyWithDefaults() *IoArgoprojWorkflowV1alpha1RetryStrategy`

NewIoArgoprojWorkflowV1alpha1RetryStrategyWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1RetryStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffinity

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetAffinity() IoArgoprojWorkflowV1alpha1RetryAffinity`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetAffinityOk() (*IoArgoprojWorkflowV1alpha1RetryAffinity, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) SetAffinity(v IoArgoprojWorkflowV1alpha1RetryAffinity)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetBackoff

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetBackoff() IoArgoprojWorkflowV1alpha1Backoff`

GetBackoff returns the Backoff field if non-nil, zero value otherwise.

### GetBackoffOk

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetBackoffOk() (*IoArgoprojWorkflowV1alpha1Backoff, bool)`

GetBackoffOk returns a tuple with the Backoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoff

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) SetBackoff(v IoArgoprojWorkflowV1alpha1Backoff)`

SetBackoff sets Backoff field to given value.

### HasBackoff

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) HasBackoff() bool`

HasBackoff returns a boolean if a field has been set.

### GetExpression

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetLimit

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetRetryPolicy

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetRetryPolicy() string`

GetRetryPolicy returns the RetryPolicy field if non-nil, zero value otherwise.

### GetRetryPolicyOk

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetRetryPolicyOk() (*string, bool)`

GetRetryPolicyOk returns a tuple with the RetryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryPolicy

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) SetRetryPolicy(v string)`

SetRetryPolicy sets RetryPolicy field to given value.

### HasRetryPolicy

`func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) HasRetryPolicy() bool`

HasRetryPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


