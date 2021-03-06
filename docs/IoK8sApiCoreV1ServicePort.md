# IoK8sApiCoreV1ServicePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the &#39;name&#39; field in the EndpointPort. Optional if only one ServicePort is defined on this service. | [optional] 
**NodePort** | Pointer to **int32** | The port on each node on which this service is exposed when type&#x3D;NodePort or LoadBalancer. Usually assigned by the system. If specified, it will be allocated to the service if unused or else creation of the service will fail. Default is to auto-allocate a port if the ServiceType of this Service requires one. More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport | [optional] 
**Port** | **int32** | The port that will be exposed by this service. | 
**Protocol** | Pointer to **string** | The IP protocol for this port. Supports \&quot;TCP\&quot;, \&quot;UDP\&quot;, and \&quot;SCTP\&quot;. Default is TCP. | [optional] 
**TargetPort** | Pointer to **string** |  | [optional] 

## Methods

### NewIoK8sApiCoreV1ServicePort

`func NewIoK8sApiCoreV1ServicePort(port int32, ) *IoK8sApiCoreV1ServicePort`

NewIoK8sApiCoreV1ServicePort instantiates a new IoK8sApiCoreV1ServicePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ServicePortWithDefaults

`func NewIoK8sApiCoreV1ServicePortWithDefaults() *IoK8sApiCoreV1ServicePort`

NewIoK8sApiCoreV1ServicePortWithDefaults instantiates a new IoK8sApiCoreV1ServicePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoK8sApiCoreV1ServicePort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoK8sApiCoreV1ServicePort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoK8sApiCoreV1ServicePort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoK8sApiCoreV1ServicePort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodePort

`func (o *IoK8sApiCoreV1ServicePort) GetNodePort() int32`

GetNodePort returns the NodePort field if non-nil, zero value otherwise.

### GetNodePortOk

`func (o *IoK8sApiCoreV1ServicePort) GetNodePortOk() (*int32, bool)`

GetNodePortOk returns a tuple with the NodePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePort

`func (o *IoK8sApiCoreV1ServicePort) SetNodePort(v int32)`

SetNodePort sets NodePort field to given value.

### HasNodePort

`func (o *IoK8sApiCoreV1ServicePort) HasNodePort() bool`

HasNodePort returns a boolean if a field has been set.

### GetPort

`func (o *IoK8sApiCoreV1ServicePort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IoK8sApiCoreV1ServicePort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IoK8sApiCoreV1ServicePort) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *IoK8sApiCoreV1ServicePort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IoK8sApiCoreV1ServicePort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IoK8sApiCoreV1ServicePort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IoK8sApiCoreV1ServicePort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTargetPort

`func (o *IoK8sApiCoreV1ServicePort) GetTargetPort() string`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *IoK8sApiCoreV1ServicePort) GetTargetPortOk() (*string, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *IoK8sApiCoreV1ServicePort) SetTargetPort(v string)`

SetTargetPort sets TargetPort field to given value.

### HasTargetPort

`func (o *IoK8sApiCoreV1ServicePort) HasTargetPort() bool`

HasTargetPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


