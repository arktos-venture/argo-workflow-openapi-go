# IoArgoprojWorkflowV1alpha1Outputs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | Pointer to [**[]IoArgoprojWorkflowV1alpha1Artifact**](IoArgoprojWorkflowV1alpha1Artifact.md) | Artifacts holds the list of output artifacts produced by a step | [optional] 
**ExitCode** | Pointer to **string** | ExitCode holds the exit code of a script template | [optional] 
**Parameters** | Pointer to [**[]IoArgoprojWorkflowV1alpha1Parameter**](IoArgoprojWorkflowV1alpha1Parameter.md) | Parameters holds the list of output parameters produced by a step | [optional] 
**Result** | Pointer to **string** | Result holds the result (stdout) of a script template | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1Outputs

`func NewIoArgoprojWorkflowV1alpha1Outputs() *IoArgoprojWorkflowV1alpha1Outputs`

NewIoArgoprojWorkflowV1alpha1Outputs instantiates a new IoArgoprojWorkflowV1alpha1Outputs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1OutputsWithDefaults

`func NewIoArgoprojWorkflowV1alpha1OutputsWithDefaults() *IoArgoprojWorkflowV1alpha1Outputs`

NewIoArgoprojWorkflowV1alpha1OutputsWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Outputs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *IoArgoprojWorkflowV1alpha1Outputs) GetArtifacts() []IoArgoprojWorkflowV1alpha1Artifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *IoArgoprojWorkflowV1alpha1Outputs) GetArtifactsOk() (*[]IoArgoprojWorkflowV1alpha1Artifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *IoArgoprojWorkflowV1alpha1Outputs) SetArtifacts(v []IoArgoprojWorkflowV1alpha1Artifact)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *IoArgoprojWorkflowV1alpha1Outputs) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetExitCode

`func (o *IoArgoprojWorkflowV1alpha1Outputs) GetExitCode() string`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *IoArgoprojWorkflowV1alpha1Outputs) GetExitCodeOk() (*string, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *IoArgoprojWorkflowV1alpha1Outputs) SetExitCode(v string)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *IoArgoprojWorkflowV1alpha1Outputs) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetParameters

`func (o *IoArgoprojWorkflowV1alpha1Outputs) GetParameters() []IoArgoprojWorkflowV1alpha1Parameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojWorkflowV1alpha1Outputs) GetParametersOk() (*[]IoArgoprojWorkflowV1alpha1Parameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojWorkflowV1alpha1Outputs) SetParameters(v []IoArgoprojWorkflowV1alpha1Parameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojWorkflowV1alpha1Outputs) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetResult

`func (o *IoArgoprojWorkflowV1alpha1Outputs) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *IoArgoprojWorkflowV1alpha1Outputs) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *IoArgoprojWorkflowV1alpha1Outputs) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *IoArgoprojWorkflowV1alpha1Outputs) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


