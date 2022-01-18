# IoArgoprojEventsV1alpha1ResourceEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTypes** | Pointer to **[]string** | EventTypes is the list of event type to watch. Possible values are - ADD, UPDATE and DELETE. | [optional] 
**Filter** | Pointer to [**IoArgoprojEventsV1alpha1ResourceFilter**](IoArgoprojEventsV1alpha1ResourceFilter.md) |  | [optional] 
**GroupVersionResource** | Pointer to [**IoK8sApimachineryPkgApisMetaV1GroupVersionResource**](IoK8sApimachineryPkgApisMetaV1GroupVersionResource.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1ResourceEventSource

`func NewIoArgoprojEventsV1alpha1ResourceEventSource() *IoArgoprojEventsV1alpha1ResourceEventSource`

NewIoArgoprojEventsV1alpha1ResourceEventSource instantiates a new IoArgoprojEventsV1alpha1ResourceEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1ResourceEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1ResourceEventSourceWithDefaults() *IoArgoprojEventsV1alpha1ResourceEventSource`

NewIoArgoprojEventsV1alpha1ResourceEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1ResourceEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTypes

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetFilter

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) GetFilter() IoArgoprojEventsV1alpha1ResourceFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) GetFilterOk() (*IoArgoprojEventsV1alpha1ResourceFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) SetFilter(v IoArgoprojEventsV1alpha1ResourceFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupVersionResource

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) GetGroupVersionResource() IoK8sApimachineryPkgApisMetaV1GroupVersionResource`

GetGroupVersionResource returns the GroupVersionResource field if non-nil, zero value otherwise.

### GetGroupVersionResourceOk

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) GetGroupVersionResourceOk() (*IoK8sApimachineryPkgApisMetaV1GroupVersionResource, bool)`

GetGroupVersionResourceOk returns a tuple with the GroupVersionResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupVersionResource

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) SetGroupVersionResource(v IoK8sApimachineryPkgApisMetaV1GroupVersionResource)`

SetGroupVersionResource sets GroupVersionResource field to given value.

### HasGroupVersionResource

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) HasGroupVersionResource() bool`

HasGroupVersionResource returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNamespace

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoArgoprojEventsV1alpha1ResourceEventSource) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


