# IoArgoprojEventsV1alpha1K8SResourcePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backoff** | Pointer to [**IoArgoprojEventsV1alpha1Backoff**](IoArgoprojEventsV1alpha1Backoff.md) |  | [optional] 
**ErrorOnBackoffTimeout** | Pointer to **bool** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1K8SResourcePolicy

`func NewIoArgoprojEventsV1alpha1K8SResourcePolicy() *IoArgoprojEventsV1alpha1K8SResourcePolicy`

NewIoArgoprojEventsV1alpha1K8SResourcePolicy instantiates a new IoArgoprojEventsV1alpha1K8SResourcePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1K8SResourcePolicyWithDefaults

`func NewIoArgoprojEventsV1alpha1K8SResourcePolicyWithDefaults() *IoArgoprojEventsV1alpha1K8SResourcePolicy`

NewIoArgoprojEventsV1alpha1K8SResourcePolicyWithDefaults instantiates a new IoArgoprojEventsV1alpha1K8SResourcePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackoff

`func (o *IoArgoprojEventsV1alpha1K8SResourcePolicy) GetBackoff() IoArgoprojEventsV1alpha1Backoff`

GetBackoff returns the Backoff field if non-nil, zero value otherwise.

### GetBackoffOk

`func (o *IoArgoprojEventsV1alpha1K8SResourcePolicy) GetBackoffOk() (*IoArgoprojEventsV1alpha1Backoff, bool)`

GetBackoffOk returns a tuple with the Backoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoff

`func (o *IoArgoprojEventsV1alpha1K8SResourcePolicy) SetBackoff(v IoArgoprojEventsV1alpha1Backoff)`

SetBackoff sets Backoff field to given value.

### HasBackoff

`func (o *IoArgoprojEventsV1alpha1K8SResourcePolicy) HasBackoff() bool`

HasBackoff returns a boolean if a field has been set.

### GetErrorOnBackoffTimeout

`func (o *IoArgoprojEventsV1alpha1K8SResourcePolicy) GetErrorOnBackoffTimeout() bool`

GetErrorOnBackoffTimeout returns the ErrorOnBackoffTimeout field if non-nil, zero value otherwise.

### GetErrorOnBackoffTimeoutOk

`func (o *IoArgoprojEventsV1alpha1K8SResourcePolicy) GetErrorOnBackoffTimeoutOk() (*bool, bool)`

GetErrorOnBackoffTimeoutOk returns a tuple with the ErrorOnBackoffTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorOnBackoffTimeout

`func (o *IoArgoprojEventsV1alpha1K8SResourcePolicy) SetErrorOnBackoffTimeout(v bool)`

SetErrorOnBackoffTimeout sets ErrorOnBackoffTimeout field to given value.

### HasErrorOnBackoffTimeout

`func (o *IoArgoprojEventsV1alpha1K8SResourcePolicy) HasErrorOnBackoffTimeout() bool`

HasErrorOnBackoffTimeout returns a boolean if a field has been set.

### GetLabels

`func (o *IoArgoprojEventsV1alpha1K8SResourcePolicy) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IoArgoprojEventsV1alpha1K8SResourcePolicy) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IoArgoprojEventsV1alpha1K8SResourcePolicy) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IoArgoprojEventsV1alpha1K8SResourcePolicy) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


