# GithubComArgoprojLabsArgoDataflowApiV1alpha1Container

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** |  | [optional] 
**Command** | Pointer to **[]string** |  | [optional] 
**Env** | Pointer to [**[]IoK8sApiCoreV1EnvVar**](IoK8sApiCoreV1EnvVar.md) |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**In** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface.md) |  | [optional] 
**Resources** | Pointer to [**IoK8sApiCoreV1ResourceRequirements**](IoK8sApiCoreV1ResourceRequirements.md) |  | [optional] 
**VolumeMounts** | Pointer to [**[]IoK8sApiCoreV1VolumeMount**](IoK8sApiCoreV1VolumeMount.md) |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Container

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Container() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Container instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Container object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1ContainerWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1ContainerWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1ContainerWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Container object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetCommand

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEnv

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetEnv() []IoK8sApiCoreV1EnvVar`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetEnvOk() (*[]IoK8sApiCoreV1EnvVar, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) SetEnv(v []IoK8sApiCoreV1EnvVar)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetImage

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetIn

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetIn() GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetInOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) SetIn(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Interface)`

SetIn sets In field to given value.

### HasIn

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetResources

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetResources() IoK8sApiCoreV1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetResourcesOk() (*IoK8sApiCoreV1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) SetResources(v IoK8sApiCoreV1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetVolumeMounts

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetVolumeMounts() []IoK8sApiCoreV1VolumeMount`

GetVolumeMounts returns the VolumeMounts field if non-nil, zero value otherwise.

### GetVolumeMountsOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) GetVolumeMountsOk() (*[]IoK8sApiCoreV1VolumeMount, bool)`

GetVolumeMountsOk returns a tuple with the VolumeMounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMounts

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) SetVolumeMounts(v []IoK8sApiCoreV1VolumeMount)`

SetVolumeMounts sets VolumeMounts field to given value.

### HasVolumeMounts

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Container) HasVolumeMounts() bool`

HasVolumeMounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


