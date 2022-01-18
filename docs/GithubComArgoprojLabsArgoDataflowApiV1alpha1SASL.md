# GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mechanism** | Pointer to **string** |  | [optional] 
**Password** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**User** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1SASL

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1SASL() *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1SASL instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1SASLWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1SASLWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1SASLWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMechanism

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL) GetMechanism() string`

GetMechanism returns the Mechanism field if non-nil, zero value otherwise.

### GetMechanismOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL) GetMechanismOk() (*string, bool)`

GetMechanismOk returns a tuple with the Mechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMechanism

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL) SetMechanism(v string)`

SetMechanism sets Mechanism field to given value.

### HasMechanism

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL) HasMechanism() bool`

HasMechanism returns a boolean if a field has been set.

### GetPassword

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL) GetPassword() IoK8sApiCoreV1SecretKeySelector`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL) GetPasswordOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL) SetPassword(v IoK8sApiCoreV1SecretKeySelector)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUser

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL) GetUser() IoK8sApiCoreV1SecretKeySelector`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL) GetUserOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL) SetUser(v IoK8sApiCoreV1SecretKeySelector)`

SetUser sets User field to given value.

### HasUser

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1SASL) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


