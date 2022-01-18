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

// GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe struct for GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe
type GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe struct {
	AbstractStep *GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractStep `json:"abstractStep,omitempty"`
	// Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.  The serialization format is:  <quantity>        ::= <signedNumber><suffix>   (Note that <suffix> may be empty, from the \"\" case in <decimalSI>.) <digit>           ::= 0 | 1 | ... | 9 <digits>          ::= <digit> | <digit><digits> <number>          ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign>            ::= \"+\" | \"-\" <signedNumber>    ::= <number> | <sign><number> <suffix>          ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei   (International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI>       ::= m | \"\" | k | M | G | T | P | E   (Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= \"e\" <signedNumber> | \"E\" <signedNumber>  No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.  When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.  Before serializing, Quantity will be put in \"canonical form\". This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:   a. No precision is lost   b. No fractional digits will be emitted   c. The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative.  Examples:   1.5 will be serialized as \"1500m\"   1.5Gi will be serialized as \"1536Mi\"  Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.  Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.)  This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
	MaxSize *string `json:"maxSize,omitempty"`
	Uid *string `json:"uid,omitempty"`
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe{}
	return &this
}

// NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DedupeWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DedupeWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe {
	this := GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe{}
	return &this
}

// GetAbstractStep returns the AbstractStep field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) GetAbstractStep() GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractStep {
	if o == nil || o.AbstractStep == nil {
		var ret GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractStep
		return ret
	}
	return *o.AbstractStep
}

// GetAbstractStepOk returns a tuple with the AbstractStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) GetAbstractStepOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractStep, bool) {
	if o == nil || o.AbstractStep == nil {
		return nil, false
	}
	return o.AbstractStep, true
}

// HasAbstractStep returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) HasAbstractStep() bool {
	if o != nil && o.AbstractStep != nil {
		return true
	}

	return false
}

// SetAbstractStep gets a reference to the given GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractStep and assigns it to the AbstractStep field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) SetAbstractStep(v GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractStep) {
	o.AbstractStep = &v
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) GetMaxSize() string {
	if o == nil || o.MaxSize == nil {
		var ret string
		return ret
	}
	return *o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) GetMaxSizeOk() (*string, bool) {
	if o == nil || o.MaxSize == nil {
		return nil, false
	}
	return o.MaxSize, true
}

// HasMaxSize returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) HasMaxSize() bool {
	if o != nil && o.MaxSize != nil {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given string and assigns it to the MaxSize field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) SetMaxSize(v string) {
	o.MaxSize = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) SetUid(v string) {
	o.Uid = &v
}

func (o GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AbstractStep != nil {
		toSerialize["abstractStep"] = o.AbstractStep
	}
	if o.MaxSize != nil {
		toSerialize["maxSize"] = o.MaxSize
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	return json.Marshal(toSerialize)
}

type NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe struct {
	value *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe
	isSet bool
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) Get() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe {
	return v.value
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) Set(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe(val *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe {
	return &NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe{value: val, isSet: true}
}

func (v NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

