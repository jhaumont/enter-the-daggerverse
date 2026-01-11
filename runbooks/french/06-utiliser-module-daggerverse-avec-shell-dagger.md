# Utiliser un module publié sur Daggerverse avec le shell Dagger

## Installer un module - Git info

Pour installer un module Dagger, vous devez initialisze l'environnement Dagger dans le répertoire `hello-sh` :
```bash
dagger init
```

Puis installer le module :
```bash
dagger install github.com/vbehar/daggerverse/git-info@v0.12.1
```

Vérifier que le module est bien chargé :
```bash
dagger <<< .help
```

Avec comme résultat :
```
✔ connect 0.2s
✔ load module 0.7s
✔ serving dependency modules 0.0s

  container-echo   Returns a container that echoes whatever string argument is provided
  grep-dir         Returns lines that match a pattern in the files of the provided Directory

  git-info      A Dagger Module to extract information about a git reference.
  ...
```

La dernière ligne montre bien que le module `git-info` est chargé.

A présent, utilisons le module pour récupérer la référence git actuelle :
```bash
dagger <<< "git-info ../ | ref"
```

Avec comme résultat :
```
✔ connect 0.2s
✔ load module 0.7s
...
✔ .ref: String! 0.0s

HEAD
```

Ça fonctionne 🎉

## Module Dagger

Pour aller plus loin, vous pouvez [créer, publier un module dans le daggerverse et l'utiliser](07-créer-publier-utiliser-module.md) dans un pipeline.
