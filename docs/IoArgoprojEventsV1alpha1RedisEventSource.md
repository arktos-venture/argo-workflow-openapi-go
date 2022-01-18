# IoArgoprojEventsV1alpha1RedisEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to **[]string** |  | [optional] 
**Db** | Pointer to **int32** |  | [optional] 
**HostAddress** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Password** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Tls** | Pointer to [**IoArgoprojEventsV1alpha1TLSConfig**](IoArgoprojEventsV1alpha1TLSConfig.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1RedisEventSource

`func NewIoArgoprojEventsV1alpha1RedisEventSource() *IoArgoprojEventsV1alpha1RedisEventSource`

NewIoArgoprojEventsV1alpha1RedisEventSource instantiates a new IoArgoprojEventsV1alpha1RedisEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1RedisEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1RedisEventSourceWithDefaults() *IoArgoprojEventsV1alpha1RedisEventSource`

NewIoArgoprojEventsV1alpha1RedisEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1RedisEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetDb

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetDb() int32`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetDbOk() (*int32, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) SetDb(v int32)`

SetDb sets Db field to given value.

### HasDb

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) HasDb() bool`

HasDb returns a boolean if a field has been set.

### GetHostAddress

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetHostAddress() string`

GetHostAddress returns the HostAddress field if non-nil, zero value otherwise.

### GetHostAddressOk

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetHostAddressOk() (*string, bool)`

GetHostAddressOk returns a tuple with the HostAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAddress

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) SetHostAddress(v string)`

SetHostAddress sets HostAddress field to given value.

### HasHostAddress

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) HasHostAddress() bool`

HasHostAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNamespace

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPassword

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetPassword() IoK8sApiCoreV1SecretKeySelector`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetPasswordOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) SetPassword(v IoK8sApiCoreV1SecretKeySelector)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTls

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetTls() IoArgoprojEventsV1alpha1TLSConfig`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) GetTlsOk() (*IoArgoprojEventsV1alpha1TLSConfig, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) SetTls(v IoArgoprojEventsV1alpha1TLSConfig)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IoArgoprojEventsV1alpha1RedisEventSource) HasTls() bool`

HasTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


