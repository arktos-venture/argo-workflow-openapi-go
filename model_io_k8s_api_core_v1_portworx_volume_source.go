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

// IoK8sApiCoreV1PortworxVolumeSource PortworxVolumeSource represents a Portworx volume resource.
type IoK8sApiCoreV1PortworxVolumeSource struct {
	// FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\". Implicitly inferred to be \"ext4\" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// VolumeID uniquely identifies a Portworx volume
	VolumeID string `json:"volumeID"`
}

// NewIoK8sApiCoreV1PortworxVolumeSource instantiates a new IoK8sApiCoreV1PortworxVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1PortworxVolumeSource(volumeID string) *IoK8sApiCoreV1PortworxVolumeSource {
	this := IoK8sApiCoreV1PortworxVolumeSource{}
	this.VolumeID = volumeID
	return &this
}

// NewIoK8sApiCoreV1PortworxVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1PortworxVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1PortworxVolumeSourceWithDefaults() *IoK8sApiCoreV1PortworxVolumeSource {
	this := IoK8sApiCoreV1PortworxVolumeSource{}
	return &this
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1PortworxVolumeSource) GetFsType() string {
	if o == nil || o.FsType == nil {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PortworxVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || o.FsType == nil {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1PortworxVolumeSource) HasFsType() bool {
	if o != nil && o.FsType != nil {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *IoK8sApiCoreV1PortworxVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1PortworxVolumeSource) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PortworxVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1PortworxVolumeSource) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *IoK8sApiCoreV1PortworxVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetVolumeID returns the VolumeID field value
func (o *IoK8sApiCoreV1PortworxVolumeSource) GetVolumeID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeID
}

// GetVolumeIDOk returns a tuple with the VolumeID field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PortworxVolumeSource) GetVolumeIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VolumeID, true
}

// SetVolumeID sets field value
func (o *IoK8sApiCoreV1PortworxVolumeSource) SetVolumeID(v string) {
	o.VolumeID = v
}

func (o IoK8sApiCoreV1PortworxVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FsType != nil {
		toSerialize["fsType"] = o.FsType
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if true {
		toSerialize["volumeID"] = o.VolumeID
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1PortworxVolumeSource struct {
	value *IoK8sApiCoreV1PortworxVolumeSource
	isSet bool
}

func (v NullableIoK8sApiCoreV1PortworxVolumeSource) Get() *IoK8sApiCoreV1PortworxVolumeSource {
	return v.value
}

func (v *NullableIoK8sApiCoreV1PortworxVolumeSource) Set(val *IoK8sApiCoreV1PortworxVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1PortworxVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1PortworxVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1PortworxVolumeSource(val *IoK8sApiCoreV1PortworxVolumeSource) *NullableIoK8sApiCoreV1PortworxVolumeSource {
	return &NullableIoK8sApiCoreV1PortworxVolumeSource{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1PortworxVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1PortworxVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


