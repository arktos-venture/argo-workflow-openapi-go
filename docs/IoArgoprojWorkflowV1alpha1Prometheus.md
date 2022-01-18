# IoArgoprojWorkflowV1alpha1Prometheus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counter** | Pointer to [**IoArgoprojWorkflowV1alpha1Counter**](IoArgoprojWorkflowV1alpha1Counter.md) |  | [optional] 
**Gauge** | Pointer to [**IoArgoprojWorkflowV1alpha1Gauge**](IoArgoprojWorkflowV1alpha1Gauge.md) |  | [optional] 
**Help** | **string** | Help is a string that describes the metric | 
**Histogram** | Pointer to [**IoArgoprojWorkflowV1alpha1Histogram**](IoArgoprojWorkflowV1alpha1Histogram.md) |  | [optional] 
**Labels** | Pointer to [**[]IoArgoprojWorkflowV1alpha1MetricLabel**](IoArgoprojWorkflowV1alpha1MetricLabel.md) | Labels is a list of metric labels | [optional] 
**Name** | **string** | Name is the name of the metric | 
**When** | Pointer to **string** | When is a conditional statement that decides when to emit the metric | [optional] 

## Methods

### NewIoArgoprojWorkflowV1alpha1Prometheus

`func NewIoArgoprojWorkflowV1alpha1Prometheus(help string, name string, ) *IoArgoprojWorkflowV1alpha1Prometheus`

NewIoArgoprojWorkflowV1alpha1Prometheus instantiates a new IoArgoprojWorkflowV1alpha1Prometheus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1PrometheusWithDefaults

`func NewIoArgoprojWorkflowV1alpha1PrometheusWithDefaults() *IoArgoprojWorkflowV1alpha1Prometheus`

NewIoArgoprojWorkflowV1alpha1PrometheusWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Prometheus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounter

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetCounter() IoArgoprojWorkflowV1alpha1Counter`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetCounterOk() (*IoArgoprojWorkflowV1alpha1Counter, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) SetCounter(v IoArgoprojWorkflowV1alpha1Counter)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### GetGauge

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetGauge() IoArgoprojWorkflowV1alpha1Gauge`

GetGauge returns the Gauge field if non-nil, zero value otherwise.

### GetGaugeOk

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetGaugeOk() (*IoArgoprojWorkflowV1alpha1Gauge, bool)`

GetGaugeOk returns a tuple with the Gauge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGauge

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) SetGauge(v IoArgoprojWorkflowV1alpha1Gauge)`

SetGauge sets Gauge field to given value.

### HasGauge

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) HasGauge() bool`

HasGauge returns a boolean if a field has been set.

### GetHelp

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) SetHelp(v string)`

SetHelp sets Help field to given value.


### GetHistogram

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetHistogram() IoArgoprojWorkflowV1alpha1Histogram`

GetHistogram returns the Histogram field if non-nil, zero value otherwise.

### GetHistogramOk

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetHistogramOk() (*IoArgoprojWorkflowV1alpha1Histogram, bool)`

GetHistogramOk returns a tuple with the Histogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogram

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) SetHistogram(v IoArgoprojWorkflowV1alpha1Histogram)`

SetHistogram sets Histogram field to given value.

### HasHistogram

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) HasHistogram() bool`

HasHistogram returns a boolean if a field has been set.

### GetLabels

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetLabels() []IoArgoprojWorkflowV1alpha1MetricLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetLabelsOk() (*[]IoArgoprojWorkflowV1alpha1MetricLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) SetLabels(v []IoArgoprojWorkflowV1alpha1MetricLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) SetName(v string)`

SetName sets Name field to given value.


### GetWhen

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetWhen() string`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) GetWhenOk() (*string, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) SetWhen(v string)`

SetWhen sets When field to given value.

### HasWhen

`func (o *IoArgoprojWorkflowV1alpha1Prometheus) HasWhen() bool`

HasWhen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


