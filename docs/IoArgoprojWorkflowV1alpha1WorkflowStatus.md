# IoArgoprojWorkflowV1alpha1WorkflowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactRepositoryRef** | Pointer to [**IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus**](IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus.md) |  | [optional] 
**CompressedNodes** | Pointer to **string** | Compressed and base64 decoded Nodes map | [optional] 
**Conditions** | Pointer to [**[]IoArgoprojWorkflowV1alpha1Condition**](IoArgoprojWorkflowV1alpha1Condition.md) | Conditions is a list of conditions the Workflow may have | [optional] 
**EstimatedDuration** | Pointer to **int32** | EstimatedDuration in seconds. | [optional] 
**FinishedAt** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** | A human readable message indicating details about why the workflow is in this condition. | [optional] 
**Nodes** | Pointer to [**map[string]IoArgoprojWorkflowV1alpha1NodeStatus**](IoArgoprojWorkflowV1alpha1NodeStatus.md) | Nodes is a mapping between a node ID and the node&#39;s status. | [optional] 
**OffloadNodeStatusVersion** | Pointer to **string** | Whether on not node status has been offloaded to a database. If exists, then Nodes and CompressedNodes will be empty. This will actually be populated with a hash of the offloaded data. | [optional] 
**Outputs** | Pointer to [**IoArgoprojWorkflowV1alpha1Outputs**](IoArgoprojWorkflowV1alpha1Outputs.md) |  | [optional] 
**PersistentVolumeClaims** | Pointer to [**[]IoK8sApiCoreV1Volume**](IoK8sApiCoreV1Volume.md) | PersistentVolumeClaims tracks all PVCs that were created as part of the io.argoproj.workflow.v1alpha1. The contents of this list are drained at the end of the workflow. | [optional] 
**Phase** | Pointer to **string** | Phase a simple, high-level summary of where the workflow is in its lifecycle. | [optional] 
**Progress** | Pointer to **string** | Progress to completion | [optional] 
**ResourcesDuration** | Pointer to **map[string]int64** | ResourcesDuration is the total for the workflow | [optional] 
**StartedAt** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**StoredTemplates** | Pointer to [**map[string]IoArgoprojWorkflowV1alpha1Template**](IoArgoprojWorkflowV1alpha1Template.md) | StoredTemplates is a mapping between a template ref and the node&#39;s status. | [optional] 
**StoredWorkflowTemplateSpec** | Pointer to [**IoArgoprojWorkflowV1alpha1WorkflowSpec**](IoArgoprojWorkflowV1alpha1WorkflowSpec.md) |  | [optional] 
**Synchronization** | Pointer to [**IoArgoprojWorkflowV1alpha1SynchronizationStatus**](IoArgoprojWorkflowV1alpha1SynchronizationStatus.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1WorkflowStatus

`func NewIoArgoprojWorkflowV1alpha1WorkflowStatus() *IoArgoprojWorkflowV1alpha1WorkflowStatus`

NewIoArgoprojWorkflowV1alpha1WorkflowStatus instantiates a new IoArgoprojWorkflowV1alpha1WorkflowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1WorkflowStatusWithDefaults

`func NewIoArgoprojWorkflowV1alpha1WorkflowStatusWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowStatus`

NewIoArgoprojWorkflowV1alpha1WorkflowStatusWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactRepositoryRef

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetArtifactRepositoryRef() IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus`

GetArtifactRepositoryRef returns the ArtifactRepositoryRef field if non-nil, zero value otherwise.

### GetArtifactRepositoryRefOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetArtifactRepositoryRefOk() (*IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus, bool)`

GetArtifactRepositoryRefOk returns a tuple with the ArtifactRepositoryRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactRepositoryRef

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetArtifactRepositoryRef(v IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus)`

SetArtifactRepositoryRef sets ArtifactRepositoryRef field to given value.

### HasArtifactRepositoryRef

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasArtifactRepositoryRef() bool`

HasArtifactRepositoryRef returns a boolean if a field has been set.

### GetCompressedNodes

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetCompressedNodes() string`

GetCompressedNodes returns the CompressedNodes field if non-nil, zero value otherwise.

### GetCompressedNodesOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetCompressedNodesOk() (*string, bool)`

GetCompressedNodesOk returns a tuple with the CompressedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedNodes

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetCompressedNodes(v string)`

SetCompressedNodes sets CompressedNodes field to given value.

### HasCompressedNodes

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasCompressedNodes() bool`

HasCompressedNodes returns a boolean if a field has been set.

### GetConditions

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetConditions() []IoArgoprojWorkflowV1alpha1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetConditionsOk() (*[]IoArgoprojWorkflowV1alpha1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetConditions(v []IoArgoprojWorkflowV1alpha1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetEstimatedDuration

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetEstimatedDuration() int32`

GetEstimatedDuration returns the EstimatedDuration field if non-nil, zero value otherwise.

### GetEstimatedDurationOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetEstimatedDurationOk() (*int32, bool)`

GetEstimatedDurationOk returns a tuple with the EstimatedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDuration

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetEstimatedDuration(v int32)`

SetEstimatedDuration sets EstimatedDuration field to given value.

### HasEstimatedDuration

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasEstimatedDuration() bool`

HasEstimatedDuration returns a boolean if a field has been set.

### GetFinishedAt

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetMessage

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNodes

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetNodes() map[string]IoArgoprojWorkflowV1alpha1NodeStatus`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetNodesOk() (*map[string]IoArgoprojWorkflowV1alpha1NodeStatus, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetNodes(v map[string]IoArgoprojWorkflowV1alpha1NodeStatus)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetOffloadNodeStatusVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetOffloadNodeStatusVersion() string`

GetOffloadNodeStatusVersion returns the OffloadNodeStatusVersion field if non-nil, zero value otherwise.

### GetOffloadNodeStatusVersionOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetOffloadNodeStatusVersionOk() (*string, bool)`

GetOffloadNodeStatusVersionOk returns a tuple with the OffloadNodeStatusVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffloadNodeStatusVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetOffloadNodeStatusVersion(v string)`

SetOffloadNodeStatusVersion sets OffloadNodeStatusVersion field to given value.

### HasOffloadNodeStatusVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasOffloadNodeStatusVersion() bool`

HasOffloadNodeStatusVersion returns a boolean if a field has been set.

### GetOutputs

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetOutputs() IoArgoprojWorkflowV1alpha1Outputs`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetOutputsOk() (*IoArgoprojWorkflowV1alpha1Outputs, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetOutputs(v IoArgoprojWorkflowV1alpha1Outputs)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetPersistentVolumeClaims

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetPersistentVolumeClaims() []IoK8sApiCoreV1Volume`

GetPersistentVolumeClaims returns the PersistentVolumeClaims field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetPersistentVolumeClaimsOk() (*[]IoK8sApiCoreV1Volume, bool)`

GetPersistentVolumeClaimsOk returns a tuple with the PersistentVolumeClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaims

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetPersistentVolumeClaims(v []IoK8sApiCoreV1Volume)`

SetPersistentVolumeClaims sets PersistentVolumeClaims field to given value.

### HasPersistentVolumeClaims

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasPersistentVolumeClaims() bool`

HasPersistentVolumeClaims returns a boolean if a field has been set.

### GetPhase

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetProgress

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetProgress(v string)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetResourcesDuration

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetResourcesDuration() map[string]int64`

GetResourcesDuration returns the ResourcesDuration field if non-nil, zero value otherwise.

### GetResourcesDurationOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetResourcesDurationOk() (*map[string]int64, bool)`

GetResourcesDurationOk returns a tuple with the ResourcesDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesDuration

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetResourcesDuration(v map[string]int64)`

SetResourcesDuration sets ResourcesDuration field to given value.

### HasResourcesDuration

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasResourcesDuration() bool`

HasResourcesDuration returns a boolean if a field has been set.

### GetStartedAt

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStoredTemplates

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetStoredTemplates() map[string]IoArgoprojWorkflowV1alpha1Template`

GetStoredTemplates returns the StoredTemplates field if non-nil, zero value otherwise.

### GetStoredTemplatesOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetStoredTemplatesOk() (*map[string]IoArgoprojWorkflowV1alpha1Template, bool)`

GetStoredTemplatesOk returns a tuple with the StoredTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredTemplates

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetStoredTemplates(v map[string]IoArgoprojWorkflowV1alpha1Template)`

SetStoredTemplates sets StoredTemplates field to given value.

### HasStoredTemplates

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasStoredTemplates() bool`

HasStoredTemplates returns a boolean if a field has been set.

### GetStoredWorkflowTemplateSpec

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetStoredWorkflowTemplateSpec() IoArgoprojWorkflowV1alpha1WorkflowSpec`

GetStoredWorkflowTemplateSpec returns the StoredWorkflowTemplateSpec field if non-nil, zero value otherwise.

### GetStoredWorkflowTemplateSpecOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetStoredWorkflowTemplateSpecOk() (*IoArgoprojWorkflowV1alpha1WorkflowSpec, bool)`

GetStoredWorkflowTemplateSpecOk returns a tuple with the StoredWorkflowTemplateSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredWorkflowTemplateSpec

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetStoredWorkflowTemplateSpec(v IoArgoprojWorkflowV1alpha1WorkflowSpec)`

SetStoredWorkflowTemplateSpec sets StoredWorkflowTemplateSpec field to given value.

### HasStoredWorkflowTemplateSpec

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasStoredWorkflowTemplateSpec() bool`

HasStoredWorkflowTemplateSpec returns a boolean if a field has been set.

### GetSynchronization

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetSynchronization() IoArgoprojWorkflowV1alpha1SynchronizationStatus`

GetSynchronization returns the Synchronization field if non-nil, zero value otherwise.

### GetSynchronizationOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) GetSynchronizationOk() (*IoArgoprojWorkflowV1alpha1SynchronizationStatus, bool)`

GetSynchronizationOk returns a tuple with the Synchronization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronization

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) SetSynchronization(v IoArgoprojWorkflowV1alpha1SynchronizationStatus)`

SetSynchronization sets Synchronization field to given value.

### HasSynchronization

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStatus) HasSynchronization() bool`

HasSynchronization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


