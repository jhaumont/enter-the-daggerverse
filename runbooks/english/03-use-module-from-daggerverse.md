# Use module from Daggerverse

## Discovering the Daggerverse

Go to the [Daggerverse](https://daggerverse.dev) website.

This page contains all Dagger's module provided by the community.

To improve our pipeline, we will search a **Go** module with some utilities functions.

Search a module with keyword [go](https://daggerverse.dev/search?q=go) which contains utilities functions like `build`, `test`, `generate`, etc.

> [!NOTE]
> As you can see, the search returns a large number of results:
> - Some correspond to older versions of the same module, which has changed directory in its source repository, or even repository,
> - Some are no longer maintained,
> 
> There is no way to know if modules are good or who create them (like [npmjs](https://www.npmjs.com))

Now, we will use this module:

![Module Dagger vito](../dagger-module-go-vito.jpg)

Direct link: https://daggerverse.dev/mod/github.com/vito/daggerverse/go@f7223d2d82fb91622cbb7954177388d995a98d59

## Go module installation

Run command:
```bash
dagger install github.com/vito/daggerverse/go@v0.0.1
```

> [!NOTE]
> It's also possible to use modules that are not available through Daggervers. More informations available in the official documentation [Using Modules from Remote Repositories](https://docs.dagger.io/api/remote-modules).

To discover this module, display help:
```bash
dagger -m go call --help
```

To display function names only, use:
```bash
dagger -m go functions
```

Another solution is to read directly the souce code: https://github.com/vito/daggerverse/blob/main/go/main.go

> [!NOTE]
> Here, we use a Go module write in Golang. This module can by write in other language (Typescript or Python).
>
> It is not necessary for the CI to be in the same language as the source code of the application to be built.

### Use Go module in application's pipeline

### Change BuildEnv function

Add a `builder` field of type `*dagger.Go` to the `Hello` structure:
```go
type Hello struct {
	builder *dagger.Go
}
```
This field allows us to manipulate an instance of the structure exposed by the Go module.

Replace `BuildEnv` function by bellow code in file `dagger/main.go`:
```go
// Build a ready-to-use development environment
func (m *Hello) BuildEnv() {
	m.builder = dag.Go().FromVersion("1.23-alpine")
}
```

Now, this function no longer returns a container, but modifies `builder` instance of the `Hello` struct with the desired version of the docker golang image.

> [!WARNING]
> The `dagger.gen.go` has compilation issue.
>
> When you change `BuildEnv` function, the interface of Dagger module has changed (removal of a variable in function).
>
> You must regenerate Dagger's module's code with command:
> ```bash
> dagger develop
> ```

### Update Build function

Replace the `Build` function by bellow code in the file `dagger/main.go`:
```go
// Build the application container
func (m *Hello) Build(source *dagger.Directory) *dagger.Container {
	m.BuildEnv()
	build := m.builder.Build(source, dagger.GoBuildOpts{Static: true})
	return dag.Container().From("debian:bookworm-slim").
		WithDirectory("/usr/bin/", build).
		WithExposedPort(666).
		WithEntrypoint([]string{"/usr/bin/hello"})
}
```

Now, this function call `BuildEnv` function to select build context.

Then, `Build` function of Go module previously imported (the module that we have chosen in the Daggerverse) to build application.

`Static: true` option is same as `WithEnvVariable("CGO_ENABLED", "0")` in previous `BuildEnv` function.

> [!NOTE]
> In Dagger, the arguments of a function can by [optionals](https://docs.dagger.io/manuals/developer/functions/#optional-arguments) by adding a comment in body of function.
> 
> It is also possibile to set default value.

> [!NOTE]
> An undocumented convention for optional's argument is to put this argument in struct.
> This struct is module's struct (here `dagger.Go`).
>
> The format is `dagger.<Package><Fonction>Opts`.
>
> Here, `dagger.GoBuildOpts` for `Static` parameter.
> The argument named `Static` is an `Opts` option of the `Build` function in the Dagger `Go` module.

### Test functions using Go module

Run `BuildEnv` function (argument `source` isn't needed):
```bash
dagger call build-env
```

Run function `Build`:
```bash
dagger call build --source=.
```

Take time to analyse the diff before/after modification in Dagger cloud traces.

Now, you will [Use module in github actions](04-use-module-in-github-actions.md).
