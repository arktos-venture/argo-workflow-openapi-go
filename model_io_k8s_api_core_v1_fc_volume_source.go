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

// IoK8sApiCoreV1FCVolumeSource Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
type IoK8sApiCoreV1FCVolumeSource struct {
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// Optional: FC target lun number
	Lun *int32 `json:"lun,omitempty"`
	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// Optional: FC target worldwide names (WWNs)
	TargetWWNs *[]string `json:"targetWWNs,omitempty"`
	// Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
	Wwids *[]string `json:"wwids,omitempty"`
}

// NewIoK8sApiCoreV1FCVolumeSource instantiates a new IoK8sApiCoreV1FCVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1FCVolumeSource() *IoK8sApiCoreV1FCVolumeSource {
	this := IoK8sApiCoreV1FCVolumeSource{}
	return &this
}

// NewIoK8sApiCoreV1FCVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1FCVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1FCVolumeSourceWithDefaults() *IoK8sApiCoreV1FCVolumeSource {
	this := IoK8sApiCoreV1FCVolumeSource{}
	return &this
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1FCVolumeSource) GetFsType() string {
	if o == nil || o.FsType == nil {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1FCVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || o.FsType == nil {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1FCVolumeSource) HasFsType() bool {
	if o != nil && o.FsType != nil {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *IoK8sApiCoreV1FCVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetLun returns the Lun field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1FCVolumeSource) GetLun() int32 {
	if o == nil || o.Lun == nil {
		var ret int32
		return ret
	}
	return *o.Lun
}

// GetLunOk returns a tuple with the Lun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1FCVolumeSource) GetLunOk() (*int32, bool) {
	if o == nil || o.Lun == nil {
		return nil, false
	}
	return o.Lun, true
}

// HasLun returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1FCVolumeSource) HasLun() bool {
	if o != nil && o.Lun != nil {
		return true
	}

	return false
}

// SetLun gets a reference to the given int32 and assigns it to the Lun field.
func (o *IoK8sApiCoreV1FCVolumeSource) SetLun(v int32) {
	o.Lun = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1FCVolumeSource) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1FCVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1FCVolumeSource) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *IoK8sApiCoreV1FCVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetTargetWWNs returns the TargetWWNs field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1FCVolumeSource) GetTargetWWNs() []string {
	if o == nil || o.TargetWWNs == nil {
		var ret []string
		return ret
	}
	return *o.TargetWWNs
}

// GetTargetWWNsOk returns a tuple with the TargetWWNs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1FCVolumeSource) GetTargetWWNsOk() (*[]string, bool) {
	if o == nil || o.TargetWWNs == nil {
		return nil, false
	}
	return o.TargetWWNs, true
}

// HasTargetWWNs returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1FCVolumeSource) HasTargetWWNs() bool {
	if o != nil && o.TargetWWNs != nil {
		return true
	}

	return false
}

// SetTargetWWNs gets a reference to the given []string and assigns it to the TargetWWNs field.
func (o *IoK8sApiCoreV1FCVolumeSource) SetTargetWWNs(v []string) {
	o.TargetWWNs = &v
}

// GetWwids returns the Wwids field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1FCVolumeSource) GetWwids() []string {
	if o == nil || o.Wwids == nil {
		var ret []string
		return ret
	}
	return *o.Wwids
}

// GetWwidsOk returns a tuple with the Wwids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1FCVolumeSource) GetWwidsOk() (*[]string, bool) {
	if o == nil || o.Wwids == nil {
		return nil, false
	}
	return o.Wwids, true
}

// HasWwids returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1FCVolumeSource) HasWwids() bool {
	if o != nil && o.Wwids != nil {
		return true
	}

	return false
}

// SetWwids gets a reference to the given []string and assigns it to the Wwids field.
func (o *IoK8sApiCoreV1FCVolumeSource) SetWwids(v []string) {
	o.Wwids = &v
}

func (o IoK8sApiCoreV1FCVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FsType != nil {
		toSerialize["fsType"] = o.FsType
	}
	if o.Lun != nil {
		toSerialize["lun"] = o.Lun
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.TargetWWNs != nil {
		toSerialize["targetWWNs"] = o.TargetWWNs
	}
	if o.Wwids != nil {
		toSerialize["wwids"] = o.Wwids
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1FCVolumeSource struct {
	value *IoK8sApiCoreV1FCVolumeSource
	isSet bool
}

func (v NullableIoK8sApiCoreV1FCVolumeSource) Get() *IoK8sApiCoreV1FCVolumeSource {
	return v.value
}

func (v *NullableIoK8sApiCoreV1FCVolumeSource) Set(val *IoK8sApiCoreV1FCVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1FCVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1FCVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1FCVolumeSource(val *IoK8sApiCoreV1FCVolumeSource) *NullableIoK8sApiCoreV1FCVolumeSource {
	return &NullableIoK8sApiCoreV1FCVolumeSource{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1FCVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1FCVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


