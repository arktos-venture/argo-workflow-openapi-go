# IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactRepository** | Pointer to [**IoArgoprojWorkflowV1alpha1ArtifactRepository**](IoArgoprojWorkflowV1alpha1ArtifactRepository.md) |  | [optional] 
**ConfigMap** | Pointer to **string** | The name of the config map. Defaults to \&quot;artifact-repositories\&quot;. | [optional] 
**Default** | Pointer to **bool** | If this ref represents the default artifact repository, rather than a config map. | [optional] 
**Key** | Pointer to **string** | The config map key. Defaults to the value of the \&quot;workflows.argoproj.io/default-artifact-repository\&quot; annotation. | [optional] 
**Namespace** | Pointer to **string** | The namespace of the config map. Defaults to the workflow&#39;s namespace, or the controller&#39;s namespace (if found). | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus

`func NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus() *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus`

NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus instantiates a new IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatusWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatusWithDefaults() *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus`

NewIoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatusWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactRepository

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) GetArtifactRepository() IoArgoprojWorkflowV1alpha1ArtifactRepository`

GetArtifactRepository returns the ArtifactRepository field if non-nil, zero value otherwise.

### GetArtifactRepositoryOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) GetArtifactRepositoryOk() (*IoArgoprojWorkflowV1alpha1ArtifactRepository, bool)`

GetArtifactRepositoryOk returns a tuple with the ArtifactRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactRepository

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) SetArtifactRepository(v IoArgoprojWorkflowV1alpha1ArtifactRepository)`

SetArtifactRepository sets ArtifactRepository field to given value.

### HasArtifactRepository

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) HasArtifactRepository() bool`

HasArtifactRepository returns a boolean if a field has been set.

### GetConfigMap

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) GetConfigMap() string`

GetConfigMap returns the ConfigMap field if non-nil, zero value otherwise.

### GetConfigMapOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) GetConfigMapOk() (*string, bool)`

GetConfigMapOk returns a tuple with the ConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMap

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) SetConfigMap(v string)`

SetConfigMap sets ConfigMap field to given value.

### HasConfigMap

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) HasConfigMap() bool`

HasConfigMap returns a boolean if a field has been set.

### GetDefault

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetKey

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNamespace

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoArgoprojWorkflowV1alpha1ArtifactRepositoryRefStatus) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


