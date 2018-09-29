//+build wireinject

package main

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/rest"
	"github.com/google/go-cloud/wire"
)

// The build tag makes sure the stub is not built in the final build.

func initializeServer() (*rest.Server, error) {
	wire.Build(wireSet)
	return &rest.Server{}, nil
}