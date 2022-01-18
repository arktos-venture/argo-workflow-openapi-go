# IoArgoprojWorkflowV1alpha1Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDeadlineSeconds** | Pointer to **string** |  | [optional] 
**Affinity** | Pointer to [**IoK8sApiCoreV1Affinity**](IoK8sApiCoreV1Affinity.md) |  | [optional] 
**ArchiveLocation** | Pointer to [**IoArgoprojWorkflowV1alpha1ArtifactLocation**](IoArgoprojWorkflowV1alpha1ArtifactLocation.md) |  | [optional] 
**AutomountServiceAccountToken** | Pointer to **bool** | AutomountServiceAccountToken indicates whether a service account token should be automatically mounted in pods. ServiceAccountName of ExecutorConfig must be specified if this value is false. | [optional] 
**Container** | Pointer to [**IoK8sApiCoreV1Container**](IoK8sApiCoreV1Container.md) |  | [optional] 
**ContainerSet** | Pointer to [**IoArgoprojWorkflowV1alpha1ContainerSetTemplate**](IoArgoprojWorkflowV1alpha1ContainerSetTemplate.md) |  | [optional] 
**Daemon** | Pointer to **bool** | Deamon will allow a workflow to proceed to the next step so long as the container reaches readiness | [optional] 
**Dag** | Pointer to [**IoArgoprojWorkflowV1alpha1DAGTemplate**](IoArgoprojWorkflowV1alpha1DAGTemplate.md) |  | [optional] 
**Data** | Pointer to [**IoArgoprojWorkflowV1alpha1Data**](IoArgoprojWorkflowV1alpha1Data.md) |  | [optional] 
**Executor** | Pointer to [**IoArgoprojWorkflowV1alpha1ExecutorConfig**](IoArgoprojWorkflowV1alpha1ExecutorConfig.md) |  | [optional] 
**FailFast** | Pointer to **bool** | FailFast, if specified, will fail this template if any of its child pods has failed. This is useful for when this template is expanded with &#x60;withItems&#x60;, etc. | [optional] 
**HostAliases** | Pointer to [**[]IoK8sApiCoreV1HostAlias**](IoK8sApiCoreV1HostAlias.md) | HostAliases is an optional list of hosts and IPs that will be injected into the pod spec | [optional] 
**Http** | Pointer to [**IoArgoprojWorkflowV1alpha1HTTP**](IoArgoprojWorkflowV1alpha1HTTP.md) |  | [optional] 
**InitContainers** | Pointer to [**[]IoArgoprojWorkflowV1alpha1UserContainer**](IoArgoprojWorkflowV1alpha1UserContainer.md) | InitContainers is a list of containers which run before the main container. | [optional] 
**Inputs** | Pointer to [**IoArgoprojWorkflowV1alpha1Inputs**](IoArgoprojWorkflowV1alpha1Inputs.md) |  | [optional] 
**Memoize** | Pointer to [**IoArgoprojWorkflowV1alpha1Memoize**](IoArgoprojWorkflowV1alpha1Memoize.md) |  | [optional] 
**Metadata** | Pointer to [**IoArgoprojWorkflowV1alpha1Metadata**](IoArgoprojWorkflowV1alpha1Metadata.md) |  | [optional] 
**Metrics** | Pointer to [**IoArgoprojWorkflowV1alpha1Metrics**](IoArgoprojWorkflowV1alpha1Metrics.md) |  | [optional] 
**Name** | Pointer to **string** | Name is the name of the template | [optional] 
**NodeSelector** | Pointer to **map[string]string** | NodeSelector is a selector to schedule this step of the workflow to be run on the selected node(s). Overrides the selector set at the workflow level. | [optional] 
**Outputs** | Pointer to [**IoArgoprojWorkflowV1alpha1Outputs**](IoArgoprojWorkflowV1alpha1Outputs.md) |  | [optional] 
**Parallelism** | Pointer to **int32** | Parallelism limits the max total parallel pods that can execute at the same time within the boundaries of this template invocation. If additional steps/dag templates are invoked, the pods created by those templates will not be counted towards this total. | [optional] 
**PodSpecPatch** | Pointer to **string** | PodSpecPatch holds strategic merge patch to apply against the pod spec. Allows parameterization of container fields which are not strings (e.g. resource limits). | [optional] 
**Priority** | Pointer to **int32** | Priority to apply to workflow pods. | [optional] 
**PriorityClassName** | Pointer to **string** | PriorityClassName to apply to workflow pods. | [optional] 
**Resource** | Pointer to [**IoArgoprojWorkflowV1alpha1ResourceTemplate**](IoArgoprojWorkflowV1alpha1ResourceTemplate.md) |  | [optional] 
**RetryStrategy** | Pointer to [**IoArgoprojWorkflowV1alpha1RetryStrategy**](IoArgoprojWorkflowV1alpha1RetryStrategy.md) |  | [optional] 
**SchedulerName** | Pointer to **string** | If specified, the pod will be dispatched by specified scheduler. Or it will be dispatched by workflow scope scheduler if specified. If neither specified, the pod will be dispatched by default scheduler. | [optional] 
**Script** | Pointer to [**IoArgoprojWorkflowV1alpha1ScriptTemplate**](IoArgoprojWorkflowV1alpha1ScriptTemplate.md) |  | [optional] 
**SecurityContext** | Pointer to [**IoK8sApiCoreV1PodSecurityContext**](IoK8sApiCoreV1PodSecurityContext.md) |  | [optional] 
**ServiceAccountName** | Pointer to **string** | ServiceAccountName to apply to workflow pods | [optional] 
**Sidecars** | Pointer to [**[]IoArgoprojWorkflowV1alpha1UserContainer**](IoArgoprojWorkflowV1alpha1UserContainer.md) | Sidecars is a list of containers which run alongside the main container Sidecars are automatically killed when the main container completes | [optional] 
**Steps** | Pointer to [**[][]IoArgoprojWorkflowV1alpha1WorkflowStep**]([]IoArgoprojWorkflowV1alpha1WorkflowStep.md) | Steps define a series of sequential/parallel workflow steps | [optional] 
**Suspend** | Pointer to [**IoArgoprojWorkflowV1alpha1SuspendTemplate**](IoArgoprojWorkflowV1alpha1SuspendTemplate.md) |  | [optional] 
**Synchronization** | Pointer to [**IoArgoprojWorkflowV1alpha1Synchronization**](IoArgoprojWorkflowV1alpha1Synchronization.md) |  | [optional] 
**Timeout** | Pointer to **string** | Timout allows to set the total node execution timeout duration counting from the node&#39;s start time. This duration also includes time in which the node spends in Pending state. This duration may not be applied to Step or DAG templates. | [optional] 
**Tolerations** | Pointer to [**[]IoK8sApiCoreV1Toleration**](IoK8sApiCoreV1Toleration.md) | Tolerations to apply to workflow pods. | [optional] 
**Volumes** | Pointer to [**[]IoK8sApiCoreV1Volume**](IoK8sApiCoreV1Volume.md) | Volumes is a list of volumes that can be mounted by containers in a template. | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1Template

`func NewIoArgoprojWorkflowV1alpha1Template() *IoArgoprojWorkflowV1alpha1Template`

NewIoArgoprojWorkflowV1alpha1Template instantiates a new IoArgoprojWorkflowV1alpha1Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1TemplateWithDefaults

`func NewIoArgoprojWorkflowV1alpha1TemplateWithDefaults() *IoArgoprojWorkflowV1alpha1Template`

NewIoArgoprojWorkflowV1alpha1TemplateWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDeadlineSeconds

`func (o *IoArgoprojWorkflowV1alpha1Template) GetActiveDeadlineSeconds() string`

GetActiveDeadlineSeconds returns the ActiveDeadlineSeconds field if non-nil, zero value otherwise.

### GetActiveDeadlineSecondsOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetActiveDeadlineSecondsOk() (*string, bool)`

GetActiveDeadlineSecondsOk returns a tuple with the ActiveDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeadlineSeconds

`func (o *IoArgoprojWorkflowV1alpha1Template) SetActiveDeadlineSeconds(v string)`

SetActiveDeadlineSeconds sets ActiveDeadlineSeconds field to given value.

### HasActiveDeadlineSeconds

`func (o *IoArgoprojWorkflowV1alpha1Template) HasActiveDeadlineSeconds() bool`

HasActiveDeadlineSeconds returns a boolean if a field has been set.

### GetAffinity

`func (o *IoArgoprojWorkflowV1alpha1Template) GetAffinity() IoK8sApiCoreV1Affinity`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetAffinityOk() (*IoK8sApiCoreV1Affinity, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *IoArgoprojWorkflowV1alpha1Template) SetAffinity(v IoK8sApiCoreV1Affinity)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *IoArgoprojWorkflowV1alpha1Template) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetArchiveLocation

`func (o *IoArgoprojWorkflowV1alpha1Template) GetArchiveLocation() IoArgoprojWorkflowV1alpha1ArtifactLocation`

GetArchiveLocation returns the ArchiveLocation field if non-nil, zero value otherwise.

### GetArchiveLocationOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetArchiveLocationOk() (*IoArgoprojWorkflowV1alpha1ArtifactLocation, bool)`

GetArchiveLocationOk returns a tuple with the ArchiveLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLocation

`func (o *IoArgoprojWorkflowV1alpha1Template) SetArchiveLocation(v IoArgoprojWorkflowV1alpha1ArtifactLocation)`

SetArchiveLocation sets ArchiveLocation field to given value.

### HasArchiveLocation

`func (o *IoArgoprojWorkflowV1alpha1Template) HasArchiveLocation() bool`

HasArchiveLocation returns a boolean if a field has been set.

### GetAutomountServiceAccountToken

`func (o *IoArgoprojWorkflowV1alpha1Template) GetAutomountServiceAccountToken() bool`

GetAutomountServiceAccountToken returns the AutomountServiceAccountToken field if non-nil, zero value otherwise.

### GetAutomountServiceAccountTokenOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetAutomountServiceAccountTokenOk() (*bool, bool)`

GetAutomountServiceAccountTokenOk returns a tuple with the AutomountServiceAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomountServiceAccountToken

`func (o *IoArgoprojWorkflowV1alpha1Template) SetAutomountServiceAccountToken(v bool)`

SetAutomountServiceAccountToken sets AutomountServiceAccountToken field to given value.

### HasAutomountServiceAccountToken

`func (o *IoArgoprojWorkflowV1alpha1Template) HasAutomountServiceAccountToken() bool`

HasAutomountServiceAccountToken returns a boolean if a field has been set.

### GetContainer

`func (o *IoArgoprojWorkflowV1alpha1Template) GetContainer() IoK8sApiCoreV1Container`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetContainerOk() (*IoK8sApiCoreV1Container, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *IoArgoprojWorkflowV1alpha1Template) SetContainer(v IoK8sApiCoreV1Container)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *IoArgoprojWorkflowV1alpha1Template) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContainerSet

`func (o *IoArgoprojWorkflowV1alpha1Template) GetContainerSet() IoArgoprojWorkflowV1alpha1ContainerSetTemplate`

GetContainerSet returns the ContainerSet field if non-nil, zero value otherwise.

### GetContainerSetOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetContainerSetOk() (*IoArgoprojWorkflowV1alpha1ContainerSetTemplate, bool)`

GetContainerSetOk returns a tuple with the ContainerSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSet

`func (o *IoArgoprojWorkflowV1alpha1Template) SetContainerSet(v IoArgoprojWorkflowV1alpha1ContainerSetTemplate)`

SetContainerSet sets ContainerSet field to given value.

### HasContainerSet

`func (o *IoArgoprojWorkflowV1alpha1Template) HasContainerSet() bool`

HasContainerSet returns a boolean if a field has been set.

### GetDaemon

`func (o *IoArgoprojWorkflowV1alpha1Template) GetDaemon() bool`

GetDaemon returns the Daemon field if non-nil, zero value otherwise.

### GetDaemonOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetDaemonOk() (*bool, bool)`

GetDaemonOk returns a tuple with the Daemon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemon

`func (o *IoArgoprojWorkflowV1alpha1Template) SetDaemon(v bool)`

SetDaemon sets Daemon field to given value.

### HasDaemon

`func (o *IoArgoprojWorkflowV1alpha1Template) HasDaemon() bool`

HasDaemon returns a boolean if a field has been set.

### GetDag

`func (o *IoArgoprojWorkflowV1alpha1Template) GetDag() IoArgoprojWorkflowV1alpha1DAGTemplate`

GetDag returns the Dag field if non-nil, zero value otherwise.

### GetDagOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetDagOk() (*IoArgoprojWorkflowV1alpha1DAGTemplate, bool)`

GetDagOk returns a tuple with the Dag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDag

`func (o *IoArgoprojWorkflowV1alpha1Template) SetDag(v IoArgoprojWorkflowV1alpha1DAGTemplate)`

SetDag sets Dag field to given value.

### HasDag

`func (o *IoArgoprojWorkflowV1alpha1Template) HasDag() bool`

HasDag returns a boolean if a field has been set.

### GetData

`func (o *IoArgoprojWorkflowV1alpha1Template) GetData() IoArgoprojWorkflowV1alpha1Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetDataOk() (*IoArgoprojWorkflowV1alpha1Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IoArgoprojWorkflowV1alpha1Template) SetData(v IoArgoprojWorkflowV1alpha1Data)`

SetData sets Data field to given value.

### HasData

`func (o *IoArgoprojWorkflowV1alpha1Template) HasData() bool`

HasData returns a boolean if a field has been set.

### GetExecutor

`func (o *IoArgoprojWorkflowV1alpha1Template) GetExecutor() IoArgoprojWorkflowV1alpha1ExecutorConfig`

GetExecutor returns the Executor field if non-nil, zero value otherwise.

### GetExecutorOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetExecutorOk() (*IoArgoprojWorkflowV1alpha1ExecutorConfig, bool)`

GetExecutorOk returns a tuple with the Executor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutor

`func (o *IoArgoprojWorkflowV1alpha1Template) SetExecutor(v IoArgoprojWorkflowV1alpha1ExecutorConfig)`

SetExecutor sets Executor field to given value.

### HasExecutor

`func (o *IoArgoprojWorkflowV1alpha1Template) HasExecutor() bool`

HasExecutor returns a boolean if a field has been set.

### GetFailFast

`func (o *IoArgoprojWorkflowV1alpha1Template) GetFailFast() bool`

GetFailFast returns the FailFast field if non-nil, zero value otherwise.

### GetFailFastOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetFailFastOk() (*bool, bool)`

GetFailFastOk returns a tuple with the FailFast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailFast

`func (o *IoArgoprojWorkflowV1alpha1Template) SetFailFast(v bool)`

SetFailFast sets FailFast field to given value.

### HasFailFast

`func (o *IoArgoprojWorkflowV1alpha1Template) HasFailFast() bool`

HasFailFast returns a boolean if a field has been set.

### GetHostAliases

`func (o *IoArgoprojWorkflowV1alpha1Template) GetHostAliases() []IoK8sApiCoreV1HostAlias`

GetHostAliases returns the HostAliases field if non-nil, zero value otherwise.

### GetHostAliasesOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetHostAliasesOk() (*[]IoK8sApiCoreV1HostAlias, bool)`

GetHostAliasesOk returns a tuple with the HostAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAliases

`func (o *IoArgoprojWorkflowV1alpha1Template) SetHostAliases(v []IoK8sApiCoreV1HostAlias)`

SetHostAliases sets HostAliases field to given value.

### HasHostAliases

`func (o *IoArgoprojWorkflowV1alpha1Template) HasHostAliases() bool`

HasHostAliases returns a boolean if a field has been set.

### GetHttp

`func (o *IoArgoprojWorkflowV1alpha1Template) GetHttp() IoArgoprojWorkflowV1alpha1HTTP`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetHttpOk() (*IoArgoprojWorkflowV1alpha1HTTP, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *IoArgoprojWorkflowV1alpha1Template) SetHttp(v IoArgoprojWorkflowV1alpha1HTTP)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *IoArgoprojWorkflowV1alpha1Template) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetInitContainers

`func (o *IoArgoprojWorkflowV1alpha1Template) GetInitContainers() []IoArgoprojWorkflowV1alpha1UserContainer`

GetInitContainers returns the InitContainers field if non-nil, zero value otherwise.

### GetInitContainersOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetInitContainersOk() (*[]IoArgoprojWorkflowV1alpha1UserContainer, bool)`

GetInitContainersOk returns a tuple with the InitContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitContainers

`func (o *IoArgoprojWorkflowV1alpha1Template) SetInitContainers(v []IoArgoprojWorkflowV1alpha1UserContainer)`

SetInitContainers sets InitContainers field to given value.

### HasInitContainers

`func (o *IoArgoprojWorkflowV1alpha1Template) HasInitContainers() bool`

HasInitContainers returns a boolean if a field has been set.

### GetInputs

`func (o *IoArgoprojWorkflowV1alpha1Template) GetInputs() IoArgoprojWorkflowV1alpha1Inputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetInputsOk() (*IoArgoprojWorkflowV1alpha1Inputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *IoArgoprojWorkflowV1alpha1Template) SetInputs(v IoArgoprojWorkflowV1alpha1Inputs)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *IoArgoprojWorkflowV1alpha1Template) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetMemoize

`func (o *IoArgoprojWorkflowV1alpha1Template) GetMemoize() IoArgoprojWorkflowV1alpha1Memoize`

GetMemoize returns the Memoize field if non-nil, zero value otherwise.

### GetMemoizeOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetMemoizeOk() (*IoArgoprojWorkflowV1alpha1Memoize, bool)`

GetMemoizeOk returns a tuple with the Memoize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoize

`func (o *IoArgoprojWorkflowV1alpha1Template) SetMemoize(v IoArgoprojWorkflowV1alpha1Memoize)`

SetMemoize sets Memoize field to given value.

### HasMemoize

`func (o *IoArgoprojWorkflowV1alpha1Template) HasMemoize() bool`

HasMemoize returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojWorkflowV1alpha1Template) GetMetadata() IoArgoprojWorkflowV1alpha1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetMetadataOk() (*IoArgoprojWorkflowV1alpha1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojWorkflowV1alpha1Template) SetMetadata(v IoArgoprojWorkflowV1alpha1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojWorkflowV1alpha1Template) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMetrics

`func (o *IoArgoprojWorkflowV1alpha1Template) GetMetrics() IoArgoprojWorkflowV1alpha1Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetMetricsOk() (*IoArgoprojWorkflowV1alpha1Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *IoArgoprojWorkflowV1alpha1Template) SetMetrics(v IoArgoprojWorkflowV1alpha1Metrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *IoArgoprojWorkflowV1alpha1Template) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetName

`func (o *IoArgoprojWorkflowV1alpha1Template) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojWorkflowV1alpha1Template) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoArgoprojWorkflowV1alpha1Template) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeSelector

`func (o *IoArgoprojWorkflowV1alpha1Template) GetNodeSelector() map[string]string`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetNodeSelectorOk() (*map[string]string, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *IoArgoprojWorkflowV1alpha1Template) SetNodeSelector(v map[string]string)`

SetNodeSelector sets NodeSelector field to given value.

### HasNodeSelector

`func (o *IoArgoprojWorkflowV1alpha1Template) HasNodeSelector() bool`

HasNodeSelector returns a boolean if a field has been set.

### GetOutputs

`func (o *IoArgoprojWorkflowV1alpha1Template) GetOutputs() IoArgoprojWorkflowV1alpha1Outputs`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetOutputsOk() (*IoArgoprojWorkflowV1alpha1Outputs, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *IoArgoprojWorkflowV1alpha1Template) SetOutputs(v IoArgoprojWorkflowV1alpha1Outputs)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *IoArgoprojWorkflowV1alpha1Template) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetParallelism

`func (o *IoArgoprojWorkflowV1alpha1Template) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *IoArgoprojWorkflowV1alpha1Template) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *IoArgoprojWorkflowV1alpha1Template) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetPodSpecPatch

`func (o *IoArgoprojWorkflowV1alpha1Template) GetPodSpecPatch() string`

GetPodSpecPatch returns the PodSpecPatch field if non-nil, zero value otherwise.

### GetPodSpecPatchOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetPodSpecPatchOk() (*string, bool)`

GetPodSpecPatchOk returns a tuple with the PodSpecPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodSpecPatch

`func (o *IoArgoprojWorkflowV1alpha1Template) SetPodSpecPatch(v string)`

SetPodSpecPatch sets PodSpecPatch field to given value.

### HasPodSpecPatch

`func (o *IoArgoprojWorkflowV1alpha1Template) HasPodSpecPatch() bool`

HasPodSpecPatch returns a boolean if a field has been set.

### GetPriority

`func (o *IoArgoprojWorkflowV1alpha1Template) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IoArgoprojWorkflowV1alpha1Template) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IoArgoprojWorkflowV1alpha1Template) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPriorityClassName

`func (o *IoArgoprojWorkflowV1alpha1Template) GetPriorityClassName() string`

GetPriorityClassName returns the PriorityClassName field if non-nil, zero value otherwise.

### GetPriorityClassNameOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetPriorityClassNameOk() (*string, bool)`

GetPriorityClassNameOk returns a tuple with the PriorityClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityClassName

`func (o *IoArgoprojWorkflowV1alpha1Template) SetPriorityClassName(v string)`

SetPriorityClassName sets PriorityClassName field to given value.

### HasPriorityClassName

`func (o *IoArgoprojWorkflowV1alpha1Template) HasPriorityClassName() bool`

HasPriorityClassName returns a boolean if a field has been set.

### GetResource

`func (o *IoArgoprojWorkflowV1alpha1Template) GetResource() IoArgoprojWorkflowV1alpha1ResourceTemplate`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetResourceOk() (*IoArgoprojWorkflowV1alpha1ResourceTemplate, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IoArgoprojWorkflowV1alpha1Template) SetResource(v IoArgoprojWorkflowV1alpha1ResourceTemplate)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IoArgoprojWorkflowV1alpha1Template) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRetryStrategy

`func (o *IoArgoprojWorkflowV1alpha1Template) GetRetryStrategy() IoArgoprojWorkflowV1alpha1RetryStrategy`

GetRetryStrategy returns the RetryStrategy field if non-nil, zero value otherwise.

### GetRetryStrategyOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetRetryStrategyOk() (*IoArgoprojWorkflowV1alpha1RetryStrategy, bool)`

GetRetryStrategyOk returns a tuple with the RetryStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryStrategy

`func (o *IoArgoprojWorkflowV1alpha1Template) SetRetryStrategy(v IoArgoprojWorkflowV1alpha1RetryStrategy)`

SetRetryStrategy sets RetryStrategy field to given value.

### HasRetryStrategy

`func (o *IoArgoprojWorkflowV1alpha1Template) HasRetryStrategy() bool`

HasRetryStrategy returns a boolean if a field has been set.

### GetSchedulerName

`func (o *IoArgoprojWorkflowV1alpha1Template) GetSchedulerName() string`

GetSchedulerName returns the SchedulerName field if non-nil, zero value otherwise.

### GetSchedulerNameOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetSchedulerNameOk() (*string, bool)`

GetSchedulerNameOk returns a tuple with the SchedulerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerName

`func (o *IoArgoprojWorkflowV1alpha1Template) SetSchedulerName(v string)`

SetSchedulerName sets SchedulerName field to given value.

### HasSchedulerName

`func (o *IoArgoprojWorkflowV1alpha1Template) HasSchedulerName() bool`

HasSchedulerName returns a boolean if a field has been set.

### GetScript

`func (o *IoArgoprojWorkflowV1alpha1Template) GetScript() IoArgoprojWorkflowV1alpha1ScriptTemplate`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetScriptOk() (*IoArgoprojWorkflowV1alpha1ScriptTemplate, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *IoArgoprojWorkflowV1alpha1Template) SetScript(v IoArgoprojWorkflowV1alpha1ScriptTemplate)`

SetScript sets Script field to given value.

### HasScript

`func (o *IoArgoprojWorkflowV1alpha1Template) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetSecurityContext

`func (o *IoArgoprojWorkflowV1alpha1Template) GetSecurityContext() IoK8sApiCoreV1PodSecurityContext`

GetSecurityContext returns the SecurityContext field if non-nil, zero value otherwise.

### GetSecurityContextOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetSecurityContextOk() (*IoK8sApiCoreV1PodSecurityContext, bool)`

GetSecurityContextOk returns a tuple with the SecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityContext

`func (o *IoArgoprojWorkflowV1alpha1Template) SetSecurityContext(v IoK8sApiCoreV1PodSecurityContext)`

SetSecurityContext sets SecurityContext field to given value.

### HasSecurityContext

`func (o *IoArgoprojWorkflowV1alpha1Template) HasSecurityContext() bool`

HasSecurityContext returns a boolean if a field has been set.

### GetServiceAccountName

`func (o *IoArgoprojWorkflowV1alpha1Template) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *IoArgoprojWorkflowV1alpha1Template) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *IoArgoprojWorkflowV1alpha1Template) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.

### GetSidecars

`func (o *IoArgoprojWorkflowV1alpha1Template) GetSidecars() []IoArgoprojWorkflowV1alpha1UserContainer`

GetSidecars returns the Sidecars field if non-nil, zero value otherwise.

### GetSidecarsOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetSidecarsOk() (*[]IoArgoprojWorkflowV1alpha1UserContainer, bool)`

GetSidecarsOk returns a tuple with the Sidecars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSidecars

`func (o *IoArgoprojWorkflowV1alpha1Template) SetSidecars(v []IoArgoprojWorkflowV1alpha1UserContainer)`

SetSidecars sets Sidecars field to given value.

### HasSidecars

`func (o *IoArgoprojWorkflowV1alpha1Template) HasSidecars() bool`

HasSidecars returns a boolean if a field has been set.

### GetSteps

`func (o *IoArgoprojWorkflowV1alpha1Template) GetSteps() [][]IoArgoprojWorkflowV1alpha1WorkflowStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetStepsOk() (*[][]IoArgoprojWorkflowV1alpha1WorkflowStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *IoArgoprojWorkflowV1alpha1Template) SetSteps(v [][]IoArgoprojWorkflowV1alpha1WorkflowStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *IoArgoprojWorkflowV1alpha1Template) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetSuspend

`func (o *IoArgoprojWorkflowV1alpha1Template) GetSuspend() IoArgoprojWorkflowV1alpha1SuspendTemplate`

GetSuspend returns the Suspend field if non-nil, zero value otherwise.

### GetSuspendOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetSuspendOk() (*IoArgoprojWorkflowV1alpha1SuspendTemplate, bool)`

GetSuspendOk returns a tuple with the Suspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspend

`func (o *IoArgoprojWorkflowV1alpha1Template) SetSuspend(v IoArgoprojWorkflowV1alpha1SuspendTemplate)`

SetSuspend sets Suspend field to given value.

### HasSuspend

`func (o *IoArgoprojWorkflowV1alpha1Template) HasSuspend() bool`

HasSuspend returns a boolean if a field has been set.

### GetSynchronization

`func (o *IoArgoprojWorkflowV1alpha1Template) GetSynchronization() IoArgoprojWorkflowV1alpha1Synchronization`

GetSynchronization returns the Synchronization field if non-nil, zero value otherwise.

### GetSynchronizationOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetSynchronizationOk() (*IoArgoprojWorkflowV1alpha1Synchronization, bool)`

GetSynchronizationOk returns a tuple with the Synchronization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronization

`func (o *IoArgoprojWorkflowV1alpha1Template) SetSynchronization(v IoArgoprojWorkflowV1alpha1Synchronization)`

SetSynchronization sets Synchronization field to given value.

### HasSynchronization

`func (o *IoArgoprojWorkflowV1alpha1Template) HasSynchronization() bool`

HasSynchronization returns a boolean if a field has been set.

### GetTimeout

`func (o *IoArgoprojWorkflowV1alpha1Template) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *IoArgoprojWorkflowV1alpha1Template) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *IoArgoprojWorkflowV1alpha1Template) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTolerations

`func (o *IoArgoprojWorkflowV1alpha1Template) GetTolerations() []IoK8sApiCoreV1Toleration`

GetTolerations returns the Tolerations field if non-nil, zero value otherwise.

### GetTolerationsOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetTolerationsOk() (*[]IoK8sApiCoreV1Toleration, bool)`

GetTolerationsOk returns a tuple with the Tolerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerations

`func (o *IoArgoprojWorkflowV1alpha1Template) SetTolerations(v []IoK8sApiCoreV1Toleration)`

SetTolerations sets Tolerations field to given value.

### HasTolerations

`func (o *IoArgoprojWorkflowV1alpha1Template) HasTolerations() bool`

HasTolerations returns a boolean if a field has been set.

### GetVolumes

`func (o *IoArgoprojWorkflowV1alpha1Template) GetVolumes() []IoK8sApiCoreV1Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *IoArgoprojWorkflowV1alpha1Template) GetVolumesOk() (*[]IoK8sApiCoreV1Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *IoArgoprojWorkflowV1alpha1Template) SetVolumes(v []IoK8sApiCoreV1Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *IoArgoprojWorkflowV1alpha1Template) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


