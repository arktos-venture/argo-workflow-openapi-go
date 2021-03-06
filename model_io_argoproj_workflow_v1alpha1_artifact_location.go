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

// IoArgoprojWorkflowV1alpha1ArtifactLocation ArtifactLocation describes a location for a single or multiple artifacts. It is used as single artifact in the context of inputs/outputs (e.g. outputs.artifacts.artname). It is also used to describe the location of multiple artifacts such as the archive location of a single workflow step, which the executor will use as a default location to store its files.
type IoArgoprojWorkflowV1alpha1ArtifactLocation struct {
	// ArchiveLogs indicates if the container logs should be archived
	ArchiveLogs *bool `json:"archiveLogs,omitempty"`
	Artifactory *IoArgoprojWorkflowV1alpha1ArtifactoryArtifact `json:"artifactory,omitempty"`
	Gcs *IoArgoprojWorkflowV1alpha1GCSArtifact `json:"gcs,omitempty"`
	Git *IoArgoprojWorkflowV1alpha1GitArtifact `json:"git,omitempty"`
	Hdfs *IoArgoprojWorkflowV1alpha1HDFSArtifact `json:"hdfs,omitempty"`
	Http *IoArgoprojWorkflowV1alpha1HTTPArtifact `json:"http,omitempty"`
	Oss *IoArgoprojWorkflowV1alpha1OSSArtifact `json:"oss,omitempty"`
	Raw *IoArgoprojWorkflowV1alpha1RawArtifact `json:"raw,omitempty"`
	S3 *IoArgoprojWorkflowV1alpha1S3Artifact `json:"s3,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1ArtifactLocation instantiates a new IoArgoprojWorkflowV1alpha1ArtifactLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1ArtifactLocation() *IoArgoprojWorkflowV1alpha1ArtifactLocation {
	this := IoArgoprojWorkflowV1alpha1ArtifactLocation{}
	return &this
}

// NewIoArgoprojWorkflowV1alpha1ArtifactLocationWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ArtifactLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1ArtifactLocationWithDefaults() *IoArgoprojWorkflowV1alpha1ArtifactLocation {
	this := IoArgoprojWorkflowV1alpha1ArtifactLocation{}
	return &this
}

// GetArchiveLogs returns the ArchiveLogs field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetArchiveLogs() bool {
	if o == nil || o.ArchiveLogs == nil {
		var ret bool
		return ret
	}
	return *o.ArchiveLogs
}

// GetArchiveLogsOk returns a tuple with the ArchiveLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetArchiveLogsOk() (*bool, bool) {
	if o == nil || o.ArchiveLogs == nil {
		return nil, false
	}
	return o.ArchiveLogs, true
}

// HasArchiveLogs returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasArchiveLogs() bool {
	if o != nil && o.ArchiveLogs != nil {
		return true
	}

	return false
}

// SetArchiveLogs gets a reference to the given bool and assigns it to the ArchiveLogs field.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetArchiveLogs(v bool) {
	o.ArchiveLogs = &v
}

// GetArtifactory returns the Artifactory field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetArtifactory() IoArgoprojWorkflowV1alpha1ArtifactoryArtifact {
	if o == nil || o.Artifactory == nil {
		var ret IoArgoprojWorkflowV1alpha1ArtifactoryArtifact
		return ret
	}
	return *o.Artifactory
}

// GetArtifactoryOk returns a tuple with the Artifactory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetArtifactoryOk() (*IoArgoprojWorkflowV1alpha1ArtifactoryArtifact, bool) {
	if o == nil || o.Artifactory == nil {
		return nil, false
	}
	return o.Artifactory, true
}

// HasArtifactory returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasArtifactory() bool {
	if o != nil && o.Artifactory != nil {
		return true
	}

	return false
}

// SetArtifactory gets a reference to the given IoArgoprojWorkflowV1alpha1ArtifactoryArtifact and assigns it to the Artifactory field.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetArtifactory(v IoArgoprojWorkflowV1alpha1ArtifactoryArtifact) {
	o.Artifactory = &v
}

// GetGcs returns the Gcs field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetGcs() IoArgoprojWorkflowV1alpha1GCSArtifact {
	if o == nil || o.Gcs == nil {
		var ret IoArgoprojWorkflowV1alpha1GCSArtifact
		return ret
	}
	return *o.Gcs
}

// GetGcsOk returns a tuple with the Gcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetGcsOk() (*IoArgoprojWorkflowV1alpha1GCSArtifact, bool) {
	if o == nil || o.Gcs == nil {
		return nil, false
	}
	return o.Gcs, true
}

// HasGcs returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasGcs() bool {
	if o != nil && o.Gcs != nil {
		return true
	}

	return false
}

// SetGcs gets a reference to the given IoArgoprojWorkflowV1alpha1GCSArtifact and assigns it to the Gcs field.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetGcs(v IoArgoprojWorkflowV1alpha1GCSArtifact) {
	o.Gcs = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetGit() IoArgoprojWorkflowV1alpha1GitArtifact {
	if o == nil || o.Git == nil {
		var ret IoArgoprojWorkflowV1alpha1GitArtifact
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetGitOk() (*IoArgoprojWorkflowV1alpha1GitArtifact, bool) {
	if o == nil || o.Git == nil {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasGit() bool {
	if o != nil && o.Git != nil {
		return true
	}

	return false
}

// SetGit gets a reference to the given IoArgoprojWorkflowV1alpha1GitArtifact and assigns it to the Git field.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetGit(v IoArgoprojWorkflowV1alpha1GitArtifact) {
	o.Git = &v
}

// GetHdfs returns the Hdfs field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetHdfs() IoArgoprojWorkflowV1alpha1HDFSArtifact {
	if o == nil || o.Hdfs == nil {
		var ret IoArgoprojWorkflowV1alpha1HDFSArtifact
		return ret
	}
	return *o.Hdfs
}

// GetHdfsOk returns a tuple with the Hdfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetHdfsOk() (*IoArgoprojWorkflowV1alpha1HDFSArtifact, bool) {
	if o == nil || o.Hdfs == nil {
		return nil, false
	}
	return o.Hdfs, true
}

// HasHdfs returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasHdfs() bool {
	if o != nil && o.Hdfs != nil {
		return true
	}

	return false
}

// SetHdfs gets a reference to the given IoArgoprojWorkflowV1alpha1HDFSArtifact and assigns it to the Hdfs field.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetHdfs(v IoArgoprojWorkflowV1alpha1HDFSArtifact) {
	o.Hdfs = &v
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetHttp() IoArgoprojWorkflowV1alpha1HTTPArtifact {
	if o == nil || o.Http == nil {
		var ret IoArgoprojWorkflowV1alpha1HTTPArtifact
		return ret
	}
	return *o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetHttpOk() (*IoArgoprojWorkflowV1alpha1HTTPArtifact, bool) {
	if o == nil || o.Http == nil {
		return nil, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasHttp() bool {
	if o != nil && o.Http != nil {
		return true
	}

	return false
}

// SetHttp gets a reference to the given IoArgoprojWorkflowV1alpha1HTTPArtifact and assigns it to the Http field.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetHttp(v IoArgoprojWorkflowV1alpha1HTTPArtifact) {
	o.Http = &v
}

// GetOss returns the Oss field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetOss() IoArgoprojWorkflowV1alpha1OSSArtifact {
	if o == nil || o.Oss == nil {
		var ret IoArgoprojWorkflowV1alpha1OSSArtifact
		return ret
	}
	return *o.Oss
}

// GetOssOk returns a tuple with the Oss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetOssOk() (*IoArgoprojWorkflowV1alpha1OSSArtifact, bool) {
	if o == nil || o.Oss == nil {
		return nil, false
	}
	return o.Oss, true
}

// HasOss returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasOss() bool {
	if o != nil && o.Oss != nil {
		return true
	}

	return false
}

// SetOss gets a reference to the given IoArgoprojWorkflowV1alpha1OSSArtifact and assigns it to the Oss field.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetOss(v IoArgoprojWorkflowV1alpha1OSSArtifact) {
	o.Oss = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetRaw() IoArgoprojWorkflowV1alpha1RawArtifact {
	if o == nil || o.Raw == nil {
		var ret IoArgoprojWorkflowV1alpha1RawArtifact
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetRawOk() (*IoArgoprojWorkflowV1alpha1RawArtifact, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given IoArgoprojWorkflowV1alpha1RawArtifact and assigns it to the Raw field.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetRaw(v IoArgoprojWorkflowV1alpha1RawArtifact) {
	o.Raw = &v
}

// GetS3 returns the S3 field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetS3() IoArgoprojWorkflowV1alpha1S3Artifact {
	if o == nil || o.S3 == nil {
		var ret IoArgoprojWorkflowV1alpha1S3Artifact
		return ret
	}
	return *o.S3
}

// GetS3Ok returns a tuple with the S3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) GetS3Ok() (*IoArgoprojWorkflowV1alpha1S3Artifact, bool) {
	if o == nil || o.S3 == nil {
		return nil, false
	}
	return o.S3, true
}

// HasS3 returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) HasS3() bool {
	if o != nil && o.S3 != nil {
		return true
	}

	return false
}

// SetS3 gets a reference to the given IoArgoprojWorkflowV1alpha1S3Artifact and assigns it to the S3 field.
func (o *IoArgoprojWorkflowV1alpha1ArtifactLocation) SetS3(v IoArgoprojWorkflowV1alpha1S3Artifact) {
	o.S3 = &v
}

func (o IoArgoprojWorkflowV1alpha1ArtifactLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArchiveLogs != nil {
		toSerialize["archiveLogs"] = o.ArchiveLogs
	}
	if o.Artifactory != nil {
		toSerialize["artifactory"] = o.Artifactory
	}
	if o.Gcs != nil {
		toSerialize["gcs"] = o.Gcs
	}
	if o.Git != nil {
		toSerialize["git"] = o.Git
	}
	if o.Hdfs != nil {
		toSerialize["hdfs"] = o.Hdfs
	}
	if o.Http != nil {
		toSerialize["http"] = o.Http
	}
	if o.Oss != nil {
		toSerialize["oss"] = o.Oss
	}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.S3 != nil {
		toSerialize["s3"] = o.S3
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1ArtifactLocation struct {
	value *IoArgoprojWorkflowV1alpha1ArtifactLocation
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1ArtifactLocation) Get() *IoArgoprojWorkflowV1alpha1ArtifactLocation {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1ArtifactLocation) Set(val *IoArgoprojWorkflowV1alpha1ArtifactLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1ArtifactLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1ArtifactLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1ArtifactLocation(val *IoArgoprojWorkflowV1alpha1ArtifactLocation) *NullableIoArgoprojWorkflowV1alpha1ArtifactLocation {
	return &NullableIoArgoprojWorkflowV1alpha1ArtifactLocation{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1ArtifactLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1ArtifactLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


