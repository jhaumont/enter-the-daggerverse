# Créer, publier un module dans le daggerverse et l'utiliser

Pour aller plus loin dans la découverte des fonctionnalités de Dagger, nous vous proposons la construction d'un module réutilisable répondant à un besoin simple : la gestion de vos dépendances.

Pour se faire, nous allons créer un module Renovate, le publier dans le Daggerverse (parce que c'est cool de partager) et l'utiliser dans le pipeline de notre application.

## Création du module Renovate

> [!NOTE]
> Continuez à travailler dans la branche de la partie précédente.

A la racine du repository du codelab, créez un répertoire `renovate` :
```bash
mkdir -p renovate
```

Initialisez votre nouveau module de Dagger :
```bash
cd renovate
dagger init --sdk=go --source=.
```

Vous allez devoir construire une fonction `RenovateScan` dans ce module.

Pour se faire, voici un squelette à copier dans le fichier `renovate/main.go` pour remplacer le code existant :

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

Les éléments à votre dispositon pour construire cette fonction :

- La commande docker à implémenter:
    ```bash
    docker run -e RENOVATE_TOKEN \
    -e "LOG_LEVEL"=info \
    -e "RENOVATE_REPOSITORIES=[\"<votre-pseudo-github>/enter-the-daggerverse\"]" \
    -e "RENOVATE_BASE_BRANCHES=[\"main\"]" \
    renovate/renovate:38 --platform=github --onboarding=false
    ```
- L'interface de la fonction `RenovateScan` :
  - `repository` : string obligatoire
  - `baseBranch` : string facultatif avec la valeur par défaut "main"
  - `renovateToken` : string facultatif. Dans notre cas, c'est un PAT (Personal Access Token) GitHub permettant d'accéder au repository
  - `logLevel` : string facultatif avec la valeur par défaut `info`
- Le résultat attendu de la fonction : les logs du scan

> [!NOTE]
> Quelques liens intéressants pour vous aider :
> - https://docs.dagger.io/manuals/developer/secrets/
> - https://pkg.go.dev/dagger.io/dagger#Container.WithExec
> - https://docs.dagger.io/manuals/developer/functions/#optional-arguments

> [!NOTE]
> Pensez à exporter votre PAT GitHub dans votre shell pour vos tests :
> ```bash
> read RENOVATE_TOKEN
> { paste GitLab access token here }
> export RENOVATE_TOKEN
> ```

> [!TIP]
> N'oubliez pas la commande `dagger develop` une fois que vous avez terminé votre fonction afin de regénérer le code d'interface Dagger.

## Publier le module renovate

> [!NOTE]
> Cette partie sera réalisée par les animateurs en direct

Suivez la documentation officielle : [Publishing Modules](https://docs.dagger.io/manuals/developer/publish-modules) et [Publish a Module](https://daggerverse.dev/publish).

## Utiliser votre module dans le pipeline

De la même manière que dans la partie précédente, nous allons construire une GitHub Actions.

> [!NOTE]
> Utilisez le module Renovate qui a déjà été publié sur le Daggerverse pour ce codelab
>
> La [documentation](https://docs.dagger.io/integrations/github) de l'intégration Dagger pour GitHub

Créez le fichier de la GitHub Actions :
```bash
touch ../.github/workflows/renovate.yaml
```

Voici le squelette à intégrer dans le fichier :

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

Construisez la step d'exécution du scan Renovate basée sur votre fonction `RenovateScan`.
