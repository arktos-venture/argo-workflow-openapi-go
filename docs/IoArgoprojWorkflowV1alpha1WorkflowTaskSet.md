# IoArgoprojWorkflowV1alpha1WorkflowTaskSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | 
**Spec** | [**IoArgoprojWorkflowV1alpha1WorkflowTaskSetSpec**](IoArgoprojWorkflowV1alpha1WorkflowTaskSetSpec.md) |  | 
**Status** | Pointer to [**IoArgoprojWorkflowV1alpha1WorkflowTaskSetStatus**](IoArgoprojWorkflowV1alpha1WorkflowTaskSetStatus.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1WorkflowTaskSet

`func NewIoArgoprojWorkflowV1alpha1WorkflowTaskSet(metadata IoK8sApimachineryPkgApisMetaV1ObjectMeta, spec IoArgoprojWorkflowV1alpha1WorkflowTaskSetSpec, ) *IoArgoprojWorkflowV1alpha1WorkflowTaskSet`

NewIoArgoprojWorkflowV1alpha1WorkflowTaskSet instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTaskSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1WorkflowTaskSetWithDefaults

`func NewIoArgoprojWorkflowV1alpha1WorkflowTaskSetWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowTaskSet`

NewIoArgoprojWorkflowV1alpha1WorkflowTaskSetWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTaskSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.


### GetSpec

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) GetSpec() IoArgoprojWorkflowV1alpha1WorkflowTaskSetSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) GetSpecOk() (*IoArgoprojWorkflowV1alpha1WorkflowTaskSetSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) SetSpec(v IoArgoprojWorkflowV1alpha1WorkflowTaskSetSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) GetStatus() IoArgoprojWorkflowV1alpha1WorkflowTaskSetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) GetStatusOk() (*IoArgoprojWorkflowV1alpha1WorkflowTaskSetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) SetStatus(v IoArgoprojWorkflowV1alpha1WorkflowTaskSetStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTaskSet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


