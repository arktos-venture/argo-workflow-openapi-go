# IoArgoprojEventsV1alpha1WebhookContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** | Port on which HTTP server is listening for incoming events. | [optional] 
**ServerCertPath** | Pointer to **string** | DeprecatedServerCertPath refers the file that contains the cert. | [optional] 
**ServerCertSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**ServerKeyPath** | Pointer to **string** |  | [optional] 
**ServerKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Url** | Pointer to **string** | URL is the url of the server. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1WebhookContext

`func NewIoArgoprojEventsV1alpha1WebhookContext() *IoArgoprojEventsV1alpha1WebhookContext`

NewIoArgoprojEventsV1alpha1WebhookContext instantiates a new IoArgoprojEventsV1alpha1WebhookContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1WebhookContextWithDefaults

`func NewIoArgoprojEventsV1alpha1WebhookContextWithDefaults() *IoArgoprojEventsV1alpha1WebhookContext`

NewIoArgoprojEventsV1alpha1WebhookContextWithDefaults instantiates a new IoArgoprojEventsV1alpha1WebhookContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthSecret

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetAuthSecret() IoK8sApiCoreV1SecretKeySelector`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetAuthSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *IoArgoprojEventsV1alpha1WebhookContext) SetAuthSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *IoArgoprojEventsV1alpha1WebhookContext) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### GetEndpoint

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *IoArgoprojEventsV1alpha1WebhookContext) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *IoArgoprojEventsV1alpha1WebhookContext) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1WebhookContext) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1WebhookContext) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMethod

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *IoArgoprojEventsV1alpha1WebhookContext) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *IoArgoprojEventsV1alpha1WebhookContext) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPort

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IoArgoprojEventsV1alpha1WebhookContext) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *IoArgoprojEventsV1alpha1WebhookContext) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServerCertPath

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerCertPath() string`

GetServerCertPath returns the ServerCertPath field if non-nil, zero value otherwise.

### GetServerCertPathOk

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerCertPathOk() (*string, bool)`

GetServerCertPathOk returns a tuple with the ServerCertPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertPath

`func (o *IoArgoprojEventsV1alpha1WebhookContext) SetServerCertPath(v string)`

SetServerCertPath sets ServerCertPath field to given value.

### HasServerCertPath

`func (o *IoArgoprojEventsV1alpha1WebhookContext) HasServerCertPath() bool`

HasServerCertPath returns a boolean if a field has been set.

### GetServerCertSecret

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerCertSecret() IoK8sApiCoreV1SecretKeySelector`

GetServerCertSecret returns the ServerCertSecret field if non-nil, zero value otherwise.

### GetServerCertSecretOk

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerCertSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetServerCertSecretOk returns a tuple with the ServerCertSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertSecret

`func (o *IoArgoprojEventsV1alpha1WebhookContext) SetServerCertSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetServerCertSecret sets ServerCertSecret field to given value.

### HasServerCertSecret

`func (o *IoArgoprojEventsV1alpha1WebhookContext) HasServerCertSecret() bool`

HasServerCertSecret returns a boolean if a field has been set.

### GetServerKeyPath

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerKeyPath() string`

GetServerKeyPath returns the ServerKeyPath field if non-nil, zero value otherwise.

### GetServerKeyPathOk

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerKeyPathOk() (*string, bool)`

GetServerKeyPathOk returns a tuple with the ServerKeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKeyPath

`func (o *IoArgoprojEventsV1alpha1WebhookContext) SetServerKeyPath(v string)`

SetServerKeyPath sets ServerKeyPath field to given value.

### HasServerKeyPath

`func (o *IoArgoprojEventsV1alpha1WebhookContext) HasServerKeyPath() bool`

HasServerKeyPath returns a boolean if a field has been set.

### GetServerKeySecret

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetServerKeySecret returns the ServerKeySecret field if non-nil, zero value otherwise.

### GetServerKeySecretOk

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetServerKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetServerKeySecretOk returns a tuple with the ServerKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKeySecret

`func (o *IoArgoprojEventsV1alpha1WebhookContext) SetServerKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetServerKeySecret sets ServerKeySecret field to given value.

### HasServerKeySecret

`func (o *IoArgoprojEventsV1alpha1WebhookContext) HasServerKeySecret() bool`

HasServerKeySecret returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojEventsV1alpha1WebhookContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojEventsV1alpha1WebhookContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoArgoprojEventsV1alpha1WebhookContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


