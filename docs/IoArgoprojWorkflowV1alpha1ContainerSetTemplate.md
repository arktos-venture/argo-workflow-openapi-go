# IoArgoprojWorkflowV1alpha1ContainerSetTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | [**[]IoArgoprojWorkflowV1alpha1ContainerNode**](IoArgoprojWorkflowV1alpha1ContainerNode.md) |  | 
**VolumeMounts** | Pointer to [**[]IoK8sApiCoreV1VolumeMount**](IoK8sApiCoreV1VolumeMount.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1ContainerSetTemplate

`func NewIoArgoprojWorkflowV1alpha1ContainerSetTemplate(containers []IoArgoprojWorkflowV1alpha1ContainerNode, ) *IoArgoprojWorkflowV1alpha1ContainerSetTemplate`

NewIoArgoprojWorkflowV1alpha1ContainerSetTemplate instantiates a new IoArgoprojWorkflowV1alpha1ContainerSetTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ContainerSetTemplateWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ContainerSetTemplateWithDefaults() *IoArgoprojWorkflowV1alpha1ContainerSetTemplate`

NewIoArgoprojWorkflowV1alpha1ContainerSetTemplateWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ContainerSetTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *IoArgoprojWorkflowV1alpha1ContainerSetTemplate) GetContainers() []IoArgoprojWorkflowV1alpha1ContainerNode`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *IoArgoprojWorkflowV1alpha1ContainerSetTemplate) GetContainersOk() (*[]IoArgoprojWorkflowV1alpha1ContainerNode, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *IoArgoprojWorkflowV1alpha1ContainerSetTemplate) SetContainers(v []IoArgoprojWorkflowV1alpha1ContainerNode)`

SetContainers sets Containers field to given value.


### GetVolumeMounts

`func (o *IoArgoprojWorkflowV1alpha1ContainerSetTemplate) GetVolumeMounts() []IoK8sApiCoreV1VolumeMount`

GetVolumeMounts returns the VolumeMounts field if non-nil, zero value otherwise.

### GetVolumeMountsOk

`func (o *IoArgoprojWorkflowV1alpha1ContainerSetTemplate) GetVolumeMountsOk() (*[]IoK8sApiCoreV1VolumeMount, bool)`

GetVolumeMountsOk returns a tuple with the VolumeMounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMounts

`func (o *IoArgoprojWorkflowV1alpha1ContainerSetTemplate) SetVolumeMounts(v []IoK8sApiCoreV1VolumeMount)`

SetVolumeMounts sets VolumeMounts field to given value.

### HasVolumeMounts

`func (o *IoArgoprojWorkflowV1alpha1ContainerSetTemplate) HasVolumeMounts() bool`

HasVolumeMounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


