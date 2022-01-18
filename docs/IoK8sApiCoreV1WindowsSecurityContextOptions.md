# IoK8sApiCoreV1WindowsSecurityContextOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GmsaCredentialSpec** | Pointer to **string** | GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field. This field is alpha-level and is only honored by servers that enable the WindowsGMSA feature flag. | [optional] 
**GmsaCredentialSpecName** | Pointer to **string** | GMSACredentialSpecName is the name of the GMSA credential spec to use. This field is alpha-level and is only honored by servers that enable the WindowsGMSA feature flag. | [optional] 
**RunAsUserName** | Pointer to **string** | The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. This field is beta-level and may be disabled with the WindowsRunAsUserName feature flag. | [optional] 

## Methods

### NewIoK8sApiCoreV1WindowsSecurityContextOptions

`func NewIoK8sApiCoreV1WindowsSecurityContextOptions() *IoK8sApiCoreV1WindowsSecurityContextOptions`

NewIoK8sApiCoreV1WindowsSecurityContextOptions instantiates a new IoK8sApiCoreV1WindowsSecurityContextOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoK8sApiCoreV1WindowsSecurityContextOptionsWithDefaults

`func NewIoK8sApiCoreV1WindowsSecurityContextOptionsWithDefaults() *IoK8sApiCoreV1WindowsSecurityContextOptions`

NewIoK8sApiCoreV1WindowsSecurityContextOptionsWithDefaults instantiates a new IoK8sApiCoreV1WindowsSecurityContextOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGmsaCredentialSpec

`func (o *IoK8sApiCoreV1WindowsSecurityContextOptions) GetGmsaCredentialSpec() string`

GetGmsaCredentialSpec returns the GmsaCredentialSpec field if non-nil, zero value otherwise.

### GetGmsaCredentialSpecOk

`func (o *IoK8sApiCoreV1WindowsSecurityContextOptions) GetGmsaCredentialSpecOk() (*string, bool)`

GetGmsaCredentialSpecOk returns a tuple with the GmsaCredentialSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmsaCredentialSpec

`func (o *IoK8sApiCoreV1WindowsSecurityContextOptions) SetGmsaCredentialSpec(v string)`

SetGmsaCredentialSpec sets GmsaCredentialSpec field to given value.

### HasGmsaCredentialSpec

`func (o *IoK8sApiCoreV1WindowsSecurityContextOptions) HasGmsaCredentialSpec() bool`

HasGmsaCredentialSpec returns a boolean if a field has been set.

### GetGmsaCredentialSpecName

`func (o *IoK8sApiCoreV1WindowsSecurityContextOptions) GetGmsaCredentialSpecName() string`

GetGmsaCredentialSpecName returns the GmsaCredentialSpecName field if non-nil, zero value otherwise.

### GetGmsaCredentialSpecNameOk

`func (o *IoK8sApiCoreV1WindowsSecurityContextOptions) GetGmsaCredentialSpecNameOk() (*string, bool)`

GetGmsaCredentialSpecNameOk returns a tuple with the GmsaCredentialSpecName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmsaCredentialSpecName

`func (o *IoK8sApiCoreV1WindowsSecurityContextOptions) SetGmsaCredentialSpecName(v string)`

SetGmsaCredentialSpecName sets GmsaCredentialSpecName field to given value.

### HasGmsaCredentialSpecName

`func (o *IoK8sApiCoreV1WindowsSecurityContextOptions) HasGmsaCredentialSpecName() bool`

HasGmsaCredentialSpecName returns a boolean if a field has been set.

### GetRunAsUserName

`func (o *IoK8sApiCoreV1WindowsSecurityContextOptions) GetRunAsUserName() string`

GetRunAsUserName returns the RunAsUserName field if non-nil, zero value otherwise.

### GetRunAsUserNameOk

`func (o *IoK8sApiCoreV1WindowsSecurityContextOptions) GetRunAsUserNameOk() (*string, bool)`

GetRunAsUserNameOk returns a tuple with the RunAsUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsUserName

`func (o *IoK8sApiCoreV1WindowsSecurityContextOptions) SetRunAsUserName(v string)`

SetRunAsUserName sets RunAsUserName field to given value.

### HasRunAsUserName

`func (o *IoK8sApiCoreV1WindowsSecurityContextOptions) HasRunAsUserName() bool`

HasRunAsUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


