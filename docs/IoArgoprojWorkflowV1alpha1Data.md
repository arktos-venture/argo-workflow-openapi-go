# IoArgoprojWorkflowV1alpha1Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**IoArgoprojWorkflowV1alpha1DataSource**](IoArgoprojWorkflowV1alpha1DataSource.md) |  | 
**Transformation** | [**[]IoArgoprojWorkflowV1alpha1TransformationStep**](IoArgoprojWorkflowV1alpha1TransformationStep.md) | Transformation applies a set of transformations | 

## Methods

### NewIoArgoprojWorkflowV1alpha1Data

`func NewIoArgoprojWorkflowV1alpha1Data(source IoArgoprojWorkflowV1alpha1DataSource, transformation []IoArgoprojWorkflowV1alpha1TransformationStep, ) *IoArgoprojWorkflowV1alpha1Data`

NewIoArgoprojWorkflowV1alpha1Data instantiates a new IoArgoprojWorkflowV1alpha1Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojWorkflowV1alpha1DataWithDefaults

`func NewIoArgoprojWorkflowV1alpha1DataWithDefaults() *IoArgoprojWorkflowV1alpha1Data`

NewIoArgoprojWorkflowV1alpha1DataWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *IoArgoprojWorkflowV1alpha1Data) GetSource() IoArgoprojWorkflowV1alpha1DataSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IoArgoprojWorkflowV1alpha1Data) GetSourceOk() (*IoArgoprojWorkflowV1alpha1DataSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IoArgoprojWorkflowV1alpha1Data) SetSource(v IoArgoprojWorkflowV1alpha1DataSource)`

SetSource sets Source field to given value.


### GetTransformation

`func (o *IoArgoprojWorkflowV1alpha1Data) GetTransformation() []IoArgoprojWorkflowV1alpha1TransformationStep`

GetTransformation returns the Transformation field if non-nil, zero value otherwise.

### GetTransformationOk

`func (o *IoArgoprojWorkflowV1alpha1Data) GetTransformationOk() (*[]IoArgoprojWorkflowV1alpha1TransformationStep, bool)`

GetTransformationOk returns a tuple with the Transformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformation

`func (o *IoArgoprojWorkflowV1alpha1Data) SetTransformation(v []IoArgoprojWorkflowV1alpha1TransformationStep)`

SetTransformation sets Transformation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


