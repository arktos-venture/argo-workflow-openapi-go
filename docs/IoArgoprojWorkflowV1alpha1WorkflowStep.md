# IoArgoprojWorkflowV1alpha1WorkflowStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to [**IoArgoprojWorkflowV1alpha1Arguments**](IoArgoprojWorkflowV1alpha1Arguments.md) |  | [optional] 
**ContinueOn** | Pointer to [**IoArgoprojWorkflowV1alpha1ContinueOn**](IoArgoprojWorkflowV1alpha1ContinueOn.md) |  | [optional] 
**Hooks** | Pointer to [**map[string]IoArgoprojWorkflowV1alpha1LifecycleHook**](IoArgoprojWorkflowV1alpha1LifecycleHook.md) | Hooks holds the lifecycle hook which is invoked at lifecycle of step, irrespective of the success, failure, or error status of the primary step | [optional] 
**Inline** | Pointer to [**IoArgoprojWorkflowV1alpha1Template**](IoArgoprojWorkflowV1alpha1Template.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the step | [optional] 
**OnExit** | Pointer to **string** | OnExit is a template reference which is invoked at the end of the template, irrespective of the success, failure, or error of the primary template. DEPRECATED: Use Hooks[exit].Template instead. | [optional] 
**Template** | Pointer to **string** | Template is the name of the template to execute as the step | [optional] 
**TemplateRef** | Pointer to [**IoArgoprojWorkflowV1alpha1TemplateRef**](IoArgoprojWorkflowV1alpha1TemplateRef.md) |  | [optional] 
**When** | Pointer to **string** | When is an expression in which the step should conditionally execute | [optional] 
**WithItems** | Pointer to **[]map[string]interface{}** | WithItems expands a step into multiple parallel steps from the items in the list | [optional] 
**WithParam** | Pointer to **string** | WithParam expands a step into multiple parallel steps from the value in the parameter, which is expected to be a JSON list. | [optional] 
**WithSequence** | Pointer to [**IoArgoprojWorkflowV1alpha1Sequence**](IoArgoprojWorkflowV1alpha1Sequence.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1WorkflowStep

`func NewIoArgoprojWorkflowV1alpha1WorkflowStep() *IoArgoprojWorkflowV1alpha1WorkflowStep`

NewIoArgoprojWorkflowV1alpha1WorkflowStep instantiates a new IoArgoprojWorkflowV1alpha1WorkflowStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1WorkflowStepWithDefaults

`func NewIoArgoprojWorkflowV1alpha1WorkflowStepWithDefaults() *IoArgoprojWorkflowV1alpha1WorkflowStep`

NewIoArgoprojWorkflowV1alpha1WorkflowStepWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1WorkflowStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetArguments() IoArgoprojWorkflowV1alpha1Arguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetArgumentsOk() (*IoArgoprojWorkflowV1alpha1Arguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) SetArguments(v IoArgoprojWorkflowV1alpha1Arguments)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetContinueOn

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetContinueOn() IoArgoprojWorkflowV1alpha1ContinueOn`

GetContinueOn returns the ContinueOn field if non-nil, zero value otherwise.

### GetContinueOnOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetContinueOnOk() (*IoArgoprojWorkflowV1alpha1ContinueOn, bool)`

GetContinueOnOk returns a tuple with the ContinueOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOn

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) SetContinueOn(v IoArgoprojWorkflowV1alpha1ContinueOn)`

SetContinueOn sets ContinueOn field to given value.

### HasContinueOn

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) HasContinueOn() bool`

HasContinueOn returns a boolean if a field has been set.

### GetHooks

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetHooks() map[string]IoArgoprojWorkflowV1alpha1LifecycleHook`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetHooksOk() (*map[string]IoArgoprojWorkflowV1alpha1LifecycleHook, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) SetHooks(v map[string]IoArgoprojWorkflowV1alpha1LifecycleHook)`

SetHooks sets Hooks field to given value.

### HasHooks

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) HasHooks() bool`

HasHooks returns a boolean if a field has been set.

### GetInline

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetInline() IoArgoprojWorkflowV1alpha1Template`

GetInline returns the Inline field if non-nil, zero value otherwise.

### GetInlineOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetInlineOk() (*IoArgoprojWorkflowV1alpha1Template, bool)`

GetInlineOk returns a tuple with the Inline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInline

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) SetInline(v IoArgoprojWorkflowV1alpha1Template)`

SetInline sets Inline field to given value.

### HasInline

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) HasInline() bool`

HasInline returns a boolean if a field has been set.

### GetName

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnExit

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetOnExit() string`

GetOnExit returns the OnExit field if non-nil, zero value otherwise.

### GetOnExitOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetOnExitOk() (*string, bool)`

GetOnExitOk returns a tuple with the OnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnExit

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) SetOnExit(v string)`

SetOnExit sets OnExit field to given value.

### HasOnExit

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) HasOnExit() bool`

HasOnExit returns a boolean if a field has been set.

### GetTemplate

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetTemplateRef() IoArgoprojWorkflowV1alpha1TemplateRef`

GetTemplateRef returns the TemplateRef field if non-nil, zero value otherwise.

### GetTemplateRefOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetTemplateRefOk() (*IoArgoprojWorkflowV1alpha1TemplateRef, bool)`

GetTemplateRefOk returns a tuple with the TemplateRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) SetTemplateRef(v IoArgoprojWorkflowV1alpha1TemplateRef)`

SetTemplateRef sets TemplateRef field to given value.

### HasTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) HasTemplateRef() bool`

HasTemplateRef returns a boolean if a field has been set.

### GetWhen

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetWhen() string`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetWhenOk() (*string, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) SetWhen(v string)`

SetWhen sets When field to given value.

### HasWhen

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetWithItems

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetWithItems() []map[string]interface{}`

GetWithItems returns the WithItems field if non-nil, zero value otherwise.

### GetWithItemsOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetWithItemsOk() (*[]map[string]interface{}, bool)`

GetWithItemsOk returns a tuple with the WithItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithItems

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) SetWithItems(v []map[string]interface{})`

SetWithItems sets WithItems field to given value.

### HasWithItems

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) HasWithItems() bool`

HasWithItems returns a boolean if a field has been set.

### GetWithParam

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetWithParam() string`

GetWithParam returns the WithParam field if non-nil, zero value otherwise.

### GetWithParamOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetWithParamOk() (*string, bool)`

GetWithParamOk returns a tuple with the WithParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithParam

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) SetWithParam(v string)`

SetWithParam sets WithParam field to given value.

### HasWithParam

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) HasWithParam() bool`

HasWithParam returns a boolean if a field has been set.

### GetWithSequence

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetWithSequence() IoArgoprojWorkflowV1alpha1Sequence`

GetWithSequence returns the WithSequence field if non-nil, zero value otherwise.

### GetWithSequenceOk

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) GetWithSequenceOk() (*IoArgoprojWorkflowV1alpha1Sequence, bool)`

GetWithSequenceOk returns a tuple with the WithSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithSequence

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) SetWithSequence(v IoArgoprojWorkflowV1alpha1Sequence)`

SetWithSequence sets WithSequence field to given value.

### HasWithSequence

`func (o *IoArgoprojWorkflowV1alpha1WorkflowStep) HasWithSequence() bool`

HasWithSequence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


