# IoArgoprojEventsV1alpha1StripeEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**CreateWebhook** | Pointer to **bool** |  | [optional] 
**EventFilter** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Webhook** | Pointer to [**IoArgoprojEventsV1alpha1WebhookContext**](IoArgoprojEventsV1alpha1WebhookContext.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1StripeEventSource

`func NewIoArgoprojEventsV1alpha1StripeEventSource() *IoArgoprojEventsV1alpha1StripeEventSource`

NewIoArgoprojEventsV1alpha1StripeEventSource instantiates a new IoArgoprojEventsV1alpha1StripeEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1StripeEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1StripeEventSourceWithDefaults() *IoArgoprojEventsV1alpha1StripeEventSource`

NewIoArgoprojEventsV1alpha1StripeEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1StripeEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) GetApiKey() IoK8sApiCoreV1SecretKeySelector`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) GetApiKeyOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) SetApiKey(v IoK8sApiCoreV1SecretKeySelector)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetCreateWebhook

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) GetCreateWebhook() bool`

GetCreateWebhook returns the CreateWebhook field if non-nil, zero value otherwise.

### GetCreateWebhookOk

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) GetCreateWebhookOk() (*bool, bool)`

GetCreateWebhookOk returns a tuple with the CreateWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateWebhook

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) SetCreateWebhook(v bool)`

SetCreateWebhook sets CreateWebhook field to given value.

### HasCreateWebhook

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) HasCreateWebhook() bool`

HasCreateWebhook returns a boolean if a field has been set.

### GetEventFilter

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) GetEventFilter() []string`

GetEventFilter returns the EventFilter field if non-nil, zero value otherwise.

### GetEventFilterOk

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) GetEventFilterOk() (*[]string, bool)`

GetEventFilterOk returns a tuple with the EventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilter

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) SetEventFilter(v []string)`

SetEventFilter sets EventFilter field to given value.

### HasEventFilter

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) HasEventFilter() bool`

HasEventFilter returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetWebhook

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) GetWebhook() IoArgoprojEventsV1alpha1WebhookContext`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) GetWebhookOk() (*IoArgoprojEventsV1alpha1WebhookContext, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) SetWebhook(v IoArgoprojEventsV1alpha1WebhookContext)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *IoArgoprojEventsV1alpha1StripeEventSource) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


