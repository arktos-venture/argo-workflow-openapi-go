# IoArgoprojWorkflowV1alpha1SubmitOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **string** | Annotations adds to metadata.labels | [optional] 
**DryRun** | Pointer to **bool** | DryRun validates the workflow on the client-side without creating it. This option is not supported in API | [optional] 
**EntryPoint** | Pointer to **string** | Entrypoint overrides spec.entrypoint | [optional] 
**GenerateName** | Pointer to **string** | GenerateName overrides metadata.generateName | [optional] 
**Labels** | Pointer to **string** | Labels adds to metadata.labels | [optional] 
**Name** | Pointer to **string** | Name overrides metadata.name | [optional] 
**OwnerReference** | Pointer to [**IoK8sApimachineryPkgApisMetaV1OwnerReference**](IoK8sApimachineryPkgApisMetaV1OwnerReference.md) |  | [optional] 
**ParameterFile** | Pointer to **string** | ParameterFile holds a reference to a parameter file. This option is not supported in API | [optional] 
**Parameters** | Pointer to **[]string** | Parameters passes input parameters to workflow | [optional] 
**ServerDryRun** | Pointer to **bool** | ServerDryRun validates the workflow on the server-side without creating it | [optional] 
**ServiceAccount** | Pointer to **string** | ServiceAccount runs all pods in the workflow using specified ServiceAccount. | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1SubmitOpts

`func NewIoArgoprojWorkflowV1alpha1SubmitOpts() *IoArgoprojWorkflowV1alpha1SubmitOpts`

NewIoArgoprojWorkflowV1alpha1SubmitOpts instantiates a new IoArgoprojWorkflowV1alpha1SubmitOpts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1SubmitOptsWithDefaults

`func NewIoArgoprojWorkflowV1alpha1SubmitOptsWithDefaults() *IoArgoprojWorkflowV1alpha1SubmitOpts`

NewIoArgoprojWorkflowV1alpha1SubmitOptsWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1SubmitOpts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetAnnotations() string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetAnnotationsOk() (*string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) SetAnnotations(v string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetDryRun

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetEntryPoint

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetEntryPoint() string`

GetEntryPoint returns the EntryPoint field if non-nil, zero value otherwise.

### GetEntryPointOk

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetEntryPointOk() (*string, bool)`

GetEntryPointOk returns a tuple with the EntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPoint

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) SetEntryPoint(v string)`

SetEntryPoint sets EntryPoint field to given value.

### HasEntryPoint

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) HasEntryPoint() bool`

HasEntryPoint returns a boolean if a field has been set.

### GetGenerateName

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetGenerateName() string`

GetGenerateName returns the GenerateName field if non-nil, zero value otherwise.

### GetGenerateNameOk

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetGenerateNameOk() (*string, bool)`

GetGenerateNameOk returns a tuple with the GenerateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateName

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) SetGenerateName(v string)`

SetGenerateName sets GenerateName field to given value.

### HasGenerateName

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) HasGenerateName() bool`

HasGenerateName returns a boolean if a field has been set.

### GetLabels

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetLabels() string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetLabelsOk() (*string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) SetLabels(v string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerReference

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetOwnerReference() IoK8sApimachineryPkgApisMetaV1OwnerReference`

GetOwnerReference returns the OwnerReference field if non-nil, zero value otherwise.

### GetOwnerReferenceOk

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetOwnerReferenceOk() (*IoK8sApimachineryPkgApisMetaV1OwnerReference, bool)`

GetOwnerReferenceOk returns a tuple with the OwnerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerReference

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) SetOwnerReference(v IoK8sApimachineryPkgApisMetaV1OwnerReference)`

SetOwnerReference sets OwnerReference field to given value.

### HasOwnerReference

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) HasOwnerReference() bool`

HasOwnerReference returns a boolean if a field has been set.

### GetParameterFile

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetParameterFile() string`

GetParameterFile returns the ParameterFile field if non-nil, zero value otherwise.

### GetParameterFileOk

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetParameterFileOk() (*string, bool)`

GetParameterFileOk returns a tuple with the ParameterFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterFile

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) SetParameterFile(v string)`

SetParameterFile sets ParameterFile field to given value.

### HasParameterFile

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) HasParameterFile() bool`

HasParameterFile returns a boolean if a field has been set.

### GetParameters

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetParameters() []string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetParametersOk() (*[]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) SetParameters(v []string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetServerDryRun

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetServerDryRun() bool`

GetServerDryRun returns the ServerDryRun field if non-nil, zero value otherwise.

### GetServerDryRunOk

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetServerDryRunOk() (*bool, bool)`

GetServerDryRunOk returns a tuple with the ServerDryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDryRun

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) SetServerDryRun(v bool)`

SetServerDryRun sets ServerDryRun field to given value.

### HasServerDryRun

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) HasServerDryRun() bool`

HasServerDryRun returns a boolean if a field has been set.

### GetServiceAccount

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *IoArgoprojWorkflowV1alpha1SubmitOpts) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


