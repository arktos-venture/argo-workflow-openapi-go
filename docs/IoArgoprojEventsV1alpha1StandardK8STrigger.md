# IoArgoprojEventsV1alpha1StandardK8STrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupVersionResource** | Pointer to [**IoK8sApimachineryPkgApisMetaV1GroupVersionResource**](IoK8sApimachineryPkgApisMetaV1GroupVersionResource.md) |  | [optional] 
**LiveObject** | Pointer to **bool** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]IoArgoprojEventsV1alpha1TriggerParameter**](IoArgoprojEventsV1alpha1TriggerParameter.md) | Parameters is the list of parameters that is applied to resolved K8s trigger object. | [optional] 
**PatchStrategy** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**IoArgoprojEventsV1alpha1ArtifactLocation**](IoArgoprojEventsV1alpha1ArtifactLocation.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1StandardK8STrigger

`func NewIoArgoprojEventsV1alpha1StandardK8STrigger() *IoArgoprojEventsV1alpha1StandardK8STrigger`

NewIoArgoprojEventsV1alpha1StandardK8STrigger instantiates a new IoArgoprojEventsV1alpha1StandardK8STrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1StandardK8STriggerWithDefaults

`func NewIoArgoprojEventsV1alpha1StandardK8STriggerWithDefaults() *IoArgoprojEventsV1alpha1StandardK8STrigger`

NewIoArgoprojEventsV1alpha1StandardK8STriggerWithDefaults instantiates a new IoArgoprojEventsV1alpha1StandardK8STrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupVersionResource

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetGroupVersionResource() IoK8sApimachineryPkgApisMetaV1GroupVersionResource`

GetGroupVersionResource returns the GroupVersionResource field if non-nil, zero value otherwise.

### GetGroupVersionResourceOk

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetGroupVersionResourceOk() (*IoK8sApimachineryPkgApisMetaV1GroupVersionResource, bool)`

GetGroupVersionResourceOk returns a tuple with the GroupVersionResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupVersionResource

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) SetGroupVersionResource(v IoK8sApimachineryPkgApisMetaV1GroupVersionResource)`

SetGroupVersionResource sets GroupVersionResource field to given value.

### HasGroupVersionResource

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) HasGroupVersionResource() bool`

HasGroupVersionResource returns a boolean if a field has been set.

### GetLiveObject

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetLiveObject() bool`

GetLiveObject returns the LiveObject field if non-nil, zero value otherwise.

### GetLiveObjectOk

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetLiveObjectOk() (*bool, bool)`

GetLiveObjectOk returns a tuple with the LiveObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveObject

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) SetLiveObject(v bool)`

SetLiveObject sets LiveObject field to given value.

### HasLiveObject

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) HasLiveObject() bool`

HasLiveObject returns a boolean if a field has been set.

### GetOperation

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetParameters

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetParameters() []IoArgoprojEventsV1alpha1TriggerParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetParametersOk() (*[]IoArgoprojEventsV1alpha1TriggerParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) SetParameters(v []IoArgoprojEventsV1alpha1TriggerParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPatchStrategy

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetPatchStrategy() string`

GetPatchStrategy returns the PatchStrategy field if non-nil, zero value otherwise.

### GetPatchStrategyOk

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetPatchStrategyOk() (*string, bool)`

GetPatchStrategyOk returns a tuple with the PatchStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchStrategy

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) SetPatchStrategy(v string)`

SetPatchStrategy sets PatchStrategy field to given value.

### HasPatchStrategy

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) HasPatchStrategy() bool`

HasPatchStrategy returns a boolean if a field has been set.

### GetSource

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetSource() IoArgoprojEventsV1alpha1ArtifactLocation`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) GetSourceOk() (*IoArgoprojEventsV1alpha1ArtifactLocation, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) SetSource(v IoArgoprojEventsV1alpha1ArtifactLocation)`

SetSource sets Source field to given value.

### HasSource

`func (o *IoArgoprojEventsV1alpha1StandardK8STrigger) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


