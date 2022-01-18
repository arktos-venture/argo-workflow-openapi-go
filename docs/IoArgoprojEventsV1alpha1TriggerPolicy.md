# IoArgoprojEventsV1alpha1TriggerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**K8s** | Pointer to [**IoArgoprojEventsV1alpha1K8SResourcePolicy**](IoArgoprojEventsV1alpha1K8SResourcePolicy.md) |  | [optional] 
**Status** | Pointer to [**IoArgoprojEventsV1alpha1StatusPolicy**](IoArgoprojEventsV1alpha1StatusPolicy.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1TriggerPolicy

`func NewIoArgoprojEventsV1alpha1TriggerPolicy() *IoArgoprojEventsV1alpha1TriggerPolicy`

NewIoArgoprojEventsV1alpha1TriggerPolicy instantiates a new IoArgoprojEventsV1alpha1TriggerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1TriggerPolicyWithDefaults

`func NewIoArgoprojEventsV1alpha1TriggerPolicyWithDefaults() *IoArgoprojEventsV1alpha1TriggerPolicy`

NewIoArgoprojEventsV1alpha1TriggerPolicyWithDefaults instantiates a new IoArgoprojEventsV1alpha1TriggerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetK8s

`func (o *IoArgoprojEventsV1alpha1TriggerPolicy) GetK8s() IoArgoprojEventsV1alpha1K8SResourcePolicy`

GetK8s returns the K8s field if non-nil, zero value otherwise.

### GetK8sOk

`func (o *IoArgoprojEventsV1alpha1TriggerPolicy) GetK8sOk() (*IoArgoprojEventsV1alpha1K8SResourcePolicy, bool)`

GetK8sOk returns a tuple with the K8s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8s

`func (o *IoArgoprojEventsV1alpha1TriggerPolicy) SetK8s(v IoArgoprojEventsV1alpha1K8SResourcePolicy)`

SetK8s sets K8s field to given value.

### HasK8s

`func (o *IoArgoprojEventsV1alpha1TriggerPolicy) HasK8s() bool`

HasK8s returns a boolean if a field has been set.

### GetStatus

`func (o *IoArgoprojEventsV1alpha1TriggerPolicy) GetStatus() IoArgoprojEventsV1alpha1StatusPolicy`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoArgoprojEventsV1alpha1TriggerPolicy) GetStatusOk() (*IoArgoprojEventsV1alpha1StatusPolicy, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoArgoprojEventsV1alpha1TriggerPolicy) SetStatus(v IoArgoprojEventsV1alpha1StatusPolicy)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoArgoprojEventsV1alpha1TriggerPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


