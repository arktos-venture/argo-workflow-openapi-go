# IoArgoprojWorkflowV1alpha1MutexStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Holding** | Pointer to [**[]IoArgoprojWorkflowV1alpha1MutexHolding**](IoArgoprojWorkflowV1alpha1MutexHolding.md) | Holding is a list of mutexes and their respective objects that are held by mutex lock for this io.argoproj.workflow.v1alpha1. | [optional] 
**Waiting** | Pointer to [**[]IoArgoprojWorkflowV1alpha1MutexHolding**](IoArgoprojWorkflowV1alpha1MutexHolding.md) | Waiting is a list of mutexes and their respective objects this workflow is waiting for. | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1MutexStatus

`func NewIoArgoprojWorkflowV1alpha1MutexStatus() *IoArgoprojWorkflowV1alpha1MutexStatus`

NewIoArgoprojWorkflowV1alpha1MutexStatus instantiates a new IoArgoprojWorkflowV1alpha1MutexStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1MutexStatusWithDefaults

`func NewIoArgoprojWorkflowV1alpha1MutexStatusWithDefaults() *IoArgoprojWorkflowV1alpha1MutexStatus`

NewIoArgoprojWorkflowV1alpha1MutexStatusWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1MutexStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHolding

`func (o *IoArgoprojWorkflowV1alpha1MutexStatus) GetHolding() []IoArgoprojWorkflowV1alpha1MutexHolding`

GetHolding returns the Holding field if non-nil, zero value otherwise.

### GetHoldingOk

`func (o *IoArgoprojWorkflowV1alpha1MutexStatus) GetHoldingOk() (*[]IoArgoprojWorkflowV1alpha1MutexHolding, bool)`

GetHoldingOk returns a tuple with the Holding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolding

`func (o *IoArgoprojWorkflowV1alpha1MutexStatus) SetHolding(v []IoArgoprojWorkflowV1alpha1MutexHolding)`

SetHolding sets Holding field to given value.

### HasHolding

`func (o *IoArgoprojWorkflowV1alpha1MutexStatus) HasHolding() bool`

HasHolding returns a boolean if a field has been set.

### GetWaiting

`func (o *IoArgoprojWorkflowV1alpha1MutexStatus) GetWaiting() []IoArgoprojWorkflowV1alpha1MutexHolding`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *IoArgoprojWorkflowV1alpha1MutexStatus) GetWaitingOk() (*[]IoArgoprojWorkflowV1alpha1MutexHolding, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *IoArgoprojWorkflowV1alpha1MutexStatus) SetWaiting(v []IoArgoprojWorkflowV1alpha1MutexHolding)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *IoArgoprojWorkflowV1alpha1MutexStatus) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


