# IoArgoprojEventsV1alpha1NATSTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) |  | [optional] 
**Payload** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) |  | [optional] 
**Subject** | Pointer to **string** | Name of the subject to put message on. | [optional] 
**Tls** | Pointer to [**IoArgoprojEventsV1alpha1TLSConfig**](IoArgoprojEventsV1alpha1TLSConfig.md) |  | [optional] 
**Url** | Pointer to **string** | URL of the NATS cluster. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1NATSTrigger

`func NewIoArgoprojEventsV1alpha1NATSTrigger() *IoArgoprojEventsV1alpha1NATSTrigger`

NewIoArgoprojEventsV1alpha1NATSTrigger instantiates a new IoArgoprojEventsV1alpha1NATSTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1NATSTriggerWithDefaults

`func NewIoArgoprojEventsV1alpha1NATSTriggerWithDefaults() *IoArgoprojEventsV1alpha1NATSTrigger`

NewIoArgoprojEventsV1alpha1NATSTriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1NATSTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPayload

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) GetPayload() []IoArgoprojEventsV1alpha1TriggerParameter`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) GetPayloadOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) SetPayload(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSubject

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTls

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) GetTls() IoArgoprojEventsV1alpha1TLSConfig`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) GetTlsOk() (*IoArgoprojEventsV1alpha1TLSConfig, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) SetTls(v IoArgoprojEventsV1alpha1TLSConfig)`

SetTls sets Tls field to given value.

### HasTls

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IoArgoprojEventsV1alpha1NATSTrigger) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


