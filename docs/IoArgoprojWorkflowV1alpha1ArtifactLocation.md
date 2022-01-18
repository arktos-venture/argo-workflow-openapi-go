# IoArgoprojWorkflowV1alpha1ArtifactLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveLogs** | Pointer to **bool** | ArchiveLogs indicates if the container logs should be archived | [optional] 
**Artifactory** | Pointer to [**IoArgoprojWorkflowV1alpha1ArtifactoryArtifact**](IoArgoprojWorkflowV1alpha1ArtifactoryArtifact.md) |  | [optional] 
**Gcs** | Pointer to [**IoArgoprojWorkflowV1alpha1GCSArtifact**](IoArgoprojWorkflowV1alpha1GCSArtifact.md) |  | [optional] 
**Git** | Pointer to [**IoArgoprojWorkflowV1alpha1GitArtifact**](IoArgoprojWorkflowV1alpha1GitArtifact.md) |  | [optional] 
**Hdfs** | Pointer to [**IoArgoprojWorkflowV1alpha1HDFSArtifact**](IoArgoprojWorkflowV1alpha1HDFSArtifact.md) |  | [optional] 
**Http** | Pointer to [**IoArgoprojWorkflowV1alpha1HTTPArtifact**](IoArgoprojWorkflowV1alpha1HTTPArtifact.md) |  | [optional] 
**Oss** | Pointer to [**IoArgoprojWorkflowV1alpha1OSSArtifact**](IoArgoprojWorkflowV1alpha1OSSArtifact.md) |  | [optional] 
**Raw** | Pointer to [**IoArgoprojWorkflowV1alpha1RawArtifact**](IoArgoprojWorkflowV1alpha1RawArtifact.md) |  | [optional] 
**S3** | Pointer to [**IoArgoprojWorkflowV1alpha1S3Artifact**](IoArgoprojWorkflowV1alpha1S3Artifact.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1ArtifactLocation

`func NewIoArgoprojWorkflowV1alpha1ArtifactLocation() *IoArgoprojWorkflowV1alpha1ArtifactLocation`

NewIoArgoprojWorkflowV1alpha1ArtifactLocation instantiates a new IoArgoprojWorkflowV1alpha1ArtifactLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ArtifactLocationWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ArtifactLocationWithDefaults() *IoArgoprojWorkflowV1alpha1ArtifactLocation`

NewIoArgoprojWorkflowV1alpha1ArtifactLocationWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ArtifactLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetArchiveLogs() bool`

GetArchiveLogs returns the ArchiveLogs field if non-nil, zero value otherwise.

### GetArchiveLogsOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetArchiveLogsOk() (*bool, bool)`

GetArchiveLogsOk returns a tuple with the ArchiveLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetArchiveLogs(v bool)`

SetArchiveLogs sets ArchiveLogs field to given value.

### HasArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasArchiveLogs() bool`

HasArchiveLogs returns a boolean if a field has been set.

### GetArtifactory

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetArtifactory() IoArgoprojWorkflowV1alpha1ArtifactoryArtifact`

GetArtifactory returns the Artifactory field if non-nil, zero value otherwise.

### GetArtifactoryOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetArtifactoryOk() (*IoArgoprojWorkflowV1alpha1ArtifactoryArtifact, bool)`

GetArtifactoryOk returns a tuple with the Artifactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactory

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetArtifactory(v IoArgoprojWorkflowV1alpha1ArtifactoryArtifact)`

SetArtifactory sets Artifactory field to given value.

### HasArtifactory

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasArtifactory() bool`

HasArtifactory returns a boolean if a field has been set.

### GetGcs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetGcs() IoArgoprojWorkflowV1alpha1GCSArtifact`

GetGcs returns the Gcs field if non-nil, zero value otherwise.

### GetGcsOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetGcsOk() (*IoArgoprojWorkflowV1alpha1GCSArtifact, bool)`

GetGcsOk returns a tuple with the Gcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetGcs(v IoArgoprojWorkflowV1alpha1GCSArtifact)`

SetGcs sets Gcs field to given value.

### HasGcs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasGcs() bool`

HasGcs returns a boolean if a field has been set.

### GetGit

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetGit() IoArgoprojWorkflowV1alpha1GitArtifact`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetGitOk() (*IoArgoprojWorkflowV1alpha1GitArtifact, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetGit(v IoArgoprojWorkflowV1alpha1GitArtifact)`

SetGit sets Git field to given value.

### HasGit

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetHdfs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetHdfs() IoArgoprojWorkflowV1alpha1HDFSArtifact`

GetHdfs returns the Hdfs field if non-nil, zero value otherwise.

### GetHdfsOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetHdfsOk() (*IoArgoprojWorkflowV1alpha1HDFSArtifact, bool)`

GetHdfsOk returns a tuple with the Hdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetHdfs(v IoArgoprojWorkflowV1alpha1HDFSArtifact)`

SetHdfs sets Hdfs field to given value.

### HasHdfs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasHdfs() bool`

HasHdfs returns a boolean if a field has been set.

### GetHttp

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetHttp() IoArgoprojWorkflowV1alpha1HTTPArtifact`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetHttpOk() (*IoArgoprojWorkflowV1alpha1HTTPArtifact, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetHttp(v IoArgoprojWorkflowV1alpha1HTTPArtifact)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetOss

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetOss() IoArgoprojWorkflowV1alpha1OSSArtifact`

GetOss returns the Oss field if non-nil, zero value otherwise.

### GetOssOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetOssOk() (*IoArgoprojWorkflowV1alpha1OSSArtifact, bool)`

GetOssOk returns a tuple with the Oss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOss

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetOss(v IoArgoprojWorkflowV1alpha1OSSArtifact)`

SetOss sets Oss field to given value.

### HasOss

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasOss() bool`

HasOss returns a boolean if a field has been set.

### GetRaw

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetRaw() IoArgoprojWorkflowV1alpha1RawArtifact`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetRawOk() (*IoArgoprojWorkflowV1alpha1RawArtifact, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetRaw(v IoArgoprojWorkflowV1alpha1RawArtifact)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetS3

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetS3() IoArgoprojWorkflowV1alpha1S3Artifact`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetS3Ok() (*IoArgoprojWorkflowV1alpha1S3Artifact, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetS3(v IoArgoprojWorkflowV1alpha1S3Artifact)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasS3() bool`

HasS3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


