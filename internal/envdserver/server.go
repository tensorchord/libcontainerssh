// Copyright 2022 The envd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package envd

import (
	"github.com/gin-gonic/gin"

	"go.containerssh.io/libcontainerssh/config"
	"go.containerssh.io/libcontainerssh/http"
	"go.containerssh.io/libcontainerssh/log"
	"go.containerssh.io/libcontainerssh/message"
	"go.containerssh.io/libcontainerssh/service"
)

// New creates a new HTTP envd server service.
func New(logger log.Logger) (service.Service, error) {
	server := &envdServerService{
		router: gin.New(),
	}
	server.bindHandlers()
	svc, err := http.NewServer(
		"envd API server",
		config.HTTPServerConfiguration{
			Listen: "0.0.0.0:8888",
		},
		server.router,
		logger,
		func(url string) {
			logger.Info(message.NewMessage(message.MEnvdServerServiceAvailable, "envd endpoint available at %s", url))
		},
	)
	if err != nil {
		return nil, err
	}

	return &envdServerService{
		Service: svc,
	}, nil
}

type envdServerService struct {
	service.Service
	router *gin.Engine
}

func (s *envdServerService) bindHandlers() {
	engine := s.router
	engine.GET("/", handlePing)
	v1 := engine.Group("/v1")
	v1.GET("/", handlePing)
}

// TODO(gaocegege): Update it.
func (s *envdServerService) ChangeStatus(ok bool) {

}
