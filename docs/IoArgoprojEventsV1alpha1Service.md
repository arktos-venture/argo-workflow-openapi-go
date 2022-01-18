# IoArgoprojEventsV1alpha1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIP** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to [**[]IoK8sApiCoreV1ServicePort**](IoK8sApiCoreV1ServicePort.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1Service

`func NewIoArgoprojEventsV1alpha1Service() *IoArgoprojEventsV1alpha1Service`

NewIoArgoprojEventsV1alpha1Service instantiates a new IoArgoprojEventsV1alpha1Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1ServiceWithDefaults

`func NewIoArgoprojEventsV1alpha1ServiceWithDefaults() *IoArgoprojEventsV1alpha1Service`

NewIoArgoprojEventsV1alpha1ServiceWithDefaults instantiates a new IoArgoprojEventsV1alpha1Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIP

`func (o *IoArgoprojEventsV1alpha1Service) GetClusterIP() string`

GetClusterIP returns the ClusterIP field if non-nil, zero value otherwise.

### GetClusterIPOk

`func (o *IoArgoprojEventsV1alpha1Service) GetClusterIPOk() (*string, bool)`

GetClusterIPOk returns a tuple with the ClusterIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIP

`func (o *IoArgoprojEventsV1alpha1Service) SetClusterIP(v string)`

SetClusterIP sets ClusterIP field to given value.

### HasClusterIP

`func (o *IoArgoprojEventsV1alpha1Service) HasClusterIP() bool`

HasClusterIP returns a boolean if a field has been set.

### GetPorts

`func (o *IoArgoprojEventsV1alpha1Service) GetPorts() []IoK8sApiCoreV1ServicePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *IoArgoprojEventsV1alpha1Service) GetPortsOk() (*[]IoK8sApiCoreV1ServicePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *IoArgoprojEventsV1alpha1Service) SetPorts(v []IoK8sApiCoreV1ServicePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *IoArgoprojEventsV1alpha1Service) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


