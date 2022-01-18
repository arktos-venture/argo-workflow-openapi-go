# IoArgoprojWorkflowV1alpha1CronWorkflowList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]IoArgoprojWorkflowV1alpha1CronWorkflow**](IoArgoprojWorkflowV1alpha1CronWorkflow.md) |  | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | [**IoK8sApimachineryPkgApisMetaV1ListMeta**](IoK8sApimachineryPkgApisMetaV1ListMeta.md) |  | 

## Methods

### NewIoArgoprojWorkflowV1alpha1CronWorkflowList

`func NewIoArgoprojWorkflowV1alpha1CronWorkflowList(items []IoArgoprojWorkflowV1alpha1CronWorkflow, metadata IoK8sApimachineryPkgApisMetaV1ListMeta, ) *IoArgoprojWorkflowV1alpha1CronWorkflowList`

NewIoArgoprojWorkflowV1alpha1CronWorkflowList instantiates a new IoArgoprojWorkflowV1alpha1CronWorkflowList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1CronWorkflowListWithDefaults

`func NewIoArgoprojWorkflowV1alpha1CronWorkflowListWithDefaults() *IoArgoprojWorkflowV1alpha1CronWorkflowList`

NewIoArgoprojWorkflowV1alpha1CronWorkflowListWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1CronWorkflowList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) GetItems() []IoArgoprojWorkflowV1alpha1CronWorkflow`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) GetItemsOk() (*[]IoArgoprojWorkflowV1alpha1CronWorkflow, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) SetItems(v []IoArgoprojWorkflowV1alpha1CronWorkflow)`

SetItems sets Items field to given value.


### GetKind

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) GetMetadata() IoK8sApimachineryPkgApisMetaV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojWorkflowV1alpha1CronWorkflowList) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ListMeta)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


