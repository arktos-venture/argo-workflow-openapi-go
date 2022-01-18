# IoArgoprojWorkflowV1alpha1Artifact

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

### NewIoArgoprojWorkflowV1alpha1Artifact

`func NewIoArgoprojWorkflowV1alpha1Artifact(name string, ) *IoArgoprojWorkflowV1alpha1Artifact`

NewIoArgoprojWorkflowV1alpha1Artifact instantiates a new IoArgoprojWorkflowV1alpha1Artifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ArtifactWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ArtifactWithDefaults() *IoArgoprojWorkflowV1alpha1Artifact`

NewIoArgoprojWorkflowV1alpha1ArtifactWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Artifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchive

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetArchive() IoArgoprojWorkflowV1alpha1ArchiveStrategy`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetArchiveOk() (*IoArgoprojWorkflowV1alpha1ArchiveStrategy, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetArchive(v IoArgoprojWorkflowV1alpha1ArchiveStrategy)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetArchiveLogs() bool`

GetArchiveLogs returns the ArchiveLogs field if non-nil, zero value otherwise.

### GetArchiveLogsOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetArchiveLogsOk() (*bool, bool)`

GetArchiveLogsOk returns a tuple with the ArchiveLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetArchiveLogs(v bool)`

SetArchiveLogs sets ArchiveLogs field to given value.

### HasArchiveLogs

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasArchiveLogs() bool`

HasArchiveLogs returns a boolean if a field has been set.

### GetArtifactory

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetArtifactory() IoArgoprojWorkflowV1alpha1ArtifactoryArtifact`

GetArtifactory returns the Artifactory field if non-nil, zero value otherwise.

### GetArtifactoryOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetArtifactoryOk() (*IoArgoprojWorkflowV1alpha1ArtifactoryArtifact, bool)`

GetArtifactoryOk returns a tuple with the Artifactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactory

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetArtifactory(v IoArgoprojWorkflowV1alpha1ArtifactoryArtifact)`

SetArtifactory sets Artifactory field to given value.

### HasArtifactory

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasArtifactory() bool`

HasArtifactory returns a boolean if a field has been set.

### GetFrom

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetFromExpression

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetFromExpression() string`

GetFromExpression returns the FromExpression field if non-nil, zero value otherwise.

### GetFromExpressionOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetFromExpressionOk() (*string, bool)`

GetFromExpressionOk returns a tuple with the FromExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromExpression

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetFromExpression(v string)`

SetFromExpression sets FromExpression field to given value.

### HasFromExpression

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasFromExpression() bool`

HasFromExpression returns a boolean if a field has been set.

### GetGcs

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetGcs() IoArgoprojWorkflowV1alpha1GCSArtifact`

GetGcs returns the Gcs field if non-nil, zero value otherwise.

### GetGcsOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetGcsOk() (*IoArgoprojWorkflowV1alpha1GCSArtifact, bool)`

GetGcsOk returns a tuple with the Gcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcs

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetGcs(v IoArgoprojWorkflowV1alpha1GCSArtifact)`

SetGcs sets Gcs field to given value.

### HasGcs

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasGcs() bool`

HasGcs returns a boolean if a field has been set.

### GetGit

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetGit() IoArgoprojWorkflowV1alpha1GitArtifact`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetGitOk() (*IoArgoprojWorkflowV1alpha1GitArtifact, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetGit(v IoArgoprojWorkflowV1alpha1GitArtifact)`

SetGit sets Git field to given value.

### HasGit

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetGlobalName

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetGlobalName() string`

GetGlobalName returns the GlobalName field if non-nil, zero value otherwise.

### GetGlobalNameOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetGlobalNameOk() (*string, bool)`

GetGlobalNameOk returns a tuple with the GlobalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalName

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetGlobalName(v string)`

SetGlobalName sets GlobalName field to given value.

### HasGlobalName

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasGlobalName() bool`

HasGlobalName returns a boolean if a field has been set.

### GetHdfs

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetHdfs() IoArgoprojWorkflowV1alpha1HDFSArtifact`

GetHdfs returns the Hdfs field if non-nil, zero value otherwise.

### GetHdfsOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetHdfsOk() (*IoArgoprojWorkflowV1alpha1HDFSArtifact, bool)`

GetHdfsOk returns a tuple with the Hdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfs

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetHdfs(v IoArgoprojWorkflowV1alpha1HDFSArtifact)`

SetHdfs sets Hdfs field to given value.

### HasHdfs

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasHdfs() bool`

HasHdfs returns a boolean if a field has been set.

### GetHttp

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetHttp() IoArgoprojWorkflowV1alpha1HTTPArtifact`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetHttpOk() (*IoArgoprojWorkflowV1alpha1HTTPArtifact, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetHttp(v IoArgoprojWorkflowV1alpha1HTTPArtifact)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetMode

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetName(v string)`

SetName sets Name field to given value.


### GetOptional

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetOss

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetOss() IoArgoprojWorkflowV1alpha1OSSArtifact`

GetOss returns the Oss field if non-nil, zero value otherwise.

### GetOssOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetOssOk() (*IoArgoprojWorkflowV1alpha1OSSArtifact, bool)`

GetOssOk returns a tuple with the Oss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOss

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetOss(v IoArgoprojWorkflowV1alpha1OSSArtifact)`

SetOss sets Oss field to given value.

### HasOss

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasOss() bool`

HasOss returns a boolean if a field has been set.

### GetPath

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRaw

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetRaw() IoArgoprojWorkflowV1alpha1RawArtifact`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetRawOk() (*IoArgoprojWorkflowV1alpha1RawArtifact, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetRaw(v IoArgoprojWorkflowV1alpha1RawArtifact)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### GetRecurseMode

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetRecurseMode() bool`

GetRecurseMode returns the RecurseMode field if non-nil, zero value otherwise.

### GetRecurseModeOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetRecurseModeOk() (*bool, bool)`

GetRecurseModeOk returns a tuple with the RecurseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurseMode

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetRecurseMode(v bool)`

SetRecurseMode sets RecurseMode field to given value.

### HasRecurseMode

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasRecurseMode() bool`

HasRecurseMode returns a boolean if a field has been set.

### GetS3

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetS3() IoArgoprojWorkflowV1alpha1S3Artifact`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetS3Ok() (*IoArgoprojWorkflowV1alpha1S3Artifact, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetS3(v IoArgoprojWorkflowV1alpha1S3Artifact)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetSubPath

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetSubPath() string`

GetSubPath returns the SubPath field if non-nil, zero value otherwise.

### GetSubPathOk

`func (o *IoArgoprojWorkflowV1alpha1Artifact) GetSubPathOk() (*string, bool)`

GetSubPathOk returns a tuple with the SubPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPath

`func (o *IoArgoprojWorkflowV1alpha1Artifact) SetSubPath(v string)`

SetSubPath sets SubPath field to given value.

### HasSubPath

`func (o *IoArgoprojWorkflowV1alpha1Artifact) HasSubPath() bool`

HasSubPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


