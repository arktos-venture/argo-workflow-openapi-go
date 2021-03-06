/*
Argo Server API

You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`

API version: VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// IoArgoprojWorkflowV1alpha1NodeStatus NodeStatus contains status information about an individual node in the workflow
type IoArgoprojWorkflowV1alpha1NodeStatus struct {
	// BoundaryID indicates the node ID of the associated template root node in which this node belongs to
	BoundaryID *string `json:"boundaryID,omitempty"`
	// Children is a list of child node IDs
	Children *[]string `json:"children,omitempty"`
	// Daemoned tracks whether or not this node was daemoned and need to be terminated
	Daemoned *bool `json:"daemoned,omitempty"`
	// DisplayName is a human readable representation of the node. Unique within a template boundary
	DisplayName *string `json:"displayName,omitempty"`
	// EstimatedDuration in seconds.
	EstimatedDuration *int32 `json:"estimatedDuration,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// HostNodeName name of the Kubernetes node on which the Pod is running, if applicable
	HostNodeName *string `json:"hostNodeName,omitempty"`
	// ID is a unique identifier of a node within the worklow It is implemented as a hash of the node name, which makes the ID deterministic
	Id string `json:"id"`
	Inputs *IoArgoprojWorkflowV1alpha1Inputs `json:"inputs,omitempty"`
	MemoizationStatus *IoArgoprojWorkflowV1alpha1MemoizationStatus `json:"memoizationStatus,omitempty"`
	// A human readable message indicating details about why the node is in this condition.
	Message *string `json:"message,omitempty"`
	// Name is unique name in the node tree used to generate the node ID
	Name string `json:"name"`
	// OutboundNodes tracks the node IDs which are considered \"outbound\" nodes to a template invocation. For every invocation of a template, there are nodes which we considered as \"outbound\". Essentially, these are last nodes in the execution sequence to run, before the template is considered completed. These nodes are then connected as parents to a following step.  In the case of single pod steps (i.e. container, script, resource templates), this list will be nil since the pod itself is already considered the \"outbound\" node. In the case of DAGs, outbound nodes are the \"target\" tasks (tasks with no children). In the case of steps, outbound nodes are all the containers involved in the last step group. NOTE: since templates are composable, the list of outbound nodes are carried upwards when a DAG/steps template invokes another DAG/steps template. In other words, the outbound nodes of a template, will be a superset of the outbound nodes of its last children.
	OutboundNodes *[]string `json:"outboundNodes,omitempty"`
	Outputs *IoArgoprojWorkflowV1alpha1Outputs `json:"outputs,omitempty"`
	// Phase a simple, high-level summary of where the node is in its lifecycle. Can be used as a state machine.
	Phase *string `json:"phase,omitempty"`
	// PodIP captures the IP of the pod for daemoned steps
	PodIP *string `json:"podIP,omitempty"`
	// Progress to completion
	Progress *string `json:"progress,omitempty"`
	// ResourcesDuration is indicative, but not accurate, resource duration. This is populated when the nodes completes.
	ResourcesDuration *map[string]int64 `json:"resourcesDuration,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	StartedAt *time.Time `json:"startedAt,omitempty"`
	SynchronizationStatus *IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus `json:"synchronizationStatus,omitempty"`
	// TemplateName is the template name which this node corresponds to. Not applicable to virtual nodes (e.g. Retry, StepGroup)
	TemplateName *string `json:"templateName,omitempty"`
	TemplateRef *IoArgoprojWorkflowV1alpha1TemplateRef `json:"templateRef,omitempty"`
	// TemplateScope is the template scope in which the template of this node was retrieved.
	TemplateScope *string `json:"templateScope,omitempty"`
	// Type indicates type of node
	Type string `json:"type"`
}

// NewIoArgoprojWorkflowV1alpha1NodeStatus instantiates a new IoArgoprojWorkflowV1alpha1NodeStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1NodeStatus(id string, name string, type_ string) *IoArgoprojWorkflowV1alpha1NodeStatus {
	this := IoArgoprojWorkflowV1alpha1NodeStatus{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewIoArgoprojWorkflowV1alpha1NodeStatusWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1NodeStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1NodeStatusWithDefaults() *IoArgoprojWorkflowV1alpha1NodeStatus {
	this := IoArgoprojWorkflowV1alpha1NodeStatus{}
	return &this
}

// GetBoundaryID returns the BoundaryID field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetBoundaryID() string {
	if o == nil || o.BoundaryID == nil {
		var ret string
		return ret
	}
	return *o.BoundaryID
}

// GetBoundaryIDOk returns a tuple with the BoundaryID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetBoundaryIDOk() (*string, bool) {
	if o == nil || o.BoundaryID == nil {
		return nil, false
	}
	return o.BoundaryID, true
}

// HasBoundaryID returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasBoundaryID() bool {
	if o != nil && o.BoundaryID != nil {
		return true
	}

	return false
}

// SetBoundaryID gets a reference to the given string and assigns it to the BoundaryID field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetBoundaryID(v string) {
	o.BoundaryID = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetChildren() []string {
	if o == nil || o.Children == nil {
		var ret []string
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetChildrenOk() (*[]string, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []string and assigns it to the Children field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetChildren(v []string) {
	o.Children = &v
}

// GetDaemoned returns the Daemoned field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetDaemoned() bool {
	if o == nil || o.Daemoned == nil {
		var ret bool
		return ret
	}
	return *o.Daemoned
}

// GetDaemonedOk returns a tuple with the Daemoned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetDaemonedOk() (*bool, bool) {
	if o == nil || o.Daemoned == nil {
		return nil, false
	}
	return o.Daemoned, true
}

// HasDaemoned returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasDaemoned() bool {
	if o != nil && o.Daemoned != nil {
		return true
	}

	return false
}

// SetDaemoned gets a reference to the given bool and assigns it to the Daemoned field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetDaemoned(v bool) {
	o.Daemoned = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEstimatedDuration returns the EstimatedDuration field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetEstimatedDuration() int32 {
	if o == nil || o.EstimatedDuration == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedDuration
}

// GetEstimatedDurationOk returns a tuple with the EstimatedDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetEstimatedDurationOk() (*int32, bool) {
	if o == nil || o.EstimatedDuration == nil {
		return nil, false
	}
	return o.EstimatedDuration, true
}

// HasEstimatedDuration returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasEstimatedDuration() bool {
	if o != nil && o.EstimatedDuration != nil {
		return true
	}

	return false
}

// SetEstimatedDuration gets a reference to the given int32 and assigns it to the EstimatedDuration field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetEstimatedDuration(v int32) {
	o.EstimatedDuration = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasFinishedAt() bool {
	if o != nil && o.FinishedAt != nil {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetHostNodeName returns the HostNodeName field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetHostNodeName() string {
	if o == nil || o.HostNodeName == nil {
		var ret string
		return ret
	}
	return *o.HostNodeName
}

// GetHostNodeNameOk returns a tuple with the HostNodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetHostNodeNameOk() (*string, bool) {
	if o == nil || o.HostNodeName == nil {
		return nil, false
	}
	return o.HostNodeName, true
}

// HasHostNodeName returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasHostNodeName() bool {
	if o != nil && o.HostNodeName != nil {
		return true
	}

	return false
}

// SetHostNodeName gets a reference to the given string and assigns it to the HostNodeName field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetHostNodeName(v string) {
	o.HostNodeName = &v
}

// GetId returns the Id field value
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetInputs() IoArgoprojWorkflowV1alpha1Inputs {
	if o == nil || o.Inputs == nil {
		var ret IoArgoprojWorkflowV1alpha1Inputs
		return ret
	}
	return *o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetInputsOk() (*IoArgoprojWorkflowV1alpha1Inputs, bool) {
	if o == nil || o.Inputs == nil {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasInputs() bool {
	if o != nil && o.Inputs != nil {
		return true
	}

	return false
}

// SetInputs gets a reference to the given IoArgoprojWorkflowV1alpha1Inputs and assigns it to the Inputs field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetInputs(v IoArgoprojWorkflowV1alpha1Inputs) {
	o.Inputs = &v
}

// GetMemoizationStatus returns the MemoizationStatus field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetMemoizationStatus() IoArgoprojWorkflowV1alpha1MemoizationStatus {
	if o == nil || o.MemoizationStatus == nil {
		var ret IoArgoprojWorkflowV1alpha1MemoizationStatus
		return ret
	}
	return *o.MemoizationStatus
}

// GetMemoizationStatusOk returns a tuple with the MemoizationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetMemoizationStatusOk() (*IoArgoprojWorkflowV1alpha1MemoizationStatus, bool) {
	if o == nil || o.MemoizationStatus == nil {
		return nil, false
	}
	return o.MemoizationStatus, true
}

// HasMemoizationStatus returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasMemoizationStatus() bool {
	if o != nil && o.MemoizationStatus != nil {
		return true
	}

	return false
}

// SetMemoizationStatus gets a reference to the given IoArgoprojWorkflowV1alpha1MemoizationStatus and assigns it to the MemoizationStatus field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetMemoizationStatus(v IoArgoprojWorkflowV1alpha1MemoizationStatus) {
	o.MemoizationStatus = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetName(v string) {
	o.Name = v
}

// GetOutboundNodes returns the OutboundNodes field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetOutboundNodes() []string {
	if o == nil || o.OutboundNodes == nil {
		var ret []string
		return ret
	}
	return *o.OutboundNodes
}

// GetOutboundNodesOk returns a tuple with the OutboundNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetOutboundNodesOk() (*[]string, bool) {
	if o == nil || o.OutboundNodes == nil {
		return nil, false
	}
	return o.OutboundNodes, true
}

// HasOutboundNodes returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasOutboundNodes() bool {
	if o != nil && o.OutboundNodes != nil {
		return true
	}

	return false
}

// SetOutboundNodes gets a reference to the given []string and assigns it to the OutboundNodes field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetOutboundNodes(v []string) {
	o.OutboundNodes = &v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetOutputs() IoArgoprojWorkflowV1alpha1Outputs {
	if o == nil || o.Outputs == nil {
		var ret IoArgoprojWorkflowV1alpha1Outputs
		return ret
	}
	return *o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetOutputsOk() (*IoArgoprojWorkflowV1alpha1Outputs, bool) {
	if o == nil || o.Outputs == nil {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasOutputs() bool {
	if o != nil && o.Outputs != nil {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given IoArgoprojWorkflowV1alpha1Outputs and assigns it to the Outputs field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetOutputs(v IoArgoprojWorkflowV1alpha1Outputs) {
	o.Outputs = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasPhase() bool {
	if o != nil && o.Phase != nil {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetPhase(v string) {
	o.Phase = &v
}

// GetPodIP returns the PodIP field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetPodIP() string {
	if o == nil || o.PodIP == nil {
		var ret string
		return ret
	}
	return *o.PodIP
}

// GetPodIPOk returns a tuple with the PodIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetPodIPOk() (*string, bool) {
	if o == nil || o.PodIP == nil {
		return nil, false
	}
	return o.PodIP, true
}

// HasPodIP returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasPodIP() bool {
	if o != nil && o.PodIP != nil {
		return true
	}

	return false
}

// SetPodIP gets a reference to the given string and assigns it to the PodIP field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetPodIP(v string) {
	o.PodIP = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetProgress() string {
	if o == nil || o.Progress == nil {
		var ret string
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetProgressOk() (*string, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given string and assigns it to the Progress field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetProgress(v string) {
	o.Progress = &v
}

// GetResourcesDuration returns the ResourcesDuration field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetResourcesDuration() map[string]int64 {
	if o == nil || o.ResourcesDuration == nil {
		var ret map[string]int64
		return ret
	}
	return *o.ResourcesDuration
}

// GetResourcesDurationOk returns a tuple with the ResourcesDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetResourcesDurationOk() (*map[string]int64, bool) {
	if o == nil || o.ResourcesDuration == nil {
		return nil, false
	}
	return o.ResourcesDuration, true
}

// HasResourcesDuration returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasResourcesDuration() bool {
	if o != nil && o.ResourcesDuration != nil {
		return true
	}

	return false
}

// SetResourcesDuration gets a reference to the given map[string]int64 and assigns it to the ResourcesDuration field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetResourcesDuration(v map[string]int64) {
	o.ResourcesDuration = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasStartedAt() bool {
	if o != nil && o.StartedAt != nil {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetSynchronizationStatus returns the SynchronizationStatus field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetSynchronizationStatus() IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus {
	if o == nil || o.SynchronizationStatus == nil {
		var ret IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus
		return ret
	}
	return *o.SynchronizationStatus
}

// GetSynchronizationStatusOk returns a tuple with the SynchronizationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetSynchronizationStatusOk() (*IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus, bool) {
	if o == nil || o.SynchronizationStatus == nil {
		return nil, false
	}
	return o.SynchronizationStatus, true
}

// HasSynchronizationStatus returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasSynchronizationStatus() bool {
	if o != nil && o.SynchronizationStatus != nil {
		return true
	}

	return false
}

// SetSynchronizationStatus gets a reference to the given IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus and assigns it to the SynchronizationStatus field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetSynchronizationStatus(v IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus) {
	o.SynchronizationStatus = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTemplateName() string {
	if o == nil || o.TemplateName == nil {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTemplateNameOk() (*string, bool) {
	if o == nil || o.TemplateName == nil {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasTemplateName() bool {
	if o != nil && o.TemplateName != nil {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetTemplateRef returns the TemplateRef field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTemplateRef() IoArgoprojWorkflowV1alpha1TemplateRef {
	if o == nil || o.TemplateRef == nil {
		var ret IoArgoprojWorkflowV1alpha1TemplateRef
		return ret
	}
	return *o.TemplateRef
}

// GetTemplateRefOk returns a tuple with the TemplateRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTemplateRefOk() (*IoArgoprojWorkflowV1alpha1TemplateRef, bool) {
	if o == nil || o.TemplateRef == nil {
		return nil, false
	}
	return o.TemplateRef, true
}

// HasTemplateRef returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasTemplateRef() bool {
	if o != nil && o.TemplateRef != nil {
		return true
	}

	return false
}

// SetTemplateRef gets a reference to the given IoArgoprojWorkflowV1alpha1TemplateRef and assigns it to the TemplateRef field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetTemplateRef(v IoArgoprojWorkflowV1alpha1TemplateRef) {
	o.TemplateRef = &v
}

// GetTemplateScope returns the TemplateScope field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTemplateScope() string {
	if o == nil || o.TemplateScope == nil {
		var ret string
		return ret
	}
	return *o.TemplateScope
}

// GetTemplateScopeOk returns a tuple with the TemplateScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTemplateScopeOk() (*string, bool) {
	if o == nil || o.TemplateScope == nil {
		return nil, false
	}
	return o.TemplateScope, true
}

// HasTemplateScope returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) HasTemplateScope() bool {
	if o != nil && o.TemplateScope != nil {
		return true
	}

	return false
}

// SetTemplateScope gets a reference to the given string and assigns it to the TemplateScope field.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetTemplateScope(v string) {
	o.TemplateScope = &v
}

// GetType returns the Type field value
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IoArgoprojWorkflowV1alpha1NodeStatus) SetType(v string) {
	o.Type = v
}

func (o IoArgoprojWorkflowV1alpha1NodeStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BoundaryID != nil {
		toSerialize["boundaryID"] = o.BoundaryID
	}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.Daemoned != nil {
		toSerialize["daemoned"] = o.Daemoned
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.EstimatedDuration != nil {
		toSerialize["estimatedDuration"] = o.EstimatedDuration
	}
	if o.FinishedAt != nil {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	if o.HostNodeName != nil {
		toSerialize["hostNodeName"] = o.HostNodeName
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Inputs != nil {
		toSerialize["inputs"] = o.Inputs
	}
	if o.MemoizationStatus != nil {
		toSerialize["memoizationStatus"] = o.MemoizationStatus
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OutboundNodes != nil {
		toSerialize["outboundNodes"] = o.OutboundNodes
	}
	if o.Outputs != nil {
		toSerialize["outputs"] = o.Outputs
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.PodIP != nil {
		toSerialize["podIP"] = o.PodIP
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	if o.ResourcesDuration != nil {
		toSerialize["resourcesDuration"] = o.ResourcesDuration
	}
	if o.StartedAt != nil {
		toSerialize["startedAt"] = o.StartedAt
	}
	if o.SynchronizationStatus != nil {
		toSerialize["synchronizationStatus"] = o.SynchronizationStatus
	}
	if o.TemplateName != nil {
		toSerialize["templateName"] = o.TemplateName
	}
	if o.TemplateRef != nil {
		toSerialize["templateRef"] = o.TemplateRef
	}
	if o.TemplateScope != nil {
		toSerialize["templateScope"] = o.TemplateScope
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1NodeStatus struct {
	value *IoArgoprojWorkflowV1alpha1NodeStatus
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1NodeStatus) Get() *IoArgoprojWorkflowV1alpha1NodeStatus {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1NodeStatus) Set(val *IoArgoprojWorkflowV1alpha1NodeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1NodeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1NodeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1NodeStatus(val *IoArgoprojWorkflowV1alpha1NodeStatus) *NullableIoArgoprojWorkflowV1alpha1NodeStatus {
	return &NullableIoArgoprojWorkflowV1alpha1NodeStatus{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1NodeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1NodeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


