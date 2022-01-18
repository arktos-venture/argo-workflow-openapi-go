# GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastScaledAt** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Selector** | Pointer to **string** |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus() *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatusWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatusWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatusWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastScaledAt

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetLastScaledAt() time.Time`

GetLastScaledAt returns the LastScaledAt field if non-nil, zero value otherwise.

### GetLastScaledAtOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetLastScaledAtOk() (*time.Time, bool)`

GetLastScaledAtOk returns a tuple with the LastScaledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScaledAt

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) SetLastScaledAt(v time.Time)`

SetLastScaledAt sets LastScaledAt field to given value.

### HasLastScaledAt

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) HasLastScaledAt() bool`

HasLastScaledAt returns a boolean if a field has been set.

### GetMessage

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPhase

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetReason

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReplicas

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetSelector

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1StepStatus) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


