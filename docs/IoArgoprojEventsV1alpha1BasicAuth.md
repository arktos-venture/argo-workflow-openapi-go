# IoArgoprojEventsV1alpha1BasicAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Username** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1BasicAuth

`func NewIoArgoprojEventsV1alpha1BasicAuth() *IoArgoprojEventsV1alpha1BasicAuth`

NewIoArgoprojEventsV1alpha1BasicAuth instantiates a new IoArgoprojEventsV1alpha1BasicAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1BasicAuthWithDefaults

`func NewIoArgoprojEventsV1alpha1BasicAuthWithDefaults() *IoArgoprojEventsV1alpha1BasicAuth`

NewIoArgoprojEventsV1alpha1BasicAuthWithDefaults instantiates a new IoArgoprojEventsV1alpha1BasicAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *IoArgoprojEventsV1alpha1BasicAuth) GetPassword() IoK8sApiCoreV1SecretKeySelector`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IoArgoprojEventsV1alpha1BasicAuth) GetPasswordOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IoArgoprojEventsV1alpha1BasicAuth) SetPassword(v IoK8sApiCoreV1SecretKeySelector)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IoArgoprojEventsV1alpha1BasicAuth) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *IoArgoprojEventsV1alpha1BasicAuth) GetUsername() IoK8sApiCoreV1SecretKeySelector`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IoArgoprojEventsV1alpha1BasicAuth) GetUsernameOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IoArgoprojEventsV1alpha1BasicAuth) SetUsername(v IoK8sApiCoreV1SecretKeySelector)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IoArgoprojEventsV1alpha1BasicAuth) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


