/*
Argo Server API

You can get examples of requests and responses by using the CLI with `--gloglevel=9`, e.g. `argo list --gloglevel=9`

API version: VERSION
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.
type IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry struct {
	// APIVersion defines the version of this resource that this field set applies to. The format is \"group/version\" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted.
	ApiVersion *string `json:"apiVersion,omitempty"`
	// FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: \"FieldsV1\"
	FieldsType *string `json:"fieldsType,omitempty"`
	// FieldsV1 stores a set of fields in a data structure like a Trie, in JSON format.  Each key is either a '.' representing the field itself, and will always map to an empty set, or a string representing a sub-field or item. The string will follow one of these four formats: 'f:<name>', where <name> is the name of a field in a struct, or key in a map 'v:<value>', where <value> is the exact json formatted value of a list item 'i:<index>', where <index> is position of a item in a list 'k:<keys>', where <keys> is a map of  a list item's key fields to their unique values If a key maps to an empty Fields value, the field that key represents is part of the set.  The exact format is defined in sigs.k8s.io/structured-merge-diff
	FieldsV1 *map[string]interface{} `json:"fieldsV1,omitempty"`
	// Manager is an identifier of the workflow managing these fields.
	Manager *string `json:"manager,omitempty"`
	// Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.
	Operation *string `json:"operation,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	Time *time.Time `json:"time,omitempty"`
}

// NewIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry instantiates a new IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry() *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry {
	this := IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry{}
	return &this
}

// NewIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntryWithDefaults instantiates a new IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntryWithDefaults() *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry {
	this := IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetFieldsType returns the FieldsType field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetFieldsType() string {
	if o == nil || o.FieldsType == nil {
		var ret string
		return ret
	}
	return *o.FieldsType
}

// GetFieldsTypeOk returns a tuple with the FieldsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetFieldsTypeOk() (*string, bool) {
	if o == nil || o.FieldsType == nil {
		return nil, false
	}
	return o.FieldsType, true
}

// HasFieldsType returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasFieldsType() bool {
	if o != nil && o.FieldsType != nil {
		return true
	}

	return false
}

// SetFieldsType gets a reference to the given string and assigns it to the FieldsType field.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetFieldsType(v string) {
	o.FieldsType = &v
}

// GetFieldsV1 returns the FieldsV1 field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetFieldsV1() map[string]interface{} {
	if o == nil || o.FieldsV1 == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.FieldsV1
}

// GetFieldsV1Ok returns a tuple with the FieldsV1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetFieldsV1Ok() (*map[string]interface{}, bool) {
	if o == nil || o.FieldsV1 == nil {
		return nil, false
	}
	return o.FieldsV1, true
}

// HasFieldsV1 returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasFieldsV1() bool {
	if o != nil && o.FieldsV1 != nil {
		return true
	}

	return false
}

// SetFieldsV1 gets a reference to the given map[string]interface{} and assigns it to the FieldsV1 field.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetFieldsV1(v map[string]interface{}) {
	o.FieldsV1 = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetManager() string {
	if o == nil || o.Manager == nil {
		var ret string
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetManagerOk() (*string, bool) {
	if o == nil || o.Manager == nil {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasManager() bool {
	if o != nil && o.Manager != nil {
		return true
	}

	return false
}

// SetManager gets a reference to the given string and assigns it to the Manager field.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetManager(v string) {
	o.Manager = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetOperation(v string) {
	o.Operation = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) SetTime(v time.Time) {
	o.Time = &v
}

func (o IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if o.FieldsType != nil {
		toSerialize["fieldsType"] = o.FieldsType
	}
	if o.FieldsV1 != nil {
		toSerialize["fieldsV1"] = o.FieldsV1
	}
	if o.Manager != nil {
		toSerialize["manager"] = o.Manager
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry struct {
	value *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry
	isSet bool
}

func (v NullableIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) Get() *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry {
	return v.value
}

func (v *NullableIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) Set(val *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry(val *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) *NullableIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry {
	return &NullableIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry{value: val, isSet: true}
}

func (v NullableIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


