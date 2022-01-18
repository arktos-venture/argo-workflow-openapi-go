# IoArgoprojWorkflowV1alpha1WorkflowEventBindingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]IoArgoprojWorkflowV1alpha1WorkflowEventBinding**](IoArgoprojWorkflowV1alpha1WorkflowEventBinding.md) |  | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.io.k8s.community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | [**IoK8sApimachineryPkgApisMetaV1ListMeta**](IoK8sApimachineryPkgApisMetaV1ListMeta.md) |  | 

## Methods

### NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingList

`func NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingList(items []IoArgoprojWorkflowV1alpha1WorkflowEventBinding, metadata IoK8sApimachineryPkgApisMetaV1ListMeta, ) *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList`

NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingList instantiates a new IoArgoprojWorkflowV1alpha1WorkflowEventBindingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingListWithDefaults

`func NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingListWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList`

NewIoArgoprojWorkflowV1alpha1WorkflowEventBindingListWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowEventBindingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetItems() []IoArgoprojWorkflowV1alpha1WorkflowEventBinding`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetItemsOk() (*[]IoArgoprojWorkflowV1alpha1WorkflowEventBinding, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) SetItems(v []IoArgoprojWorkflowV1alpha1WorkflowEventBinding)`

SetItems sets Items field to given value.


### GetKind

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetMetadata() IoK8sApimachineryPkgApisMetaV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojWorkflowV1alpha1WorkflowEventBindingList) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ListMeta)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


