# IoArgoprojWorkflowV1alpha1DAGTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailFast** | Pointer to **bool** | This flag is for DAG logic. The DAG logic has a built-in \&quot;fail fast\&quot; feature to stop scheduling new steps, as soon as it detects that one of the DAG nodes is failed. Then it waits until all DAG nodes are completed before failing the DAG itself. The FailFast flag default is true,  if set to false, it will allow a DAG to run all branches of the DAG to completion (either success or failure), regardless of the failed outcomes of branches in the DAG. More info and example about this feature at https://github.com/argoproj/argo-workflows/issues/1442 | [optional] 
**Target** | Pointer to **string** | Target are one or more names of targets to execute in a DAG | [optional] 
**Tasks** | [**[]IoArgoprojWorkflowV1alpha1DAGTask**](IoArgoprojWorkflowV1alpha1DAGTask.md) | Tasks are a list of DAG tasks | 

## Methods

### NewIoArgoprojWorkflowV1alpha1DAGTemplate

`func NewIoArgoprojWorkflowV1alpha1DAGTemplate(tasks []IoArgoprojWorkflowV1alpha1DAGTask, ) *IoArgoprojWorkflowV1alpha1DAGTemplate`

NewIoArgoprojWorkflowV1alpha1DAGTemplate instantiates a new IoArgoprojWorkflowV1alpha1DAGTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1DAGTemplateWithDefaults

`func NewIoArgoprojWorkflowV1alpha1DAGTemplateWithDefaults() *IoArgoprojWorkflowV1alpha1DAGTemplate`

NewIoArgoprojWorkflowV1alpha1DAGTemplateWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1DAGTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailFast

`func (o *IoArgoprojWorkflowV1alpha1DAGTemplate) GetFailFast() bool`

GetFailFast returns the FailFast field if non-nil, zero value otherwise.

### GetFailFastOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTemplate) GetFailFastOk() (*bool, bool)`

GetFailFastOk returns a tuple with the FailFast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailFast

`func (o *IoArgoprojWorkflowV1alpha1DAGTemplate) SetFailFast(v bool)`

SetFailFast sets FailFast field to given value.

### HasFailFast

`func (o *IoArgoprojWorkflowV1alpha1DAGTemplate) HasFailFast() bool`

HasFailFast returns a boolean if a field has been set.

### GetTarget

`func (o *IoArgoprojWorkflowV1alpha1DAGTemplate) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTemplate) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *IoArgoprojWorkflowV1alpha1DAGTemplate) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *IoArgoprojWorkflowV1alpha1DAGTemplate) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTasks

`func (o *IoArgoprojWorkflowV1alpha1DAGTemplate) GetTasks() []IoArgoprojWorkflowV1alpha1DAGTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTemplate) GetTasksOk() (*[]IoArgoprojWorkflowV1alpha1DAGTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *IoArgoprojWorkflowV1alpha1DAGTemplate) SetTasks(v []IoArgoprojWorkflowV1alpha1DAGTask)`

SetTasks sets Tasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


