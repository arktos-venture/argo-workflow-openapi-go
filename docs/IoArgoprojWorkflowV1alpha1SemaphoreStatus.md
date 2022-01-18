# IoArgoprojWorkflowV1alpha1SemaphoreStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Holding** | Pointer to [**[]IoArgoprojWorkflowV1alpha1SemaphoreHolding**](IoArgoprojWorkflowV1alpha1SemaphoreHolding.md) | Holding stores the list of resource acquired synchronization lock for workflows. | [optional] 
**Waiting** | Pointer to [**[]IoArgoprojWorkflowV1alpha1SemaphoreHolding**](IoArgoprojWorkflowV1alpha1SemaphoreHolding.md) | Waiting indicates the list of current synchronization lock holders. | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1SemaphoreStatus

`func NewIoArgoprojWorkflowV1alpha1SemaphoreStatus() *IoArgoprojWorkflowV1alpha1SemaphoreStatus`

NewIoArgoprojWorkflowV1alpha1SemaphoreStatus instantiates a new IoArgoprojWorkflowV1alpha1SemaphoreStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1SemaphoreStatusWithDefaults

`func NewIoArgoprojWorkflowV1alpha1SemaphoreStatusWithDefaults() *IoArgoprojWorkflowV1alpha1SemaphoreStatus`

NewIoArgoprojWorkflowV1alpha1SemaphoreStatusWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1SemaphoreStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHolding

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreStatus) GetHolding() []IoArgoprojWorkflowV1alpha1SemaphoreHolding`

GetHolding returns the Holding field if non-nil, zero value otherwise.

### GetHoldingOk

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreStatus) GetHoldingOk() (*[]IoArgoprojWorkflowV1alpha1SemaphoreHolding, bool)`

GetHoldingOk returns a tuple with the Holding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolding

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreStatus) SetHolding(v []IoArgoprojWorkflowV1alpha1SemaphoreHolding)`

SetHolding sets Holding field to given value.

### HasHolding

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreStatus) HasHolding() bool`

HasHolding returns a boolean if a field has been set.

### GetWaiting

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreStatus) GetWaiting() []IoArgoprojWorkflowV1alpha1SemaphoreHolding`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreStatus) GetWaitingOk() (*[]IoArgoprojWorkflowV1alpha1SemaphoreHolding, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreStatus) SetWaiting(v []IoArgoprojWorkflowV1alpha1SemaphoreHolding)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreStatus) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


