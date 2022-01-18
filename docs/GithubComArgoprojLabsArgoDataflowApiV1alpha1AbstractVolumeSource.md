# GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsElasticBlockStore** | Pointer to [**IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource**](IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource.md) |  | [optional] 
**AzureDisk** | Pointer to [**IoK8sApiCoreV1AzureDiskVolumeSource**](IoK8sApiCoreV1AzureDiskVolumeSource.md) |  | [optional] 
**AzureFile** | Pointer to [**IoK8sApiCoreV1AzureFileVolumeSource**](IoK8sApiCoreV1AzureFileVolumeSource.md) |  | [optional] 
**Cephfs** | Pointer to [**IoK8sApiCoreV1CephFSVolumeSource**](IoK8sApiCoreV1CephFSVolumeSource.md) |  | [optional] 
**Cinder** | Pointer to [**IoK8sApiCoreV1CinderVolumeSource**](IoK8sApiCoreV1CinderVolumeSource.md) |  | [optional] 
**ConfigMap** | Pointer to [**IoK8sApiCoreV1ConfigMapVolumeSource**](IoK8sApiCoreV1ConfigMapVolumeSource.md) |  | [optional] 
**Csi** | Pointer to [**IoK8sApiCoreV1CSIVolumeSource**](IoK8sApiCoreV1CSIVolumeSource.md) |  | [optional] 
**DownwardAPI** | Pointer to [**IoK8sApiCoreV1DownwardAPIVolumeSource**](IoK8sApiCoreV1DownwardAPIVolumeSource.md) |  | [optional] 
**EmptyDir** | Pointer to [**IoK8sApiCoreV1EmptyDirVolumeSource**](IoK8sApiCoreV1EmptyDirVolumeSource.md) |  | [optional] 
**Ephemeral** | Pointer to [**IoK8sApiCoreV1EphemeralVolumeSource**](IoK8sApiCoreV1EphemeralVolumeSource.md) |  | [optional] 
**Fc** | Pointer to [**IoK8sApiCoreV1FCVolumeSource**](IoK8sApiCoreV1FCVolumeSource.md) |  | [optional] 
**FlexVolume** | Pointer to [**IoK8sApiCoreV1FlexVolumeSource**](IoK8sApiCoreV1FlexVolumeSource.md) |  | [optional] 
**Flocker** | Pointer to [**IoK8sApiCoreV1FlockerVolumeSource**](IoK8sApiCoreV1FlockerVolumeSource.md) |  | [optional] 
**GcePersistentDisk** | Pointer to [**IoK8sApiCoreV1GCEPersistentDiskVolumeSource**](IoK8sApiCoreV1GCEPersistentDiskVolumeSource.md) |  | [optional] 
**GitRepo** | Pointer to [**IoK8sApiCoreV1GitRepoVolumeSource**](IoK8sApiCoreV1GitRepoVolumeSource.md) |  | [optional] 
**Glusterfs** | Pointer to [**IoK8sApiCoreV1GlusterfsVolumeSource**](IoK8sApiCoreV1GlusterfsVolumeSource.md) |  | [optional] 
**HostPath** | Pointer to [**IoK8sApiCoreV1HostPathVolumeSource**](IoK8sApiCoreV1HostPathVolumeSource.md) |  | [optional] 
**Iscsi** | Pointer to [**IoK8sApiCoreV1ISCSIVolumeSource**](IoK8sApiCoreV1ISCSIVolumeSource.md) |  | [optional] 
**Nfs** | Pointer to [**IoK8sApiCoreV1NFSVolumeSource**](IoK8sApiCoreV1NFSVolumeSource.md) |  | [optional] 
**PersistentVolumeClaim** | Pointer to [**IoK8sApiCoreV1PersistentVolumeClaimVolumeSource**](IoK8sApiCoreV1PersistentVolumeClaimVolumeSource.md) |  | [optional] 
**PhotonPersistentDisk** | Pointer to [**IoK8sApiCoreV1PhotonPersistentDiskVolumeSource**](IoK8sApiCoreV1PhotonPersistentDiskVolumeSource.md) |  | [optional] 
**PortworxVolume** | Pointer to [**IoK8sApiCoreV1PortworxVolumeSource**](IoK8sApiCoreV1PortworxVolumeSource.md) |  | [optional] 
**Projected** | Pointer to [**IoK8sApiCoreV1ProjectedVolumeSource**](IoK8sApiCoreV1ProjectedVolumeSource.md) |  | [optional] 
**Quobyte** | Pointer to [**IoK8sApiCoreV1QuobyteVolumeSource**](IoK8sApiCoreV1QuobyteVolumeSource.md) |  | [optional] 
**Rbd** | Pointer to [**IoK8sApiCoreV1RBDVolumeSource**](IoK8sApiCoreV1RBDVolumeSource.md) |  | [optional] 
**ScaleIO** | Pointer to [**IoK8sApiCoreV1ScaleIOVolumeSource**](IoK8sApiCoreV1ScaleIOVolumeSource.md) |  | [optional] 
**Secret** | Pointer to [**IoK8sApiCoreV1SecretVolumeSource**](IoK8sApiCoreV1SecretVolumeSource.md) |  | [optional] 
**Storageos** | Pointer to [**IoK8sApiCoreV1StorageOSVolumeSource**](IoK8sApiCoreV1StorageOSVolumeSource.md) |  | [optional] 
**VsphereVolume** | Pointer to [**IoK8sApiCoreV1VsphereVirtualDiskVolumeSource**](IoK8sApiCoreV1VsphereVirtualDiskVolumeSource.md) |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource() *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSourceWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSourceWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSourceWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsElasticBlockStore

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetAwsElasticBlockStore() IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource`

GetAwsElasticBlockStore returns the AwsElasticBlockStore field if non-nil, zero value otherwise.

### GetAwsElasticBlockStoreOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetAwsElasticBlockStoreOk() (*IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource, bool)`

GetAwsElasticBlockStoreOk returns a tuple with the AwsElasticBlockStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsElasticBlockStore

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetAwsElasticBlockStore(v IoK8sApiCoreV1AWSElasticBlockStoreVolumeSource)`

SetAwsElasticBlockStore sets AwsElasticBlockStore field to given value.

### HasAwsElasticBlockStore

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasAwsElasticBlockStore() bool`

HasAwsElasticBlockStore returns a boolean if a field has been set.

### GetAzureDisk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetAzureDisk() IoK8sApiCoreV1AzureDiskVolumeSource`

GetAzureDisk returns the AzureDisk field if non-nil, zero value otherwise.

### GetAzureDiskOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetAzureDiskOk() (*IoK8sApiCoreV1AzureDiskVolumeSource, bool)`

GetAzureDiskOk returns a tuple with the AzureDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDisk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetAzureDisk(v IoK8sApiCoreV1AzureDiskVolumeSource)`

SetAzureDisk sets AzureDisk field to given value.

### HasAzureDisk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasAzureDisk() bool`

HasAzureDisk returns a boolean if a field has been set.

### GetAzureFile

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetAzureFile() IoK8sApiCoreV1AzureFileVolumeSource`

GetAzureFile returns the AzureFile field if non-nil, zero value otherwise.

### GetAzureFileOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetAzureFileOk() (*IoK8sApiCoreV1AzureFileVolumeSource, bool)`

GetAzureFileOk returns a tuple with the AzureFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureFile

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetAzureFile(v IoK8sApiCoreV1AzureFileVolumeSource)`

SetAzureFile sets AzureFile field to given value.

### HasAzureFile

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasAzureFile() bool`

HasAzureFile returns a boolean if a field has been set.

### GetCephfs

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetCephfs() IoK8sApiCoreV1CephFSVolumeSource`

GetCephfs returns the Cephfs field if non-nil, zero value otherwise.

### GetCephfsOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetCephfsOk() (*IoK8sApiCoreV1CephFSVolumeSource, bool)`

GetCephfsOk returns a tuple with the Cephfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCephfs

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetCephfs(v IoK8sApiCoreV1CephFSVolumeSource)`

SetCephfs sets Cephfs field to given value.

### HasCephfs

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasCephfs() bool`

HasCephfs returns a boolean if a field has been set.

### GetCinder

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetCinder() IoK8sApiCoreV1CinderVolumeSource`

GetCinder returns the Cinder field if non-nil, zero value otherwise.

### GetCinderOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetCinderOk() (*IoK8sApiCoreV1CinderVolumeSource, bool)`

GetCinderOk returns a tuple with the Cinder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCinder

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetCinder(v IoK8sApiCoreV1CinderVolumeSource)`

SetCinder sets Cinder field to given value.

### HasCinder

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasCinder() bool`

HasCinder returns a boolean if a field has been set.

### GetConfigMap

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetConfigMap() IoK8sApiCoreV1ConfigMapVolumeSource`

GetConfigMap returns the ConfigMap field if non-nil, zero value otherwise.

### GetConfigMapOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetConfigMapOk() (*IoK8sApiCoreV1ConfigMapVolumeSource, bool)`

GetConfigMapOk returns a tuple with the ConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMap

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetConfigMap(v IoK8sApiCoreV1ConfigMapVolumeSource)`

SetConfigMap sets ConfigMap field to given value.

### HasConfigMap

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasConfigMap() bool`

HasConfigMap returns a boolean if a field has been set.

### GetCsi

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetCsi() IoK8sApiCoreV1CSIVolumeSource`

GetCsi returns the Csi field if non-nil, zero value otherwise.

### GetCsiOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetCsiOk() (*IoK8sApiCoreV1CSIVolumeSource, bool)`

GetCsiOk returns a tuple with the Csi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsi

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetCsi(v IoK8sApiCoreV1CSIVolumeSource)`

SetCsi sets Csi field to given value.

### HasCsi

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasCsi() bool`

HasCsi returns a boolean if a field has been set.

### GetDownwardAPI

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetDownwardAPI() IoK8sApiCoreV1DownwardAPIVolumeSource`

GetDownwardAPI returns the DownwardAPI field if non-nil, zero value otherwise.

### GetDownwardAPIOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetDownwardAPIOk() (*IoK8sApiCoreV1DownwardAPIVolumeSource, bool)`

GetDownwardAPIOk returns a tuple with the DownwardAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownwardAPI

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetDownwardAPI(v IoK8sApiCoreV1DownwardAPIVolumeSource)`

SetDownwardAPI sets DownwardAPI field to given value.

### HasDownwardAPI

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasDownwardAPI() bool`

HasDownwardAPI returns a boolean if a field has been set.

### GetEmptyDir

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetEmptyDir() IoK8sApiCoreV1EmptyDirVolumeSource`

GetEmptyDir returns the EmptyDir field if non-nil, zero value otherwise.

### GetEmptyDirOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetEmptyDirOk() (*IoK8sApiCoreV1EmptyDirVolumeSource, bool)`

GetEmptyDirOk returns a tuple with the EmptyDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyDir

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetEmptyDir(v IoK8sApiCoreV1EmptyDirVolumeSource)`

SetEmptyDir sets EmptyDir field to given value.

### HasEmptyDir

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasEmptyDir() bool`

HasEmptyDir returns a boolean if a field has been set.

### GetEphemeral

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetEphemeral() IoK8sApiCoreV1EphemeralVolumeSource`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetEphemeralOk() (*IoK8sApiCoreV1EphemeralVolumeSource, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetEphemeral(v IoK8sApiCoreV1EphemeralVolumeSource)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetFc

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetFc() IoK8sApiCoreV1FCVolumeSource`

GetFc returns the Fc field if non-nil, zero value otherwise.

### GetFcOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetFcOk() (*IoK8sApiCoreV1FCVolumeSource, bool)`

GetFcOk returns a tuple with the Fc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFc

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetFc(v IoK8sApiCoreV1FCVolumeSource)`

SetFc sets Fc field to given value.

### HasFc

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasFc() bool`

HasFc returns a boolean if a field has been set.

### GetFlexVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetFlexVolume() IoK8sApiCoreV1FlexVolumeSource`

GetFlexVolume returns the FlexVolume field if non-nil, zero value otherwise.

### GetFlexVolumeOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetFlexVolumeOk() (*IoK8sApiCoreV1FlexVolumeSource, bool)`

GetFlexVolumeOk returns a tuple with the FlexVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetFlexVolume(v IoK8sApiCoreV1FlexVolumeSource)`

SetFlexVolume sets FlexVolume field to given value.

### HasFlexVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasFlexVolume() bool`

HasFlexVolume returns a boolean if a field has been set.

### GetFlocker

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetFlocker() IoK8sApiCoreV1FlockerVolumeSource`

GetFlocker returns the Flocker field if non-nil, zero value otherwise.

### GetFlockerOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetFlockerOk() (*IoK8sApiCoreV1FlockerVolumeSource, bool)`

GetFlockerOk returns a tuple with the Flocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlocker

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetFlocker(v IoK8sApiCoreV1FlockerVolumeSource)`

SetFlocker sets Flocker field to given value.

### HasFlocker

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasFlocker() bool`

HasFlocker returns a boolean if a field has been set.

### GetGcePersistentDisk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetGcePersistentDisk() IoK8sApiCoreV1GCEPersistentDiskVolumeSource`

GetGcePersistentDisk returns the GcePersistentDisk field if non-nil, zero value otherwise.

### GetGcePersistentDiskOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetGcePersistentDiskOk() (*IoK8sApiCoreV1GCEPersistentDiskVolumeSource, bool)`

GetGcePersistentDiskOk returns a tuple with the GcePersistentDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcePersistentDisk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetGcePersistentDisk(v IoK8sApiCoreV1GCEPersistentDiskVolumeSource)`

SetGcePersistentDisk sets GcePersistentDisk field to given value.

### HasGcePersistentDisk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasGcePersistentDisk() bool`

HasGcePersistentDisk returns a boolean if a field has been set.

### GetGitRepo

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetGitRepo() IoK8sApiCoreV1GitRepoVolumeSource`

GetGitRepo returns the GitRepo field if non-nil, zero value otherwise.

### GetGitRepoOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetGitRepoOk() (*IoK8sApiCoreV1GitRepoVolumeSource, bool)`

GetGitRepoOk returns a tuple with the GitRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepo

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetGitRepo(v IoK8sApiCoreV1GitRepoVolumeSource)`

SetGitRepo sets GitRepo field to given value.

### HasGitRepo

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasGitRepo() bool`

HasGitRepo returns a boolean if a field has been set.

### GetGlusterfs

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetGlusterfs() IoK8sApiCoreV1GlusterfsVolumeSource`

GetGlusterfs returns the Glusterfs field if non-nil, zero value otherwise.

### GetGlusterfsOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetGlusterfsOk() (*IoK8sApiCoreV1GlusterfsVolumeSource, bool)`

GetGlusterfsOk returns a tuple with the Glusterfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlusterfs

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetGlusterfs(v IoK8sApiCoreV1GlusterfsVolumeSource)`

SetGlusterfs sets Glusterfs field to given value.

### HasGlusterfs

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasGlusterfs() bool`

HasGlusterfs returns a boolean if a field has been set.

### GetHostPath

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetHostPath() IoK8sApiCoreV1HostPathVolumeSource`

GetHostPath returns the HostPath field if non-nil, zero value otherwise.

### GetHostPathOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetHostPathOk() (*IoK8sApiCoreV1HostPathVolumeSource, bool)`

GetHostPathOk returns a tuple with the HostPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPath

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetHostPath(v IoK8sApiCoreV1HostPathVolumeSource)`

SetHostPath sets HostPath field to given value.

### HasHostPath

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasHostPath() bool`

HasHostPath returns a boolean if a field has been set.

### GetIscsi

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetIscsi() IoK8sApiCoreV1ISCSIVolumeSource`

GetIscsi returns the Iscsi field if non-nil, zero value otherwise.

### GetIscsiOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetIscsiOk() (*IoK8sApiCoreV1ISCSIVolumeSource, bool)`

GetIscsiOk returns a tuple with the Iscsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsi

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetIscsi(v IoK8sApiCoreV1ISCSIVolumeSource)`

SetIscsi sets Iscsi field to given value.

### HasIscsi

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasIscsi() bool`

HasIscsi returns a boolean if a field has been set.

### GetNfs

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetNfs() IoK8sApiCoreV1NFSVolumeSource`

GetNfs returns the Nfs field if non-nil, zero value otherwise.

### GetNfsOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetNfsOk() (*IoK8sApiCoreV1NFSVolumeSource, bool)`

GetNfsOk returns a tuple with the Nfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfs

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetNfs(v IoK8sApiCoreV1NFSVolumeSource)`

SetNfs sets Nfs field to given value.

### HasNfs

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasNfs() bool`

HasNfs returns a boolean if a field has been set.

### GetPersistentVolumeClaim

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetPersistentVolumeClaim() IoK8sApiCoreV1PersistentVolumeClaimVolumeSource`

GetPersistentVolumeClaim returns the PersistentVolumeClaim field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetPersistentVolumeClaimOk() (*IoK8sApiCoreV1PersistentVolumeClaimVolumeSource, bool)`

GetPersistentVolumeClaimOk returns a tuple with the PersistentVolumeClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaim

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetPersistentVolumeClaim(v IoK8sApiCoreV1PersistentVolumeClaimVolumeSource)`

SetPersistentVolumeClaim sets PersistentVolumeClaim field to given value.

### HasPersistentVolumeClaim

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasPersistentVolumeClaim() bool`

HasPersistentVolumeClaim returns a boolean if a field has been set.

### GetPhotonPersistentDisk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetPhotonPersistentDisk() IoK8sApiCoreV1PhotonPersistentDiskVolumeSource`

GetPhotonPersistentDisk returns the PhotonPersistentDisk field if non-nil, zero value otherwise.

### GetPhotonPersistentDiskOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetPhotonPersistentDiskOk() (*IoK8sApiCoreV1PhotonPersistentDiskVolumeSource, bool)`

GetPhotonPersistentDiskOk returns a tuple with the PhotonPersistentDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotonPersistentDisk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetPhotonPersistentDisk(v IoK8sApiCoreV1PhotonPersistentDiskVolumeSource)`

SetPhotonPersistentDisk sets PhotonPersistentDisk field to given value.

### HasPhotonPersistentDisk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasPhotonPersistentDisk() bool`

HasPhotonPersistentDisk returns a boolean if a field has been set.

### GetPortworxVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetPortworxVolume() IoK8sApiCoreV1PortworxVolumeSource`

GetPortworxVolume returns the PortworxVolume field if non-nil, zero value otherwise.

### GetPortworxVolumeOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetPortworxVolumeOk() (*IoK8sApiCoreV1PortworxVolumeSource, bool)`

GetPortworxVolumeOk returns a tuple with the PortworxVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortworxVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetPortworxVolume(v IoK8sApiCoreV1PortworxVolumeSource)`

SetPortworxVolume sets PortworxVolume field to given value.

### HasPortworxVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasPortworxVolume() bool`

HasPortworxVolume returns a boolean if a field has been set.

### GetProjected

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetProjected() IoK8sApiCoreV1ProjectedVolumeSource`

GetProjected returns the Projected field if non-nil, zero value otherwise.

### GetProjectedOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetProjectedOk() (*IoK8sApiCoreV1ProjectedVolumeSource, bool)`

GetProjectedOk returns a tuple with the Projected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjected

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetProjected(v IoK8sApiCoreV1ProjectedVolumeSource)`

SetProjected sets Projected field to given value.

### HasProjected

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasProjected() bool`

HasProjected returns a boolean if a field has been set.

### GetQuobyte

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetQuobyte() IoK8sApiCoreV1QuobyteVolumeSource`

GetQuobyte returns the Quobyte field if non-nil, zero value otherwise.

### GetQuobyteOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetQuobyteOk() (*IoK8sApiCoreV1QuobyteVolumeSource, bool)`

GetQuobyteOk returns a tuple with the Quobyte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuobyte

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetQuobyte(v IoK8sApiCoreV1QuobyteVolumeSource)`

SetQuobyte sets Quobyte field to given value.

### HasQuobyte

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasQuobyte() bool`

HasQuobyte returns a boolean if a field has been set.

### GetRbd

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetRbd() IoK8sApiCoreV1RBDVolumeSource`

GetRbd returns the Rbd field if non-nil, zero value otherwise.

### GetRbdOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetRbdOk() (*IoK8sApiCoreV1RBDVolumeSource, bool)`

GetRbdOk returns a tuple with the Rbd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbd

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetRbd(v IoK8sApiCoreV1RBDVolumeSource)`

SetRbd sets Rbd field to given value.

### HasRbd

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasRbd() bool`

HasRbd returns a boolean if a field has been set.

### GetScaleIO

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetScaleIO() IoK8sApiCoreV1ScaleIOVolumeSource`

GetScaleIO returns the ScaleIO field if non-nil, zero value otherwise.

### GetScaleIOOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetScaleIOOk() (*IoK8sApiCoreV1ScaleIOVolumeSource, bool)`

GetScaleIOOk returns a tuple with the ScaleIO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleIO

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetScaleIO(v IoK8sApiCoreV1ScaleIOVolumeSource)`

SetScaleIO sets ScaleIO field to given value.

### HasScaleIO

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasScaleIO() bool`

HasScaleIO returns a boolean if a field has been set.

### GetSecret

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetSecret() IoK8sApiCoreV1SecretVolumeSource`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetSecretOk() (*IoK8sApiCoreV1SecretVolumeSource, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetSecret(v IoK8sApiCoreV1SecretVolumeSource)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStorageos

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetStorageos() IoK8sApiCoreV1StorageOSVolumeSource`

GetStorageos returns the Storageos field if non-nil, zero value otherwise.

### GetStorageosOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetStorageosOk() (*IoK8sApiCoreV1StorageOSVolumeSource, bool)`

GetStorageosOk returns a tuple with the Storageos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageos

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetStorageos(v IoK8sApiCoreV1StorageOSVolumeSource)`

SetStorageos sets Storageos field to given value.

### HasStorageos

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasStorageos() bool`

HasStorageos returns a boolean if a field has been set.

### GetVsphereVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetVsphereVolume() IoK8sApiCoreV1VsphereVirtualDiskVolumeSource`

GetVsphereVolume returns the VsphereVolume field if non-nil, zero value otherwise.

### GetVsphereVolumeOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) GetVsphereVolumeOk() (*IoK8sApiCoreV1VsphereVirtualDiskVolumeSource, bool)`

GetVsphereVolumeOk returns a tuple with the VsphereVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsphereVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) SetVsphereVolume(v IoK8sApiCoreV1VsphereVirtualDiskVolumeSource)`

SetVsphereVolume sets VsphereVolume field to given value.

### HasVsphereVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractVolumeSource) HasVsphereVolume() bool`

HasVsphereVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


