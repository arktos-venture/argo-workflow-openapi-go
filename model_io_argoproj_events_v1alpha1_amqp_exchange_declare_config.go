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

// IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig struct for IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig
type IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig struct {
	AutoDelete *bool `json:"autoDelete,omitempty"`
	Durable *bool `json:"durable,omitempty"`
	Internal *bool `json:"internal,omitempty"`
	NoWait *bool `json:"noWait,omitempty"`
}

// NewIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig instantiates a new IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig() *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig {
	this := IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig{}
	return &this
}

// NewIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfigWithDefaults instantiates a new IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfigWithDefaults() *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig {
	this := IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig{}
	return &this
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) GetAutoDelete() bool {
	if o == nil || o.AutoDelete == nil {
		var ret bool
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) GetAutoDeleteOk() (*bool, bool) {
	if o == nil || o.AutoDelete == nil {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) HasAutoDelete() bool {
	if o != nil && o.AutoDelete != nil {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given bool and assigns it to the AutoDelete field.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) SetAutoDelete(v bool) {
	o.AutoDelete = &v
}

// GetDurable returns the Durable field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) GetDurable() bool {
	if o == nil || o.Durable == nil {
		var ret bool
		return ret
	}
	return *o.Durable
}

// GetDurableOk returns a tuple with the Durable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) GetDurableOk() (*bool, bool) {
	if o == nil || o.Durable == nil {
		return nil, false
	}
	return o.Durable, true
}

// HasDurable returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) HasDurable() bool {
	if o != nil && o.Durable != nil {
		return true
	}

	return false
}

// SetDurable gets a reference to the given bool and assigns it to the Durable field.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) SetDurable(v bool) {
	o.Durable = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) SetInternal(v bool) {
	o.Internal = &v
}

// GetNoWait returns the NoWait field value if set, zero value otherwise.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) GetNoWait() bool {
	if o == nil || o.NoWait == nil {
		var ret bool
		return ret
	}
	return *o.NoWait
}

// GetNoWaitOk returns a tuple with the NoWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) GetNoWaitOk() (*bool, bool) {
	if o == nil || o.NoWait == nil {
		return nil, false
	}
	return o.NoWait, true
}

// HasNoWait returns a boolean if a field has been set.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) HasNoWait() bool {
	if o != nil && o.NoWait != nil {
		return true
	}

	return false
}

// SetNoWait gets a reference to the given bool and assigns it to the NoWait field.
func (o *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) SetNoWait(v bool) {
	o.NoWait = &v
}

func (o IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoDelete != nil {
		toSerialize["autoDelete"] = o.AutoDelete
	}
	if o.Durable != nil {
		toSerialize["durable"] = o.Durable
	}
	if o.Internal != nil {
		toSerialize["internal"] = o.Internal
	}
	if o.NoWait != nil {
		toSerialize["noWait"] = o.NoWait
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig struct {
	value *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig
	isSet bool
}

func (v NullableIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) Get() *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig {
	return v.value
}

func (v *NullableIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) Set(val *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig(val *IoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) *NullableIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig {
	return &NullableIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig{value: val, isSet: true}
}

func (v NullableIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojEventsV1alpha1AMQPExchangeDeclareConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


