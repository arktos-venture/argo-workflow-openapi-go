# IoArgoprojWorkflowV1alpha1HTTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is content of the HTTP Request | [optional] 
**Headers** | Pointer to [**[]IoArgoprojWorkflowV1alpha1HTTPHeader**](IoArgoprojWorkflowV1alpha1HTTPHeader.md) | Headers are an optional list of headers to send with HTTP requests | [optional] 
**Method** | Pointer to **string** | Method is HTTP methods for HTTP Request | [optional] 
**TimeoutSeconds** | Pointer to **int32** | TimeoutSeconds is request timeout for HTTP Request. Default is 30 seconds | [optional] 
**Url** | **string** | URL of the HTTP Request | 

## Methods

### NewIoArgoprojWorkflowV1alpha1HTTP

`func NewIoArgoprojWorkflowV1alpha1HTTP(url string, ) *IoArgoprojWorkflowV1alpha1HTTP`

NewIoArgoprojWorkflowV1alpha1HTTP instantiates a new IoArgoprojWorkflowV1alpha1HTTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1HTTPWithDefaults

`func NewIoArgoprojWorkflowV1alpha1HTTPWithDefaults() *IoArgoprojWorkflowV1alpha1HTTP`

NewIoArgoprojWorkflowV1alpha1HTTPWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1HTTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *IoArgoprojWorkflowV1alpha1HTTP) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *IoArgoprojWorkflowV1alpha1HTTP) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *IoArgoprojWorkflowV1alpha1HTTP) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *IoArgoprojWorkflowV1alpha1HTTP) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeaders

`func (o *IoArgoprojWorkflowV1alpha1HTTP) GetHeaders() []IoArgoprojWorkflowV1alpha1HTTPHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *IoArgoprojWorkflowV1alpha1HTTP) GetHeadersOk() (*[]IoArgoprojWorkflowV1alpha1HTTPHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *IoArgoprojWorkflowV1alpha1HTTP) SetHeaders(v []IoArgoprojWorkflowV1alpha1HTTPHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *IoArgoprojWorkflowV1alpha1HTTP) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *IoArgoprojWorkflowV1alpha1HTTP) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *IoArgoprojWorkflowV1alpha1HTTP) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *IoArgoprojWorkflowV1alpha1HTTP) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *IoArgoprojWorkflowV1alpha1HTTP) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *IoArgoprojWorkflowV1alpha1HTTP) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *IoArgoprojWorkflowV1alpha1HTTP) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *IoArgoprojWorkflowV1alpha1HTTP) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *IoArgoprojWorkflowV1alpha1HTTP) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetUrl

`func (o *IoArgoprojWorkflowV1alpha1HTTP) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojWorkflowV1alpha1HTTP) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojWorkflowV1alpha1HTTP) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


