package main

import (
	"context"
	"dagger/hello/internal/dagger"
	"fmt"
	"math"
	"math/rand/v2"
)

type Hello struct {
	builder *dagger.Go
}

// Publish the application container after building and testing it on-the-fly
func (m *Hello) Publish(ctx context.Context, source *dagger.Directory) (string, error) {
	return m.Build(source).
		Publish(ctx, fmt.Sprintf("ttl.sh/hello-%.0f", math.Floor(rand.Float64()*10000000))) //#nosec
}

// Build the application container
// func (m *Hello) Build(source *dagger.Directory) *dagger.Container {
// 	build := m.BuildEnv(source).
// 		WithExec([]string{"go", "build", "-o", "bin/"}).
// 		Directory("./bin")
// 	return dag.Container().From("debian:bookworm-slim").
// 		WithDirectory("/usr/bin/", build).
// 		WithExposedPort(666).
// 		WithEntrypoint([]string{"/usr/bin/hello"})
// }
// Build the application container
func (m *Hello) Build(source *dagger.Directory) *dagger.Container {
	m.BuildEnv()
	build := m.builder.Build(source, dagger.GoBuildOpts{Static: true})
	return dag.Container().From("debian:bookworm-slim").
		WithDirectory("/usr/bin/", build).
		WithExposedPort(666).
		WithEntrypoint([]string{"/usr/bin/hello"})
}

// Build a ready-to-use development environment
//
//	func (m *Hello) BuildEnv(source *dagger.Directory) *dagger.Container {
//		return dag.Container().
//			From("golang:1.23-alpine").
//			WithDirectory("/src", source).
//			WithWorkdir("/src").
//			WithEnvVariable("CGO_ENABLED", "0")
//	}
//
// Build a ready-to-use development environment
func (m *Hello) BuildEnv() {
	m.builder = dag.Go().FromVersion("1.23-alpine")
}
