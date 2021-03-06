# IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | Pointer to **string** | Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \&quot;ext4\&quot;, \&quot;xfs\&quot;, \&quot;ntfs\&quot;. Implicitly inferred to be \&quot;ext4\&quot; if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | [optional] 
**Partition** | Pointer to **int32** | The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as \&quot;1\&quot;. Similarly, the volume partition for /dev/sda is \&quot;0\&quot; (or you can leave the property empty). | [optional] 
**ReadOnly** | Pointer to **bool** | Specify \&quot;true\&quot; to force and set the ReadOnly property in VolumeMounts to \&quot;true\&quot;. If omitted, the default is \&quot;false\&quot;. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | [optional] 
**VolumeID** | **string** | Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore | 

## Methods

### NewIoK8sApiCoreV1AWSElasticBlockStoreVolumeSource

`func NewIoK8sApiCoreV1AWSElasticBlockStoreVolumeSource(volumeID string, ) *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource`

NewIoK8sApiCoreV1AWSElasticBlockStoreVolumeSource instantiates a new IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1AWSElasticBlockStoreVolumeSourceWithDefaults

`func NewIoK8sApiCoreV1AWSElasticBlockStoreVolumeSourceWithDefaults() *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource`

NewIoK8sApiCoreV1AWSElasticBlockStoreVolumeSourceWithDefaults instantiates a new IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsType

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetPartition

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetReadOnly

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetVolumeID

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) GetVolumeID() string`

GetVolumeID returns the VolumeID field if non-nil, zero value otherwise.

### GetVolumeIDOk

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) GetVolumeIDOk() (*string, bool)`

GetVolumeIDOk returns a tuple with the VolumeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeID

`func (o *IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource) SetVolumeID(v string)`

SetVolumeID sets VolumeID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


