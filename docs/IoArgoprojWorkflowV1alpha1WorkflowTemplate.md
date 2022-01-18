# IoArgoprojWorkflowV1alpha1WorkflowTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | 
**Spec** | [**IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec**](IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec.md) |  | 

## Methods

### NewIoArgoprojWorkflowV1alpha1WorkflowTemplate

`func NewIoArgoprojWorkflowV1alpha1WorkflowTemplate(metadata IoK8sApimachineryPkgApisMetaV1ObjectMeta, spec IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec, ) *IoArgoprojWorkflowV1alpha1WorkflowTemplate`

NewIoArgoprojWorkflowV1alpha1WorkflowTemplate instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1WorkflowTemplateWithDefaults

`func NewIoArgoprojWorkflowV1alpha1WorkflowTemplateWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowTemplate`

NewIoArgoprojWorkflowV1alpha1WorkflowTemplateWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) GetSpec() IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) GetSpecOk() (*IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplate) SetSpec(v IoArgoprojWorkflowV1alpha1WorkflowTemplateSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


