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

// IoArgoprojWorkflowV1alpha1ContainerNode struct for IoArgoprojWorkflowV1alpha1ContainerNode
type IoArgoprojWorkflowV1alpha1ContainerNode struct {
	// Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Args *[]string `json:"args,omitempty"`
	// Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Command *[]string `json:"command,omitempty"`
	Dependencies *[]string `json:"dependencies,omitempty"`
	// List of environment variables to set in the container. Cannot be updated.
	Env *[]IoK8sApiCoreV1EnvVar `json:"env,omitempty"`
	// List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
	EnvFrom *[]IoK8sApiCoreV1EnvFromSource `json:"envFrom,omitempty"`
	// Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.
	Image *string `json:"image,omitempty"`
	// Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty"`
	Lifecycle *IoK8sApiCoreV1Lifecycle `json:"lifecycle,omitempty"`
	LivenessProbe *IoK8sApiCoreV1Probe `json:"livenessProbe,omitempty"`
	// Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
	Name string `json:"name"`
	// List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default \"0.0.0.0\" address inside a container will be accessible from the network. Cannot be updated.
	Ports *[]IoK8sApiCoreV1ContainerPort `json:"ports,omitempty"`
	ReadinessProbe *IoK8sApiCoreV1Probe `json:"readinessProbe,omitempty"`
	Resources *IoK8sApiCoreV1ResourceRequirements `json:"resources,omitempty"`
	SecurityContext *IoK8sApiCoreV1SecurityContext `json:"securityContext,omitempty"`
	StartupProbe *IoK8sApiCoreV1Probe `json:"startupProbe,omitempty"`
	// Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
	Stdin *bool `json:"stdin,omitempty"`
	// Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
	StdinOnce *bool `json:"stdinOnce,omitempty"`
	// Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
	TerminationMessagePath *string `json:"terminationMessagePath,omitempty"`
	// Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
	TerminationMessagePolicy *string `json:"terminationMessagePolicy,omitempty"`
	// Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
	Tty *bool `json:"tty,omitempty"`
	// volumeDevices is the list of block devices to be used by the container.
	VolumeDevices *[]IoK8sApiCoreV1VolumeDevice `json:"volumeDevices,omitempty"`
	// Pod volumes to mount into the container's filesystem. Cannot be updated.
	VolumeMounts *[]IoK8sApiCoreV1VolumeMount `json:"volumeMounts,omitempty"`
	// Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	WorkingDir *string `json:"workingDir,omitempty"`
}

// NewIoArgoprojWorkflowV1alpha1ContainerNode instantiates a new IoArgoprojWorkflowV1alpha1ContainerNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIoArgoprojWorkflowV1alpha1ContainerNode(name string) *IoArgoprojWorkflowV1alpha1ContainerNode {
	this := IoArgoprojWorkflowV1alpha1ContainerNode{}
	this.Name = name
	return &this
}

// NewIoArgoprojWorkflowV1alpha1ContainerNodeWithDefaults instantiates a new IoArgoprojWorkflowV1alpha1ContainerNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIoArgoprojWorkflowV1alpha1ContainerNodeWithDefaults() *IoArgoprojWorkflowV1alpha1ContainerNode {
	this := IoArgoprojWorkflowV1alpha1ContainerNode{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetArgs() []string {
	if o == nil || o.Args == nil {
		var ret []string
		return ret
	}
	return *o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetArgsOk() (*[]string, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the Args field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetArgs(v []string) {
	o.Args = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetCommand() []string {
	if o == nil || o.Command == nil {
		var ret []string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetCommandOk() (*[]string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasCommand() bool {
	if o != nil && o.Command != nil {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetCommand(v []string) {
	o.Command = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetDependencies() []string {
	if o == nil || o.Dependencies == nil {
		var ret []string
		return ret
	}
	return *o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetDependenciesOk() (*[]string, bool) {
	if o == nil || o.Dependencies == nil {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasDependencies() bool {
	if o != nil && o.Dependencies != nil {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given []string and assigns it to the Dependencies field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetDependencies(v []string) {
	o.Dependencies = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetEnv() []IoK8sApiCoreV1EnvVar {
	if o == nil || o.Env == nil {
		var ret []IoK8sApiCoreV1EnvVar
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetEnvOk() (*[]IoK8sApiCoreV1EnvVar, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasEnv() bool {
	if o != nil && o.Env != nil {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []IoK8sApiCoreV1EnvVar and assigns it to the Env field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetEnv(v []IoK8sApiCoreV1EnvVar) {
	o.Env = &v
}

// GetEnvFrom returns the EnvFrom field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetEnvFrom() []IoK8sApiCoreV1EnvFromSource {
	if o == nil || o.EnvFrom == nil {
		var ret []IoK8sApiCoreV1EnvFromSource
		return ret
	}
	return *o.EnvFrom
}

// GetEnvFromOk returns a tuple with the EnvFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetEnvFromOk() (*[]IoK8sApiCoreV1EnvFromSource, bool) {
	if o == nil || o.EnvFrom == nil {
		return nil, false
	}
	return o.EnvFrom, true
}

// HasEnvFrom returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasEnvFrom() bool {
	if o != nil && o.EnvFrom != nil {
		return true
	}

	return false
}

// SetEnvFrom gets a reference to the given []IoK8sApiCoreV1EnvFromSource and assigns it to the EnvFrom field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetEnvFrom(v []IoK8sApiCoreV1EnvFromSource) {
	o.EnvFrom = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetImage(v string) {
	o.Image = &v
}

// GetImagePullPolicy returns the ImagePullPolicy field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetImagePullPolicy() string {
	if o == nil || o.ImagePullPolicy == nil {
		var ret string
		return ret
	}
	return *o.ImagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetImagePullPolicyOk() (*string, bool) {
	if o == nil || o.ImagePullPolicy == nil {
		return nil, false
	}
	return o.ImagePullPolicy, true
}

// HasImagePullPolicy returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasImagePullPolicy() bool {
	if o != nil && o.ImagePullPolicy != nil {
		return true
	}

	return false
}

// SetImagePullPolicy gets a reference to the given string and assigns it to the ImagePullPolicy field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetImagePullPolicy(v string) {
	o.ImagePullPolicy = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetLifecycle() IoK8sApiCoreV1Lifecycle {
	if o == nil || o.Lifecycle == nil {
		var ret IoK8sApiCoreV1Lifecycle
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetLifecycleOk() (*IoK8sApiCoreV1Lifecycle, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasLifecycle() bool {
	if o != nil && o.Lifecycle != nil {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given IoK8sApiCoreV1Lifecycle and assigns it to the Lifecycle field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetLifecycle(v IoK8sApiCoreV1Lifecycle) {
	o.Lifecycle = &v
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetLivenessProbe() IoK8sApiCoreV1Probe {
	if o == nil || o.LivenessProbe == nil {
		var ret IoK8sApiCoreV1Probe
		return ret
	}
	return *o.LivenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetLivenessProbeOk() (*IoK8sApiCoreV1Probe, bool) {
	if o == nil || o.LivenessProbe == nil {
		return nil, false
	}
	return o.LivenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasLivenessProbe() bool {
	if o != nil && o.LivenessProbe != nil {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given IoK8sApiCoreV1Probe and assigns it to the LivenessProbe field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetLivenessProbe(v IoK8sApiCoreV1Probe) {
	o.LivenessProbe = &v
}

// GetName returns the Name field value
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetName(v string) {
	o.Name = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetPorts() []IoK8sApiCoreV1ContainerPort {
	if o == nil || o.Ports == nil {
		var ret []IoK8sApiCoreV1ContainerPort
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetPortsOk() (*[]IoK8sApiCoreV1ContainerPort, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []IoK8sApiCoreV1ContainerPort and assigns it to the Ports field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetPorts(v []IoK8sApiCoreV1ContainerPort) {
	o.Ports = &v
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetReadinessProbe() IoK8sApiCoreV1Probe {
	if o == nil || o.ReadinessProbe == nil {
		var ret IoK8sApiCoreV1Probe
		return ret
	}
	return *o.ReadinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetReadinessProbeOk() (*IoK8sApiCoreV1Probe, bool) {
	if o == nil || o.ReadinessProbe == nil {
		return nil, false
	}
	return o.ReadinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasReadinessProbe() bool {
	if o != nil && o.ReadinessProbe != nil {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given IoK8sApiCoreV1Probe and assigns it to the ReadinessProbe field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetReadinessProbe(v IoK8sApiCoreV1Probe) {
	o.ReadinessProbe = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetResources() IoK8sApiCoreV1ResourceRequirements {
	if o == nil || o.Resources == nil {
		var ret IoK8sApiCoreV1ResourceRequirements
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetResourcesOk() (*IoK8sApiCoreV1ResourceRequirements, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given IoK8sApiCoreV1ResourceRequirements and assigns it to the Resources field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetResources(v IoK8sApiCoreV1ResourceRequirements) {
	o.Resources = &v
}

// GetSecurityContext returns the SecurityContext field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetSecurityContext() IoK8sApiCoreV1SecurityContext {
	if o == nil || o.SecurityContext == nil {
		var ret IoK8sApiCoreV1SecurityContext
		return ret
	}
	return *o.SecurityContext
}

// GetSecurityContextOk returns a tuple with the SecurityContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetSecurityContextOk() (*IoK8sApiCoreV1SecurityContext, bool) {
	if o == nil || o.SecurityContext == nil {
		return nil, false
	}
	return o.SecurityContext, true
}

// HasSecurityContext returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasSecurityContext() bool {
	if o != nil && o.SecurityContext != nil {
		return true
	}

	return false
}

// SetSecurityContext gets a reference to the given IoK8sApiCoreV1SecurityContext and assigns it to the SecurityContext field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetSecurityContext(v IoK8sApiCoreV1SecurityContext) {
	o.SecurityContext = &v
}

// GetStartupProbe returns the StartupProbe field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetStartupProbe() IoK8sApiCoreV1Probe {
	if o == nil || o.StartupProbe == nil {
		var ret IoK8sApiCoreV1Probe
		return ret
	}
	return *o.StartupProbe
}

// GetStartupProbeOk returns a tuple with the StartupProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetStartupProbeOk() (*IoK8sApiCoreV1Probe, bool) {
	if o == nil || o.StartupProbe == nil {
		return nil, false
	}
	return o.StartupProbe, true
}

// HasStartupProbe returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasStartupProbe() bool {
	if o != nil && o.StartupProbe != nil {
		return true
	}

	return false
}

// SetStartupProbe gets a reference to the given IoK8sApiCoreV1Probe and assigns it to the StartupProbe field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetStartupProbe(v IoK8sApiCoreV1Probe) {
	o.StartupProbe = &v
}

// GetStdin returns the Stdin field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetStdin() bool {
	if o == nil || o.Stdin == nil {
		var ret bool
		return ret
	}
	return *o.Stdin
}

// GetStdinOk returns a tuple with the Stdin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetStdinOk() (*bool, bool) {
	if o == nil || o.Stdin == nil {
		return nil, false
	}
	return o.Stdin, true
}

// HasStdin returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasStdin() bool {
	if o != nil && o.Stdin != nil {
		return true
	}

	return false
}

// SetStdin gets a reference to the given bool and assigns it to the Stdin field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetStdin(v bool) {
	o.Stdin = &v
}

// GetStdinOnce returns the StdinOnce field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetStdinOnce() bool {
	if o == nil || o.StdinOnce == nil {
		var ret bool
		return ret
	}
	return *o.StdinOnce
}

// GetStdinOnceOk returns a tuple with the StdinOnce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetStdinOnceOk() (*bool, bool) {
	if o == nil || o.StdinOnce == nil {
		return nil, false
	}
	return o.StdinOnce, true
}

// HasStdinOnce returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasStdinOnce() bool {
	if o != nil && o.StdinOnce != nil {
		return true
	}

	return false
}

// SetStdinOnce gets a reference to the given bool and assigns it to the StdinOnce field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetStdinOnce(v bool) {
	o.StdinOnce = &v
}

// GetTerminationMessagePath returns the TerminationMessagePath field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetTerminationMessagePath() string {
	if o == nil || o.TerminationMessagePath == nil {
		var ret string
		return ret
	}
	return *o.TerminationMessagePath
}

// GetTerminationMessagePathOk returns a tuple with the TerminationMessagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetTerminationMessagePathOk() (*string, bool) {
	if o == nil || o.TerminationMessagePath == nil {
		return nil, false
	}
	return o.TerminationMessagePath, true
}

// HasTerminationMessagePath returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasTerminationMessagePath() bool {
	if o != nil && o.TerminationMessagePath != nil {
		return true
	}

	return false
}

// SetTerminationMessagePath gets a reference to the given string and assigns it to the TerminationMessagePath field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetTerminationMessagePath(v string) {
	o.TerminationMessagePath = &v
}

// GetTerminationMessagePolicy returns the TerminationMessagePolicy field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetTerminationMessagePolicy() string {
	if o == nil || o.TerminationMessagePolicy == nil {
		var ret string
		return ret
	}
	return *o.TerminationMessagePolicy
}

// GetTerminationMessagePolicyOk returns a tuple with the TerminationMessagePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetTerminationMessagePolicyOk() (*string, bool) {
	if o == nil || o.TerminationMessagePolicy == nil {
		return nil, false
	}
	return o.TerminationMessagePolicy, true
}

// HasTerminationMessagePolicy returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasTerminationMessagePolicy() bool {
	if o != nil && o.TerminationMessagePolicy != nil {
		return true
	}

	return false
}

// SetTerminationMessagePolicy gets a reference to the given string and assigns it to the TerminationMessagePolicy field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetTerminationMessagePolicy(v string) {
	o.TerminationMessagePolicy = &v
}

// GetTty returns the Tty field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetTty() bool {
	if o == nil || o.Tty == nil {
		var ret bool
		return ret
	}
	return *o.Tty
}

// GetTtyOk returns a tuple with the Tty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetTtyOk() (*bool, bool) {
	if o == nil || o.Tty == nil {
		return nil, false
	}
	return o.Tty, true
}

// HasTty returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasTty() bool {
	if o != nil && o.Tty != nil {
		return true
	}

	return false
}

// SetTty gets a reference to the given bool and assigns it to the Tty field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetTty(v bool) {
	o.Tty = &v
}

// GetVolumeDevices returns the VolumeDevices field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetVolumeDevices() []IoK8sApiCoreV1VolumeDevice {
	if o == nil || o.VolumeDevices == nil {
		var ret []IoK8sApiCoreV1VolumeDevice
		return ret
	}
	return *o.VolumeDevices
}

// GetVolumeDevicesOk returns a tuple with the VolumeDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetVolumeDevicesOk() (*[]IoK8sApiCoreV1VolumeDevice, bool) {
	if o == nil || o.VolumeDevices == nil {
		return nil, false
	}
	return o.VolumeDevices, true
}

// HasVolumeDevices returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasVolumeDevices() bool {
	if o != nil && o.VolumeDevices != nil {
		return true
	}

	return false
}

// SetVolumeDevices gets a reference to the given []IoK8sApiCoreV1VolumeDevice and assigns it to the VolumeDevices field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetVolumeDevices(v []IoK8sApiCoreV1VolumeDevice) {
	o.VolumeDevices = &v
}

// GetVolumeMounts returns the VolumeMounts field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetVolumeMounts() []IoK8sApiCoreV1VolumeMount {
	if o == nil || o.VolumeMounts == nil {
		var ret []IoK8sApiCoreV1VolumeMount
		return ret
	}
	return *o.VolumeMounts
}

// GetVolumeMountsOk returns a tuple with the VolumeMounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetVolumeMountsOk() (*[]IoK8sApiCoreV1VolumeMount, bool) {
	if o == nil || o.VolumeMounts == nil {
		return nil, false
	}
	return o.VolumeMounts, true
}

// HasVolumeMounts returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasVolumeMounts() bool {
	if o != nil && o.VolumeMounts != nil {
		return true
	}

	return false
}

// SetVolumeMounts gets a reference to the given []IoK8sApiCoreV1VolumeMount and assigns it to the VolumeMounts field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetVolumeMounts(v []IoK8sApiCoreV1VolumeMount) {
	o.VolumeMounts = &v
}

// GetWorkingDir returns the WorkingDir field value if set, zero value otherwise.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetWorkingDir() string {
	if o == nil || o.WorkingDir == nil {
		var ret string
		return ret
	}
	return *o.WorkingDir
}

// GetWorkingDirOk returns a tuple with the WorkingDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) GetWorkingDirOk() (*string, bool) {
	if o == nil || o.WorkingDir == nil {
		return nil, false
	}
	return o.WorkingDir, true
}

// HasWorkingDir returns a boolean if a field has been set.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) HasWorkingDir() bool {
	if o != nil && o.WorkingDir != nil {
		return true
	}

	return false
}

// SetWorkingDir gets a reference to the given string and assigns it to the WorkingDir field.
func (o *IoArgoprojWorkflowV1alpha1ContainerNode) SetWorkingDir(v string) {
	o.WorkingDir = &v
}

func (o IoArgoprojWorkflowV1alpha1ContainerNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.EnvFrom != nil {
		toSerialize["envFrom"] = o.EnvFrom
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.ImagePullPolicy != nil {
		toSerialize["imagePullPolicy"] = o.ImagePullPolicy
	}
	if o.Lifecycle != nil {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	if o.LivenessProbe != nil {
		toSerialize["livenessProbe"] = o.LivenessProbe
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if o.ReadinessProbe != nil {
		toSerialize["readinessProbe"] = o.ReadinessProbe
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.SecurityContext != nil {
		toSerialize["securityContext"] = o.SecurityContext
	}
	if o.StartupProbe != nil {
		toSerialize["startupProbe"] = o.StartupProbe
	}
	if o.Stdin != nil {
		toSerialize["stdin"] = o.Stdin
	}
	if o.StdinOnce != nil {
		toSerialize["stdinOnce"] = o.StdinOnce
	}
	if o.TerminationMessagePath != nil {
		toSerialize["terminationMessagePath"] = o.TerminationMessagePath
	}
	if o.TerminationMessagePolicy != nil {
		toSerialize["terminationMessagePolicy"] = o.TerminationMessagePolicy
	}
	if o.Tty != nil {
		toSerialize["tty"] = o.Tty
	}
	if o.VolumeDevices != nil {
		toSerialize["volumeDevices"] = o.VolumeDevices
	}
	if o.VolumeMounts != nil {
		toSerialize["volumeMounts"] = o.VolumeMounts
	}
	if o.WorkingDir != nil {
		toSerialize["workingDir"] = o.WorkingDir
	}
	return json.Marshal(toSerialize)
}

type NullableIoArgoprojWorkflowV1alpha1ContainerNode struct {
	value *IoArgoprojWorkflowV1alpha1ContainerNode
	isSet bool
}

func (v NullableIoArgoprojWorkflowV1alpha1ContainerNode) Get() *IoArgoprojWorkflowV1alpha1ContainerNode {
	return v.value
}

func (v *NullableIoArgoprojWorkflowV1alpha1ContainerNode) Set(val *IoArgoprojWorkflowV1alpha1ContainerNode) {
	v.value = val
	v.isSet = true
}

func (v NullableIoArgoprojWorkflowV1alpha1ContainerNode) IsSet() bool {
	return v.isSet
}

func (v *NullableIoArgoprojWorkflowV1alpha1ContainerNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIoArgoprojWorkflowV1alpha1ContainerNode(val *IoArgoprojWorkflowV1alpha1ContainerNode) *NullableIoArgoprojWorkflowV1alpha1ContainerNode {
	return &NullableIoArgoprojWorkflowV1alpha1ContainerNode{value: val, isSet: true}
}

func (v NullableIoArgoprojWorkflowV1alpha1ContainerNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIoArgoprojWorkflowV1alpha1ContainerNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


