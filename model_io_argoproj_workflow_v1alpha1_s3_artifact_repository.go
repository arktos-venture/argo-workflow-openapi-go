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

// IoArgoprojWorkflowV1alpha1S3ArtifactRepository S3ArtifactRepository defines the controller configuration for an S3 artifact repository
type IoArgoprojWorkflowV1alpha1S3ArtifactRepository struct {
	AccessKeySecret *IoK8sApiCoreV1SecretKeySelector `json:"accessKeySecret,omitempty"`
	// Bucket is the name of the bucket
	Bucket *string `json:"bucket,omitempty"`
	CreateBucketIfNotPresent *IoArgoprojWorkflowV1alpha1CreateS3BucketOptions `json:"createBucketIfNotPresent,omitempty"`
	EncryptionOptions *IoArgoprojWorkflowV1alpha1S3EncryptionOptions `json:"encryptionOptions,omitempty"`
	// Endpoint is the hostname of the bucket endpoint
	Endpoint *string `json:"endpoint,omitempty"`
	// Insecure will connect to the service with TLS
	Insecure *bool `json:"insecure,omitempty"`
	// KeyFormat is defines the format of how to store keys. Can reference workflow variables
	KeyFormat *string `json:"keyFormat,omitempty"`
	// KeyPrefix is prefix used as part of the bucket key in which the controller will store artifacts. DEPRECATED. Use KeyFormat instead
	KeyPrefix *string `json:"keyPrefix,omitempty"`
	// Region contains the optional bucket region
	Region *string `json:"region,omitempty"`
	// RoleARN is the Amazon Resource Name (ARN) of the role to assume.
	RoleARN *string `json:"roleARN,omitempty"`
	SecretKeySecret *IoK8sApiCoreV1SecretKeySelector `json:"secretKeySecret,omitempty"`
	// UseSDKCreds tells the driver to figure out credentials based on sdk defaults.
	UseSDKCreds *bool `json:"useSDKCreds,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1S3ArtifactRepository instantiates a new IoArgoprojWorkflowV1alpha1S3ArtifactRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1S3ArtifactRepository() *IoArgoprojWorkflowV1alpha1S3ArtifactRepository {
	this := IoArgoprojWorkflowV1alpha1S3ArtifactRepository{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1S3ArtifactRepositoryWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1S3ArtifactRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1S3ArtifactRepositoryWithDefaults() *IoArgoprojWorkflowV1alpha1S3ArtifactRepository {
	this := IoArgoprojWorkflowV1alpha1S3ArtifactRepository{}
	return &this
}

// GetAccessKeySecret returns the AccessKeySecret field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetAccessKeySecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.AccessKeySecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.AccessKeySecret
}

// GetAccessKeySecretOk returns a tuple with the AccessKeySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetAccessKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.AccessKeySecret == nil {
		return nil, false
	}
	return o.AccessKeySecret, true
}

// HasAccessKeySecret returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasAccessKeySecret() bool {
	if o != nil && o.AccessKeySecret != nil {
		return true
	}

	return false
}

// SetAccessKeySecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the AccessKeySecret field.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetAccessKeySecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.AccessKeySecret = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetBucket(v string) {
	o.Bucket = &v
}

// GetCreateBucketIfNotPresent returns the CreateBucketIfNotPresent field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetCreateBucketIfNotPresent() IoArgoprojWorkflowV1alpha1CreateS3BucketOptions {
	if o == nil || o.CreateBucketIfNotPresent == nil {
		var ret IoArgoprojWorkflowV1alpha1CreateS3BucketOptions
		return ret
	}
	return *o.CreateBucketIfNotPresent
}

// GetCreateBucketIfNotPresentOk returns a tuple with the CreateBucketIfNotPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetCreateBucketIfNotPresentOk() (*IoArgoprojWorkflowV1alpha1CreateS3BucketOptions, bool) {
	if o == nil || o.CreateBucketIfNotPresent == nil {
		return nil, false
	}
	return o.CreateBucketIfNotPresent, true
}

// HasCreateBucketIfNotPresent returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasCreateBucketIfNotPresent() bool {
	if o != nil && o.CreateBucketIfNotPresent != nil {
		return true
	}

	return false
}

// SetCreateBucketIfNotPresent gets a reference to the given IoArgoprojWorkflowV1alpha1CreateS3BucketOptions and assigns it to the CreateBucketIfNotPresent field.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetCreateBucketIfNotPresent(v IoArgoprojWorkflowV1alpha1CreateS3BucketOptions) {
	o.CreateBucketIfNotPresent = &v
}

// GetEncryptionOptions returns the EncryptionOptions field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetEncryptionOptions() IoArgoprojWorkflowV1alpha1S3EncryptionOptions {
	if o == nil || o.EncryptionOptions == nil {
		var ret IoArgoprojWorkflowV1alpha1S3EncryptionOptions
		return ret
	}
	return *o.EncryptionOptions
}

// GetEncryptionOptionsOk returns a tuple with the EncryptionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetEncryptionOptionsOk() (*IoArgoprojWorkflowV1alpha1S3EncryptionOptions, bool) {
	if o == nil || o.EncryptionOptions == nil {
		return nil, false
	}
	return o.EncryptionOptions, true
}

// HasEncryptionOptions returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasEncryptionOptions() bool {
	if o != nil && o.EncryptionOptions != nil {
		return true
	}

	return false
}

// SetEncryptionOptions gets a reference to the given IoArgoprojWorkflowV1alpha1S3EncryptionOptions and assigns it to the EncryptionOptions field.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetEncryptionOptions(v IoArgoprojWorkflowV1alpha1S3EncryptionOptions) {
	o.EncryptionOptions = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetInsecure returns the Insecure field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetInsecure() bool {
	if o == nil || o.Insecure == nil {
		var ret bool
		return ret
	}
	return *o.Insecure
}

// GetInsecureOk returns a tuple with the Insecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetInsecureOk() (*bool, bool) {
	if o == nil || o.Insecure == nil {
		return nil, false
	}
	return o.Insecure, true
}

// HasInsecure returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasInsecure() bool {
	if o != nil && o.Insecure != nil {
		return true
	}

	return false
}

// SetInsecure gets a reference to the given bool and assigns it to the Insecure field.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetInsecure(v bool) {
	o.Insecure = &v
}

// GetKeyFormat returns the KeyFormat field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetKeyFormat() string {
	if o == nil || o.KeyFormat == nil {
		var ret string
		return ret
	}
	return *o.KeyFormat
}

// GetKeyFormatOk returns a tuple with the KeyFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetKeyFormatOk() (*string, bool) {
	if o == nil || o.KeyFormat == nil {
		return nil, false
	}
	return o.KeyFormat, true
}

// HasKeyFormat returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasKeyFormat() bool {
	if o != nil && o.KeyFormat != nil {
		return true
	}

	return false
}

// SetKeyFormat gets a reference to the given string and assigns it to the KeyFormat field.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetKeyFormat(v string) {
	o.KeyFormat = &v
}

// GetKeyPrefix returns the KeyPrefix field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetKeyPrefix() string {
	if o == nil || o.KeyPrefix == nil {
		var ret string
		return ret
	}
	return *o.KeyPrefix
}

// GetKeyPrefixOk returns a tuple with the KeyPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetKeyPrefixOk() (*string, bool) {
	if o == nil || o.KeyPrefix == nil {
		return nil, false
	}
	return o.KeyPrefix, true
}

// HasKeyPrefix returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasKeyPrefix() bool {
	if o != nil && o.KeyPrefix != nil {
		return true
	}

	return false
}

// SetKeyPrefix gets a reference to the given string and assigns it to the KeyPrefix field.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetKeyPrefix(v string) {
	o.KeyPrefix = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetRegion(v string) {
	o.Region = &v
}

// GetRoleARN returns the RoleARN field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetRoleARN() string {
	if o == nil || o.RoleARN == nil {
		var ret string
		return ret
	}
	return *o.RoleARN
}

// GetRoleARNOk returns a tuple with the RoleARN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetRoleARNOk() (*string, bool) {
	if o == nil || o.RoleARN == nil {
		return nil, false
	}
	return o.RoleARN, true
}

// HasRoleARN returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasRoleARN() bool {
	if o != nil && o.RoleARN != nil {
		return true
	}

	return false
}

// SetRoleARN gets a reference to the given string and assigns it to the RoleARN field.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetRoleARN(v string) {
	o.RoleARN = &v
}

// GetSecretKeySecret returns the SecretKeySecret field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetSecretKeySecret() IoK8sApiCoreV1SecretKeySelector {
	if o == nil || o.SecretKeySecret == nil {
		var ret IoK8sApiCoreV1SecretKeySelector
		return ret
	}
	return *o.SecretKeySecret
}

// GetSecretKeySecretOk returns a tuple with the SecretKeySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetSecretKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool) {
	if o == nil || o.SecretKeySecret == nil {
		return nil, false
	}
	return o.SecretKeySecret, true
}

// HasSecretKeySecret returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasSecretKeySecret() bool {
	if o != nil && o.SecretKeySecret != nil {
		return true
	}

	return false
}

// SetSecretKeySecret gets a reference to the given IoK8sApiCoreV1SecretKeySelector and assigns it to the SecretKeySecret field.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetSecretKeySecret(v IoK8sApiCoreV1SecretKeySelector) {
	o.SecretKeySecret = &v
}

// GetUseSDKCreds returns the UseSDKCreds field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetUseSDKCreds() bool {
	if o == nil || o.UseSDKCreds == nil {
		var ret bool
		return ret
	}
	return *o.UseSDKCreds
}

// GetUseSDKCredsOk returns a tuple with the UseSDKCreds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) GetUseSDKCredsOk() (*bool, bool) {
	if o == nil || o.UseSDKCreds == nil {
		return nil, false
	}
	return o.UseSDKCreds, true
}

// HasUseSDKCreds returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) HasUseSDKCreds() bool {
	if o != nil && o.UseSDKCreds != nil {
		return true
	}

	return false
}

// SetUseSDKCreds gets a reference to the given bool and assigns it to the UseSDKCreds field.
func (o *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) SetUseSDKCreds(v bool) {
	o.UseSDKCreds = &v
}

func (o IoArgoprojWorkflowV1alpha1S3ArtifactRepository) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKeySecret != nil {
		toSerialize["accessKeySecret"] = o.AccessKeySecret
	}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.CreateBucketIfNotPresent != nil {
		toSerialize["createBucketIfNotPresent"] = o.CreateBucketIfNotPresent
	}
	if o.EncryptionOptions != nil {
		toSerialize["encryptionOptions"] = o.EncryptionOptions
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Insecure != nil {
		toSerialize["insecure"] = o.Insecure
	}
	if o.KeyFormat != nil {
		toSerialize["keyFormat"] = o.KeyFormat
	}
	if o.KeyPrefix != nil {
		toSerialize["keyPrefix"] = o.KeyPrefix
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.RoleARN != nil {
		toSerialize["roleARN"] = o.RoleARN
	}
	if o.SecretKeySecret != nil {
		toSerialize["secretKeySecret"] = o.SecretKeySecret
	}
	if o.UseSDKCreds != nil {
		toSerialize["useSDKCreds"] = o.UseSDKCreds
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1S3ArtifactRepository struct {
	value *IoArgoprojWorkflowV1alpha1S3ArtifactRepository
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1S3ArtifactRepository) Get() *IoArgoprojWorkflowV1alpha1S3ArtifactRepository {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1S3ArtifactRepository) Set(val *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1S3ArtifactRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1S3ArtifactRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1S3ArtifactRepository(val *IoArgoprojWorkflowV1alpha1S3ArtifactRepository) *NullableIoArgoprojWorkflowV1alpha1S3ArtifactRepository {
	return &NullableIoArgoprojWorkflowV1alpha1S3ArtifactRepository{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1S3ArtifactRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1S3ArtifactRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


