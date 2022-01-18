# IoArgoprojWorkflowV1alpha1OSSLifecycleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarkDeletionAfterDays** | Pointer to **int32** | MarkDeletionAfterDays is the number of days before we delete objects in the bucket | [optional] 
**MarkInfrequentAccessAfterDays** | Pointer to **int32** | MarkInfrequentAccessAfterDays is the number of days before we convert the objects in the bucket to Infrequent Access (IA) storage type | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1OSSLifecycleRule

`func NewIoArgoprojWorkflowV1alpha1OSSLifecycleRule() *IoArgoprojWorkflowV1alpha1OSSLifecycleRule`

NewIoArgoprojWorkflowV1alpha1OSSLifecycleRule instantiates a new IoArgoprojWorkflowV1alpha1OSSLifecycleRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1OSSLifecycleRuleWithDefaults

`func NewIoArgoprojWorkflowV1alpha1OSSLifecycleRuleWithDefaults() *IoArgoprojWorkflowV1alpha1OSSLifecycleRule`

NewIoArgoprojWorkflowV1alpha1OSSLifecycleRuleWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1OSSLifecycleRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarkDeletionAfterDays

`func (o *IoArgoprojWorkflowV1alpha1OSSLifecycleRule) GetMarkDeletionAfterDays() int32`

GetMarkDeletionAfterDays returns the MarkDeletionAfterDays field if non-nil, zero value otherwise.

### GetMarkDeletionAfterDaysOk

`func (o *IoArgoprojWorkflowV1alpha1OSSLifecycleRule) GetMarkDeletionAfterDaysOk() (*int32, bool)`

GetMarkDeletionAfterDaysOk returns a tuple with the MarkDeletionAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkDeletionAfterDays

`func (o *IoArgoprojWorkflowV1alpha1OSSLifecycleRule) SetMarkDeletionAfterDays(v int32)`

SetMarkDeletionAfterDays sets MarkDeletionAfterDays field to given value.

### HasMarkDeletionAfterDays

`func (o *IoArgoprojWorkflowV1alpha1OSSLifecycleRule) HasMarkDeletionAfterDays() bool`

HasMarkDeletionAfterDays returns a boolean if a field has been set.

### GetMarkInfrequentAccessAfterDays

`func (o *IoArgoprojWorkflowV1alpha1OSSLifecycleRule) GetMarkInfrequentAccessAfterDays() int32`

GetMarkInfrequentAccessAfterDays returns the MarkInfrequentAccessAfterDays field if non-nil, zero value otherwise.

### GetMarkInfrequentAccessAfterDaysOk

`func (o *IoArgoprojWorkflowV1alpha1OSSLifecycleRule) GetMarkInfrequentAccessAfterDaysOk() (*int32, bool)`

GetMarkInfrequentAccessAfterDaysOk returns a tuple with the MarkInfrequentAccessAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkInfrequentAccessAfterDays

`func (o *IoArgoprojWorkflowV1alpha1OSSLifecycleRule) SetMarkInfrequentAccessAfterDays(v int32)`

SetMarkInfrequentAccessAfterDays sets MarkInfrequentAccessAfterDays field to given value.

### HasMarkInfrequentAccessAfterDays

`func (o *IoArgoprojWorkflowV1alpha1OSSLifecycleRule) HasMarkInfrequentAccessAfterDays() bool`

HasMarkInfrequentAccessAfterDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


