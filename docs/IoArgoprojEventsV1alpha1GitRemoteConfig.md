# IoArgoprojEventsV1alpha1GitRemoteConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the remote to fetch from. | [optional] 
**Urls** | Pointer to **[]string** | URLs the URLs of a remote repository. It must be non-empty. Fetch will always use the first URL, while push will use all of them. | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1GitRemoteConfig

`func NewIoArgoprojEventsV1alpha1GitRemoteConfig() *IoArgoprojEventsV1alpha1GitRemoteConfig`

NewIoArgoprojEventsV1alpha1GitRemoteConfig instantiates a new IoArgoprojEventsV1alpha1GitRemoteConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1GitRemoteConfigWithDefaults

`func NewIoArgoprojEventsV1alpha1GitRemoteConfigWithDefaults() *IoArgoprojEventsV1alpha1GitRemoteConfig`

NewIoArgoprojEventsV1alpha1GitRemoteConfigWithDefaults instantiates a new IoArgoprojEventsV1alpha1GitRemoteConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrls

`func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *IoArgoprojEventsV1alpha1GitRemoteConfig) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


