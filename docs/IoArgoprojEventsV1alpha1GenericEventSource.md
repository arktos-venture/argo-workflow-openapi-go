# IoArgoprojEventsV1alpha1GenericEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Config** | Pointer to **string** |  | [optional] 
**Insecure** | Pointer to **bool** | Insecure determines the type of connection. | [optional] 
**JsonBody** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Url** | Pointer to **string** | URL of the gRPC server that implements the event source. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1GenericEventSource

`func NewIoArgoprojEventsV1alpha1GenericEventSource() *IoArgoprojEventsV1alpha1GenericEventSource`

NewIoArgoprojEventsV1alpha1GenericEventSource instantiates a new IoArgoprojEventsV1alpha1GenericEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1GenericEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1GenericEventSourceWithDefaults() *IoArgoprojEventsV1alpha1GenericEventSource`

NewIoArgoprojEventsV1alpha1GenericEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1GenericEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthSecret

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) GetAuthSecret() IoK8sApiCoreV1SecretKeySelector`

GetAuthSecret returns the AuthSecret field if non-nil, zero value otherwise.

### GetAuthSecretOk

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) GetAuthSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetAuthSecretOk returns a tuple with the AuthSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSecret

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) SetAuthSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetAuthSecret sets AuthSecret field to given value.

### HasAuthSecret

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) HasAuthSecret() bool`

HasAuthSecret returns a boolean if a field has been set.

### GetConfig

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetInsecure

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetJsonBody

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) GetJsonBody() bool`

GetJsonBody returns the JsonBody field if non-nil, zero value otherwise.

### GetJsonBodyOk

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) GetJsonBodyOk() (*bool, bool)`

GetJsonBodyOk returns a tuple with the JsonBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBody

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) SetJsonBody(v bool)`

SetJsonBody sets JsonBody field to given value.

### HasJsonBody

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) HasJsonBody() bool`

HasJsonBody returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoArgoprojEventsV1alpha1GenericEventSource) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


