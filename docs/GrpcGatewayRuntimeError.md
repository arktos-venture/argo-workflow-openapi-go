# GrpcGatewayRuntimeError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Details** | Pointer to [**[]GoogleProtobufAny**](GoogleProtobufAny.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewGrpcGatewayRuntimeError

`func NewGrpcGatewayRuntimeError() *GrpcGatewayRuntimeError`

NewGrpcGatewayRuntimeError instantiates a new GrpcGatewayRuntimeError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrpcGatewayRuntimeErrorWithDefaults

`func NewGrpcGatewayRuntimeErrorWithDefaults() *GrpcGatewayRuntimeError`

NewGrpcGatewayRuntimeErrorWithDefaults instantiates a new GrpcGatewayRuntimeError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GrpcGatewayRuntimeError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GrpcGatewayRuntimeError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GrpcGatewayRuntimeError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GrpcGatewayRuntimeError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *GrpcGatewayRuntimeError) GetDetails() []GoogleProtobufAny`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GrpcGatewayRuntimeError) GetDetailsOk() (*[]GoogleProtobufAny, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GrpcGatewayRuntimeError) SetDetails(v []GoogleProtobufAny)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GrpcGatewayRuntimeError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetError

`func (o *GrpcGatewayRuntimeError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GrpcGatewayRuntimeError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GrpcGatewayRuntimeError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GrpcGatewayRuntimeError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *GrpcGatewayRuntimeError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GrpcGatewayRuntimeError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GrpcGatewayRuntimeError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GrpcGatewayRuntimeError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


