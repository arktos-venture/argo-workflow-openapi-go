# IoArgoprojEventsV1alpha1NATSEventsSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**IoArgoprojEventsV1alpha1NATSAuth**](IoArgoprojEventsV1alpha1NATSAuth.md) |  | [optional] 
**ConnectionBackoff** | Pointer to [**IoArgoprojEventsV1alpha1Backoff**](IoArgoprojEventsV1alpha1Backoff.md) |  | [optional] 
**JsonBody** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Tls** | Pointer to [**IoArgoprojEventsV1alpha1TLSConfig**](IoArgoprojEventsV1alpha1TLSConfig.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1NATSEventsSource

`func NewIoArgoprojEventsV1alpha1NATSEventsSource() *IoArgoprojEventsV1alpha1NATSEventsSource`

NewIoArgoprojEventsV1alpha1NATSEventsSource instantiates a new IoArgoprojEventsV1alpha1NATSEventsSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1NATSEventsSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1NATSEventsSourceWithDefaults() *IoArgoprojEventsV1alpha1NATSEventsSource`

NewIoArgoprojEventsV1alpha1NATSEventsSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1NATSEventsSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetAuth() IoArgoprojEventsV1alpha1NATSAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetAuthOk() (*IoArgoprojEventsV1alpha1NATSAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) SetAuth(v IoArgoprojEventsV1alpha1NATSAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetConnectionBackoff() IoArgoprojEventsV1alpha1Backoff`

GetConnectionBackoff returns the ConnectionBackoff field if non-nil, zero value otherwise.

### GetConnectionBackoffOk

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetConnectionBackoffOk() (*IoArgoprojEventsV1alpha1Backoff, bool)`

GetConnectionBackoffOk returns a tuple with the ConnectionBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) SetConnectionBackoff(v IoArgoprojEventsV1alpha1Backoff)`

SetConnectionBackoff sets ConnectionBackoff field to given value.

### HasConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) HasConnectionBackoff() bool`

HasConnectionBackoff returns a boolean if a field has been set.

### GetJsonBody

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetJsonBody() bool`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetJsonBodyOk() (*bool, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) SetJsonBody(v bool)`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSubject

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTls

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetTls() IoArgoprojEventsV1alpha1TLSConfig`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetTlsOk() (*IoArgoprojEventsV1alpha1TLSConfig, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) SetTls(v IoArgoprojEventsV1alpha1TLSConfig)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoArgoprojEventsV1alpha1NATSEventsSource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


