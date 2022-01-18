# IoArgoprojEventsV1alpha1PubSubEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**CredentialsFile** | Pointer to **string** |  | [optional] 
**DeleteSubscriptionOnFinish** | Pointer to **bool** |  | [optional] 
**JsonBody** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**ProjectID** | Pointer to **string** |  | [optional] 
**SubscriptionID** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**TopicProjectID** | Pointer to **string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1PubSubEventSource

`func NewIoArgoprojEventsV1alpha1PubSubEventSource() *IoArgoprojEventsV1alpha1PubSubEventSource`

NewIoArgoprojEventsV1alpha1PubSubEventSource instantiates a new IoArgoprojEventsV1alpha1PubSubEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1PubSubEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1PubSubEventSourceWithDefaults() *IoArgoprojEventsV1alpha1PubSubEventSource`

NewIoArgoprojEventsV1alpha1PubSubEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1PubSubEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialSecret

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetCredentialSecret() IoK8sApiCoreV1SecretKeySelector`

GetCredentialSecret returns the CredentialSecret field if non-nil, zero value otherwise.

### GetCredentialSecretOk

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetCredentialSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetCredentialSecretOk returns a tuple with the CredentialSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialSecret

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) SetCredentialSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetCredentialSecret sets CredentialSecret field to given value.

### HasCredentialSecret

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) HasCredentialSecret() bool`

HasCredentialSecret returns a boolean if a field has been set.

### GetCredentialsFile

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetCredentialsFile() string`

GetCredentialsFile returns the CredentialsFile field if non-nil, zero value otherwise.

### GetCredentialsFileOk

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetCredentialsFileOk() (*string, bool)`

GetCredentialsFileOk returns a tuple with the CredentialsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsFile

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) SetCredentialsFile(v string)`

SetCredentialsFile sets CredentialsFile field to given value.

### HasCredentialsFile

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) HasCredentialsFile() bool`

HasCredentialsFile returns a boolean if a field has been set.

### GetDeleteSubscriptionOnFinish

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetDeleteSubscriptionOnFinish() bool`

GetDeleteSubscriptionOnFinish returns the DeleteSubscriptionOnFinish field if non-nil, zero value otherwise.

### GetDeleteSubscriptionOnFinishOk

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetDeleteSubscriptionOnFinishOk() (*bool, bool)`

GetDeleteSubscriptionOnFinishOk returns a tuple with the DeleteSubscriptionOnFinish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSubscriptionOnFinish

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) SetDeleteSubscriptionOnFinish(v bool)`

SetDeleteSubscriptionOnFinish sets DeleteSubscriptionOnFinish field to given value.

### HasDeleteSubscriptionOnFinish

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) HasDeleteSubscriptionOnFinish() bool`

HasDeleteSubscriptionOnFinish returns a boolean if a field has been set.

### GetJsonBody

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetJsonBody() bool`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetJsonBodyOk() (*bool, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) SetJsonBody(v bool)`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProjectID

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### GetSubscriptionID

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetSubscriptionID() string`

GetSubscriptionID returns the SubscriptionID field if non-nil, zero value otherwise.

### GetSubscriptionIDOk

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetSubscriptionIDOk() (*string, bool)`

GetSubscriptionIDOk returns a tuple with the SubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionID

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) SetSubscriptionID(v string)`

SetSubscriptionID sets SubscriptionID field to given value.

### HasSubscriptionID

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) HasSubscriptionID() bool`

HasSubscriptionID returns a boolean if a field has been set.

### GetTopic

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetTopicProjectID

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetTopicProjectID() string`

GetTopicProjectID returns the TopicProjectID field if non-nil, zero value otherwise.

### GetTopicProjectIDOk

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) GetTopicProjectIDOk() (*string, bool)`

GetTopicProjectIDOk returns a tuple with the TopicProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicProjectID

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) SetTopicProjectID(v string)`

SetTopicProjectID sets TopicProjectID field to given value.

### HasTopicProjectID

`func (o *IoArgoprojEventsV1alpha1PubSubEventSource) HasTopicProjectID() bool`

HasTopicProjectID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


