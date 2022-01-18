# IoArgoprojEventsV1alpha1ArgoWorkflowTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupVersionResource** | Pointer to [**IoK8sApimachineryPkgApisMetaV1GroupVersionResource**](IoK8sApimachineryPkgApisMetaV1GroupVersionResource.md) |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) |  | [optional] 
**Source** | Pointer to [**IoArgoprojEventsV1alpha1ArtifactLocation**](IoArgoprojEventsV1alpha1ArtifactLocation.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1ArgoWorkflowTrigger

`func NewIoArgoprojEventsV1alpha1ArgoWorkflowTrigger() *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger`

NewIoArgoprojEventsV1alpha1ArgoWorkflowTrigger instantiates a new IoArgoprojEventsV1alpha1ArgoWorkflowTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1ArgoWorkflowTriggerWithDefaults

`func NewIoArgoprojEventsV1alpha1ArgoWorkflowTriggerWithDefaults() *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger`

NewIoArgoprojEventsV1alpha1ArgoWorkflowTriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1ArgoWorkflowTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupVersionResource

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) GetGroupVersionResource() IoK8sApimachineryPkgApisMetaV1GroupVersionResource`

GetGroupVersionResource returns the GroupVersionResource field if non-nil, zero value otherwise.

### GetGroupVersionResourceOk

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) GetGroupVersionResourceOk() (*IoK8sApimachineryPkgApisMetaV1GroupVersionResource, bool)`

GetGroupVersionResourceOk returns a tuple with the GroupVersionResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupVersionResource

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) SetGroupVersionResource(v IoK8sApimachineryPkgApisMetaV1GroupVersionResource)`

SetGroupVersionResource sets GroupVersionResource field to given value.

### HasGroupVersionResource

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) HasGroupVersionResource() bool`

HasGroupVersionResource returns a boolean if a field has been set.

### GetOperation

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetParameters

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSource

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) GetSource() IoArgoprojEventsV1alpha1ArtifactLocation`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) GetSourceOk() (*IoArgoprojEventsV1alpha1ArtifactLocation, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) SetSource(v IoArgoprojEventsV1alpha1ArtifactLocation)`

SetSource sets Source field to given value.

### HasSource

`func (o *IoArgoprojEventsV1alpha1ArgoWorkflowTrigger) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


