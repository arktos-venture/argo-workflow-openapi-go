# IoArgoprojEventsV1alpha1GitlabEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**DeleteHookOnFinish** | Pointer to **bool** |  | [optional] 
**EnableSSLVerification** | Pointer to **bool** |  | [optional] 
**Events** | Pointer to **[]string** | Events are gitlab event to listen to. Refer https://github.com/xanzy/go-gitlab/blob/bf34eca5d13a9f4c3f501d8a97b8ac226d55e4d9/projects.go#L794. | [optional] 
**GitlabBaseURL** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**ProjectID** | Pointer to **string** |  | [optional] 
**Webhook** | Pointer to [**IoArgoprojEventsV1alpha1WebhookContext**](IoArgoprojEventsV1alpha1WebhookContext.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1GitlabEventSource

`func NewIoArgoprojEventsV1alpha1GitlabEventSource() *IoArgoprojEventsV1alpha1GitlabEventSource`

NewIoArgoprojEventsV1alpha1GitlabEventSource instantiates a new IoArgoprojEventsV1alpha1GitlabEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1GitlabEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1GitlabEventSourceWithDefaults() *IoArgoprojEventsV1alpha1GitlabEventSource`

NewIoArgoprojEventsV1alpha1GitlabEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1GitlabEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetAccessToken() IoK8sApiCoreV1SecretKeySelector`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetAccessTokenOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) SetAccessToken(v IoK8sApiCoreV1SecretKeySelector)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetDeleteHookOnFinish

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetDeleteHookOnFinish() bool`

GetDeleteHookOnFinish returns the DeleteHookOnFinish field if non-nil, zero value otherwise.

### GetDeleteHookOnFinishOk

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetDeleteHookOnFinishOk() (*bool, bool)`

GetDeleteHookOnFinishOk returns a tuple with the DeleteHookOnFinish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteHookOnFinish

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) SetDeleteHookOnFinish(v bool)`

SetDeleteHookOnFinish sets DeleteHookOnFinish field to given value.

### HasDeleteHookOnFinish

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) HasDeleteHookOnFinish() bool`

HasDeleteHookOnFinish returns a boolean if a field has been set.

### GetEnableSSLVerification

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetEnableSSLVerification() bool`

GetEnableSSLVerification returns the EnableSSLVerification field if non-nil, zero value otherwise.

### GetEnableSSLVerificationOk

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetEnableSSLVerificationOk() (*bool, bool)`

GetEnableSSLVerificationOk returns a tuple with the EnableSSLVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSSLVerification

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) SetEnableSSLVerification(v bool)`

SetEnableSSLVerification sets EnableSSLVerification field to given value.

### HasEnableSSLVerification

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) HasEnableSSLVerification() bool`

HasEnableSSLVerification returns a boolean if a field has been set.

### GetEvents

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetGitlabBaseURL

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetGitlabBaseURL() string`

GetGitlabBaseURL returns the GitlabBaseURL field if non-nil, zero value otherwise.

### GetGitlabBaseURLOk

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetGitlabBaseURLOk() (*string, bool)`

GetGitlabBaseURLOk returns a tuple with the GitlabBaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabBaseURL

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) SetGitlabBaseURL(v string)`

SetGitlabBaseURL sets GitlabBaseURL field to given value.

### HasGitlabBaseURL

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) HasGitlabBaseURL() bool`

HasGitlabBaseURL returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProjectID

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### GetWebhook

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetWebhook() IoArgoprojEventsV1alpha1WebhookContext`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) GetWebhookOk() (*IoArgoprojEventsV1alpha1WebhookContext, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) SetWebhook(v IoArgoprojEventsV1alpha1WebhookContext)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *IoArgoprojEventsV1alpha1GitlabEventSource) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


