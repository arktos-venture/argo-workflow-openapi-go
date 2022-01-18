/*
Argo Server API

You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`

API version: VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IoArgoprojWorkflowV1alpha1DAGTask DAGTask represents a node in the graph during DAG execution
type IoArgoprojWorkflowV1alpha1DAGTask struct {
	Arguments *IoArgoprojWorkflowV1alpha1Arguments `json:"arguments,omitempty"`
	ContinueOn *IoArgoprojWorkflowV1alpha1ContinueOn `json:"continueOn,omitempty"`
	// Dependencies are name of other targets which this depends on
	Dependencies *[]string `json:"dependencies,omitempty"`
	// Depends are name of other targets which this depends on
	Depends *string `json:"depends,omitempty"`
	// Hooks hold the lifecycle hook which is invoked at lifecycle of task, irrespective of the success, failure, or error status of the primary task
	Hooks *map[string]IoArgoprojWorkflowV1alpha1LifecycleHook `json:"hooks,omitempty"`
	Inline *IoArgoprojWorkflowV1alpha1Template `json:"inline,omitempty"`
	// Name is the name of the target
	Name string `json:"name"`
	// OnExit is a template reference which is invoked at the end of the template, irrespective of the success, failure, or error of the primary template. DEPRECATED: Use Hooks[exit].Template instead.
	OnExit *string `json:"onExit,omitempty"`
	// Name of template to execute
	Template *string `json:"template,omitempty"`
	TemplateRef *IoArgoprojWorkflowV1alpha1TemplateRef `json:"templateRef,omitempty"`
	// When is an expression in which the task should conditionally execute
	When *string `json:"when,omitempty"`
	// WithItems expands a task into multiple parallel tasks from the items in the list
	WithItems *[]map[string]interface{} `json:"withItems,omitempty"`
	// WithParam expands a task into multiple parallel tasks from the value in the parameter, which is expected to be a JSON list.
	WithParam *string `json:"withParam,omitempty"`
	WithSequence *IoArgoprojWorkflowV1alpha1Sequence `json:"withSequence,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1DAGTask instantiates a new IoArgoprojWorkflowV1alpha1DAGTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1DAGTask(name string) *IoArgoprojWorkflowV1alpha1DAGTask {
	this := IoArgoprojWorkflowV1alpha1DAGTask{}
	this.Name = name
	return &this
}

// NewIoArgoprojWorkflowV1alpha1DAGTaskWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1DAGTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1DAGTaskWithDefaults() *IoArgoprojWorkflowV1alpha1DAGTask {
	this := IoArgoprojWorkflowV1alpha1DAGTask{}
	return &this
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetArguments() IoArgoprojWorkflowV1alpha1Arguments {
	if o == nil || o.Arguments == nil {
		var ret IoArgoprojWorkflowV1alpha1Arguments
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetArgumentsOk() (*IoArgoprojWorkflowV1alpha1Arguments, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given IoArgoprojWorkflowV1alpha1Arguments and assigns it to the Arguments field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetArguments(v IoArgoprojWorkflowV1alpha1Arguments) {
	o.Arguments = &v
}

// GetContinueOn returns the ContinueOn field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetContinueOn() IoArgoprojWorkflowV1alpha1ContinueOn {
	if o == nil || o.ContinueOn == nil {
		var ret IoArgoprojWorkflowV1alpha1ContinueOn
		return ret
	}
	return *o.ContinueOn
}

// GetContinueOnOk returns a tuple with the ContinueOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetContinueOnOk() (*IoArgoprojWorkflowV1alpha1ContinueOn, bool) {
	if o == nil || o.ContinueOn == nil {
		return nil, false
	}
	return o.ContinueOn, true
}

// HasContinueOn returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasContinueOn() bool {
	if o != nil && o.ContinueOn != nil {
		return true
	}

	return false
}

// SetContinueOn gets a reference to the given IoArgoprojWorkflowV1alpha1ContinueOn and assigns it to the ContinueOn field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetContinueOn(v IoArgoprojWorkflowV1alpha1ContinueOn) {
	o.ContinueOn = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetDependencies() []string {
	if o == nil || o.Dependencies == nil {
		var ret []string
		return ret
	}
	return *o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetDependenciesOk() (*[]string, bool) {
	if o == nil || o.Dependencies == nil {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasDependencies() bool {
	if o != nil && o.Dependencies != nil {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given []string and assigns it to the Dependencies field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetDependencies(v []string) {
	o.Dependencies = &v
}

// GetDepends returns the Depends field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetDepends() string {
	if o == nil || o.Depends == nil {
		var ret string
		return ret
	}
	return *o.Depends
}

// GetDependsOk returns a tuple with the Depends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetDependsOk() (*string, bool) {
	if o == nil || o.Depends == nil {
		return nil, false
	}
	return o.Depends, true
}

// HasDepends returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasDepends() bool {
	if o != nil && o.Depends != nil {
		return true
	}

	return false
}

// SetDepends gets a reference to the given string and assigns it to the Depends field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetDepends(v string) {
	o.Depends = &v
}

// GetHooks returns the Hooks field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetHooks() map[string]IoArgoprojWorkflowV1alpha1LifecycleHook {
	if o == nil || o.Hooks == nil {
		var ret map[string]IoArgoprojWorkflowV1alpha1LifecycleHook
		return ret
	}
	return *o.Hooks
}

// GetHooksOk returns a tuple with the Hooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetHooksOk() (*map[string]IoArgoprojWorkflowV1alpha1LifecycleHook, bool) {
	if o == nil || o.Hooks == nil {
		return nil, false
	}
	return o.Hooks, true
}

// HasHooks returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasHooks() bool {
	if o != nil && o.Hooks != nil {
		return true
	}

	return false
}

// SetHooks gets a reference to the given map[string]IoArgoprojWorkflowV1alpha1LifecycleHook and assigns it to the Hooks field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetHooks(v map[string]IoArgoprojWorkflowV1alpha1LifecycleHook) {
	o.Hooks = &v
}

// GetInline returns the Inline field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetInline() IoArgoprojWorkflowV1alpha1Template {
	if o == nil || o.Inline == nil {
		var ret IoArgoprojWorkflowV1alpha1Template
		return ret
	}
	return *o.Inline
}

// GetInlineOk returns a tuple with the Inline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetInlineOk() (*IoArgoprojWorkflowV1alpha1Template, bool) {
	if o == nil || o.Inline == nil {
		return nil, false
	}
	return o.Inline, true
}

// HasInline returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasInline() bool {
	if o != nil && o.Inline != nil {
		return true
	}

	return false
}

// SetInline gets a reference to the given IoArgoprojWorkflowV1alpha1Template and assigns it to the Inline field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetInline(v IoArgoprojWorkflowV1alpha1Template) {
	o.Inline = &v
}

// GetName returns the Name field value
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetName(v string) {
	o.Name = v
}

// GetOnExit returns the OnExit field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetOnExit() string {
	if o == nil || o.OnExit == nil {
		var ret string
		return ret
	}
	return *o.OnExit
}

// GetOnExitOk returns a tuple with the OnExit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetOnExitOk() (*string, bool) {
	if o == nil || o.OnExit == nil {
		return nil, false
	}
	return o.OnExit, true
}

// HasOnExit returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasOnExit() bool {
	if o != nil && o.OnExit != nil {
		return true
	}

	return false
}

// SetOnExit gets a reference to the given string and assigns it to the OnExit field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetOnExit(v string) {
	o.OnExit = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetTemplate(v string) {
	o.Template = &v
}

// GetTemplateRef returns the TemplateRef field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetTemplateRef() IoArgoprojWorkflowV1alpha1TemplateRef {
	if o == nil || o.TemplateRef == nil {
		var ret IoArgoprojWorkflowV1alpha1TemplateRef
		return ret
	}
	return *o.TemplateRef
}

// GetTemplateRefOk returns a tuple with the TemplateRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetTemplateRefOk() (*IoArgoprojWorkflowV1alpha1TemplateRef, bool) {
	if o == nil || o.TemplateRef == nil {
		return nil, false
	}
	return o.TemplateRef, true
}

// HasTemplateRef returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasTemplateRef() bool {
	if o != nil && o.TemplateRef != nil {
		return true
	}

	return false
}

// SetTemplateRef gets a reference to the given IoArgoprojWorkflowV1alpha1TemplateRef and assigns it to the TemplateRef field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetTemplateRef(v IoArgoprojWorkflowV1alpha1TemplateRef) {
	o.TemplateRef = &v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWhen() string {
	if o == nil || o.When == nil {
		var ret string
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWhenOk() (*string, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasWhen() bool {
	if o != nil && o.When != nil {
		return true
	}

	return false
}

// SetWhen gets a reference to the given string and assigns it to the When field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetWhen(v string) {
	o.When = &v
}

// GetWithItems returns the WithItems field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWithItems() []map[string]interface{} {
	if o == nil || o.WithItems == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.WithItems
}

// GetWithItemsOk returns a tuple with the WithItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWithItemsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.WithItems == nil {
		return nil, false
	}
	return o.WithItems, true
}

// HasWithItems returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasWithItems() bool {
	if o != nil && o.WithItems != nil {
		return true
	}

	return false
}

// SetWithItems gets a reference to the given []map[string]interface{} and assigns it to the WithItems field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetWithItems(v []map[string]interface{}) {
	o.WithItems = &v
}

// GetWithParam returns the WithParam field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWithParam() string {
	if o == nil || o.WithParam == nil {
		var ret string
		return ret
	}
	return *o.WithParam
}

// GetWithParamOk returns a tuple with the WithParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWithParamOk() (*string, bool) {
	if o == nil || o.WithParam == nil {
		return nil, false
	}
	return o.WithParam, true
}

// HasWithParam returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasWithParam() bool {
	if o != nil && o.WithParam != nil {
		return true
	}

	return false
}

// SetWithParam gets a reference to the given string and assigns it to the WithParam field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetWithParam(v string) {
	o.WithParam = &v
}

// GetWithSequence returns the WithSequence field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWithSequence() IoArgoprojWorkflowV1alpha1Sequence {
	if o == nil || o.WithSequence == nil {
		var ret IoArgoprojWorkflowV1alpha1Sequence
		return ret
	}
	return *o.WithSequence
}

// GetWithSequenceOk returns a tuple with the WithSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWithSequenceOk() (*IoArgoprojWorkflowV1alpha1Sequence, bool) {
	if o == nil || o.WithSequence == nil {
		return nil, false
	}
	return o.WithSequence, true
}

// HasWithSequence returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasWithSequence() bool {
	if o != nil && o.WithSequence != nil {
		return true
	}

	return false
}

// SetWithSequence gets a reference to the given IoArgoprojWorkflowV1alpha1Sequence and assigns it to the WithSequence field.
func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetWithSequence(v IoArgoprojWorkflowV1alpha1Sequence) {
	o.WithSequence = &v
}

func (o IoArgoprojWorkflowV1alpha1DAGTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.ContinueOn != nil {
		toSerialize["continueOn"] = o.ContinueOn
	}
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}
	if o.Depends != nil {
		toSerialize["depends"] = o.Depends
	}
	if o.Hooks != nil {
		toSerialize["hooks"] = o.Hooks
	}
	if o.Inline != nil {
		toSerialize["inline"] = o.Inline
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OnExit != nil {
		toSerialize["onExit"] = o.OnExit
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.TemplateRef != nil {
		toSerialize["templateRef"] = o.TemplateRef
	}
	if o.When != nil {
		toSerialize["when"] = o.When
	}
	if o.WithItems != nil {
		toSerialize["withItems"] = o.WithItems
	}
	if o.WithParam != nil {
		toSerialize["withParam"] = o.WithParam
	}
	if o.WithSequence != nil {
		toSerialize["withSequence"] = o.WithSequence
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1DAGTask struct {
	value *IoArgoprojWorkflowV1alpha1DAGTask
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1DAGTask) Get() *IoArgoprojWorkflowV1alpha1DAGTask {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1DAGTask) Set(val *IoArgoprojWorkflowV1alpha1DAGTask) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1DAGTask) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1DAGTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1DAGTask(val *IoArgoprojWorkflowV1alpha1DAGTask) *NullableIoArgoprojWorkflowV1alpha1DAGTask {
	return &NullableIoArgoprojWorkflowV1alpha1DAGTask{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1DAGTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1DAGTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

