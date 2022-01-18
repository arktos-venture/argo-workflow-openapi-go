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

// IoK8sApiCoreV1DownwardAPIVolumeFile DownwardAPIVolumeFile represents information to create the file containing the pod field
type IoK8sApiCoreV1DownwardAPIVolumeFile struct {
	FieldRef *IoK8sApiCoreV1ObjectFieldSelector `json:"fieldRef,omitempty"`
	// Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Mode *int32 `json:"mode,omitempty"`
	// Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	Path string `json:"path"`
	ResourceFieldRef *IoK8sApiCoreV1ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
}

// NewIoK8sApiCoreV1DownwardAPIVolumeFile instantiates a new IoK8sApiCoreV1DownwardAPIVolumeFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1DownwardAPIVolumeFile(path string) *IoK8sApiCoreV1DownwardAPIVolumeFile {
	this := IoK8sApiCoreV1DownwardAPIVolumeFile{}
	this.Path = path
	return &this
}

// NewIoK8sApiCoreV1DownwardAPIVolumeFileWithDefaults instantiates a new IoK8sApiCoreV1DownwardAPIVolumeFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1DownwardAPIVolumeFileWithDefaults() *IoK8sApiCoreV1DownwardAPIVolumeFile {
	this := IoK8sApiCoreV1DownwardAPIVolumeFile{}
	return &this
}

// GetFieldRef returns the FieldRef field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) GetFieldRef() IoK8sApiCoreV1ObjectFieldSelector {
	if o == nil || o.FieldRef == nil {
		var ret IoK8sApiCoreV1ObjectFieldSelector
		return ret
	}
	return *o.FieldRef
}

// GetFieldRefOk returns a tuple with the FieldRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) GetFieldRefOk() (*IoK8sApiCoreV1ObjectFieldSelector, bool) {
	if o == nil || o.FieldRef == nil {
		return nil, false
	}
	return o.FieldRef, true
}

// HasFieldRef returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) HasFieldRef() bool {
	if o != nil && o.FieldRef != nil {
		return true
	}

	return false
}

// SetFieldRef gets a reference to the given IoK8sApiCoreV1ObjectFieldSelector and assigns it to the FieldRef field.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) SetFieldRef(v IoK8sApiCoreV1ObjectFieldSelector) {
	o.FieldRef = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) GetMode() int32 {
	if o == nil || o.Mode == nil {
		var ret int32
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) GetModeOk() (*int32, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given int32 and assigns it to the Mode field.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) SetMode(v int32) {
	o.Mode = &v
}

// GetPath returns the Path field value
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) SetPath(v string) {
	o.Path = v
}

// GetResourceFieldRef returns the ResourceFieldRef field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) GetResourceFieldRef() IoK8sApiCoreV1ResourceFieldSelector {
	if o == nil || o.ResourceFieldRef == nil {
		var ret IoK8sApiCoreV1ResourceFieldSelector
		return ret
	}
	return *o.ResourceFieldRef
}

// GetResourceFieldRefOk returns a tuple with the ResourceFieldRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) GetResourceFieldRefOk() (*IoK8sApiCoreV1ResourceFieldSelector, bool) {
	if o == nil || o.ResourceFieldRef == nil {
		return nil, false
	}
	return o.ResourceFieldRef, true
}

// HasResourceFieldRef returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) HasResourceFieldRef() bool {
	if o != nil && o.ResourceFieldRef != nil {
		return true
	}

	return false
}

// SetResourceFieldRef gets a reference to the given IoK8sApiCoreV1ResourceFieldSelector and assigns it to the ResourceFieldRef field.
func (o *IoK8sApiCoreV1DownwardAPIVolumeFile) SetResourceFieldRef(v IoK8sApiCoreV1ResourceFieldSelector) {
	o.ResourceFieldRef = &v
}

func (o IoK8sApiCoreV1DownwardAPIVolumeFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FieldRef != nil {
		toSerialize["fieldRef"] = o.FieldRef
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.ResourceFieldRef != nil {
		toSerialize["resourceFieldRef"] = o.ResourceFieldRef
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1DownwardAPIVolumeFile struct {
	value *IoK8sApiCoreV1DownwardAPIVolumeFile
	isSet bool
}

func (v NullableIoK8sApiCoreV1DownwardAPIVolumeFile) Get() *IoK8sApiCoreV1DownwardAPIVolumeFile {
	return v.value
}

func (v *NullableIoK8sApiCoreV1DownwardAPIVolumeFile) Set(val *IoK8sApiCoreV1DownwardAPIVolumeFile) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1DownwardAPIVolumeFile) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1DownwardAPIVolumeFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1DownwardAPIVolumeFile(val *IoK8sApiCoreV1DownwardAPIVolumeFile) *NullableIoK8sApiCoreV1DownwardAPIVolumeFile {
	return &NullableIoK8sApiCoreV1DownwardAPIVolumeFile{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1DownwardAPIVolumeFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1DownwardAPIVolumeFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


