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

// IoArgoprojWorkflowV1alpha1ResourceTemplate ResourceTemplate is a template subtype to manipulate kubernetes resources
type IoArgoprojWorkflowV1alpha1ResourceTemplate struct {
	// Action is the action to perform to the resource. Must be one of: get, create, apply, delete, replace, patch
	Action string `json:"action"`
	// FailureCondition is a label selector expression which describes the conditions of the k8s resource in which the step was considered failed
	FailureCondition *string `json:"failureCondition,omitempty"`
	// Flags is a set of additional options passed to kubectl before submitting a resource I.e. to disable resource validation: flags: [  \"--validate=false\"  # disable resource validation ]
	Flags *[]string `json:"flags,omitempty"`
	// Manifest contains the kubernetes manifest
	Manifest *string `json:"manifest,omitempty"`
	// MergeStrategy is the strategy used to merge a patch. It defaults to \"strategic\" Must be one of: strategic, merge, json
	MergeStrategy *string `json:"mergeStrategy,omitempty"`
	// SetOwnerReference sets the reference to the workflow on the OwnerReference of generated resource.
	SetOwnerReference *bool `json:"setOwnerReference,omitempty"`
	// SuccessCondition is a label selector expression which describes the conditions of the k8s resource in which it is acceptable to proceed to the following step
	SuccessCondition *string `json:"successCondition,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1ResourceTemplate instantiates a new IoArgoprojWorkflowV1alpha1ResourceTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1ResourceTemplate(action string) *IoArgoprojWorkflowV1alpha1ResourceTemplate {
	this := IoArgoprojWorkflowV1alpha1ResourceTemplate{}
	this.Action = action
	return &this
}

// NewIoArgoprojWorkflowV1alpha1ResourceTemplateWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ResourceTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1ResourceTemplateWithDefaults() *IoArgoprojWorkflowV1alpha1ResourceTemplate {
	this := IoArgoprojWorkflowV1alpha1ResourceTemplate{}
	return &this
}

// GetAction returns the Action field value
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetAction(v string) {
	o.Action = v
}

// GetFailureCondition returns the FailureCondition field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetFailureCondition() string {
	if o == nil || o.FailureCondition == nil {
		var ret string
		return ret
	}
	return *o.FailureCondition
}

// GetFailureConditionOk returns a tuple with the FailureCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetFailureConditionOk() (*string, bool) {
	if o == nil || o.FailureCondition == nil {
		return nil, false
	}
	return o.FailureCondition, true
}

// HasFailureCondition returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) HasFailureCondition() bool {
	if o != nil && o.FailureCondition != nil {
		return true
	}

	return false
}

// SetFailureCondition gets a reference to the given string and assigns it to the FailureCondition field.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetFailureCondition(v string) {
	o.FailureCondition = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetFlags() []string {
	if o == nil || o.Flags == nil {
		var ret []string
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetFlagsOk() (*[]string, bool) {
	if o == nil || o.Flags == nil {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetFlags(v []string) {
	o.Flags = &v
}

// GetManifest returns the Manifest field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetManifest() string {
	if o == nil || o.Manifest == nil {
		var ret string
		return ret
	}
	return *o.Manifest
}

// GetManifestOk returns a tuple with the Manifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetManifestOk() (*string, bool) {
	if o == nil || o.Manifest == nil {
		return nil, false
	}
	return o.Manifest, true
}

// HasManifest returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) HasManifest() bool {
	if o != nil && o.Manifest != nil {
		return true
	}

	return false
}

// SetManifest gets a reference to the given string and assigns it to the Manifest field.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetManifest(v string) {
	o.Manifest = &v
}

// GetMergeStrategy returns the MergeStrategy field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetMergeStrategy() string {
	if o == nil || o.MergeStrategy == nil {
		var ret string
		return ret
	}
	return *o.MergeStrategy
}

// GetMergeStrategyOk returns a tuple with the MergeStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetMergeStrategyOk() (*string, bool) {
	if o == nil || o.MergeStrategy == nil {
		return nil, false
	}
	return o.MergeStrategy, true
}

// HasMergeStrategy returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) HasMergeStrategy() bool {
	if o != nil && o.MergeStrategy != nil {
		return true
	}

	return false
}

// SetMergeStrategy gets a reference to the given string and assigns it to the MergeStrategy field.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetMergeStrategy(v string) {
	o.MergeStrategy = &v
}

// GetSetOwnerReference returns the SetOwnerReference field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetSetOwnerReference() bool {
	if o == nil || o.SetOwnerReference == nil {
		var ret bool
		return ret
	}
	return *o.SetOwnerReference
}

// GetSetOwnerReferenceOk returns a tuple with the SetOwnerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetSetOwnerReferenceOk() (*bool, bool) {
	if o == nil || o.SetOwnerReference == nil {
		return nil, false
	}
	return o.SetOwnerReference, true
}

// HasSetOwnerReference returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) HasSetOwnerReference() bool {
	if o != nil && o.SetOwnerReference != nil {
		return true
	}

	return false
}

// SetSetOwnerReference gets a reference to the given bool and assigns it to the SetOwnerReference field.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetSetOwnerReference(v bool) {
	o.SetOwnerReference = &v
}

// GetSuccessCondition returns the SuccessCondition field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetSuccessCondition() string {
	if o == nil || o.SuccessCondition == nil {
		var ret string
		return ret
	}
	return *o.SuccessCondition
}

// GetSuccessConditionOk returns a tuple with the SuccessCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetSuccessConditionOk() (*string, bool) {
	if o == nil || o.SuccessCondition == nil {
		return nil, false
	}
	return o.SuccessCondition, true
}

// HasSuccessCondition returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) HasSuccessCondition() bool {
	if o != nil && o.SuccessCondition != nil {
		return true
	}

	return false
}

// SetSuccessCondition gets a reference to the given string and assigns it to the SuccessCondition field.
func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetSuccessCondition(v string) {
	o.SuccessCondition = &v
}

func (o IoArgoprojWorkflowV1alpha1ResourceTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}
	if o.FailureCondition != nil {
		toSerialize["failureCondition"] = o.FailureCondition
	}
	if o.Flags != nil {
		toSerialize["flags"] = o.Flags
	}
	if o.Manifest != nil {
		toSerialize["manifest"] = o.Manifest
	}
	if o.MergeStrategy != nil {
		toSerialize["mergeStrategy"] = o.MergeStrategy
	}
	if o.SetOwnerReference != nil {
		toSerialize["setOwnerReference"] = o.SetOwnerReference
	}
	if o.SuccessCondition != nil {
		toSerialize["successCondition"] = o.SuccessCondition
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1ResourceTemplate struct {
	value *IoArgoprojWorkflowV1alpha1ResourceTemplate
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1ResourceTemplate) Get() *IoArgoprojWorkflowV1alpha1ResourceTemplate {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1ResourceTemplate) Set(val *IoArgoprojWorkflowV1alpha1ResourceTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1ResourceTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1ResourceTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1ResourceTemplate(val *IoArgoprojWorkflowV1alpha1ResourceTemplate) *NullableIoArgoprojWorkflowV1alpha1ResourceTemplate {
	return &NullableIoArgoprojWorkflowV1alpha1ResourceTemplate{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1ResourceTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1ResourceTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

