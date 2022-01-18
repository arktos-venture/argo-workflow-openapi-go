# GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitInterval** | Pointer to [**IoK8sApimachineryPkgApisMetaV1Duration**](IoK8sApimachineryPkgApisMetaV1Duration.md) |  | [optional] 
**Database** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1Database**](GithubComArgoprojLabsArgoDataflowApiV1alpha1Database.md) |  | [optional] 
**InitSchema** | Pointer to **bool** |  | [optional] 
**OffsetColumn** | Pointer to **string** |  | [optional] 
**PollInterval** | Pointer to [**IoK8sApimachineryPkgApisMetaV1Duration**](IoK8sApimachineryPkgApisMetaV1Duration.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource() *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSourceWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSourceWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSourceWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitInterval

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetCommitInterval() IoK8sApimachineryPkgApisMetaV1Duration`

GetCommitInterval returns the CommitInterval field if non-nil, zero value otherwise.

### GetCommitIntervalOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetCommitIntervalOk() (*IoK8sApimachineryPkgApisMetaV1Duration, bool)`

GetCommitIntervalOk returns a tuple with the CommitInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitInterval

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) SetCommitInterval(v IoK8sApimachineryPkgApisMetaV1Duration)`

SetCommitInterval sets CommitInterval field to given value.

### HasCommitInterval

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) HasCommitInterval() bool`

HasCommitInterval returns a boolean if a field has been set.

### GetDatabase

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetDatabase() GithubComArgoprojLabsArgoDataflowApiV1alpha1Database`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetDatabaseOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Database, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) SetDatabase(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Database)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetInitSchema

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetInitSchema() bool`

GetInitSchema returns the InitSchema field if non-nil, zero value otherwise.

### GetInitSchemaOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetInitSchemaOk() (*bool, bool)`

GetInitSchemaOk returns a tuple with the InitSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitSchema

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) SetInitSchema(v bool)`

SetInitSchema sets InitSchema field to given value.

### HasInitSchema

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) HasInitSchema() bool`

HasInitSchema returns a boolean if a field has been set.

### GetOffsetColumn

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetOffsetColumn() string`

GetOffsetColumn returns the OffsetColumn field if non-nil, zero value otherwise.

### GetOffsetColumnOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetOffsetColumnOk() (*string, bool)`

GetOffsetColumnOk returns a tuple with the OffsetColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetColumn

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) SetOffsetColumn(v string)`

SetOffsetColumn sets OffsetColumn field to given value.

### HasOffsetColumn

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) HasOffsetColumn() bool`

HasOffsetColumn returns a boolean if a field has been set.

### GetPollInterval

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetPollInterval() IoK8sApimachineryPkgApisMetaV1Duration`

GetPollInterval returns the PollInterval field if non-nil, zero value otherwise.

### GetPollIntervalOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetPollIntervalOk() (*IoK8sApimachineryPkgApisMetaV1Duration, bool)`

GetPollIntervalOk returns a tuple with the PollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInterval

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) SetPollInterval(v IoK8sApimachineryPkgApisMetaV1Duration)`

SetPollInterval sets PollInterval field to given value.

### HasPollInterval

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) HasPollInterval() bool`

HasPollInterval returns a boolean if a field has been set.

### GetQuery

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


