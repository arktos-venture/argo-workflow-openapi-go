# IoArgoprojWorkflowV1alpha1ArtifactRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveLogs** | Pointer to **bool** | ArchiveLogs enables log archiving | [optional] 
**Artifactory** | Pointer to [**IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository**](IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository.md) |  | [optional] 
**Gcs** | Pointer to [**IoArgoprojWorkflowV1alpha1GCSArtifactRepository**](IoArgoprojWorkflowV1alpha1GCSArtifactRepository.md) |  | [optional] 
**Hdfs** | Pointer to [**IoArgoprojWorkflowV1alpha1HDFSArtifactRepository**](IoArgoprojWorkflowV1alpha1HDFSArtifactRepository.md) |  | [optional] 
**Oss** | Pointer to [**IoArgoprojWorkflowV1alpha1OSSArtifactRepository**](IoArgoprojWorkflowV1alpha1OSSArtifactRepository.md) |  | [optional] 
**S3** | Pointer to [**IoArgoprojWorkflowV1alpha1S3ArtifactRepository**](IoArgoprojWorkflowV1alpha1S3ArtifactRepository.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1ArtifactRepository

`func NewIoArgoprojWorkflowV1alpha1ArtifactRepository() *IoArgoprojWorkflowV1alpha1ArtifactRepository`

NewIoArgoprojWorkflowV1alpha1ArtifactRepository instantiates a new IoArgoprojWorkflowV1alpha1ArtifactRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryWithDefaults() *IoArgoprojWorkflowV1alpha1ArtifactRepository`

NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ArtifactRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) GetArchiveLogs() bool`

GetArchiveLogs returns the ArchiveLogs field if non-nil, zero value otherwise.

### GetArchiveLogsOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) GetArchiveLogsOk() (*bool, bool)`

GetArchiveLogsOk returns a tuple with the ArchiveLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) SetArchiveLogs(v bool)`

SetArchiveLogs sets ArchiveLogs field to given value.

### HasArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) HasArchiveLogs() bool`

HasArchiveLogs returns a boolean if a field has been set.

### GetArtifactory

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) GetArtifactory() IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository`

GetArtifactory returns the Artifactory field if non-nil, zero value otherwise.

### GetArtifactoryOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) GetArtifactoryOk() (*IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository, bool)`

GetArtifactoryOk returns a tuple with the Artifactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactory

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) SetArtifactory(v IoArgoprojWorkflowV1alpha1ArtifactoryArtifactRepository)`

SetArtifactory sets Artifactory field to given value.

### HasArtifactory

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) HasArtifactory() bool`

HasArtifactory returns a boolean if a field has been set.

### GetGcs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) GetGcs() IoArgoprojWorkflowV1alpha1GCSArtifactRepository`

GetGcs returns the Gcs field if non-nil, zero value otherwise.

### GetGcsOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) GetGcsOk() (*IoArgoprojWorkflowV1alpha1GCSArtifactRepository, bool)`

GetGcsOk returns a tuple with the Gcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) SetGcs(v IoArgoprojWorkflowV1alpha1GCSArtifactRepository)`

SetGcs sets Gcs field to given value.

### HasGcs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) HasGcs() bool`

HasGcs returns a boolean if a field has been set.

### GetHdfs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) GetHdfs() IoArgoprojWorkflowV1alpha1HDFSArtifactRepository`

GetHdfs returns the Hdfs field if non-nil, zero value otherwise.

### GetHdfsOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) GetHdfsOk() (*IoArgoprojWorkflowV1alpha1HDFSArtifactRepository, bool)`

GetHdfsOk returns a tuple with the Hdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) SetHdfs(v IoArgoprojWorkflowV1alpha1HDFSArtifactRepository)`

SetHdfs sets Hdfs field to given value.

### HasHdfs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) HasHdfs() bool`

HasHdfs returns a boolean if a field has been set.

### GetOss

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) GetOss() IoArgoprojWorkflowV1alpha1OSSArtifactRepository`

GetOss returns the Oss field if non-nil, zero value otherwise.

### GetOssOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) GetOssOk() (*IoArgoprojWorkflowV1alpha1OSSArtifactRepository, bool)`

GetOssOk returns a tuple with the Oss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOss

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) SetOss(v IoArgoprojWorkflowV1alpha1OSSArtifactRepository)`

SetOss sets Oss field to given value.

### HasOss

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) HasOss() bool`

HasOss returns a boolean if a field has been set.

### GetS3

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) GetS3() IoArgoprojWorkflowV1alpha1S3ArtifactRepository`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) GetS3Ok() (*IoArgoprojWorkflowV1alpha1S3ArtifactRepository, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) SetS3(v IoArgoprojWorkflowV1alpha1S3ArtifactRepository)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepository) HasS3() bool`

HasS3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


