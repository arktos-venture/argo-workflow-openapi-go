# IoArgoprojEventsV1alpha1GithubEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**ApiToken** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**DeleteHookOnFinish** | Pointer to **bool** |  | [optional] 
**Events** | Pointer to **[]string** |  | [optional] 
**GithubBaseURL** | Pointer to **string** |  | [optional] 
**GithubUploadURL** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Insecure** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Repositories** | Pointer to [**[]IoArgoprojEventsV1alpha1OwnedRepositories**](IoArgoprojEventsV1alpha1OwnedRepositories.md) |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**Webhook** | Pointer to [**IoArgoprojEventsV1alpha1WebhookContext**](IoArgoprojEventsV1alpha1WebhookContext.md) |  | [optional] 
**WebhookSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1GithubEventSource

`func NewIoArgoprojEventsV1alpha1GithubEventSource() *IoArgoprojEventsV1alpha1GithubEventSource`

NewIoArgoprojEventsV1alpha1GithubEventSource instantiates a new IoArgoprojEventsV1alpha1GithubEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1GithubEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1GithubEventSourceWithDefaults() *IoArgoprojEventsV1alpha1GithubEventSource`

NewIoArgoprojEventsV1alpha1GithubEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1GithubEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApiToken

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetApiToken() IoK8sApiCoreV1SecretKeySelector`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetApiTokenOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetApiToken(v IoK8sApiCoreV1SecretKeySelector)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### GetContentType

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDeleteHookOnFinish

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetDeleteHookOnFinish() bool`

GetDeleteHookOnFinish returns the DeleteHookOnFinish field if non-nil, zero value otherwise.

### GetDeleteHookOnFinishOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetDeleteHookOnFinishOk() (*bool, bool)`

GetDeleteHookOnFinishOk returns a tuple with the DeleteHookOnFinish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteHookOnFinish

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetDeleteHookOnFinish(v bool)`

SetDeleteHookOnFinish sets DeleteHookOnFinish field to given value.

### HasDeleteHookOnFinish

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasDeleteHookOnFinish() bool`

HasDeleteHookOnFinish returns a boolean if a field has been set.

### GetEvents

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetGithubBaseURL

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetGithubBaseURL() string`

GetGithubBaseURL returns the GithubBaseURL field if non-nil, zero value otherwise.

### GetGithubBaseURLOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetGithubBaseURLOk() (*string, bool)`

GetGithubBaseURLOk returns a tuple with the GithubBaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubBaseURL

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetGithubBaseURL(v string)`

SetGithubBaseURL sets GithubBaseURL field to given value.

### HasGithubBaseURL

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasGithubBaseURL() bool`

HasGithubBaseURL returns a boolean if a field has been set.

### GetGithubUploadURL

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetGithubUploadURL() string`

GetGithubUploadURL returns the GithubUploadURL field if non-nil, zero value otherwise.

### GetGithubUploadURLOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetGithubUploadURLOk() (*string, bool)`

GetGithubUploadURLOk returns a tuple with the GithubUploadURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUploadURL

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetGithubUploadURL(v string)`

SetGithubUploadURL sets GithubUploadURL field to given value.

### HasGithubUploadURL

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasGithubUploadURL() bool`

HasGithubUploadURL returns a boolean if a field has been set.

### GetId

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsecure

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOwner

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRepositories

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetRepositories() []IoArgoprojEventsV1alpha1OwnedRepositories`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetRepositoriesOk() (*[]IoArgoprojEventsV1alpha1OwnedRepositories, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetRepositories(v []IoArgoprojEventsV1alpha1OwnedRepositories)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetRepository

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetWebhook

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetWebhook() IoArgoprojEventsV1alpha1WebhookContext`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetWebhookOk() (*IoArgoprojEventsV1alpha1WebhookContext, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetWebhook(v IoArgoprojEventsV1alpha1WebhookContext)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetWebhookSecret() IoK8sApiCoreV1SecretKeySelector`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) GetWebhookSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) SetWebhookSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *IoArgoprojEventsV1alpha1GithubEventSource) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


