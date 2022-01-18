# IoArgoprojWorkflowV1alpha1Link

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the link, E.g. \&quot;Workflow Logs\&quot; or \&quot;Pod Logs\&quot; | 
**Scope** | **string** | \&quot;workflow\&quot;, \&quot;pod\&quot;, \&quot;pod-logs\&quot;, \&quot;event-source-logs\&quot;, \&quot;sensor-logs\&quot; or \&quot;chat\&quot; | 
**Url** | **string** | The URL. Can contain \&quot;${metadata.namespace}\&quot;, \&quot;${metadata.name}\&quot;, \&quot;${status.startedAt}\&quot;, \&quot;${status.finishedAt}\&quot; or any other element in workflow yaml, e.g. \&quot;${io.argoproj.workflow.v1alpha1.metadata.annotations.userDefinedKey}\&quot; | 

## Methods

### NewIoArgoprojWorkflowV1alpha1Link

`func NewIoArgoprojWorkflowV1alpha1Link(name string, scope string, url string, ) *IoArgoprojWorkflowV1alpha1Link`

NewIoArgoprojWorkflowV1alpha1Link instantiates a new IoArgoprojWorkflowV1alpha1Link object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1LinkWithDefaults

`func NewIoArgoprojWorkflowV1alpha1LinkWithDefaults() *IoArgoprojWorkflowV1alpha1Link`

NewIoArgoprojWorkflowV1alpha1LinkWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Link object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoArgoprojWorkflowV1alpha1Link) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojWorkflowV1alpha1Link) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojWorkflowV1alpha1Link) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *IoArgoprojWorkflowV1alpha1Link) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IoArgoprojWorkflowV1alpha1Link) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IoArgoprojWorkflowV1alpha1Link) SetScope(v string)`

SetScope sets Scope field to given value.


### GetUrl

`func (o *IoArgoprojWorkflowV1alpha1Link) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IoArgoprojWorkflowV1alpha1Link) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IoArgoprojWorkflowV1alpha1Link) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


