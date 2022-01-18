# IoArgoprojWorkflowV1alpha1TemplateRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterScope** | Pointer to **bool** | ClusterScope indicates the referred template is cluster scoped (i.e. a ClusterWorkflowTemplate). | [optional] 
**Name** | Pointer to **string** | Name is the resource name of the template. | [optional] 
**Template** | Pointer to **string** | Template is the name of referred template in the resource. | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1TemplateRef

`func NewIoArgoprojWorkflowV1alpha1TemplateRef() *IoArgoprojWorkflowV1alpha1TemplateRef`

NewIoArgoprojWorkflowV1alpha1TemplateRef instantiates a new IoArgoprojWorkflowV1alpha1TemplateRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1TemplateRefWithDefaults

`func NewIoArgoprojWorkflowV1alpha1TemplateRefWithDefaults() *IoArgoprojWorkflowV1alpha1TemplateRef`

NewIoArgoprojWorkflowV1alpha1TemplateRefWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1TemplateRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterScope

`func (o *IoArgoprojWorkflowV1alpha1TemplateRef) GetClusterScope() bool`

GetClusterScope returns the ClusterScope field if non-nil, zero value otherwise.

### GetClusterScopeOk

`func (o *IoArgoprojWorkflowV1alpha1TemplateRef) GetClusterScopeOk() (*bool, bool)`

GetClusterScopeOk returns a tuple with the ClusterScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterScope

`func (o *IoArgoprojWorkflowV1alpha1TemplateRef) SetClusterScope(v bool)`

SetClusterScope sets ClusterScope field to given value.

### HasClusterScope

`func (o *IoArgoprojWorkflowV1alpha1TemplateRef) HasClusterScope() bool`

HasClusterScope returns a boolean if a field has been set.

### GetName

`func (o *IoArgoprojWorkflowV1alpha1TemplateRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojWorkflowV1alpha1TemplateRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojWorkflowV1alpha1TemplateRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoArgoprojWorkflowV1alpha1TemplateRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplate

`func (o *IoArgoprojWorkflowV1alpha1TemplateRef) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoArgoprojWorkflowV1alpha1TemplateRef) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoArgoprojWorkflowV1alpha1TemplateRef) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IoArgoprojWorkflowV1alpha1TemplateRef) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


