# IoArgoprojEventsV1alpha1KafkaEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionBackoff** | Pointer to [**IoArgoprojEventsV1alpha1Backoff**](IoArgoprojEventsV1alpha1Backoff.md) |  | [optional] 
**ConsumerGroup** | Pointer to [**IoArgoprojEventsV1alpha1KafkaConsumerGroup**](IoArgoprojEventsV1alpha1KafkaConsumerGroup.md) |  | [optional] 
**JsonBody** | Pointer to **bool** |  | [optional] 
**LimitEventsPerSecond** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Partition** | Pointer to **string** |  | [optional] 
**Sasl** | Pointer to [**IoArgoprojEventsV1alpha1SASLConfig**](IoArgoprojEventsV1alpha1SASLConfig.md) |  | [optional] 
**Tls** | Pointer to [**IoArgoprojEventsV1alpha1TLSConfig**](IoArgoprojEventsV1alpha1TLSConfig.md) |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1KafkaEventSource

`func NewIoArgoprojEventsV1alpha1KafkaEventSource() *IoArgoprojEventsV1alpha1KafkaEventSource`

NewIoArgoprojEventsV1alpha1KafkaEventSource instantiates a new IoArgoprojEventsV1alpha1KafkaEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1KafkaEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1KafkaEventSourceWithDefaults() *IoArgoprojEventsV1alpha1KafkaEventSource`

NewIoArgoprojEventsV1alpha1KafkaEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1KafkaEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetConnectionBackoff() IoArgoprojEventsV1alpha1Backoff`

GetConnectionBackoff returns the ConnectionBackoff field if non-nil, zero value otherwise.

### GetConnectionBackoffOk

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetConnectionBackoffOk() (*IoArgoprojEventsV1alpha1Backoff, bool)`

GetConnectionBackoffOk returns a tuple with the ConnectionBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) SetConnectionBackoff(v IoArgoprojEventsV1alpha1Backoff)`

SetConnectionBackoff sets ConnectionBackoff field to given value.

### HasConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) HasConnectionBackoff() bool`

HasConnectionBackoff returns a boolean if a field has been set.

### GetConsumerGroup

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetConsumerGroup() IoArgoprojEventsV1alpha1KafkaConsumerGroup`

GetConsumerGroup returns the ConsumerGroup field if non-nil, zero value otherwise.

### GetConsumerGroupOk

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetConsumerGroupOk() (*IoArgoprojEventsV1alpha1KafkaConsumerGroup, bool)`

GetConsumerGroupOk returns a tuple with the ConsumerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerGroup

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) SetConsumerGroup(v IoArgoprojEventsV1alpha1KafkaConsumerGroup)`

SetConsumerGroup sets ConsumerGroup field to given value.

### HasConsumerGroup

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) HasConsumerGroup() bool`

HasConsumerGroup returns a boolean if a field has been set.

### GetJsonBody

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetJsonBody() bool`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetJsonBodyOk() (*bool, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) SetJsonBody(v bool)`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.

### GetLimitEventsPerSecond

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetLimitEventsPerSecond() string`

GetLimitEventsPerSecond returns the LimitEventsPerSecond field if non-nil, zero value otherwise.

### GetLimitEventsPerSecondOk

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetLimitEventsPerSecondOk() (*string, bool)`

GetLimitEventsPerSecondOk returns a tuple with the LimitEventsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitEventsPerSecond

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) SetLimitEventsPerSecond(v string)`

SetLimitEventsPerSecond sets LimitEventsPerSecond field to given value.

### HasLimitEventsPerSecond

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) HasLimitEventsPerSecond() bool`

HasLimitEventsPerSecond returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPartition

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetSasl

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetSasl() IoArgoprojEventsV1alpha1SASLConfig`

GetSasl returns the Sasl field if non-nil, zero value otherwise.

### GetSaslOk

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetSaslOk() (*IoArgoprojEventsV1alpha1SASLConfig, bool)`

GetSaslOk returns a tuple with the Sasl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasl

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) SetSasl(v IoArgoprojEventsV1alpha1SASLConfig)`

SetSasl sets Sasl field to given value.

### HasSasl

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) HasSasl() bool`

HasSasl returns a boolean if a field has been set.

### GetTls

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetTls() IoArgoprojEventsV1alpha1TLSConfig`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetTlsOk() (*IoArgoprojEventsV1alpha1TLSConfig, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) SetTls(v IoArgoprojEventsV1alpha1TLSConfig)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetTopic

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVersion

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IoArgoprojEventsV1alpha1KafkaEventSource) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


