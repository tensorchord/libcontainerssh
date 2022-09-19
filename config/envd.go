package config

// EnvdConfig is the base configuration structure for envd backend.
type EnvdConfig struct {
	// Connection configures the connection to the Kubernetes cluster.
	Connection KubernetesConnectionConfig `json:"connection,omitempty" yaml:"connection" comment:"Kubernetes configuration options"`
	// Pod contains the spec and specific settings for creating the pod.
	Pod KubernetesPodConfig `json:"pod,omitempty" yaml:"pod" comment:"Container configuration"`
	// Timeout specifies how long to wait for the Pod to come up.
	Timeouts KubernetesTimeoutConfig `json:"timeouts,omitempty" yaml:"timeouts" comment:"Timeout for pod creation"`
}

// Validate checks the configuration options and returns an error if the configuration is invalid.
func (c EnvdConfig) Validate() error {
	if err := c.Connection.Validate(); err != nil {
		return wrap(err, "connection")
	}
	if err := c.Pod.Validate(); err != nil {
		return wrap(err, "pod")
	}
	if err := c.Timeouts.Validate(); err != nil {
		return wrap(err, "timeouts")
	}
	return nil
}
