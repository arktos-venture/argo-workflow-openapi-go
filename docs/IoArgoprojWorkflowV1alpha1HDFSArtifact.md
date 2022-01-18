# IoArgoprojWorkflowV1alpha1HDFSArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** | Addresses is accessible addresses of HDFS name nodes | [optional] 
**Force** | Pointer to **bool** | Force copies a file forcibly even if it exists (default: false) | [optional] 
**HdfsUser** | Pointer to **string** | HDFSUser is the user to access HDFS file system. It is ignored if either ccache or keytab is used. | [optional] 
**KrbCCacheSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**KrbConfigConfigMap** | Pointer to [**IoK8sApiCoreV1ConfigMapKeySelector**](IoK8sApiCoreV1ConfigMapKeySelector.md) |  | [optional] 
**KrbKeytabSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**KrbRealm** | Pointer to **string** | KrbRealm is the Kerberos realm used with Kerberos keytab It must be set if keytab is used. | [optional] 
**KrbServicePrincipalName** | Pointer to **string** | KrbServicePrincipalName is the principal name of Kerberos service It must be set if either ccache or keytab is used. | [optional] 
**KrbUsername** | Pointer to **string** | KrbUsername is the Kerberos username used with Kerberos keytab It must be set if keytab is used. | [optional] 
**Path** | **string** | Path is a file path in HDFS | 

## Methods

### NewIoArgoprojWorkflowV1alpha1HDFSArtifact

`func NewIoArgoprojWorkflowV1alpha1HDFSArtifact(path string, ) *IoArgoprojWorkflowV1alpha1HDFSArtifact`

NewIoArgoprojWorkflowV1alpha1HDFSArtifact instantiates a new IoArgoprojWorkflowV1alpha1HDFSArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1HDFSArtifactWithDefaults

`func NewIoArgoprojWorkflowV1alpha1HDFSArtifactWithDefaults() *IoArgoprojWorkflowV1alpha1HDFSArtifact`

NewIoArgoprojWorkflowV1alpha1HDFSArtifactWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1HDFSArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetForce

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetHdfsUser

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetHdfsUser() string`

GetHdfsUser returns the HdfsUser field if non-nil, zero value otherwise.

### GetHdfsUserOk

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetHdfsUserOk() (*string, bool)`

GetHdfsUserOk returns a tuple with the HdfsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsUser

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetHdfsUser(v string)`

SetHdfsUser sets HdfsUser field to given value.

### HasHdfsUser

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasHdfsUser() bool`

HasHdfsUser returns a boolean if a field has been set.

### GetKrbCCacheSecret

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbCCacheSecret() IoK8sApiCoreV1SecretKeySelector`

GetKrbCCacheSecret returns the KrbCCacheSecret field if non-nil, zero value otherwise.

### GetKrbCCacheSecretOk

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbCCacheSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetKrbCCacheSecretOk returns a tuple with the KrbCCacheSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbCCacheSecret

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetKrbCCacheSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetKrbCCacheSecret sets KrbCCacheSecret field to given value.

### HasKrbCCacheSecret

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasKrbCCacheSecret() bool`

HasKrbCCacheSecret returns a boolean if a field has been set.

### GetKrbConfigConfigMap

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbConfigConfigMap() IoK8sApiCoreV1ConfigMapKeySelector`

GetKrbConfigConfigMap returns the KrbConfigConfigMap field if non-nil, zero value otherwise.

### GetKrbConfigConfigMapOk

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbConfigConfigMapOk() (*IoK8sApiCoreV1ConfigMapKeySelector, bool)`

GetKrbConfigConfigMapOk returns a tuple with the KrbConfigConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbConfigConfigMap

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetKrbConfigConfigMap(v IoK8sApiCoreV1ConfigMapKeySelector)`

SetKrbConfigConfigMap sets KrbConfigConfigMap field to given value.

### HasKrbConfigConfigMap

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasKrbConfigConfigMap() bool`

HasKrbConfigConfigMap returns a boolean if a field has been set.

### GetKrbKeytabSecret

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbKeytabSecret() IoK8sApiCoreV1SecretKeySelector`

GetKrbKeytabSecret returns the KrbKeytabSecret field if non-nil, zero value otherwise.

### GetKrbKeytabSecretOk

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbKeytabSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetKrbKeytabSecretOk returns a tuple with the KrbKeytabSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbKeytabSecret

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetKrbKeytabSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetKrbKeytabSecret sets KrbKeytabSecret field to given value.

### HasKrbKeytabSecret

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasKrbKeytabSecret() bool`

HasKrbKeytabSecret returns a boolean if a field has been set.

### GetKrbRealm

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbRealm() string`

GetKrbRealm returns the KrbRealm field if non-nil, zero value otherwise.

### GetKrbRealmOk

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbRealmOk() (*string, bool)`

GetKrbRealmOk returns a tuple with the KrbRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbRealm

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetKrbRealm(v string)`

SetKrbRealm sets KrbRealm field to given value.

### HasKrbRealm

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasKrbRealm() bool`

HasKrbRealm returns a boolean if a field has been set.

### GetKrbServicePrincipalName

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbServicePrincipalName() string`

GetKrbServicePrincipalName returns the KrbServicePrincipalName field if non-nil, zero value otherwise.

### GetKrbServicePrincipalNameOk

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbServicePrincipalNameOk() (*string, bool)`

GetKrbServicePrincipalNameOk returns a tuple with the KrbServicePrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbServicePrincipalName

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetKrbServicePrincipalName(v string)`

SetKrbServicePrincipalName sets KrbServicePrincipalName field to given value.

### HasKrbServicePrincipalName

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasKrbServicePrincipalName() bool`

HasKrbServicePrincipalName returns a boolean if a field has been set.

### GetKrbUsername

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbUsername() string`

GetKrbUsername returns the KrbUsername field if non-nil, zero value otherwise.

### GetKrbUsernameOk

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetKrbUsernameOk() (*string, bool)`

GetKrbUsernameOk returns a tuple with the KrbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbUsername

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetKrbUsername(v string)`

SetKrbUsername sets KrbUsername field to given value.

### HasKrbUsername

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) HasKrbUsername() bool`

HasKrbUsername returns a boolean if a field has been set.

### GetPath

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoArgoprojWorkflowV1alpha1HDFSArtifact) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


