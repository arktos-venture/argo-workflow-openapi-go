# IoArgoprojWorkflowV1alpha1RetryAffinity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeAntiAffinity** | Pointer to **map[string]interface{}** | RetryNodeAntiAffinity is a placeholder for future expansion, only empty nodeAntiAffinity is allowed. In order to prevent running steps on the same host, it uses \&quot;kubernetes.io/hostname\&quot;. | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1RetryAffinity

`func NewIoArgoprojWorkflowV1alpha1RetryAffinity() *IoArgoprojWorkflowV1alpha1RetryAffinity`

NewIoArgoprojWorkflowV1alpha1RetryAffinity instantiates a new IoArgoprojWorkflowV1alpha1RetryAffinity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1RetryAffinityWithDefaults

`func NewIoArgoprojWorkflowV1alpha1RetryAffinityWithDefaults() *IoArgoprojWorkflowV1alpha1RetryAffinity`

NewIoArgoprojWorkflowV1alpha1RetryAffinityWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1RetryAffinity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeAntiAffinity

`func (o *IoArgoprojWorkflowV1alpha1RetryAffinity) GetNodeAntiAffinity() map[string]interface{}`

GetNodeAntiAffinity returns the NodeAntiAffinity field if non-nil, zero value otherwise.

### GetNodeAntiAffinityOk

`func (o *IoArgoprojWorkflowV1alpha1RetryAffinity) GetNodeAntiAffinityOk() (*map[string]interface{}, bool)`

GetNodeAntiAffinityOk returns a tuple with the NodeAntiAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAntiAffinity

`func (o *IoArgoprojWorkflowV1alpha1RetryAffinity) SetNodeAntiAffinity(v map[string]interface{})`

SetNodeAntiAffinity sets NodeAntiAffinity field to given value.

### HasNodeAntiAffinity

`func (o *IoArgoprojWorkflowV1alpha1RetryAffinity) HasNodeAntiAffinity() bool`

HasNodeAntiAffinity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


