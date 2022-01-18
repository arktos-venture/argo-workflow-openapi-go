# GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Db** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink**](GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink.md) |  | [optional] 
**Http** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSink**](GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSink.md) |  | [optional] 
**Kafka** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSink**](GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSink.md) |  | [optional] 
**Log** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Log**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Log.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**S3** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1S3Sink**](GithubComArgoprojLabsArgoDataflowApiV1alpha1S3Sink.md) |  | [optional] 
**Stan** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN**](GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN.md) |  | [optional] 
**Volume** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1VolumeSink**](GithubComArgoprojLabsArgoDataflowApiV1alpha1VolumeSink.md) |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Sink

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Sink() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Sink instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1SinkWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1SinkWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1SinkWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDb

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetDb() GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetDbOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) SetDb(v GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSink)`

SetDb sets Db field to given value.

### HasDb

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) HasDb() bool`

HasDb returns a boolean if a field has been set.

### GetHttp

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetHttp() GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSink`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetHttpOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSink, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) SetHttp(v GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSink)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetKafka

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetKafka() GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSink`

GetKafka returns the Kafka field if non-nil, zero value otherwise.

### GetKafkaOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetKafkaOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSink, bool)`

GetKafkaOk returns a tuple with the Kafka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafka

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) SetKafka(v GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSink)`

SetKafka sets Kafka field to given value.

### HasKafka

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) HasKafka() bool`

HasKafka returns a boolean if a field has been set.

### GetLog

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetLog() GithubComArgoprojLabsArgoDataflowApiV1alpha1Log`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetLogOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Log, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) SetLog(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Log)`

SetLog sets Log field to given value.

### HasLog

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetS3

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetS3() GithubComArgoprojLabsArgoDataflowApiV1alpha1S3Sink`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetS3Ok() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1S3Sink, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) SetS3(v GithubComArgoprojLabsArgoDataflowApiV1alpha1S3Sink)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetStan

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetStan() GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN`

GetStan returns the Stan field if non-nil, zero value otherwise.

### GetStanOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetStanOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN, bool)`

GetStanOk returns a tuple with the Stan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStan

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) SetStan(v GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN)`

SetStan sets Stan field to given value.

### HasStan

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) HasStan() bool`

HasStan returns a boolean if a field has been set.

### GetVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetVolume() GithubComArgoprojLabsArgoDataflowApiV1alpha1VolumeSink`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) GetVolumeOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1VolumeSink, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) SetVolume(v GithubComArgoprojLabsArgoDataflowApiV1alpha1VolumeSink)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Sink) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


