# IoArgoprojWorkflowV1alpha1DAGTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to [**IoArgoprojWorkflowV1alpha1Arguments**](IoArgoprojWorkflowV1alpha1Arguments.md) |  | [optional] 
**ContinueOn** | Pointer to [**IoArgoprojWorkflowV1alpha1ContinueOn**](IoArgoprojWorkflowV1alpha1ContinueOn.md) |  | [optional] 
**Dependencies** | Pointer to **[]string** | Dependencies are name of other targets which this depends on | [optional] 
**Depends** | Pointer to **string** | Depends are name of other targets which this depends on | [optional] 
**Hooks** | Pointer to [**map[string]IoArgoprojWorkflowV1alpha1LifecycleHook**](IoArgoprojWorkflowV1alpha1LifecycleHook.md) | Hooks hold the lifecycle hook which is invoked at lifecycle of task, irrespective of the success, failure, or error status of the primary task | [optional] 
**Inline** | Pointer to [**IoArgoprojWorkflowV1alpha1Template**](IoArgoprojWorkflowV1alpha1Template.md) |  | [optional] 
**Name** | **string** | Name is the name of the target | 
**OnExit** | Pointer to **string** | OnExit is a template reference which is invoked at the end of the template, irrespective of the success, failure, or error of the primary template. DEPRECATED: Use Hooks[exit].Template instead. | [optional] 
**Template** | Pointer to **string** | Name of template to execute | [optional] 
**TemplateRef** | Pointer to [**IoArgoprojWorkflowV1alpha1TemplateRef**](IoArgoprojWorkflowV1alpha1TemplateRef.md) |  | [optional] 
**When** | Pointer to **string** | When is an expression in which the task should conditionally execute | [optional] 
**WithItems** | Pointer to **[]map[string]interface{}** | WithItems expands a task into multiple parallel tasks from the items in the list | [optional] 
**WithParam** | Pointer to **string** | WithParam expands a task into multiple parallel tasks from the value in the parameter, which is expected to be a JSON list. | [optional] 
**WithSequence** | Pointer to [**IoArgoprojWorkflowV1alpha1Sequence**](IoArgoprojWorkflowV1alpha1Sequence.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1DAGTask

`func NewIoArgoprojWorkflowV1alpha1DAGTask(name string, ) *IoArgoprojWorkflowV1alpha1DAGTask`

NewIoArgoprojWorkflowV1alpha1DAGTask instantiates a new IoArgoprojWorkflowV1alpha1DAGTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1DAGTaskWithDefaults

`func NewIoArgoprojWorkflowV1alpha1DAGTaskWithDefaults() *IoArgoprojWorkflowV1alpha1DAGTask`

NewIoArgoprojWorkflowV1alpha1DAGTaskWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1DAGTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetArguments() IoArgoprojWorkflowV1alpha1Arguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetArgumentsOk() (*IoArgoprojWorkflowV1alpha1Arguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetArguments(v IoArgoprojWorkflowV1alpha1Arguments)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetContinueOn

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetContinueOn() IoArgoprojWorkflowV1alpha1ContinueOn`

GetContinueOn returns the ContinueOn field if non-nil, zero value otherwise.

### GetContinueOnOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetContinueOnOk() (*IoArgoprojWorkflowV1alpha1ContinueOn, bool)`

GetContinueOnOk returns a tuple with the ContinueOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOn

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetContinueOn(v IoArgoprojWorkflowV1alpha1ContinueOn)`

SetContinueOn sets ContinueOn field to given value.

### HasContinueOn

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasContinueOn() bool`

HasContinueOn returns a boolean if a field has been set.

### GetDependencies

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetDependencies() []string`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetDependenciesOk() (*[]string, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetDependencies(v []string)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetDepends

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetDepends() string`

GetDepends returns the Depends field if non-nil, zero value otherwise.

### GetDependsOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetDependsOk() (*string, bool)`

GetDependsOk returns a tuple with the Depends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepends

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetDepends(v string)`

SetDepends sets Depends field to given value.

### HasDepends

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasDepends() bool`

HasDepends returns a boolean if a field has been set.

### GetHooks

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetHooks() map[string]IoArgoprojWorkflowV1alpha1LifecycleHook`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetHooksOk() (*map[string]IoArgoprojWorkflowV1alpha1LifecycleHook, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetHooks(v map[string]IoArgoprojWorkflowV1alpha1LifecycleHook)`

SetHooks sets Hooks field to given value.

### HasHooks

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasHooks() bool`

HasHooks returns a boolean if a field has been set.

### GetInline

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetInline() IoArgoprojWorkflowV1alpha1Template`

GetInline returns the Inline field if non-nil, zero value otherwise.

### GetInlineOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetInlineOk() (*IoArgoprojWorkflowV1alpha1Template, bool)`

GetInlineOk returns a tuple with the Inline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInline

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetInline(v IoArgoprojWorkflowV1alpha1Template)`

SetInline sets Inline field to given value.

### HasInline

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasInline() bool`

HasInline returns a boolean if a field has been set.

### GetName

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetName(v string)`

SetName sets Name field to given value.


### GetOnExit

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetOnExit() string`

GetOnExit returns the OnExit field if non-nil, zero value otherwise.

### GetOnExitOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetOnExitOk() (*string, bool)`

GetOnExitOk returns a tuple with the OnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnExit

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetOnExit(v string)`

SetOnExit sets OnExit field to given value.

### HasOnExit

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasOnExit() bool`

HasOnExit returns a boolean if a field has been set.

### GetTemplate

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetTemplateRef() IoArgoprojWorkflowV1alpha1TemplateRef`

GetTemplateRef returns the TemplateRef field if non-nil, zero value otherwise.

### GetTemplateRefOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetTemplateRefOk() (*IoArgoprojWorkflowV1alpha1TemplateRef, bool)`

GetTemplateRefOk returns a tuple with the TemplateRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetTemplateRef(v IoArgoprojWorkflowV1alpha1TemplateRef)`

SetTemplateRef sets TemplateRef field to given value.

### HasTemplateRef

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasTemplateRef() bool`

HasTemplateRef returns a boolean if a field has been set.

### GetWhen

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWhen() string`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWhenOk() (*string, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetWhen(v string)`

SetWhen sets When field to given value.

### HasWhen

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetWithItems

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWithItems() []map[string]interface{}`

GetWithItems returns the WithItems field if non-nil, zero value otherwise.

### GetWithItemsOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWithItemsOk() (*[]map[string]interface{}, bool)`

GetWithItemsOk returns a tuple with the WithItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithItems

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetWithItems(v []map[string]interface{})`

SetWithItems sets WithItems field to given value.

### HasWithItems

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasWithItems() bool`

HasWithItems returns a boolean if a field has been set.

### GetWithParam

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWithParam() string`

GetWithParam returns the WithParam field if non-nil, zero value otherwise.

### GetWithParamOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWithParamOk() (*string, bool)`

GetWithParamOk returns a tuple with the WithParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithParam

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetWithParam(v string)`

SetWithParam sets WithParam field to given value.

### HasWithParam

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasWithParam() bool`

HasWithParam returns a boolean if a field has been set.

### GetWithSequence

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWithSequence() IoArgoprojWorkflowV1alpha1Sequence`

GetWithSequence returns the WithSequence field if non-nil, zero value otherwise.

### GetWithSequenceOk

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) GetWithSequenceOk() (*IoArgoprojWorkflowV1alpha1Sequence, bool)`

GetWithSequenceOk returns a tuple with the WithSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithSequence

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) SetWithSequence(v IoArgoprojWorkflowV1alpha1Sequence)`

SetWithSequence sets WithSequence field to given value.

### HasWithSequence

`func (o *IoArgoprojWorkflowV1alpha1DAGTask) HasWithSequence() bool`

HasWithSequence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


