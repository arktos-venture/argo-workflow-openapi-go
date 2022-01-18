# IoArgoprojEventsV1alpha1EventSourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amqp** | Pointer to [**map[string]IoArgoprojEventsV1alpha1AMQPEventSource**](IoArgoprojEventsV1alpha1AMQPEventSource.md) |  | [optional] 
**AzureEventsHub** | Pointer to [**map[string]IoArgoprojEventsV1alpha1AzureEventsHubEventSource**](IoArgoprojEventsV1alpha1AzureEventsHubEventSource.md) |  | [optional] 
**Calendar** | Pointer to [**map[string]IoArgoprojEventsV1alpha1CalendarEventSource**](IoArgoprojEventsV1alpha1CalendarEventSource.md) |  | [optional] 
**Emitter** | Pointer to [**map[string]IoArgoprojEventsV1alpha1EmitterEventSource**](IoArgoprojEventsV1alpha1EmitterEventSource.md) |  | [optional] 
**EventBusName** | Pointer to **string** |  | [optional] 
**File** | Pointer to [**map[string]IoArgoprojEventsV1alpha1FileEventSource**](IoArgoprojEventsV1alpha1FileEventSource.md) |  | [optional] 
**Generic** | Pointer to [**map[string]IoArgoprojEventsV1alpha1GenericEventSource**](IoArgoprojEventsV1alpha1GenericEventSource.md) |  | [optional] 
**Github** | Pointer to [**map[string]IoArgoprojEventsV1alpha1GithubEventSource**](IoArgoprojEventsV1alpha1GithubEventSource.md) |  | [optional] 
**Gitlab** | Pointer to [**map[string]IoArgoprojEventsV1alpha1GitlabEventSource**](IoArgoprojEventsV1alpha1GitlabEventSource.md) |  | [optional] 
**Hdfs** | Pointer to [**map[string]IoArgoprojEventsV1alpha1HDFSEventSource**](IoArgoprojEventsV1alpha1HDFSEventSource.md) |  | [optional] 
**Kafka** | Pointer to [**map[string]IoArgoprojEventsV1alpha1KafkaEventSource**](IoArgoprojEventsV1alpha1KafkaEventSource.md) |  | [optional] 
**Minio** | Pointer to [**map[string]IoArgoprojEventsV1alpha1S3Artifact**](IoArgoprojEventsV1alpha1S3Artifact.md) |  | [optional] 
**Mqtt** | Pointer to [**map[string]IoArgoprojEventsV1alpha1MQTTEventSource**](IoArgoprojEventsV1alpha1MQTTEventSource.md) |  | [optional] 
**Nats** | Pointer to [**map[string]IoArgoprojEventsV1alpha1NATSEventsSource**](IoArgoprojEventsV1alpha1NATSEventsSource.md) |  | [optional] 
**Nsq** | Pointer to [**map[string]IoArgoprojEventsV1alpha1NSQEventSource**](IoArgoprojEventsV1alpha1NSQEventSource.md) |  | [optional] 
**PubSub** | Pointer to [**map[string]IoArgoprojEventsV1alpha1PubSubEventSource**](IoArgoprojEventsV1alpha1PubSubEventSource.md) |  | [optional] 
**Pulsar** | Pointer to [**map[string]IoArgoprojEventsV1alpha1PulsarEventSource**](IoArgoprojEventsV1alpha1PulsarEventSource.md) |  | [optional] 
**Redis** | Pointer to [**map[string]IoArgoprojEventsV1alpha1RedisEventSource**](IoArgoprojEventsV1alpha1RedisEventSource.md) |  | [optional] 
**Replica** | Pointer to **int32** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Resource** | Pointer to [**map[string]IoArgoprojEventsV1alpha1ResourceEventSource**](IoArgoprojEventsV1alpha1ResourceEventSource.md) |  | [optional] 
**Service** | Pointer to [**IoArgoprojEventsV1alpha1Service**](IoArgoprojEventsV1alpha1Service.md) |  | [optional] 
**Slack** | Pointer to [**map[string]IoArgoprojEventsV1alpha1SlackEventSource**](IoArgoprojEventsV1alpha1SlackEventSource.md) |  | [optional] 
**Sns** | Pointer to [**map[string]IoArgoprojEventsV1alpha1SNSEventSource**](IoArgoprojEventsV1alpha1SNSEventSource.md) |  | [optional] 
**Sqs** | Pointer to [**map[string]IoArgoprojEventsV1alpha1SQSEventSource**](IoArgoprojEventsV1alpha1SQSEventSource.md) |  | [optional] 
**StorageGrid** | Pointer to [**map[string]IoArgoprojEventsV1alpha1StorageGridEventSource**](IoArgoprojEventsV1alpha1StorageGridEventSource.md) |  | [optional] 
**Stripe** | Pointer to [**map[string]IoArgoprojEventsV1alpha1StripeEventSource**](IoArgoprojEventsV1alpha1StripeEventSource.md) |  | [optional] 
**Template** | Pointer to [**IoArgoprojEventsV1alpha1Template**](IoArgoprojEventsV1alpha1Template.md) |  | [optional] 
**Webhook** | Pointer to [**map[string]IoArgoprojEventsV1alpha1WebhookContext**](IoArgoprojEventsV1alpha1WebhookContext.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1EventSourceSpec

`func NewIoArgoprojEventsV1alpha1EventSourceSpec() *IoArgoprojEventsV1alpha1EventSourceSpec`

NewIoArgoprojEventsV1alpha1EventSourceSpec instantiates a new IoArgoprojEventsV1alpha1EventSourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1EventSourceSpecWithDefaults

`func NewIoArgoprojEventsV1alpha1EventSourceSpecWithDefaults() *IoArgoprojEventsV1alpha1EventSourceSpec`

NewIoArgoprojEventsV1alpha1EventSourceSpecWithDefaults instantiates a new IoArgoprojEventsV1alpha1EventSourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmqp

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetAmqp() map[string]IoArgoprojEventsV1alpha1AMQPEventSource`

GetAmqp returns the Amqp field if non-nil, zero value otherwise.

### GetAmqpOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetAmqpOk() (*map[string]IoArgoprojEventsV1alpha1AMQPEventSource, bool)`

GetAmqpOk returns a tuple with the Amqp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmqp

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetAmqp(v map[string]IoArgoprojEventsV1alpha1AMQPEventSource)`

SetAmqp sets Amqp field to given value.

### HasAmqp

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasAmqp() bool`

HasAmqp returns a boolean if a field has been set.

### GetAzureEventsHub

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetAzureEventsHub() map[string]IoArgoprojEventsV1alpha1AzureEventsHubEventSource`

GetAzureEventsHub returns the AzureEventsHub field if non-nil, zero value otherwise.

### GetAzureEventsHubOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetAzureEventsHubOk() (*map[string]IoArgoprojEventsV1alpha1AzureEventsHubEventSource, bool)`

GetAzureEventsHubOk returns a tuple with the AzureEventsHub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureEventsHub

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetAzureEventsHub(v map[string]IoArgoprojEventsV1alpha1AzureEventsHubEventSource)`

SetAzureEventsHub sets AzureEventsHub field to given value.

### HasAzureEventsHub

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasAzureEventsHub() bool`

HasAzureEventsHub returns a boolean if a field has been set.

### GetCalendar

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetCalendar() map[string]IoArgoprojEventsV1alpha1CalendarEventSource`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetCalendarOk() (*map[string]IoArgoprojEventsV1alpha1CalendarEventSource, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetCalendar(v map[string]IoArgoprojEventsV1alpha1CalendarEventSource)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### GetEmitter

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetEmitter() map[string]IoArgoprojEventsV1alpha1EmitterEventSource`

GetEmitter returns the Emitter field if non-nil, zero value otherwise.

### GetEmitterOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetEmitterOk() (*map[string]IoArgoprojEventsV1alpha1EmitterEventSource, bool)`

GetEmitterOk returns a tuple with the Emitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmitter

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetEmitter(v map[string]IoArgoprojEventsV1alpha1EmitterEventSource)`

SetEmitter sets Emitter field to given value.

### HasEmitter

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasEmitter() bool`

HasEmitter returns a boolean if a field has been set.

### GetEventBusName

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetEventBusName() string`

GetEventBusName returns the EventBusName field if non-nil, zero value otherwise.

### GetEventBusNameOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetEventBusNameOk() (*string, bool)`

GetEventBusNameOk returns a tuple with the EventBusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventBusName

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetEventBusName(v string)`

SetEventBusName sets EventBusName field to given value.

### HasEventBusName

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasEventBusName() bool`

HasEventBusName returns a boolean if a field has been set.

### GetFile

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetFile() map[string]IoArgoprojEventsV1alpha1FileEventSource`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetFileOk() (*map[string]IoArgoprojEventsV1alpha1FileEventSource, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetFile(v map[string]IoArgoprojEventsV1alpha1FileEventSource)`

SetFile sets File field to given value.

### HasFile

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetGeneric

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetGeneric() map[string]IoArgoprojEventsV1alpha1GenericEventSource`

GetGeneric returns the Generic field if non-nil, zero value otherwise.

### GetGenericOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetGenericOk() (*map[string]IoArgoprojEventsV1alpha1GenericEventSource, bool)`

GetGenericOk returns a tuple with the Generic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneric

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetGeneric(v map[string]IoArgoprojEventsV1alpha1GenericEventSource)`

SetGeneric sets Generic field to given value.

### HasGeneric

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasGeneric() bool`

HasGeneric returns a boolean if a field has been set.

### GetGithub

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetGithub() map[string]IoArgoprojEventsV1alpha1GithubEventSource`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetGithubOk() (*map[string]IoArgoprojEventsV1alpha1GithubEventSource, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetGithub(v map[string]IoArgoprojEventsV1alpha1GithubEventSource)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetGitlab

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetGitlab() map[string]IoArgoprojEventsV1alpha1GitlabEventSource`

GetGitlab returns the Gitlab field if non-nil, zero value otherwise.

### GetGitlabOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetGitlabOk() (*map[string]IoArgoprojEventsV1alpha1GitlabEventSource, bool)`

GetGitlabOk returns a tuple with the Gitlab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlab

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetGitlab(v map[string]IoArgoprojEventsV1alpha1GitlabEventSource)`

SetGitlab sets Gitlab field to given value.

### HasGitlab

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasGitlab() bool`

HasGitlab returns a boolean if a field has been set.

### GetHdfs

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetHdfs() map[string]IoArgoprojEventsV1alpha1HDFSEventSource`

GetHdfs returns the Hdfs field if non-nil, zero value otherwise.

### GetHdfsOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetHdfsOk() (*map[string]IoArgoprojEventsV1alpha1HDFSEventSource, bool)`

GetHdfsOk returns a tuple with the Hdfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdfs

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetHdfs(v map[string]IoArgoprojEventsV1alpha1HDFSEventSource)`

SetHdfs sets Hdfs field to given value.

### HasHdfs

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasHdfs() bool`

HasHdfs returns a boolean if a field has been set.

### GetKafka

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetKafka() map[string]IoArgoprojEventsV1alpha1KafkaEventSource`

GetKafka returns the Kafka field if non-nil, zero value otherwise.

### GetKafkaOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetKafkaOk() (*map[string]IoArgoprojEventsV1alpha1KafkaEventSource, bool)`

GetKafkaOk returns a tuple with the Kafka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafka

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetKafka(v map[string]IoArgoprojEventsV1alpha1KafkaEventSource)`

SetKafka sets Kafka field to given value.

### HasKafka

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasKafka() bool`

HasKafka returns a boolean if a field has been set.

### GetMinio

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetMinio() map[string]IoArgoprojEventsV1alpha1S3Artifact`

GetMinio returns the Minio field if non-nil, zero value otherwise.

### GetMinioOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetMinioOk() (*map[string]IoArgoprojEventsV1alpha1S3Artifact, bool)`

GetMinioOk returns a tuple with the Minio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinio

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetMinio(v map[string]IoArgoprojEventsV1alpha1S3Artifact)`

SetMinio sets Minio field to given value.

### HasMinio

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasMinio() bool`

HasMinio returns a boolean if a field has been set.

### GetMqtt

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetMqtt() map[string]IoArgoprojEventsV1alpha1MQTTEventSource`

GetMqtt returns the Mqtt field if non-nil, zero value otherwise.

### GetMqttOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetMqttOk() (*map[string]IoArgoprojEventsV1alpha1MQTTEventSource, bool)`

GetMqttOk returns a tuple with the Mqtt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqtt

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetMqtt(v map[string]IoArgoprojEventsV1alpha1MQTTEventSource)`

SetMqtt sets Mqtt field to given value.

### HasMqtt

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasMqtt() bool`

HasMqtt returns a boolean if a field has been set.

### GetNats

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetNats() map[string]IoArgoprojEventsV1alpha1NATSEventsSource`

GetNats returns the Nats field if non-nil, zero value otherwise.

### GetNatsOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetNatsOk() (*map[string]IoArgoprojEventsV1alpha1NATSEventsSource, bool)`

GetNatsOk returns a tuple with the Nats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNats

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetNats(v map[string]IoArgoprojEventsV1alpha1NATSEventsSource)`

SetNats sets Nats field to given value.

### HasNats

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasNats() bool`

HasNats returns a boolean if a field has been set.

### GetNsq

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetNsq() map[string]IoArgoprojEventsV1alpha1NSQEventSource`

GetNsq returns the Nsq field if non-nil, zero value otherwise.

### GetNsqOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetNsqOk() (*map[string]IoArgoprojEventsV1alpha1NSQEventSource, bool)`

GetNsqOk returns a tuple with the Nsq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsq

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetNsq(v map[string]IoArgoprojEventsV1alpha1NSQEventSource)`

SetNsq sets Nsq field to given value.

### HasNsq

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasNsq() bool`

HasNsq returns a boolean if a field has been set.

### GetPubSub

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetPubSub() map[string]IoArgoprojEventsV1alpha1PubSubEventSource`

GetPubSub returns the PubSub field if non-nil, zero value otherwise.

### GetPubSubOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetPubSubOk() (*map[string]IoArgoprojEventsV1alpha1PubSubEventSource, bool)`

GetPubSubOk returns a tuple with the PubSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubSub

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetPubSub(v map[string]IoArgoprojEventsV1alpha1PubSubEventSource)`

SetPubSub sets PubSub field to given value.

### HasPubSub

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasPubSub() bool`

HasPubSub returns a boolean if a field has been set.

### GetPulsar

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetPulsar() map[string]IoArgoprojEventsV1alpha1PulsarEventSource`

GetPulsar returns the Pulsar field if non-nil, zero value otherwise.

### GetPulsarOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetPulsarOk() (*map[string]IoArgoprojEventsV1alpha1PulsarEventSource, bool)`

GetPulsarOk returns a tuple with the Pulsar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulsar

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetPulsar(v map[string]IoArgoprojEventsV1alpha1PulsarEventSource)`

SetPulsar sets Pulsar field to given value.

### HasPulsar

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasPulsar() bool`

HasPulsar returns a boolean if a field has been set.

### GetRedis

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetRedis() map[string]IoArgoprojEventsV1alpha1RedisEventSource`

GetRedis returns the Redis field if non-nil, zero value otherwise.

### GetRedisOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetRedisOk() (*map[string]IoArgoprojEventsV1alpha1RedisEventSource, bool)`

GetRedisOk returns a tuple with the Redis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedis

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetRedis(v map[string]IoArgoprojEventsV1alpha1RedisEventSource)`

SetRedis sets Redis field to given value.

### HasRedis

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasRedis() bool`

HasRedis returns a boolean if a field has been set.

### GetReplica

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetReplica() int32`

GetReplica returns the Replica field if non-nil, zero value otherwise.

### GetReplicaOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetReplicaOk() (*int32, bool)`

GetReplicaOk returns a tuple with the Replica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplica

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetReplica(v int32)`

SetReplica sets Replica field to given value.

### HasReplica

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasReplica() bool`

HasReplica returns a boolean if a field has been set.

### GetReplicas

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetResource

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetResource() map[string]IoArgoprojEventsV1alpha1ResourceEventSource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetResourceOk() (*map[string]IoArgoprojEventsV1alpha1ResourceEventSource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetResource(v map[string]IoArgoprojEventsV1alpha1ResourceEventSource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetService

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetService() IoArgoprojEventsV1alpha1Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetServiceOk() (*IoArgoprojEventsV1alpha1Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetService(v IoArgoprojEventsV1alpha1Service)`

SetService sets Service field to given value.

### HasService

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSlack

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetSlack() map[string]IoArgoprojEventsV1alpha1SlackEventSource`

GetSlack returns the Slack field if non-nil, zero value otherwise.

### GetSlackOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetSlackOk() (*map[string]IoArgoprojEventsV1alpha1SlackEventSource, bool)`

GetSlackOk returns a tuple with the Slack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlack

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetSlack(v map[string]IoArgoprojEventsV1alpha1SlackEventSource)`

SetSlack sets Slack field to given value.

### HasSlack

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasSlack() bool`

HasSlack returns a boolean if a field has been set.

### GetSns

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetSns() map[string]IoArgoprojEventsV1alpha1SNSEventSource`

GetSns returns the Sns field if non-nil, zero value otherwise.

### GetSnsOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetSnsOk() (*map[string]IoArgoprojEventsV1alpha1SNSEventSource, bool)`

GetSnsOk returns a tuple with the Sns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSns

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetSns(v map[string]IoArgoprojEventsV1alpha1SNSEventSource)`

SetSns sets Sns field to given value.

### HasSns

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasSns() bool`

HasSns returns a boolean if a field has been set.

### GetSqs

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetSqs() map[string]IoArgoprojEventsV1alpha1SQSEventSource`

GetSqs returns the Sqs field if non-nil, zero value otherwise.

### GetSqsOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetSqsOk() (*map[string]IoArgoprojEventsV1alpha1SQSEventSource, bool)`

GetSqsOk returns a tuple with the Sqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqs

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetSqs(v map[string]IoArgoprojEventsV1alpha1SQSEventSource)`

SetSqs sets Sqs field to given value.

### HasSqs

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasSqs() bool`

HasSqs returns a boolean if a field has been set.

### GetStorageGrid

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetStorageGrid() map[string]IoArgoprojEventsV1alpha1StorageGridEventSource`

GetStorageGrid returns the StorageGrid field if non-nil, zero value otherwise.

### GetStorageGridOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetStorageGridOk() (*map[string]IoArgoprojEventsV1alpha1StorageGridEventSource, bool)`

GetStorageGridOk returns a tuple with the StorageGrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGrid

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetStorageGrid(v map[string]IoArgoprojEventsV1alpha1StorageGridEventSource)`

SetStorageGrid sets StorageGrid field to given value.

### HasStorageGrid

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasStorageGrid() bool`

HasStorageGrid returns a boolean if a field has been set.

### GetStripe

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetStripe() map[string]IoArgoprojEventsV1alpha1StripeEventSource`

GetStripe returns the Stripe field if non-nil, zero value otherwise.

### GetStripeOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetStripeOk() (*map[string]IoArgoprojEventsV1alpha1StripeEventSource, bool)`

GetStripeOk returns a tuple with the Stripe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripe

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetStripe(v map[string]IoArgoprojEventsV1alpha1StripeEventSource)`

SetStripe sets Stripe field to given value.

### HasStripe

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasStripe() bool`

HasStripe returns a boolean if a field has been set.

### GetTemplate

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetTemplate() IoArgoprojEventsV1alpha1Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetTemplateOk() (*IoArgoprojEventsV1alpha1Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetTemplate(v IoArgoprojEventsV1alpha1Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetWebhook

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetWebhook() map[string]IoArgoprojEventsV1alpha1WebhookContext`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetWebhookOk() (*map[string]IoArgoprojEventsV1alpha1WebhookContext, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetWebhook(v map[string]IoArgoprojEventsV1alpha1WebhookContext)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


