# IoArgoprojEventsV1alpha1KafkaTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compress** | Pointer to **bool** |  | [optional] 
**FlushFrequency** | Pointer to **int32** |  | [optional] 
**Parameters** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) | Parameters is the list of parameters that is applied to resolved Kafka trigger object. | [optional] 
**Partition** | Pointer to **int32** | Partition to write data to. | [optional] 
**PartitioningKey** | Pointer to **string** | The partitioning key for the messages put on the Kafka topic. Defaults to broker url. +optional. | [optional] 
**Payload** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) | Payload is the list of key-value extracted from an event payload to construct the request payload. | [optional] 
**RequiredAcks** | Pointer to **int32** | RequiredAcks used in producer to tell the broker how many replica acknowledgements Defaults to 1 (Only wait for the leader to ack). +optional. | [optional] 
**Sasl** | Pointer to [**IoArgoprojEventsV1alpha1SASLConfig**](IoArgoprojEventsV1alpha1SASLConfig.md) |  | [optional] 
**Tls** | Pointer to [**IoArgoprojEventsV1alpha1TLSConfig**](IoArgoprojEventsV1alpha1TLSConfig.md) |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL of the Kafka broker, multiple URLs separated by comma. | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1KafkaTrigger

`func NewIoArgoprojEventsV1alpha1KafkaTrigger() *IoArgoprojEventsV1alpha1KafkaTrigger`

NewIoArgoprojEventsV1alpha1KafkaTrigger instantiates a new IoArgoprojEventsV1alpha1KafkaTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1KafkaTriggerWithDefaults

`func NewIoArgoprojEventsV1alpha1KafkaTriggerWithDefaults() *IoArgoprojEventsV1alpha1KafkaTrigger`

NewIoArgoprojEventsV1alpha1KafkaTriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1KafkaTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompress

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetFlushFrequency

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetFlushFrequency() int32`

GetFlushFrequency returns the FlushFrequency field if non-nil, zero value otherwise.

### GetFlushFrequencyOk

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetFlushFrequencyOk() (*int32, bool)`

GetFlushFrequencyOk returns a tuple with the FlushFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlushFrequency

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) SetFlushFrequency(v int32)`

SetFlushFrequency sets FlushFrequency field to given value.

### HasFlushFrequency

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) HasFlushFrequency() bool`

HasFlushFrequency returns a boolean if a field has been set.

### GetParameters

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPartition

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetPartitioningKey

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetPartitioningKey() string`

GetPartitioningKey returns the PartitioningKey field if non-nil, zero value otherwise.

### GetPartitioningKeyOk

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetPartitioningKeyOk() (*string, bool)`

GetPartitioningKeyOk returns a tuple with the PartitioningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitioningKey

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) SetPartitioningKey(v string)`

SetPartitioningKey sets PartitioningKey field to given value.

### HasPartitioningKey

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) HasPartitioningKey() bool`

HasPartitioningKey returns a boolean if a field has been set.

### GetPayload

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetPayload() []IoArgoprojEventsV1alpha1TriggerParameter`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetPayloadOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) SetPayload(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetRequiredAcks

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetRequiredAcks() int32`

GetRequiredAcks returns the RequiredAcks field if non-nil, zero value otherwise.

### GetRequiredAcksOk

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetRequiredAcksOk() (*int32, bool)`

GetRequiredAcksOk returns a tuple with the RequiredAcks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAcks

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) SetRequiredAcks(v int32)`

SetRequiredAcks sets RequiredAcks field to given value.

### HasRequiredAcks

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) HasRequiredAcks() bool`

HasRequiredAcks returns a boolean if a field has been set.

### GetSasl

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetSasl() IoArgoprojEventsV1alpha1SASLConfig`

GetSasl returns the Sasl field if non-nil, zero value otherwise.

### GetSaslOk

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetSaslOk() (*IoArgoprojEventsV1alpha1SASLConfig, bool)`

GetSaslOk returns a tuple with the Sasl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasl

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) SetSasl(v IoArgoprojEventsV1alpha1SASLConfig)`

SetSasl sets Sasl field to given value.

### HasSasl

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) HasSasl() bool`

HasSasl returns a boolean if a field has been set.

### GetTls

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetTls() IoArgoprojEventsV1alpha1TLSConfig`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetTlsOk() (*IoArgoprojEventsV1alpha1TLSConfig, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) SetTls(v IoArgoprojEventsV1alpha1TLSConfig)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetTopic

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVersion

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IoArgoprojEventsV1alpha1KafkaTrigger) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


