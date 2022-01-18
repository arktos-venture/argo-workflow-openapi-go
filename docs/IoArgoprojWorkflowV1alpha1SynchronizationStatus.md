# IoArgoprojWorkflowV1alpha1SynchronizationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mutex** | Pointer to [**IoArgoprojWorkflowV1alpha1MutexStatus**](IoArgoprojWorkflowV1alpha1MutexStatus.md) |  | [optional] 
**Semaphore** | Pointer to [**IoArgoprojWorkflowV1alpha1SemaphoreStatus**](IoArgoprojWorkflowV1alpha1SemaphoreStatus.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1SynchronizationStatus

`func NewIoArgoprojWorkflowV1alpha1SynchronizationStatus() *IoArgoprojWorkflowV1alpha1SynchronizationStatus`

NewIoArgoprojWorkflowV1alpha1SynchronizationStatus instantiates a new IoArgoprojWorkflowV1alpha1SynchronizationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1SynchronizationStatusWithDefaults

`func NewIoArgoprojWorkflowV1alpha1SynchronizationStatusWithDefaults() *IoArgoprojWorkflowV1alpha1SynchronizationStatus`

NewIoArgoprojWorkflowV1alpha1SynchronizationStatusWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1SynchronizationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMutex

`func (o *IoArgoprojWorkflowV1alpha1SynchronizationStatus) GetMutex() IoArgoprojWorkflowV1alpha1MutexStatus`

GetMutex returns the Mutex field if non-nil, zero value otherwise.

### GetMutexOk

`func (o *IoArgoprojWorkflowV1alpha1SynchronizationStatus) GetMutexOk() (*IoArgoprojWorkflowV1alpha1MutexStatus, bool)`

GetMutexOk returns a tuple with the Mutex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutex

`func (o *IoArgoprojWorkflowV1alpha1SynchronizationStatus) SetMutex(v IoArgoprojWorkflowV1alpha1MutexStatus)`

SetMutex sets Mutex field to given value.

### HasMutex

`func (o *IoArgoprojWorkflowV1alpha1SynchronizationStatus) HasMutex() bool`

HasMutex returns a boolean if a field has been set.

### GetSemaphore

`func (o *IoArgoprojWorkflowV1alpha1SynchronizationStatus) GetSemaphore() IoArgoprojWorkflowV1alpha1SemaphoreStatus`

GetSemaphore returns the Semaphore field if non-nil, zero value otherwise.

### GetSemaphoreOk

`func (o *IoArgoprojWorkflowV1alpha1SynchronizationStatus) GetSemaphoreOk() (*IoArgoprojWorkflowV1alpha1SemaphoreStatus, bool)`

GetSemaphoreOk returns a tuple with the Semaphore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemaphore

`func (o *IoArgoprojWorkflowV1alpha1SynchronizationStatus) SetSemaphore(v IoArgoprojWorkflowV1alpha1SemaphoreStatus)`

SetSemaphore sets Semaphore field to given value.

### HasSemaphore

`func (o *IoArgoprojWorkflowV1alpha1SynchronizationStatus) HasSemaphore() bool`

HasSemaphore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


