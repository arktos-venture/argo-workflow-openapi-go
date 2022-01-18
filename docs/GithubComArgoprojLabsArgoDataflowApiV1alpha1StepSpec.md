# GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affinity** | Pointer to [**IoK8sApiCoreV1Affinity**](IoK8sApiCoreV1Affinity.md) |  | [optional] 
**Cat** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Cat**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Cat.md) |  | [optional] 
**Code** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Code**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Code.md) |  | [optional] 
**Container** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Container**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Container.md) |  | [optional] 
**Dedupe** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe.md) |  | [optional] 
**Expand** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Expand**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Expand.md) |  | [optional] 
**Filter** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Filter**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Filter.md) |  | [optional] 
**Flatten** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Flatten**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Flatten.md) |  | [optional] 
**Git** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Git**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Git.md) |  | [optional] 
**Group** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Group**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Group.md) |  | [optional] 
**ImagePullSecrets** | Pointer to [**[]IoK8sApiCoreV1LocalObjectReference**](IoK8sApiCoreV1LocalObjectReference.md) |  | [optional] 
**Map** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Map**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Map.md) |  | [optional] 
**Metadata** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Metadata**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Metadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeSelector** | Pointer to **map[string]string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**RestartPolicy** | Pointer to **string** |  | [optional] 
**Scale** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Scale**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Scale.md) |  | [optional] 
**ServiceAccountName** | Pointer to **string** |  | [optional] 
**Sidecar** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Sidecar**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Sidecar.md) |  | [optional] 
**Sinks** | Pointer to [**[]GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink.md) |  | [optional] 
**Sources** | Pointer to [**[]GithubComArgoprojLabsArgoDataflowApiV1alpha1Source**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Source.md) |  | [optional] 
**Terminator** | Pointer to **bool** |  | [optional] 
**Tolerations** | Pointer to [**[]IoK8sApiCoreV1Toleration**](IoK8sApiCoreV1Toleration.md) |  | [optional] 
**Volumes** | Pointer to [**[]IoK8sApiCoreV1Volume**](IoK8sApiCoreV1Volume.md) |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec() *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpecWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpecWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpecWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffinity

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetAffinity() IoK8sApiCoreV1Affinity`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetAffinityOk() (*IoK8sApiCoreV1Affinity, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetAffinity(v IoK8sApiCoreV1Affinity)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetCat

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetCat() GithubComArgoprojLabsArgoDataflowApiV1alpha1Cat`

GetCat returns the Cat field if non-nil, zero value otherwise.

### GetCatOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetCatOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Cat, bool)`

GetCatOk returns a tuple with the Cat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCat

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetCat(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Cat)`

SetCat sets Cat field to given value.

### HasCat

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasCat() bool`

HasCat returns a boolean if a field has been set.

### GetCode

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetCode() GithubComArgoprojLabsArgoDataflowApiV1alpha1Code`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetCodeOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Code, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetCode(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Code)`

SetCode sets Code field to given value.

### HasCode

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainer

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetContainer() GithubComArgoprojLabsArgoDataflowApiV1alpha1Container`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetContainerOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Container, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetContainer(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Container)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetDedupe

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetDedupe() GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe`

GetDedupe returns the Dedupe field if non-nil, zero value otherwise.

### GetDedupeOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetDedupeOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe, bool)`

GetDedupeOk returns a tuple with the Dedupe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedupe

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetDedupe(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe)`

SetDedupe sets Dedupe field to given value.

### HasDedupe

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasDedupe() bool`

HasDedupe returns a boolean if a field has been set.

### GetExpand

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetExpand() GithubComArgoprojLabsArgoDataflowApiV1alpha1Expand`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetExpandOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Expand, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetExpand(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Expand)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFilter

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetFilter() GithubComArgoprojLabsArgoDataflowApiV1alpha1Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetFilterOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetFilter(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFlatten

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetFlatten() GithubComArgoprojLabsArgoDataflowApiV1alpha1Flatten`

GetFlatten returns the Flatten field if non-nil, zero value otherwise.

### GetFlattenOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetFlattenOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Flatten, bool)`

GetFlattenOk returns a tuple with the Flatten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatten

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetFlatten(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Flatten)`

SetFlatten sets Flatten field to given value.

### HasFlatten

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasFlatten() bool`

HasFlatten returns a boolean if a field has been set.

### GetGit

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetGit() GithubComArgoprojLabsArgoDataflowApiV1alpha1Git`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetGitOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Git, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetGit(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Git)`

SetGit sets Git field to given value.

### HasGit

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetGroup

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetGroup() GithubComArgoprojLabsArgoDataflowApiV1alpha1Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetGroupOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetGroup(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Group)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetImagePullSecrets

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetImagePullSecrets() []IoK8sApiCoreV1LocalObjectReference`

GetImagePullSecrets returns the ImagePullSecrets field if non-nil, zero value otherwise.

### GetImagePullSecretsOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetImagePullSecretsOk() (*[]IoK8sApiCoreV1LocalObjectReference, bool)`

GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullSecrets

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetImagePullSecrets(v []IoK8sApiCoreV1LocalObjectReference)`

SetImagePullSecrets sets ImagePullSecrets field to given value.

### HasImagePullSecrets

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasImagePullSecrets() bool`

HasImagePullSecrets returns a boolean if a field has been set.

### GetMap

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetMap() GithubComArgoprojLabsArgoDataflowApiV1alpha1Map`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetMapOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Map, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetMap(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Map)`

SetMap sets Map field to given value.

### HasMap

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasMap() bool`

HasMap returns a boolean if a field has been set.

### GetMetadata

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetMetadata() GithubComArgoprojLabsArgoDataflowApiV1alpha1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetMetadataOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetMetadata(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeSelector

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetNodeSelector() map[string]string`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetNodeSelectorOk() (*map[string]string, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetNodeSelector(v map[string]string)`

SetNodeSelector sets NodeSelector field to given value.

### HasNodeSelector

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasNodeSelector() bool`

HasNodeSelector returns a boolean if a field has been set.

### GetReplicas

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRestartPolicy

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetRestartPolicy() string`

GetRestartPolicy returns the RestartPolicy field if non-nil, zero value otherwise.

### GetRestartPolicyOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetRestartPolicyOk() (*string, bool)`

GetRestartPolicyOk returns a tuple with the RestartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartPolicy

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetRestartPolicy(v string)`

SetRestartPolicy sets RestartPolicy field to given value.

### HasRestartPolicy

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasRestartPolicy() bool`

HasRestartPolicy returns a boolean if a field has been set.

### GetScale

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetScale() GithubComArgoprojLabsArgoDataflowApiV1alpha1Scale`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetScaleOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Scale, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetScale(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Scale)`

SetScale sets Scale field to given value.

### HasScale

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetServiceAccountName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.

### GetSidecar

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetSidecar() GithubComArgoprojLabsArgoDataflowApiV1alpha1Sidecar`

GetSidecar returns the Sidecar field if non-nil, zero value otherwise.

### GetSidecarOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetSidecarOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Sidecar, bool)`

GetSidecarOk returns a tuple with the Sidecar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSidecar

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetSidecar(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Sidecar)`

SetSidecar sets Sidecar field to given value.

### HasSidecar

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasSidecar() bool`

HasSidecar returns a boolean if a field has been set.

### GetSinks

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetSinks() []GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink`

GetSinks returns the Sinks field if non-nil, zero value otherwise.

### GetSinksOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetSinksOk() (*[]GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink, bool)`

GetSinksOk returns a tuple with the Sinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinks

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetSinks(v []GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink)`

SetSinks sets Sinks field to given value.

### HasSinks

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasSinks() bool`

HasSinks returns a boolean if a field has been set.

### GetSources

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetSources() []GithubComArgoprojLabsArgoDataflowApiV1alpha1Source`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetSourcesOk() (*[]GithubComArgoprojLabsArgoDataflowApiV1alpha1Source, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetSources(v []GithubComArgoprojLabsArgoDataflowApiV1alpha1Source)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTerminator

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetTerminator() bool`

GetTerminator returns the Terminator field if non-nil, zero value otherwise.

### GetTerminatorOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetTerminatorOk() (*bool, bool)`

GetTerminatorOk returns a tuple with the Terminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminator

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetTerminator(v bool)`

SetTerminator sets Terminator field to given value.

### HasTerminator

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasTerminator() bool`

HasTerminator returns a boolean if a field has been set.

### GetTolerations

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetTolerations() []IoK8sApiCoreV1Toleration`

GetTolerations returns the Tolerations field if non-nil, zero value otherwise.

### GetTolerationsOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetTolerationsOk() (*[]IoK8sApiCoreV1Toleration, bool)`

GetTolerationsOk returns a tuple with the Tolerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerations

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetTolerations(v []IoK8sApiCoreV1Toleration)`

SetTolerations sets Tolerations field to given value.

### HasTolerations

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasTolerations() bool`

HasTolerations returns a boolean if a field has been set.

### GetVolumes

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetVolumes() []IoK8sApiCoreV1Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) GetVolumesOk() (*[]IoK8sApiCoreV1Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) SetVolumes(v []IoK8sApiCoreV1Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


