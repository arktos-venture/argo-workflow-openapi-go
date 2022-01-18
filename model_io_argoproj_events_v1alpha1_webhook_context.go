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

// IoArgoprojEventsV1alpha1WebhookContext struct for IoArgoprojEventsV1alpha1WebhookContext
type IoArgoprojEventsV1alpha1WebhookContext struct {
	AuthSecret *IoK8sApiCoreV1SecretKeySelector `json:"authSecret,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	Method *string `json:"method,omitempty"`
	// Port on which HTTP server is listening for incoming events.
	Port *string `json:"port,omitempty"`
	// DeprecatedServerCertPath refers the file that contains the cert.
	ServerCertPath *string `json:"serverCertPath,omitempty"`
	ServerCertSecret *IoK8sApiCoreV1SecretKeySelector `json:"serverCertSecret,omitempty"`
	ServerKeyPath *string `json:"serverKeyPath,omitempty"`
	ServerKeySecret *IoK8sApiCoreV1SecretKeySelector `json:"serverKeySecret,omitempty"`
	// URL is the url of the server.
	Url *string `json:"url,omitempty"`
}

// NewIoArgoprojEventsV1alpha1WebhookContext instantiates a new IoArgoprojEventsV1alpha1WebhookContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1WebhookContext() *IoArgoprojEventsV1alpha1WebhookContext {
	this := IoArgoprojEventsV1alpha1WebhookContext{}
	return &this
}

// NewIoArgoprojEventsV1alpha1WebhookContextWithDefaults instantiates a new IoArgoprojEventsV1alpha1WebhookContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1WebhookContextWithDefaults() *IoArgoprojEventsV1alpha1WebhookContext {
	this := IoArgoprojEventsV1alpha1WebhookContext{}
	return &this
}

// GetAuthSecret returns the AuthSecret field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetAuthSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.AuthSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.AuthSecret
}

// GetAuthSecretOk returns a tuple with the AuthSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetAuthSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.AuthSecret == nil {
		return nil, false
	}
	return o.AuthSecret, true
}

// HasAuthSecret returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) HasAuthSecret() bool {
	if o != nil && o.AuthSecret != nil {
		return true
	}

	return false
}

// SetAuthSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the AuthSecret field.
func (o *IoArgoprojEventsV1alpha1WebhookContext) SetAuthSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.AuthSecret = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *IoArgoprojEventsV1alpha1WebhookContext) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *IoArgoprojEventsV1alpha1WebhookContext) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *IoArgoprojEventsV1alpha1WebhookContext) SetMethod(v string) {
	o.Method = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *IoArgoprojEventsV1alpha1WebhookContext) SetPort(v string) {
	o.Port = &v
}

// GetServerCertPath returns the ServerCertPath field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerCertPath() string {
	if o == nil || o.ServerCertPath == nil {
		var ret string
		return ret
	}
	return *o.ServerCertPath
}

// GetServerCertPathOk returns a tuple with the ServerCertPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerCertPathOk() (*string, bool) {
	if o == nil || o.ServerCertPath == nil {
		return nil, false
	}
	return o.ServerCertPath, true
}

// HasServerCertPath returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) HasServerCertPath() bool {
	if o != nil && o.ServerCertPath != nil {
		return true
	}

	return false
}

// SetServerCertPath gets a reference to the given string and assigns it to the ServerCertPath field.
func (o *IoArgoprojEventsV1alpha1WebhookContext) SetServerCertPath(v string) {
	o.ServerCertPath = &v
}

// GetServerCertSecret returns the ServerCertSecret field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerCertSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.ServerCertSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.ServerCertSecret
}

// GetServerCertSecretOk returns a tuple with the ServerCertSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerCertSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.ServerCertSecret == nil {
		return nil, false
	}
	return o.ServerCertSecret, true
}

// HasServerCertSecret returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) HasServerCertSecret() bool {
	if o != nil && o.ServerCertSecret != nil {
		return true
	}

	return false
}

// SetServerCertSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the ServerCertSecret field.
func (o *IoArgoprojEventsV1alpha1WebhookContext) SetServerCertSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.ServerCertSecret = &v
}

// GetServerKeyPath returns the ServerKeyPath field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerKeyPath() string {
	if o == nil || o.ServerKeyPath == nil {
		var ret string
		return ret
	}
	return *o.ServerKeyPath
}

// GetServerKeyPathOk returns a tuple with the ServerKeyPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerKeyPathOk() (*string, bool) {
	if o == nil || o.ServerKeyPath == nil {
		return nil, false
	}
	return o.ServerKeyPath, true
}

// HasServerKeyPath returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) HasServerKeyPath() bool {
	if o != nil && o.ServerKeyPath != nil {
		return true
	}

	return false
}

// SetServerKeyPath gets a reference to the given string and assigns it to the ServerKeyPath field.
func (o *IoArgoprojEventsV1alpha1WebhookContext) SetServerKeyPath(v string) {
	o.ServerKeyPath = &v
}

// GetServerKeySecret returns the ServerKeySecret field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerKeySecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.ServerKeySecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.ServerKeySecret
}

// GetServerKeySecretOk returns a tuple with the ServerKeySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.ServerKeySecret == nil {
		return nil, false
	}
	return o.ServerKeySecret, true
}

// HasServerKeySecret returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) HasServerKeySecret() bool {
	if o != nil && o.ServerKeySecret != nil {
		return true
	}

	return false
}

// SetServerKeySecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the ServerKeySecret field.
func (o *IoArgoprojEventsV1alpha1WebhookContext) SetServerKeySecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.ServerKeySecret = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1WebhookContext) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *IoArgoprojEventsV1alpha1WebhookContext) SetUrl(v string) {
	o.Url = &v
}

func (o IoArgoprojEventsV1alpha1WebhookContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthSecret != nil {
		toSerialize["authSecret"] = o.AuthSecret
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.ServerCertPath != nil {
		toSerialize["serverCertPath"] = o.ServerCertPath
	}
	if o.ServerCertSecret != nil {
		toSerialize["serverCertSecret"] = o.ServerCertSecret
	}
	if o.ServerKeyPath != nil {
		toSerialize["serverKeyPath"] = o.ServerKeyPath
	}
	if o.ServerKeySecret != nil {
		toSerialize["serverKeySecret"] = o.ServerKeySecret
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1WebhookContext struct {
	value *IoArgoprojEventsV1alpha1WebhookContext
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1WebhookContext) Get() *IoArgoprojEventsV1alpha1WebhookContext {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1WebhookContext) Set(val *IoArgoprojEventsV1alpha1WebhookContext) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1WebhookContext) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1WebhookContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1WebhookContext(val *IoArgoprojEventsV1alpha1WebhookContext) *NullableIoArgoprojEventsV1alpha1WebhookContext {
	return &NullableIoArgoprojEventsV1alpha1WebhookContext{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1WebhookContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1WebhookContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


