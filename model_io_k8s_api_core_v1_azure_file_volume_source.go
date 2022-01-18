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

// IoK8sApiCoreV1AzureFileVolumeSource AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type IoK8sApiCoreV1AzureFileVolumeSource struct {
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// the name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName"`
	// Share Name
	ShareName string `json:"shareName"`
}

// NewIoK8sApiCoreV1AzureFileVolumeSource instantiates a new IoK8sApiCoreV1AzureFileVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1AzureFileVolumeSource(secretName string, shareName string) *IoK8sApiCoreV1AzureFileVolumeSource {
	this := IoK8sApiCoreV1AzureFileVolumeSource{}
	this.SecretName = secretName
	this.ShareName = shareName
	return &this
}

// NewIoK8sApiCoreV1AzureFileVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1AzureFileVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1AzureFileVolumeSourceWithDefaults() *IoK8sApiCoreV1AzureFileVolumeSource {
	this := IoK8sApiCoreV1AzureFileVolumeSource{}
	return &this
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1AzureFileVolumeSource) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1AzureFileVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1AzureFileVolumeSource) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *IoK8sApiCoreV1AzureFileVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetSecretName returns the SecretName field value
func (o *IoK8sApiCoreV1AzureFileVolumeSource) GetSecretName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1AzureFileVolumeSource) GetSecretNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecretName, true
}

// SetSecretName sets field value
func (o *IoK8sApiCoreV1AzureFileVolumeSource) SetSecretName(v string) {
	o.SecretName = v
}

// GetShareName returns the ShareName field value
func (o *IoK8sApiCoreV1AzureFileVolumeSource) GetShareName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShareName
}

// GetShareNameOk returns a tuple with the ShareName field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1AzureFileVolumeSource) GetShareNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ShareName, true
}

// SetShareName sets field value
func (o *IoK8sApiCoreV1AzureFileVolumeSource) SetShareName(v string) {
	o.ShareName = v
}

func (o IoK8sApiCoreV1AzureFileVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if true {
		toSerialize["secretName"] = o.SecretName
	}
	if true {
		toSerialize["shareName"] = o.ShareName
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1AzureFileVolumeSource struct {
	value *IoK8sApiCoreV1AzureFileVolumeSource
	isSet bool
}

func (v NullableIoK8sApiCoreV1AzureFileVolumeSource) Get() *IoK8sApiCoreV1AzureFileVolumeSource {
	return v.value
}

func (v *NullableIoK8sApiCoreV1AzureFileVolumeSource) Set(val *IoK8sApiCoreV1AzureFileVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1AzureFileVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1AzureFileVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1AzureFileVolumeSource(val *IoK8sApiCoreV1AzureFileVolumeSource) *NullableIoK8sApiCoreV1AzureFileVolumeSource {
	return &NullableIoK8sApiCoreV1AzureFileVolumeSource{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1AzureFileVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1AzureFileVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


