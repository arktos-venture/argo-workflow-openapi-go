# SensorUpdateSensorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Sensor** | Pointer to [**IoArgoprojEventsV1alpha1Sensor**](IoArgoprojEventsV1alpha1Sensor.md) |  | [optional] 

## Methods

### NewSensorUpdateSensorRequest

`func NewSensorUpdateSensorRequest() *SensorUpdateSensorRequest`

NewSensorUpdateSensorRequest instantiates a new SensorUpdateSensorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorUpdateSensorRequestWithDefaults

`func NewSensorUpdateSensorRequestWithDefaults() *SensorUpdateSensorRequest`

NewSensorUpdateSensorRequestWithDefaults instantiates a new SensorUpdateSensorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SensorUpdateSensorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SensorUpdateSensorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SensorUpdateSensorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SensorUpdateSensorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *SensorUpdateSensorRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SensorUpdateSensorRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SensorUpdateSensorRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SensorUpdateSensorRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetSensor

`func (o *SensorUpdateSensorRequest) GetSensor() IoArgoprojEventsV1alpha1Sensor`

GetSensor returns the Sensor field if non-nil, zero value otherwise.

### GetSensorOk

`func (o *SensorUpdateSensorRequest) GetSensorOk() (*IoArgoprojEventsV1alpha1Sensor, bool)`

GetSensorOk returns a tuple with the Sensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensor

`func (o *SensorUpdateSensorRequest) SetSensor(v IoArgoprojEventsV1alpha1Sensor)`

SetSensor sets Sensor field to given value.

### HasSensor

`func (o *SensorUpdateSensorRequest) HasSensor() bool`

HasSensor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


