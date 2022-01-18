# IoArgoprojEventsV1alpha1DataFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparator** | Pointer to **string** | Comparator compares the event data with a user given value. Can be \&quot;&gt;&#x3D;\&quot;, \&quot;&gt;\&quot;, \&quot;&#x3D;\&quot;, \&quot;!&#x3D;\&quot;, \&quot;&lt;\&quot;, or \&quot;&lt;&#x3D;\&quot;. Is optional, and if left blank treated as equality \&quot;&#x3D;\&quot;. | [optional] 
**Path** | Pointer to **string** | Path is the JSONPath of the event&#39;s (JSON decoded) data key Path is a series of keys separated by a dot. A key may contain wildcard characters &#39;*&#39; and &#39;?&#39;. To access an array value use the index as the key. The dot and wildcard characters can be escaped with &#39;\\\\&#39;. See https://github.com/tidwall/gjson#path-syntax for more information on how to use this. | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1DataFilter

`func NewIoArgoprojEventsV1alpha1DataFilter() *IoArgoprojEventsV1alpha1DataFilter`

NewIoArgoprojEventsV1alpha1DataFilter instantiates a new IoArgoprojEventsV1alpha1DataFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1DataFilterWithDefaults

`func NewIoArgoprojEventsV1alpha1DataFilterWithDefaults() *IoArgoprojEventsV1alpha1DataFilter`

NewIoArgoprojEventsV1alpha1DataFilterWithDefaults instantiates a new IoArgoprojEventsV1alpha1DataFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparator

`func (o *IoArgoprojEventsV1alpha1DataFilter) GetComparator() string`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *IoArgoprojEventsV1alpha1DataFilter) GetComparatorOk() (*string, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *IoArgoprojEventsV1alpha1DataFilter) SetComparator(v string)`

SetComparator sets Comparator field to given value.

### HasComparator

`func (o *IoArgoprojEventsV1alpha1DataFilter) HasComparator() bool`

HasComparator returns a boolean if a field has been set.

### GetPath

`func (o *IoArgoprojEventsV1alpha1DataFilter) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *IoArgoprojEventsV1alpha1DataFilter) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *IoArgoprojEventsV1alpha1DataFilter) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *IoArgoprojEventsV1alpha1DataFilter) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetTemplate

`func (o *IoArgoprojEventsV1alpha1DataFilter) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *IoArgoprojEventsV1alpha1DataFilter) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *IoArgoprojEventsV1alpha1DataFilter) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *IoArgoprojEventsV1alpha1DataFilter) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetType

`func (o *IoArgoprojEventsV1alpha1DataFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IoArgoprojEventsV1alpha1DataFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IoArgoprojEventsV1alpha1DataFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IoArgoprojEventsV1alpha1DataFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *IoArgoprojEventsV1alpha1DataFilter) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IoArgoprojEventsV1alpha1DataFilter) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IoArgoprojEventsV1alpha1DataFilter) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IoArgoprojEventsV1alpha1DataFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


