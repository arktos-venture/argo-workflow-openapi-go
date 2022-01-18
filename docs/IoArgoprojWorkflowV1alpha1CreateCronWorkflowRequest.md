# IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateOptions** | Pointer to [**IoK8sApimachineryPkgApisMetaV1CreateOptions**](IoK8sApimachineryPkgApisMetaV1CreateOptions.md) |  | [optional] 
**CronWorkflow** | Pointer to [**IoArgoprojWorkflowV1alpha1CronWorkflow**](IoArgoprojWorkflowV1alpha1CronWorkflow.md) |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest

`func NewIoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest() *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest`

NewIoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest instantiates a new IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1CreateCronWorkflowRequestWithDefaults

`func NewIoArgoprojWorkflowV1alpha1CreateCronWorkflowRequestWithDefaults() *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest`

NewIoArgoprojWorkflowV1alpha1CreateCronWorkflowRequestWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateOptions

`func (o *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest) GetCreateOptions() IoK8sApimachineryPkgApisMetaV1CreateOptions`

GetCreateOptions returns the CreateOptions field if non-nil, zero value otherwise.

### GetCreateOptionsOk

`func (o *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest) GetCreateOptionsOk() (*IoK8sApimachineryPkgApisMetaV1CreateOptions, bool)`

GetCreateOptionsOk returns a tuple with the CreateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOptions

`func (o *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest) SetCreateOptions(v IoK8sApimachineryPkgApisMetaV1CreateOptions)`

SetCreateOptions sets CreateOptions field to given value.

### HasCreateOptions

`func (o *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest) HasCreateOptions() bool`

HasCreateOptions returns a boolean if a field has been set.

### GetCronWorkflow

`func (o *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest) GetCronWorkflow() IoArgoprojWorkflowV1alpha1CronWorkflow`

GetCronWorkflow returns the CronWorkflow field if non-nil, zero value otherwise.

### GetCronWorkflowOk

`func (o *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest) GetCronWorkflowOk() (*IoArgoprojWorkflowV1alpha1CronWorkflow, bool)`

GetCronWorkflowOk returns a tuple with the CronWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronWorkflow

`func (o *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest) SetCronWorkflow(v IoArgoprojWorkflowV1alpha1CronWorkflow)`

SetCronWorkflow sets CronWorkflow field to given value.

### HasCronWorkflow

`func (o *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest) HasCronWorkflow() bool`

HasCronWorkflow returns a boolean if a field has been set.

### GetNamespace

`func (o *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoArgoprojWorkflowV1alpha1CreateCronWorkflowRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


