# IoArgoprojWorkflowV1alpha1OSSArtifactRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Bucket** | Pointer to **string** | Bucket is the name of the bucket | [optional] 
**CreateBucketIfNotPresent** | Pointer to **bool** | CreateBucketIfNotPresent tells the driver to attempt to create the OSS bucket for output artifacts, if it doesn&#39;t exist | [optional] 
**Endpoint** | Pointer to **string** | Endpoint is the hostname of the bucket endpoint | [optional] 
**KeyFormat** | Pointer to **string** | KeyFormat is defines the format of how to store keys. Can reference workflow variables | [optional] 
**LifecycleRule** | Pointer to [**IoArgoprojWorkflowV1alpha1OSSLifecycleRule**](IoArgoprojWorkflowV1alpha1OSSLifecycleRule.md) |  | [optional] 
**SecretKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**SecurityToken** | Pointer to **string** | SecurityToken is the user&#39;s temporary security token. For more details, check out: https://www.alibabacloud.com/help/doc-detail/100624.htm | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1OSSArtifactRepository

`func NewIoArgoprojWorkflowV1alpha1OSSArtifactRepository() *IoArgoprojWorkflowV1alpha1OSSArtifactRepository`

NewIoArgoprojWorkflowV1alpha1OSSArtifactRepository instantiates a new IoArgoprojWorkflowV1alpha1OSSArtifactRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1OSSArtifactRepositoryWithDefaults

`func NewIoArgoprojWorkflowV1alpha1OSSArtifactRepositoryWithDefaults() *IoArgoprojWorkflowV1alpha1OSSArtifactRepository`

NewIoArgoprojWorkflowV1alpha1OSSArtifactRepositoryWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1OSSArtifactRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeySecret

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetAccessKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetAccessKeySecret returns the AccessKeySecret field if non-nil, zero value otherwise.

### GetAccessKeySecretOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetAccessKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetAccessKeySecretOk returns a tuple with the AccessKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeySecret

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) SetAccessKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetAccessKeySecret sets AccessKeySecret field to given value.

### HasAccessKeySecret

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) HasAccessKeySecret() bool`

HasAccessKeySecret returns a boolean if a field has been set.

### GetBucket

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCreateBucketIfNotPresent

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetCreateBucketIfNotPresent() bool`

GetCreateBucketIfNotPresent returns the CreateBucketIfNotPresent field if non-nil, zero value otherwise.

### GetCreateBucketIfNotPresentOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetCreateBucketIfNotPresentOk() (*bool, bool)`

GetCreateBucketIfNotPresentOk returns a tuple with the CreateBucketIfNotPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBucketIfNotPresent

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) SetCreateBucketIfNotPresent(v bool)`

SetCreateBucketIfNotPresent sets CreateBucketIfNotPresent field to given value.

### HasCreateBucketIfNotPresent

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) HasCreateBucketIfNotPresent() bool`

HasCreateBucketIfNotPresent returns a boolean if a field has been set.

### GetEndpoint

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetKeyFormat

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetKeyFormat() string`

GetKeyFormat returns the KeyFormat field if non-nil, zero value otherwise.

### GetKeyFormatOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetKeyFormatOk() (*string, bool)`

GetKeyFormatOk returns a tuple with the KeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFormat

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) SetKeyFormat(v string)`

SetKeyFormat sets KeyFormat field to given value.

### HasKeyFormat

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) HasKeyFormat() bool`

HasKeyFormat returns a boolean if a field has been set.

### GetLifecycleRule

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetLifecycleRule() IoArgoprojWorkflowV1alpha1OSSLifecycleRule`

GetLifecycleRule returns the LifecycleRule field if non-nil, zero value otherwise.

### GetLifecycleRuleOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetLifecycleRuleOk() (*IoArgoprojWorkflowV1alpha1OSSLifecycleRule, bool)`

GetLifecycleRuleOk returns a tuple with the LifecycleRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleRule

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) SetLifecycleRule(v IoArgoprojWorkflowV1alpha1OSSLifecycleRule)`

SetLifecycleRule sets LifecycleRule field to given value.

### HasLifecycleRule

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) HasLifecycleRule() bool`

HasLifecycleRule returns a boolean if a field has been set.

### GetSecretKeySecret

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetSecretKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetSecretKeySecret returns the SecretKeySecret field if non-nil, zero value otherwise.

### GetSecretKeySecretOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetSecretKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSecretKeySecretOk returns a tuple with the SecretKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeySecret

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) SetSecretKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetSecretKeySecret sets SecretKeySecret field to given value.

### HasSecretKeySecret

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) HasSecretKeySecret() bool`

HasSecretKeySecret returns a boolean if a field has been set.

### GetSecurityToken

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifactRepository) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


