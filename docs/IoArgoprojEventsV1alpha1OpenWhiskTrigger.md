# IoArgoprojEventsV1alpha1OpenWhiskTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionName** | Pointer to **string** | Name of the action/function. | [optional] 
**AuthToken** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Host** | Pointer to **string** | Host URL of the OpenWhisk. | [optional] 
**Namespace** | Pointer to **string** | Namespace for the action. Defaults to \&quot;_\&quot;. +optional. | [optional] 
**Parameters** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) |  | [optional] 
**Payload** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) | Payload is the list of key-value extracted from an event payload to construct the request payload. | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1OpenWhiskTrigger

`func NewIoArgoprojEventsV1alpha1OpenWhiskTrigger() *IoArgoprojEventsV1alpha1OpenWhiskTrigger`

NewIoArgoprojEventsV1alpha1OpenWhiskTrigger instantiates a new IoArgoprojEventsV1alpha1OpenWhiskTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1OpenWhiskTriggerWithDefaults

`func NewIoArgoprojEventsV1alpha1OpenWhiskTriggerWithDefaults() *IoArgoprojEventsV1alpha1OpenWhiskTrigger`

NewIoArgoprojEventsV1alpha1OpenWhiskTriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1OpenWhiskTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionName

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetActionName() string`

GetActionName returns the ActionName field if non-nil, zero value otherwise.

### GetActionNameOk

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetActionNameOk() (*string, bool)`

GetActionNameOk returns a tuple with the ActionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionName

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) SetActionName(v string)`

SetActionName sets ActionName field to given value.

### HasActionName

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) HasActionName() bool`

HasActionName returns a boolean if a field has been set.

### GetAuthToken

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetAuthToken() IoK8sApiCoreV1SecretKeySelector`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetAuthTokenOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) SetAuthToken(v IoK8sApiCoreV1SecretKeySelector)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetHost

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetNamespace

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetParameters

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPayload

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetPayload() []IoArgoprojEventsV1alpha1TriggerParameter`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetPayloadOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) SetPayload(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetVersion

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IoArgoprojEventsV1alpha1OpenWhiskTrigger) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


