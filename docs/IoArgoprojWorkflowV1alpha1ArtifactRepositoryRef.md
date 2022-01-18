# IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMap** | Pointer to **string** | The name of the config map. Defaults to \&quot;artifact-repositories\&quot;. | [optional] 
**Key** | Pointer to **string** | The config map key. Defaults to the value of the \&quot;workflows.argoproj.io/default-artifact-repository\&quot; annotation. | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryRef

`func NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryRef() *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef`

NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryRef instantiates a new IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryRefWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryRefWithDefaults() *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef`

NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryRefWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMap

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef) GetConfigMap() string`

GetConfigMap returns the ConfigMap field if non-nil, zero value otherwise.

### GetConfigMapOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef) GetConfigMapOk() (*string, bool)`

GetConfigMapOk returns a tuple with the ConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMap

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef) SetConfigMap(v string)`

SetConfigMap sets ConfigMap field to given value.

### HasConfigMap

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef) HasConfigMap() bool`

HasConfigMap returns a boolean if a field has been set.

### GetKey

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRef) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


