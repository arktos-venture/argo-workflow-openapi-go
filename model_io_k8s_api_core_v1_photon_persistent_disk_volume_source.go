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

// IoK8sApiCoreV1PhotonPersistentDiskVolumeSource Represents a Photon Controller persistent disk resource.
type IoK8sApiCoreV1PhotonPersistentDiskVolumeSource struct {
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// ID that identifies Photon Controller persistent disk
	PdID string `json:"pdID"`
}

// NewIoK8sApiCoreV1PhotonPersistentDiskVolumeSource instantiates a new IoK8sApiCoreV1PhotonPersistentDiskVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1PhotonPersistentDiskVolumeSource(pdID string) *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource {
	this := IoK8sApiCoreV1PhotonPersistentDiskVolumeSource{}
	this.PdID = pdID
	return &this
}

// NewIoK8sApiCoreV1PhotonPersistentDiskVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1PhotonPersistentDiskVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1PhotonPersistentDiskVolumeSourceWithDefaults() *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource {
	this := IoK8sApiCoreV1PhotonPersistentDiskVolumeSource{}
	return &this
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource) GetFsType() string {
	if o == nil || o.FsType == nil {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || o.FsType == nil {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource) HasFsType() bool {
	if o != nil && o.FsType != nil {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetPdID returns the PdID field value
func (o *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource) GetPdID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PdID
}

// GetPdIDOk returns a tuple with the PdID field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource) GetPdIDOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PdID, true
}

// SetPdID sets field value
func (o *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource) SetPdID(v string) {
	o.PdID = v
}

func (o IoK8sApiCoreV1PhotonPersistentDiskVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FsType != nil {
		toSerialize["fsType"] = o.FsType
	}
	if true {
		toSerialize["pdID"] = o.PdID
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1PhotonPersistentDiskVolumeSource struct {
	value *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource
	isSet bool
}

func (v NullableIoK8sApiCoreV1PhotonPersistentDiskVolumeSource) Get() *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource {
	return v.value
}

func (v *NullableIoK8sApiCoreV1PhotonPersistentDiskVolumeSource) Set(val *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1PhotonPersistentDiskVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1PhotonPersistentDiskVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1PhotonPersistentDiskVolumeSource(val *IoK8sApiCoreV1PhotonPersistentDiskVolumeSource) *NullableIoK8sApiCoreV1PhotonPersistentDiskVolumeSource {
	return &NullableIoK8sApiCoreV1PhotonPersistentDiskVolumeSource{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1PhotonPersistentDiskVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1PhotonPersistentDiskVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


