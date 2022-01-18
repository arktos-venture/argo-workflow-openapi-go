# GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]IoK8sApimachineryPkgApisMetaV1Condition**](IoK8sApimachineryPkgApisMetaV1Condition.md) |  | [optional] 
**LastUpdated** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus() *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatusWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatusWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatusWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetConditions() []IoK8sApimachineryPkgApisMetaV1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetConditionsOk() (*[]IoK8sApimachineryPkgApisMetaV1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) SetConditions(v []IoK8sApimachineryPkgApisMetaV1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMessage

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPhase

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


