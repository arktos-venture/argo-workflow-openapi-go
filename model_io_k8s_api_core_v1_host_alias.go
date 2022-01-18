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

// IoK8sApiCoreV1HostAlias HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
type IoK8sApiCoreV1HostAlias struct {
	// Hostnames for the above IP address.
	Hostnames *[]string `json:"hostnames,omitempty"`
	// IP address of the host file entry.
	Ip *string `json:"ip,omitempty"`
}

// NewIoK8sApiCoreV1HostAlias instantiates a new IoK8sApiCoreV1HostAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1HostAlias() *IoK8sApiCoreV1HostAlias {
	this := IoK8sApiCoreV1HostAlias{}
	return &this
}

// NewIoK8sApiCoreV1HostAliasWithDefaults instantiates a new IoK8sApiCoreV1HostAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1HostAliasWithDefaults() *IoK8sApiCoreV1HostAlias {
	this := IoK8sApiCoreV1HostAlias{}
	return &this
}

// GetHostnames returns the Hostnames field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1HostAlias) GetHostnames() []string {
	if o == nil || o.Hostnames == nil {
		var ret []string
		return ret
	}
	return *o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1HostAlias) GetHostnamesOk() (*[]string, bool) {
	if o == nil || o.Hostnames == nil {
		return nil, false
	}
	return o.Hostnames, true
}

// HasHostnames returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1HostAlias) HasHostnames() bool {
	if o != nil && o.Hostnames != nil {
		return true
	}

	return false
}

// SetHostnames gets a reference to the given []string and assigns it to the Hostnames field.
func (o *IoK8sApiCoreV1HostAlias) SetHostnames(v []string) {
	o.Hostnames = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1HostAlias) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1HostAlias) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1HostAlias) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *IoK8sApiCoreV1HostAlias) SetIp(v string) {
	o.Ip = &v
}

func (o IoK8sApiCoreV1HostAlias) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hostnames != nil {
		toSerialize["hostnames"] = o.Hostnames
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1HostAlias struct {
	value *IoK8sApiCoreV1HostAlias
	isSet bool
}

func (v NullableIoK8sApiCoreV1HostAlias) Get() *IoK8sApiCoreV1HostAlias {
	return v.value
}

func (v *NullableIoK8sApiCoreV1HostAlias) Set(val *IoK8sApiCoreV1HostAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1HostAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1HostAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1HostAlias(val *IoK8sApiCoreV1HostAlias) *NullableIoK8sApiCoreV1HostAlias {
	return &NullableIoK8sApiCoreV1HostAlias{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1HostAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1HostAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


