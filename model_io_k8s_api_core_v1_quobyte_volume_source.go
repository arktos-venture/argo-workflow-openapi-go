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

// IoK8sApiCoreV1QuobyteVolumeSource Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.
type IoK8sApiCoreV1QuobyteVolumeSource struct {
	// Group to map volume access to Default is no group
	Group *string `json:"group,omitempty"`
	// ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes
	Registry string `json:"registry"`
	// Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin
	Tenant *string `json:"tenant,omitempty"`
	// User to map volume access to Defaults to serivceaccount user
	User *string `json:"user,omitempty"`
	// Volume is a string that references an already created Quobyte volume by name.
	Volume string `json:"volume"`
}

// NewIoK8sApiCoreV1QuobyteVolumeSource instantiates a new IoK8sApiCoreV1QuobyteVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1QuobyteVolumeSource(registry string, volume string) *IoK8sApiCoreV1QuobyteVolumeSource {
	this := IoK8sApiCoreV1QuobyteVolumeSource{}
	this.Registry = registry
	this.Volume = volume
	return &this
}

// NewIoK8sApiCoreV1QuobyteVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1QuobyteVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1QuobyteVolumeSourceWithDefaults() *IoK8sApiCoreV1QuobyteVolumeSource {
	this := IoK8sApiCoreV1QuobyteVolumeSource{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) SetGroup(v string) {
	o.Group = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetRegistry returns the Registry field value
func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetRegistry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetRegistryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Registry, true
}

// SetRegistry sets field value
func (o *IoK8sApiCoreV1QuobyteVolumeSource) SetRegistry(v string) {
	o.Registry = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) SetTenant(v string) {
	o.Tenant = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) SetUser(v string) {
	o.User = &v
}

// GetVolume returns the Volume field value
func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetVolume() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetVolumeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Volume, true
}

// SetVolume sets field value
func (o *IoK8sApiCoreV1QuobyteVolumeSource) SetVolume(v string) {
	o.Volume = v
}

func (o IoK8sApiCoreV1QuobyteVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if true {
		toSerialize["registry"] = o.Registry
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["volume"] = o.Volume
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1QuobyteVolumeSource struct {
	value *IoK8sApiCoreV1QuobyteVolumeSource
	isSet bool
}

func (v NullableIoK8sApiCoreV1QuobyteVolumeSource) Get() *IoK8sApiCoreV1QuobyteVolumeSource {
	return v.value
}

func (v *NullableIoK8sApiCoreV1QuobyteVolumeSource) Set(val *IoK8sApiCoreV1QuobyteVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1QuobyteVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1QuobyteVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1QuobyteVolumeSource(val *IoK8sApiCoreV1QuobyteVolumeSource) *NullableIoK8sApiCoreV1QuobyteVolumeSource {
	return &NullableIoK8sApiCoreV1QuobyteVolumeSource{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1QuobyteVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1QuobyteVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


