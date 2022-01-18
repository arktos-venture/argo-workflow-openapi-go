/*
Argo Server API

You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`

API version: VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry struct for StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry
type StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry struct {
	Error *GrpcGatewayRuntimeStreamError `json:"error,omitempty"`
	Result *IoArgoprojWorkflowV1alpha1LogEntry `json:"result,omitempty"`
}

// NewStreamResultOfIoArgoprojWorkflowV1alpha1LogEntry instantiates a new StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamResultOfIoArgoprojWorkflowV1alpha1LogEntry() *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry {
	this := StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry{}
	return &this
}

// NewStreamResultOfIoArgoprojWorkflowV1alpha1LogEntryWithDefaults instantiates a new StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamResultOfIoArgoprojWorkflowV1alpha1LogEntryWithDefaults() *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry {
	this := StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) GetError() GrpcGatewayRuntimeStreamError {
	if o == nil || o.Error == nil {
		var ret GrpcGatewayRuntimeStreamError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) GetErrorOk() (*GrpcGatewayRuntimeStreamError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given GrpcGatewayRuntimeStreamError and assigns it to the Error field.
func (o *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) SetError(v GrpcGatewayRuntimeStreamError) {
	o.Error = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) GetResult() IoArgoprojWorkflowV1alpha1LogEntry {
	if o == nil || o.Result == nil {
		var ret IoArgoprojWorkflowV1alpha1LogEntry
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) GetResultOk() (*IoArgoprojWorkflowV1alpha1LogEntry, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given IoArgoprojWorkflowV1alpha1LogEntry and assigns it to the Result field.
func (o *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) SetResult(v IoArgoprojWorkflowV1alpha1LogEntry) {
	o.Result = &v
}

func (o StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableStreamResultOfIoArgoprojWorkflowV1alpha1LogEntry struct {
	value *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry
	isSet bool
}

func (v NullableStreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) Get() *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry {
	return v.value
}

func (v *NullableStreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) Set(val *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamResultOfIoArgoprojWorkflowV1alpha1LogEntry(val *StreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) *NullableStreamResultOfIoArgoprojWorkflowV1alpha1LogEntry {
	return &NullableStreamResultOfIoArgoprojWorkflowV1alpha1LogEntry{value: val, isSet: true}
}

func (v NullableStreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamResultOfIoArgoprojWorkflowV1alpha1LogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


