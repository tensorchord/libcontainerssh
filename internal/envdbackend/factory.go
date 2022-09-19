package envdbackend

import (
	"net"
	"sync"

	publicconfig "go.containerssh.io/libcontainerssh/config"
	"go.containerssh.io/libcontainerssh/internal/metrics"
	"go.containerssh.io/libcontainerssh/internal/sshserver"
	"go.containerssh.io/libcontainerssh/log"
	"go.containerssh.io/libcontainerssh/message"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func New(
	client net.TCPAddr,
	connectionID string,
	config publicconfig.EnvdConfig,
	logger log.Logger,
	backendRequestsMetric metrics.SimpleCounter,
	backendFailuresMetric metrics.SimpleCounter,
) (
	sshserver.NetworkConnectionHandler,
	error,
) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", config.Connection.KubeConfigPath)
	if err != nil {
		return nil, err
	}

	cli, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		err = message.WrapUser(
			err,
			message.EKubernetesConfigError,
			UserMessageInitializeSSHSession,
			"Failed to initialize Kubernetes client.",
		)
		logger.Error(err)
		return nil, err
	}

	return &networkHandler{
		mutex:        &sync.Mutex{},
		client:       client,
		connectionID: connectionID,
		config:       config,
		cli:          cli,
		labels:       nil,
		logger:       logger,
		disconnected: false,
		done:         make(chan struct{}),
	}, nil
}
