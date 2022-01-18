# IoArgoprojWorkflowV1alpha1ResourceTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action is the action to perform to the resource. Must be one of: get, create, apply, delete, replace, patch | 
**FailureCondition** | Pointer to **string** | FailureCondition is a label selector expression which describes the conditions of the k8s resource in which the step was considered failed | [optional] 
**Flags** | Pointer to **[]string** | Flags is a set of additional options passed to kubectl before submitting a resource I.e. to disable resource validation: flags: [  \&quot;--validate&#x3D;false\&quot;  # disable resource validation ] | [optional] 
**Manifest** | Pointer to **string** | Manifest contains the kubernetes manifest | [optional] 
**MergeStrategy** | Pointer to **string** | MergeStrategy is the strategy used to merge a patch. It defaults to \&quot;strategic\&quot; Must be one of: strategic, merge, json | [optional] 
**SetOwnerReference** | Pointer to **bool** | SetOwnerReference sets the reference to the workflow on the OwnerReference of generated resource. | [optional] 
**SuccessCondition** | Pointer to **string** | SuccessCondition is a label selector expression which describes the conditions of the k8s resource in which it is acceptable to proceed to the following step | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1ResourceTemplate

`func NewIoArgoprojWorkflowV1alpha1ResourceTemplate(action string, ) *IoArgoprojWorkflowV1alpha1ResourceTemplate`

NewIoArgoprojWorkflowV1alpha1ResourceTemplate instantiates a new IoArgoprojWorkflowV1alpha1ResourceTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ResourceTemplateWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ResourceTemplateWithDefaults() *IoArgoprojWorkflowV1alpha1ResourceTemplate`

NewIoArgoprojWorkflowV1alpha1ResourceTemplateWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ResourceTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetAction(v string)`

SetAction sets Action field to given value.


### GetFailureCondition

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetFailureCondition() string`

GetFailureCondition returns the FailureCondition field if non-nil, zero value otherwise.

### GetFailureConditionOk

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetFailureConditionOk() (*string, bool)`

GetFailureConditionOk returns a tuple with the FailureCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCondition

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetFailureCondition(v string)`

SetFailureCondition sets FailureCondition field to given value.

### HasFailureCondition

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) HasFailureCondition() bool`

HasFailureCondition returns a boolean if a field has been set.

### GetFlags

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetManifest

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetManifest() string`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetManifestOk() (*string, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetManifest(v string)`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetMergeStrategy

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetMergeStrategy() string`

GetMergeStrategy returns the MergeStrategy field if non-nil, zero value otherwise.

### GetMergeStrategyOk

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetMergeStrategyOk() (*string, bool)`

GetMergeStrategyOk returns a tuple with the MergeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeStrategy

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetMergeStrategy(v string)`

SetMergeStrategy sets MergeStrategy field to given value.

### HasMergeStrategy

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) HasMergeStrategy() bool`

HasMergeStrategy returns a boolean if a field has been set.

### GetSetOwnerReference

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetSetOwnerReference() bool`

GetSetOwnerReference returns the SetOwnerReference field if non-nil, zero value otherwise.

### GetSetOwnerReferenceOk

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetSetOwnerReferenceOk() (*bool, bool)`

GetSetOwnerReferenceOk returns a tuple with the SetOwnerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetOwnerReference

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetSetOwnerReference(v bool)`

SetSetOwnerReference sets SetOwnerReference field to given value.

### HasSetOwnerReference

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) HasSetOwnerReference() bool`

HasSetOwnerReference returns a boolean if a field has been set.

### GetSuccessCondition

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetSuccessCondition() string`

GetSuccessCondition returns the SuccessCondition field if non-nil, zero value otherwise.

### GetSuccessConditionOk

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) GetSuccessConditionOk() (*string, bool)`

GetSuccessConditionOk returns a tuple with the SuccessCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCondition

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) SetSuccessCondition(v string)`

SetSuccessCondition sets SuccessCondition field to given value.

### HasSuccessCondition

`func (o *IoArgoprojWorkflowV1alpha1ResourceTemplate) HasSuccessCondition() bool`

HasSuccessCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


