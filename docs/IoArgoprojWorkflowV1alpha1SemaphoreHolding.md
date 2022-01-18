# IoArgoprojWorkflowV1alpha1SemaphoreHolding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Holders** | Pointer to **[]string** | Holders stores the list of current holder names in the io.argoproj.workflow.v1alpha1. | [optional] 
**Semaphore** | Pointer to **string** | Semaphore stores the semaphore name. | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1SemaphoreHolding

`func NewIoArgoprojWorkflowV1alpha1SemaphoreHolding() *IoArgoprojWorkflowV1alpha1SemaphoreHolding`

NewIoArgoprojWorkflowV1alpha1SemaphoreHolding instantiates a new IoArgoprojWorkflowV1alpha1SemaphoreHolding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1SemaphoreHoldingWithDefaults

`func NewIoArgoprojWorkflowV1alpha1SemaphoreHoldingWithDefaults() *IoArgoprojWorkflowV1alpha1SemaphoreHolding`

NewIoArgoprojWorkflowV1alpha1SemaphoreHoldingWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1SemaphoreHolding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHolders

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreHolding) GetHolders() []string`

GetHolders returns the Holders field if non-nil, zero value otherwise.

### GetHoldersOk

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreHolding) GetHoldersOk() (*[]string, bool)`

GetHoldersOk returns a tuple with the Holders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolders

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreHolding) SetHolders(v []string)`

SetHolders sets Holders field to given value.

### HasHolders

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreHolding) HasHolders() bool`

HasHolders returns a boolean if a field has been set.

### GetSemaphore

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreHolding) GetSemaphore() string`

GetSemaphore returns the Semaphore field if non-nil, zero value otherwise.

### GetSemaphoreOk

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreHolding) GetSemaphoreOk() (*string, bool)`

GetSemaphoreOk returns a tuple with the Semaphore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemaphore

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreHolding) SetSemaphore(v string)`

SetSemaphore sets Semaphore field to given value.

### HasSemaphore

`func (o *IoArgoprojWorkflowV1alpha1SemaphoreHolding) HasSemaphore() bool`

HasSemaphore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


