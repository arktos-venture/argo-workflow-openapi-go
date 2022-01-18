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

// IoK8sApiCoreV1TCPSocketAction TCPSocketAction describes an action based on opening a socket
type IoK8sApiCoreV1TCPSocketAction struct {
	// Optional: Host name to connect to, defaults to the pod IP.
	Host *string `json:"host,omitempty"`
	Port string `json:"port"`
}

// NewIoK8sApiCoreV1TCPSocketAction instantiates a new IoK8sApiCoreV1TCPSocketAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1TCPSocketAction(port string) *IoK8sApiCoreV1TCPSocketAction {
	this := IoK8sApiCoreV1TCPSocketAction{}
	this.Port = port
	return &this
}

// NewIoK8sApiCoreV1TCPSocketActionWithDefaults instantiates a new IoK8sApiCoreV1TCPSocketAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1TCPSocketActionWithDefaults() *IoK8sApiCoreV1TCPSocketAction {
	this := IoK8sApiCoreV1TCPSocketAction{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1TCPSocketAction) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1TCPSocketAction) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1TCPSocketAction) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *IoK8sApiCoreV1TCPSocketAction) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value
func (o *IoK8sApiCoreV1TCPSocketAction) GetPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1TCPSocketAction) GetPortOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *IoK8sApiCoreV1TCPSocketAction) SetPort(v string) {
	o.Port = v
}

func (o IoK8sApiCoreV1TCPSocketAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1TCPSocketAction struct {
	value *IoK8sApiCoreV1TCPSocketAction
	isSet bool
}

func (v NullableIoK8sApiCoreV1TCPSocketAction) Get() *IoK8sApiCoreV1TCPSocketAction {
	return v.value
}

func (v *NullableIoK8sApiCoreV1TCPSocketAction) Set(val *IoK8sApiCoreV1TCPSocketAction) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1TCPSocketAction) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1TCPSocketAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1TCPSocketAction(val *IoK8sApiCoreV1TCPSocketAction) *NullableIoK8sApiCoreV1TCPSocketAction {
	return &NullableIoK8sApiCoreV1TCPSocketAction{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1TCPSocketAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1TCPSocketAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


