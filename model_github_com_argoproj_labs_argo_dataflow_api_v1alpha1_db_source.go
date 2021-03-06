/*
Argo Server API

You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`

API version: VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource
type GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource struct {
	CommitInterval *IoK8sApimachineryPkgApisMetaV1Duration `json:"commitInterval,omitempty"`
	Database *GithubComArgoprojLabsArgoDataflowApiV1alpha1Database `json:"database,omitempty"`
	InitSchema *bool `json:"initSchema,omitempty"`
	OffsetColumn *string `json:"offsetColumn,omitempty"`
	PollInterval *IoK8sApimachineryPkgApisMetaV1Duration `json:"pollInterval,omitempty"`
	Query *string `json:"query,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource() *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSourceWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSourceWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource{}
	return &this
}

// GetCommitInterval returns the CommitInterval field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetCommitInterval() IoK8sApimachineryPkgApisMetaV1Duration {
	if o == nil || o.CommitInterval == nil {
		var ret IoK8sApimachineryPkgApisMetaV1Duration
		return ret
	}
	return *o.CommitInterval
}

// GetCommitIntervalOk returns a tuple with the CommitInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetCommitIntervalOk() (*IoK8sApimachineryPkgApisMetaV1Duration, bool) {
	if o == nil || o.CommitInterval == nil {
		return nil, false
	}
	return o.CommitInterval, true
}

// HasCommitInterval returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) HasCommitInterval() bool {
	if o != nil && o.CommitInterval != nil {
		return true
	}

	return false
}

// SetCommitInterval gets a reference to the given IoK8sApimachineryPkgApisMetaV1Duration and assigns it to the CommitInterval field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) SetCommitInterval(v IoK8sApimachineryPkgApisMetaV1Duration) {
	o.CommitInterval = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetDatabase() GithubComArgoprojLabsArgoDataflowApiV1alpha1Database {
	if o == nil || o.Database == nil {
		var ret GithubComArgoprojLabsArgoDataflowApiV1alpha1Database
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetDatabaseOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1Database, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) HasDatabase() bool {
	if o != nil && o.Database != nil {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given GithubComArgoprojLabsArgoDataflowApiV1alpha1Database and assigns it to the Database field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) SetDatabase(v GithubComArgoprojLabsArgoDataflowApiV1alpha1Database) {
	o.Database = &v
}

// GetInitSchema returns the InitSchema field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetInitSchema() bool {
	if o == nil || o.InitSchema == nil {
		var ret bool
		return ret
	}
	return *o.InitSchema
}

// GetInitSchemaOk returns a tuple with the InitSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetInitSchemaOk() (*bool, bool) {
	if o == nil || o.InitSchema == nil {
		return nil, false
	}
	return o.InitSchema, true
}

// HasInitSchema returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) HasInitSchema() bool {
	if o != nil && o.InitSchema != nil {
		return true
	}

	return false
}

// SetInitSchema gets a reference to the given bool and assigns it to the InitSchema field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) SetInitSchema(v bool) {
	o.InitSchema = &v
}

// GetOffsetColumn returns the OffsetColumn field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetOffsetColumn() string {
	if o == nil || o.OffsetColumn == nil {
		var ret string
		return ret
	}
	return *o.OffsetColumn
}

// GetOffsetColumnOk returns a tuple with the OffsetColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetOffsetColumnOk() (*string, bool) {
	if o == nil || o.OffsetColumn == nil {
		return nil, false
	}
	return o.OffsetColumn, true
}

// HasOffsetColumn returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) HasOffsetColumn() bool {
	if o != nil && o.OffsetColumn != nil {
		return true
	}

	return false
}

// SetOffsetColumn gets a reference to the given string and assigns it to the OffsetColumn field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) SetOffsetColumn(v string) {
	o.OffsetColumn = &v
}

// GetPollInterval returns the PollInterval field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetPollInterval() IoK8sApimachineryPkgApisMetaV1Duration {
	if o == nil || o.PollInterval == nil {
		var ret IoK8sApimachineryPkgApisMetaV1Duration
		return ret
	}
	return *o.PollInterval
}

// GetPollIntervalOk returns a tuple with the PollInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetPollIntervalOk() (*IoK8sApimachineryPkgApisMetaV1Duration, bool) {
	if o == nil || o.PollInterval == nil {
		return nil, false
	}
	return o.PollInterval, true
}

// HasPollInterval returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) HasPollInterval() bool {
	if o != nil && o.PollInterval != nil {
		return true
	}

	return false
}

// SetPollInterval gets a reference to the given IoK8sApimachineryPkgApisMetaV1Duration and assigns it to the PollInterval field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) SetPollInterval(v IoK8sApimachineryPkgApisMetaV1Duration) {
	o.PollInterval = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) SetQuery(v string) {
	o.Query = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommitInterval != nil {
		toSerialize["commitInterval"] = o.CommitInterval
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.InitSchema != nil {
		toSerialize["initSchema"] = o.InitSchema
	}
	if o.OffsetColumn != nil {
		toSerialize["offsetColumn"] = o.OffsetColumn
	}
	if o.PollInterval != nil {
		toSerialize["pollInterval"] = o.PollInterval
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1DBSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


