# IoArgoprojWorkflowV1alpha1Memoize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | [**IoArgoprojWorkflowV1alpha1Cache**](IoArgoprojWorkflowV1alpha1Cache.md) |  | 
**Key** | **string** | Key is the key to use as the caching key | 
**MaxAge** | **string** | MaxAge is the maximum age (e.g. \&quot;180s\&quot;, \&quot;24h\&quot;) of an entry that is still considered valid. If an entry is older than the MaxAge, it will be ignored. | 

## Methods

### NewIoArgoprojWorkflowV1alpha1Memoize

`func NewIoArgoprojWorkflowV1alpha1Memoize(cache IoArgoprojWorkflowV1alpha1Cache, key string, maxAge string, ) *IoArgoprojWorkflowV1alpha1Memoize`

NewIoArgoprojWorkflowV1alpha1Memoize instantiates a new IoArgoprojWorkflowV1alpha1Memoize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1MemoizeWithDefaults

`func NewIoArgoprojWorkflowV1alpha1MemoizeWithDefaults() *IoArgoprojWorkflowV1alpha1Memoize`

NewIoArgoprojWorkflowV1alpha1MemoizeWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Memoize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *IoArgoprojWorkflowV1alpha1Memoize) GetCache() IoArgoprojWorkflowV1alpha1Cache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *IoArgoprojWorkflowV1alpha1Memoize) GetCacheOk() (*IoArgoprojWorkflowV1alpha1Cache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *IoArgoprojWorkflowV1alpha1Memoize) SetCache(v IoArgoprojWorkflowV1alpha1Cache)`

SetCache sets Cache field to given value.


### GetKey

`func (o *IoArgoprojWorkflowV1alpha1Memoize) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IoArgoprojWorkflowV1alpha1Memoize) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IoArgoprojWorkflowV1alpha1Memoize) SetKey(v string)`

SetKey sets Key field to given value.


### GetMaxAge

`func (o *IoArgoprojWorkflowV1alpha1Memoize) GetMaxAge() string`

GetMaxAge returns the MaxAge field if non-nil, zero value otherwise.

### GetMaxAgeOk

`func (o *IoArgoprojWorkflowV1alpha1Memoize) GetMaxAgeOk() (*string, bool)`

GetMaxAgeOk returns a tuple with the MaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAge

`func (o *IoArgoprojWorkflowV1alpha1Memoize) SetMaxAge(v string)`

SetMaxAge sets MaxAge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


