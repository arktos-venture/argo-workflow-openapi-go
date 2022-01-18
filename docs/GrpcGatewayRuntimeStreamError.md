# GrpcGatewayRuntimeStreamError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]GoogleProtobufAny**](GoogleProtobufAny.md) |  | [optional] 
**GrpcCode** | Pointer to **int32** |  | [optional] 
**HttpCode** | Pointer to **int32** |  | [optional] 
**HttpStatus** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewGrpcGatewayRuntimeStreamError

`func NewGrpcGatewayRuntimeStreamError() *GrpcGatewayRuntimeStreamError`

NewGrpcGatewayRuntimeStreamError instantiates a new GrpcGatewayRuntimeStreamError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrpcGatewayRuntimeStreamErrorWithDefaults

`func NewGrpcGatewayRuntimeStreamErrorWithDefaults() *GrpcGatewayRuntimeStreamError`

NewGrpcGatewayRuntimeStreamErrorWithDefaults instantiates a new GrpcGatewayRuntimeStreamError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *GrpcGatewayRuntimeStreamError) GetDetails() []GoogleProtobufAny`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GrpcGatewayRuntimeStreamError) GetDetailsOk() (*[]GoogleProtobufAny, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GrpcGatewayRuntimeStreamError) SetDetails(v []GoogleProtobufAny)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GrpcGatewayRuntimeStreamError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetGrpcCode

`func (o *GrpcGatewayRuntimeStreamError) GetGrpcCode() int32`

GetGrpcCode returns the GrpcCode field if non-nil, zero value otherwise.

### GetGrpcCodeOk

`func (o *GrpcGatewayRuntimeStreamError) GetGrpcCodeOk() (*int32, bool)`

GetGrpcCodeOk returns a tuple with the GrpcCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpcCode

`func (o *GrpcGatewayRuntimeStreamError) SetGrpcCode(v int32)`

SetGrpcCode sets GrpcCode field to given value.

### HasGrpcCode

`func (o *GrpcGatewayRuntimeStreamError) HasGrpcCode() bool`

HasGrpcCode returns a boolean if a field has been set.

### GetHttpCode

`func (o *GrpcGatewayRuntimeStreamError) GetHttpCode() int32`

GetHttpCode returns the HttpCode field if non-nil, zero value otherwise.

### GetHttpCodeOk

`func (o *GrpcGatewayRuntimeStreamError) GetHttpCodeOk() (*int32, bool)`

GetHttpCodeOk returns a tuple with the HttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCode

`func (o *GrpcGatewayRuntimeStreamError) SetHttpCode(v int32)`

SetHttpCode sets HttpCode field to given value.

### HasHttpCode

`func (o *GrpcGatewayRuntimeStreamError) HasHttpCode() bool`

HasHttpCode returns a boolean if a field has been set.

### GetHttpStatus

`func (o *GrpcGatewayRuntimeStreamError) GetHttpStatus() string`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *GrpcGatewayRuntimeStreamError) GetHttpStatusOk() (*string, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *GrpcGatewayRuntimeStreamError) SetHttpStatus(v string)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *GrpcGatewayRuntimeStreamError) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetMessage

`func (o *GrpcGatewayRuntimeStreamError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GrpcGatewayRuntimeStreamError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GrpcGatewayRuntimeStreamError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GrpcGatewayRuntimeStreamError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


