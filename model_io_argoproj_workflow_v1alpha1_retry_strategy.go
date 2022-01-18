/*
Argo Server API

You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`

API version: VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IoArgoprojWorkflowV1alpha1RetryStrategy RetryStrategy provides controls on how to retry a workflow step
type IoArgoprojWorkflowV1alpha1RetryStrategy struct {
	Affinity *IoArgoprojWorkflowV1alpha1RetryAffinity `json:"affinity,omitempty"`
	Backoff *IoArgoprojWorkflowV1alpha1Backoff `json:"backoff,omitempty"`
	// Expression is a condition expression for when a node will be retried. If it evaluates to false, the node will not be retried and the retry strategy will be ignored/
	Expression *string `json:"expression,omitempty"`
	Limit *string `json:"limit,omitempty"`
	// RetryPolicy is a policy of NodePhase statuses that will be retried
	RetryPolicy *string `json:"retryPolicy,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1RetryStrategy instantiates a new IoArgoprojWorkflowV1alpha1RetryStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1RetryStrategy() *IoArgoprojWorkflowV1alpha1RetryStrategy {
	this := IoArgoprojWorkflowV1alpha1RetryStrategy{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1RetryStrategyWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1RetryStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1RetryStrategyWithDefaults() *IoArgoprojWorkflowV1alpha1RetryStrategy {
	this := IoArgoprojWorkflowV1alpha1RetryStrategy{}
	return &this
}

// GetAffinity returns the Affinity field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetAffinity() IoArgoprojWorkflowV1alpha1RetryAffinity {
	if o == nil || o.Affinity == nil {
		var ret IoArgoprojWorkflowV1alpha1RetryAffinity
		return ret
	}
	return *o.Affinity
}

// GetAffinityOk returns a tuple with the Affinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetAffinityOk() (*IoArgoprojWorkflowV1alpha1RetryAffinity, bool) {
	if o == nil || o.Affinity == nil {
		return nil, false
	}
	return o.Affinity, true
}

// HasAffinity returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) HasAffinity() bool {
	if o != nil && o.Affinity != nil {
		return true
	}

	return false
}

// SetAffinity gets a reference to the given IoArgoprojWorkflowV1alpha1RetryAffinity and assigns it to the Affinity field.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) SetAffinity(v IoArgoprojWorkflowV1alpha1RetryAffinity) {
	o.Affinity = &v
}

// GetBackoff returns the Backoff field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetBackoff() IoArgoprojWorkflowV1alpha1Backoff {
	if o == nil || o.Backoff == nil {
		var ret IoArgoprojWorkflowV1alpha1Backoff
		return ret
	}
	return *o.Backoff
}

// GetBackoffOk returns a tuple with the Backoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetBackoffOk() (*IoArgoprojWorkflowV1alpha1Backoff, bool) {
	if o == nil || o.Backoff == nil {
		return nil, false
	}
	return o.Backoff, true
}

// HasBackoff returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) HasBackoff() bool {
	if o != nil && o.Backoff != nil {
		return true
	}

	return false
}

// SetBackoff gets a reference to the given IoArgoprojWorkflowV1alpha1Backoff and assigns it to the Backoff field.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) SetBackoff(v IoArgoprojWorkflowV1alpha1Backoff) {
	o.Backoff = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) SetExpression(v string) {
	o.Expression = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetLimit() string {
	if o == nil || o.Limit == nil {
		var ret string
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetLimitOk() (*string, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given string and assigns it to the Limit field.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) SetLimit(v string) {
	o.Limit = &v
}

// GetRetryPolicy returns the RetryPolicy field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetRetryPolicy() string {
	if o == nil || o.RetryPolicy == nil {
		var ret string
		return ret
	}
	return *o.RetryPolicy
}

// GetRetryPolicyOk returns a tuple with the RetryPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) GetRetryPolicyOk() (*string, bool) {
	if o == nil || o.RetryPolicy == nil {
		return nil, false
	}
	return o.RetryPolicy, true
}

// HasRetryPolicy returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) HasRetryPolicy() bool {
	if o != nil && o.RetryPolicy != nil {
		return true
	}

	return false
}

// SetRetryPolicy gets a reference to the given string and assigns it to the RetryPolicy field.
func (o *IoArgoprojWorkflowV1alpha1RetryStrategy) SetRetryPolicy(v string) {
	o.RetryPolicy = &v
}

func (o IoArgoprojWorkflowV1alpha1RetryStrategy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Affinity != nil {
		toSerialize["affinity"] = o.Affinity
	}
	if o.Backoff != nil {
		toSerialize["backoff"] = o.Backoff
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.RetryPolicy != nil {
		toSerialize["retryPolicy"] = o.RetryPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1RetryStrategy struct {
	value *IoArgoprojWorkflowV1alpha1RetryStrategy
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1RetryStrategy) Get() *IoArgoprojWorkflowV1alpha1RetryStrategy {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1RetryStrategy) Set(val *IoArgoprojWorkflowV1alpha1RetryStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1RetryStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1RetryStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1RetryStrategy(val *IoArgoprojWorkflowV1alpha1RetryStrategy) *NullableIoArgoprojWorkflowV1alpha1RetryStrategy {
	return &NullableIoArgoprojWorkflowV1alpha1RetryStrategy{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1RetryStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1RetryStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


