# IoArgoprojWorkflowV1alpha1PodGC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelSelector** | Pointer to [**IoK8sApimachineryPkgApisMetaV1LabelSelector**](IoK8sApimachineryPkgApisMetaV1LabelSelector.md) |  | [optional] 
**Strategy** | Pointer to **string** | Strategy is the strategy to use. One of \&quot;OnPodCompletion\&quot;, \&quot;OnPodSuccess\&quot;, \&quot;OnWorkflowCompletion\&quot;, \&quot;OnWorkflowSuccess\&quot; | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1PodGC

`func NewIoArgoprojWorkflowV1alpha1PodGC() *IoArgoprojWorkflowV1alpha1PodGC`

NewIoArgoprojWorkflowV1alpha1PodGC instantiates a new IoArgoprojWorkflowV1alpha1PodGC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1PodGCWithDefaults

`func NewIoArgoprojWorkflowV1alpha1PodGCWithDefaults() *IoArgoprojWorkflowV1alpha1PodGC`

NewIoArgoprojWorkflowV1alpha1PodGCWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1PodGC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelSelector

`func (o *IoArgoprojWorkflowV1alpha1PodGC) GetLabelSelector() IoK8sApimachineryPkgApisMetaV1LabelSelector`

GetLabelSelector returns the LabelSelector field if non-nil, zero value otherwise.

### GetLabelSelectorOk

`func (o *IoArgoprojWorkflowV1alpha1PodGC) GetLabelSelectorOk() (*IoK8sApimachineryPkgApisMetaV1LabelSelector, bool)`

GetLabelSelectorOk returns a tuple with the LabelSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSelector

`func (o *IoArgoprojWorkflowV1alpha1PodGC) SetLabelSelector(v IoK8sApimachineryPkgApisMetaV1LabelSelector)`

SetLabelSelector sets LabelSelector field to given value.

### HasLabelSelector

`func (o *IoArgoprojWorkflowV1alpha1PodGC) HasLabelSelector() bool`

HasLabelSelector returns a boolean if a field has been set.

### GetStrategy

`func (o *IoArgoprojWorkflowV1alpha1PodGC) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *IoArgoprojWorkflowV1alpha1PodGC) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *IoArgoprojWorkflowV1alpha1PodGC) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *IoArgoprojWorkflowV1alpha1PodGC) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


