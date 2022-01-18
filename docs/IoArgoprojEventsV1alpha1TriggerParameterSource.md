# IoArgoprojEventsV1alpha1TriggerParameterSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextKey** | Pointer to **string** | ContextKey is the JSONPath of the event&#39;s (JSON decoded) context key ContextKey is a series of keys separated by a dot. A key may contain wildcard characters &#39;*&#39; and &#39;?&#39;. To access an array value use the index as the key. The dot and wildcard characters can be escaped with &#39;\\\\&#39;. See https://github.com/tidwall/gjson#path-syntax for more information on how to use this. | [optional] 
**ContextTemplate** | Pointer to **string** |  | [optional] 
**DataKey** | Pointer to **string** | DataKey is the JSONPath of the event&#39;s (JSON decoded) data key DataKey is a series of keys separated by a dot. A key may contain wildcard characters &#39;*&#39; and &#39;?&#39;. To access an array value use the index as the key. The dot and wildcard characters can be escaped with &#39;\\\\&#39;. See https://github.com/tidwall/gjson#path-syntax for more information on how to use this. | [optional] 
**DataTemplate** | Pointer to **string** |  | [optional] 
**DependencyName** | Pointer to **string** | DependencyName refers to the name of the dependency. The event which is stored for this dependency is used as payload for the parameterization. Make sure to refer to one of the dependencies you have defined under Dependencies list. | [optional] 
**Value** | Pointer to **string** | Value is the default literal value to use for this parameter source This is only used if the DataKey is invalid. If the DataKey is invalid and this is not defined, this param source will produce an error. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1TriggerParameterSource

`func NewIoArgoprojEventsV1alpha1TriggerParameterSource() *IoArgoprojEventsV1alpha1TriggerParameterSource`

NewIoArgoprojEventsV1alpha1TriggerParameterSource instantiates a new IoArgoprojEventsV1alpha1TriggerParameterSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1TriggerParameterSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1TriggerParameterSourceWithDefaults() *IoArgoprojEventsV1alpha1TriggerParameterSource`

NewIoArgoprojEventsV1alpha1TriggerParameterSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1TriggerParameterSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextKey

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) GetContextKey() string`

GetContextKey returns the ContextKey field if non-nil, zero value otherwise.

### GetContextKeyOk

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) GetContextKeyOk() (*string, bool)`

GetContextKeyOk returns a tuple with the ContextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKey

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) SetContextKey(v string)`

SetContextKey sets ContextKey field to given value.

### HasContextKey

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) HasContextKey() bool`

HasContextKey returns a boolean if a field has been set.

### GetContextTemplate

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) GetContextTemplate() string`

GetContextTemplate returns the ContextTemplate field if non-nil, zero value otherwise.

### GetContextTemplateOk

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) GetContextTemplateOk() (*string, bool)`

GetContextTemplateOk returns a tuple with the ContextTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextTemplate

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) SetContextTemplate(v string)`

SetContextTemplate sets ContextTemplate field to given value.

### HasContextTemplate

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) HasContextTemplate() bool`

HasContextTemplate returns a boolean if a field has been set.

### GetDataKey

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) GetDataKey() string`

GetDataKey returns the DataKey field if non-nil, zero value otherwise.

### GetDataKeyOk

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) GetDataKeyOk() (*string, bool)`

GetDataKeyOk returns a tuple with the DataKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataKey

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) SetDataKey(v string)`

SetDataKey sets DataKey field to given value.

### HasDataKey

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) HasDataKey() bool`

HasDataKey returns a boolean if a field has been set.

### GetDataTemplate

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) GetDataTemplate() string`

GetDataTemplate returns the DataTemplate field if non-nil, zero value otherwise.

### GetDataTemplateOk

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) GetDataTemplateOk() (*string, bool)`

GetDataTemplateOk returns a tuple with the DataTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTemplate

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) SetDataTemplate(v string)`

SetDataTemplate sets DataTemplate field to given value.

### HasDataTemplate

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) HasDataTemplate() bool`

HasDataTemplate returns a boolean if a field has been set.

### GetDependencyName

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) GetDependencyName() string`

GetDependencyName returns the DependencyName field if non-nil, zero value otherwise.

### GetDependencyNameOk

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) GetDependencyNameOk() (*string, bool)`

GetDependencyNameOk returns a tuple with the DependencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyName

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) SetDependencyName(v string)`

SetDependencyName sets DependencyName field to given value.

### HasDependencyName

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) HasDependencyName() bool`

HasDependencyName returns a boolean if a field has been set.

### GetValue

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IoArgoprojEventsV1alpha1TriggerParameterSource) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


