# Create, publish a module in Daggerverse and use it

To dig more with Dagger functionalities, we propose to build a new module to answer a simple question: how manage dependencies?

To do so, we will create a Renovate module, pulish it in Daggerverse (because it's cool to share) and use it in our application's pipeline.

## Renovate module creation 

> [!NOTE]
> Please, use the same git branch than before.

At the root level of repository's codelab, create a `renovate` folder:
```bash
mkdir -p renovate
```

Init Dagger module:
```bash
cd renovate
dagger init --sdk=go --source=.
```

You must create a `RenovateScan` function in this module.

To do so, you can use code's skeleton below and copy it in the `renovate/main.go` file:

```go
package main

import (
	"context"
	"dagger/renovate/internal/dagger"
)

type Renovate struct{}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Renovate) RenovateScan(
	ctx context.Context,
	//TODO
) (string, error) {
	return //TODO
}
```

The elements at your disposal to build this function:

- The Docker command to implement:
    ```bash
    docker run -e RENOVATE_TOKEN \
    -e "LOG_LEVEL"=info \
    -e "RENOVATE_REPOSITORIES=[\"<votre-pseudo-github>/enter-the-daggerverse\"]" \
    -e "RENOVATE_BASE_BRANCHES=[\"main\"]" \
    renovate/renovate:38 --platform=github --onboarding=false
    ```
- `RenovateScan` function interface:
  - `repository`: mandatory string 
  - `baseBranche`: optional string with `main`  as default value
  - `renovateToken`: mandatory string. In our case, it's a GitHub's PAT (Personal Access Token) which allow the access to our repository
  - `logLevel`: optional string with `info` as default value
- The function returns scan logs

> [!NOTE]
> Somes usefull link to help:
> - https://docs.dagger.io/manuals/developer/secrets/
> - https://pkg.go.dev/dagger.io/dagger#Container.WithExec
> - https://docs.dagger.io/manuals/developer/functions/#optional-arguments

> [!NOTE]
> Don't forget to export your GitHub PAT in your shell to run the tests:
> ```bash
> read RENOVATE_TOKEN
> { paste GitLab access token here }
> export RENOVATE_TOKEN
> ```

> [!TIP]
> Don't forget to run `dagger develop` after function creation to refresh Dagger interface.

## Publish renovate's module

> [!NOTE]
> This part will be do by speakers in live

Follow official documentation: [Publishing Modules](https://docs.dagger.io/manuals/developer/publish-modules) and [Publish a Module](https://daggerverse.dev/publish).

## Use your module in the pipeline

Like before, we create a GitHub Action.

> [!NOTE]
> Use Renovate module already published on Daggerverse for this codelab
> 
> The [documentation](https://docs.dagger.io/integrations/github) of Dagger integration for GitHub

Create GitHub Actions' file:
```bash
touch ../.github/workflows/renovate.yaml
```

This is the skeleton of the pipeline to use in the file:

```yaml
name: Renovate Scan
on:
  ## We will not use the schedule but it will be the good practice :)
  # schedule:
    ## The "*" (#42, asterisk) character has special semantics in YAML, so this
    ## string has to be quoted.
    # - cron: '0/15 * * * *'
  pull_request:
    branches:
      - 'main'

jobs:
  renovate:
    name: Renovate scan
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
```

Build the execution step of Renovate scan based on your `RenovateScan` function.
