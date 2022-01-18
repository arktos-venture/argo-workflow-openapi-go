# IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxUnavailable** | Pointer to **string** |  | [optional] 
**MinAvailable** | Pointer to **string** |  | [optional] 
**Selector** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 

## Methods

### NewIoK8sApiPolicyV1beta1PodDisruptionBudgetSpec

`func NewIoK8sApiPolicyV1beta1PodDisruptionBudgetSpec() *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec`

NewIoK8sApiPolicyV1beta1PodDisruptionBudgetSpec instantiates a new IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiPolicyV1beta1PodDisruptionBudgetSpecWithDefaults

`func NewIoK8sApiPolicyV1beta1PodDisruptionBudgetSpecWithDefaults() *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec`

NewIoK8sApiPolicyV1beta1PodDisruptionBudgetSpecWithDefaults instantiates a new IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxUnavailable

`func (o *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec) GetMaxUnavailable() string`

GetMaxUnavailable returns the MaxUnavailable field if non-nil, zero value otherwise.

### GetMaxUnavailableOk

`func (o *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec) GetMaxUnavailableOk() (*string, bool)`

GetMaxUnavailableOk returns a tuple with the MaxUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnavailable

`func (o *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec) SetMaxUnavailable(v string)`

SetMaxUnavailable sets MaxUnavailable field to given value.

### HasMaxUnavailable

`func (o *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec) HasMaxUnavailable() bool`

HasMaxUnavailable returns a boolean if a field has been set.

### GetMinAvailable

`func (o *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec) GetMinAvailable() string`

GetMinAvailable returns the MinAvailable field if non-nil, zero value otherwise.

### GetMinAvailableOk

`func (o *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec) GetMinAvailableOk() (*string, bool)`

GetMinAvailableOk returns a tuple with the MinAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAvailable

`func (o *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec) SetMinAvailable(v string)`

SetMinAvailable sets MinAvailable field to given value.

### HasMinAvailable

`func (o *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec) HasMinAvailable() bool`

HasMinAvailable returns a boolean if a field has been set.

### GetSelector

`func (o *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec) GetSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec) GetSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec) SetSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


