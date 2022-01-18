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

// IoK8sApiCoreV1AzureDiskVolumeSource AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
type IoK8sApiCoreV1AzureDiskVolumeSource struct {
	// Host Caching mode: None, Read Only, Read Write.
	CachingMode *string `json:"cachingMode,omitempty"`
	// The Name of the data disk in the blob storage
	DiskName string `json:"diskName"`
	// The URI the data disk in the blob storage
	DiskURI string `json:"diskURI"`
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared
	Kind *string `json:"kind,omitempty"`
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
}

// NewIoK8sApiCoreV1AzureDiskVolumeSource instantiates a new IoK8sApiCoreV1AzureDiskVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1AzureDiskVolumeSource(diskName string, diskURI string) *IoK8sApiCoreV1AzureDiskVolumeSource {
	this := IoK8sApiCoreV1AzureDiskVolumeSource{}
	this.DiskName = diskName
	this.DiskURI = diskURI
	return &this
}

// NewIoK8sApiCoreV1AzureDiskVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1AzureDiskVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1AzureDiskVolumeSourceWithDefaults() *IoK8sApiCoreV1AzureDiskVolumeSource {
	this := IoK8sApiCoreV1AzureDiskVolumeSource{}
	return &this
}

// GetCachingMode returns the CachingMode field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetCachingMode() string {
	if o == nil || o.CachingMode == nil {
		var ret string
		return ret
	}
	return *o.CachingMode
}

// GetCachingModeOk returns a tuple with the CachingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetCachingModeOk() (*string, bool) {
	if o == nil || o.CachingMode == nil {
		return nil, false
	}
	return o.CachingMode, true
}

// HasCachingMode returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) HasCachingMode() bool {
	if o != nil && o.CachingMode != nil {
		return true
	}

	return false
}

// SetCachingMode gets a reference to the given string and assigns it to the CachingMode field.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) SetCachingMode(v string) {
	o.CachingMode = &v
}

// GetDiskName returns the DiskName field value
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetDiskName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiskName
}

// GetDiskNameOk returns a tuple with the DiskName field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetDiskNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DiskName, true
}

// SetDiskName sets field value
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) SetDiskName(v string) {
	o.DiskName = v
}

// GetDiskURI returns the DiskURI field value
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetDiskURI() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiskURI
}

// GetDiskURIOk returns a tuple with the DiskURI field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetDiskURIOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DiskURI, true
}

// SetDiskURI sets field value
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) SetDiskURI(v string) {
	o.DiskURI = v
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetFsType() string {
	if o == nil || o.FsType == nil {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || o.FsType == nil {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) HasFsType() bool {
	if o != nil && o.FsType != nil {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) SetKind(v string) {
	o.Kind = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *IoK8sApiCoreV1AzureDiskVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o IoK8sApiCoreV1AzureDiskVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CachingMode != nil {
		toSerialize["cachingMode"] = o.CachingMode
	}
	if true {
		toSerialize["diskName"] = o.DiskName
	}
	if true {
		toSerialize["diskURI"] = o.DiskURI
	}
	if o.FsType != nil {
		toSerialize["fsType"] = o.FsType
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1AzureDiskVolumeSource struct {
	value *IoK8sApiCoreV1AzureDiskVolumeSource
	isSet bool
}

func (v NullableIoK8sApiCoreV1AzureDiskVolumeSource) Get() *IoK8sApiCoreV1AzureDiskVolumeSource {
	return v.value
}

func (v *NullableIoK8sApiCoreV1AzureDiskVolumeSource) Set(val *IoK8sApiCoreV1AzureDiskVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1AzureDiskVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1AzureDiskVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1AzureDiskVolumeSource(val *IoK8sApiCoreV1AzureDiskVolumeSource) *NullableIoK8sApiCoreV1AzureDiskVolumeSource {
	return &NullableIoK8sApiCoreV1AzureDiskVolumeSource{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1AzureDiskVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1AzureDiskVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


