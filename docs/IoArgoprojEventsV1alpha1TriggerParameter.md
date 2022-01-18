# IoArgoprojEventsV1alpha1TriggerParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dest** | Pointer to **string** | Dest is the JSONPath of a resource key. A path is a series of keys separated by a dot. The colon character can be escaped with &#39;.&#39; The -1 key can be used to append a value to an existing array. See https://github.com/tidwall/sjson#path-syntax for more information about how this is used. | [optional] 
**Operation** | Pointer to **string** | Operation is what to do with the existing value at Dest, whether to &#39;prepend&#39;, &#39;overwrite&#39;, or &#39;append&#39; it. | [optional] 
**Src** | Pointer to [**IoArgoprojEventsV1alpha1TriggerParameterSource**](IoArgoprojEventsV1alpha1TriggerParameterSource.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1TriggerParameter

`func NewIoArgoprojEventsV1alpha1TriggerParameter() *IoArgoprojEventsV1alpha1TriggerParameter`

NewIoArgoprojEventsV1alpha1TriggerParameter instantiates a new IoArgoprojEventsV1alpha1TriggerParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1TriggerParameterWithDefaults

`func NewIoArgoprojEventsV1alpha1TriggerParameterWithDefaults() *IoArgoprojEventsV1alpha1TriggerParameter`

NewIoArgoprojEventsV1alpha1TriggerParameterWithDefaults instantiates a new IoArgoprojEventsV1alpha1TriggerParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDest

`func (o *IoArgoprojEventsV1alpha1TriggerParameter) GetDest() string`

GetDest returns the Dest field if non-nil, zero value otherwise.

### GetDestOk

`func (o *IoArgoprojEventsV1alpha1TriggerParameter) GetDestOk() (*string, bool)`

GetDestOk returns a tuple with the Dest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDest

`func (o *IoArgoprojEventsV1alpha1TriggerParameter) SetDest(v string)`

SetDest sets Dest field to given value.

### HasDest

`func (o *IoArgoprojEventsV1alpha1TriggerParameter) HasDest() bool`

HasDest returns a boolean if a field has been set.

### GetOperation

`func (o *IoArgoprojEventsV1alpha1TriggerParameter) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *IoArgoprojEventsV1alpha1TriggerParameter) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *IoArgoprojEventsV1alpha1TriggerParameter) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *IoArgoprojEventsV1alpha1TriggerParameter) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetSrc

`func (o *IoArgoprojEventsV1alpha1TriggerParameter) GetSrc() IoArgoprojEventsV1alpha1TriggerParameterSource`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *IoArgoprojEventsV1alpha1TriggerParameter) GetSrcOk() (*IoArgoprojEventsV1alpha1TriggerParameterSource, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *IoArgoprojEventsV1alpha1TriggerParameter) SetSrc(v IoArgoprojEventsV1alpha1TriggerParameterSource)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *IoArgoprojEventsV1alpha1TriggerParameter) HasSrc() bool`

HasSrc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


