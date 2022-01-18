# IoArgoprojEventsV1alpha1SlackTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) |  | [optional] 
**SlackToken** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1SlackTrigger

`func NewIoArgoprojEventsV1alpha1SlackTrigger() *IoArgoprojEventsV1alpha1SlackTrigger`

NewIoArgoprojEventsV1alpha1SlackTrigger instantiates a new IoArgoprojEventsV1alpha1SlackTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1SlackTriggerWithDefaults

`func NewIoArgoprojEventsV1alpha1SlackTriggerWithDefaults() *IoArgoprojEventsV1alpha1SlackTrigger`

NewIoArgoprojEventsV1alpha1SlackTriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1SlackTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetMessage

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetParameters

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSlackToken

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) GetSlackToken() IoK8sApiCoreV1SecretKeySelector`

GetSlackToken returns the SlackToken field if non-nil, zero value otherwise.

### GetSlackTokenOk

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) GetSlackTokenOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSlackTokenOk returns a tuple with the SlackToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackToken

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) SetSlackToken(v IoK8sApiCoreV1SecretKeySelector)`

SetSlackToken sets SlackToken field to given value.

### HasSlackToken

`func (o *IoArgoprojEventsV1alpha1SlackTrigger) HasSlackToken() bool`

HasSlackToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


