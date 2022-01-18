# IoArgoprojEventsV1alpha1SensorSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Circuit** | Pointer to **string** | Circuit is a boolean expression of dependency groups Deprecated: will be removed in v1.5, use Switch in triggers instead. | [optional] 
**Dependencies** | Pointer to [**[]IoArgoprojEventsV1alpha1EventDependency**](IoArgoprojEventsV1alpha1EventDependency.md) | Dependencies is a list of the events that this sensor is dependent on. | [optional] 
**DependencyGroups** | Pointer to [**[]IoArgoprojEventsV1alpha1DependencyGroup**](IoArgoprojEventsV1alpha1DependencyGroup.md) | DependencyGroups is a list of the groups of events. | [optional] 
**ErrorOnFailedRound** | Pointer to **bool** | ErrorOnFailedRound if set to true, marks sensor state as &#x60;error&#x60; if the previous trigger round fails. Once sensor state is set to &#x60;error&#x60;, no further triggers will be processed. | [optional] 
**EventBusName** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Template** | Pointer to [**IoArgoprojEventsV1alpha1Template**](IoArgoprojEventsV1alpha1Template.md) |  | [optional] 
**Triggers** | Pointer to [**[]IoArgoprojEventsV1alpha1Trigger**](IoArgoprojEventsV1alpha1Trigger.md) | Triggers is a list of the things that this sensor evokes. These are the outputs from this sensor. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1SensorSpec

`func NewIoArgoprojEventsV1alpha1SensorSpec() *IoArgoprojEventsV1alpha1SensorSpec`

NewIoArgoprojEventsV1alpha1SensorSpec instantiates a new IoArgoprojEventsV1alpha1SensorSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1SensorSpecWithDefaults

`func NewIoArgoprojEventsV1alpha1SensorSpecWithDefaults() *IoArgoprojEventsV1alpha1SensorSpec`

NewIoArgoprojEventsV1alpha1SensorSpecWithDefaults instantiates a new IoArgoprojEventsV1alpha1SensorSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuit

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetCircuit() string`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetCircuitOk() (*string, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *IoArgoprojEventsV1alpha1SensorSpec) SetCircuit(v string)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *IoArgoprojEventsV1alpha1SensorSpec) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetDependencies

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetDependencies() []IoArgoprojEventsV1alpha1EventDependency`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetDependenciesOk() (*[]IoArgoprojEventsV1alpha1EventDependency, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *IoArgoprojEventsV1alpha1SensorSpec) SetDependencies(v []IoArgoprojEventsV1alpha1EventDependency)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *IoArgoprojEventsV1alpha1SensorSpec) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetDependencyGroups

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetDependencyGroups() []IoArgoprojEventsV1alpha1DependencyGroup`

GetDependencyGroups returns the DependencyGroups field if non-nil, zero value otherwise.

### GetDependencyGroupsOk

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetDependencyGroupsOk() (*[]IoArgoprojEventsV1alpha1DependencyGroup, bool)`

GetDependencyGroupsOk returns a tuple with the DependencyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyGroups

`func (o *IoArgoprojEventsV1alpha1SensorSpec) SetDependencyGroups(v []IoArgoprojEventsV1alpha1DependencyGroup)`

SetDependencyGroups sets DependencyGroups field to given value.

### HasDependencyGroups

`func (o *IoArgoprojEventsV1alpha1SensorSpec) HasDependencyGroups() bool`

HasDependencyGroups returns a boolean if a field has been set.

### GetErrorOnFailedRound

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetErrorOnFailedRound() bool`

GetErrorOnFailedRound returns the ErrorOnFailedRound field if non-nil, zero value otherwise.

### GetErrorOnFailedRoundOk

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetErrorOnFailedRoundOk() (*bool, bool)`

GetErrorOnFailedRoundOk returns a tuple with the ErrorOnFailedRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorOnFailedRound

`func (o *IoArgoprojEventsV1alpha1SensorSpec) SetErrorOnFailedRound(v bool)`

SetErrorOnFailedRound sets ErrorOnFailedRound field to given value.

### HasErrorOnFailedRound

`func (o *IoArgoprojEventsV1alpha1SensorSpec) HasErrorOnFailedRound() bool`

HasErrorOnFailedRound returns a boolean if a field has been set.

### GetEventBusName

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetEventBusName() string`

GetEventBusName returns the EventBusName field if non-nil, zero value otherwise.

### GetEventBusNameOk

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetEventBusNameOk() (*string, bool)`

GetEventBusNameOk returns a tuple with the EventBusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventBusName

`func (o *IoArgoprojEventsV1alpha1SensorSpec) SetEventBusName(v string)`

SetEventBusName sets EventBusName field to given value.

### HasEventBusName

`func (o *IoArgoprojEventsV1alpha1SensorSpec) HasEventBusName() bool`

HasEventBusName returns a boolean if a field has been set.

### GetReplicas

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *IoArgoprojEventsV1alpha1SensorSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *IoArgoprojEventsV1alpha1SensorSpec) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetTemplate

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetTemplate() IoArgoprojEventsV1alpha1Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetTemplateOk() (*IoArgoprojEventsV1alpha1Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoArgoprojEventsV1alpha1SensorSpec) SetTemplate(v IoArgoprojEventsV1alpha1Template)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IoArgoprojEventsV1alpha1SensorSpec) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTriggers

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetTriggers() []IoArgoprojEventsV1alpha1Trigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *IoArgoprojEventsV1alpha1SensorSpec) GetTriggersOk() (*[]IoArgoprojEventsV1alpha1Trigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *IoArgoprojEventsV1alpha1SensorSpec) SetTriggers(v []IoArgoprojEventsV1alpha1Trigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *IoArgoprojEventsV1alpha1SensorSpec) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


