# IoArgoprojEventsV1alpha1TLSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCertPath** | Pointer to **string** |  | [optional] 
**CaCertSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**ClientCertPath** | Pointer to **string** |  | [optional] 
**ClientCertSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**ClientKeyPath** | Pointer to **string** |  | [optional] 
**ClientKeySecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1TLSConfig

`func NewIoArgoprojEventsV1alpha1TLSConfig() *IoArgoprojEventsV1alpha1TLSConfig`

NewIoArgoprojEventsV1alpha1TLSConfig instantiates a new IoArgoprojEventsV1alpha1TLSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1TLSConfigWithDefaults

`func NewIoArgoprojEventsV1alpha1TLSConfigWithDefaults() *IoArgoprojEventsV1alpha1TLSConfig`

NewIoArgoprojEventsV1alpha1TLSConfigWithDefaults instantiates a new IoArgoprojEventsV1alpha1TLSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCertPath

`func (o *IoArgoprojEventsV1alpha1TLSConfig) GetCaCertPath() string`

GetCaCertPath returns the CaCertPath field if non-nil, zero value otherwise.

### GetCaCertPathOk

`func (o *IoArgoprojEventsV1alpha1TLSConfig) GetCaCertPathOk() (*string, bool)`

GetCaCertPathOk returns a tuple with the CaCertPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertPath

`func (o *IoArgoprojEventsV1alpha1TLSConfig) SetCaCertPath(v string)`

SetCaCertPath sets CaCertPath field to given value.

### HasCaCertPath

`func (o *IoArgoprojEventsV1alpha1TLSConfig) HasCaCertPath() bool`

HasCaCertPath returns a boolean if a field has been set.

### GetCaCertSecret

`func (o *IoArgoprojEventsV1alpha1TLSConfig) GetCaCertSecret() IoK8sApiCoreV1SecretKeySelector`

GetCaCertSecret returns the CaCertSecret field if non-nil, zero value otherwise.

### GetCaCertSecretOk

`func (o *IoArgoprojEventsV1alpha1TLSConfig) GetCaCertSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetCaCertSecretOk returns a tuple with the CaCertSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertSecret

`func (o *IoArgoprojEventsV1alpha1TLSConfig) SetCaCertSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetCaCertSecret sets CaCertSecret field to given value.

### HasCaCertSecret

`func (o *IoArgoprojEventsV1alpha1TLSConfig) HasCaCertSecret() bool`

HasCaCertSecret returns a boolean if a field has been set.

### GetClientCertPath

`func (o *IoArgoprojEventsV1alpha1TLSConfig) GetClientCertPath() string`

GetClientCertPath returns the ClientCertPath field if non-nil, zero value otherwise.

### GetClientCertPathOk

`func (o *IoArgoprojEventsV1alpha1TLSConfig) GetClientCertPathOk() (*string, bool)`

GetClientCertPathOk returns a tuple with the ClientCertPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertPath

`func (o *IoArgoprojEventsV1alpha1TLSConfig) SetClientCertPath(v string)`

SetClientCertPath sets ClientCertPath field to given value.

### HasClientCertPath

`func (o *IoArgoprojEventsV1alpha1TLSConfig) HasClientCertPath() bool`

HasClientCertPath returns a boolean if a field has been set.

### GetClientCertSecret

`func (o *IoArgoprojEventsV1alpha1TLSConfig) GetClientCertSecret() IoK8sApiCoreV1SecretKeySelector`

GetClientCertSecret returns the ClientCertSecret field if non-nil, zero value otherwise.

### GetClientCertSecretOk

`func (o *IoArgoprojEventsV1alpha1TLSConfig) GetClientCertSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetClientCertSecretOk returns a tuple with the ClientCertSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertSecret

`func (o *IoArgoprojEventsV1alpha1TLSConfig) SetClientCertSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetClientCertSecret sets ClientCertSecret field to given value.

### HasClientCertSecret

`func (o *IoArgoprojEventsV1alpha1TLSConfig) HasClientCertSecret() bool`

HasClientCertSecret returns a boolean if a field has been set.

### GetClientKeyPath

`func (o *IoArgoprojEventsV1alpha1TLSConfig) GetClientKeyPath() string`

GetClientKeyPath returns the ClientKeyPath field if non-nil, zero value otherwise.

### GetClientKeyPathOk

`func (o *IoArgoprojEventsV1alpha1TLSConfig) GetClientKeyPathOk() (*string, bool)`

GetClientKeyPathOk returns a tuple with the ClientKeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKeyPath

`func (o *IoArgoprojEventsV1alpha1TLSConfig) SetClientKeyPath(v string)`

SetClientKeyPath sets ClientKeyPath field to given value.

### HasClientKeyPath

`func (o *IoArgoprojEventsV1alpha1TLSConfig) HasClientKeyPath() bool`

HasClientKeyPath returns a boolean if a field has been set.

### GetClientKeySecret

`func (o *IoArgoprojEventsV1alpha1TLSConfig) GetClientKeySecret() IoK8sApiCoreV1SecretKeySelector`

GetClientKeySecret returns the ClientKeySecret field if non-nil, zero value otherwise.

### GetClientKeySecretOk

`func (o *IoArgoprojEventsV1alpha1TLSConfig) GetClientKeySecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetClientKeySecretOk returns a tuple with the ClientKeySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKeySecret

`func (o *IoArgoprojEventsV1alpha1TLSConfig) SetClientKeySecret(v IoK8sApiCoreV1SecretKeySelector)`

SetClientKeySecret sets ClientKeySecret field to given value.

### HasClientKeySecret

`func (o *IoArgoprojEventsV1alpha1TLSConfig) HasClientKeySecret() bool`

HasClientKeySecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


