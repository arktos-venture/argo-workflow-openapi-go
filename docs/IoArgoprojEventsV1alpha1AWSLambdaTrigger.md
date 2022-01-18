# IoArgoprojEventsV1alpha1AWSLambdaTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**FunctionName** | Pointer to **string** | FunctionName refers to the name of the function to invoke. | [optional] 
**InvocationType** | Pointer to **string** | Choose from the following options.     * RequestResponse (default) - Invoke the function synchronously. Keep    the connection open until the function returns a response or times out.    The API response includes the function response and additional data.     * Event - Invoke the function asynchronously. Send events that fail multiple    times to the function&#39;s dead-letter queue (if it&#39;s configured). The API    response only includes a status code.     * DryRun - Validate parameter values and verify that the user or role    has permission to invoke the function. +optional | [optional] 
**Parameters** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) |  | [optional] 
**Payload** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) | Payload is the list of key-value extracted from an event payload to construct the request payload. | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1AWSLambdaTrigger

`func NewIoArgoprojEventsV1alpha1AWSLambdaTrigger() *IoArgoprojEventsV1alpha1AWSLambdaTrigger`

NewIoArgoprojEventsV1alpha1AWSLambdaTrigger instantiates a new IoArgoprojEventsV1alpha1AWSLambdaTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1AWSLambdaTriggerWithDefaults

`func NewIoArgoprojEventsV1alpha1AWSLambdaTriggerWithDefaults() *IoArgoprojEventsV1alpha1AWSLambdaTrigger`

NewIoArgoprojEventsV1alpha1AWSLambdaTriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1AWSLambdaTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetAccessKey() IoK8sApiCoreV1SecretKeySelector`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetAccessKeyOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) SetAccessKey(v IoK8sApiCoreV1SecretKeySelector)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetFunctionName

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetInvocationType

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetInvocationType() string`

GetInvocationType returns the InvocationType field if non-nil, zero value otherwise.

### GetInvocationTypeOk

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetInvocationTypeOk() (*string, bool)`

GetInvocationTypeOk returns a tuple with the InvocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationType

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) SetInvocationType(v string)`

SetInvocationType sets InvocationType field to given value.

### HasInvocationType

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) HasInvocationType() bool`

HasInvocationType returns a boolean if a field has been set.

### GetParameters

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPayload

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetPayload() []IoArgoprojEventsV1alpha1TriggerParameter`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetPayloadOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) SetPayload(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetRegion

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretKey

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetSecretKey() IoK8sApiCoreV1SecretKeySelector`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) GetSecretKeyOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) SetSecretKey(v IoK8sApiCoreV1SecretKeySelector)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *IoArgoprojEventsV1alpha1AWSLambdaTrigger) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


