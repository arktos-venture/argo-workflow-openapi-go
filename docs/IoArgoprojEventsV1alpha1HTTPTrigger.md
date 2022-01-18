# IoArgoprojEventsV1alpha1HTTPTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicAuth** | Pointer to [**IoArgoprojEventsV1alpha1BasicAuth**](IoArgoprojEventsV1alpha1BasicAuth.md) |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) | Parameters is the list of key-value extracted from event&#39;s payload that are applied to the HTTP trigger resource. | [optional] 
**Payload** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) |  | [optional] 
**SecureHeaders** | Pointer to [**[]IoArgoprojEventsV1alpha1SecureHeader**](IoArgoprojEventsV1alpha1SecureHeader.md) |  | [optional] 
**Timeout** | Pointer to **string** |  | [optional] 
**Tls** | Pointer to [**IoArgoprojEventsV1alpha1TLSConfig**](IoArgoprojEventsV1alpha1TLSConfig.md) |  | [optional] 
**Url** | Pointer to **string** | URL refers to the URL to send HTTP request to. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1HTTPTrigger

`func NewIoArgoprojEventsV1alpha1HTTPTrigger() *IoArgoprojEventsV1alpha1HTTPTrigger`

NewIoArgoprojEventsV1alpha1HTTPTrigger instantiates a new IoArgoprojEventsV1alpha1HTTPTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1HTTPTriggerWithDefaults

`func NewIoArgoprojEventsV1alpha1HTTPTriggerWithDefaults() *IoArgoprojEventsV1alpha1HTTPTrigger`

NewIoArgoprojEventsV1alpha1HTTPTriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1HTTPTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicAuth

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetBasicAuth() IoArgoprojEventsV1alpha1BasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetBasicAuthOk() (*IoArgoprojEventsV1alpha1BasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) SetBasicAuth(v IoArgoprojEventsV1alpha1BasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetHeaders

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetParameters

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPayload

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetPayload() []IoArgoprojEventsV1alpha1TriggerParameter`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetPayloadOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) SetPayload(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSecureHeaders

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetSecureHeaders() []IoArgoprojEventsV1alpha1SecureHeader`

GetSecureHeaders returns the SecureHeaders field if non-nil, zero value otherwise.

### GetSecureHeadersOk

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetSecureHeadersOk() (*[]IoArgoprojEventsV1alpha1SecureHeader, bool)`

GetSecureHeadersOk returns a tuple with the SecureHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureHeaders

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) SetSecureHeaders(v []IoArgoprojEventsV1alpha1SecureHeader)`

SetSecureHeaders sets SecureHeaders field to given value.

### HasSecureHeaders

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) HasSecureHeaders() bool`

HasSecureHeaders returns a boolean if a field has been set.

### GetTimeout

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTls

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetTls() IoArgoprojEventsV1alpha1TLSConfig`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetTlsOk() (*IoArgoprojEventsV1alpha1TLSConfig, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) SetTls(v IoArgoprojEventsV1alpha1TLSConfig)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoArgoprojEventsV1alpha1HTTPTrigger) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


