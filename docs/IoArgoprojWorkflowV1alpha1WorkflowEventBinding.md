# IoArgoprojWorkflowV1alpha1WorkflowEventBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | 
**Spec** | [**IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec**](IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec.md) |  | 

## Methods

### NewIoArgoprojWorkflowV1alpha1WorkflowEventBinding

`func NewIoArgoprojWorkflowV1alpha1WorkflowEventBinding(metadata IoK8sApimachineryPkgApisMetaV1ObjectMeta, spec IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec, ) *IoArgoprojWorkflowV1alpha1WorkflowEventBinding`

NewIoArgoprojWorkflowV1alpha1WorkflowEventBinding instantiates a new IoArgoprojWorkflowV1alpha1WorkflowEventBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingWithDefaults

`func NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowEventBinding`

NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowEventBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetSpec() IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) GetSpecOk() (*IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBinding) SetSpec(v IoArgoprojWorkflowV1alpha1WorkflowEventBindingSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


