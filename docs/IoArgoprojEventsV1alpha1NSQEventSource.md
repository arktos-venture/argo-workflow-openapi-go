# IoArgoprojEventsV1alpha1NSQEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** |  | [optional] 
**ConnectionBackoff** | Pointer to [**IoArgoprojEventsV1alpha1Backoff**](IoArgoprojEventsV1alpha1Backoff.md) |  | [optional] 
**HostAddress** | Pointer to **string** |  | [optional] 
**JsonBody** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Tls** | Pointer to [**IoArgoprojEventsV1alpha1TLSConfig**](IoArgoprojEventsV1alpha1TLSConfig.md) |  | [optional] 
**Topic** | Pointer to **string** | Topic to subscribe to. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1NSQEventSource

`func NewIoArgoprojEventsV1alpha1NSQEventSource() *IoArgoprojEventsV1alpha1NSQEventSource`

NewIoArgoprojEventsV1alpha1NSQEventSource instantiates a new IoArgoprojEventsV1alpha1NSQEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1NSQEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1NSQEventSourceWithDefaults() *IoArgoprojEventsV1alpha1NSQEventSource`

NewIoArgoprojEventsV1alpha1NSQEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1NSQEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetConnectionBackoff() IoArgoprojEventsV1alpha1Backoff`

GetConnectionBackoff returns the ConnectionBackoff field if non-nil, zero value otherwise.

### GetConnectionBackoffOk

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetConnectionBackoffOk() (*IoArgoprojEventsV1alpha1Backoff, bool)`

GetConnectionBackoffOk returns a tuple with the ConnectionBackoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) SetConnectionBackoff(v IoArgoprojEventsV1alpha1Backoff)`

SetConnectionBackoff sets ConnectionBackoff field to given value.

### HasConnectionBackoff

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) HasConnectionBackoff() bool`

HasConnectionBackoff returns a boolean if a field has been set.

### GetHostAddress

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetHostAddress() string`

GetHostAddress returns the HostAddress field if non-nil, zero value otherwise.

### GetHostAddressOk

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetHostAddressOk() (*string, bool)`

GetHostAddressOk returns a tuple with the HostAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAddress

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) SetHostAddress(v string)`

SetHostAddress sets HostAddress field to given value.

### HasHostAddress

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) HasHostAddress() bool`

HasHostAddress returns a boolean if a field has been set.

### GetJsonBody

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetJsonBody() bool`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetJsonBodyOk() (*bool, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) SetJsonBody(v bool)`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTls

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetTls() IoArgoprojEventsV1alpha1TLSConfig`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetTlsOk() (*IoArgoprojEventsV1alpha1TLSConfig, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) SetTls(v IoArgoprojEventsV1alpha1TLSConfig)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetTopic

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *IoArgoprojEventsV1alpha1NSQEventSource) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


