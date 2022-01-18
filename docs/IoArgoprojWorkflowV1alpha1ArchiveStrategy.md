# IoArgoprojWorkflowV1alpha1ArchiveStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**None** | Pointer to **map[string]interface{}** | NoneStrategy indicates to skip tar process and upload the files or directory tree as independent files. Note that if the artifact is a directory, the artifact driver must support the ability to save/load the directory appropriately. | [optional] 
**Tar** | Pointer to [**IoArgoprojWorkflowV1alpha1TarStrategy**](IoArgoprojWorkflowV1alpha1TarStrategy.md) |  | [optional] 
**Zip** | Pointer to **map[string]interface{}** | ZipStrategy will unzip zipped input artifacts | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1ArchiveStrategy

`func NewIoArgoprojWorkflowV1alpha1ArchiveStrategy() *IoArgoprojWorkflowV1alpha1ArchiveStrategy`

NewIoArgoprojWorkflowV1alpha1ArchiveStrategy instantiates a new IoArgoprojWorkflowV1alpha1ArchiveStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ArchiveStrategyWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ArchiveStrategyWithDefaults() *IoArgoprojWorkflowV1alpha1ArchiveStrategy`

NewIoArgoprojWorkflowV1alpha1ArchiveStrategyWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ArchiveStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNone

`func (o *IoArgoprojWorkflowV1alpha1ArchiveStrategy) GetNone() map[string]interface{}`

GetNone returns the None field if non-nil, zero value otherwise.

### GetNoneOk

`func (o *IoArgoprojWorkflowV1alpha1ArchiveStrategy) GetNoneOk() (*map[string]interface{}, bool)`

GetNoneOk returns a tuple with the None field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNone

`func (o *IoArgoprojWorkflowV1alpha1ArchiveStrategy) SetNone(v map[string]interface{})`

SetNone sets None field to given value.

### HasNone

`func (o *IoArgoprojWorkflowV1alpha1ArchiveStrategy) HasNone() bool`

HasNone returns a boolean if a field has been set.

### GetTar

`func (o *IoArgoprojWorkflowV1alpha1ArchiveStrategy) GetTar() IoArgoprojWorkflowV1alpha1TarStrategy`

GetTar returns the Tar field if non-nil, zero value otherwise.

### GetTarOk

`func (o *IoArgoprojWorkflowV1alpha1ArchiveStrategy) GetTarOk() (*IoArgoprojWorkflowV1alpha1TarStrategy, bool)`

GetTarOk returns a tuple with the Tar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTar

`func (o *IoArgoprojWorkflowV1alpha1ArchiveStrategy) SetTar(v IoArgoprojWorkflowV1alpha1TarStrategy)`

SetTar sets Tar field to given value.

### HasTar

`func (o *IoArgoprojWorkflowV1alpha1ArchiveStrategy) HasTar() bool`

HasTar returns a boolean if a field has been set.

### GetZip

`func (o *IoArgoprojWorkflowV1alpha1ArchiveStrategy) GetZip() map[string]interface{}`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *IoArgoprojWorkflowV1alpha1ArchiveStrategy) GetZipOk() (*map[string]interface{}, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *IoArgoprojWorkflowV1alpha1ArchiveStrategy) SetZip(v map[string]interface{})`

SetZip sets Zip field to given value.

### HasZip

`func (o *IoArgoprojWorkflowV1alpha1ArchiveStrategy) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


