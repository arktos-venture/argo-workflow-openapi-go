# IoArgoprojWorkflowV1alpha1HTTPArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to [**[]IoArgoprojWorkflowV1alpha1Header**](IoArgoprojWorkflowV1alpha1Header.md) | Headers are an optional list of headers to send with HTTP requests for artifacts | [optional] 
**Url** | **string** | URL of the artifact | 

## Methods

### NewIoArgoprojWorkflowV1alpha1HTTPArtifact

`func NewIoArgoprojWorkflowV1alpha1HTTPArtifact(url string, ) *IoArgoprojWorkflowV1alpha1HTTPArtifact`

NewIoArgoprojWorkflowV1alpha1HTTPArtifact instantiates a new IoArgoprojWorkflowV1alpha1HTTPArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1HTTPArtifactWithDefaults

`func NewIoArgoprojWorkflowV1alpha1HTTPArtifactWithDefaults() *IoArgoprojWorkflowV1alpha1HTTPArtifact`

NewIoArgoprojWorkflowV1alpha1HTTPArtifactWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1HTTPArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *IoArgoprojWorkflowV1alpha1HTTPArtifact) GetHeaders() []IoArgoprojWorkflowV1alpha1Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *IoArgoprojWorkflowV1alpha1HTTPArtifact) GetHeadersOk() (*[]IoArgoprojWorkflowV1alpha1Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *IoArgoprojWorkflowV1alpha1HTTPArtifact) SetHeaders(v []IoArgoprojWorkflowV1alpha1Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *IoArgoprojWorkflowV1alpha1HTTPArtifact) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojWorkflowV1alpha1HTTPArtifact) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojWorkflowV1alpha1HTTPArtifact) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojWorkflowV1alpha1HTTPArtifact) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


