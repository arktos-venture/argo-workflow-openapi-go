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

// IoArgoprojWorkflowV1alpha1HDFSArtifact HDFSArtifact is the location of an HDFS artifact
type IoArgoprojWorkflowV1alpha1HDFSArtifact struct {
	// Addresses is accessible addresses of HDFS name nodes
	Addresses *[]string `json:"addresses,omitempty"`
	// Force copies a file forcibly even if it exists (default: false)
	Force *bool `json:"force,omitempty"`
	// HDFSUser is the user to access HDFS file system. It is ignored if either ccache or keytab is used.
	HdfsUser *string `json:"hdfsUser,omitempty"`
	KrbCCacheSecret *IoK8sApiCoreV1SecretKeySelector `json:"krbCCacheSecret,omitempty"`
	KrbConfigConfigMap *IoK8sApiCoreV1ConfigMapKeySelector `json:"krbConfigConfigMap,omitempty"`
	KrbKeytabSecret *IoK8sApiCoreV1SecretKeySelector `json:"krbKeytabSecret,omitempty"`
	// KrbRealm is the Kerberos realm used with Kerberos keytab It must be set if keytab is used.
	KrbRealm *string `json:"krbRealm,omitempty"`
	// KrbServicePrincipalName is the principal name of Kerberos service It must be set if either ccache or keytab is used.
	KrbServicePrincipalName *string `json:"krbServicePrincipalName,omitempty"`
	// KrbUsername is the Kerberos username used with Kerberos keytab It must be set if keytab is used.
	KrbUsername *string `json:"krbUsername,omitempty"`
	// Path is a file path in HDFS
	Path string `json:"path"`
}

// NewIoArgoprojWorkflowV1alpha1HDFSArtifact instantiates a new IoArgoprojWorkflowV1alpha1HDFSArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1HDFSArtifact(path string) *IoArgoprojWorkflowV1alpha1HDFSArtifact {
	this := IoArgoprojWorkflowV1alpha1HDFSArtifact{}
	this.Path = path
	return &this
}

// NewIoArgoprojWorkflowV1alpha1HDFSArtifactWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1HDFSArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1HDFSArtifactWithDefaults() *IoArgoprojWorkflowV1alpha1HDFSArtifact {
	this := IoArgoprojWorkflowV1alpha1HDFSArtifact{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetAddresses() []string {
	if o == nil || o.Addresses == nil {
		var ret []string
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetAddressesOk() (*[]string, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetAddresses(v []string) {
	o.Addresses = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetForce(v bool) {
	o.Force = &v
}

// GetHdfsUser returns the HdfsUser field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetHdfsUser() string {
	if o == nil || o.HdfsUser == nil {
		var ret string
		return ret
	}
	return *o.HdfsUser
}

// GetHdfsUserOk returns a tuple with the HdfsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetHdfsUserOk() (*string, bool) {
	if o == nil || o.HdfsUser == nil {
		return nil, false
	}
	return o.HdfsUser, true
}

// HasHdfsUser returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasHdfsUser() bool {
	if o != nil && o.HdfsUser != nil {
		return true
	}

	return false
}

// SetHdfsUser gets a reference to the given string and assigns it to the HdfsUser field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetHdfsUser(v string) {
	o.HdfsUser = &v
}

// GetKrbCCacheSecret returns the KrbCCacheSecret field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbCCacheSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.KrbCCacheSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.KrbCCacheSecret
}

// GetKrbCCacheSecretOk returns a tuple with the KrbCCacheSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbCCacheSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.KrbCCacheSecret == nil {
		return nil, false
	}
	return o.KrbCCacheSecret, true
}

// HasKrbCCacheSecret returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasKrbCCacheSecret() bool {
	if o != nil && o.KrbCCacheSecret != nil {
		return true
	}

	return false
}

// SetKrbCCacheSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the KrbCCacheSecret field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetKrbCCacheSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.KrbCCacheSecret = &v
}

// GetKrbConfigConfigMap returns the KrbConfigConfigMap field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbConfigConfigMap() IoK8sApiCoreV1ConfigMapKeySelector {
	if o == nil || o.KrbConfigConfigMap == nil {
		var ret IoK8sApiCoreV1ConfigMapKeySelector
		return ret
	}
	return *o.KrbConfigConfigMap
}

// GetKrbConfigConfigMapOk returns a tuple with the KrbConfigConfigMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbConfigConfigMapOk() (*IoK8sApiCoreV1ConfigMapKeySelector, bool) {
	if o == nil || o.KrbConfigConfigMap == nil {
		return nil, false
	}
	return o.KrbConfigConfigMap, true
}

// HasKrbConfigConfigMap returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasKrbConfigConfigMap() bool {
	if o != nil && o.KrbConfigConfigMap != nil {
		return true
	}

	return false
}

// SetKrbConfigConfigMap gets a reference to the given IoK8sApiCoreV1ConfigMapKeySelector and assigns it to the KrbConfigConfigMap field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetKrbConfigConfigMap(v IoK8sApiCoreV1ConfigMapKeySelector) {
	o.KrbConfigConfigMap = &v
}

// GetKrbKeytabSecret returns the KrbKeytabSecret field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbKeytabSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.KrbKeytabSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.KrbKeytabSecret
}

// GetKrbKeytabSecretOk returns a tuple with the KrbKeytabSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbKeytabSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.KrbKeytabSecret == nil {
		return nil, false
	}
	return o.KrbKeytabSecret, true
}

// HasKrbKeytabSecret returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasKrbKeytabSecret() bool {
	if o != nil && o.KrbKeytabSecret != nil {
		return true
	}

	return false
}

// SetKrbKeytabSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the KrbKeytabSecret field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetKrbKeytabSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.KrbKeytabSecret = &v
}

// GetKrbRealm returns the KrbRealm field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbRealm() string {
	if o == nil || o.KrbRealm == nil {
		var ret string
		return ret
	}
	return *o.KrbRealm
}

// GetKrbRealmOk returns a tuple with the KrbRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbRealmOk() (*string, bool) {
	if o == nil || o.KrbRealm == nil {
		return nil, false
	}
	return o.KrbRealm, true
}

// HasKrbRealm returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasKrbRealm() bool {
	if o != nil && o.KrbRealm != nil {
		return true
	}

	return false
}

// SetKrbRealm gets a reference to the given string and assigns it to the KrbRealm field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetKrbRealm(v string) {
	o.KrbRealm = &v
}

// GetKrbServicePrincipalName returns the KrbServicePrincipalName field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbServicePrincipalName() string {
	if o == nil || o.KrbServicePrincipalName == nil {
		var ret string
		return ret
	}
	return *o.KrbServicePrincipalName
}

// GetKrbServicePrincipalNameOk returns a tuple with the KrbServicePrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbServicePrincipalNameOk() (*string, bool) {
	if o == nil || o.KrbServicePrincipalName == nil {
		return nil, false
	}
	return o.KrbServicePrincipalName, true
}

// HasKrbServicePrincipalName returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasKrbServicePrincipalName() bool {
	if o != nil && o.KrbServicePrincipalName != nil {
		return true
	}

	return false
}

// SetKrbServicePrincipalName gets a reference to the given string and assigns it to the KrbServicePrincipalName field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetKrbServicePrincipalName(v string) {
	o.KrbServicePrincipalName = &v
}

// GetKrbUsername returns the KrbUsername field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbUsername() string {
	if o == nil || o.KrbUsername == nil {
		var ret string
		return ret
	}
	return *o.KrbUsername
}

// GetKrbUsernameOk returns a tuple with the KrbUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbUsernameOk() (*string, bool) {
	if o == nil || o.KrbUsername == nil {
		return nil, false
	}
	return o.KrbUsername, true
}

// HasKrbUsername returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasKrbUsername() bool {
	if o != nil && o.KrbUsername != nil {
		return true
	}

	return false
}

// SetKrbUsername gets a reference to the given string and assigns it to the KrbUsername field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetKrbUsername(v string) {
	o.KrbUsername = &v
}

// GetPath returns the Path field value
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetPath(v string) {
	o.Path = v
}

func (o IoArgoprojWorkflowV1alpha1HDFSArtifact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	if o.HdfsUser != nil {
		toSerialize["hdfsUser"] = o.HdfsUser
	}
	if o.KrbCCacheSecret != nil {
		toSerialize["krbCCacheSecret"] = o.KrbCCacheSecret
	}
	if o.KrbConfigConfigMap != nil {
		toSerialize["krbConfigConfigMap"] = o.KrbConfigConfigMap
	}
	if o.KrbKeytabSecret != nil {
		toSerialize["krbKeytabSecret"] = o.KrbKeytabSecret
	}
	if o.KrbRealm != nil {
		toSerialize["krbRealm"] = o.KrbRealm
	}
	if o.KrbServicePrincipalName != nil {
		toSerialize["krbServicePrincipalName"] = o.KrbServicePrincipalName
	}
	if o.KrbUsername != nil {
		toSerialize["krbUsername"] = o.KrbUsername
	}
	if true {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1HDFSArtifact struct {
	value *IoArgoprojWorkflowV1alpha1HDFSArtifact
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1HDFSArtifact) Get() *IoArgoprojWorkflowV1alpha1HDFSArtifact {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1HDFSArtifact) Set(val *IoArgoprojWorkflowV1alpha1HDFSArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1HDFSArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1HDFSArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1HDFSArtifact(val *IoArgoprojWorkflowV1alpha1HDFSArtifact) *NullableIoArgoprojWorkflowV1alpha1HDFSArtifact {
	return &NullableIoArgoprojWorkflowV1alpha1HDFSArtifact{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1HDFSArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1HDFSArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


