package testbed

import "github.com/containerd/containerd"

// DefaultContainer exports a default containerd instance
func DefaultContainer(addr, namespace string) (*containerd.Client, error) {
	return containerd.New(addr,
		containerd.WithDefaultNamespace(namespace),
		containerd.WithDefaultRuntime("io.containerd.runc.v1"),
	)
}
