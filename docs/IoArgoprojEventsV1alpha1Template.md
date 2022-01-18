# IoArgoprojEventsV1alpha1Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affinity** | Pointer to [**IoK8sApiCoreV1Affinity**](IoK8sApiCoreV1Affinity.md) |  | [optional] 
**Container** | Pointer to [**IoK8sApiCoreV1Container**](IoK8sApiCoreV1Container.md) |  | [optional] 
**ImagePullSecrets** | Pointer to [**[]IoK8sApiCoreV1LocalObjectReference**](IoK8sApiCoreV1LocalObjectReference.md) |  | [optional] 
**Metadata** | Pointer to [**IoArgoprojEventsV1alpha1Metadata**](IoArgoprojEventsV1alpha1Metadata.md) |  | [optional] 
**NodeSelector** | Pointer to **map[string]string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**PriorityClassName** | Pointer to **string** |  | [optional] 
**SecurityContext** | Pointer to [**IoK8sApiCoreV1PodSecurityContext**](IoK8sApiCoreV1PodSecurityContext.md) |  | [optional] 
**ServiceAccountName** | Pointer to **string** |  | [optional] 
**Tolerations** | Pointer to [**[]IoK8sApiCoreV1Toleration**](IoK8sApiCoreV1Toleration.md) |  | [optional] 
**Volumes** | Pointer to [**[]IoK8sApiCoreV1Volume**](IoK8sApiCoreV1Volume.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1Template

`func NewIoArgoprojEventsV1alpha1Template() *IoArgoprojEventsV1alpha1Template`

NewIoArgoprojEventsV1alpha1Template instantiates a new IoArgoprojEventsV1alpha1Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1TemplateWithDefaults

`func NewIoArgoprojEventsV1alpha1TemplateWithDefaults() *IoArgoprojEventsV1alpha1Template`

NewIoArgoprojEventsV1alpha1TemplateWithDefaults instantiates a new IoArgoprojEventsV1alpha1Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffinity

`func (o *IoArgoprojEventsV1alpha1Template) GetAffinity() IoK8sApiCoreV1Affinity`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *IoArgoprojEventsV1alpha1Template) GetAffinityOk() (*IoK8sApiCoreV1Affinity, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *IoArgoprojEventsV1alpha1Template) SetAffinity(v IoK8sApiCoreV1Affinity)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *IoArgoprojEventsV1alpha1Template) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetContainer

`func (o *IoArgoprojEventsV1alpha1Template) GetContainer() IoK8sApiCoreV1Container`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *IoArgoprojEventsV1alpha1Template) GetContainerOk() (*IoK8sApiCoreV1Container, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *IoArgoprojEventsV1alpha1Template) SetContainer(v IoK8sApiCoreV1Container)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *IoArgoprojEventsV1alpha1Template) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetImagePullSecrets

`func (o *IoArgoprojEventsV1alpha1Template) GetImagePullSecrets() []IoK8sApiCoreV1LocalObjectReference`

GetImagePullSecrets returns the ImagePullSecrets field if non-nil, zero value otherwise.

### GetImagePullSecretsOk

`func (o *IoArgoprojEventsV1alpha1Template) GetImagePullSecretsOk() (*[]IoK8sApiCoreV1LocalObjectReference, bool)`

GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullSecrets

`func (o *IoArgoprojEventsV1alpha1Template) SetImagePullSecrets(v []IoK8sApiCoreV1LocalObjectReference)`

SetImagePullSecrets sets ImagePullSecrets field to given value.

### HasImagePullSecrets

`func (o *IoArgoprojEventsV1alpha1Template) HasImagePullSecrets() bool`

HasImagePullSecrets returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1Template) GetMetadata() IoArgoprojEventsV1alpha1Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1Template) GetMetadataOk() (*IoArgoprojEventsV1alpha1Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1Template) SetMetadata(v IoArgoprojEventsV1alpha1Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1Template) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNodeSelector

`func (o *IoArgoprojEventsV1alpha1Template) GetNodeSelector() map[string]string`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *IoArgoprojEventsV1alpha1Template) GetNodeSelectorOk() (*map[string]string, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *IoArgoprojEventsV1alpha1Template) SetNodeSelector(v map[string]string)`

SetNodeSelector sets NodeSelector field to given value.

### HasNodeSelector

`func (o *IoArgoprojEventsV1alpha1Template) HasNodeSelector() bool`

HasNodeSelector returns a boolean if a field has been set.

### GetPriority

`func (o *IoArgoprojEventsV1alpha1Template) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IoArgoprojEventsV1alpha1Template) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IoArgoprojEventsV1alpha1Template) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IoArgoprojEventsV1alpha1Template) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPriorityClassName

`func (o *IoArgoprojEventsV1alpha1Template) GetPriorityClassName() string`

GetPriorityClassName returns the PriorityClassName field if non-nil, zero value otherwise.

### GetPriorityClassNameOk

`func (o *IoArgoprojEventsV1alpha1Template) GetPriorityClassNameOk() (*string, bool)`

GetPriorityClassNameOk returns a tuple with the PriorityClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityClassName

`func (o *IoArgoprojEventsV1alpha1Template) SetPriorityClassName(v string)`

SetPriorityClassName sets PriorityClassName field to given value.

### HasPriorityClassName

`func (o *IoArgoprojEventsV1alpha1Template) HasPriorityClassName() bool`

HasPriorityClassName returns a boolean if a field has been set.

### GetSecurityContext

`func (o *IoArgoprojEventsV1alpha1Template) GetSecurityContext() IoK8sApiCoreV1PodSecurityContext`

GetSecurityContext returns the SecurityContext field if non-nil, zero value otherwise.

### GetSecurityContextOk

`func (o *IoArgoprojEventsV1alpha1Template) GetSecurityContextOk() (*IoK8sApiCoreV1PodSecurityContext, bool)`

GetSecurityContextOk returns a tuple with the SecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityContext

`func (o *IoArgoprojEventsV1alpha1Template) SetSecurityContext(v IoK8sApiCoreV1PodSecurityContext)`

SetSecurityContext sets SecurityContext field to given value.

### HasSecurityContext

`func (o *IoArgoprojEventsV1alpha1Template) HasSecurityContext() bool`

HasSecurityContext returns a boolean if a field has been set.

### GetServiceAccountName

`func (o *IoArgoprojEventsV1alpha1Template) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *IoArgoprojEventsV1alpha1Template) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *IoArgoprojEventsV1alpha1Template) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *IoArgoprojEventsV1alpha1Template) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.

### GetTolerations

`func (o *IoArgoprojEventsV1alpha1Template) GetTolerations() []IoK8sApiCoreV1Toleration`

GetTolerations returns the Tolerations field if non-nil, zero value otherwise.

### GetTolerationsOk

`func (o *IoArgoprojEventsV1alpha1Template) GetTolerationsOk() (*[]IoK8sApiCoreV1Toleration, bool)`

GetTolerationsOk returns a tuple with the Tolerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerations

`func (o *IoArgoprojEventsV1alpha1Template) SetTolerations(v []IoK8sApiCoreV1Toleration)`

SetTolerations sets Tolerations field to given value.

### HasTolerations

`func (o *IoArgoprojEventsV1alpha1Template) HasTolerations() bool`

HasTolerations returns a boolean if a field has been set.

### GetVolumes

`func (o *IoArgoprojEventsV1alpha1Template) GetVolumes() []IoK8sApiCoreV1Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *IoArgoprojEventsV1alpha1Template) GetVolumesOk() (*[]IoK8sApiCoreV1Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *IoArgoprojEventsV1alpha1Template) SetVolumes(v []IoK8sApiCoreV1Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *IoArgoprojEventsV1alpha1Template) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


