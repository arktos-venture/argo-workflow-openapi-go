# IoArgoprojEventsV1alpha1SQSEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**JsonBody** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Queue** | Pointer to **string** |  | [optional] 
**QueueAccountId** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RoleARN** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**WaitTimeSeconds** | Pointer to **string** | WaitTimeSeconds is The duration (in seconds) for which the call waits for a message to arrive in the queue before returning. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1SQSEventSource

`func NewIoArgoprojEventsV1alpha1SQSEventSource() *IoArgoprojEventsV1alpha1SQSEventSource`

NewIoArgoprojEventsV1alpha1SQSEventSource instantiates a new IoArgoprojEventsV1alpha1SQSEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1SQSEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1SQSEventSourceWithDefaults() *IoArgoprojEventsV1alpha1SQSEventSource`

NewIoArgoprojEventsV1alpha1SQSEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1SQSEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetAccessKey() IoK8sApiCoreV1SecretKeySelector`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetAccessKeyOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) SetAccessKey(v IoK8sApiCoreV1SecretKeySelector)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetJsonBody

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetJsonBody() bool`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetJsonBodyOk() (*bool, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) SetJsonBody(v bool)`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQueue

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetQueueAccountId

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetQueueAccountId() string`

GetQueueAccountId returns the QueueAccountId field if non-nil, zero value otherwise.

### GetQueueAccountIdOk

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetQueueAccountIdOk() (*string, bool)`

GetQueueAccountIdOk returns a tuple with the QueueAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueAccountId

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) SetQueueAccountId(v string)`

SetQueueAccountId sets QueueAccountId field to given value.

### HasQueueAccountId

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) HasQueueAccountId() bool`

HasQueueAccountId returns a boolean if a field has been set.

### GetRegion

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoleARN

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetRoleARN() string`

GetRoleARN returns the RoleARN field if non-nil, zero value otherwise.

### GetRoleARNOk

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetRoleARNOk() (*string, bool)`

GetRoleARNOk returns a tuple with the RoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleARN

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) SetRoleARN(v string)`

SetRoleARN sets RoleARN field to given value.

### HasRoleARN

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) HasRoleARN() bool`

HasRoleARN returns a boolean if a field has been set.

### GetSecretKey

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetSecretKey() IoK8sApiCoreV1SecretKeySelector`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetSecretKeyOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) SetSecretKey(v IoK8sApiCoreV1SecretKeySelector)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetWaitTimeSeconds

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetWaitTimeSeconds() string`

GetWaitTimeSeconds returns the WaitTimeSeconds field if non-nil, zero value otherwise.

### GetWaitTimeSecondsOk

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) GetWaitTimeSecondsOk() (*string, bool)`

GetWaitTimeSecondsOk returns a tuple with the WaitTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimeSeconds

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) SetWaitTimeSeconds(v string)`

SetWaitTimeSeconds sets WaitTimeSeconds field to given value.

### HasWaitTimeSeconds

`func (o *IoArgoprojEventsV1alpha1SQSEventSource) HasWaitTimeSeconds() bool`

HasWaitTimeSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


