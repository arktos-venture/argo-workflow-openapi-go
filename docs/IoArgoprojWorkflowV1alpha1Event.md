# IoArgoprojWorkflowV1alpha1Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selector** | **string** | Selector (https://github.com/antonmedv/expr) that we must must match the io.argoproj.workflow.v1alpha1. E.g. &#x60;payload.message &#x3D;&#x3D; \&quot;test\&quot;&#x60; | 

## Methods

### NewIoArgoprojWorkflowV1alpha1Event

`func NewIoArgoprojWorkflowV1alpha1Event(selector string, ) *IoArgoprojWorkflowV1alpha1Event`

NewIoArgoprojWorkflowV1alpha1Event instantiates a new IoArgoprojWorkflowV1alpha1Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1EventWithDefaults

`func NewIoArgoprojWorkflowV1alpha1EventWithDefaults() *IoArgoprojWorkflowV1alpha1Event`

NewIoArgoprojWorkflowV1alpha1EventWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelector

`func (o *IoArgoprojWorkflowV1alpha1Event) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *IoArgoprojWorkflowV1alpha1Event) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *IoArgoprojWorkflowV1alpha1Event) SetSelector(v string)`

SetSelector sets Selector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


