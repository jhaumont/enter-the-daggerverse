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
	baseBranch string,
	renovateToken *dagger.Secret,
	// +optional
	// +default="info"
	logLevel string,
) (string, error) {
	return dag.Container().
		From("renovate/renovate:38").
		WithSecretVariable("RENOVATE_TOKEN", renovateToken).
		WithEnvVariable("RENOVATE_REPOSITORIES", repository).
		WithEnvVariable("RENOVATE_BASE_BRANCHES", baseBranch).
		WithEnvVariable("LOG_LEVEL", logLevel).
		WithExec([]string{"--platform=github", "--onboarding=false"}, dagger.ContainerWithExecOpts{UseEntrypoint: true}).
		Stdout(ctx)
}
