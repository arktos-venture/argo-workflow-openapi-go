# IoArgoprojEventsV1alpha1AMQPEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**IoArgoprojEventsV1alpha1BasicAuth**](IoArgoprojEventsV1alpha1BasicAuth.md) |  | [optional] 
**ConnectionBackoff** | Pointer to [**IoArgoprojEventsV1alpha1Backoff**](IoArgoprojEventsV1alpha1Backoff.md) |  | [optional] 
**Consume** | Pointer to [**IoArgoprojEventsV1alpha1AMQPConsumeConfig**](IoArgoprojEventsV1alpha1AMQPConsumeConfig.md) |  | [optional] 
**ExchangeDeclare** | Pointer to [**IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig**](IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig.md) |  | [optional] 
**ExchangeName** | Pointer to **string** |  | [optional] 
**ExchangeType** | Pointer to **string** |  | [optional] 
**JsonBody** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**QueueBind** | Pointer to [**IoArgoprojEventsV1alpha1AMQPQueueBindConfig**](IoArgoprojEventsV1alpha1AMQPQueueBindConfig.md) |  | [optional] 
**QueueDeclare** | Pointer to [**IoArgoprojEventsV1alpha1AMQPQueueDeclareConfig**](IoArgoprojEventsV1alpha1AMQPQueueDeclareConfig.md) |  | [optional] 
**RoutingKey** | Pointer to **string** |  | [optional] 
**Tls** | Pointer to [**IoArgoprojEventsV1alpha1TLSConfig**](IoArgoprojEventsV1alpha1TLSConfig.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1AMQPEventSource

`func NewIoArgoprojEventsV1alpha1AMQPEventSource() *IoArgoprojEventsV1alpha1AMQPEventSource`

NewIoArgoprojEventsV1alpha1AMQPEventSource instantiates a new IoArgoprojEventsV1alpha1AMQPEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1AMQPEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1AMQPEventSourceWithDefaults() *IoArgoprojEventsV1alpha1AMQPEventSource`

NewIoArgoprojEventsV1alpha1AMQPEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1AMQPEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetAuth() IoArgoprojEventsV1alpha1BasicAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetAuthOk() (*IoArgoprojEventsV1alpha1BasicAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetAuth(v IoArgoprojEventsV1alpha1BasicAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetConnectionBackoff() IoArgoprojEventsV1alpha1Backoff`

GetConnectionBackoff returns the ConnectionBackoff field if non-nil, zero value otherwise.

### GetConnectionBackoffOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetConnectionBackoffOk() (*IoArgoprojEventsV1alpha1Backoff, bool)`

GetConnectionBackoffOk returns a tuple with the ConnectionBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetConnectionBackoff(v IoArgoprojEventsV1alpha1Backoff)`

SetConnectionBackoff sets ConnectionBackoff field to given value.

### HasConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasConnectionBackoff() bool`

HasConnectionBackoff returns a boolean if a field has been set.

### GetConsume

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetConsume() IoArgoprojEventsV1alpha1AMQPConsumeConfig`

GetConsume returns the Consume field if non-nil, zero value otherwise.

### GetConsumeOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetConsumeOk() (*IoArgoprojEventsV1alpha1AMQPConsumeConfig, bool)`

GetConsumeOk returns a tuple with the Consume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsume

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetConsume(v IoArgoprojEventsV1alpha1AMQPConsumeConfig)`

SetConsume sets Consume field to given value.

### HasConsume

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasConsume() bool`

HasConsume returns a boolean if a field has been set.

### GetExchangeDeclare

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetExchangeDeclare() IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig`

GetExchangeDeclare returns the ExchangeDeclare field if non-nil, zero value otherwise.

### GetExchangeDeclareOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetExchangeDeclareOk() (*IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig, bool)`

GetExchangeDeclareOk returns a tuple with the ExchangeDeclare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeDeclare

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetExchangeDeclare(v IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig)`

SetExchangeDeclare sets ExchangeDeclare field to given value.

### HasExchangeDeclare

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasExchangeDeclare() bool`

HasExchangeDeclare returns a boolean if a field has been set.

### GetExchangeName

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetExchangeType

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetExchangeType() string`

GetExchangeType returns the ExchangeType field if non-nil, zero value otherwise.

### GetExchangeTypeOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetExchangeTypeOk() (*string, bool)`

GetExchangeTypeOk returns a tuple with the ExchangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeType

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetExchangeType(v string)`

SetExchangeType sets ExchangeType field to given value.

### HasExchangeType

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasExchangeType() bool`

HasExchangeType returns a boolean if a field has been set.

### GetJsonBody

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetJsonBody() bool`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetJsonBodyOk() (*bool, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetJsonBody(v bool)`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQueueBind

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetQueueBind() IoArgoprojEventsV1alpha1AMQPQueueBindConfig`

GetQueueBind returns the QueueBind field if non-nil, zero value otherwise.

### GetQueueBindOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetQueueBindOk() (*IoArgoprojEventsV1alpha1AMQPQueueBindConfig, bool)`

GetQueueBindOk returns a tuple with the QueueBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueBind

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetQueueBind(v IoArgoprojEventsV1alpha1AMQPQueueBindConfig)`

SetQueueBind sets QueueBind field to given value.

### HasQueueBind

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasQueueBind() bool`

HasQueueBind returns a boolean if a field has been set.

### GetQueueDeclare

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetQueueDeclare() IoArgoprojEventsV1alpha1AMQPQueueDeclareConfig`

GetQueueDeclare returns the QueueDeclare field if non-nil, zero value otherwise.

### GetQueueDeclareOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetQueueDeclareOk() (*IoArgoprojEventsV1alpha1AMQPQueueDeclareConfig, bool)`

GetQueueDeclareOk returns a tuple with the QueueDeclare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDeclare

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetQueueDeclare(v IoArgoprojEventsV1alpha1AMQPQueueDeclareConfig)`

SetQueueDeclare sets QueueDeclare field to given value.

### HasQueueDeclare

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasQueueDeclare() bool`

HasQueueDeclare returns a boolean if a field has been set.

### GetRoutingKey

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetRoutingKey() string`

GetRoutingKey returns the RoutingKey field if non-nil, zero value otherwise.

### GetRoutingKeyOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetRoutingKeyOk() (*string, bool)`

GetRoutingKeyOk returns a tuple with the RoutingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKey

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetRoutingKey(v string)`

SetRoutingKey sets RoutingKey field to given value.

### HasRoutingKey

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasRoutingKey() bool`

HasRoutingKey returns a boolean if a field has been set.

### GetTls

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetTls() IoArgoprojEventsV1alpha1TLSConfig`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetTlsOk() (*IoArgoprojEventsV1alpha1TLSConfig, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetTls(v IoArgoprojEventsV1alpha1TLSConfig)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoArgoprojEventsV1alpha1AMQPEventSource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


