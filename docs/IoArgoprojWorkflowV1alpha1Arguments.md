# IoArgoprojWorkflowV1alpha1Arguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | Pointer to [**[]IoArgoprojWorkflowV1alpha1Artifact**](IoArgoprojWorkflowV1alpha1Artifact.md) | Artifacts is the list of artifacts to pass to the template or workflow | [optional] 
**Parameters** | Pointer to [**[]IoArgoprojWorkflowV1alpha1Parameter**](IoArgoprojWorkflowV1alpha1Parameter.md) | Parameters is the list of parameters to pass to the template or workflow | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1Arguments

`func NewIoArgoprojWorkflowV1alpha1Arguments() *IoArgoprojWorkflowV1alpha1Arguments`

NewIoArgoprojWorkflowV1alpha1Arguments instantiates a new IoArgoprojWorkflowV1alpha1Arguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ArgumentsWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ArgumentsWithDefaults() *IoArgoprojWorkflowV1alpha1Arguments`

NewIoArgoprojWorkflowV1alpha1ArgumentsWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Arguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *IoArgoprojWorkflowV1alpha1Arguments) GetArtifacts() []IoArgoprojWorkflowV1alpha1Artifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *IoArgoprojWorkflowV1alpha1Arguments) GetArtifactsOk() (*[]IoArgoprojWorkflowV1alpha1Artifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *IoArgoprojWorkflowV1alpha1Arguments) SetArtifacts(v []IoArgoprojWorkflowV1alpha1Artifact)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *IoArgoprojWorkflowV1alpha1Arguments) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetParameters

`func (o *IoArgoprojWorkflowV1alpha1Arguments) GetParameters() []IoArgoprojWorkflowV1alpha1Parameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojWorkflowV1alpha1Arguments) GetParametersOk() (*[]IoArgoprojWorkflowV1alpha1Parameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojWorkflowV1alpha1Arguments) SetParameters(v []IoArgoprojWorkflowV1alpha1Parameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojWorkflowV1alpha1Arguments) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


