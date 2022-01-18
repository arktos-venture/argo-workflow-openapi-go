# IoArgoprojWorkflowV1alpha1S3EncryptionOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableEncryption** | Pointer to **bool** | EnableEncryption tells the driver to encrypt objects if set to true. If kmsKeyId and serverSideCustomerKeySecret are not set, SSE-S3 will be used | [optional] 
**KmsEncryptionContext** | Pointer to **string** | KmsEncryptionContext is a json blob that contains an encryption context. See https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context for more information | [optional] 
**KmsKeyId** | Pointer to **string** | KMSKeyId tells the driver to encrypt the object using the specified KMS Key. | [optional] 
**ServerSideCustomerKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1S3EncryptionOptions

`func NewIoArgoprojWorkflowV1alpha1S3EncryptionOptions() *IoArgoprojWorkflowV1alpha1S3EncryptionOptions`

NewIoArgoprojWorkflowV1alpha1S3EncryptionOptions instantiates a new IoArgoprojWorkflowV1alpha1S3EncryptionOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1S3EncryptionOptionsWithDefaults

`func NewIoArgoprojWorkflowV1alpha1S3EncryptionOptionsWithDefaults() *IoArgoprojWorkflowV1alpha1S3EncryptionOptions`

NewIoArgoprojWorkflowV1alpha1S3EncryptionOptionsWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1S3EncryptionOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableEncryption

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) GetEnableEncryption() bool`

GetEnableEncryption returns the EnableEncryption field if non-nil, zero value otherwise.

### GetEnableEncryptionOk

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) GetEnableEncryptionOk() (*bool, bool)`

GetEnableEncryptionOk returns a tuple with the EnableEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEncryption

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) SetEnableEncryption(v bool)`

SetEnableEncryption sets EnableEncryption field to given value.

### HasEnableEncryption

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) HasEnableEncryption() bool`

HasEnableEncryption returns a boolean if a field has been set.

### GetKmsEncryptionContext

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) GetKmsEncryptionContext() string`

GetKmsEncryptionContext returns the KmsEncryptionContext field if non-nil, zero value otherwise.

### GetKmsEncryptionContextOk

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) GetKmsEncryptionContextOk() (*string, bool)`

GetKmsEncryptionContextOk returns a tuple with the KmsEncryptionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsEncryptionContext

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) SetKmsEncryptionContext(v string)`

SetKmsEncryptionContext sets KmsEncryptionContext field to given value.

### HasKmsEncryptionContext

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) HasKmsEncryptionContext() bool`

HasKmsEncryptionContext returns a boolean if a field has been set.

### GetKmsKeyId

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.

### GetServerSideCustomerKeySecret

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) GetServerSideCustomerKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetServerSideCustomerKeySecret returns the ServerSideCustomerKeySecret field if non-nil, zero value otherwise.

### GetServerSideCustomerKeySecretOk

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) GetServerSideCustomerKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetServerSideCustomerKeySecretOk returns a tuple with the ServerSideCustomerKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSideCustomerKeySecret

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) SetServerSideCustomerKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetServerSideCustomerKeySecret sets ServerSideCustomerKeySecret field to given value.

### HasServerSideCustomerKeySecret

`func (o *IoArgoprojWorkflowV1alpha1S3EncryptionOptions) HasServerSideCustomerKeySecret() bool`

HasServerSideCustomerKeySecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


