# IoArgoprojWorkflowV1alpha1Submit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to [**IoArgoprojWorkflowV1alpha1Arguments**](IoArgoprojWorkflowV1alpha1Arguments.md) |  | [optional] 
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**WorkflowTemplateRef** | [**IoArgoprojWorkflowV1alpha1WorkflowTemplateRef**](IoArgoprojWorkflowV1alpha1WorkflowTemplateRef.md) |  | 

## Methods

### NewIoArgoprojWorkflowV1alpha1Submit

`func NewIoArgoprojWorkflowV1alpha1Submit(workflowTemplateRef IoArgoprojWorkflowV1alpha1WorkflowTemplateRef, ) *IoArgoprojWorkflowV1alpha1Submit`

NewIoArgoprojWorkflowV1alpha1Submit instantiates a new IoArgoprojWorkflowV1alpha1Submit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1SubmitWithDefaults

`func NewIoArgoprojWorkflowV1alpha1SubmitWithDefaults() *IoArgoprojWorkflowV1alpha1Submit`

NewIoArgoprojWorkflowV1alpha1SubmitWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Submit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *IoArgoprojWorkflowV1alpha1Submit) GetArguments() IoArgoprojWorkflowV1alpha1Arguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *IoArgoprojWorkflowV1alpha1Submit) GetArgumentsOk() (*IoArgoprojWorkflowV1alpha1Arguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *IoArgoprojWorkflowV1alpha1Submit) SetArguments(v IoArgoprojWorkflowV1alpha1Arguments)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *IoArgoprojWorkflowV1alpha1Submit) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojWorkflowV1alpha1Submit) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojWorkflowV1alpha1Submit) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojWorkflowV1alpha1Submit) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojWorkflowV1alpha1Submit) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetWorkflowTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1Submit) GetWorkflowTemplateRef() IoArgoprojWorkflowV1alpha1WorkflowTemplateRef`

GetWorkflowTemplateRef returns the WorkflowTemplateRef field if non-nil, zero value otherwise.

### GetWorkflowTemplateRefOk

`func (o *IoArgoprojWorkflowV1alpha1Submit) GetWorkflowTemplateRefOk() (*IoArgoprojWorkflowV1alpha1WorkflowTemplateRef, bool)`

GetWorkflowTemplateRefOk returns a tuple with the WorkflowTemplateRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1Submit) SetWorkflowTemplateRef(v IoArgoprojWorkflowV1alpha1WorkflowTemplateRef)`

SetWorkflowTemplateRef sets WorkflowTemplateRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


