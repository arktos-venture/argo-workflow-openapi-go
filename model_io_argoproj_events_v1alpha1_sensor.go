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

// IoArgoprojEventsV1alpha1Sensor struct for IoArgoprojEventsV1alpha1Sensor
type IoArgoprojEventsV1alpha1Sensor struct {
	Metadata *IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata,omitempty"`
	Spec *IoArgoprojEventsV1alpha1SensorSpec `json:"spec,omitempty"`
	Status *IoArgoprojEventsV1alpha1SensorStatus `json:"status,omitempty"`
}

// NewIoArgoprojEventsV1alpha1Sensor instantiates a new IoArgoprojEventsV1alpha1Sensor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1Sensor() *IoArgoprojEventsV1alpha1Sensor {
	this := IoArgoprojEventsV1alpha1Sensor{}
	return &this
}

// NewIoArgoprojEventsV1alpha1SensorWithDefaults instantiates a new IoArgoprojEventsV1alpha1Sensor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1SensorWithDefaults() *IoArgoprojEventsV1alpha1Sensor {
	this := IoArgoprojEventsV1alpha1Sensor{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Sensor) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret IoK8sApimachineryPkgApisMetaV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Sensor) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Sensor) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given IoK8sApimachineryPkgApisMetaV1ObjectMeta and assigns it to the Metadata field.
func (o *IoArgoprojEventsV1alpha1Sensor) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Sensor) GetSpec() IoArgoprojEventsV1alpha1SensorSpec {
	if o == nil || o.Spec == nil {
		var ret IoArgoprojEventsV1alpha1SensorSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Sensor) GetSpecOk() (*IoArgoprojEventsV1alpha1SensorSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Sensor) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given IoArgoprojEventsV1alpha1SensorSpec and assigns it to the Spec field.
func (o *IoArgoprojEventsV1alpha1Sensor) SetSpec(v IoArgoprojEventsV1alpha1SensorSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Sensor) GetStatus() IoArgoprojEventsV1alpha1SensorStatus {
	if o == nil || o.Status == nil {
		var ret IoArgoprojEventsV1alpha1SensorStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Sensor) GetStatusOk() (*IoArgoprojEventsV1alpha1SensorStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Sensor) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IoArgoprojEventsV1alpha1SensorStatus and assigns it to the Status field.
func (o *IoArgoprojEventsV1alpha1Sensor) SetStatus(v IoArgoprojEventsV1alpha1SensorStatus) {
	o.Status = &v
}

func (o IoArgoprojEventsV1alpha1Sensor) MarshalJSON() ([]byte, error) {
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

type NullableIoArgoprojEventsV1alpha1Sensor struct {
	value *IoArgoprojEventsV1alpha1Sensor
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1Sensor) Get() *IoArgoprojEventsV1alpha1Sensor {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1Sensor) Set(val *IoArgoprojEventsV1alpha1Sensor) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1Sensor) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1Sensor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1Sensor(val *IoArgoprojEventsV1alpha1Sensor) *NullableIoArgoprojEventsV1alpha1Sensor {
	return &NullableIoArgoprojEventsV1alpha1Sensor{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1Sensor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1Sensor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


