# IoArgoprojEventsV1alpha1MQTTEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**ConnectionBackoff** | Pointer to [**IoArgoprojEventsV1alpha1Backoff**](IoArgoprojEventsV1alpha1Backoff.md) |  | [optional] 
**JsonBody** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Tls** | Pointer to [**IoArgoprojEventsV1alpha1TLSConfig**](IoArgoprojEventsV1alpha1TLSConfig.md) |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1MQTTEventSource

`func NewIoArgoprojEventsV1alpha1MQTTEventSource() *IoArgoprojEventsV1alpha1MQTTEventSource`

NewIoArgoprojEventsV1alpha1MQTTEventSource instantiates a new IoArgoprojEventsV1alpha1MQTTEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1MQTTEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1MQTTEventSourceWithDefaults() *IoArgoprojEventsV1alpha1MQTTEventSource`

NewIoArgoprojEventsV1alpha1MQTTEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1MQTTEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetConnectionBackoff() IoArgoprojEventsV1alpha1Backoff`

GetConnectionBackoff returns the ConnectionBackoff field if non-nil, zero value otherwise.

### GetConnectionBackoffOk

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetConnectionBackoffOk() (*IoArgoprojEventsV1alpha1Backoff, bool)`

GetConnectionBackoffOk returns a tuple with the ConnectionBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) SetConnectionBackoff(v IoArgoprojEventsV1alpha1Backoff)`

SetConnectionBackoff sets ConnectionBackoff field to given value.

### HasConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) HasConnectionBackoff() bool`

HasConnectionBackoff returns a boolean if a field has been set.

### GetJsonBody

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetJsonBody() bool`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetJsonBodyOk() (*bool, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) SetJsonBody(v bool)`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTls

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetTls() IoArgoprojEventsV1alpha1TLSConfig`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetTlsOk() (*IoArgoprojEventsV1alpha1TLSConfig, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) SetTls(v IoArgoprojEventsV1alpha1TLSConfig)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetTopic

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoArgoprojEventsV1alpha1MQTTEventSource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


