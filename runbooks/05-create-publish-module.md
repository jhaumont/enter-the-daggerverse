# Create, publish a module in Daggerverse and use it

To dig more functionalities of Dagger, we propose to build a new module to answer simple question: how manage dependencies?

To do that, we will create a Renovate module, pulish it in Daggerverse and use it in our application's pipeline.

## Module creation for Renovate

> [!NOTE]
> Please, use the same previously git branch.

In te root of repository's codelab, create `renovate`:
A la racine du repository du codelab, créez un répertoire renovate :
```bash
mkdir -p renovate
```

Init Dagger module:
```bash
cd renovate
dagger init --sdk=go --source=.
```

You must create a `RenovateScan` in this module.

To do this, you can use code's skeleton below and copy it in `renovate/main.go` file:

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

The following items are available for this function:

- Docker command to implemente:
    ```bash
    docker run -e RENOVATE_TOKEN \
    -e "LOG_LEVEL"=info \
    -e "RENOVATE_REPOSITORIES=[\"<votre-pseudo-github>/enter-the-daggerverse\"]" \
    -e "RENOVATE_BASE_BRANCHES=[\"main\"]" \
    renovate/renovate:38 --platform=github --onboarding=false
    ```
- `RenovateScan` interface's function:
  - `repository` : string mandatory
  - `baseBranche` : string optional with `main` default value
  - `renovateToken` : string mandatory. In our case, it's a GitHub's PAT (Personal Access Token) allow to access on our repository
  - `logLevel` : string optional with `info` default value
- The function return logs of scan

> [!NOTE]
> Somes usefull link to help you:
> - https://docs.dagger.io/manuals/developer/secrets/
> - https://pkg.go.dev/dagger.io/dagger#Container.WithExec
> - https://docs.dagger.io/manuals/developer/functions/#optional-arguments

> [!NOTE]
> Don't forget to expert your GitHub PAT in you shell to run your tests:
> ```bash
> read RENOVATE_TOKEN
> { paste GitLab access token here }
> export RENOVATE_TOKEN
> ```

> [!TIP]
> Don't forget to run `dagger develop` after you finish with your function to regenerate Dagger interface.

## Publish renovate's module

> [!NOTE]
> This part will be do by speakers in live

Follow official documentation: [Publishing Modules](https://docs.dagger.io/manuals/developer/publish-modules) et [Publish a Module](https://daggerverse.dev/publish).

## Use your module in pipeline

Like previously, we create a GitHub Actions.

> [!NOTE]
> Use Renovate module that already published on Daggerverse for this codelab
> 
> The [documentation](https://docs.dagger.io/integrations/github) of integration of Dagger for GitHub

Create GitHub Actions' file below:
```bash
touch ../.github/workflows/renovate.yaml
```

This is the skeleton of pipeline to push in the file:

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