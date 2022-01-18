# IoArgoprojWorkflowV1alpha1ValueFrom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMapKeyRef** | Pointer to [**IoK8sApiCoreV1ConfigMapKeySelector**](IoK8sApiCoreV1ConfigMapKeySelector.md) |  | [optional] 
**Default** | Pointer to **string** | Default specifies a value to be used if retrieving the value from the specified source fails | [optional] 
**Event** | Pointer to **string** | Selector (https://github.com/antonmedv/expr) that is evaluated against the event to get the value of the parameter. E.g. &#x60;payload.message&#x60; | [optional] 
**Expression** | Pointer to **string** | Expression, if defined, is evaluated to specify the value for the parameter | [optional] 
**JqFilter** | Pointer to **string** | JQFilter expression against the resource object in resource templates | [optional] 
**JsonPath** | Pointer to **string** | JSONPath of a resource to retrieve an output parameter value from in resource templates | [optional] 
**Parameter** | Pointer to **string** | Parameter reference to a step or dag task in which to retrieve an output parameter value from (e.g. &#39;{{steps.mystep.outputs.myparam}}&#39;) | [optional] 
**Path** | Pointer to **string** | Path in the container to retrieve an output parameter value from in container templates | [optional] 
**Supplied** | Pointer to **map[string]interface{}** | SuppliedValueFrom is a placeholder for a value to be filled in directly, either through the CLI, API, etc. | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1ValueFrom

`func NewIoArgoprojWorkflowV1alpha1ValueFrom() *IoArgoprojWorkflowV1alpha1ValueFrom`

NewIoArgoprojWorkflowV1alpha1ValueFrom instantiates a new IoArgoprojWorkflowV1alpha1ValueFrom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ValueFromWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ValueFromWithDefaults() *IoArgoprojWorkflowV1alpha1ValueFrom`

NewIoArgoprojWorkflowV1alpha1ValueFromWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ValueFrom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMapKeyRef

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetConfigMapKeyRef() IoK8sApiCoreV1ConfigMapKeySelector`

GetConfigMapKeyRef returns the ConfigMapKeyRef field if non-nil, zero value otherwise.

### GetConfigMapKeyRefOk

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetConfigMapKeyRefOk() (*IoK8sApiCoreV1ConfigMapKeySelector, bool)`

GetConfigMapKeyRefOk returns a tuple with the ConfigMapKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapKeyRef

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) SetConfigMapKeyRef(v IoK8sApiCoreV1ConfigMapKeySelector)`

SetConfigMapKeyRef sets ConfigMapKeyRef field to given value.

### HasConfigMapKeyRef

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) HasConfigMapKeyRef() bool`

HasConfigMapKeyRef returns a boolean if a field has been set.

### GetDefault

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEvent

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetExpression

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetJqFilter

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetJqFilter() string`

GetJqFilter returns the JqFilter field if non-nil, zero value otherwise.

### GetJqFilterOk

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetJqFilterOk() (*string, bool)`

GetJqFilterOk returns a tuple with the JqFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJqFilter

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) SetJqFilter(v string)`

SetJqFilter sets JqFilter field to given value.

### HasJqFilter

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) HasJqFilter() bool`

HasJqFilter returns a boolean if a field has been set.

### GetJsonPath

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetJsonPath() string`

GetJsonPath returns the JsonPath field if non-nil, zero value otherwise.

### GetJsonPathOk

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetJsonPathOk() (*string, bool)`

GetJsonPathOk returns a tuple with the JsonPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonPath

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) SetJsonPath(v string)`

SetJsonPath sets JsonPath field to given value.

### HasJsonPath

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) HasJsonPath() bool`

HasJsonPath returns a boolean if a field has been set.

### GetParameter

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetPath

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSupplied

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetSupplied() map[string]interface{}`

GetSupplied returns the Supplied field if non-nil, zero value otherwise.

### GetSuppliedOk

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) GetSuppliedOk() (*map[string]interface{}, bool)`

GetSuppliedOk returns a tuple with the Supplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplied

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) SetSupplied(v map[string]interface{})`

SetSupplied sets Supplied field to given value.

### HasSupplied

`func (o *IoArgoprojWorkflowV1alpha1ValueFrom) HasSupplied() bool`

HasSupplied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


