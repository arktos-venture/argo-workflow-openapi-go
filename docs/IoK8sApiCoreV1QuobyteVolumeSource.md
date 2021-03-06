# IoK8sApiCoreV1QuobyteVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Group to map volume access to Default is no group | [optional] 
**ReadOnly** | Pointer to **bool** | ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false. | [optional] 
**Registry** | **string** | Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes | 
**Tenant** | Pointer to **string** | Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin | [optional] 
**User** | Pointer to **string** | User to map volume access to Defaults to serivceaccount user | [optional] 
**Volume** | **string** | Volume is a string that references an already created Quobyte volume by name. | 

## Methods

### NewIoK8sApiCoreV1QuobyteVolumeSource

`func NewIoK8sApiCoreV1QuobyteVolumeSource(registry string, volume string, ) *IoK8sApiCoreV1QuobyteVolumeSource`

NewIoK8sApiCoreV1QuobyteVolumeSource instantiates a new IoK8sApiCoreV1QuobyteVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1QuobyteVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1QuobyteVolumeSourceWithDefaults() *IoK8sApiCoreV1QuobyteVolumeSource`

NewIoK8sApiCoreV1QuobyteVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1QuobyteVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetReadOnly

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRegistry

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### GetTenant

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUser

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVolume

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *IoK8sApiCoreV1QuobyteVolumeSource) SetVolume(v string)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


