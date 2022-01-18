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

// IoK8sApiCoreV1PersistentVolumeClaim PersistentVolumeClaim is a user's request for and claim to a persistent volume
type IoK8sApiCoreV1PersistentVolumeClaim struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata,omitempty"`
	Spec *IoK8sApiCoreV1PersistentVolumeClaimSpec `json:"spec,omitempty"`
	Status *IoK8sApiCoreV1PersistentVolumeClaimStatus `json:"status,omitempty"`
}

// NewIoK8sApiCoreV1PersistentVolumeClaim instantiates a new IoK8sApiCoreV1PersistentVolumeClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1PersistentVolumeClaim() *IoK8sApiCoreV1PersistentVolumeClaim {
	this := IoK8sApiCoreV1PersistentVolumeClaim{}
	return &this
}

// NewIoK8sApiCoreV1PersistentVolumeClaimWithDefaults instantiates a new IoK8sApiCoreV1PersistentVolumeClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1PersistentVolumeClaimWithDefaults() *IoK8sApiCoreV1PersistentVolumeClaim {
	this := IoK8sApiCoreV1PersistentVolumeClaim{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret IoK8sApimachineryPkgApisMetaV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given IoK8sApimachineryPkgApisMetaV1ObjectMeta and assigns it to the Metadata field.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) GetSpec() IoK8sApiCoreV1PersistentVolumeClaimSpec {
	if o == nil || o.Spec == nil {
		var ret IoK8sApiCoreV1PersistentVolumeClaimSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) GetSpecOk() (*IoK8sApiCoreV1PersistentVolumeClaimSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given IoK8sApiCoreV1PersistentVolumeClaimSpec and assigns it to the Spec field.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) SetSpec(v IoK8sApiCoreV1PersistentVolumeClaimSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) GetStatus() IoK8sApiCoreV1PersistentVolumeClaimStatus {
	if o == nil || o.Status == nil {
		var ret IoK8sApiCoreV1PersistentVolumeClaimStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) GetStatusOk() (*IoK8sApiCoreV1PersistentVolumeClaimStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IoK8sApiCoreV1PersistentVolumeClaimStatus and assigns it to the Status field.
func (o *IoK8sApiCoreV1PersistentVolumeClaim) SetStatus(v IoK8sApiCoreV1PersistentVolumeClaimStatus) {
	o.Status = &v
}

func (o IoK8sApiCoreV1PersistentVolumeClaim) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
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

type NullableIoK8sApiCoreV1PersistentVolumeClaim struct {
	value *IoK8sApiCoreV1PersistentVolumeClaim
	isSet bool
}

func (v NullableIoK8sApiCoreV1PersistentVolumeClaim) Get() *IoK8sApiCoreV1PersistentVolumeClaim {
	return v.value
}

func (v *NullableIoK8sApiCoreV1PersistentVolumeClaim) Set(val *IoK8sApiCoreV1PersistentVolumeClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1PersistentVolumeClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1PersistentVolumeClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1PersistentVolumeClaim(val *IoK8sApiCoreV1PersistentVolumeClaim) *NullableIoK8sApiCoreV1PersistentVolumeClaim {
	return &NullableIoK8sApiCoreV1PersistentVolumeClaim{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1PersistentVolumeClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1PersistentVolumeClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

