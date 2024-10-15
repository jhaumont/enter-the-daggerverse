# Utiliser un module du Daggerverse

## Découverte du Daggerverse

Aller sur la page https://daggerverse.dev

C'est l'index (gratuit) des modules Dagger mis à disposition par la communauté.

Afin d'améliorer notre pipeline, nous allons regarder si un module GO avec des fonctions utilitaires existent.

Chercher un module avec le mot clé `go`: https://daggerverse.dev/search?q=go.

> [!WARNING]
> Il y a de très nombreux résultats:
> - Certains correspondent à d'anciennes versions du même module qui a changé de répertoire dans son repository source voir de repository.
> - Certains ne sont plus maintenus.
> - Aucun moyen de connaitre la fiabilité du développeur à l'origine du module.

Pour la suite, nous allons utiliser ce module qui correspond à nos besoins https://daggerverse.dev/mod/github.com/vito/daggerverse/go@f7223d2d82fb91622cbb7954177388d995a98d59 (avec des fonctions utilitaires de type `build`, `test`, `generate`, etc).

## Installation du module Go

Lancer la commande :

```bash
dagger install github.com/vito/daggerverse/go@v0.0.1
```

Pour découvrir le module, afficher son aide :

```bash
dagger -m go call --help
```

Pour n'avoir que la liste des fonctions, utiliser:

```bash
dagger -m go functions
```

Ou autre solution, regarder directement le code source : https://github.com/vito/daggerverse/blob/main/go/main.go

> [!NOTE]
> Ici on utilise un module Go programmé Go mais au final, pour un autre besoin, le module aurait pu être écrit dans un tout autre language (Typescript ou Python)

## Utiliser le module Go dans le pipeline de l'application

### Modifier la fonction BuildEnv

Ajouter un champ `builder` de type `Go` dans la structure `Hello`:

```go
type Hello struct {
	builder *dagger.Go
}
```

Il va nous permettre de manipuler une instance de la structure exposée par le module Go

Remplacer la fonction BuildEnv par le code suivant dans le ficher `dagger/main.go`:

```go
// Build a ready-to-use development environment
func (m *Hello) BuildEnv() {
	m.builder = dag.Go().FromVersion("1.23-alpine")
}
```

Dorénavant, cette fonction ne retourne plus un container mais modifie l'instance `builder` de la structure `Go` avec la version désirée de l'image docker goland.

> [!WARNING]
> Le fichier `dagger.gen.go` a un problème de compilation.
>
> En modifiant la fonction BuildEnv, l'interface du module Dagger a changé (suppression de la variable source de cette fonction).
> 
> Regénérer le code Dagger du module:
>
> ```bash
> dagger develop
> ```


### Modifier la fonction Build

Remplacer la fonction Build par le code suivant dans le ficher `dagger/main.go`:

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

Dorénavant, cette fonction utilise la fonction `BuildEnv` pour sélectionner son contexte de build puis la fonction `Build` du module Go pour réaliser le build en tant que tel. L'option `Static: true` correspond à la ligne `WithEnvVariable("CGO_ENABLED", "0")` dans la version précédent du `BuildEnv`.

> [!NOTE]
> Dans Dagger, les arguments d'une fonction peuvent être optionnels en l'indiquant par un commentaire dans le code de la fonction. Exemple: https://docs.dagger.io/manuals/developer/functions/#optional-arguments
> Même chose pour la valeur par défaut.

> [!NOTE]
> Une convention (non documentée) pour les argument optionnels, est de mettre l'argument dans une strucutre.
> Cette structure est celle du module.
> Elle est du format `dagger.<Package><Fonction>Opts`. Exemple ici: `dagger.GoBuildOpts` pour utiliser le paramètre `Static`.
> On lit donc l'argument `Static` est une option `Opts` de la fonction `Build` du module Dagger `Go`.

### Tester les fonctions utilisant le module Go

Lancer la fonction BuildEnv (il n'y a plus besoin de l'argument `source`) :

```bash
dagger call build-env
```

Lancer la fonction Build :
```bash
dagger call build --source=.
```

Prendre le temps d'analyser les différences avant/après modifications dans les traces sur le Dagger cloud.

Pour la suite, vous allez utiliser votre module dans une GitHub Action - [cliquer ici](04-utiliser-module-github-actions.md).
