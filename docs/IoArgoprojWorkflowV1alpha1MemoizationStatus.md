# IoArgoprojWorkflowV1alpha1MemoizationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheName** | **string** | Cache is the name of the cache that was used | 
**Hit** | **bool** | Hit indicates whether this node was created from a cache entry | 
**Key** | **string** | Key is the name of the key used for this node&#39;s cache | 

## Methods

### NewIoArgoprojWorkflowV1alpha1MemoizationStatus

`func NewIoArgoprojWorkflowV1alpha1MemoizationStatus(cacheName string, hit bool, key string, ) *IoArgoprojWorkflowV1alpha1MemoizationStatus`

NewIoArgoprojWorkflowV1alpha1MemoizationStatus instantiates a new IoArgoprojWorkflowV1alpha1MemoizationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1MemoizationStatusWithDefaults

`func NewIoArgoprojWorkflowV1alpha1MemoizationStatusWithDefaults() *IoArgoprojWorkflowV1alpha1MemoizationStatus`

NewIoArgoprojWorkflowV1alpha1MemoizationStatusWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1MemoizationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheName

`func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) GetCacheName() string`

GetCacheName returns the CacheName field if non-nil, zero value otherwise.

### GetCacheNameOk

`func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) GetCacheNameOk() (*string, bool)`

GetCacheNameOk returns a tuple with the CacheName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheName

`func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) SetCacheName(v string)`

SetCacheName sets CacheName field to given value.


### GetHit

`func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) GetHit() bool`

GetHit returns the Hit field if non-nil, zero value otherwise.

### GetHitOk

`func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) GetHitOk() (*bool, bool)`

GetHitOk returns a tuple with the Hit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHit

`func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) SetHit(v bool)`

SetHit sets Hit field to given value.


### GetKey

`func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IoArgoprojWorkflowV1alpha1MemoizationStatus) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


