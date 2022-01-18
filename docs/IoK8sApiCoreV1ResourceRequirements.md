# IoK8sApiCoreV1ResourceRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | Pointer to **map[string]string** | Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ | [optional] 
**Requests** | Pointer to **map[string]string** | Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/ | [optional] 

## Methods

### NewIoK8sApiCoreV1ResourceRequirements

`func NewIoK8sApiCoreV1ResourceRequirements() *IoK8sApiCoreV1ResourceRequirements`

NewIoK8sApiCoreV1ResourceRequirements instantiates a new IoK8sApiCoreV1ResourceRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1ResourceRequirementsWithDefaults

`func NewIoK8sApiCoreV1ResourceRequirementsWithDefaults() *IoK8sApiCoreV1ResourceRequirements`

NewIoK8sApiCoreV1ResourceRequirementsWithDefaults instantiates a new IoK8sApiCoreV1ResourceRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimits

`func (o *IoK8sApiCoreV1ResourceRequirements) GetLimits() map[string]string`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *IoK8sApiCoreV1ResourceRequirements) GetLimitsOk() (*map[string]string, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *IoK8sApiCoreV1ResourceRequirements) SetLimits(v map[string]string)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *IoK8sApiCoreV1ResourceRequirements) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetRequests

`func (o *IoK8sApiCoreV1ResourceRequirements) GetRequests() map[string]string`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *IoK8sApiCoreV1ResourceRequirements) GetRequestsOk() (*map[string]string, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *IoK8sApiCoreV1ResourceRequirements) SetRequests(v map[string]string)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *IoK8sApiCoreV1ResourceRequirements) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


