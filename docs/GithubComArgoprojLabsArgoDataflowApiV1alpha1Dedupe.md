# GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbstractStep** | Pointer to [**GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractStep**](GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractStep.md) |  | [optional] 
**MaxSize** | Pointer to **string** | Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors.  The serialization format is:  &lt;quantity&gt;        ::&#x3D; &lt;signedNumber&gt;&lt;suffix&gt;   (Note that &lt;suffix&gt; may be empty, from the \&quot;\&quot; case in &lt;decimalSI&gt;.) &lt;digit&gt;           ::&#x3D; 0 | 1 | ... | 9 &lt;digits&gt;          ::&#x3D; &lt;digit&gt; | &lt;digit&gt;&lt;digits&gt; &lt;number&gt;          ::&#x3D; &lt;digits&gt; | &lt;digits&gt;.&lt;digits&gt; | &lt;digits&gt;. | .&lt;digits&gt; &lt;sign&gt;            ::&#x3D; \&quot;+\&quot; | \&quot;-\&quot; &lt;signedNumber&gt;    ::&#x3D; &lt;number&gt; | &lt;sign&gt;&lt;number&gt; &lt;suffix&gt;          ::&#x3D; &lt;binarySI&gt; | &lt;decimalExponent&gt; | &lt;decimalSI&gt; &lt;binarySI&gt;        ::&#x3D; Ki | Mi | Gi | Ti | Pi | Ei   (International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) &lt;decimalSI&gt;       ::&#x3D; m | \&quot;\&quot; | k | M | G | T | P | E   (Note that 1024 &#x3D; 1Ki but 1000 &#x3D; 1k; I didn&#39;t choose the capitalization.) &lt;decimalExponent&gt; ::&#x3D; \&quot;e\&quot; &lt;signedNumber&gt; | \&quot;E\&quot; &lt;signedNumber&gt;  No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities.  When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized.  Before serializing, Quantity will be put in \&quot;canonical form\&quot;. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that:   a. No precision is lost   b. No fractional digits will be emitted   c. The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative.  Examples:   1.5 will be serialized as \&quot;1500m\&quot;   1.5Gi will be serialized as \&quot;1536Mi\&quot;  Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise.  Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don&#39;t diff.)  This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation. | [optional] 
**Uid** | Pointer to **string** |  | [optional] 

## Methods

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DedupeWithDefaults

`func NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DedupeWithDefaults() *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe`

NewGithubComArgoprojLabsArgoDataflowApiV1alpha1DedupeWithDefaults instantiates a new GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstractStep

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) GetAbstractStep() GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractStep`

GetAbstractStep returns the AbstractStep field if non-nil, zero value otherwise.

### GetAbstractStepOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) GetAbstractStepOk() (*GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractStep, bool)`

GetAbstractStepOk returns a tuple with the AbstractStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstractStep

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) SetAbstractStep(v GithubComArgoprojLabsArgoDataflowApiV1alpha1AbstractStep)`

SetAbstractStep sets AbstractStep field to given value.

### HasAbstractStep

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) HasAbstractStep() bool`

HasAbstractStep returns a boolean if a field has been set.

### GetMaxSize

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) GetMaxSize() string`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) GetMaxSizeOk() (*string, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) SetMaxSize(v string)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetUid

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *GithubComArgoprojLabsArgoDataflowApiV1alpha1Dedupe) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


