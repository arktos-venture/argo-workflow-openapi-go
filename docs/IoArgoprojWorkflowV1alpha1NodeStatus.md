# IoArgoprojWorkflowV1alpha1NodeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundaryID** | Pointer to **string** | BoundaryID indicates the node ID of the associated template root node in which this node belongs to | [optional] 
**Children** | Pointer to **[]string** | Children is a list of child node IDs | [optional] 
**Daemoned** | Pointer to **bool** | Daemoned tracks whether or not this node was daemoned and need to be terminated | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is a human readable representation of the node. Unique within a template boundary | [optional] 
**EstimatedDuration** | Pointer to **int32** | EstimatedDuration in seconds. | [optional] 
**FinishedAt** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**HostNodeName** | Pointer to **string** | HostNodeName name of the Kubernetes node on which the Pod is running, if applicable | [optional] 
**Id** | **string** | ID is a unique identifier of a node within the worklow It is implemented as a hash of the node name, which makes the ID deterministic | 
**Inputs** | Pointer to [**IoArgoprojWorkflowV1alpha1Inputs**](IoArgoprojWorkflowV1alpha1Inputs.md) |  | [optional] 
**MemoizationStatus** | Pointer to [**IoArgoprojWorkflowV1alpha1MemoizationStatus**](IoArgoprojWorkflowV1alpha1MemoizationStatus.md) |  | [optional] 
**Message** | Pointer to **string** | A human readable message indicating details about why the node is in this condition. | [optional] 
**Name** | **string** | Name is unique name in the node tree used to generate the node ID | 
**OutboundNodes** | Pointer to **[]string** | OutboundNodes tracks the node IDs which are considered \&quot;outbound\&quot; nodes to a template invocation. For every invocation of a template, there are nodes which we considered as \&quot;outbound\&quot;. Essentially, these are last nodes in the execution sequence to run, before the template is considered completed. These nodes are then connected as parents to a following step.  In the case of single pod steps (i.e. container, script, resource templates), this list will be nil since the pod itself is already considered the \&quot;outbound\&quot; node. In the case of DAGs, outbound nodes are the \&quot;target\&quot; tasks (tasks with no children). In the case of steps, outbound nodes are all the containers involved in the last step group. NOTE: since templates are composable, the list of outbound nodes are carried upwards when a DAG/steps template invokes another DAG/steps template. In other words, the outbound nodes of a template, will be a superset of the outbound nodes of its last children. | [optional] 
**Outputs** | Pointer to [**IoArgoprojWorkflowV1alpha1Outputs**](IoArgoprojWorkflowV1alpha1Outputs.md) |  | [optional] 
**Phase** | Pointer to **string** | Phase a simple, high-level summary of where the node is in its lifecycle. Can be used as a state machine. | [optional] 
**PodIP** | Pointer to **string** | PodIP captures the IP of the pod for daemoned steps | [optional] 
**Progress** | Pointer to **string** | Progress to completion | [optional] 
**ResourcesDuration** | Pointer to **map[string]int64** | ResourcesDuration is indicative, but not accurate, resource duration. This is populated when the nodes completes. | [optional] 
**StartedAt** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**SynchronizationStatus** | Pointer to [**IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus**](IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus.md) |  | [optional] 
**TemplateName** | Pointer to **string** | TemplateName is the template name which this node corresponds to. Not applicable to virtual nodes (e.g. Retry, StepGroup) | [optional] 
**TemplateRef** | Pointer to [**IoArgoprojWorkflowV1alpha1TemplateRef**](IoArgoprojWorkflowV1alpha1TemplateRef.md) |  | [optional] 
**TemplateScope** | Pointer to **string** | TemplateScope is the template scope in which the template of this node was retrieved. | [optional] 
**Type** | **string** | Type indicates type of node | 

## Methods

### NewIoArgoprojWorkflowV1alpha1NodeStatus

`func NewIoArgoprojWorkflowV1alpha1NodeStatus(id string, name string, type_ string, ) *IoArgoprojWorkflowV1alpha1NodeStatus`

NewIoArgoprojWorkflowV1alpha1NodeStatus instantiates a new IoArgoprojWorkflowV1alpha1NodeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1NodeStatusWithDefaults

`func NewIoArgoprojWorkflowV1alpha1NodeStatusWithDefaults() *IoArgoprojWorkflowV1alpha1NodeStatus`

NewIoArgoprojWorkflowV1alpha1NodeStatusWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1NodeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundaryID

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetBoundaryID() string`

GetBoundaryID returns the BoundaryID field if non-nil, zero value otherwise.

### GetBoundaryIDOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetBoundaryIDOk() (*string, bool)`

GetBoundaryIDOk returns a tuple with the BoundaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundaryID

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetBoundaryID(v string)`

SetBoundaryID sets BoundaryID field to given value.

### HasBoundaryID

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasBoundaryID() bool`

HasBoundaryID returns a boolean if a field has been set.

### GetChildren

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetChildren() []string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetChildrenOk() (*[]string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetChildren(v []string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetDaemoned

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetDaemoned() bool`

GetDaemoned returns the Daemoned field if non-nil, zero value otherwise.

### GetDaemonedOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetDaemonedOk() (*bool, bool)`

GetDaemonedOk returns a tuple with the Daemoned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemoned

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetDaemoned(v bool)`

SetDaemoned sets Daemoned field to given value.

### HasDaemoned

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasDaemoned() bool`

HasDaemoned returns a boolean if a field has been set.

### GetDisplayName

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEstimatedDuration

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetEstimatedDuration() int32`

GetEstimatedDuration returns the EstimatedDuration field if non-nil, zero value otherwise.

### GetEstimatedDurationOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetEstimatedDurationOk() (*int32, bool)`

GetEstimatedDurationOk returns a tuple with the EstimatedDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDuration

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetEstimatedDuration(v int32)`

SetEstimatedDuration sets EstimatedDuration field to given value.

### HasEstimatedDuration

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasEstimatedDuration() bool`

HasEstimatedDuration returns a boolean if a field has been set.

### GetFinishedAt

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetHostNodeName

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetHostNodeName() string`

GetHostNodeName returns the HostNodeName field if non-nil, zero value otherwise.

### GetHostNodeNameOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetHostNodeNameOk() (*string, bool)`

GetHostNodeNameOk returns a tuple with the HostNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNodeName

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetHostNodeName(v string)`

SetHostNodeName sets HostNodeName field to given value.

### HasHostNodeName

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasHostNodeName() bool`

HasHostNodeName returns a boolean if a field has been set.

### GetId

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetId(v string)`

SetId sets Id field to given value.


### GetInputs

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetInputs() IoArgoprojWorkflowV1alpha1Inputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetInputsOk() (*IoArgoprojWorkflowV1alpha1Inputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetInputs(v IoArgoprojWorkflowV1alpha1Inputs)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetMemoizationStatus

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetMemoizationStatus() IoArgoprojWorkflowV1alpha1MemoizationStatus`

GetMemoizationStatus returns the MemoizationStatus field if non-nil, zero value otherwise.

### GetMemoizationStatusOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetMemoizationStatusOk() (*IoArgoprojWorkflowV1alpha1MemoizationStatus, bool)`

GetMemoizationStatusOk returns a tuple with the MemoizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoizationStatus

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetMemoizationStatus(v IoArgoprojWorkflowV1alpha1MemoizationStatus)`

SetMemoizationStatus sets MemoizationStatus field to given value.

### HasMemoizationStatus

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasMemoizationStatus() bool`

HasMemoizationStatus returns a boolean if a field has been set.

### GetMessage

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetName(v string)`

SetName sets Name field to given value.


### GetOutboundNodes

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetOutboundNodes() []string`

GetOutboundNodes returns the OutboundNodes field if non-nil, zero value otherwise.

### GetOutboundNodesOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetOutboundNodesOk() (*[]string, bool)`

GetOutboundNodesOk returns a tuple with the OutboundNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundNodes

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetOutboundNodes(v []string)`

SetOutboundNodes sets OutboundNodes field to given value.

### HasOutboundNodes

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasOutboundNodes() bool`

HasOutboundNodes returns a boolean if a field has been set.

### GetOutputs

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetOutputs() IoArgoprojWorkflowV1alpha1Outputs`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetOutputsOk() (*IoArgoprojWorkflowV1alpha1Outputs, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetOutputs(v IoArgoprojWorkflowV1alpha1Outputs)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetPhase

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPodIP

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetPodIP() string`

GetPodIP returns the PodIP field if non-nil, zero value otherwise.

### GetPodIPOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetPodIPOk() (*string, bool)`

GetPodIPOk returns a tuple with the PodIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodIP

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetPodIP(v string)`

SetPodIP sets PodIP field to given value.

### HasPodIP

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasPodIP() bool`

HasPodIP returns a boolean if a field has been set.

### GetProgress

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetProgress(v string)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetResourcesDuration

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetResourcesDuration() map[string]int64`

GetResourcesDuration returns the ResourcesDuration field if non-nil, zero value otherwise.

### GetResourcesDurationOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetResourcesDurationOk() (*map[string]int64, bool)`

GetResourcesDurationOk returns a tuple with the ResourcesDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesDuration

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetResourcesDuration(v map[string]int64)`

SetResourcesDuration sets ResourcesDuration field to given value.

### HasResourcesDuration

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasResourcesDuration() bool`

HasResourcesDuration returns a boolean if a field has been set.

### GetStartedAt

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetSynchronizationStatus

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetSynchronizationStatus() IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus`

GetSynchronizationStatus returns the SynchronizationStatus field if non-nil, zero value otherwise.

### GetSynchronizationStatusOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetSynchronizationStatusOk() (*IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus, bool)`

GetSynchronizationStatusOk returns a tuple with the SynchronizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationStatus

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetSynchronizationStatus(v IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus)`

SetSynchronizationStatus sets SynchronizationStatus field to given value.

### HasSynchronizationStatus

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasSynchronizationStatus() bool`

HasSynchronizationStatus returns a boolean if a field has been set.

### GetTemplateName

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTemplateRef() IoArgoprojWorkflowV1alpha1TemplateRef`

GetTemplateRef returns the TemplateRef field if non-nil, zero value otherwise.

### GetTemplateRefOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTemplateRefOk() (*IoArgoprojWorkflowV1alpha1TemplateRef, bool)`

GetTemplateRefOk returns a tuple with the TemplateRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetTemplateRef(v IoArgoprojWorkflowV1alpha1TemplateRef)`

SetTemplateRef sets TemplateRef field to given value.

### HasTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasTemplateRef() bool`

HasTemplateRef returns a boolean if a field has been set.

### GetTemplateScope

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTemplateScope() string`

GetTemplateScope returns the TemplateScope field if non-nil, zero value otherwise.

### GetTemplateScopeOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTemplateScopeOk() (*string, bool)`

GetTemplateScopeOk returns a tuple with the TemplateScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateScope

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetTemplateScope(v string)`

SetTemplateScope sets TemplateScope field to given value.

### HasTemplateScope

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasTemplateScope() bool`

HasTemplateScope returns a boolean if a field has been set.

### GetType

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


