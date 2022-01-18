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

// IoK8sApiCoreV1ServiceAccountTokenProjection ServiceAccountTokenProjection represents a projected service account token volume. This projection can be used to insert a service account token into the pods runtime filesystem for use against APIs (Kubernetes API Server or otherwise).
type IoK8sApiCoreV1ServiceAccountTokenProjection struct {
	// Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.
	Audience *string `json:"audience,omitempty"`
	// ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.
	ExpirationSeconds *int32 `json:"expirationSeconds,omitempty"`
	// Path is the path relative to the mount point of the file to project the token into.
	Path string `json:"path"`
}

// NewIoK8sApiCoreV1ServiceAccountTokenProjection instantiates a new IoK8sApiCoreV1ServiceAccountTokenProjection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1ServiceAccountTokenProjection(path string) *IoK8sApiCoreV1ServiceAccountTokenProjection {
	this := IoK8sApiCoreV1ServiceAccountTokenProjection{}
	this.Path = path
	return &this
}

// NewIoK8sApiCoreV1ServiceAccountTokenProjectionWithDefaults instantiates a new IoK8sApiCoreV1ServiceAccountTokenProjection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1ServiceAccountTokenProjectionWithDefaults() *IoK8sApiCoreV1ServiceAccountTokenProjection {
	this := IoK8sApiCoreV1ServiceAccountTokenProjection{}
	return &this
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1ServiceAccountTokenProjection) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1ServiceAccountTokenProjection) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1ServiceAccountTokenProjection) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *IoK8sApiCoreV1ServiceAccountTokenProjection) SetAudience(v string) {
	o.Audience = &v
}

// GetExpirationSeconds returns the ExpirationSeconds field value if set, zero value otherwise.
func (o *IoK8sApiCoreV1ServiceAccountTokenProjection) GetExpirationSeconds() int32 {
	if o == nil || o.ExpirationSeconds == nil {
		var ret int32
		return ret
	}
	return *o.ExpirationSeconds
}

// GetExpirationSecondsOk returns a tuple with the ExpirationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1ServiceAccountTokenProjection) GetExpirationSecondsOk() (*int32, bool) {
	if o == nil || o.ExpirationSeconds == nil {
		return nil, false
	}
	return o.ExpirationSeconds, true
}

// HasExpirationSeconds returns a boolean if a field has been set.
func (o *IoK8sApiCoreV1ServiceAccountTokenProjection) HasExpirationSeconds() bool {
	if o != nil && o.ExpirationSeconds != nil {
		return true
	}

	return false
}

// SetExpirationSeconds gets a reference to the given int32 and assigns it to the ExpirationSeconds field.
func (o *IoK8sApiCoreV1ServiceAccountTokenProjection) SetExpirationSeconds(v int32) {
	o.ExpirationSeconds = &v
}

// GetPath returns the Path field value
func (o *IoK8sApiCoreV1ServiceAccountTokenProjection) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1ServiceAccountTokenProjection) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *IoK8sApiCoreV1ServiceAccountTokenProjection) SetPath(v string) {
	o.Path = v
}

func (o IoK8sApiCoreV1ServiceAccountTokenProjection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.ExpirationSeconds != nil {
		toSerialize["expirationSeconds"] = o.ExpirationSeconds
	}
	if true {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1ServiceAccountTokenProjection struct {
	value *IoK8sApiCoreV1ServiceAccountTokenProjection
	isSet bool
}

func (v NullableIoK8sApiCoreV1ServiceAccountTokenProjection) Get() *IoK8sApiCoreV1ServiceAccountTokenProjection {
	return v.value
}

func (v *NullableIoK8sApiCoreV1ServiceAccountTokenProjection) Set(val *IoK8sApiCoreV1ServiceAccountTokenProjection) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1ServiceAccountTokenProjection) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1ServiceAccountTokenProjection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1ServiceAccountTokenProjection(val *IoK8sApiCoreV1ServiceAccountTokenProjection) *NullableIoK8sApiCoreV1ServiceAccountTokenProjection {
	return &NullableIoK8sApiCoreV1ServiceAccountTokenProjection{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1ServiceAccountTokenProjection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1ServiceAccountTokenProjection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


