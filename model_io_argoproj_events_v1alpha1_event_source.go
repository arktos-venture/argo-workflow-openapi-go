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

// IoArgoprojEventsV1alpha1EventSource struct for IoArgoprojEventsV1alpha1EventSource
type IoArgoprojEventsV1alpha1EventSource struct {
	Metadata *IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata,omitempty"`
	Spec *IoArgoprojEventsV1alpha1EventSourceSpec `json:"spec,omitempty"`
	Status *IoArgoprojEventsV1alpha1EventSourceStatus `json:"status,omitempty"`
}

// NewIoArgoprojEventsV1alpha1EventSource instantiates a new IoArgoprojEventsV1alpha1EventSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1EventSource() *IoArgoprojEventsV1alpha1EventSource {
	this := IoArgoprojEventsV1alpha1EventSource{}
	return &this
}

// NewIoArgoprojEventsV1alpha1EventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1EventSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1EventSourceWithDefaults() *IoArgoprojEventsV1alpha1EventSource {
	this := IoArgoprojEventsV1alpha1EventSource{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSource) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret IoK8sApimachineryPkgApisMetaV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSource) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSource) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given IoK8sApimachineryPkgApisMetaV1ObjectMeta and assigns it to the Metadata field.
func (o *IoArgoprojEventsV1alpha1EventSource) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSource) GetSpec() IoArgoprojEventsV1alpha1EventSourceSpec {
	if o == nil || o.Spec == nil {
		var ret IoArgoprojEventsV1alpha1EventSourceSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSource) GetSpecOk() (*IoArgoprojEventsV1alpha1EventSourceSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSource) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given IoArgoprojEventsV1alpha1EventSourceSpec and assigns it to the Spec field.
func (o *IoArgoprojEventsV1alpha1EventSource) SetSpec(v IoArgoprojEventsV1alpha1EventSourceSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSource) GetStatus() IoArgoprojEventsV1alpha1EventSourceStatus {
	if o == nil || o.Status == nil {
		var ret IoArgoprojEventsV1alpha1EventSourceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSource) GetStatusOk() (*IoArgoprojEventsV1alpha1EventSourceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSource) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IoArgoprojEventsV1alpha1EventSourceStatus and assigns it to the Status field.
func (o *IoArgoprojEventsV1alpha1EventSource) SetStatus(v IoArgoprojEventsV1alpha1EventSourceStatus) {
	o.Status = &v
}

func (o IoArgoprojEventsV1alpha1EventSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1EventSource struct {
	value *IoArgoprojEventsV1alpha1EventSource
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1EventSource) Get() *IoArgoprojEventsV1alpha1EventSource {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1EventSource) Set(val *IoArgoprojEventsV1alpha1EventSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1EventSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1EventSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1EventSource(val *IoArgoprojEventsV1alpha1EventSource) *NullableIoArgoprojEventsV1alpha1EventSource {
	return &NullableIoArgoprojEventsV1alpha1EventSource{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1EventSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1EventSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


