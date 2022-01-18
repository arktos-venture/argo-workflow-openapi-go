# IoArgoprojEventsV1alpha1Sensor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**IoK8sApimachineryPkgApisMetaV1ObjectMeta**](IoK8sApimachineryPkgApisMetaV1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**IoArgoprojEventsV1alpha1SensorSpec**](IoArgoprojEventsV1alpha1SensorSpec.md) |  | [optional] 
**Status** | Pointer to [**IoArgoprojEventsV1alpha1SensorStatus**](IoArgoprojEventsV1alpha1SensorStatus.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1Sensor

`func NewIoArgoprojEventsV1alpha1Sensor() *IoArgoprojEventsV1alpha1Sensor`

NewIoArgoprojEventsV1alpha1Sensor instantiates a new IoArgoprojEventsV1alpha1Sensor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1SensorWithDefaults

`func NewIoArgoprojEventsV1alpha1SensorWithDefaults() *IoArgoprojEventsV1alpha1Sensor`

NewIoArgoprojEventsV1alpha1SensorWithDefaults instantiates a new IoArgoprojEventsV1alpha1Sensor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *IoArgoprojEventsV1alpha1Sensor) GetMetadata() IoK8sApimachineryPkgApisMetaV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IoArgoprojEventsV1alpha1Sensor) GetMetadataOk() (*IoK8sApimachineryPkgApisMetaV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IoArgoprojEventsV1alpha1Sensor) SetMetadata(v IoK8sApimachineryPkgApisMetaV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IoArgoprojEventsV1alpha1Sensor) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *IoArgoprojEventsV1alpha1Sensor) GetSpec() IoArgoprojEventsV1alpha1SensorSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *IoArgoprojEventsV1alpha1Sensor) GetSpecOk() (*IoArgoprojEventsV1alpha1SensorSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *IoArgoprojEventsV1alpha1Sensor) SetSpec(v IoArgoprojEventsV1alpha1SensorSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *IoArgoprojEventsV1alpha1Sensor) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *IoArgoprojEventsV1alpha1Sensor) GetStatus() IoArgoprojEventsV1alpha1SensorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IoArgoprojEventsV1alpha1Sensor) GetStatusOk() (*IoArgoprojEventsV1alpha1SensorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IoArgoprojEventsV1alpha1Sensor) SetStatus(v IoArgoprojEventsV1alpha1SensorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IoArgoprojEventsV1alpha1Sensor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


