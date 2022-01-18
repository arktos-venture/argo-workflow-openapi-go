/*
Argo Server API

You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`

API version: VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IoArgoprojEventsV1alpha1EventSourceSpec struct for IoArgoprojEventsV1alpha1EventSourceSpec
type IoArgoprojEventsV1alpha1EventSourceSpec struct {
	Amqp *map[string]IoArgoprojEventsV1alpha1AMQPEventSource `json:"amqp,omitempty"`
	AzureEventsHub *map[string]IoArgoprojEventsV1alpha1AzureEventsHubEventSource `json:"azureEventsHub,omitempty"`
	Calendar *map[string]IoArgoprojEventsV1alpha1CalendarEventSource `json:"calendar,omitempty"`
	Emitter *map[string]IoArgoprojEventsV1alpha1EmitterEventSource `json:"emitter,omitempty"`
	EventBusName *string `json:"eventBusName,omitempty"`
	File *map[string]IoArgoprojEventsV1alpha1FileEventSource `json:"file,omitempty"`
	Generic *map[string]IoArgoprojEventsV1alpha1GenericEventSource `json:"generic,omitempty"`
	Github *map[string]IoArgoprojEventsV1alpha1GithubEventSource `json:"github,omitempty"`
	Gitlab *map[string]IoArgoprojEventsV1alpha1GitlabEventSource `json:"gitlab,omitempty"`
	Hdfs *map[string]IoArgoprojEventsV1alpha1HDFSEventSource `json:"hdfs,omitempty"`
	Kafka *map[string]IoArgoprojEventsV1alpha1KafkaEventSource `json:"kafka,omitempty"`
	Minio *map[string]IoArgoprojEventsV1alpha1S3Artifact `json:"minio,omitempty"`
	Mqtt *map[string]IoArgoprojEventsV1alpha1MQTTEventSource `json:"mqtt,omitempty"`
	Nats *map[string]IoArgoprojEventsV1alpha1NATSEventsSource `json:"nats,omitempty"`
	Nsq *map[string]IoArgoprojEventsV1alpha1NSQEventSource `json:"nsq,omitempty"`
	PubSub *map[string]IoArgoprojEventsV1alpha1PubSubEventSource `json:"pubSub,omitempty"`
	Pulsar *map[string]IoArgoprojEventsV1alpha1PulsarEventSource `json:"pulsar,omitempty"`
	Redis *map[string]IoArgoprojEventsV1alpha1RedisEventSource `json:"redis,omitempty"`
	Replica *int32 `json:"replica,omitempty"`
	Replicas *int32 `json:"replicas,omitempty"`
	Resource *map[string]IoArgoprojEventsV1alpha1ResourceEventSource `json:"resource,omitempty"`
	Service *IoArgoprojEventsV1alpha1Service `json:"service,omitempty"`
	Slack *map[string]IoArgoprojEventsV1alpha1SlackEventSource `json:"slack,omitempty"`
	Sns *map[string]IoArgoprojEventsV1alpha1SNSEventSource `json:"sns,omitempty"`
	Sqs *map[string]IoArgoprojEventsV1alpha1SQSEventSource `json:"sqs,omitempty"`
	StorageGrid *map[string]IoArgoprojEventsV1alpha1StorageGridEventSource `json:"storageGrid,omitempty"`
	Stripe *map[string]IoArgoprojEventsV1alpha1StripeEventSource `json:"stripe,omitempty"`
	Template *IoArgoprojEventsV1alpha1Template `json:"template,omitempty"`
	Webhook *map[string]IoArgoprojEventsV1alpha1WebhookContext `json:"webhook,omitempty"`
}

// NewIoArgoprojEventsV1alpha1EventSourceSpec instantiates a new IoArgoprojEventsV1alpha1EventSourceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1EventSourceSpec() *IoArgoprojEventsV1alpha1EventSourceSpec {
	this := IoArgoprojEventsV1alpha1EventSourceSpec{}
	return &this
}

// NewIoArgoprojEventsV1alpha1EventSourceSpecWithDefaults instantiates a new IoArgoprojEventsV1alpha1EventSourceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1EventSourceSpecWithDefaults() *IoArgoprojEventsV1alpha1EventSourceSpec {
	this := IoArgoprojEventsV1alpha1EventSourceSpec{}
	return &this
}

// GetAmqp returns the Amqp field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetAmqp() map[string]IoArgoprojEventsV1alpha1AMQPEventSource {
	if o == nil || o.Amqp == nil {
		var ret map[string]IoArgoprojEventsV1alpha1AMQPEventSource
		return ret
	}
	return *o.Amqp
}

// GetAmqpOk returns a tuple with the Amqp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetAmqpOk() (*map[string]IoArgoprojEventsV1alpha1AMQPEventSource, bool) {
	if o == nil || o.Amqp == nil {
		return nil, false
	}
	return o.Amqp, true
}

// HasAmqp returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasAmqp() bool {
	if o != nil && o.Amqp != nil {
		return true
	}

	return false
}

// SetAmqp gets a reference to the given map[string]IoArgoprojEventsV1alpha1AMQPEventSource and assigns it to the Amqp field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetAmqp(v map[string]IoArgoprojEventsV1alpha1AMQPEventSource) {
	o.Amqp = &v
}

// GetAzureEventsHub returns the AzureEventsHub field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetAzureEventsHub() map[string]IoArgoprojEventsV1alpha1AzureEventsHubEventSource {
	if o == nil || o.AzureEventsHub == nil {
		var ret map[string]IoArgoprojEventsV1alpha1AzureEventsHubEventSource
		return ret
	}
	return *o.AzureEventsHub
}

// GetAzureEventsHubOk returns a tuple with the AzureEventsHub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetAzureEventsHubOk() (*map[string]IoArgoprojEventsV1alpha1AzureEventsHubEventSource, bool) {
	if o == nil || o.AzureEventsHub == nil {
		return nil, false
	}
	return o.AzureEventsHub, true
}

// HasAzureEventsHub returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasAzureEventsHub() bool {
	if o != nil && o.AzureEventsHub != nil {
		return true
	}

	return false
}

// SetAzureEventsHub gets a reference to the given map[string]IoArgoprojEventsV1alpha1AzureEventsHubEventSource and assigns it to the AzureEventsHub field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetAzureEventsHub(v map[string]IoArgoprojEventsV1alpha1AzureEventsHubEventSource) {
	o.AzureEventsHub = &v
}

// GetCalendar returns the Calendar field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetCalendar() map[string]IoArgoprojEventsV1alpha1CalendarEventSource {
	if o == nil || o.Calendar == nil {
		var ret map[string]IoArgoprojEventsV1alpha1CalendarEventSource
		return ret
	}
	return *o.Calendar
}

// GetCalendarOk returns a tuple with the Calendar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetCalendarOk() (*map[string]IoArgoprojEventsV1alpha1CalendarEventSource, bool) {
	if o == nil || o.Calendar == nil {
		return nil, false
	}
	return o.Calendar, true
}

// HasCalendar returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasCalendar() bool {
	if o != nil && o.Calendar != nil {
		return true
	}

	return false
}

// SetCalendar gets a reference to the given map[string]IoArgoprojEventsV1alpha1CalendarEventSource and assigns it to the Calendar field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetCalendar(v map[string]IoArgoprojEventsV1alpha1CalendarEventSource) {
	o.Calendar = &v
}

// GetEmitter returns the Emitter field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetEmitter() map[string]IoArgoprojEventsV1alpha1EmitterEventSource {
	if o == nil || o.Emitter == nil {
		var ret map[string]IoArgoprojEventsV1alpha1EmitterEventSource
		return ret
	}
	return *o.Emitter
}

// GetEmitterOk returns a tuple with the Emitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetEmitterOk() (*map[string]IoArgoprojEventsV1alpha1EmitterEventSource, bool) {
	if o == nil || o.Emitter == nil {
		return nil, false
	}
	return o.Emitter, true
}

// HasEmitter returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasEmitter() bool {
	if o != nil && o.Emitter != nil {
		return true
	}

	return false
}

// SetEmitter gets a reference to the given map[string]IoArgoprojEventsV1alpha1EmitterEventSource and assigns it to the Emitter field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetEmitter(v map[string]IoArgoprojEventsV1alpha1EmitterEventSource) {
	o.Emitter = &v
}

// GetEventBusName returns the EventBusName field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetEventBusName() string {
	if o == nil || o.EventBusName == nil {
		var ret string
		return ret
	}
	return *o.EventBusName
}

// GetEventBusNameOk returns a tuple with the EventBusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetEventBusNameOk() (*string, bool) {
	if o == nil || o.EventBusName == nil {
		return nil, false
	}
	return o.EventBusName, true
}

// HasEventBusName returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasEventBusName() bool {
	if o != nil && o.EventBusName != nil {
		return true
	}

	return false
}

// SetEventBusName gets a reference to the given string and assigns it to the EventBusName field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetEventBusName(v string) {
	o.EventBusName = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetFile() map[string]IoArgoprojEventsV1alpha1FileEventSource {
	if o == nil || o.File == nil {
		var ret map[string]IoArgoprojEventsV1alpha1FileEventSource
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetFileOk() (*map[string]IoArgoprojEventsV1alpha1FileEventSource, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given map[string]IoArgoprojEventsV1alpha1FileEventSource and assigns it to the File field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetFile(v map[string]IoArgoprojEventsV1alpha1FileEventSource) {
	o.File = &v
}

// GetGeneric returns the Generic field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetGeneric() map[string]IoArgoprojEventsV1alpha1GenericEventSource {
	if o == nil || o.Generic == nil {
		var ret map[string]IoArgoprojEventsV1alpha1GenericEventSource
		return ret
	}
	return *o.Generic
}

// GetGenericOk returns a tuple with the Generic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetGenericOk() (*map[string]IoArgoprojEventsV1alpha1GenericEventSource, bool) {
	if o == nil || o.Generic == nil {
		return nil, false
	}
	return o.Generic, true
}

// HasGeneric returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasGeneric() bool {
	if o != nil && o.Generic != nil {
		return true
	}

	return false
}

// SetGeneric gets a reference to the given map[string]IoArgoprojEventsV1alpha1GenericEventSource and assigns it to the Generic field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetGeneric(v map[string]IoArgoprojEventsV1alpha1GenericEventSource) {
	o.Generic = &v
}

// GetGithub returns the Github field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetGithub() map[string]IoArgoprojEventsV1alpha1GithubEventSource {
	if o == nil || o.Github == nil {
		var ret map[string]IoArgoprojEventsV1alpha1GithubEventSource
		return ret
	}
	return *o.Github
}

// GetGithubOk returns a tuple with the Github field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetGithubOk() (*map[string]IoArgoprojEventsV1alpha1GithubEventSource, bool) {
	if o == nil || o.Github == nil {
		return nil, false
	}
	return o.Github, true
}

// HasGithub returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasGithub() bool {
	if o != nil && o.Github != nil {
		return true
	}

	return false
}

// SetGithub gets a reference to the given map[string]IoArgoprojEventsV1alpha1GithubEventSource and assigns it to the Github field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetGithub(v map[string]IoArgoprojEventsV1alpha1GithubEventSource) {
	o.Github = &v
}

// GetGitlab returns the Gitlab field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetGitlab() map[string]IoArgoprojEventsV1alpha1GitlabEventSource {
	if o == nil || o.Gitlab == nil {
		var ret map[string]IoArgoprojEventsV1alpha1GitlabEventSource
		return ret
	}
	return *o.Gitlab
}

// GetGitlabOk returns a tuple with the Gitlab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetGitlabOk() (*map[string]IoArgoprojEventsV1alpha1GitlabEventSource, bool) {
	if o == nil || o.Gitlab == nil {
		return nil, false
	}
	return o.Gitlab, true
}

// HasGitlab returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasGitlab() bool {
	if o != nil && o.Gitlab != nil {
		return true
	}

	return false
}

// SetGitlab gets a reference to the given map[string]IoArgoprojEventsV1alpha1GitlabEventSource and assigns it to the Gitlab field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetGitlab(v map[string]IoArgoprojEventsV1alpha1GitlabEventSource) {
	o.Gitlab = &v
}

// GetHdfs returns the Hdfs field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetHdfs() map[string]IoArgoprojEventsV1alpha1HDFSEventSource {
	if o == nil || o.Hdfs == nil {
		var ret map[string]IoArgoprojEventsV1alpha1HDFSEventSource
		return ret
	}
	return *o.Hdfs
}

// GetHdfsOk returns a tuple with the Hdfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetHdfsOk() (*map[string]IoArgoprojEventsV1alpha1HDFSEventSource, bool) {
	if o == nil || o.Hdfs == nil {
		return nil, false
	}
	return o.Hdfs, true
}

// HasHdfs returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasHdfs() bool {
	if o != nil && o.Hdfs != nil {
		return true
	}

	return false
}

// SetHdfs gets a reference to the given map[string]IoArgoprojEventsV1alpha1HDFSEventSource and assigns it to the Hdfs field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetHdfs(v map[string]IoArgoprojEventsV1alpha1HDFSEventSource) {
	o.Hdfs = &v
}

// GetKafka returns the Kafka field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetKafka() map[string]IoArgoprojEventsV1alpha1KafkaEventSource {
	if o == nil || o.Kafka == nil {
		var ret map[string]IoArgoprojEventsV1alpha1KafkaEventSource
		return ret
	}
	return *o.Kafka
}

// GetKafkaOk returns a tuple with the Kafka field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetKafkaOk() (*map[string]IoArgoprojEventsV1alpha1KafkaEventSource, bool) {
	if o == nil || o.Kafka == nil {
		return nil, false
	}
	return o.Kafka, true
}

// HasKafka returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasKafka() bool {
	if o != nil && o.Kafka != nil {
		return true
	}

	return false
}

// SetKafka gets a reference to the given map[string]IoArgoprojEventsV1alpha1KafkaEventSource and assigns it to the Kafka field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetKafka(v map[string]IoArgoprojEventsV1alpha1KafkaEventSource) {
	o.Kafka = &v
}

// GetMinio returns the Minio field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetMinio() map[string]IoArgoprojEventsV1alpha1S3Artifact {
	if o == nil || o.Minio == nil {
		var ret map[string]IoArgoprojEventsV1alpha1S3Artifact
		return ret
	}
	return *o.Minio
}

// GetMinioOk returns a tuple with the Minio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetMinioOk() (*map[string]IoArgoprojEventsV1alpha1S3Artifact, bool) {
	if o == nil || o.Minio == nil {
		return nil, false
	}
	return o.Minio, true
}

// HasMinio returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasMinio() bool {
	if o != nil && o.Minio != nil {
		return true
	}

	return false
}

// SetMinio gets a reference to the given map[string]IoArgoprojEventsV1alpha1S3Artifact and assigns it to the Minio field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetMinio(v map[string]IoArgoprojEventsV1alpha1S3Artifact) {
	o.Minio = &v
}

// GetMqtt returns the Mqtt field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetMqtt() map[string]IoArgoprojEventsV1alpha1MQTTEventSource {
	if o == nil || o.Mqtt == nil {
		var ret map[string]IoArgoprojEventsV1alpha1MQTTEventSource
		return ret
	}
	return *o.Mqtt
}

// GetMqttOk returns a tuple with the Mqtt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetMqttOk() (*map[string]IoArgoprojEventsV1alpha1MQTTEventSource, bool) {
	if o == nil || o.Mqtt == nil {
		return nil, false
	}
	return o.Mqtt, true
}

// HasMqtt returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasMqtt() bool {
	if o != nil && o.Mqtt != nil {
		return true
	}

	return false
}

// SetMqtt gets a reference to the given map[string]IoArgoprojEventsV1alpha1MQTTEventSource and assigns it to the Mqtt field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetMqtt(v map[string]IoArgoprojEventsV1alpha1MQTTEventSource) {
	o.Mqtt = &v
}

// GetNats returns the Nats field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetNats() map[string]IoArgoprojEventsV1alpha1NATSEventsSource {
	if o == nil || o.Nats == nil {
		var ret map[string]IoArgoprojEventsV1alpha1NATSEventsSource
		return ret
	}
	return *o.Nats
}

// GetNatsOk returns a tuple with the Nats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetNatsOk() (*map[string]IoArgoprojEventsV1alpha1NATSEventsSource, bool) {
	if o == nil || o.Nats == nil {
		return nil, false
	}
	return o.Nats, true
}

// HasNats returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasNats() bool {
	if o != nil && o.Nats != nil {
		return true
	}

	return false
}

// SetNats gets a reference to the given map[string]IoArgoprojEventsV1alpha1NATSEventsSource and assigns it to the Nats field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetNats(v map[string]IoArgoprojEventsV1alpha1NATSEventsSource) {
	o.Nats = &v
}

// GetNsq returns the Nsq field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetNsq() map[string]IoArgoprojEventsV1alpha1NSQEventSource {
	if o == nil || o.Nsq == nil {
		var ret map[string]IoArgoprojEventsV1alpha1NSQEventSource
		return ret
	}
	return *o.Nsq
}

// GetNsqOk returns a tuple with the Nsq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetNsqOk() (*map[string]IoArgoprojEventsV1alpha1NSQEventSource, bool) {
	if o == nil || o.Nsq == nil {
		return nil, false
	}
	return o.Nsq, true
}

// HasNsq returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasNsq() bool {
	if o != nil && o.Nsq != nil {
		return true
	}

	return false
}

// SetNsq gets a reference to the given map[string]IoArgoprojEventsV1alpha1NSQEventSource and assigns it to the Nsq field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetNsq(v map[string]IoArgoprojEventsV1alpha1NSQEventSource) {
	o.Nsq = &v
}

// GetPubSub returns the PubSub field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetPubSub() map[string]IoArgoprojEventsV1alpha1PubSubEventSource {
	if o == nil || o.PubSub == nil {
		var ret map[string]IoArgoprojEventsV1alpha1PubSubEventSource
		return ret
	}
	return *o.PubSub
}

// GetPubSubOk returns a tuple with the PubSub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetPubSubOk() (*map[string]IoArgoprojEventsV1alpha1PubSubEventSource, bool) {
	if o == nil || o.PubSub == nil {
		return nil, false
	}
	return o.PubSub, true
}

// HasPubSub returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasPubSub() bool {
	if o != nil && o.PubSub != nil {
		return true
	}

	return false
}

// SetPubSub gets a reference to the given map[string]IoArgoprojEventsV1alpha1PubSubEventSource and assigns it to the PubSub field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetPubSub(v map[string]IoArgoprojEventsV1alpha1PubSubEventSource) {
	o.PubSub = &v
}

// GetPulsar returns the Pulsar field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetPulsar() map[string]IoArgoprojEventsV1alpha1PulsarEventSource {
	if o == nil || o.Pulsar == nil {
		var ret map[string]IoArgoprojEventsV1alpha1PulsarEventSource
		return ret
	}
	return *o.Pulsar
}

// GetPulsarOk returns a tuple with the Pulsar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetPulsarOk() (*map[string]IoArgoprojEventsV1alpha1PulsarEventSource, bool) {
	if o == nil || o.Pulsar == nil {
		return nil, false
	}
	return o.Pulsar, true
}

// HasPulsar returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasPulsar() bool {
	if o != nil && o.Pulsar != nil {
		return true
	}

	return false
}

// SetPulsar gets a reference to the given map[string]IoArgoprojEventsV1alpha1PulsarEventSource and assigns it to the Pulsar field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetPulsar(v map[string]IoArgoprojEventsV1alpha1PulsarEventSource) {
	o.Pulsar = &v
}

// GetRedis returns the Redis field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetRedis() map[string]IoArgoprojEventsV1alpha1RedisEventSource {
	if o == nil || o.Redis == nil {
		var ret map[string]IoArgoprojEventsV1alpha1RedisEventSource
		return ret
	}
	return *o.Redis
}

// GetRedisOk returns a tuple with the Redis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetRedisOk() (*map[string]IoArgoprojEventsV1alpha1RedisEventSource, bool) {
	if o == nil || o.Redis == nil {
		return nil, false
	}
	return o.Redis, true
}

// HasRedis returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasRedis() bool {
	if o != nil && o.Redis != nil {
		return true
	}

	return false
}

// SetRedis gets a reference to the given map[string]IoArgoprojEventsV1alpha1RedisEventSource and assigns it to the Redis field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetRedis(v map[string]IoArgoprojEventsV1alpha1RedisEventSource) {
	o.Redis = &v
}

// GetReplica returns the Replica field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetReplica() int32 {
	if o == nil || o.Replica == nil {
		var ret int32
		return ret
	}
	return *o.Replica
}

// GetReplicaOk returns a tuple with the Replica field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetReplicaOk() (*int32, bool) {
	if o == nil || o.Replica == nil {
		return nil, false
	}
	return o.Replica, true
}

// HasReplica returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasReplica() bool {
	if o != nil && o.Replica != nil {
		return true
	}

	return false
}

// SetReplica gets a reference to the given int32 and assigns it to the Replica field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetReplica(v int32) {
	o.Replica = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasReplicas() bool {
	if o != nil && o.Replicas != nil {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetResource() map[string]IoArgoprojEventsV1alpha1ResourceEventSource {
	if o == nil || o.Resource == nil {
		var ret map[string]IoArgoprojEventsV1alpha1ResourceEventSource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetResourceOk() (*map[string]IoArgoprojEventsV1alpha1ResourceEventSource, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given map[string]IoArgoprojEventsV1alpha1ResourceEventSource and assigns it to the Resource field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetResource(v map[string]IoArgoprojEventsV1alpha1ResourceEventSource) {
	o.Resource = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetService() IoArgoprojEventsV1alpha1Service {
	if o == nil || o.Service == nil {
		var ret IoArgoprojEventsV1alpha1Service
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetServiceOk() (*IoArgoprojEventsV1alpha1Service, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given IoArgoprojEventsV1alpha1Service and assigns it to the Service field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetService(v IoArgoprojEventsV1alpha1Service) {
	o.Service = &v
}

// GetSlack returns the Slack field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetSlack() map[string]IoArgoprojEventsV1alpha1SlackEventSource {
	if o == nil || o.Slack == nil {
		var ret map[string]IoArgoprojEventsV1alpha1SlackEventSource
		return ret
	}
	return *o.Slack
}

// GetSlackOk returns a tuple with the Slack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetSlackOk() (*map[string]IoArgoprojEventsV1alpha1SlackEventSource, bool) {
	if o == nil || o.Slack == nil {
		return nil, false
	}
	return o.Slack, true
}

// HasSlack returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasSlack() bool {
	if o != nil && o.Slack != nil {
		return true
	}

	return false
}

// SetSlack gets a reference to the given map[string]IoArgoprojEventsV1alpha1SlackEventSource and assigns it to the Slack field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetSlack(v map[string]IoArgoprojEventsV1alpha1SlackEventSource) {
	o.Slack = &v
}

// GetSns returns the Sns field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetSns() map[string]IoArgoprojEventsV1alpha1SNSEventSource {
	if o == nil || o.Sns == nil {
		var ret map[string]IoArgoprojEventsV1alpha1SNSEventSource
		return ret
	}
	return *o.Sns
}

// GetSnsOk returns a tuple with the Sns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetSnsOk() (*map[string]IoArgoprojEventsV1alpha1SNSEventSource, bool) {
	if o == nil || o.Sns == nil {
		return nil, false
	}
	return o.Sns, true
}

// HasSns returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasSns() bool {
	if o != nil && o.Sns != nil {
		return true
	}

	return false
}

// SetSns gets a reference to the given map[string]IoArgoprojEventsV1alpha1SNSEventSource and assigns it to the Sns field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetSns(v map[string]IoArgoprojEventsV1alpha1SNSEventSource) {
	o.Sns = &v
}

// GetSqs returns the Sqs field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetSqs() map[string]IoArgoprojEventsV1alpha1SQSEventSource {
	if o == nil || o.Sqs == nil {
		var ret map[string]IoArgoprojEventsV1alpha1SQSEventSource
		return ret
	}
	return *o.Sqs
}

// GetSqsOk returns a tuple with the Sqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetSqsOk() (*map[string]IoArgoprojEventsV1alpha1SQSEventSource, bool) {
	if o == nil || o.Sqs == nil {
		return nil, false
	}
	return o.Sqs, true
}

// HasSqs returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasSqs() bool {
	if o != nil && o.Sqs != nil {
		return true
	}

	return false
}

// SetSqs gets a reference to the given map[string]IoArgoprojEventsV1alpha1SQSEventSource and assigns it to the Sqs field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetSqs(v map[string]IoArgoprojEventsV1alpha1SQSEventSource) {
	o.Sqs = &v
}

// GetStorageGrid returns the StorageGrid field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetStorageGrid() map[string]IoArgoprojEventsV1alpha1StorageGridEventSource {
	if o == nil || o.StorageGrid == nil {
		var ret map[string]IoArgoprojEventsV1alpha1StorageGridEventSource
		return ret
	}
	return *o.StorageGrid
}

// GetStorageGridOk returns a tuple with the StorageGrid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetStorageGridOk() (*map[string]IoArgoprojEventsV1alpha1StorageGridEventSource, bool) {
	if o == nil || o.StorageGrid == nil {
		return nil, false
	}
	return o.StorageGrid, true
}

// HasStorageGrid returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasStorageGrid() bool {
	if o != nil && o.StorageGrid != nil {
		return true
	}

	return false
}

// SetStorageGrid gets a reference to the given map[string]IoArgoprojEventsV1alpha1StorageGridEventSource and assigns it to the StorageGrid field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetStorageGrid(v map[string]IoArgoprojEventsV1alpha1StorageGridEventSource) {
	o.StorageGrid = &v
}

// GetStripe returns the Stripe field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetStripe() map[string]IoArgoprojEventsV1alpha1StripeEventSource {
	if o == nil || o.Stripe == nil {
		var ret map[string]IoArgoprojEventsV1alpha1StripeEventSource
		return ret
	}
	return *o.Stripe
}

// GetStripeOk returns a tuple with the Stripe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetStripeOk() (*map[string]IoArgoprojEventsV1alpha1StripeEventSource, bool) {
	if o == nil || o.Stripe == nil {
		return nil, false
	}
	return o.Stripe, true
}

// HasStripe returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasStripe() bool {
	if o != nil && o.Stripe != nil {
		return true
	}

	return false
}

// SetStripe gets a reference to the given map[string]IoArgoprojEventsV1alpha1StripeEventSource and assigns it to the Stripe field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetStripe(v map[string]IoArgoprojEventsV1alpha1StripeEventSource) {
	o.Stripe = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetTemplate() IoArgoprojEventsV1alpha1Template {
	if o == nil || o.Template == nil {
		var ret IoArgoprojEventsV1alpha1Template
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetTemplateOk() (*IoArgoprojEventsV1alpha1Template, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given IoArgoprojEventsV1alpha1Template and assigns it to the Template field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetTemplate(v IoArgoprojEventsV1alpha1Template) {
	o.Template = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetWebhook() map[string]IoArgoprojEventsV1alpha1WebhookContext {
	if o == nil || o.Webhook == nil {
		var ret map[string]IoArgoprojEventsV1alpha1WebhookContext
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) GetWebhookOk() (*map[string]IoArgoprojEventsV1alpha1WebhookContext, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given map[string]IoArgoprojEventsV1alpha1WebhookContext and assigns it to the Webhook field.
func (o *IoArgoprojEventsV1alpha1EventSourceSpec) SetWebhook(v map[string]IoArgoprojEventsV1alpha1WebhookContext) {
	o.Webhook = &v
}

func (o IoArgoprojEventsV1alpha1EventSourceSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amqp != nil {
		toSerialize["amqp"] = o.Amqp
	}
	if o.AzureEventsHub != nil {
		toSerialize["azureEventsHub"] = o.AzureEventsHub
	}
	if o.Calendar != nil {
		toSerialize["calendar"] = o.Calendar
	}
	if o.Emitter != nil {
		toSerialize["emitter"] = o.Emitter
	}
	if o.EventBusName != nil {
		toSerialize["eventBusName"] = o.EventBusName
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Generic != nil {
		toSerialize["generic"] = o.Generic
	}
	if o.Github != nil {
		toSerialize["github"] = o.Github
	}
	if o.Gitlab != nil {
		toSerialize["gitlab"] = o.Gitlab
	}
	if o.Hdfs != nil {
		toSerialize["hdfs"] = o.Hdfs
	}
	if o.Kafka != nil {
		toSerialize["kafka"] = o.Kafka
	}
	if o.Minio != nil {
		toSerialize["minio"] = o.Minio
	}
	if o.Mqtt != nil {
		toSerialize["mqtt"] = o.Mqtt
	}
	if o.Nats != nil {
		toSerialize["nats"] = o.Nats
	}
	if o.Nsq != nil {
		toSerialize["nsq"] = o.Nsq
	}
	if o.PubSub != nil {
		toSerialize["pubSub"] = o.PubSub
	}
	if o.Pulsar != nil {
		toSerialize["pulsar"] = o.Pulsar
	}
	if o.Redis != nil {
		toSerialize["redis"] = o.Redis
	}
	if o.Replica != nil {
		toSerialize["replica"] = o.Replica
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Slack != nil {
		toSerialize["slack"] = o.Slack
	}
	if o.Sns != nil {
		toSerialize["sns"] = o.Sns
	}
	if o.Sqs != nil {
		toSerialize["sqs"] = o.Sqs
	}
	if o.StorageGrid != nil {
		toSerialize["storageGrid"] = o.StorageGrid
	}
	if o.Stripe != nil {
		toSerialize["stripe"] = o.Stripe
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1EventSourceSpec struct {
	value *IoArgoprojEventsV1alpha1EventSourceSpec
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1EventSourceSpec) Get() *IoArgoprojEventsV1alpha1EventSourceSpec {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1EventSourceSpec) Set(val *IoArgoprojEventsV1alpha1EventSourceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1EventSourceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1EventSourceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1EventSourceSpec(val *IoArgoprojEventsV1alpha1EventSourceSpec) *NullableIoArgoprojEventsV1alpha1EventSourceSpec {
	return &NullableIoArgoprojEventsV1alpha1EventSourceSpec{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1EventSourceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1EventSourceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


