# IoArgoprojEventsV1alpha1StorageGridEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiURL** | Pointer to **string** | APIURL is the url of the storagegrid api. | [optional] 
**AuthToken** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Bucket** | Pointer to **string** | Name of the bucket to register notifications for. | [optional] 
**Events** | Pointer to **[]string** |  | [optional] 
**Filter** | Pointer to [**IoArgoprojEventsV1alpha1StorageGridFilter**](IoArgoprojEventsV1alpha1StorageGridFilter.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**TopicArn** | Pointer to **string** |  | [optional] 
**Webhook** | Pointer to [**IoArgoprojEventsV1alpha1WebhookContext**](IoArgoprojEventsV1alpha1WebhookContext.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1StorageGridEventSource

`func NewIoArgoprojEventsV1alpha1StorageGridEventSource() *IoArgoprojEventsV1alpha1StorageGridEventSource`

NewIoArgoprojEventsV1alpha1StorageGridEventSource instantiates a new IoArgoprojEventsV1alpha1StorageGridEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1StorageGridEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1StorageGridEventSourceWithDefaults() *IoArgoprojEventsV1alpha1StorageGridEventSource`

NewIoArgoprojEventsV1alpha1StorageGridEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1StorageGridEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiURL

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.

### HasApiURL

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasApiURL() bool`

HasApiURL returns a boolean if a field has been set.

### GetAuthToken

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetAuthToken() IoK8sApiCoreV1SecretKeySelector`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetAuthTokenOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetAuthToken(v IoK8sApiCoreV1SecretKeySelector)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetBucket

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetEvents

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetFilter

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetFilter() IoArgoprojEventsV1alpha1StorageGridFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetFilterOk() (*IoArgoprojEventsV1alpha1StorageGridFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetFilter(v IoArgoprojEventsV1alpha1StorageGridFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRegion

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetTopicArn

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetTopicArn() string`

GetTopicArn returns the TopicArn field if non-nil, zero value otherwise.

### GetTopicArnOk

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetTopicArnOk() (*string, bool)`

GetTopicArnOk returns a tuple with the TopicArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicArn

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetTopicArn(v string)`

SetTopicArn sets TopicArn field to given value.

### HasTopicArn

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasTopicArn() bool`

HasTopicArn returns a boolean if a field has been set.

### GetWebhook

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetWebhook() IoArgoprojEventsV1alpha1WebhookContext`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) GetWebhookOk() (*IoArgoprojEventsV1alpha1WebhookContext, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) SetWebhook(v IoArgoprojEventsV1alpha1WebhookContext)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *IoArgoprojEventsV1alpha1StorageGridEventSource) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


