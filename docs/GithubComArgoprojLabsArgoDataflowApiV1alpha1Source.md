# GithubComArgoprojLabsArgoDataflowApiV1alpha1Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cron** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron.md) |  | [optional] 
**Db** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource**](GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource.md) |  | [optional] 
**Http** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource**](GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource.md) |  | [optional] 
**Kafka** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource**](GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Retry** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Backoff**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Backoff.md) |  | [optional] 
**S3** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1S3Source**](GithubComArgoprojLabsArgoDataflowApiV1alpha1S3Source.md) |  | [optional] 
**Stan** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN**](GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN.md) |  | [optional] 
**Volume** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1VolumeSource**](GithubComArgoprojLabsArgoDataflowApiV1alpha1VolumeSource.md) |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Source

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Source() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Source instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1SourceWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1SourceWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1SourceWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCron

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetCron() GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetCronOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) SetCron(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Cron)`

SetCron sets Cron field to given value.

### HasCron

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetDb

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetDb() GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetDbOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) SetDb(v GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource)`

SetDb sets Db field to given value.

### HasDb

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) HasDb() bool`

HasDb returns a boolean if a field has been set.

### GetHttp

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetHttp() GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetHttpOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) SetHttp(v GithubComArgoprojLabsArgoDataflowApiV1alpha1HTTPSource)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetKafka

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetKafka() GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource`

GetKafka returns the Kafka field if non-nil, zero value otherwise.

### GetKafkaOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetKafkaOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource, bool)`

GetKafkaOk returns a tuple with the Kafka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafka

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) SetKafka(v GithubComArgoprojLabsArgoDataflowApiV1alpha1KafkaSource)`

SetKafka sets Kafka field to given value.

### HasKafka

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) HasKafka() bool`

HasKafka returns a boolean if a field has been set.

### GetName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetry

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetRetry() GithubComArgoprojLabsArgoDataflowApiV1alpha1Backoff`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetRetryOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Backoff, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) SetRetry(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Backoff)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetS3

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetS3() GithubComArgoprojLabsArgoDataflowApiV1alpha1S3Source`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetS3Ok() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1S3Source, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) SetS3(v GithubComArgoprojLabsArgoDataflowApiV1alpha1S3Source)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetStan

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetStan() GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN`

GetStan returns the Stan field if non-nil, zero value otherwise.

### GetStanOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetStanOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN, bool)`

GetStanOk returns a tuple with the Stan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStan

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) SetStan(v GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN)`

SetStan sets Stan field to given value.

### HasStan

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) HasStan() bool`

HasStan returns a boolean if a field has been set.

### GetVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetVolume() GithubComArgoprojLabsArgoDataflowApiV1alpha1VolumeSource`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) GetVolumeOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1VolumeSource, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) SetVolume(v GithubComArgoprojLabsArgoDataflowApiV1alpha1VolumeSource)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Source) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


