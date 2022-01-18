# IoArgoprojEventsV1alpha1S3Artifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 
**Bucket** | Pointer to [**IoArgoprojEventsV1alpha1S3Bucket**](IoArgoprojEventsV1alpha1S3Bucket.md) |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **[]string** |  | [optional] 
**Filter** | Pointer to [**IoArgoprojEventsV1alpha1S3Filter**](IoArgoprojEventsV1alpha1S3Filter.md) |  | [optional] 
**Insecure** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1S3Artifact

`func NewIoArgoprojEventsV1alpha1S3Artifact() *IoArgoprojEventsV1alpha1S3Artifact`

NewIoArgoprojEventsV1alpha1S3Artifact instantiates a new IoArgoprojEventsV1alpha1S3Artifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1S3ArtifactWithDefaults

`func NewIoArgoprojEventsV1alpha1S3ArtifactWithDefaults() *IoArgoprojEventsV1alpha1S3Artifact`

NewIoArgoprojEventsV1alpha1S3ArtifactWithDefaults instantiates a new IoArgoprojEventsV1alpha1S3Artifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetAccessKey() IoK8sApiCoreV1SecretKeySelector`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetAccessKeyOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *IoArgoprojEventsV1alpha1S3Artifact) SetAccessKey(v IoK8sApiCoreV1SecretKeySelector)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *IoArgoprojEventsV1alpha1S3Artifact) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetBucket

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetBucket() IoArgoprojEventsV1alpha1S3Bucket`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetBucketOk() (*IoArgoprojEventsV1alpha1S3Bucket, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IoArgoprojEventsV1alpha1S3Artifact) SetBucket(v IoArgoprojEventsV1alpha1S3Bucket)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *IoArgoprojEventsV1alpha1S3Artifact) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetEndpoint

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *IoArgoprojEventsV1alpha1S3Artifact) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *IoArgoprojEventsV1alpha1S3Artifact) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEvents

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *IoArgoprojEventsV1alpha1S3Artifact) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *IoArgoprojEventsV1alpha1S3Artifact) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetFilter

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetFilter() IoArgoprojEventsV1alpha1S3Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetFilterOk() (*IoArgoprojEventsV1alpha1S3Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IoArgoprojEventsV1alpha1S3Artifact) SetFilter(v IoArgoprojEventsV1alpha1S3Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IoArgoprojEventsV1alpha1S3Artifact) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetInsecure

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *IoArgoprojEventsV1alpha1S3Artifact) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *IoArgoprojEventsV1alpha1S3Artifact) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1S3Artifact) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1S3Artifact) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRegion

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *IoArgoprojEventsV1alpha1S3Artifact) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *IoArgoprojEventsV1alpha1S3Artifact) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretKey

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetSecretKey() IoK8sApiCoreV1SecretKeySelector`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *IoArgoprojEventsV1alpha1S3Artifact) GetSecretKeyOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *IoArgoprojEventsV1alpha1S3Artifact) SetSecretKey(v IoK8sApiCoreV1SecretKeySelector)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *IoArgoprojEventsV1alpha1S3Artifact) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


