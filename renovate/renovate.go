// A generated module for Renovate functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/renovate/internal/dagger"
)

type Renovate struct{}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Renovate) RenovateScan(
	ctx context.Context,
	repository string,
	// +optional
	// +default="main"
	baseBranche string,
	renovateToken *dagger.Secret,
	// +optional
	// +default="info"
	logLevel string,
) (string, error) {
	return dag.Container().
		From("renovate/renovate:38").
		WithSecretVariable("RENOVATE_TOKEN", renovateToken).
		WithEnvVariable("RENOVATE_REPOSITORIES", repository).
		WithEnvVariable("RENOVATE_BASE_BRANCHES", baseBranche).
		WithEnvVariable("LOG_LEVEL", logLevel).
		WithExec([]string{"--platform=github", "--onboarding=false"}, dagger.ContainerWithExecOpts{UseEntrypoint: true}).
		Stdout(ctx)
}
