# IoArgoprojWorkflowV1alpha1CronWorkflowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | [**[]IoK8sApiCoreV1ObjectReference**](IoK8sApiCoreV1ObjectReference.md) | Active is a list of active workflows stemming from this CronWorkflow | 
**Conditions** | [**[]IoArgoprojWorkflowV1alpha1Condition**](IoArgoprojWorkflowV1alpha1Condition.md) | Conditions is a list of conditions the CronWorkflow may have | 
**LastScheduledTime** | **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | 

## Methods

### NewIoArgoprojWorkflowV1alpha1CronWorkflowStatus

`func NewIoArgoprojWorkflowV1alpha1CronWorkflowStatus(active []IoK8sApiCoreV1ObjectReference, conditions []IoArgoprojWorkflowV1alpha1Condition, lastScheduledTime time.Time, ) *IoArgoprojWorkflowV1alpha1CronWorkflowStatus`

NewIoArgoprojWorkflowV1alpha1CronWorkflowStatus instantiates a new IoArgoprojWorkflowV1alpha1CronWorkflowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1CronWorkflowStatusWithDefaults

`func NewIoArgoprojWorkflowV1alpha1CronWorkflowStatusWithDefaults() *IoArgoprojWorkflowV1alpha1CronWorkflowStatus`

NewIoArgoprojWorkflowV1alpha1CronWorkflowStatusWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1CronWorkflowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) GetActive() []IoK8sApiCoreV1ObjectReference`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) GetActiveOk() (*[]IoK8sApiCoreV1ObjectReference, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) SetActive(v []IoK8sApiCoreV1ObjectReference)`

SetActive sets Active field to given value.


### GetConditions

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) GetConditions() []IoArgoprojWorkflowV1alpha1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) GetConditionsOk() (*[]IoArgoprojWorkflowV1alpha1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) SetConditions(v []IoArgoprojWorkflowV1alpha1Condition)`

SetConditions sets Conditions field to given value.


### GetLastScheduledTime

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) GetLastScheduledTime() time.Time`

GetLastScheduledTime returns the LastScheduledTime field if non-nil, zero value otherwise.

### GetLastScheduledTimeOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) GetLastScheduledTimeOk() (*time.Time, bool)`

GetLastScheduledTimeOk returns a tuple with the LastScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScheduledTime

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowStatus) SetLastScheduledTime(v time.Time)`

SetLastScheduledTime sets LastScheduledTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


