# IoArgoprojWorkflowV1alpha1Synchronization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mutex** | Pointer to [**IoArgoprojWorkflowV1alpha1Mutex**](IoArgoprojWorkflowV1alpha1Mutex.md) |  | [optional] 
**Semaphore** | Pointer to [**IoArgoprojWorkflowV1alpha1SemaphoreRef**](IoArgoprojWorkflowV1alpha1SemaphoreRef.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1Synchronization

`func NewIoArgoprojWorkflowV1alpha1Synchronization() *IoArgoprojWorkflowV1alpha1Synchronization`

NewIoArgoprojWorkflowV1alpha1Synchronization instantiates a new IoArgoprojWorkflowV1alpha1Synchronization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1SynchronizationWithDefaults

`func NewIoArgoprojWorkflowV1alpha1SynchronizationWithDefaults() *IoArgoprojWorkflowV1alpha1Synchronization`

NewIoArgoprojWorkflowV1alpha1SynchronizationWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Synchronization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMutex

`func (o *IoArgoprojWorkflowV1alpha1Synchronization) GetMutex() IoArgoprojWorkflowV1alpha1Mutex`

GetMutex returns the Mutex field if non-nil, zero value otherwise.

### GetMutexOk

`func (o *IoArgoprojWorkflowV1alpha1Synchronization) GetMutexOk() (*IoArgoprojWorkflowV1alpha1Mutex, bool)`

GetMutexOk returns a tuple with the Mutex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutex

`func (o *IoArgoprojWorkflowV1alpha1Synchronization) SetMutex(v IoArgoprojWorkflowV1alpha1Mutex)`

SetMutex sets Mutex field to given value.

### HasMutex

`func (o *IoArgoprojWorkflowV1alpha1Synchronization) HasMutex() bool`

HasMutex returns a boolean if a field has been set.

### GetSemaphore

`func (o *IoArgoprojWorkflowV1alpha1Synchronization) GetSemaphore() IoArgoprojWorkflowV1alpha1SemaphoreRef`

GetSemaphore returns the Semaphore field if non-nil, zero value otherwise.

### GetSemaphoreOk

`func (o *IoArgoprojWorkflowV1alpha1Synchronization) GetSemaphoreOk() (*IoArgoprojWorkflowV1alpha1SemaphoreRef, bool)`

GetSemaphoreOk returns a tuple with the Semaphore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemaphore

`func (o *IoArgoprojWorkflowV1alpha1Synchronization) SetSemaphore(v IoArgoprojWorkflowV1alpha1SemaphoreRef)`

SetSemaphore sets Semaphore field to given value.

### HasSemaphore

`func (o *IoArgoprojWorkflowV1alpha1Synchronization) HasSemaphore() bool`

HasSemaphore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


