# IoArgoprojEventsV1alpha1EventPersistence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catchup** | Pointer to [**IoArgoprojEventsV1alpha1CatchupConfiguration**](IoArgoprojEventsV1alpha1CatchupConfiguration.md) |  | [optional] 
**ConfigMap** | Pointer to [**IoArgoprojEventsV1alpha1ConfigMapPersistence**](IoArgoprojEventsV1alpha1ConfigMapPersistence.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1EventPersistence

`func NewIoArgoprojEventsV1alpha1EventPersistence() *IoArgoprojEventsV1alpha1EventPersistence`

NewIoArgoprojEventsV1alpha1EventPersistence instantiates a new IoArgoprojEventsV1alpha1EventPersistence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1EventPersistenceWithDefaults

`func NewIoArgoprojEventsV1alpha1EventPersistenceWithDefaults() *IoArgoprojEventsV1alpha1EventPersistence`

NewIoArgoprojEventsV1alpha1EventPersistenceWithDefaults instantiates a new IoArgoprojEventsV1alpha1EventPersistence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatchup

`func (o *IoArgoprojEventsV1alpha1EventPersistence) GetCatchup() IoArgoprojEventsV1alpha1CatchupConfiguration`

GetCatchup returns the Catchup field if non-nil, zero value otherwise.

### GetCatchupOk

`func (o *IoArgoprojEventsV1alpha1EventPersistence) GetCatchupOk() (*IoArgoprojEventsV1alpha1CatchupConfiguration, bool)`

GetCatchupOk returns a tuple with the Catchup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchup

`func (o *IoArgoprojEventsV1alpha1EventPersistence) SetCatchup(v IoArgoprojEventsV1alpha1CatchupConfiguration)`

SetCatchup sets Catchup field to given value.

### HasCatchup

`func (o *IoArgoprojEventsV1alpha1EventPersistence) HasCatchup() bool`

HasCatchup returns a boolean if a field has been set.

### GetConfigMap

`func (o *IoArgoprojEventsV1alpha1EventPersistence) GetConfigMap() IoArgoprojEventsV1alpha1ConfigMapPersistence`

GetConfigMap returns the ConfigMap field if non-nil, zero value otherwise.

### GetConfigMapOk

`func (o *IoArgoprojEventsV1alpha1EventPersistence) GetConfigMapOk() (*IoArgoprojEventsV1alpha1ConfigMapPersistence, bool)`

GetConfigMapOk returns a tuple with the ConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMap

`func (o *IoArgoprojEventsV1alpha1EventPersistence) SetConfigMap(v IoArgoprojEventsV1alpha1ConfigMapPersistence)`

SetConfigMap sets ConfigMap field to given value.

### HasConfigMap

`func (o *IoArgoprojEventsV1alpha1EventPersistence) HasConfigMap() bool`

HasConfigMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


