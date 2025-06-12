# Use module from Daggerverse

## Discovering the Daggerverse

Go to the [Daggerverse](https://daggerverse.dev) website.

This page contains all Dagger modules published by the community.

To improve our pipeline, we will search for a **Go** module with some utilities functions.

Search for a module with the keyword [go](https://daggerverse.dev/search?q=go) and find one which contains functions like `build`, `test`, `generate`, etc.

> [!NOTE]
> As you can see, the search returns a large number of results:
> - Some correspond to older versions of the same module, which has changed directory in its source repository, or even repository,
> - Some are no longer maintained,
>
> There is currently no way to know if modules are good or who uses them in their projects (like [npmjs](https://www.npmjs.com))

Now, we will use this module:

![Module Dagger vito](../dagger-module-go-vito.jpg)

Direct link: https://daggerverse.dev/mod/github.com/vito/daggerverse/go@f7223d2d82fb91622cbb7954177388d995a98d59

## Go module installation

Run command:
```bash
dagger install github.com/vito/daggerverse/go@v0.0.1
```

> [!NOTE]
> It's also possible to use modules that are not available through Daggerverse (e.g. private git repositories). More information is available in the documentation [Using Modules from Remote Repositories](https://docs.dagger.io/api/remote-modules).

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
> Here, we use a module written in Golang. We can use module written in other supported languages (TypeScript or Python) and still install in our module - regardless of the other modules programming language.
>
> Also, it is not necessary for the module to be the same language as the source code of the application to be built. We could use the Dagger Go SDK to build a Python application, for example.

### Use Go module in application's pipeline

### Change BuildEnv function

Add a `builder` field of type `*dagger.Go` to the `Hello` struct in file `.dagger/main.go`:
```go
type Hello struct {
	builder *dagger.Go
}
```
This field allows us to use functions exposed by the Go module.

Replace `BuildEnv` function by bellow code in file same file:
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
> You must regenerate the Dagger module's code with command:
> ```bash
> dagger develop
> ```

### Update Build function

Replace the `Build` function by bellow code in the file `.dagger/main.go`:
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

Now, this function calls the `BuildEnv` function to build the environment. Then, `Build` function of Go module previously imported (the module that we have chosen in the Daggerverse) to build the application.

`Static: true` option is same as `WithEnvVariable("CGO_ENABLED", "0")` in previous `BuildEnv` function.

> [!NOTE]
> In Dagger, the arguments of a function can be [optional](https://docs.dagger.io/manuals/developer/functions/#optional-arguments) by adding a comment in body of function.
>
> It is also possibile to set [default values](https://docs.dagger.io/api/arguments/#default-values).

> [!NOTE]
> An undocumented convention for optional's argument is to put this argument in struct.
> This struct is module's struct (here `dagger.Go`).
>
> The format is `dagger.<Package><Function>Opts`.
>
> Here, `dagger.GoBuildOpts` for `Static` parameter.
> The argument named `Static` is an `Opts` option of the `Build` function in the Dagger `Go` module.

### Test functions using the Go module

Run the `BuildEnv` function (note the argument `source` isn't needed):
```bash
dagger call build-env
```

Run function `Build`:
```bash
dagger call build --source=.
```

Take time to analyse the diff before/after modification in Dagger cloud.

Now, you will [Use our module in GitHub Actions](04-use-module-in-github-actions.md).
