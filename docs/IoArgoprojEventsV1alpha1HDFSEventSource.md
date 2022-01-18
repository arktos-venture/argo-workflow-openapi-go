# IoArgoprojEventsV1alpha1HDFSEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** |  | [optional] 
**CheckInterval** | Pointer to **string** |  | [optional] 
**HdfsUser** | Pointer to **string** | HDFSUser is the user to access HDFS file system. It is ignored if either ccache or keytab is used. | [optional] 
**KrbCCacheSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**KrbConfigConfigMap** | Pointer to [**IoK8sApiCoreV1ConfigMapKeySelector**](IoK8sApiCoreV1ConfigMapKeySelector.md) |  | [optional] 
**KrbKeytabSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**KrbRealm** | Pointer to **string** | KrbRealm is the Kerberos realm used with Kerberos keytab It must be set if keytab is used. | [optional] 
**KrbServicePrincipalName** | Pointer to **string** | KrbServicePrincipalName is the principal name of Kerberos service It must be set if either ccache or keytab is used. | [optional] 
**KrbUsername** | Pointer to **string** | KrbUsername is the Kerberos username used with Kerberos keytab It must be set if keytab is used. | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**WatchPathConfig** | Pointer to [**IoArgoprojEventsV1alpha1WatchPathConfig**](IoArgoprojEventsV1alpha1WatchPathConfig.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1HDFSEventSource

`func NewIoArgoprojEventsV1alpha1HDFSEventSource() *IoArgoprojEventsV1alpha1HDFSEventSource`

NewIoArgoprojEventsV1alpha1HDFSEventSource instantiates a new IoArgoprojEventsV1alpha1HDFSEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1HDFSEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1HDFSEventSourceWithDefaults() *IoArgoprojEventsV1alpha1HDFSEventSource`

NewIoArgoprojEventsV1alpha1HDFSEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1HDFSEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetCheckInterval

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetCheckInterval() string`

GetCheckInterval returns the CheckInterval field if non-nil, zero value otherwise.

### GetCheckIntervalOk

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetCheckIntervalOk() (*string, bool)`

GetCheckIntervalOk returns a tuple with the CheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckInterval

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetCheckInterval(v string)`

SetCheckInterval sets CheckInterval field to given value.

### HasCheckInterval

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasCheckInterval() bool`

HasCheckInterval returns a boolean if a field has been set.

### GetHdfsUser

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetHdfsUser() string`

GetHdfsUser returns the HdfsUser field if non-nil, zero value otherwise.

### GetHdfsUserOk

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetHdfsUserOk() (*string, bool)`

GetHdfsUserOk returns a tuple with the HdfsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfsUser

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetHdfsUser(v string)`

SetHdfsUser sets HdfsUser field to given value.

### HasHdfsUser

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasHdfsUser() bool`

HasHdfsUser returns a boolean if a field has been set.

### GetKrbCCacheSecret

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbCCacheSecret() IoK8sApiCoreV1SecretKeySelector`

GetKrbCCacheSecret returns the KrbCCacheSecret field if non-nil, zero value otherwise.

### GetKrbCCacheSecretOk

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbCCacheSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetKrbCCacheSecretOk returns a tuple with the KrbCCacheSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbCCacheSecret

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetKrbCCacheSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetKrbCCacheSecret sets KrbCCacheSecret field to given value.

### HasKrbCCacheSecret

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasKrbCCacheSecret() bool`

HasKrbCCacheSecret returns a boolean if a field has been set.

### GetKrbConfigConfigMap

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbConfigConfigMap() IoK8sApiCoreV1ConfigMapKeySelector`

GetKrbConfigConfigMap returns the KrbConfigConfigMap field if non-nil, zero value otherwise.

### GetKrbConfigConfigMapOk

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbConfigConfigMapOk() (*IoK8sApiCoreV1ConfigMapKeySelector, bool)`

GetKrbConfigConfigMapOk returns a tuple with the KrbConfigConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbConfigConfigMap

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetKrbConfigConfigMap(v IoK8sApiCoreV1ConfigMapKeySelector)`

SetKrbConfigConfigMap sets KrbConfigConfigMap field to given value.

### HasKrbConfigConfigMap

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasKrbConfigConfigMap() bool`

HasKrbConfigConfigMap returns a boolean if a field has been set.

### GetKrbKeytabSecret

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbKeytabSecret() IoK8sApiCoreV1SecretKeySelector`

GetKrbKeytabSecret returns the KrbKeytabSecret field if non-nil, zero value otherwise.

### GetKrbKeytabSecretOk

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbKeytabSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetKrbKeytabSecretOk returns a tuple with the KrbKeytabSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbKeytabSecret

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetKrbKeytabSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetKrbKeytabSecret sets KrbKeytabSecret field to given value.

### HasKrbKeytabSecret

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasKrbKeytabSecret() bool`

HasKrbKeytabSecret returns a boolean if a field has been set.

### GetKrbRealm

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbRealm() string`

GetKrbRealm returns the KrbRealm field if non-nil, zero value otherwise.

### GetKrbRealmOk

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbRealmOk() (*string, bool)`

GetKrbRealmOk returns a tuple with the KrbRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbRealm

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetKrbRealm(v string)`

SetKrbRealm sets KrbRealm field to given value.

### HasKrbRealm

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasKrbRealm() bool`

HasKrbRealm returns a boolean if a field has been set.

### GetKrbServicePrincipalName

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbServicePrincipalName() string`

GetKrbServicePrincipalName returns the KrbServicePrincipalName field if non-nil, zero value otherwise.

### GetKrbServicePrincipalNameOk

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbServicePrincipalNameOk() (*string, bool)`

GetKrbServicePrincipalNameOk returns a tuple with the KrbServicePrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbServicePrincipalName

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetKrbServicePrincipalName(v string)`

SetKrbServicePrincipalName sets KrbServicePrincipalName field to given value.

### HasKrbServicePrincipalName

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasKrbServicePrincipalName() bool`

HasKrbServicePrincipalName returns a boolean if a field has been set.

### GetKrbUsername

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbUsername() string`

GetKrbUsername returns the KrbUsername field if non-nil, zero value otherwise.

### GetKrbUsernameOk

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetKrbUsernameOk() (*string, bool)`

GetKrbUsernameOk returns a tuple with the KrbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbUsername

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetKrbUsername(v string)`

SetKrbUsername sets KrbUsername field to given value.

### HasKrbUsername

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasKrbUsername() bool`

HasKrbUsername returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetType

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWatchPathConfig

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetWatchPathConfig() IoArgoprojEventsV1alpha1WatchPathConfig`

GetWatchPathConfig returns the WatchPathConfig field if non-nil, zero value otherwise.

### GetWatchPathConfigOk

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) GetWatchPathConfigOk() (*IoArgoprojEventsV1alpha1WatchPathConfig, bool)`

GetWatchPathConfigOk returns a tuple with the WatchPathConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchPathConfig

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) SetWatchPathConfig(v IoArgoprojEventsV1alpha1WatchPathConfig)`

SetWatchPathConfig sets WatchPathConfig field to given value.

### HasWatchPathConfig

`func (o *IoArgoprojEventsV1alpha1HDFSEventSource) HasWatchPathConfig() bool`

HasWatchPathConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


