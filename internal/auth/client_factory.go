package auth

import (
	"fmt"

	"github.com/containerssh/libcontainerssh/config"
	"github.com/containerssh/libcontainerssh/internal/metrics"
	"github.com/containerssh/libcontainerssh/log"
	"github.com/containerssh/libcontainerssh/service"
)

func NewClient(
	cfg config.AuthConfig,
	logger log.Logger,
	metrics metrics.Collector,
) (Client, service.Service, error) {
	if err := cfg.Validate(); err != nil {
		return nil, nil, err
	}
	switch cfg.Method {
	case config.AuthMethodWebhook:
		client, err := NewHttpAuthClient(cfg, logger, metrics)
		return client, nil, err
	case config.AuthMethodOAuth2:
		return NewOAuth2Client(cfg, logger, metrics)
	default:
		return nil, nil, fmt.Errorf("unsupported method: %s", cfg.Method)
	}
}
