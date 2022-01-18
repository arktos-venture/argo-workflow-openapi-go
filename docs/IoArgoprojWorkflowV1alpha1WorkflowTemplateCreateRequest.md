# IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateOptions** | Pointer to [**IoK8sApimachineryPkgApisMetaV1CreateOptions**](IoK8sApimachineryPkgApisMetaV1CreateOptions.md) |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Template** | Pointer to [**IoArgoprojWorkflowV1alpha1WorkflowTemplate**](IoArgoprojWorkflowV1alpha1WorkflowTemplate.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest

`func NewIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest() *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest`

NewIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequestWithDefaults

`func NewIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequestWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest`

NewIoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequestWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateOptions

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) GetCreateOptions() IoK8sApimachineryPkgApisMetaV1CreateOptions`

GetCreateOptions returns the CreateOptions field if non-nil, zero value otherwise.

### GetCreateOptionsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) GetCreateOptionsOk() (*IoK8sApimachineryPkgApisMetaV1CreateOptions, bool)`

GetCreateOptionsOk returns a tuple with the CreateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOptions

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) SetCreateOptions(v IoK8sApimachineryPkgApisMetaV1CreateOptions)`

SetCreateOptions sets CreateOptions field to given value.

### HasCreateOptions

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) HasCreateOptions() bool`

HasCreateOptions returns a boolean if a field has been set.

### GetNamespace

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTemplate

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) GetTemplate() IoArgoprojWorkflowV1alpha1WorkflowTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) GetTemplateOk() (*IoArgoprojWorkflowV1alpha1WorkflowTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) SetTemplate(v IoArgoprojWorkflowV1alpha1WorkflowTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IoArgoprojWorkflowV1alpha1WorkflowTemplateCreateRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


