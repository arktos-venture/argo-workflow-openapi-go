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

// IoArgoprojWorkflowV1alpha1WorkflowCreateRequest struct for IoArgoprojWorkflowV1alpha1WorkflowCreateRequest
type IoArgoprojWorkflowV1alpha1WorkflowCreateRequest struct {
	CreateOptions *IoK8sApimachineryPkgApisMetaV1CreateOptions `json:"createOptions,omitempty"`
	// This field is no longer used.
	InstanceID *string `json:"instanceID,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	ServerDryRun *bool `json:"serverDryRun,omitempty"`
	Workflow *IoArgoprojWorkflowV1alpha1Workflow `json:"workflow,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1WorkflowCreateRequest instantiates a new IoArgoprojWorkflowV1alpha1WorkflowCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1WorkflowCreateRequest() *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest {
	this := IoArgoprojWorkflowV1alpha1WorkflowCreateRequest{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1WorkflowCreateRequestWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1WorkflowCreateRequestWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest {
	this := IoArgoprojWorkflowV1alpha1WorkflowCreateRequest{}
	return &this
}

// GetCreateOptions returns the CreateOptions field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetCreateOptions() IoK8sApimachineryPkgApisMetaV1CreateOptions {
	if o == nil || o.CreateOptions == nil {
		var ret IoK8sApimachineryPkgApisMetaV1CreateOptions
		return ret
	}
	return *o.CreateOptions
}

// GetCreateOptionsOk returns a tuple with the CreateOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetCreateOptionsOk() (*IoK8sApimachineryPkgApisMetaV1CreateOptions, bool) {
	if o == nil || o.CreateOptions == nil {
		return nil, false
	}
	return o.CreateOptions, true
}

// HasCreateOptions returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) HasCreateOptions() bool {
	if o != nil && o.CreateOptions != nil {
		return true
	}

	return false
}

// SetCreateOptions gets a reference to the given IoK8sApimachineryPkgApisMetaV1CreateOptions and assigns it to the CreateOptions field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) SetCreateOptions(v IoK8sApimachineryPkgApisMetaV1CreateOptions) {
	o.CreateOptions = &v
}

// GetInstanceID returns the InstanceID field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetInstanceID() string {
	if o == nil || o.InstanceID == nil {
		var ret string
		return ret
	}
	return *o.InstanceID
}

// GetInstanceIDOk returns a tuple with the InstanceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetInstanceIDOk() (*string, bool) {
	if o == nil || o.InstanceID == nil {
		return nil, false
	}
	return o.InstanceID, true
}

// HasInstanceID returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) HasInstanceID() bool {
	if o != nil && o.InstanceID != nil {
		return true
	}

	return false
}

// SetInstanceID gets a reference to the given string and assigns it to the InstanceID field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) SetInstanceID(v string) {
	o.InstanceID = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) SetNamespace(v string) {
	o.Namespace = &v
}

// GetServerDryRun returns the ServerDryRun field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetServerDryRun() bool {
	if o == nil || o.ServerDryRun == nil {
		var ret bool
		return ret
	}
	return *o.ServerDryRun
}

// GetServerDryRunOk returns a tuple with the ServerDryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetServerDryRunOk() (*bool, bool) {
	if o == nil || o.ServerDryRun == nil {
		return nil, false
	}
	return o.ServerDryRun, true
}

// HasServerDryRun returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) HasServerDryRun() bool {
	if o != nil && o.ServerDryRun != nil {
		return true
	}

	return false
}

// SetServerDryRun gets a reference to the given bool and assigns it to the ServerDryRun field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) SetServerDryRun(v bool) {
	o.ServerDryRun = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetWorkflow() IoArgoprojWorkflowV1alpha1Workflow {
	if o == nil || o.Workflow == nil {
		var ret IoArgoprojWorkflowV1alpha1Workflow
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetWorkflowOk() (*IoArgoprojWorkflowV1alpha1Workflow, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given IoArgoprojWorkflowV1alpha1Workflow and assigns it to the Workflow field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) SetWorkflow(v IoArgoprojWorkflowV1alpha1Workflow) {
	o.Workflow = &v
}

func (o IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateOptions != nil {
		toSerialize["createOptions"] = o.CreateOptions
	}
	if o.InstanceID != nil {
		toSerialize["instanceID"] = o.InstanceID
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.ServerDryRun != nil {
		toSerialize["serverDryRun"] = o.ServerDryRun
	}
	if o.Workflow != nil {
		toSerialize["workflow"] = o.Workflow
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1WorkflowCreateRequest struct {
	value *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowCreateRequest) Get() *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowCreateRequest) Set(val *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1WorkflowCreateRequest(val *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) *NullableIoArgoprojWorkflowV1alpha1WorkflowCreateRequest {
	return &NullableIoArgoprojWorkflowV1alpha1WorkflowCreateRequest{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


