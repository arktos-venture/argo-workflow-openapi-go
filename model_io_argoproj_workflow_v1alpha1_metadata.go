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

// IoArgoprojWorkflowV1alpha1Metadata Pod metdata
type IoArgoprojWorkflowV1alpha1Metadata struct {
	Annotations *map[string]string `json:"annotations,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1Metadata instantiates a new IoArgoprojWorkflowV1alpha1Metadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1Metadata() *IoArgoprojWorkflowV1alpha1Metadata {
	this := IoArgoprojWorkflowV1alpha1Metadata{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1MetadataWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Metadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1MetadataWithDefaults() *IoArgoprojWorkflowV1alpha1Metadata {
	this := IoArgoprojWorkflowV1alpha1Metadata{}
	return &this
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1Metadata) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1Metadata) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1Metadata) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *IoArgoprojWorkflowV1alpha1Metadata) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1Metadata) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1Metadata) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1Metadata) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *IoArgoprojWorkflowV1alpha1Metadata) SetLabels(v map[string]string) {
	o.Labels = &v
}

func (o IoArgoprojWorkflowV1alpha1Metadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1Metadata struct {
	value *IoArgoprojWorkflowV1alpha1Metadata
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1Metadata) Get() *IoArgoprojWorkflowV1alpha1Metadata {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1Metadata) Set(val *IoArgoprojWorkflowV1alpha1Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1Metadata) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1Metadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1Metadata(val *IoArgoprojWorkflowV1alpha1Metadata) *NullableIoArgoprojWorkflowV1alpha1Metadata {
	return &NullableIoArgoprojWorkflowV1alpha1Metadata{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1Metadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1Metadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


