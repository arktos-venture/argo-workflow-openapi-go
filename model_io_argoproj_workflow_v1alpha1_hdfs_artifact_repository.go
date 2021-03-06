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

// IoArgoprojWorkflowV1alpha1HDFSArtifactRepository HDFSArtifactRepository defines the controller configuration for an HDFS artifact repository
type IoArgoprojWorkflowV1alpha1HDFSArtifactRepository struct {
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
	// PathFormat is defines the format of path to store a file. Can reference workflow variables
	PathFormat *string `json:"pathFormat,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1HDFSArtifactRepository instantiates a new IoArgoprojWorkflowV1alpha1HDFSArtifactRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1HDFSArtifactRepository() *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository {
	this := IoArgoprojWorkflowV1alpha1HDFSArtifactRepository{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1HDFSArtifactRepositoryWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1HDFSArtifactRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1HDFSArtifactRepositoryWithDefaults() *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository {
	this := IoArgoprojWorkflowV1alpha1HDFSArtifactRepository{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetAddresses() []string {
	if o == nil || o.Addresses == nil {
		var ret []string
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetAddressesOk() (*[]string, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) SetAddresses(v []string) {
	o.Addresses = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) SetForce(v bool) {
	o.Force = &v
}

// GetHdfsUser returns the HdfsUser field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetHdfsUser() string {
	if o == nil || o.HdfsUser == nil {
		var ret string
		return ret
	}
	return *o.HdfsUser
}

// GetHdfsUserOk returns a tuple with the HdfsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetHdfsUserOk() (*string, bool) {
	if o == nil || o.HdfsUser == nil {
		return nil, false
	}
	return o.HdfsUser, true
}

// HasHdfsUser returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) HasHdfsUser() bool {
	if o != nil && o.HdfsUser != nil {
		return true
	}

	return false
}

// SetHdfsUser gets a reference to the given string and assigns it to the HdfsUser field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) SetHdfsUser(v string) {
	o.HdfsUser = &v
}

// GetKrbCCacheSecret returns the KrbCCacheSecret field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetKrbCCacheSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.KrbCCacheSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.KrbCCacheSecret
}

// GetKrbCCacheSecretOk returns a tuple with the KrbCCacheSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetKrbCCacheSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.KrbCCacheSecret == nil {
		return nil, false
	}
	return o.KrbCCacheSecret, true
}

// HasKrbCCacheSecret returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) HasKrbCCacheSecret() bool {
	if o != nil && o.KrbCCacheSecret != nil {
		return true
	}

	return false
}

// SetKrbCCacheSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the KrbCCacheSecret field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) SetKrbCCacheSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.KrbCCacheSecret = &v
}

// GetKrbConfigConfigMap returns the KrbConfigConfigMap field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetKrbConfigConfigMap() IoK8sApiCoreV1ConfigMapKeySelector {
	if o == nil || o.KrbConfigConfigMap == nil {
		var ret IoK8sApiCoreV1ConfigMapKeySelector
		return ret
	}
	return *o.KrbConfigConfigMap
}

// GetKrbConfigConfigMapOk returns a tuple with the KrbConfigConfigMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetKrbConfigConfigMapOk() (*IoK8sApiCoreV1ConfigMapKeySelector, bool) {
	if o == nil || o.KrbConfigConfigMap == nil {
		return nil, false
	}
	return o.KrbConfigConfigMap, true
}

// HasKrbConfigConfigMap returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) HasKrbConfigConfigMap() bool {
	if o != nil && o.KrbConfigConfigMap != nil {
		return true
	}

	return false
}

// SetKrbConfigConfigMap gets a reference to the given IoK8sApiCoreV1ConfigMapKeySelector and assigns it to the KrbConfigConfigMap field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) SetKrbConfigConfigMap(v IoK8sApiCoreV1ConfigMapKeySelector) {
	o.KrbConfigConfigMap = &v
}

// GetKrbKeytabSecret returns the KrbKeytabSecret field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetKrbKeytabSecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.KrbKeytabSecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.KrbKeytabSecret
}

// GetKrbKeytabSecretOk returns a tuple with the KrbKeytabSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetKrbKeytabSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.KrbKeytabSecret == nil {
		return nil, false
	}
	return o.KrbKeytabSecret, true
}

// HasKrbKeytabSecret returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) HasKrbKeytabSecret() bool {
	if o != nil && o.KrbKeytabSecret != nil {
		return true
	}

	return false
}

// SetKrbKeytabSecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the KrbKeytabSecret field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) SetKrbKeytabSecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.KrbKeytabSecret = &v
}

// GetKrbRealm returns the KrbRealm field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetKrbRealm() string {
	if o == nil || o.KrbRealm == nil {
		var ret string
		return ret
	}
	return *o.KrbRealm
}

// GetKrbRealmOk returns a tuple with the KrbRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetKrbRealmOk() (*string, bool) {
	if o == nil || o.KrbRealm == nil {
		return nil, false
	}
	return o.KrbRealm, true
}

// HasKrbRealm returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) HasKrbRealm() bool {
	if o != nil && o.KrbRealm != nil {
		return true
	}

	return false
}

// SetKrbRealm gets a reference to the given string and assigns it to the KrbRealm field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) SetKrbRealm(v string) {
	o.KrbRealm = &v
}

// GetKrbServicePrincipalName returns the KrbServicePrincipalName field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetKrbServicePrincipalName() string {
	if o == nil || o.KrbServicePrincipalName == nil {
		var ret string
		return ret
	}
	return *o.KrbServicePrincipalName
}

// GetKrbServicePrincipalNameOk returns a tuple with the KrbServicePrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetKrbServicePrincipalNameOk() (*string, bool) {
	if o == nil || o.KrbServicePrincipalName == nil {
		return nil, false
	}
	return o.KrbServicePrincipalName, true
}

// HasKrbServicePrincipalName returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) HasKrbServicePrincipalName() bool {
	if o != nil && o.KrbServicePrincipalName != nil {
		return true
	}

	return false
}

// SetKrbServicePrincipalName gets a reference to the given string and assigns it to the KrbServicePrincipalName field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) SetKrbServicePrincipalName(v string) {
	o.KrbServicePrincipalName = &v
}

// GetKrbUsername returns the KrbUsername field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetKrbUsername() string {
	if o == nil || o.KrbUsername == nil {
		var ret string
		return ret
	}
	return *o.KrbUsername
}

// GetKrbUsernameOk returns a tuple with the KrbUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetKrbUsernameOk() (*string, bool) {
	if o == nil || o.KrbUsername == nil {
		return nil, false
	}
	return o.KrbUsername, true
}

// HasKrbUsername returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) HasKrbUsername() bool {
	if o != nil && o.KrbUsername != nil {
		return true
	}

	return false
}

// SetKrbUsername gets a reference to the given string and assigns it to the KrbUsername field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) SetKrbUsername(v string) {
	o.KrbUsername = &v
}

// GetPathFormat returns the PathFormat field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetPathFormat() string {
	if o == nil || o.PathFormat == nil {
		var ret string
		return ret
	}
	return *o.PathFormat
}

// GetPathFormatOk returns a tuple with the PathFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) GetPathFormatOk() (*string, bool) {
	if o == nil || o.PathFormat == nil {
		return nil, false
	}
	return o.PathFormat, true
}

// HasPathFormat returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) HasPathFormat() bool {
	if o != nil && o.PathFormat != nil {
		return true
	}

	return false
}

// SetPathFormat gets a reference to the given string and assigns it to the PathFormat field.
func (o *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) SetPathFormat(v string) {
	o.PathFormat = &v
}

func (o IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) MarshalJSON() ([]byte, error) {
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
	if o.PathFormat != nil {
		toSerialize["pathFormat"] = o.PathFormat
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1HDFSArtifactRepository struct {
	value *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1HDFSArtifactRepository) Get() *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1HDFSArtifactRepository) Set(val *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1HDFSArtifactRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1HDFSArtifactRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1HDFSArtifactRepository(val *IoArgoprojWorkflowV1alpha1HDFSArtifactRepository) *NullableIoArgoprojWorkflowV1alpha1HDFSArtifactRepository {
	return &NullableIoArgoprojWorkflowV1alpha1HDFSArtifactRepository{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1HDFSArtifactRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1HDFSArtifactRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


