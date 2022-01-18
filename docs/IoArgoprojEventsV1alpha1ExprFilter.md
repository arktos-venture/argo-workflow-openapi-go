# IoArgoprojEventsV1alpha1ExprFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expr** | Pointer to **string** | Expr refers to the expression that determines the outcome of the filter. | [optional] 
**Fields** | Pointer to [**[]IoArgoprojEventsV1alpha1PayloadField**](IoArgoprojEventsV1alpha1PayloadField.md) | Fields refers to set of keys that refer to the paths within event payload. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1ExprFilter

`func NewIoArgoprojEventsV1alpha1ExprFilter() *IoArgoprojEventsV1alpha1ExprFilter`

NewIoArgoprojEventsV1alpha1ExprFilter instantiates a new IoArgoprojEventsV1alpha1ExprFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1ExprFilterWithDefaults

`func NewIoArgoprojEventsV1alpha1ExprFilterWithDefaults() *IoArgoprojEventsV1alpha1ExprFilter`

NewIoArgoprojEventsV1alpha1ExprFilterWithDefaults instantiates a new IoArgoprojEventsV1alpha1ExprFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpr

`func (o *IoArgoprojEventsV1alpha1ExprFilter) GetExpr() string`

GetExpr returns the Expr field if non-nil, zero value otherwise.

### GetExprOk

`func (o *IoArgoprojEventsV1alpha1ExprFilter) GetExprOk() (*string, bool)`

GetExprOk returns a tuple with the Expr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpr

`func (o *IoArgoprojEventsV1alpha1ExprFilter) SetExpr(v string)`

SetExpr sets Expr field to given value.

### HasExpr

`func (o *IoArgoprojEventsV1alpha1ExprFilter) HasExpr() bool`

HasExpr returns a boolean if a field has been set.

### GetFields

`func (o *IoArgoprojEventsV1alpha1ExprFilter) GetFields() []IoArgoprojEventsV1alpha1PayloadField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IoArgoprojEventsV1alpha1ExprFilter) GetFieldsOk() (*[]IoArgoprojEventsV1alpha1PayloadField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IoArgoprojEventsV1alpha1ExprFilter) SetFields(v []IoArgoprojEventsV1alpha1PayloadField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *IoArgoprojEventsV1alpha1ExprFilter) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


