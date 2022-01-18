# IoArgoprojEventsV1alpha1ArtifactLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configmap** | Pointer to [**IoK8sApiCoreV1ConfigMapKeySelector**](IoK8sApiCoreV1ConfigMapKeySelector.md) |  | [optional] 
**File** | Pointer to [**IoArgoprojEventsV1alpha1FileArtifact**](IoArgoprojEventsV1alpha1FileArtifact.md) |  | [optional] 
**Git** | Pointer to [**IoArgoprojEventsV1alpha1GitArtifact**](IoArgoprojEventsV1alpha1GitArtifact.md) |  | [optional] 
**Inline** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to [**IoArgoprojEventsV1alpha1Resource**](IoArgoprojEventsV1alpha1Resource.md) |  | [optional] 
**S3** | Pointer to [**IoArgoprojEventsV1alpha1S3Artifact**](IoArgoprojEventsV1alpha1S3Artifact.md) |  | [optional] 
**Url** | Pointer to [**IoArgoprojEventsV1alpha1URLArtifact**](IoArgoprojEventsV1alpha1URLArtifact.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1ArtifactLocation

`func NewIoArgoprojEventsV1alpha1ArtifactLocation() *IoArgoprojEventsV1alpha1ArtifactLocation`

NewIoArgoprojEventsV1alpha1ArtifactLocation instantiates a new IoArgoprojEventsV1alpha1ArtifactLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1ArtifactLocationWithDefaults

`func NewIoArgoprojEventsV1alpha1ArtifactLocationWithDefaults() *IoArgoprojEventsV1alpha1ArtifactLocation`

NewIoArgoprojEventsV1alpha1ArtifactLocationWithDefaults instantiates a new IoArgoprojEventsV1alpha1ArtifactLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigmap

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetConfigmap() IoK8sApiCoreV1ConfigMapKeySelector`

GetConfigmap returns the Configmap field if non-nil, zero value otherwise.

### GetConfigmapOk

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetConfigmapOk() (*IoK8sApiCoreV1ConfigMapKeySelector, bool)`

GetConfigmapOk returns a tuple with the Configmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigmap

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetConfigmap(v IoK8sApiCoreV1ConfigMapKeySelector)`

SetConfigmap sets Configmap field to given value.

### HasConfigmap

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasConfigmap() bool`

HasConfigmap returns a boolean if a field has been set.

### GetFile

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetFile() IoArgoprojEventsV1alpha1FileArtifact`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetFileOk() (*IoArgoprojEventsV1alpha1FileArtifact, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetFile(v IoArgoprojEventsV1alpha1FileArtifact)`

SetFile sets File field to given value.

### HasFile

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetGit

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetGit() IoArgoprojEventsV1alpha1GitArtifact`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetGitOk() (*IoArgoprojEventsV1alpha1GitArtifact, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetGit(v IoArgoprojEventsV1alpha1GitArtifact)`

SetGit sets Git field to given value.

### HasGit

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetInline

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetInline() string`

GetInline returns the Inline field if non-nil, zero value otherwise.

### GetInlineOk

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetInlineOk() (*string, bool)`

GetInlineOk returns a tuple with the Inline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInline

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetInline(v string)`

SetInline sets Inline field to given value.

### HasInline

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasInline() bool`

HasInline returns a boolean if a field has been set.

### GetResource

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetResource() IoArgoprojEventsV1alpha1Resource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetResourceOk() (*IoArgoprojEventsV1alpha1Resource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetResource(v IoArgoprojEventsV1alpha1Resource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetS3

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetS3() IoArgoprojEventsV1alpha1S3Artifact`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetS3Ok() (*IoArgoprojEventsV1alpha1S3Artifact, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetS3(v IoArgoprojEventsV1alpha1S3Artifact)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetUrl() IoArgoprojEventsV1alpha1URLArtifact`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) GetUrlOk() (*IoArgoprojEventsV1alpha1URLArtifact, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) SetUrl(v IoArgoprojEventsV1alpha1URLArtifact)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoArgoprojEventsV1alpha1ArtifactLocation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


