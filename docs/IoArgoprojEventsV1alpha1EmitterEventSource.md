# IoArgoprojEventsV1alpha1EmitterEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Broker** | Pointer to **string** | Broker URI to connect to. | [optional] 
**ChannelKey** | Pointer to **string** |  | [optional] 
**ChannelName** | Pointer to **string** |  | [optional] 
**ConnectionBackoff** | Pointer to [**IoArgoprojEventsV1alpha1Backoff**](IoArgoprojEventsV1alpha1Backoff.md) |  | [optional] 
**JsonBody** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Password** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Tls** | Pointer to [**IoArgoprojEventsV1alpha1TLSConfig**](IoArgoprojEventsV1alpha1TLSConfig.md) |  | [optional] 
**Username** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1EmitterEventSource

`func NewIoArgoprojEventsV1alpha1EmitterEventSource() *IoArgoprojEventsV1alpha1EmitterEventSource`

NewIoArgoprojEventsV1alpha1EmitterEventSource instantiates a new IoArgoprojEventsV1alpha1EmitterEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1EmitterEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1EmitterEventSourceWithDefaults() *IoArgoprojEventsV1alpha1EmitterEventSource`

NewIoArgoprojEventsV1alpha1EmitterEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1EmitterEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBroker

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetBroker() string`

GetBroker returns the Broker field if non-nil, zero value otherwise.

### GetBrokerOk

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetBrokerOk() (*string, bool)`

GetBrokerOk returns a tuple with the Broker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroker

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) SetBroker(v string)`

SetBroker sets Broker field to given value.

### HasBroker

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) HasBroker() bool`

HasBroker returns a boolean if a field has been set.

### GetChannelKey

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetChannelKey() string`

GetChannelKey returns the ChannelKey field if non-nil, zero value otherwise.

### GetChannelKeyOk

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetChannelKeyOk() (*string, bool)`

GetChannelKeyOk returns a tuple with the ChannelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelKey

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) SetChannelKey(v string)`

SetChannelKey sets ChannelKey field to given value.

### HasChannelKey

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) HasChannelKey() bool`

HasChannelKey returns a boolean if a field has been set.

### GetChannelName

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetConnectionBackoff() IoArgoprojEventsV1alpha1Backoff`

GetConnectionBackoff returns the ConnectionBackoff field if non-nil, zero value otherwise.

### GetConnectionBackoffOk

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetConnectionBackoffOk() (*IoArgoprojEventsV1alpha1Backoff, bool)`

GetConnectionBackoffOk returns a tuple with the ConnectionBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) SetConnectionBackoff(v IoArgoprojEventsV1alpha1Backoff)`

SetConnectionBackoff sets ConnectionBackoff field to given value.

### HasConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) HasConnectionBackoff() bool`

HasConnectionBackoff returns a boolean if a field has been set.

### GetJsonBody

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetJsonBody() bool`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetJsonBodyOk() (*bool, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) SetJsonBody(v bool)`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPassword

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetPassword() IoK8sApiCoreV1SecretKeySelector`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetPasswordOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) SetPassword(v IoK8sApiCoreV1SecretKeySelector)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTls

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetTls() IoArgoprojEventsV1alpha1TLSConfig`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetTlsOk() (*IoArgoprojEventsV1alpha1TLSConfig, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) SetTls(v IoArgoprojEventsV1alpha1TLSConfig)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetUsername

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetUsername() IoK8sApiCoreV1SecretKeySelector`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) GetUsernameOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) SetUsername(v IoK8sApiCoreV1SecretKeySelector)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IoArgoprojEventsV1alpha1EmitterEventSource) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


