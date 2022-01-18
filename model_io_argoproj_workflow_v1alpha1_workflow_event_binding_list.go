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

// IoArgoprojWorkflowV1alpha1WorkflowEventBindingList WorkflowEventBindingList is list of event resources
type IoArgoprojWorkflowV1alpha1WorkflowEventBindingList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	Items []IoArgoprojWorkflowV1alpha1WorkflowEventBinding `json:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata IoK8sApimachineryPkgApisMetaV1ListMeta `json:"metadata"`
}

// NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingList instantiates a new IoArgoprojWorkflowV1alpha1WorkflowEventBindingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingList(items []IoArgoprojWorkflowV1alpha1WorkflowEventBinding, metadata IoK8sApimachineryPkgApisMetaV1ListMeta) *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList {
	this := IoArgoprojWorkflowV1alpha1WorkflowEventBindingList{}
	this.Items = items
	this.Metadata = metadata
	return &this
}

// NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingListWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowEventBindingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingListWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList {
	this := IoArgoprojWorkflowV1alpha1WorkflowEventBindingList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetItems() []IoArgoprojWorkflowV1alpha1WorkflowEventBinding {
	if o == nil {
		var ret []IoArgoprojWorkflowV1alpha1WorkflowEventBinding
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetItemsOk() (*[]IoArgoprojWorkflowV1alpha1WorkflowEventBinding, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) SetItems(v []IoArgoprojWorkflowV1alpha1WorkflowEventBinding) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetMetadata() IoK8sApimachineryPkgApisMetaV1ListMeta {
	if o == nil {
		var ret IoK8sApimachineryPkgApisMetaV1ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ListMeta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ListMeta) {
	o.Metadata = v
}

func (o IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if true {
		toSerialize["items"] = o.Items
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingList struct {
	value *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingList) Get() *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingList) Set(val *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingList) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingList(val *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) *NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingList {
	return &NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingList{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1WorkflowEventBindingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


