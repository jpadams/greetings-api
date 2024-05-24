package main

import (
	"context"
	"runtime"
)

type Backend struct{}

// Run the unit tests for the backend
func (b *Backend) UnitTest(ctx context.Context, source *Directory) (string, error) {
	return dag.
		Golang().
		WithProject(source).
		Test(ctx)
}

// Lint the backend Go code
func (b *Backend) Lint(ctx context.Context, source *Directory) (string, error) {
	return dag.
		Golang().
		WithProject(source).
		GolangciLint(ctx)
}

// Build the backend
func (b *Backend) Build(
	source *Directory,
	// +optional
	arch string,
) *Directory {
	if arch == "" {
		arch = runtime.GOARCH
	}
	return dag.
		Golang().
		WithProject(source).
		Build([]string{}, GolangBuildOpts{Arch: arch})
}

// Return the compiled backend binary for a particular architecture
func (b *Backend) Binary(
	source *Directory,
	// +optional
	arch string,
) *File {
	d := b.Build(source, arch)
	return d.File("greetings-api")
}

// Get a container ready to run the backend
func (b *Backend) Container(
	source *Directory,
	// +optional
	arch string,
) *Container {
	if arch == "" {
		arch = runtime.GOARCH
	}
	bin := b.Binary(source, arch)
	return dag.
		Container(ContainerOpts{Platform: Platform(arch)}).
		From("cgr.dev/chainguard/wolfi-base:latest@sha256:a8c9c2888304e62c133af76f520c9c9e6b3ce6f1a45e3eaa57f6639eb8053c90").
		WithFile("/bin/greetings-api", bin).
		WithEntrypoint([]string{"/bin/greetings-api"}).
		WithExposedPort(8080)
}

// Get a Service to run the backend
func (b *Backend) Serve(source *Directory) *Service {
	return b.Container(source, runtime.GOARCH).AsService()
}
