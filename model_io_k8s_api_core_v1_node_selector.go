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

// IoK8sApiCoreV1NodeSelector A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
type IoK8sApiCoreV1NodeSelector struct {
	// Required. A list of node selector terms. The terms are ORed.
	NodeSelectorTerms []IoK8sApiCoreV1NodeSelectorTerm `json:"nodeSelectorTerms"`
}

// NewIoK8sApiCoreV1NodeSelector instantiates a new IoK8sApiCoreV1NodeSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApiCoreV1NodeSelector(nodeSelectorTerms []IoK8sApiCoreV1NodeSelectorTerm) *IoK8sApiCoreV1NodeSelector {
	this := IoK8sApiCoreV1NodeSelector{}
	this.NodeSelectorTerms = nodeSelectorTerms
	return &this
}

// NewIoK8sApiCoreV1NodeSelectorWithDefaults instantiates a new IoK8sApiCoreV1NodeSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApiCoreV1NodeSelectorWithDefaults() *IoK8sApiCoreV1NodeSelector {
	this := IoK8sApiCoreV1NodeSelector{}
	return &this
}

// GetNodeSelectorTerms returns the NodeSelectorTerms field value
func (o *IoK8sApiCoreV1NodeSelector) GetNodeSelectorTerms() []IoK8sApiCoreV1NodeSelectorTerm {
	if o == nil {
		var ret []IoK8sApiCoreV1NodeSelectorTerm
		return ret
	}

	return o.NodeSelectorTerms
}

// GetNodeSelectorTermsOk returns a tuple with the NodeSelectorTerms field value
// and a boolean to check if the value has been set.
func (o *IoK8sApiCoreV1NodeSelector) GetNodeSelectorTermsOk() (*[]IoK8sApiCoreV1NodeSelectorTerm, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NodeSelectorTerms, true
}

// SetNodeSelectorTerms sets field value
func (o *IoK8sApiCoreV1NodeSelector) SetNodeSelectorTerms(v []IoK8sApiCoreV1NodeSelectorTerm) {
	o.NodeSelectorTerms = v
}

func (o IoK8sApiCoreV1NodeSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nodeSelectorTerms"] = o.NodeSelectorTerms
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApiCoreV1NodeSelector struct {
	value *IoK8sApiCoreV1NodeSelector
	isSet bool
}

func (v NullableIoK8sApiCoreV1NodeSelector) Get() *IoK8sApiCoreV1NodeSelector {
	return v.value
}

func (v *NullableIoK8sApiCoreV1NodeSelector) Set(val *IoK8sApiCoreV1NodeSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApiCoreV1NodeSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApiCoreV1NodeSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApiCoreV1NodeSelector(val *IoK8sApiCoreV1NodeSelector) *NullableIoK8sApiCoreV1NodeSelector {
	return &NullableIoK8sApiCoreV1NodeSelector{value: val, isSet: true}
}

func (v NullableIoK8sApiCoreV1NodeSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApiCoreV1NodeSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


