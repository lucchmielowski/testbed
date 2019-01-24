package runtime

import (
	"context"

	"github.com/containerd/containerd"
	api "github.com/lucchmielowski/testbed/api/services/runtime/v1"

	"github.com/lucchmielowski/testbed"
	"github.com/lucchmielowski/testbed/services"
	"google.golang.org/grpc"
)

const (
	serviceID = "testbed.services.runtime.v1"
)

type service struct {
	containerAddr string
	namespace     string
	bridge        string
	dataDir       string
	stateDir      string
	config        *testbed.Config
}

// New create a new instance of the runtime service
func New(cfg *testbed.Config) (services.Service, error) {
	return &service{
		containerAddr: cfg.ContainerAddr,
		namespace:     cfg.Namespace,
		bridge:        cfg.Bridge,
		dataDir:       cfg.DataDir,
		stateDir:      cfg.StateDir,
		config:        cfg,
	}, nil
}

func (s *service) Register(server *grpc.Server) error {
	api.RegisterNodeServer(server, s)
	return nil
}

func (s *service) ID() string {
	return serviceID
}

func (s *service) Type() services.Type {
	return services.RuntimeService
}

func (s *service) Requires() []services.Type {
	return nil
}

func (s *service) Info(ctx context.Context, req *api.InfoRequest) (*api.InfoResponse, error) {
	return &api.InfoResponse{
		ID: serviceID,
	}, nil
}

func (s *service) Start() error {
	return nil
}

func (s *service) Stop() error {
	return nil
}

func (s *service) containerd() (*containerd.Client, error) {
	return testbed.DefaultContainerd(s.containerAddr, s.namespace)
}
