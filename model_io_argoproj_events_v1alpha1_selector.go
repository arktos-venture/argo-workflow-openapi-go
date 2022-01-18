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

// IoArgoprojEventsV1alpha1Selector Selector represents conditional operation to select K8s objects.
type IoArgoprojEventsV1alpha1Selector struct {
	Key *string `json:"key,omitempty"`
	Operation *string `json:"operation,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewIoArgoprojEventsV1alpha1Selector instantiates a new IoArgoprojEventsV1alpha1Selector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1Selector() *IoArgoprojEventsV1alpha1Selector {
	this := IoArgoprojEventsV1alpha1Selector{}
	return &this
}

// NewIoArgoprojEventsV1alpha1SelectorWithDefaults instantiates a new IoArgoprojEventsV1alpha1Selector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1SelectorWithDefaults() *IoArgoprojEventsV1alpha1Selector {
	this := IoArgoprojEventsV1alpha1Selector{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Selector) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Selector) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Selector) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *IoArgoprojEventsV1alpha1Selector) SetKey(v string) {
	o.Key = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Selector) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Selector) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Selector) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *IoArgoprojEventsV1alpha1Selector) SetOperation(v string) {
	o.Operation = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1Selector) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1Selector) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1Selector) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IoArgoprojEventsV1alpha1Selector) SetValue(v string) {
	o.Value = &v
}

func (o IoArgoprojEventsV1alpha1Selector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1Selector struct {
	value *IoArgoprojEventsV1alpha1Selector
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1Selector) Get() *IoArgoprojEventsV1alpha1Selector {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1Selector) Set(val *IoArgoprojEventsV1alpha1Selector) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1Selector) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1Selector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1Selector(val *IoArgoprojEventsV1alpha1Selector) *NullableIoArgoprojEventsV1alpha1Selector {
	return &NullableIoArgoprojEventsV1alpha1Selector{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1Selector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1Selector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

