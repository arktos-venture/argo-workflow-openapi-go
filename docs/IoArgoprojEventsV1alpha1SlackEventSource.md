# IoArgoprojEventsV1alpha1SlackEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**SigningSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Token** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Webhook** | Pointer to [**IoArgoprojEventsV1alpha1WebhookContext**](IoArgoprojEventsV1alpha1WebhookContext.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1SlackEventSource

`func NewIoArgoprojEventsV1alpha1SlackEventSource() *IoArgoprojEventsV1alpha1SlackEventSource`

NewIoArgoprojEventsV1alpha1SlackEventSource instantiates a new IoArgoprojEventsV1alpha1SlackEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1SlackEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1SlackEventSourceWithDefaults() *IoArgoprojEventsV1alpha1SlackEventSource`

NewIoArgoprojEventsV1alpha1SlackEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1SlackEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSigningSecret

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) GetSigningSecret() IoK8sApiCoreV1SecretKeySelector`

GetSigningSecret returns the SigningSecret field if non-nil, zero value otherwise.

### GetSigningSecretOk

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) GetSigningSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSigningSecretOk returns a tuple with the SigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecret

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) SetSigningSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetSigningSecret sets SigningSecret field to given value.

### HasSigningSecret

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) HasSigningSecret() bool`

HasSigningSecret returns a boolean if a field has been set.

### GetToken

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) GetToken() IoK8sApiCoreV1SecretKeySelector`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) GetTokenOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) SetToken(v IoK8sApiCoreV1SecretKeySelector)`

SetToken sets Token field to given value.

### HasToken

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetWebhook

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) GetWebhook() IoArgoprojEventsV1alpha1WebhookContext`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) GetWebhookOk() (*IoArgoprojEventsV1alpha1WebhookContext, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) SetWebhook(v IoArgoprojEventsV1alpha1WebhookContext)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *IoArgoprojEventsV1alpha1SlackEventSource) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


