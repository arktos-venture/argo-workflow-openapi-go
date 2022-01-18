# IoArgoprojEventsV1alpha1PulsarEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionBackoff** | Pointer to [**IoArgoprojEventsV1alpha1Backoff**](IoArgoprojEventsV1alpha1Backoff.md) |  | [optional] 
**JsonBody** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Tls** | Pointer to [**IoArgoprojEventsV1alpha1TLSConfig**](IoArgoprojEventsV1alpha1TLSConfig.md) |  | [optional] 
**TlsAllowInsecureConnection** | Pointer to **bool** |  | [optional] 
**TlsTrustCertsSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**TlsValidateHostname** | Pointer to **bool** |  | [optional] 
**Topics** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1PulsarEventSource

`func NewIoArgoprojEventsV1alpha1PulsarEventSource() *IoArgoprojEventsV1alpha1PulsarEventSource`

NewIoArgoprojEventsV1alpha1PulsarEventSource instantiates a new IoArgoprojEventsV1alpha1PulsarEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1PulsarEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1PulsarEventSourceWithDefaults() *IoArgoprojEventsV1alpha1PulsarEventSource`

NewIoArgoprojEventsV1alpha1PulsarEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1PulsarEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetConnectionBackoff() IoArgoprojEventsV1alpha1Backoff`

GetConnectionBackoff returns the ConnectionBackoff field if non-nil, zero value otherwise.

### GetConnectionBackoffOk

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetConnectionBackoffOk() (*IoArgoprojEventsV1alpha1Backoff, bool)`

GetConnectionBackoffOk returns a tuple with the ConnectionBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) SetConnectionBackoff(v IoArgoprojEventsV1alpha1Backoff)`

SetConnectionBackoff sets ConnectionBackoff field to given value.

### HasConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) HasConnectionBackoff() bool`

HasConnectionBackoff returns a boolean if a field has been set.

### GetJsonBody

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetJsonBody() bool`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetJsonBodyOk() (*bool, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) SetJsonBody(v bool)`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTls

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetTls() IoArgoprojEventsV1alpha1TLSConfig`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetTlsOk() (*IoArgoprojEventsV1alpha1TLSConfig, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) SetTls(v IoArgoprojEventsV1alpha1TLSConfig)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetTlsAllowInsecureConnection

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetTlsAllowInsecureConnection() bool`

GetTlsAllowInsecureConnection returns the TlsAllowInsecureConnection field if non-nil, zero value otherwise.

### GetTlsAllowInsecureConnectionOk

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetTlsAllowInsecureConnectionOk() (*bool, bool)`

GetTlsAllowInsecureConnectionOk returns a tuple with the TlsAllowInsecureConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAllowInsecureConnection

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) SetTlsAllowInsecureConnection(v bool)`

SetTlsAllowInsecureConnection sets TlsAllowInsecureConnection field to given value.

### HasTlsAllowInsecureConnection

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) HasTlsAllowInsecureConnection() bool`

HasTlsAllowInsecureConnection returns a boolean if a field has been set.

### GetTlsTrustCertsSecret

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetTlsTrustCertsSecret() IoK8sApiCoreV1SecretKeySelector`

GetTlsTrustCertsSecret returns the TlsTrustCertsSecret field if non-nil, zero value otherwise.

### GetTlsTrustCertsSecretOk

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetTlsTrustCertsSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetTlsTrustCertsSecretOk returns a tuple with the TlsTrustCertsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTrustCertsSecret

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) SetTlsTrustCertsSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetTlsTrustCertsSecret sets TlsTrustCertsSecret field to given value.

### HasTlsTrustCertsSecret

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) HasTlsTrustCertsSecret() bool`

HasTlsTrustCertsSecret returns a boolean if a field has been set.

### GetTlsValidateHostname

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetTlsValidateHostname() bool`

GetTlsValidateHostname returns the TlsValidateHostname field if non-nil, zero value otherwise.

### GetTlsValidateHostnameOk

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetTlsValidateHostnameOk() (*bool, bool)`

GetTlsValidateHostnameOk returns a tuple with the TlsValidateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsValidateHostname

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) SetTlsValidateHostname(v bool)`

SetTlsValidateHostname sets TlsValidateHostname field to given value.

### HasTlsValidateHostname

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) HasTlsValidateHostname() bool`

HasTlsValidateHostname returns a boolean if a field has been set.

### GetTopics

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetType

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoArgoprojEventsV1alpha1PulsarEventSource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


