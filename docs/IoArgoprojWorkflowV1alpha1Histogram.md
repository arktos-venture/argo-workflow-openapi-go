# IoArgoprojWorkflowV1alpha1Histogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buckets** | **[]float32** | Buckets is a list of bucket divisors for the histogram | 
**Value** | **string** | Value is the value of the metric | 

## Methods

### NewIoArgoprojWorkflowV1alpha1Histogram

`func NewIoArgoprojWorkflowV1alpha1Histogram(buckets []float32, value string, ) *IoArgoprojWorkflowV1alpha1Histogram`

NewIoArgoprojWorkflowV1alpha1Histogram instantiates a new IoArgoprojWorkflowV1alpha1Histogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1HistogramWithDefaults

`func NewIoArgoprojWorkflowV1alpha1HistogramWithDefaults() *IoArgoprojWorkflowV1alpha1Histogram`

NewIoArgoprojWorkflowV1alpha1HistogramWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Histogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *IoArgoprojWorkflowV1alpha1Histogram) GetBuckets() []float32`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *IoArgoprojWorkflowV1alpha1Histogram) GetBucketsOk() (*[]float32, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *IoArgoprojWorkflowV1alpha1Histogram) SetBuckets(v []float32)`

SetBuckets sets Buckets field to given value.


### GetValue

`func (o *IoArgoprojWorkflowV1alpha1Histogram) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IoArgoprojWorkflowV1alpha1Histogram) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IoArgoprojWorkflowV1alpha1Histogram) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


