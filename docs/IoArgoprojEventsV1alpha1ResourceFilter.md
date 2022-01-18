# IoArgoprojEventsV1alpha1ResourceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterStart** | Pointer to **bool** |  | [optional] 
**CreatedBy** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] 
**Fields** | Pointer to [**[]IoArgoprojEventsV1alpha1Selector**](IoArgoprojEventsV1alpha1Selector.md) |  | [optional] 
**Labels** | Pointer to [**[]IoArgoprojEventsV1alpha1Selector**](IoArgoprojEventsV1alpha1Selector.md) |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1ResourceFilter

`func NewIoArgoprojEventsV1alpha1ResourceFilter() *IoArgoprojEventsV1alpha1ResourceFilter`

NewIoArgoprojEventsV1alpha1ResourceFilter instantiates a new IoArgoprojEventsV1alpha1ResourceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1ResourceFilterWithDefaults

`func NewIoArgoprojEventsV1alpha1ResourceFilterWithDefaults() *IoArgoprojEventsV1alpha1ResourceFilter`

NewIoArgoprojEventsV1alpha1ResourceFilterWithDefaults instantiates a new IoArgoprojEventsV1alpha1ResourceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfterStart

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) GetAfterStart() bool`

GetAfterStart returns the AfterStart field if non-nil, zero value otherwise.

### GetAfterStartOk

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) GetAfterStartOk() (*bool, bool)`

GetAfterStartOk returns a tuple with the AfterStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterStart

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) SetAfterStart(v bool)`

SetAfterStart sets AfterStart field to given value.

### HasAfterStart

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) HasAfterStart() bool`

HasAfterStart returns a boolean if a field has been set.

### GetCreatedBy

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) GetCreatedBy() time.Time`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) GetCreatedByOk() (*time.Time, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) SetCreatedBy(v time.Time)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetFields

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) GetFields() []IoArgoprojEventsV1alpha1Selector`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) GetFieldsOk() (*[]IoArgoprojEventsV1alpha1Selector, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) SetFields(v []IoArgoprojEventsV1alpha1Selector)`

SetFields sets Fields field to given value.

### HasFields

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetLabels

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) GetLabels() []IoArgoprojEventsV1alpha1Selector`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) GetLabelsOk() (*[]IoArgoprojEventsV1alpha1Selector, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) SetLabels(v []IoArgoprojEventsV1alpha1Selector)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPrefix

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *IoArgoprojEventsV1alpha1ResourceFilter) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


