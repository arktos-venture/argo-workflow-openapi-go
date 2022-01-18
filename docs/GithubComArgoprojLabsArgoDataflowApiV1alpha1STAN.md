# GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1STANAuth**](GithubComArgoprojLabsArgoDataflowApiV1alpha1STANAuth.md) |  | [optional] 
**ClusterId** | Pointer to **string** |  | [optional] 
**MaxInflight** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NatsMonitoringUrl** | Pointer to **string** |  | [optional] 
**NatsUrl** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**SubjectPrefix** | Pointer to **string** |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1STAN

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1STAN() *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1STAN instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1STANWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1STANWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1STANWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetAuth() GithubComArgoprojLabsArgoDataflowApiV1alpha1STANAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetAuthOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1STANAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) SetAuth(v GithubComArgoprojLabsArgoDataflowApiV1alpha1STANAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetClusterId

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetMaxInflight

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetMaxInflight() int32`

GetMaxInflight returns the MaxInflight field if non-nil, zero value otherwise.

### GetMaxInflightOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetMaxInflightOk() (*int32, bool)`

GetMaxInflightOk returns a tuple with the MaxInflight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInflight

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) SetMaxInflight(v int32)`

SetMaxInflight sets MaxInflight field to given value.

### HasMaxInflight

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) HasMaxInflight() bool`

HasMaxInflight returns a boolean if a field has been set.

### GetName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNatsMonitoringUrl

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetNatsMonitoringUrl() string`

GetNatsMonitoringUrl returns the NatsMonitoringUrl field if non-nil, zero value otherwise.

### GetNatsMonitoringUrlOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetNatsMonitoringUrlOk() (*string, bool)`

GetNatsMonitoringUrlOk returns a tuple with the NatsMonitoringUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatsMonitoringUrl

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) SetNatsMonitoringUrl(v string)`

SetNatsMonitoringUrl sets NatsMonitoringUrl field to given value.

### HasNatsMonitoringUrl

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) HasNatsMonitoringUrl() bool`

HasNatsMonitoringUrl returns a boolean if a field has been set.

### GetNatsUrl

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetNatsUrl() string`

GetNatsUrl returns the NatsUrl field if non-nil, zero value otherwise.

### GetNatsUrlOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetNatsUrlOk() (*string, bool)`

GetNatsUrlOk returns a tuple with the NatsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatsUrl

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) SetNatsUrl(v string)`

SetNatsUrl sets NatsUrl field to given value.

### HasNatsUrl

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) HasNatsUrl() bool`

HasNatsUrl returns a boolean if a field has been set.

### GetSubject

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectPrefix

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetSubjectPrefix() string`

GetSubjectPrefix returns the SubjectPrefix field if non-nil, zero value otherwise.

### GetSubjectPrefixOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) GetSubjectPrefixOk() (*string, bool)`

GetSubjectPrefixOk returns a tuple with the SubjectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectPrefix

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) SetSubjectPrefix(v string)`

SetSubjectPrefix sets SubjectPrefix field to given value.

### HasSubjectPrefix

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1STAN) HasSubjectPrefix() bool`

HasSubjectPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


