# IoArgoprojEventsV1alpha1PayloadField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name acts as key that holds the value at the path. | [optional] 
**Path** | Pointer to **string** | Path is the JSONPath of the event&#39;s (JSON decoded) data key Path is a series of keys separated by a dot. A key may contain wildcard characters &#39;*&#39; and &#39;?&#39;. To access an array value use the index as the key. The dot and wildcard characters can be escaped with &#39;\\\\&#39;. See https://github.com/tidwall/gjson#path-syntax for more information on how to use this. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1PayloadField

`func NewIoArgoprojEventsV1alpha1PayloadField() *IoArgoprojEventsV1alpha1PayloadField`

NewIoArgoprojEventsV1alpha1PayloadField instantiates a new IoArgoprojEventsV1alpha1PayloadField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1PayloadFieldWithDefaults

`func NewIoArgoprojEventsV1alpha1PayloadFieldWithDefaults() *IoArgoprojEventsV1alpha1PayloadField`

NewIoArgoprojEventsV1alpha1PayloadFieldWithDefaults instantiates a new IoArgoprojEventsV1alpha1PayloadField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoArgoprojEventsV1alpha1PayloadField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojEventsV1alpha1PayloadField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojEventsV1alpha1PayloadField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoArgoprojEventsV1alpha1PayloadField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *IoArgoprojEventsV1alpha1PayloadField) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoArgoprojEventsV1alpha1PayloadField) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoArgoprojEventsV1alpha1PayloadField) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IoArgoprojEventsV1alpha1PayloadField) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


