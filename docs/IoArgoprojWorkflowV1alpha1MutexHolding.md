# IoArgoprojWorkflowV1alpha1MutexHolding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Holder** | Pointer to **string** | Holder is a reference to the object which holds the Mutex. Holding Scenario:   1. Current workflow&#39;s NodeID which is holding the lock.      e.g: ${NodeID} Waiting Scenario:   1. Current workflow or other workflow NodeID which is holding the lock.      e.g: ${WorkflowName}/${NodeID} | [optional] 
**Mutex** | Pointer to **string** | Reference for the mutex e.g: ${namespace}/mutex/${mutexName} | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1MutexHolding

`func NewIoArgoprojWorkflowV1alpha1MutexHolding() *IoArgoprojWorkflowV1alpha1MutexHolding`

NewIoArgoprojWorkflowV1alpha1MutexHolding instantiates a new IoArgoprojWorkflowV1alpha1MutexHolding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1MutexHoldingWithDefaults

`func NewIoArgoprojWorkflowV1alpha1MutexHoldingWithDefaults() *IoArgoprojWorkflowV1alpha1MutexHolding`

NewIoArgoprojWorkflowV1alpha1MutexHoldingWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1MutexHolding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHolder

`func (o *IoArgoprojWorkflowV1alpha1MutexHolding) GetHolder() string`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *IoArgoprojWorkflowV1alpha1MutexHolding) GetHolderOk() (*string, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *IoArgoprojWorkflowV1alpha1MutexHolding) SetHolder(v string)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *IoArgoprojWorkflowV1alpha1MutexHolding) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetMutex

`func (o *IoArgoprojWorkflowV1alpha1MutexHolding) GetMutex() string`

GetMutex returns the Mutex field if non-nil, zero value otherwise.

### GetMutexOk

`func (o *IoArgoprojWorkflowV1alpha1MutexHolding) GetMutexOk() (*string, bool)`

GetMutexOk returns a tuple with the Mutex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutex

`func (o *IoArgoprojWorkflowV1alpha1MutexHolding) SetMutex(v string)`

SetMutex sets Mutex field to given value.

### HasMutex

`func (o *IoArgoprojWorkflowV1alpha1MutexHolding) HasMutex() bool`

HasMutex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


