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

// IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
type IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement struct {
	// key is the label key that the selector applies to.
	Key string `json:"key"`
	// operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator string `json:"operator"`
	// values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	Values *[]string `json:"values,omitempty"`
}

// NewIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement instantiates a new IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement(key string, operator string) *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement {
	this := IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement{}
	this.Key = key
	this.Operator = operator
	return &this
}

// NewIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirementWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirementWithDefaults() *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement {
	this := IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement{}
	return &this
}

// GetKey returns the Key field value
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) SetKey(v string) {
	o.Key = v
}

// GetOperator returns the Operator field value
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) GetOperatorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) SetOperator(v string) {
	o.Operator = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) SetValues(v []string) {
	o.Values = &v
}

func (o IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement struct {
	value *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement
	isSet bool
}

func (v NullableIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) Get() *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement {
	return v.value
}

func (v *NullableIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) Set(val *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement(val *IoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) *NullableIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement {
	return &NullableIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement{value: val, isSet: true}
}

func (v NullableIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApimachineryPkgApisMetaV1LabelSelectorRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


