package services

import "google.golang.org/grpc"

// Type type is the string representing the service name
type Type string

const (
	// RuntimeService is the service that controls container handling
	RuntimeService Type = "testbed.services.runtime.v1"
)

// Service is the interface that all testbed services must implement
type Service interface {
	ID() string
	Type() Type
	Register(*grpc.Server) error
	Requires() []Type
	Start() error
	Stop() error
}
