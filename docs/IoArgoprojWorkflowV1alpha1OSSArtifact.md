# IoArgoprojWorkflowV1alpha1OSSArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Bucket** | Pointer to **string** | Bucket is the name of the bucket | [optional] 
**CreateBucketIfNotPresent** | Pointer to **bool** | CreateBucketIfNotPresent tells the driver to attempt to create the OSS bucket for output artifacts, if it doesn&#39;t exist | [optional] 
**Endpoint** | Pointer to **string** | Endpoint is the hostname of the bucket endpoint | [optional] 
**Key** | **string** | Key is the path in the bucket where the artifact resides | 
**LifecycleRule** | Pointer to [**IoArgoprojWorkflowV1alpha1OSSLifecycleRule**](IoArgoprojWorkflowV1alpha1OSSLifecycleRule.md) |  | [optional] 
**SecretKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**SecurityToken** | Pointer to **string** | SecurityToken is the user&#39;s temporary security token. For more details, check out: https://www.alibabacloud.com/help/doc-detail/100624.htm | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1OSSArtifact

`func NewIoArgoprojWorkflowV1alpha1OSSArtifact(key string, ) *IoArgoprojWorkflowV1alpha1OSSArtifact`

NewIoArgoprojWorkflowV1alpha1OSSArtifact instantiates a new IoArgoprojWorkflowV1alpha1OSSArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1OSSArtifactWithDefaults

`func NewIoArgoprojWorkflowV1alpha1OSSArtifactWithDefaults() *IoArgoprojWorkflowV1alpha1OSSArtifact`

NewIoArgoprojWorkflowV1alpha1OSSArtifactWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1OSSArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeySecret

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetAccessKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetAccessKeySecret returns the AccessKeySecret field if non-nil, zero value otherwise.

### GetAccessKeySecretOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetAccessKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetAccessKeySecretOk returns a tuple with the AccessKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeySecret

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) SetAccessKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetAccessKeySecret sets AccessKeySecret field to given value.

### HasAccessKeySecret

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) HasAccessKeySecret() bool`

HasAccessKeySecret returns a boolean if a field has been set.

### GetBucket

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCreateBucketIfNotPresent

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetCreateBucketIfNotPresent() bool`

GetCreateBucketIfNotPresent returns the CreateBucketIfNotPresent field if non-nil, zero value otherwise.

### GetCreateBucketIfNotPresentOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetCreateBucketIfNotPresentOk() (*bool, bool)`

GetCreateBucketIfNotPresentOk returns a tuple with the CreateBucketIfNotPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBucketIfNotPresent

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) SetCreateBucketIfNotPresent(v bool)`

SetCreateBucketIfNotPresent sets CreateBucketIfNotPresent field to given value.

### HasCreateBucketIfNotPresent

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) HasCreateBucketIfNotPresent() bool`

HasCreateBucketIfNotPresent returns a boolean if a field has been set.

### GetEndpoint

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetKey

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) SetKey(v string)`

SetKey sets Key field to given value.


### GetLifecycleRule

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetLifecycleRule() IoArgoprojWorkflowV1alpha1OSSLifecycleRule`

GetLifecycleRule returns the LifecycleRule field if non-nil, zero value otherwise.

### GetLifecycleRuleOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetLifecycleRuleOk() (*IoArgoprojWorkflowV1alpha1OSSLifecycleRule, bool)`

GetLifecycleRuleOk returns a tuple with the LifecycleRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleRule

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) SetLifecycleRule(v IoArgoprojWorkflowV1alpha1OSSLifecycleRule)`

SetLifecycleRule sets LifecycleRule field to given value.

### HasLifecycleRule

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) HasLifecycleRule() bool`

HasLifecycleRule returns a boolean if a field has been set.

### GetSecretKeySecret

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetSecretKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetSecretKeySecret returns the SecretKeySecret field if non-nil, zero value otherwise.

### GetSecretKeySecretOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetSecretKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSecretKeySecretOk returns a tuple with the SecretKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeySecret

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) SetSecretKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetSecretKeySecret sets SecretKeySecret field to given value.

### HasSecretKeySecret

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) HasSecretKeySecret() bool`

HasSecretKeySecret returns a boolean if a field has been set.

### GetSecurityToken

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *IoArgoprojWorkflowV1alpha1OSSArtifact) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


