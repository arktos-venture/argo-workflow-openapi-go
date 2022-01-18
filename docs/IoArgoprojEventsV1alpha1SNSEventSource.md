# IoArgoprojEventsV1alpha1SNSEventSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RoleARN** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**TopicArn** | Pointer to **string** |  | [optional] 
**ValidateSignature** | Pointer to **bool** |  | [optional] 
**Webhook** | Pointer to [**IoArgoprojEventsV1alpha1WebhookContext**](IoArgoprojEventsV1alpha1WebhookContext.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1SNSEventSource

`func NewIoArgoprojEventsV1alpha1SNSEventSource() *IoArgoprojEventsV1alpha1SNSEventSource`

NewIoArgoprojEventsV1alpha1SNSEventSource instantiates a new IoArgoprojEventsV1alpha1SNSEventSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1SNSEventSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1SNSEventSourceWithDefaults() *IoArgoprojEventsV1alpha1SNSEventSource`

NewIoArgoprojEventsV1alpha1SNSEventSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1SNSEventSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetAccessKey() IoK8sApiCoreV1SecretKeySelector`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetAccessKeyOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) SetAccessKey(v IoK8sApiCoreV1SecretKeySelector)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRegion

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoleARN

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetRoleARN() string`

GetRoleARN returns the RoleARN field if non-nil, zero value otherwise.

### GetRoleARNOk

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetRoleARNOk() (*string, bool)`

GetRoleARNOk returns a tuple with the RoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleARN

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) SetRoleARN(v string)`

SetRoleARN sets RoleARN field to given value.

### HasRoleARN

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) HasRoleARN() bool`

HasRoleARN returns a boolean if a field has been set.

### GetSecretKey

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetSecretKey() IoK8sApiCoreV1SecretKeySelector`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetSecretKeyOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) SetSecretKey(v IoK8sApiCoreV1SecretKeySelector)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetTopicArn

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetTopicArn() string`

GetTopicArn returns the TopicArn field if non-nil, zero value otherwise.

### GetTopicArnOk

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetTopicArnOk() (*string, bool)`

GetTopicArnOk returns a tuple with the TopicArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicArn

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) SetTopicArn(v string)`

SetTopicArn sets TopicArn field to given value.

### HasTopicArn

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) HasTopicArn() bool`

HasTopicArn returns a boolean if a field has been set.

### GetValidateSignature

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetValidateSignature() bool`

GetValidateSignature returns the ValidateSignature field if non-nil, zero value otherwise.

### GetValidateSignatureOk

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetValidateSignatureOk() (*bool, bool)`

GetValidateSignatureOk returns a tuple with the ValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateSignature

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) SetValidateSignature(v bool)`

SetValidateSignature sets ValidateSignature field to given value.

### HasValidateSignature

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) HasValidateSignature() bool`

HasValidateSignature returns a boolean if a field has been set.

### GetWebhook

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetWebhook() IoArgoprojEventsV1alpha1WebhookContext`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) GetWebhookOk() (*IoArgoprojEventsV1alpha1WebhookContext, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) SetWebhook(v IoArgoprojEventsV1alpha1WebhookContext)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *IoArgoprojEventsV1alpha1SNSEventSource) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


