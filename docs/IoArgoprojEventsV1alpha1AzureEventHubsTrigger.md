# IoArgoprojEventsV1alpha1AzureEventHubsTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** |  | [optional] 
**HubName** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) |  | [optional] 
**Payload** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) | Payload is the list of key-value extracted from an event payload to construct the request payload. | [optional] 
**SharedAccessKey** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**SharedAccessKeyName** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1AzureEventHubsTrigger

`func NewIoArgoprojEventsV1alpha1AzureEventHubsTrigger() *IoArgoprojEventsV1alpha1AzureEventHubsTrigger`

NewIoArgoprojEventsV1alpha1AzureEventHubsTrigger instantiates a new IoArgoprojEventsV1alpha1AzureEventHubsTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1AzureEventHubsTriggerWithDefaults

`func NewIoArgoprojEventsV1alpha1AzureEventHubsTriggerWithDefaults() *IoArgoprojEventsV1alpha1AzureEventHubsTrigger`

NewIoArgoprojEventsV1alpha1AzureEventHubsTriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1AzureEventHubsTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetHubName

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) GetHubName() string`

GetHubName returns the HubName field if non-nil, zero value otherwise.

### GetHubNameOk

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) GetHubNameOk() (*string, bool)`

GetHubNameOk returns a tuple with the HubName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubName

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) SetHubName(v string)`

SetHubName sets HubName field to given value.

### HasHubName

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) HasHubName() bool`

HasHubName returns a boolean if a field has been set.

### GetParameters

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPayload

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) GetPayload() []IoArgoprojEventsV1alpha1TriggerParameter`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) GetPayloadOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) SetPayload(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSharedAccessKey

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) GetSharedAccessKey() IoK8sApiCoreV1SecretKeySelector`

GetSharedAccessKey returns the SharedAccessKey field if non-nil, zero value otherwise.

### GetSharedAccessKeyOk

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) GetSharedAccessKeyOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSharedAccessKeyOk returns a tuple with the SharedAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAccessKey

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) SetSharedAccessKey(v IoK8sApiCoreV1SecretKeySelector)`

SetSharedAccessKey sets SharedAccessKey field to given value.

### HasSharedAccessKey

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) HasSharedAccessKey() bool`

HasSharedAccessKey returns a boolean if a field has been set.

### GetSharedAccessKeyName

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) GetSharedAccessKeyName() IoK8sApiCoreV1SecretKeySelector`

GetSharedAccessKeyName returns the SharedAccessKeyName field if non-nil, zero value otherwise.

### GetSharedAccessKeyNameOk

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) GetSharedAccessKeyNameOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSharedAccessKeyNameOk returns a tuple with the SharedAccessKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAccessKeyName

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) SetSharedAccessKeyName(v IoK8sApiCoreV1SecretKeySelector)`

SetSharedAccessKeyName sets SharedAccessKeyName field to given value.

### HasSharedAccessKeyName

`func (o *IoArgoprojEventsV1alpha1AzureEventHubsTrigger) HasSharedAccessKeyName() bool`

HasSharedAccessKeyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


