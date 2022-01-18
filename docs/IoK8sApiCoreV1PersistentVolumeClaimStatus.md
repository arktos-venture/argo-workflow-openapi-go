# IoK8sApiCoreV1PersistentVolumeClaimStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to **[]string** | AccessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | [optional] 
**Capacity** | Pointer to **map[string]string** | Represents the actual resources of the underlying volume. | [optional] 
**Conditions** | Pointer to [**[]IoK8sApiCoreV1PersistentVolumeClaimCondition**](IoK8sApiCoreV1PersistentVolumeClaimCondition.md) | Current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to &#39;ResizeStarted&#39;. | [optional] 
**Phase** | Pointer to **string** | Phase represents the current phase of PersistentVolumeClaim. | [optional] 

## Methods

### NewIoK8sApiCoreV1PersistentVolumeClaimStatus

`func NewIoK8sApiCoreV1PersistentVolumeClaimStatus() *IoK8sApiCoreV1PersistentVolumeClaimStatus`

NewIoK8sApiCoreV1PersistentVolumeClaimStatus instantiates a new IoK8sApiCoreV1PersistentVolumeClaimStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1PersistentVolumeClaimStatusWithDefaults

`func NewIoK8sApiCoreV1PersistentVolumeClaimStatusWithDefaults() *IoK8sApiCoreV1PersistentVolumeClaimStatus`

NewIoK8sApiCoreV1PersistentVolumeClaimStatusWithDefaults instantiates a new IoK8sApiCoreV1PersistentVolumeClaimStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetCapacity

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) GetCapacity() map[string]string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) GetCapacityOk() (*map[string]string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) SetCapacity(v map[string]string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetConditions

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) GetConditions() []IoK8sApiCoreV1PersistentVolumeClaimCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) GetConditionsOk() (*[]IoK8sApiCoreV1PersistentVolumeClaimCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) SetConditions(v []IoK8sApiCoreV1PersistentVolumeClaimCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPhase

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *IoK8sApiCoreV1PersistentVolumeClaimStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


