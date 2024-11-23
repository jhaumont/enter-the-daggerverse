# Use module from Daggervers

## Discovering the Daggerverse

Goto the [Daggerverse](https://daggerverse.dev) website.

This page contains all Dagger's module provide by community.

To improve our pipeline, we will search a **Go** module with some utilities functions.

Search a module with keyword [go](https://daggerverse.dev/search?q=go) and contains utilities functions like `build`, `test`, `generate`, etc.

> [!NOTE]
> As you can see, many results return by search:
> - somes modules are old version of same module that have changed folder in source repository or repoository,
> - somes are not maintain any more.
> 
> There is no way to know if module are good or who is create this module (like [npmjs](https://www.npmjs.com))

Now, we will use this module:

![Module Dagger vito](./dagger-module-go-vito.jpg)

Direct link: https://daggerverse.dev/mod/github.com/vito/daggerverse/go@f7223d2d82fb91622cbb7954177388d995a98d59

## Go module installation

Run command:
```bash
dagger install github.com/vito/daggerverse/go@v0.0.1
```

> [!NOTE]
> It's also possible to use modules that are not available through Daggervers. More informations available in official documentation [Using Modules from Remote Repositories](https://docs.dagger.io/api/remote-modules).

To discover this module, show its help:
```bash
dagger -m go call --help
```

To show only functions name, use:
```bash
dagger -m go functions
```

Another solution is to directly read souce code: https://github.com/vito/daggerverse/blob/main/go/main.go

> [!NOTE]
> Here, we use a Go module write in Go lang. This module can by write in other language (Typescript ou Python).
>
> Your CI can use another language that your application to build it.

### Use Go module in application's pipeline

### Change BuildEnv function

Add a field named `builder` with `*dagger.Go` type in struct called `Hello`:
```go
type Hello struct {
	builder *dagger.Go
}
```

This module help use to manage an instance exposed by Go module.

Replace `BuildEnv` function by bellow code in file `dagger/main.go`:
```go
// Build a ready-to-use development environment
func (m *Hello) BuildEnv() {
	m.builder = dag.Go().FromVersion("1.23-alpine")
}
```

Now, this function don't return a container but update `builder` field of struct `Hello`.

> [!WARNING]
> The `dagger.gen.go` as an issue to build.
>
> When you change `BuildEnv` function, the interface of Dogger module has changed (remove a variable in function).
>
> You must regenerate SDK helper:
> ```bash
> dagger develop
> ```

### Update Build function

Replace the `Build` function by bellow code in file `dagger/main.go`:
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

Then, `Build` function of Go previously imported (the module that we choose from Daggerverse) to build application.

`Static: true` option is same as `WithEnvVariable("CGO_ENABLED", "0")` in previous `BuildEnv` function.

> [!NOTE]
> In Dagger, the arguments of function can by [optionals](https://docs.dagger.io/manuals/developer/functions/#optional-arguments) by adding a comment in body of function.
> 
> There is also possibile to set default value.

> [!NOTE]
> An undocumented convention for optional's argument is to put this argument in struct.
> This struct is module's struct (here `dagger.Go`).
>
> Struct use format `dagger.<Package><Fonction>Opts`.
>
> Here, `dagger.GoBuildOpts` for `Static` parameter.
> We read `Static` argument is an option `Opts` of function `Build` of `Go` Dagger module.

### Test functions using Go module

Run `BuildEnv` function (argument `source` isn't needed):
```bash
dagger call build-env
```

Run function `Build`:
```bash
dagger call build --source=.
```

Take time to analyse diff before/after modification un Dagger cloud traces.

To the next, please go to the next page [Use module in github actions](04-use-module-in-github-actions.md).
