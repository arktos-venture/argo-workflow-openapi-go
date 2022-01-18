# IoArgoprojWorkflowV1alpha1ArtifactPaths

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archive** | Pointer to [**IoArgoprojWorkflowV1alpha1ArchiveStrategy**](IoArgoprojWorkflowV1alpha1ArchiveStrategy.md) |  | [optional] 
**ArchiveLogs** | Pointer to **bool** | ArchiveLogs indicates if the container logs should be archived | [optional] 
**Artifactory** | Pointer to [**IoArgoprojWorkflowV1alpha1ArtifactoryArtifact**](IoArgoprojWorkflowV1alpha1ArtifactoryArtifact.md) |  | [optional] 
**From** | Pointer to **string** | From allows an artifact to reference an artifact from a previous step | [optional] 
**FromExpression** | Pointer to **string** | FromExpression, if defined, is evaluated to specify the value for the artifact | [optional] 
**Gcs** | Pointer to [**IoArgoprojWorkflowV1alpha1GCSArtifact**](IoArgoprojWorkflowV1alpha1GCSArtifact.md) |  | [optional] 
**Git** | Pointer to [**IoArgoprojWorkflowV1alpha1GitArtifact**](IoArgoprojWorkflowV1alpha1GitArtifact.md) |  | [optional] 
**GlobalName** | Pointer to **string** | GlobalName exports an output artifact to the global scope, making it available as &#39;{{io.argoproj.workflow.v1alpha1.outputs.artifacts.XXXX}} and in workflow.status.outputs.artifacts | [optional] 
**Hdfs** | Pointer to [**IoArgoprojWorkflowV1alpha1HDFSArtifact**](IoArgoprojWorkflowV1alpha1HDFSArtifact.md) |  | [optional] 
**Http** | Pointer to [**IoArgoprojWorkflowV1alpha1HTTPArtifact**](IoArgoprojWorkflowV1alpha1HTTPArtifact.md) |  | [optional] 
**Mode** | Pointer to **int32** | mode bits to use on this file, must be a value between 0 and 0777 set when loading input artifacts. | [optional] 
**Name** | **string** | name of the artifact. must be unique within a template&#39;s inputs/outputs. | 
**Optional** | Pointer to **bool** | Make Artifacts optional, if Artifacts doesn&#39;t generate or exist | [optional] 
**Oss** | Pointer to [**IoArgoprojWorkflowV1alpha1OSSArtifact**](IoArgoprojWorkflowV1alpha1OSSArtifact.md) |  | [optional] 
**Path** | Pointer to **string** | Path is the container path to the artifact | [optional] 
**Raw** | Pointer to [**IoArgoprojWorkflowV1alpha1RawArtifact**](IoArgoprojWorkflowV1alpha1RawArtifact.md) |  | [optional] 
**RecurseMode** | Pointer to **bool** | If mode is set, apply the permission recursively into the artifact if it is a folder | [optional] 
**S3** | Pointer to [**IoArgoprojWorkflowV1alpha1S3Artifact**](IoArgoprojWorkflowV1alpha1S3Artifact.md) |  | [optional] 
**SubPath** | Pointer to **string** | SubPath allows an artifact to be sourced from a subpath within the specified source | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1ArtifactPaths

`func NewIoArgoprojWorkflowV1alpha1ArtifactPaths(name string, ) *IoArgoprojWorkflowV1alpha1ArtifactPaths`

NewIoArgoprojWorkflowV1alpha1ArtifactPaths instantiates a new IoArgoprojWorkflowV1alpha1ArtifactPaths object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ArtifactPathsWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ArtifactPathsWithDefaults() *IoArgoprojWorkflowV1alpha1ArtifactPaths`

NewIoArgoprojWorkflowV1alpha1ArtifactPathsWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ArtifactPaths object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchive

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetArchive() IoArgoprojWorkflowV1alpha1ArchiveStrategy`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetArchiveOk() (*IoArgoprojWorkflowV1alpha1ArchiveStrategy, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetArchive(v IoArgoprojWorkflowV1alpha1ArchiveStrategy)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetArchiveLogs() bool`

GetArchiveLogs returns the ArchiveLogs field if non-nil, zero value otherwise.

### GetArchiveLogsOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetArchiveLogsOk() (*bool, bool)`

GetArchiveLogsOk returns a tuple with the ArchiveLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetArchiveLogs(v bool)`

SetArchiveLogs sets ArchiveLogs field to given value.

### HasArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasArchiveLogs() bool`

HasArchiveLogs returns a boolean if a field has been set.

### GetArtifactory

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetArtifactory() IoArgoprojWorkflowV1alpha1ArtifactoryArtifact`

GetArtifactory returns the Artifactory field if non-nil, zero value otherwise.

### GetArtifactoryOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetArtifactoryOk() (*IoArgoprojWorkflowV1alpha1ArtifactoryArtifact, bool)`

GetArtifactoryOk returns a tuple with the Artifactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactory

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetArtifactory(v IoArgoprojWorkflowV1alpha1ArtifactoryArtifact)`

SetArtifactory sets Artifactory field to given value.

### HasArtifactory

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasArtifactory() bool`

HasArtifactory returns a boolean if a field has been set.

### GetFrom

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetFromExpression

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetFromExpression() string`

GetFromExpression returns the FromExpression field if non-nil, zero value otherwise.

### GetFromExpressionOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetFromExpressionOk() (*string, bool)`

GetFromExpressionOk returns a tuple with the FromExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromExpression

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetFromExpression(v string)`

SetFromExpression sets FromExpression field to given value.

### HasFromExpression

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasFromExpression() bool`

HasFromExpression returns a boolean if a field has been set.

### GetGcs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetGcs() IoArgoprojWorkflowV1alpha1GCSArtifact`

GetGcs returns the Gcs field if non-nil, zero value otherwise.

### GetGcsOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetGcsOk() (*IoArgoprojWorkflowV1alpha1GCSArtifact, bool)`

GetGcsOk returns a tuple with the Gcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetGcs(v IoArgoprojWorkflowV1alpha1GCSArtifact)`

SetGcs sets Gcs field to given value.

### HasGcs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasGcs() bool`

HasGcs returns a boolean if a field has been set.

### GetGit

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetGit() IoArgoprojWorkflowV1alpha1GitArtifact`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetGitOk() (*IoArgoprojWorkflowV1alpha1GitArtifact, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetGit(v IoArgoprojWorkflowV1alpha1GitArtifact)`

SetGit sets Git field to given value.

### HasGit

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetGlobalName

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetGlobalName() string`

GetGlobalName returns the GlobalName field if non-nil, zero value otherwise.

### GetGlobalNameOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetGlobalNameOk() (*string, bool)`

GetGlobalNameOk returns a tuple with the GlobalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalName

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetGlobalName(v string)`

SetGlobalName sets GlobalName field to given value.

### HasGlobalName

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasGlobalName() bool`

HasGlobalName returns a boolean if a field has been set.

### GetHdfs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetHdfs() IoArgoprojWorkflowV1alpha1HDFSArtifact`

GetHdfs returns the Hdfs field if non-nil, zero value otherwise.

### GetHdfsOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetHdfsOk() (*IoArgoprojWorkflowV1alpha1HDFSArtifact, bool)`

GetHdfsOk returns a tuple with the Hdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetHdfs(v IoArgoprojWorkflowV1alpha1HDFSArtifact)`

SetHdfs sets Hdfs field to given value.

### HasHdfs

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasHdfs() bool`

HasHdfs returns a boolean if a field has been set.

### GetHttp

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetHttp() IoArgoprojWorkflowV1alpha1HTTPArtifact`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetHttpOk() (*IoArgoprojWorkflowV1alpha1HTTPArtifact, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetHttp(v IoArgoprojWorkflowV1alpha1HTTPArtifact)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetMode

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetName(v string)`

SetName sets Name field to given value.


### GetOptional

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetOss

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetOss() IoArgoprojWorkflowV1alpha1OSSArtifact`

GetOss returns the Oss field if non-nil, zero value otherwise.

### GetOssOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetOssOk() (*IoArgoprojWorkflowV1alpha1OSSArtifact, bool)`

GetOssOk returns a tuple with the Oss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOss

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetOss(v IoArgoprojWorkflowV1alpha1OSSArtifact)`

SetOss sets Oss field to given value.

### HasOss

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasOss() bool`

HasOss returns a boolean if a field has been set.

### GetPath

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRaw

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetRaw() IoArgoprojWorkflowV1alpha1RawArtifact`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetRawOk() (*IoArgoprojWorkflowV1alpha1RawArtifact, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetRaw(v IoArgoprojWorkflowV1alpha1RawArtifact)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetRecurseMode

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetRecurseMode() bool`

GetRecurseMode returns the RecurseMode field if non-nil, zero value otherwise.

### GetRecurseModeOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetRecurseModeOk() (*bool, bool)`

GetRecurseModeOk returns a tuple with the RecurseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurseMode

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetRecurseMode(v bool)`

SetRecurseMode sets RecurseMode field to given value.

### HasRecurseMode

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasRecurseMode() bool`

HasRecurseMode returns a boolean if a field has been set.

### GetS3

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetS3() IoArgoprojWorkflowV1alpha1S3Artifact`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetS3Ok() (*IoArgoprojWorkflowV1alpha1S3Artifact, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetS3(v IoArgoprojWorkflowV1alpha1S3Artifact)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetSubPath

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetSubPath() string`

GetSubPath returns the SubPath field if non-nil, zero value otherwise.

### GetSubPathOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) GetSubPathOk() (*string, bool)`

GetSubPathOk returns a tuple with the SubPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPath

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) SetSubPath(v string)`

SetSubPath sets SubPath field to given value.

### HasSubPath

`func (o *IoArgoprojWorkflowV1alpha1ArtifactPaths) HasSubPath() bool`

HasSubPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


