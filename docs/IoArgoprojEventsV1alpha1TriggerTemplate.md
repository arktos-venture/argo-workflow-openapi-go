# IoArgoprojEventsV1alpha1TriggerTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArgoWorkflow** | Pointer to [**IoArgoprojEventsV1alpha1ArgoWorkflowTrigger**](IoArgoprojEventsV1alpha1ArgoWorkflowTrigger.md) |  | [optional] 
**AwsLambda** | Pointer to [**IoArgoprojEventsV1alpha1AWSLambdaTrigger**](IoArgoprojEventsV1alpha1AWSLambdaTrigger.md) |  | [optional] 
**AzureEventHubs** | Pointer to [**IoArgoprojEventsV1alpha1AzureEventHubsTrigger**](IoArgoprojEventsV1alpha1AzureEventHubsTrigger.md) |  | [optional] 
**Conditions** | Pointer to **string** |  | [optional] 
**Custom** | Pointer to [**IoArgoprojEventsV1alpha1CustomTrigger**](IoArgoprojEventsV1alpha1CustomTrigger.md) |  | [optional] 
**Http** | Pointer to [**IoArgoprojEventsV1alpha1HTTPTrigger**](IoArgoprojEventsV1alpha1HTTPTrigger.md) |  | [optional] 
**K8s** | Pointer to [**IoArgoprojEventsV1alpha1StandardK8STrigger**](IoArgoprojEventsV1alpha1StandardK8STrigger.md) |  | [optional] 
**Kafka** | Pointer to [**IoArgoprojEventsV1alpha1KafkaTrigger**](IoArgoprojEventsV1alpha1KafkaTrigger.md) |  | [optional] 
**Log** | Pointer to [**IoArgoprojEventsV1alpha1LogTrigger**](IoArgoprojEventsV1alpha1LogTrigger.md) |  | [optional] 
**Name** | Pointer to **string** | Name is a unique name of the action to take. | [optional] 
**Nats** | Pointer to [**IoArgoprojEventsV1alpha1NATSTrigger**](IoArgoprojEventsV1alpha1NATSTrigger.md) |  | [optional] 
**OpenWhisk** | Pointer to [**IoArgoprojEventsV1alpha1OpenWhiskTrigger**](IoArgoprojEventsV1alpha1OpenWhiskTrigger.md) |  | [optional] 
**Slack** | Pointer to [**IoArgoprojEventsV1alpha1SlackTrigger**](IoArgoprojEventsV1alpha1SlackTrigger.md) |  | [optional] 
**Switch** | Pointer to [**IoArgoprojEventsV1alpha1TriggerSwitch**](IoArgoprojEventsV1alpha1TriggerSwitch.md) |  | [optional] 

## Methods

### NewIoArgoprojEventsV1alpha1TriggerTemplate

`func NewIoArgoprojEventsV1alpha1TriggerTemplate() *IoArgoprojEventsV1alpha1TriggerTemplate`

NewIoArgoprojEventsV1alpha1TriggerTemplate instantiates a new IoArgoprojEventsV1alpha1TriggerTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIoArgoprojEventsV1alpha1TriggerTemplateWithDefaults

`func NewIoArgoprojEventsV1alpha1TriggerTemplateWithDefaults() *IoArgoprojEventsV1alpha1TriggerTemplate`

NewIoArgoprojEventsV1alpha1TriggerTemplateWithDefaults instantiates a new IoArgoprojEventsV1alpha1TriggerTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgoWorkflow

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetArgoWorkflow() IoArgoprojEventsV1alpha1ArgoWorkflowTrigger`

GetArgoWorkflow returns the ArgoWorkflow field if non-nil, zero value otherwise.

### GetArgoWorkflowOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetArgoWorkflowOk() (*IoArgoprojEventsV1alpha1ArgoWorkflowTrigger, bool)`

GetArgoWorkflowOk returns a tuple with the ArgoWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgoWorkflow

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetArgoWorkflow(v IoArgoprojEventsV1alpha1ArgoWorkflowTrigger)`

SetArgoWorkflow sets ArgoWorkflow field to given value.

### HasArgoWorkflow

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasArgoWorkflow() bool`

HasArgoWorkflow returns a boolean if a field has been set.

### GetAwsLambda

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetAwsLambda() IoArgoprojEventsV1alpha1AWSLambdaTrigger`

GetAwsLambda returns the AwsLambda field if non-nil, zero value otherwise.

### GetAwsLambdaOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetAwsLambdaOk() (*IoArgoprojEventsV1alpha1AWSLambdaTrigger, bool)`

GetAwsLambdaOk returns a tuple with the AwsLambda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLambda

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetAwsLambda(v IoArgoprojEventsV1alpha1AWSLambdaTrigger)`

SetAwsLambda sets AwsLambda field to given value.

### HasAwsLambda

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasAwsLambda() bool`

HasAwsLambda returns a boolean if a field has been set.

### GetAzureEventHubs

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetAzureEventHubs() IoArgoprojEventsV1alpha1AzureEventHubsTrigger`

GetAzureEventHubs returns the AzureEventHubs field if non-nil, zero value otherwise.

### GetAzureEventHubsOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetAzureEventHubsOk() (*IoArgoprojEventsV1alpha1AzureEventHubsTrigger, bool)`

GetAzureEventHubsOk returns a tuple with the AzureEventHubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureEventHubs

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetAzureEventHubs(v IoArgoprojEventsV1alpha1AzureEventHubsTrigger)`

SetAzureEventHubs sets AzureEventHubs field to given value.

### HasAzureEventHubs

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasAzureEventHubs() bool`

HasAzureEventHubs returns a boolean if a field has been set.

### GetConditions

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetConditions() string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetConditionsOk() (*string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetConditions(v string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCustom

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetCustom() IoArgoprojEventsV1alpha1CustomTrigger`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetCustomOk() (*IoArgoprojEventsV1alpha1CustomTrigger, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetCustom(v IoArgoprojEventsV1alpha1CustomTrigger)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetHttp

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetHttp() IoArgoprojEventsV1alpha1HTTPTrigger`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetHttpOk() (*IoArgoprojEventsV1alpha1HTTPTrigger, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetHttp(v IoArgoprojEventsV1alpha1HTTPTrigger)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetK8s

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetK8s() IoArgoprojEventsV1alpha1StandardK8STrigger`

GetK8s returns the K8s field if non-nil, zero value otherwise.

### GetK8sOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetK8sOk() (*IoArgoprojEventsV1alpha1StandardK8STrigger, bool)`

GetK8sOk returns a tuple with the K8s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8s

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetK8s(v IoArgoprojEventsV1alpha1StandardK8STrigger)`

SetK8s sets K8s field to given value.

### HasK8s

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasK8s() bool`

HasK8s returns a boolean if a field has been set.

### GetKafka

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetKafka() IoArgoprojEventsV1alpha1KafkaTrigger`

GetKafka returns the Kafka field if non-nil, zero value otherwise.

### GetKafkaOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetKafkaOk() (*IoArgoprojEventsV1alpha1KafkaTrigger, bool)`

GetKafkaOk returns a tuple with the Kafka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafka

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetKafka(v IoArgoprojEventsV1alpha1KafkaTrigger)`

SetKafka sets Kafka field to given value.

### HasKafka

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasKafka() bool`

HasKafka returns a boolean if a field has been set.

### GetLog

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetLog() IoArgoprojEventsV1alpha1LogTrigger`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetLogOk() (*IoArgoprojEventsV1alpha1LogTrigger, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetLog(v IoArgoprojEventsV1alpha1LogTrigger)`

SetLog sets Log field to given value.

### HasLog

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetName

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNats

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetNats() IoArgoprojEventsV1alpha1NATSTrigger`

GetNats returns the Nats field if non-nil, zero value otherwise.

### GetNatsOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetNatsOk() (*IoArgoprojEventsV1alpha1NATSTrigger, bool)`

GetNatsOk returns a tuple with the Nats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNats

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetNats(v IoArgoprojEventsV1alpha1NATSTrigger)`

SetNats sets Nats field to given value.

### HasNats

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasNats() bool`

HasNats returns a boolean if a field has been set.

### GetOpenWhisk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetOpenWhisk() IoArgoprojEventsV1alpha1OpenWhiskTrigger`

GetOpenWhisk returns the OpenWhisk field if non-nil, zero value otherwise.

### GetOpenWhiskOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetOpenWhiskOk() (*IoArgoprojEventsV1alpha1OpenWhiskTrigger, bool)`

GetOpenWhiskOk returns a tuple with the OpenWhisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenWhisk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetOpenWhisk(v IoArgoprojEventsV1alpha1OpenWhiskTrigger)`

SetOpenWhisk sets OpenWhisk field to given value.

### HasOpenWhisk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasOpenWhisk() bool`

HasOpenWhisk returns a boolean if a field has been set.

### GetSlack

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetSlack() IoArgoprojEventsV1alpha1SlackTrigger`

GetSlack returns the Slack field if non-nil, zero value otherwise.

### GetSlackOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetSlackOk() (*IoArgoprojEventsV1alpha1SlackTrigger, bool)`

GetSlackOk returns a tuple with the Slack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlack

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetSlack(v IoArgoprojEventsV1alpha1SlackTrigger)`

SetSlack sets Slack field to given value.

### HasSlack

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasSlack() bool`

HasSlack returns a boolean if a field has been set.

### GetSwitch

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetSwitch() IoArgoprojEventsV1alpha1TriggerSwitch`

GetSwitch returns the Switch field if non-nil, zero value otherwise.

### GetSwitchOk

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) GetSwitchOk() (*IoArgoprojEventsV1alpha1TriggerSwitch, bool)`

GetSwitchOk returns a tuple with the Switch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitch

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) SetSwitch(v IoArgoprojEventsV1alpha1TriggerSwitch)`

SetSwitch sets Switch field to given value.

### HasSwitch

`func (o *IoArgoprojEventsV1alpha1TriggerTemplate) HasSwitch() bool`

HasSwitch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


