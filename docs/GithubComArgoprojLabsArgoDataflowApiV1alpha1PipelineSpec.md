# GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletionDelay** | Pointer to [**IoK8sApimachineryPkgApisMetaV1Duration**](IoK8sApimachineryPkgApisMetaV1Duration.md) |  | [optional] 
**Steps** | Pointer to [**[]GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec**](GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec.md) |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec() *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpecWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpecWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpecWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletionDelay

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec) GetDeletionDelay() IoK8sApimachineryPkgApisMetaV1Duration`

GetDeletionDelay returns the DeletionDelay field if non-nil, zero value otherwise.

### GetDeletionDelayOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec) GetDeletionDelayOk() (*IoK8sApimachineryPkgApisMetaV1Duration, bool)`

GetDeletionDelayOk returns a tuple with the DeletionDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionDelay

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec) SetDeletionDelay(v IoK8sApimachineryPkgApisMetaV1Duration)`

SetDeletionDelay sets DeletionDelay field to given value.

### HasDeletionDelay

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec) HasDeletionDelay() bool`

HasDeletionDelay returns a boolean if a field has been set.

### GetSteps

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec) GetSteps() []GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec) GetStepsOk() (*[]GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec) SetSteps(v []GithubComArgoprojLabsArgoDataflowApiV1alpha1StepSpec)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1PipelineSpec) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


