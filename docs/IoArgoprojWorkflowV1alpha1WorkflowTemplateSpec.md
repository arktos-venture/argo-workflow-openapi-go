# IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDeadlineSeconds** | Pointer to **int32** | Optional duration in seconds relative to the workflow start time which the workflow is allowed to run before the controller terminates the io.argoproj.workflow.v1alpha1. A value of zero is used to terminate a Running workflow | [optional] 
**Affinity** | Pointer to [**IoK8sApiCoreV1Affinity**](IoK8sApiCoreV1Affinity.md) |  | [optional] 
**ArchiveLogs** | Pointer to **bool** | ArchiveLogs indicates if the container logs should be archived | [optional] 
**Arguments** | Pointer to [**IoArgoprojWorkflowV1alpha1Arguments**](IoArgoprojWorkflowV1alpha1Arguments.md) |  | [optional] 
**ArtifactRepositoryRef** | Pointer to [**IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef**](IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef.md) |  | [optional] 
**AutomountServiceAccountToken** | Pointer to **bool** | AutomountServiceAccountToken indicates whether a service account token should be automatically mounted in pods. ServiceAccountName of ExecutorConfig must be specified if this value is false. | [optional] 
**DnsConfig** | Pointer to [**IoK8sApiCoreV1PodDNSConfig**](IoK8sApiCoreV1PodDNSConfig.md) |  | [optional] 
**DnsPolicy** | Pointer to **string** | Set DNS policy for the pod. Defaults to \&quot;ClusterFirst\&quot;. Valid values are &#39;ClusterFirstWithHostNet&#39;, &#39;ClusterFirst&#39;, &#39;Default&#39; or &#39;None&#39;. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to &#39;ClusterFirstWithHostNet&#39;. | [optional] 
**Entrypoint** | Pointer to **string** | Entrypoint is a template reference to the starting point of the io.argoproj.workflow.v1alpha1. | [optional] 
**Executor** | Pointer to [**IoArgoprojWorkflowV1alpha1ExecutorConfig**](IoArgoprojWorkflowV1alpha1ExecutorConfig.md) |  | [optional] 
**HostAliases** | Pointer to [**[]IoK8sApiCoreV1HostAlias**](IoK8sApiCoreV1HostAlias.md) |  | [optional] 
**HostNetwork** | Pointer to **bool** | Host networking requested for this workflow pod. Default to false. | [optional] 
**ImagePullSecrets** | Pointer to [**[]IoK8sApiCoreV1LocalObjectReference**](IoK8sApiCoreV1LocalObjectReference.md) | ImagePullSecrets is a list of references to secrets in the same namespace to use for pulling any images in pods that reference this ServiceAccount. ImagePullSecrets are distinct from Secrets because Secrets can be mounted in the pod, but ImagePullSecrets are only accessed by the kubelet. More info: https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod | [optional] 
**Metrics** | Pointer to [**IoArgoprojWorkflowV1alpha1Metrics**](IoArgoprojWorkflowV1alpha1Metrics.md) |  | [optional] 
**NodeSelector** | Pointer to **map[string]string** | NodeSelector is a selector which will result in all pods of the workflow to be scheduled on the selected node(s). This is able to be overridden by a nodeSelector specified in the template. | [optional] 
**OnExit** | Pointer to **string** | OnExit is a template reference which is invoked at the end of the workflow, irrespective of the success, failure, or error of the primary io.argoproj.workflow.v1alpha1. | [optional] 
**Parallelism** | Pointer to **int32** | Parallelism limits the max total parallel pods that can execute at the same time in a workflow | [optional] 
**PodDisruptionBudget** | Pointer to [**IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec**](IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec.md) |  | [optional] 
**PodGC** | Pointer to [**IoArgoprojWorkflowV1alpha1PodGC**](IoArgoprojWorkflowV1alpha1PodGC.md) |  | [optional] 
**PodMetadata** | Pointer to [**IoArgoprojWorkflowV1alpha1Metadata**](IoArgoprojWorkflowV1alpha1Metadata.md) |  | [optional] 
**PodPriority** | Pointer to **int32** | Priority to apply to workflow pods. | [optional] 
**PodPriorityClassName** | Pointer to **string** | PriorityClassName to apply to workflow pods. | [optional] 
**PodSpecPatch** | Pointer to **string** | PodSpecPatch holds strategic merge patch to apply against the pod spec. Allows parameterization of container fields which are not strings (e.g. resource limits). | [optional] 
**Priority** | Pointer to **int32** | Priority is used if controller is configured to process limited number of workflows in parallel. Workflows with higher priority are processed first. | [optional] 
**RetryStrategy** | Pointer to [**IoArgoprojWorkflowV1alpha1RetryStrategy**](IoArgoprojWorkflowV1alpha1RetryStrategy.md) |  | [optional] 
**SchedulerName** | Pointer to **string** | Set scheduler name for all pods. Will be overridden if container/script template&#39;s scheduler name is set. Default scheduler will be used if neither specified. | [optional] 
**SecurityContext** | Pointer to [**IoK8sApiCoreV1PodSecurityContext**](IoK8sApiCoreV1PodSecurityContext.md) |  | [optional] 
**ServiceAccountName** | Pointer to **string** | ServiceAccountName is the name of the ServiceAccount to run all pods of the workflow as. | [optional] 
**Shutdown** | Pointer to **string** | Shutdown will shutdown the workflow according to its ShutdownStrategy | [optional] 
**Suspend** | Pointer to **bool** | Suspend will suspend the workflow and prevent execution of any future steps in the workflow | [optional] 
**Synchronization** | Pointer to [**IoArgoprojWorkflowV1alpha1Synchronization**](IoArgoprojWorkflowV1alpha1Synchronization.md) |  | [optional] 
**TemplateDefaults** | Pointer to [**IoArgoprojWorkflowV1alpha1Template**](IoArgoprojWorkflowV1alpha1Template.md) |  | [optional] 
**Templates** | Pointer to [**[]IoArgoprojWorkflowV1alpha1Template**](IoArgoprojWorkflowV1alpha1Template.md) | Templates is a list of workflow templates used in a workflow | [optional] 
**Tolerations** | Pointer to [**[]IoK8sApiCoreV1Toleration**](IoK8sApiCoreV1Toleration.md) | Tolerations to apply to workflow pods. | [optional] 
**TtlStrategy** | Pointer to [**IoArgoprojWorkflowV1alpha1TTLStrategy**](IoArgoprojWorkflowV1alpha1TTLStrategy.md) |  | [optional] 
**VolumeClaimGC** | Pointer to [**IoArgoprojWorkflowV1alpha1VolumeClaimGC**](IoArgoprojWorkflowV1alpha1VolumeClaimGC.md) |  | [optional] 
**VolumeClaimTemplates** | Pointer to [**[]IoK8sApiCoreV1PersistentVolumeClaim**](IoK8sApiCoreV1PersistentVolumeClaim.md) | VolumeClaimTemplates is a list of claims that containers are allowed to reference. The Workflow controller will create the claims at the beginning of the workflow and delete the claims upon completion of the workflow | [optional] 
**Volumes** | Pointer to [**[]IoK8sApiCoreV1Volume**](IoK8sApiCoreV1Volume.md) | Volumes is a list of volumes that can be mounted by containers in a io.argoproj.workflow.v1alpha1. | [optional] 
**WorkflowMetadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**WorkflowTemplateRef** | Pointer to [**IoArgoprojWorkflowV1alpha1WorkflowTemplateRef**](IoArgoprojWorkflowV1alpha1WorkflowTemplateRef.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1WorkflowTemplateSpec

`func NewIoArgoprojWorkflowV1alpha1WorkflowTemplateSpec() *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec`

NewIoArgoprojWorkflowV1alpha1WorkflowTemplateSpec instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1WorkflowTemplateSpecWithDefaults

`func NewIoArgoprojWorkflowV1alpha1WorkflowTemplateSpecWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec`

NewIoArgoprojWorkflowV1alpha1WorkflowTemplateSpecWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDeadlineSeconds

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetActiveDeadlineSeconds() int32`

GetActiveDeadlineSeconds returns the ActiveDeadlineSeconds field if non-nil, zero value otherwise.

### GetActiveDeadlineSecondsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetActiveDeadlineSecondsOk() (*int32, bool)`

GetActiveDeadlineSecondsOk returns a tuple with the ActiveDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDeadlineSeconds

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetActiveDeadlineSeconds(v int32)`

SetActiveDeadlineSeconds sets ActiveDeadlineSeconds field to given value.

### HasActiveDeadlineSeconds

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasActiveDeadlineSeconds() bool`

HasActiveDeadlineSeconds returns a boolean if a field has been set.

### GetAffinity

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetAffinity() IoK8sApiCoreV1Affinity`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetAffinityOk() (*IoK8sApiCoreV1Affinity, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetAffinity(v IoK8sApiCoreV1Affinity)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetArchiveLogs() bool`

GetArchiveLogs returns the ArchiveLogs field if non-nil, zero value otherwise.

### GetArchiveLogsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetArchiveLogsOk() (*bool, bool)`

GetArchiveLogsOk returns a tuple with the ArchiveLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetArchiveLogs(v bool)`

SetArchiveLogs sets ArchiveLogs field to given value.

### HasArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasArchiveLogs() bool`

HasArchiveLogs returns a boolean if a field has been set.

### GetArguments

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetArguments() IoArgoprojWorkflowV1alpha1Arguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetArgumentsOk() (*IoArgoprojWorkflowV1alpha1Arguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetArguments(v IoArgoprojWorkflowV1alpha1Arguments)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetArtifactRepositoryRef

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetArtifactRepositoryRef() IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef`

GetArtifactRepositoryRef returns the ArtifactRepositoryRef field if non-nil, zero value otherwise.

### GetArtifactRepositoryRefOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetArtifactRepositoryRefOk() (*IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef, bool)`

GetArtifactRepositoryRefOk returns a tuple with the ArtifactRepositoryRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactRepositoryRef

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetArtifactRepositoryRef(v IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef)`

SetArtifactRepositoryRef sets ArtifactRepositoryRef field to given value.

### HasArtifactRepositoryRef

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasArtifactRepositoryRef() bool`

HasArtifactRepositoryRef returns a boolean if a field has been set.

### GetAutomountServiceAccountToken

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetAutomountServiceAccountToken() bool`

GetAutomountServiceAccountToken returns the AutomountServiceAccountToken field if non-nil, zero value otherwise.

### GetAutomountServiceAccountTokenOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetAutomountServiceAccountTokenOk() (*bool, bool)`

GetAutomountServiceAccountTokenOk returns a tuple with the AutomountServiceAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomountServiceAccountToken

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetAutomountServiceAccountToken(v bool)`

SetAutomountServiceAccountToken sets AutomountServiceAccountToken field to given value.

### HasAutomountServiceAccountToken

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasAutomountServiceAccountToken() bool`

HasAutomountServiceAccountToken returns a boolean if a field has been set.

### GetDnsConfig

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetDnsConfig() IoK8sApiCoreV1PodDNSConfig`

GetDnsConfig returns the DnsConfig field if non-nil, zero value otherwise.

### GetDnsConfigOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetDnsConfigOk() (*IoK8sApiCoreV1PodDNSConfig, bool)`

GetDnsConfigOk returns a tuple with the DnsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfig

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetDnsConfig(v IoK8sApiCoreV1PodDNSConfig)`

SetDnsConfig sets DnsConfig field to given value.

### HasDnsConfig

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasDnsConfig() bool`

HasDnsConfig returns a boolean if a field has been set.

### GetDnsPolicy

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetDnsPolicy() string`

GetDnsPolicy returns the DnsPolicy field if non-nil, zero value otherwise.

### GetDnsPolicyOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetDnsPolicyOk() (*string, bool)`

GetDnsPolicyOk returns a tuple with the DnsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPolicy

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetDnsPolicy(v string)`

SetDnsPolicy sets DnsPolicy field to given value.

### HasDnsPolicy

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasDnsPolicy() bool`

HasDnsPolicy returns a boolean if a field has been set.

### GetEntrypoint

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetExecutor

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetExecutor() IoArgoprojWorkflowV1alpha1ExecutorConfig`

GetExecutor returns the Executor field if non-nil, zero value otherwise.

### GetExecutorOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetExecutorOk() (*IoArgoprojWorkflowV1alpha1ExecutorConfig, bool)`

GetExecutorOk returns a tuple with the Executor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutor

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetExecutor(v IoArgoprojWorkflowV1alpha1ExecutorConfig)`

SetExecutor sets Executor field to given value.

### HasExecutor

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasExecutor() bool`

HasExecutor returns a boolean if a field has been set.

### GetHostAliases

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetHostAliases() []IoK8sApiCoreV1HostAlias`

GetHostAliases returns the HostAliases field if non-nil, zero value otherwise.

### GetHostAliasesOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetHostAliasesOk() (*[]IoK8sApiCoreV1HostAlias, bool)`

GetHostAliasesOk returns a tuple with the HostAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAliases

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetHostAliases(v []IoK8sApiCoreV1HostAlias)`

SetHostAliases sets HostAliases field to given value.

### HasHostAliases

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasHostAliases() bool`

HasHostAliases returns a boolean if a field has been set.

### GetHostNetwork

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetHostNetwork() bool`

GetHostNetwork returns the HostNetwork field if non-nil, zero value otherwise.

### GetHostNetworkOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetHostNetworkOk() (*bool, bool)`

GetHostNetworkOk returns a tuple with the HostNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNetwork

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetHostNetwork(v bool)`

SetHostNetwork sets HostNetwork field to given value.

### HasHostNetwork

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasHostNetwork() bool`

HasHostNetwork returns a boolean if a field has been set.

### GetImagePullSecrets

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetImagePullSecrets() []IoK8sApiCoreV1LocalObjectReference`

GetImagePullSecrets returns the ImagePullSecrets field if non-nil, zero value otherwise.

### GetImagePullSecretsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetImagePullSecretsOk() (*[]IoK8sApiCoreV1LocalObjectReference, bool)`

GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullSecrets

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetImagePullSecrets(v []IoK8sApiCoreV1LocalObjectReference)`

SetImagePullSecrets sets ImagePullSecrets field to given value.

### HasImagePullSecrets

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasImagePullSecrets() bool`

HasImagePullSecrets returns a boolean if a field has been set.

### GetMetrics

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetMetrics() IoArgoprojWorkflowV1alpha1Metrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetMetricsOk() (*IoArgoprojWorkflowV1alpha1Metrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetMetrics(v IoArgoprojWorkflowV1alpha1Metrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetNodeSelector

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetNodeSelector() map[string]string`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetNodeSelectorOk() (*map[string]string, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetNodeSelector(v map[string]string)`

SetNodeSelector sets NodeSelector field to given value.

### HasNodeSelector

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasNodeSelector() bool`

HasNodeSelector returns a boolean if a field has been set.

### GetOnExit

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetOnExit() string`

GetOnExit returns the OnExit field if non-nil, zero value otherwise.

### GetOnExitOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetOnExitOk() (*string, bool)`

GetOnExitOk returns a tuple with the OnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnExit

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetOnExit(v string)`

SetOnExit sets OnExit field to given value.

### HasOnExit

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasOnExit() bool`

HasOnExit returns a boolean if a field has been set.

### GetParallelism

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetPodDisruptionBudget

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPodDisruptionBudget() IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec`

GetPodDisruptionBudget returns the PodDisruptionBudget field if non-nil, zero value otherwise.

### GetPodDisruptionBudgetOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPodDisruptionBudgetOk() (*IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec, bool)`

GetPodDisruptionBudgetOk returns a tuple with the PodDisruptionBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodDisruptionBudget

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetPodDisruptionBudget(v IoK8sApiPolicyV1beta1PodDisruptionBudgetSpec)`

SetPodDisruptionBudget sets PodDisruptionBudget field to given value.

### HasPodDisruptionBudget

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasPodDisruptionBudget() bool`

HasPodDisruptionBudget returns a boolean if a field has been set.

### GetPodGC

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPodGC() IoArgoprojWorkflowV1alpha1PodGC`

GetPodGC returns the PodGC field if non-nil, zero value otherwise.

### GetPodGCOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPodGCOk() (*IoArgoprojWorkflowV1alpha1PodGC, bool)`

GetPodGCOk returns a tuple with the PodGC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodGC

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetPodGC(v IoArgoprojWorkflowV1alpha1PodGC)`

SetPodGC sets PodGC field to given value.

### HasPodGC

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasPodGC() bool`

HasPodGC returns a boolean if a field has been set.

### GetPodMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPodMetadata() IoArgoprojWorkflowV1alpha1Metadata`

GetPodMetadata returns the PodMetadata field if non-nil, zero value otherwise.

### GetPodMetadataOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPodMetadataOk() (*IoArgoprojWorkflowV1alpha1Metadata, bool)`

GetPodMetadataOk returns a tuple with the PodMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetPodMetadata(v IoArgoprojWorkflowV1alpha1Metadata)`

SetPodMetadata sets PodMetadata field to given value.

### HasPodMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasPodMetadata() bool`

HasPodMetadata returns a boolean if a field has been set.

### GetPodPriority

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPodPriority() int32`

GetPodPriority returns the PodPriority field if non-nil, zero value otherwise.

### GetPodPriorityOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPodPriorityOk() (*int32, bool)`

GetPodPriorityOk returns a tuple with the PodPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPriority

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetPodPriority(v int32)`

SetPodPriority sets PodPriority field to given value.

### HasPodPriority

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasPodPriority() bool`

HasPodPriority returns a boolean if a field has been set.

### GetPodPriorityClassName

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPodPriorityClassName() string`

GetPodPriorityClassName returns the PodPriorityClassName field if non-nil, zero value otherwise.

### GetPodPriorityClassNameOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPodPriorityClassNameOk() (*string, bool)`

GetPodPriorityClassNameOk returns a tuple with the PodPriorityClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPriorityClassName

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetPodPriorityClassName(v string)`

SetPodPriorityClassName sets PodPriorityClassName field to given value.

### HasPodPriorityClassName

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasPodPriorityClassName() bool`

HasPodPriorityClassName returns a boolean if a field has been set.

### GetPodSpecPatch

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPodSpecPatch() string`

GetPodSpecPatch returns the PodSpecPatch field if non-nil, zero value otherwise.

### GetPodSpecPatchOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPodSpecPatchOk() (*string, bool)`

GetPodSpecPatchOk returns a tuple with the PodSpecPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodSpecPatch

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetPodSpecPatch(v string)`

SetPodSpecPatch sets PodSpecPatch field to given value.

### HasPodSpecPatch

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasPodSpecPatch() bool`

HasPodSpecPatch returns a boolean if a field has been set.

### GetPriority

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRetryStrategy

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetRetryStrategy() IoArgoprojWorkflowV1alpha1RetryStrategy`

GetRetryStrategy returns the RetryStrategy field if non-nil, zero value otherwise.

### GetRetryStrategyOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetRetryStrategyOk() (*IoArgoprojWorkflowV1alpha1RetryStrategy, bool)`

GetRetryStrategyOk returns a tuple with the RetryStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryStrategy

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetRetryStrategy(v IoArgoprojWorkflowV1alpha1RetryStrategy)`

SetRetryStrategy sets RetryStrategy field to given value.

### HasRetryStrategy

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasRetryStrategy() bool`

HasRetryStrategy returns a boolean if a field has been set.

### GetSchedulerName

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetSchedulerName() string`

GetSchedulerName returns the SchedulerName field if non-nil, zero value otherwise.

### GetSchedulerNameOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetSchedulerNameOk() (*string, bool)`

GetSchedulerNameOk returns a tuple with the SchedulerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerName

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetSchedulerName(v string)`

SetSchedulerName sets SchedulerName field to given value.

### HasSchedulerName

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasSchedulerName() bool`

HasSchedulerName returns a boolean if a field has been set.

### GetSecurityContext

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetSecurityContext() IoK8sApiCoreV1PodSecurityContext`

GetSecurityContext returns the SecurityContext field if non-nil, zero value otherwise.

### GetSecurityContextOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetSecurityContextOk() (*IoK8sApiCoreV1PodSecurityContext, bool)`

GetSecurityContextOk returns a tuple with the SecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityContext

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetSecurityContext(v IoK8sApiCoreV1PodSecurityContext)`

SetSecurityContext sets SecurityContext field to given value.

### HasSecurityContext

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasSecurityContext() bool`

HasSecurityContext returns a boolean if a field has been set.

### GetServiceAccountName

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.

### GetShutdown

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetShutdown() string`

GetShutdown returns the Shutdown field if non-nil, zero value otherwise.

### GetShutdownOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetShutdownOk() (*string, bool)`

GetShutdownOk returns a tuple with the Shutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutdown

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetShutdown(v string)`

SetShutdown sets Shutdown field to given value.

### HasShutdown

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasShutdown() bool`

HasShutdown returns a boolean if a field has been set.

### GetSuspend

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetSuspend() bool`

GetSuspend returns the Suspend field if non-nil, zero value otherwise.

### GetSuspendOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetSuspendOk() (*bool, bool)`

GetSuspendOk returns a tuple with the Suspend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspend

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetSuspend(v bool)`

SetSuspend sets Suspend field to given value.

### HasSuspend

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasSuspend() bool`

HasSuspend returns a boolean if a field has been set.

### GetSynchronization

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetSynchronization() IoArgoprojWorkflowV1alpha1Synchronization`

GetSynchronization returns the Synchronization field if non-nil, zero value otherwise.

### GetSynchronizationOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetSynchronizationOk() (*IoArgoprojWorkflowV1alpha1Synchronization, bool)`

GetSynchronizationOk returns a tuple with the Synchronization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronization

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetSynchronization(v IoArgoprojWorkflowV1alpha1Synchronization)`

SetSynchronization sets Synchronization field to given value.

### HasSynchronization

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasSynchronization() bool`

HasSynchronization returns a boolean if a field has been set.

### GetTemplateDefaults

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetTemplateDefaults() IoArgoprojWorkflowV1alpha1Template`

GetTemplateDefaults returns the TemplateDefaults field if non-nil, zero value otherwise.

### GetTemplateDefaultsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetTemplateDefaultsOk() (*IoArgoprojWorkflowV1alpha1Template, bool)`

GetTemplateDefaultsOk returns a tuple with the TemplateDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDefaults

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetTemplateDefaults(v IoArgoprojWorkflowV1alpha1Template)`

SetTemplateDefaults sets TemplateDefaults field to given value.

### HasTemplateDefaults

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasTemplateDefaults() bool`

HasTemplateDefaults returns a boolean if a field has been set.

### GetTemplates

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetTemplates() []IoArgoprojWorkflowV1alpha1Template`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetTemplatesOk() (*[]IoArgoprojWorkflowV1alpha1Template, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetTemplates(v []IoArgoprojWorkflowV1alpha1Template)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetTolerations

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetTolerations() []IoK8sApiCoreV1Toleration`

GetTolerations returns the Tolerations field if non-nil, zero value otherwise.

### GetTolerationsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetTolerationsOk() (*[]IoK8sApiCoreV1Toleration, bool)`

GetTolerationsOk returns a tuple with the Tolerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerations

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetTolerations(v []IoK8sApiCoreV1Toleration)`

SetTolerations sets Tolerations field to given value.

### HasTolerations

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasTolerations() bool`

HasTolerations returns a boolean if a field has been set.

### GetTtlStrategy

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetTtlStrategy() IoArgoprojWorkflowV1alpha1TTLStrategy`

GetTtlStrategy returns the TtlStrategy field if non-nil, zero value otherwise.

### GetTtlStrategyOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetTtlStrategyOk() (*IoArgoprojWorkflowV1alpha1TTLStrategy, bool)`

GetTtlStrategyOk returns a tuple with the TtlStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlStrategy

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetTtlStrategy(v IoArgoprojWorkflowV1alpha1TTLStrategy)`

SetTtlStrategy sets TtlStrategy field to given value.

### HasTtlStrategy

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasTtlStrategy() bool`

HasTtlStrategy returns a boolean if a field has been set.

### GetVolumeClaimGC

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetVolumeClaimGC() IoArgoprojWorkflowV1alpha1VolumeClaimGC`

GetVolumeClaimGC returns the VolumeClaimGC field if non-nil, zero value otherwise.

### GetVolumeClaimGCOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetVolumeClaimGCOk() (*IoArgoprojWorkflowV1alpha1VolumeClaimGC, bool)`

GetVolumeClaimGCOk returns a tuple with the VolumeClaimGC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeClaimGC

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetVolumeClaimGC(v IoArgoprojWorkflowV1alpha1VolumeClaimGC)`

SetVolumeClaimGC sets VolumeClaimGC field to given value.

### HasVolumeClaimGC

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasVolumeClaimGC() bool`

HasVolumeClaimGC returns a boolean if a field has been set.

### GetVolumeClaimTemplates

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetVolumeClaimTemplates() []IoK8sApiCoreV1PersistentVolumeClaim`

GetVolumeClaimTemplates returns the VolumeClaimTemplates field if non-nil, zero value otherwise.

### GetVolumeClaimTemplatesOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetVolumeClaimTemplatesOk() (*[]IoK8sApiCoreV1PersistentVolumeClaim, bool)`

GetVolumeClaimTemplatesOk returns a tuple with the VolumeClaimTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeClaimTemplates

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetVolumeClaimTemplates(v []IoK8sApiCoreV1PersistentVolumeClaim)`

SetVolumeClaimTemplates sets VolumeClaimTemplates field to given value.

### HasVolumeClaimTemplates

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasVolumeClaimTemplates() bool`

HasVolumeClaimTemplates returns a boolean if a field has been set.

### GetVolumes

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetVolumes() []IoK8sApiCoreV1Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetVolumesOk() (*[]IoK8sApiCoreV1Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetVolumes(v []IoK8sApiCoreV1Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetWorkflowMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetWorkflowMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetWorkflowMetadata returns the WorkflowMetadata field if non-nil, zero value otherwise.

### GetWorkflowMetadataOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetWorkflowMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetWorkflowMetadataOk returns a tuple with the WorkflowMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetWorkflowMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetWorkflowMetadata sets WorkflowMetadata field to given value.

### HasWorkflowMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasWorkflowMetadata() bool`

HasWorkflowMetadata returns a boolean if a field has been set.

### GetWorkflowTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetWorkflowTemplateRef() IoArgoprojWorkflowV1alpha1WorkflowTemplateRef`

GetWorkflowTemplateRef returns the WorkflowTemplateRef field if non-nil, zero value otherwise.

### GetWorkflowTemplateRefOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) GetWorkflowTemplateRefOk() (*IoArgoprojWorkflowV1alpha1WorkflowTemplateRef, bool)`

GetWorkflowTemplateRefOk returns a tuple with the WorkflowTemplateRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) SetWorkflowTemplateRef(v IoArgoprojWorkflowV1alpha1WorkflowTemplateRef)`

SetWorkflowTemplateRef sets WorkflowTemplateRef field to given value.

### HasWorkflowTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec) HasWorkflowTemplateRef() bool`

HasWorkflowTemplateRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


