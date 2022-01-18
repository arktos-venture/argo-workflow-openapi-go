# IoArgoprojEventsV1alpha1SASLConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mechanism** | Pointer to **string** |  | [optional] 
**Password** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**User** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1SASLConfig

`func NewIoArgoprojEventsV1alpha1SASLConfig() *IoArgoprojEventsV1alpha1SASLConfig`

NewIoArgoprojEventsV1alpha1SASLConfig instantiates a new IoArgoprojEventsV1alpha1SASLConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1SASLConfigWithDefaults

`func NewIoArgoprojEventsV1alpha1SASLConfigWithDefaults() *IoArgoprojEventsV1alpha1SASLConfig`

NewIoArgoprojEventsV1alpha1SASLConfigWithDefaults instantiates a new IoArgoprojEventsV1alpha1SASLConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMechanism

`func (o *IoArgoprojEventsV1alpha1SASLConfig) GetMechanism() string`

GetMechanism returns the Mechanism field if non-nil, zero value otherwise.

### GetMechanismOk

`func (o *IoArgoprojEventsV1alpha1SASLConfig) GetMechanismOk() (*string, bool)`

GetMechanismOk returns a tuple with the Mechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMechanism

`func (o *IoArgoprojEventsV1alpha1SASLConfig) SetMechanism(v string)`

SetMechanism sets Mechanism field to given value.

### HasMechanism

`func (o *IoArgoprojEventsV1alpha1SASLConfig) HasMechanism() bool`

HasMechanism returns a boolean if a field has been set.

### GetPassword

`func (o *IoArgoprojEventsV1alpha1SASLConfig) GetPassword() IoK8sApiCoreV1SecretKeySelector`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IoArgoprojEventsV1alpha1SASLConfig) GetPasswordOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IoArgoprojEventsV1alpha1SASLConfig) SetPassword(v IoK8sApiCoreV1SecretKeySelector)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IoArgoprojEventsV1alpha1SASLConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUser

`func (o *IoArgoprojEventsV1alpha1SASLConfig) GetUser() IoK8sApiCoreV1SecretKeySelector`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IoArgoprojEventsV1alpha1SASLConfig) GetUserOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IoArgoprojEventsV1alpha1SASLConfig) SetUser(v IoK8sApiCoreV1SecretKeySelector)`

SetUser sets User field to given value.

### HasUser

`func (o *IoArgoprojEventsV1alpha1SASLConfig) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


