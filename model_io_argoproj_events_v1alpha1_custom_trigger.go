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

// IoArgoprojEventsV1alpha1CustomTrigger CustomTrigger refers to the specification of the custom trigger.
type IoArgoprojEventsV1alpha1CustomTrigger struct {
	CertFilePath *string `json:"certFilePath,omitempty"`
	CertSecret *IoK8sApiCoreV1SecretKeySelector `json:"certSecret,omitempty"`
	// Parameters is the list of parameters that is applied to resolved custom trigger trigger object.
	Parameters *[]IoArgoprojEventsV1alpha1TriggerParameter `json:"parameters,omitempty"`
	// Payload is the list of key-value extracted from an event payload to construct the request payload.
	Payload *[]IoArgoprojEventsV1alpha1TriggerParameter `json:"payload,omitempty"`
	Secure *bool `json:"secure,omitempty"`
	// ServerNameOverride for the secure connection between sensor and custom trigger gRPC server.
	ServerNameOverride *string `json:"serverNameOverride,omitempty"`
	ServerURL *string `json:"serverURL,omitempty"`
	// Spec is the custom trigger resource specification that custom trigger gRPC server knows how to interpret.
	Spec *map[string]string `json:"spec,omitempty"`
}

// NewIoArgoprojEventsV1alpha1CustomTrigger instantiates a new IoArgoprojEventsV1alpha1CustomTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1CustomTrigger() *IoArgoprojEventsV1alpha1CustomTrigger {
	this := IoArgoprojEventsV1alpha1CustomTrigger{}
	return &this
}

// NewIoArgoprojEventsV1alpha1CustomTriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1CustomTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1CustomTriggerWithDefaults() *IoArgoprojEventsV1alpha1CustomTrigger {
	this := IoArgoprojEventsV1alpha1CustomTrigger{}
	return &this
}

// GetCertFilePath returns the CertFilePath field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetCertFilePath() string {
	if o == nil || o.CertFilePath == nil {
		var ret string
		return ret
	}
	return *o.CertFilePath
}

// GetCertFilePathOk returns a tuple with the CertFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetCertFilePathOk() (*string, bool) {
	if o == nil || o.CertFilePath == nil {
		return nil, false
	}
	return o.CertFilePath, true
}

// HasCertFilePath returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasCertFilePath() bool {
	if o != nil && o.CertFilePath != nil {
		return true
	}

	return false
}

// SetCertFilePath gets a reference to the given string and assigns it to the CertFilePath field.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetCertFilePath(v string) {
	o.CertFilePath = &v
}

// GetCertSecret returns the CertSecret field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetCertSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.CertSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.CertSecret
}

// GetCertSecretOk returns a tuple with the CertSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetCertSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.CertSecret == nil {
		return nil, false
	}
	return o.CertSecret, true
}

// HasCertSecret returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasCertSecret() bool {
	if o != nil && o.CertSecret != nil {
		return true
	}

	return false
}

// SetCertSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the CertSecret field.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetCertSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.CertSecret = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter {
	if o == nil || o.Parameters == nil {
		var ret []IoArgoprojEventsV1alpha1TriggerParameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []IoArgoprojEventsV1alpha1TriggerParameter and assigns it to the Parameters field.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter) {
	o.Parameters = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetPayload() []IoArgoprojEventsV1alpha1TriggerParameter {
	if o == nil || o.Payload == nil {
		var ret []IoArgoprojEventsV1alpha1TriggerParameter
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetPayloadOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given []IoArgoprojEventsV1alpha1TriggerParameter and assigns it to the Payload field.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetPayload(v []IoArgoprojEventsV1alpha1TriggerParameter) {
	o.Payload = &v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetSecure() bool {
	if o == nil || o.Secure == nil {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetSecureOk() (*bool, bool) {
	if o == nil || o.Secure == nil {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasSecure() bool {
	if o != nil && o.Secure != nil {
		return true
	}

	return false
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetSecure(v bool) {
	o.Secure = &v
}

// GetServerNameOverride returns the ServerNameOverride field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetServerNameOverride() string {
	if o == nil || o.ServerNameOverride == nil {
		var ret string
		return ret
	}
	return *o.ServerNameOverride
}

// GetServerNameOverrideOk returns a tuple with the ServerNameOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetServerNameOverrideOk() (*string, bool) {
	if o == nil || o.ServerNameOverride == nil {
		return nil, false
	}
	return o.ServerNameOverride, true
}

// HasServerNameOverride returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasServerNameOverride() bool {
	if o != nil && o.ServerNameOverride != nil {
		return true
	}

	return false
}

// SetServerNameOverride gets a reference to the given string and assigns it to the ServerNameOverride field.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetServerNameOverride(v string) {
	o.ServerNameOverride = &v
}

// GetServerURL returns the ServerURL field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetServerURL() string {
	if o == nil || o.ServerURL == nil {
		var ret string
		return ret
	}
	return *o.ServerURL
}

// GetServerURLOk returns a tuple with the ServerURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetServerURLOk() (*string, bool) {
	if o == nil || o.ServerURL == nil {
		return nil, false
	}
	return o.ServerURL, true
}

// HasServerURL returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasServerURL() bool {
	if o != nil && o.ServerURL != nil {
		return true
	}

	return false
}

// SetServerURL gets a reference to the given string and assigns it to the ServerURL field.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetServerURL(v string) {
	o.ServerURL = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetSpec() map[string]string {
	if o == nil || o.Spec == nil {
		var ret map[string]string
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetSpecOk() (*map[string]string, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given map[string]string and assigns it to the Spec field.
func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetSpec(v map[string]string) {
	o.Spec = &v
}

func (o IoArgoprojEventsV1alpha1CustomTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CertFilePath != nil {
		toSerialize["certFilePath"] = o.CertFilePath
	}
	if o.CertSecret != nil {
		toSerialize["certSecret"] = o.CertSecret
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.Secure != nil {
		toSerialize["secure"] = o.Secure
	}
	if o.ServerNameOverride != nil {
		toSerialize["serverNameOverride"] = o.ServerNameOverride
	}
	if o.ServerURL != nil {
		toSerialize["serverURL"] = o.ServerURL
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1CustomTrigger struct {
	value *IoArgoprojEventsV1alpha1CustomTrigger
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1CustomTrigger) Get() *IoArgoprojEventsV1alpha1CustomTrigger {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1CustomTrigger) Set(val *IoArgoprojEventsV1alpha1CustomTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1CustomTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1CustomTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1CustomTrigger(val *IoArgoprojEventsV1alpha1CustomTrigger) *NullableIoArgoprojEventsV1alpha1CustomTrigger {
	return &NullableIoArgoprojEventsV1alpha1CustomTrigger{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1CustomTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1CustomTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


