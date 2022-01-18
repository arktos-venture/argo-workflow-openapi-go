# IoArgoprojWorkflowV1alpha1Parameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **string** | Default is the default value to use for an input parameter if a value was not supplied | [optional] 
**Enum** | Pointer to **[]string** | Enum holds a list of string values to choose from, for the actual value of the parameter | [optional] 
**GlobalName** | Pointer to **string** | GlobalName exports an output parameter to the global scope, making it available as &#39;{{io.argoproj.workflow.v1alpha1.outputs.parameters.XXXX}} and in workflow.status.outputs.parameters | [optional] 
**Name** | **string** | Name is the parameter name | 
**Value** | Pointer to **string** | Value is the literal value to use for the parameter. If specified in the context of an input parameter, the value takes precedence over any passed values | [optional] 
**ValueFrom** | Pointer to [**IoArgoprojWorkflowV1alpha1ValueFrom**](IoArgoprojWorkflowV1alpha1ValueFrom.md) |  | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1Parameter

`func NewIoArgoprojWorkflowV1alpha1Parameter(name string, ) *IoArgoprojWorkflowV1alpha1Parameter`

NewIoArgoprojWorkflowV1alpha1Parameter instantiates a new IoArgoprojWorkflowV1alpha1Parameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1ParameterWithDefaults

`func NewIoArgoprojWorkflowV1alpha1ParameterWithDefaults() *IoArgoprojWorkflowV1alpha1Parameter`

NewIoArgoprojWorkflowV1alpha1ParameterWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Parameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *IoArgoprojWorkflowV1alpha1Parameter) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *IoArgoprojWorkflowV1alpha1Parameter) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *IoArgoprojWorkflowV1alpha1Parameter) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *IoArgoprojWorkflowV1alpha1Parameter) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnum

`func (o *IoArgoprojWorkflowV1alpha1Parameter) GetEnum() []string`

GetEnum returns the Enum field if non-nil, zero value otherwise.

### GetEnumOk

`func (o *IoArgoprojWorkflowV1alpha1Parameter) GetEnumOk() (*[]string, bool)`

GetEnumOk returns a tuple with the Enum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnum

`func (o *IoArgoprojWorkflowV1alpha1Parameter) SetEnum(v []string)`

SetEnum sets Enum field to given value.

### HasEnum

`func (o *IoArgoprojWorkflowV1alpha1Parameter) HasEnum() bool`

HasEnum returns a boolean if a field has been set.

### GetGlobalName

`func (o *IoArgoprojWorkflowV1alpha1Parameter) GetGlobalName() string`

GetGlobalName returns the GlobalName field if non-nil, zero value otherwise.

### GetGlobalNameOk

`func (o *IoArgoprojWorkflowV1alpha1Parameter) GetGlobalNameOk() (*string, bool)`

GetGlobalNameOk returns a tuple with the GlobalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalName

`func (o *IoArgoprojWorkflowV1alpha1Parameter) SetGlobalName(v string)`

SetGlobalName sets GlobalName field to given value.

### HasGlobalName

`func (o *IoArgoprojWorkflowV1alpha1Parameter) HasGlobalName() bool`

HasGlobalName returns a boolean if a field has been set.

### GetName

`func (o *IoArgoprojWorkflowV1alpha1Parameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojWorkflowV1alpha1Parameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojWorkflowV1alpha1Parameter) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *IoArgoprojWorkflowV1alpha1Parameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IoArgoprojWorkflowV1alpha1Parameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IoArgoprojWorkflowV1alpha1Parameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IoArgoprojWorkflowV1alpha1Parameter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueFrom

`func (o *IoArgoprojWorkflowV1alpha1Parameter) GetValueFrom() IoArgoprojWorkflowV1alpha1ValueFrom`

GetValueFrom returns the ValueFrom field if non-nil, zero value otherwise.

### GetValueFromOk

`func (o *IoArgoprojWorkflowV1alpha1Parameter) GetValueFromOk() (*IoArgoprojWorkflowV1alpha1ValueFrom, bool)`

GetValueFromOk returns a tuple with the ValueFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFrom

`func (o *IoArgoprojWorkflowV1alpha1Parameter) SetValueFrom(v IoArgoprojWorkflowV1alpha1ValueFrom)`

SetValueFrom sets ValueFrom field to given value.

### HasValueFrom

`func (o *IoArgoprojWorkflowV1alpha1Parameter) HasValueFrom() bool`

HasValueFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


