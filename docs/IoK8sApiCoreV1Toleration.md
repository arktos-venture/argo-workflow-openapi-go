# IoK8sApiCoreV1Toleration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | Pointer to **string** | Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute. | [optional] 
**Key** | Pointer to **string** | Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys. | [optional] 
**Operator** | Pointer to **string** | Operator represents a key&#39;s relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category. | [optional] 
**TolerationSeconds** | Pointer to **int32** | TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system. | [optional] 
**Value** | Pointer to **string** | Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string. | [optional] 

## Methods

### NewIoK8sApiCoreV1Toleration

`func NewIoK8sApiCoreV1Toleration() *IoK8sApiCoreV1Toleration`

NewIoK8sApiCoreV1Toleration instantiates a new IoK8sApiCoreV1Toleration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1TolerationWithDefaults

`func NewIoK8sApiCoreV1TolerationWithDefaults() *IoK8sApiCoreV1Toleration`

NewIoK8sApiCoreV1TolerationWithDefaults instantiates a new IoK8sApiCoreV1Toleration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *IoK8sApiCoreV1Toleration) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *IoK8sApiCoreV1Toleration) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *IoK8sApiCoreV1Toleration) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *IoK8sApiCoreV1Toleration) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetKey

`func (o *IoK8sApiCoreV1Toleration) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IoK8sApiCoreV1Toleration) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IoK8sApiCoreV1Toleration) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IoK8sApiCoreV1Toleration) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOperator

`func (o *IoK8sApiCoreV1Toleration) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *IoK8sApiCoreV1Toleration) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *IoK8sApiCoreV1Toleration) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *IoK8sApiCoreV1Toleration) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetTolerationSeconds

`func (o *IoK8sApiCoreV1Toleration) GetTolerationSeconds() int32`

GetTolerationSeconds returns the TolerationSeconds field if non-nil, zero value otherwise.

### GetTolerationSecondsOk

`func (o *IoK8sApiCoreV1Toleration) GetTolerationSecondsOk() (*int32, bool)`

GetTolerationSecondsOk returns a tuple with the TolerationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerationSeconds

`func (o *IoK8sApiCoreV1Toleration) SetTolerationSeconds(v int32)`

SetTolerationSeconds sets TolerationSeconds field to given value.

### HasTolerationSeconds

`func (o *IoK8sApiCoreV1Toleration) HasTolerationSeconds() bool`

HasTolerationSeconds returns a boolean if a field has been set.

### GetValue

`func (o *IoK8sApiCoreV1Toleration) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IoK8sApiCoreV1Toleration) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IoK8sApiCoreV1Toleration) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IoK8sApiCoreV1Toleration) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


