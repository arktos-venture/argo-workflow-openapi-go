# IoArgoprojWorkflowV1alpha1CronWorkflowSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConcurrencyPolicy** | Pointer to **string** | ConcurrencyPolicy is the K8s-style concurrency policy that will be used | [optional] 
**FailedJobsHistoryLimit** | Pointer to **int32** | FailedJobsHistoryLimit is the number of failed jobs to be kept at a time | [optional] 
**Schedule** | **string** | Schedule is a schedule to run the Workflow in Cron format | 
**StartingDeadlineSeconds** | Pointer to **int32** | StartingDeadlineSeconds is the K8s-style deadline that will limit the time a CronWorkflow will be run after its original scheduled time if it is missed. | [optional] 
**SuccessfulJobsHistoryLimit** | Pointer to **int32** | SuccessfulJobsHistoryLimit is the number of successful jobs to be kept at a time | [optional] 
**Suspend** | Pointer to **bool** | Suspend is a flag that will stop new CronWorkflows from running if set to true | [optional] 
**Timezone** | Pointer to **string** | Timezone is the timezone against which the cron schedule will be calculated, e.g. \&quot;Asia/Tokyo\&quot;. Default is machine&#39;s local time. | [optional] 
**WorkflowMetadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**WorkflowSpec** | [**IoArgoprojWorkflowV1alpha1WorkflowSpec**](IoArgoprojWorkflowV1alpha1WorkflowSpec.md) |  | 

## Methods

### NewIoArgoprojWorkflowV1alpha1CronWorkflowSpec

`func NewIoArgoprojWorkflowV1alpha1CronWorkflowSpec(schedule string, workflowSpec IoArgoprojWorkflowV1alpha1WorkflowSpec, ) *IoArgoprojWorkflowV1alpha1CronWorkflowSpec`

NewIoArgoprojWorkflowV1alpha1CronWorkflowSpec instantiates a new IoArgoprojWorkflowV1alpha1CronWorkflowSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1CronWorkflowSpecWithDefaults

`func NewIoArgoprojWorkflowV1alpha1CronWorkflowSpecWithDefaults() *IoArgoprojWorkflowV1alpha1CronWorkflowSpec`

NewIoArgoprojWorkflowV1alpha1CronWorkflowSpecWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1CronWorkflowSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcurrencyPolicy

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetConcurrencyPolicy() string`

GetConcurrencyPolicy returns the ConcurrencyPolicy field if non-nil, zero value otherwise.

### GetConcurrencyPolicyOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetConcurrencyPolicyOk() (*string, bool)`

GetConcurrencyPolicyOk returns a tuple with the ConcurrencyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrencyPolicy

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) SetConcurrencyPolicy(v string)`

SetConcurrencyPolicy sets ConcurrencyPolicy field to given value.

### HasConcurrencyPolicy

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) HasConcurrencyPolicy() bool`

HasConcurrencyPolicy returns a boolean if a field has been set.

### GetFailedJobsHistoryLimit

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetFailedJobsHistoryLimit() int32`

GetFailedJobsHistoryLimit returns the FailedJobsHistoryLimit field if non-nil, zero value otherwise.

### GetFailedJobsHistoryLimitOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetFailedJobsHistoryLimitOk() (*int32, bool)`

GetFailedJobsHistoryLimitOk returns a tuple with the FailedJobsHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedJobsHistoryLimit

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) SetFailedJobsHistoryLimit(v int32)`

SetFailedJobsHistoryLimit sets FailedJobsHistoryLimit field to given value.

### HasFailedJobsHistoryLimit

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) HasFailedJobsHistoryLimit() bool`

HasFailedJobsHistoryLimit returns a boolean if a field has been set.

### GetSchedule

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetStartingDeadlineSeconds

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetStartingDeadlineSeconds() int32`

GetStartingDeadlineSeconds returns the StartingDeadlineSeconds field if non-nil, zero value otherwise.

### GetStartingDeadlineSecondsOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetStartingDeadlineSecondsOk() (*int32, bool)`

GetStartingDeadlineSecondsOk returns a tuple with the StartingDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingDeadlineSeconds

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) SetStartingDeadlineSeconds(v int32)`

SetStartingDeadlineSeconds sets StartingDeadlineSeconds field to given value.

### HasStartingDeadlineSeconds

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) HasStartingDeadlineSeconds() bool`

HasStartingDeadlineSeconds returns a boolean if a field has been set.

### GetSuccessfulJobsHistoryLimit

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetSuccessfulJobsHistoryLimit() int32`

GetSuccessfulJobsHistoryLimit returns the SuccessfulJobsHistoryLimit field if non-nil, zero value otherwise.

### GetSuccessfulJobsHistoryLimitOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetSuccessfulJobsHistoryLimitOk() (*int32, bool)`

GetSuccessfulJobsHistoryLimitOk returns a tuple with the SuccessfulJobsHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulJobsHistoryLimit

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) SetSuccessfulJobsHistoryLimit(v int32)`

SetSuccessfulJobsHistoryLimit sets SuccessfulJobsHistoryLimit field to given value.

### HasSuccessfulJobsHistoryLimit

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) HasSuccessfulJobsHistoryLimit() bool`

HasSuccessfulJobsHistoryLimit returns a boolean if a field has been set.

### GetSuspend

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetSuspend() bool`

GetSuspend returns the Suspend field if non-nil, zero value otherwise.

### GetSuspendOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetSuspendOk() (*bool, bool)`

GetSuspendOk returns a tuple with the Suspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspend

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) SetSuspend(v bool)`

SetSuspend sets Suspend field to given value.

### HasSuspend

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) HasSuspend() bool`

HasSuspend returns a boolean if a field has been set.

### GetTimezone

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetWorkflowMetadata

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetWorkflowMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetWorkflowMetadata returns the WorkflowMetadata field if non-nil, zero value otherwise.

### GetWorkflowMetadataOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetWorkflowMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetWorkflowMetadataOk returns a tuple with the WorkflowMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowMetadata

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) SetWorkflowMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetWorkflowMetadata sets WorkflowMetadata field to given value.

### HasWorkflowMetadata

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) HasWorkflowMetadata() bool`

HasWorkflowMetadata returns a boolean if a field has been set.

### GetWorkflowSpec

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetWorkflowSpec() IoArgoprojWorkflowV1alpha1WorkflowSpec`

GetWorkflowSpec returns the WorkflowSpec field if non-nil, zero value otherwise.

### GetWorkflowSpecOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) GetWorkflowSpecOk() (*IoArgoprojWorkflowV1alpha1WorkflowSpec, bool)`

GetWorkflowSpecOk returns a tuple with the WorkflowSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSpec

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowSpec) SetWorkflowSpec(v IoArgoprojWorkflowV1alpha1WorkflowSpec)`

SetWorkflowSpec sets WorkflowSpec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


