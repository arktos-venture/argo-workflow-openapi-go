# IoArgoprojEventsV1alpha1CustomTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertFilePath** | Pointer to **string** |  | [optional] 
**CertSecret** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Parameters** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) | Parameters is the list of parameters that is applied to resolved custom trigger trigger object. | [optional] 
**Payload** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) | Payload is the list of key-value extracted from an event payload to construct the request payload. | [optional] 
**Secure** | Pointer to **bool** |  | [optional] 
**ServerNameOverride** | Pointer to **string** | ServerNameOverride for the secure connection between sensor and custom trigger gRPC server. | [optional] 
**ServerURL** | Pointer to **string** |  | [optional] 
**Spec** | Pointer to **map[string]string** | Spec is the custom trigger resource specification that custom trigger gRPC server knows how to interpret. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1CustomTrigger

`func NewIoArgoprojEventsV1alpha1CustomTrigger() *IoArgoprojEventsV1alpha1CustomTrigger`

NewIoArgoprojEventsV1alpha1CustomTrigger instantiates a new IoArgoprojEventsV1alpha1CustomTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1CustomTriggerWithDefaults

`func NewIoArgoprojEventsV1alpha1CustomTriggerWithDefaults() *IoArgoprojEventsV1alpha1CustomTrigger`

NewIoArgoprojEventsV1alpha1CustomTriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1CustomTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertFilePath

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetCertFilePath() string`

GetCertFilePath returns the CertFilePath field if non-nil, zero value otherwise.

### GetCertFilePathOk

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetCertFilePathOk() (*string, bool)`

GetCertFilePathOk returns a tuple with the CertFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFilePath

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetCertFilePath(v string)`

SetCertFilePath sets CertFilePath field to given value.

### HasCertFilePath

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasCertFilePath() bool`

HasCertFilePath returns a boolean if a field has been set.

### GetCertSecret

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetCertSecret() IoK8sApiCoreV1SecretKeySelector`

GetCertSecret returns the CertSecret field if non-nil, zero value otherwise.

### GetCertSecretOk

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetCertSecretOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetCertSecretOk returns a tuple with the CertSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertSecret

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetCertSecret(v IoK8sApiCoreV1SecretKeySelector)`

SetCertSecret sets CertSecret field to given value.

### HasCertSecret

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasCertSecret() bool`

HasCertSecret returns a boolean if a field has been set.

### GetParameters

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPayload

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetPayload() []IoArgoprojEventsV1alpha1TriggerParameter`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetPayloadOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetPayload(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSecure

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetServerNameOverride

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetServerNameOverride() string`

GetServerNameOverride returns the ServerNameOverride field if non-nil, zero value otherwise.

### GetServerNameOverrideOk

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetServerNameOverrideOk() (*string, bool)`

GetServerNameOverrideOk returns a tuple with the ServerNameOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNameOverride

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetServerNameOverride(v string)`

SetServerNameOverride sets ServerNameOverride field to given value.

### HasServerNameOverride

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasServerNameOverride() bool`

HasServerNameOverride returns a boolean if a field has been set.

### GetServerURL

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetServerURL() string`

GetServerURL returns the ServerURL field if non-nil, zero value otherwise.

### GetServerURLOk

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetServerURLOk() (*string, bool)`

GetServerURLOk returns a tuple with the ServerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerURL

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetServerURL(v string)`

SetServerURL sets ServerURL field to given value.

### HasServerURL

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasServerURL() bool`

HasServerURL returns a boolean if a field has been set.

### GetSpec

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetSpec() map[string]string`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) GetSpecOk() (*map[string]string, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) SetSpec(v map[string]string)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *IoArgoprojEventsV1alpha1CustomTrigger) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


