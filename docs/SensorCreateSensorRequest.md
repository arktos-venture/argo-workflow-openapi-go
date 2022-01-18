# SensorCreateSensorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateOptions** | Pointer to [**IoK8sApimachineryPkgApisMetaV1CreateOptions**](IoK8sApimachineryPkgApisMetaV1CreateOptions.md) |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Sensor** | Pointer to [**IoArgoprojEventsV1alpha1Sensor**](IoArgoprojEventsV1alpha1Sensor.md) |  | [optional] 

## Methods

### NewSensorCreateSensorRequest

`func NewSensorCreateSensorRequest() *SensorCreateSensorRequest`

NewSensorCreateSensorRequest instantiates a new SensorCreateSensorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorCreateSensorRequestWithDefaults

`func NewSensorCreateSensorRequestWithDefaults() *SensorCreateSensorRequest`

NewSensorCreateSensorRequestWithDefaults instantiates a new SensorCreateSensorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateOptions

`func (o *SensorCreateSensorRequest) GetCreateOptions() IoK8sApimachineryPkgApisMetaV1CreateOptions`

GetCreateOptions returns the CreateOptions field if non-nil, zero value otherwise.

### GetCreateOptionsOk

`func (o *SensorCreateSensorRequest) GetCreateOptionsOk() (*IoK8sApimachineryPkgApisMetaV1CreateOptions, bool)`

GetCreateOptionsOk returns a tuple with the CreateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOptions

`func (o *SensorCreateSensorRequest) SetCreateOptions(v IoK8sApimachineryPkgApisMetaV1CreateOptions)`

SetCreateOptions sets CreateOptions field to given value.

### HasCreateOptions

`func (o *SensorCreateSensorRequest) HasCreateOptions() bool`

HasCreateOptions returns a boolean if a field has been set.

### GetNamespace

`func (o *SensorCreateSensorRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SensorCreateSensorRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SensorCreateSensorRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SensorCreateSensorRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSensor

`func (o *SensorCreateSensorRequest) GetSensor() IoArgoprojEventsV1alpha1Sensor`

GetSensor returns the Sensor field if non-nil, zero value otherwise.

### GetSensorOk

`func (o *SensorCreateSensorRequest) GetSensorOk() (*IoArgoprojEventsV1alpha1Sensor, bool)`

GetSensorOk returns a tuple with the Sensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensor

`func (o *SensorCreateSensorRequest) SetSensor(v IoArgoprojEventsV1alpha1Sensor)`

SetSensor sets Sensor field to given value.

### HasSensor

`func (o *SensorCreateSensorRequest) HasSensor() bool`

HasSensor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


