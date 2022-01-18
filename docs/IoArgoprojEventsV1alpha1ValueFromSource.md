# IoArgoprojEventsV1alpha1ValueFromSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMapKeyRef** | Pointer to [**IoK8sApiCoreV1ConfigMapKeySelector**](IoK8sApiCoreV1ConfigMapKeySelector.md) |  | [optional] 
**SecretKeyRef** | Pointer to [**IoK8sApiCoreV1SecretKeySelector**](IoK8sApiCoreV1SecretKeySelector.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1ValueFromSource

`func NewIoArgoprojEventsV1alpha1ValueFromSource() *IoArgoprojEventsV1alpha1ValueFromSource`

NewIoArgoprojEventsV1alpha1ValueFromSource instantiates a new IoArgoprojEventsV1alpha1ValueFromSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1ValueFromSourceWithDefaults

`func NewIoArgoprojEventsV1alpha1ValueFromSourceWithDefaults() *IoArgoprojEventsV1alpha1ValueFromSource`

NewIoArgoprojEventsV1alpha1ValueFromSourceWithDefaults instantiates a new IoArgoprojEventsV1alpha1ValueFromSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMapKeyRef

`func (o *IoArgoprojEventsV1alpha1ValueFromSource) GetConfigMapKeyRef() IoK8sApiCoreV1ConfigMapKeySelector`

GetConfigMapKeyRef returns the ConfigMapKeyRef field if non-nil, zero value otherwise.

### GetConfigMapKeyRefOk

`func (o *IoArgoprojEventsV1alpha1ValueFromSource) GetConfigMapKeyRefOk() (*IoK8sApiCoreV1ConfigMapKeySelector, bool)`

GetConfigMapKeyRefOk returns a tuple with the ConfigMapKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapKeyRef

`func (o *IoArgoprojEventsV1alpha1ValueFromSource) SetConfigMapKeyRef(v IoK8sApiCoreV1ConfigMapKeySelector)`

SetConfigMapKeyRef sets ConfigMapKeyRef field to given value.

### HasConfigMapKeyRef

`func (o *IoArgoprojEventsV1alpha1ValueFromSource) HasConfigMapKeyRef() bool`

HasConfigMapKeyRef returns a boolean if a field has been set.

### GetSecretKeyRef

`func (o *IoArgoprojEventsV1alpha1ValueFromSource) GetSecretKeyRef() IoK8sApiCoreV1SecretKeySelector`

GetSecretKeyRef returns the SecretKeyRef field if non-nil, zero value otherwise.

### GetSecretKeyRefOk

`func (o *IoArgoprojEventsV1alpha1ValueFromSource) GetSecretKeyRefOk() (*IoK8sApiCoreV1SecretKeySelector, bool)`

GetSecretKeyRefOk returns a tuple with the SecretKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyRef

`func (o *IoArgoprojEventsV1alpha1ValueFromSource) SetSecretKeyRef(v IoK8sApiCoreV1SecretKeySelector)`

SetSecretKeyRef sets SecretKeyRef field to given value.

### HasSecretKeyRef

`func (o *IoArgoprojEventsV1alpha1ValueFromSource) HasSecretKeyRef() bool`

HasSecretKeyRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


