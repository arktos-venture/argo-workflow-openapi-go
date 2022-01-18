# IoArgoprojWorkflowV1alpha1WorkflowCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateOptions** | Pointer to [**IoK8sApimachineryPkgApisMetaV1CreateOptions**](IoK8sApimachineryPkgApisMetaV1CreateOptions.md) |  | [optional] 
**InstanceID** | Pointer to **string** | This field is no longer used. | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**ServerDryRun** | Pointer to **bool** |  | [optional] 
**Workflow** | Pointer to [**IoArgoprojWorkflowV1alpha1Workflow**](IoArgoprojWorkflowV1alpha1Workflow.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1WorkflowCreateRequest

`func NewIoArgoprojWorkflowV1alpha1WorkflowCreateRequest() *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest`

NewIoArgoprojWorkflowV1alpha1WorkflowCreateRequest instantiates a new IoArgoprojWorkflowV1alpha1WorkflowCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1WorkflowCreateRequestWithDefaults

`func NewIoArgoprojWorkflowV1alpha1WorkflowCreateRequestWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest`

NewIoArgoprojWorkflowV1alpha1WorkflowCreateRequestWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateOptions

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetCreateOptions() IoK8sApimachineryPkgApisMetaV1CreateOptions`

GetCreateOptions returns the CreateOptions field if non-nil, zero value otherwise.

### GetCreateOptionsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetCreateOptionsOk() (*IoK8sApimachineryPkgApisMetaV1CreateOptions, bool)`

GetCreateOptionsOk returns a tuple with the CreateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOptions

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) SetCreateOptions(v IoK8sApimachineryPkgApisMetaV1CreateOptions)`

SetCreateOptions sets CreateOptions field to given value.

### HasCreateOptions

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) HasCreateOptions() bool`

HasCreateOptions returns a boolean if a field has been set.

### GetInstanceID

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetInstanceID() string`

GetInstanceID returns the InstanceID field if non-nil, zero value otherwise.

### GetInstanceIDOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetInstanceIDOk() (*string, bool)`

GetInstanceIDOk returns a tuple with the InstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceID

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) SetInstanceID(v string)`

SetInstanceID sets InstanceID field to given value.

### HasInstanceID

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) HasInstanceID() bool`

HasInstanceID returns a boolean if a field has been set.

### GetNamespace

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetServerDryRun

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetServerDryRun() bool`

GetServerDryRun returns the ServerDryRun field if non-nil, zero value otherwise.

### GetServerDryRunOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetServerDryRunOk() (*bool, bool)`

GetServerDryRunOk returns a tuple with the ServerDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDryRun

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) SetServerDryRun(v bool)`

SetServerDryRun sets ServerDryRun field to given value.

### HasServerDryRun

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) HasServerDryRun() bool`

HasServerDryRun returns a boolean if a field has been set.

### GetWorkflow

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetWorkflow() IoArgoprojWorkflowV1alpha1Workflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) GetWorkflowOk() (*IoArgoprojWorkflowV1alpha1Workflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) SetWorkflow(v IoArgoprojWorkflowV1alpha1Workflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *IoArgoprojWorkflowV1alpha1WorkflowCreateRequest) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


