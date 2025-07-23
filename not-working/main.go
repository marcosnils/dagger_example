// A generated module for NotWorking functions
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
	"dagger/not-working/internal/dagger"
)

type NotWorking struct {
	// +private
	Src *dagger.Directory
}

func New(
	// +defaultPath="."
	src *dagger.Directory,
) *NotWorking {
	// Create a new NotWorking module with a default source directory
	return &NotWorking{
		Src: src,
	}
}

// Create a base container that builds from a Dockerfile
func (m *NotWorking) BaseContainer() *dagger.Container {
	return dag.Container().Build(
		m.Src,
		dagger.ContainerBuildOpts{Dockerfile: "nw.Dockerfile"},
	)
}

// Call a function from the other module
func (m *NotWorking) SayHello(ctx context.Context) string {
	ctr := m.BaseContainer()
	out, err := dag.OtherModule().ContainerEcho(ctr).Stdout(ctx)
	if err != nil {
		return "Error calling OtherModule: " + err.Error()
	}
	return out
}
