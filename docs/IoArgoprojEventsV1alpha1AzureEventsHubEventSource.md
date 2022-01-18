# IoArgoprojEventsV1alpha1AzureEventsHubEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** |  | [optional] 
**HubName** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**SharedAccessKey** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**SharedAccessKeyName** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1AzureEventsHubEventSource

`func NewIoArgoprojEventsV1alpha1AzureEventsHubEventSource() *IoArgoprojEventsV1alpha1AzureEventsHubEventSource`

NewIoArgoprojEventsV1alpha1AzureEventsHubEventSource instantiates a new IoArgoprojEventsV1alpha1AzureEventsHubEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1AzureEventsHubEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1AzureEventsHubEventSourceWithDefaults() *IoArgoprojEventsV1alpha1AzureEventsHubEventSource`

NewIoArgoprojEventsV1alpha1AzureEventsHubEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1AzureEventsHubEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetHubName

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) GetHubName() string`

GetHubName returns the HubName field if non-nil, zero value otherwise.

### GetHubNameOk

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) GetHubNameOk() (*string, bool)`

GetHubNameOk returns a tuple with the HubName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubName

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) SetHubName(v string)`

SetHubName sets HubName field to given value.

### HasHubName

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) HasHubName() bool`

HasHubName returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSharedAccessKey

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) GetSharedAccessKey() IoK8sApiCoreV1SecretKeySelector`

GetSharedAccessKey returns the SharedAccessKey field if non-nil, zero value otherwise.

### GetSharedAccessKeyOk

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) GetSharedAccessKeyOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSharedAccessKeyOk returns a tuple with the SharedAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAccessKey

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) SetSharedAccessKey(v IoK8sApiCoreV1SecretKeySelector)`

SetSharedAccessKey sets SharedAccessKey field to given value.

### HasSharedAccessKey

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) HasSharedAccessKey() bool`

HasSharedAccessKey returns a boolean if a field has been set.

### GetSharedAccessKeyName

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) GetSharedAccessKeyName() IoK8sApiCoreV1SecretKeySelector`

GetSharedAccessKeyName returns the SharedAccessKeyName field if non-nil, zero value otherwise.

### GetSharedAccessKeyNameOk

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) GetSharedAccessKeyNameOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSharedAccessKeyNameOk returns a tuple with the SharedAccessKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAccessKeyName

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) SetSharedAccessKeyName(v IoK8sApiCoreV1SecretKeySelector)`

SetSharedAccessKeyName sets SharedAccessKeyName field to given value.

### HasSharedAccessKeyName

`func (o *IoArgoprojEventsV1alpha1AzureEventsHubEventSource) HasSharedAccessKeyName() bool`

HasSharedAccessKeyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


