# Introduction à Dagger Shell

Le 26 Mars 2025, l'équipe de Dagger.io [a annoncé](https://dagger.io/blog/a-shell-for-the-container-age-introducing-dagger-shell) une nouvelle fonctionnalité : un shell interactif pour Dagger.

Cette nouvelle fonctionnalité est très pratique pour tester Dagger.

Le terme "shell" peut prêter à confusion. Il ne s'agit pas d'un shell comme Bash, Zsh, mais plus d'un langage très simple d'interaction avec Dagger, similaire à la syntaxe [jq](https://jqlang.org).

## Langement du shell

Ouvrez le répertoire `hello-sh` :
```bash
cd /workspaces/enter-the-daggerverse/hello-sh
```

Pour démarrer le shell Dagger, exécutez simplement la commande :
```bash
dagger
```

Avec pour résultat :
```
Dagger interactive shell. Type ".help" for more information. Press Ctrl+D to exit.
⋈
─ esc nav mode · > run prompt ────────────────────────────────────────────────────
```

Vous pouvez quitter le shell avec `Ctrl+C`, `Ctrl+D` ou `exit`.

# Obtenir de l'aide

Lorsque vous êtes dans le shell Dagger, vous pouvez utiliser la commande `.help` pour afficher de l'aide :
```
✔ .help 0.0s
  cache-volume   Constructs a cache volume for a given cache key.
  container      Creates a scratch container.
  directory      Creates an empty directory.
  ...
```

Vous pouvez avoir aussi des informations sur un module ou une commande. Reportez vous à la documentation officielle : [Built-in help](https://docs.dagger.io/features/shell/#built-in-help).

# Exécuter votre première commande

Une fois le shell Dagger lancé, vous pouvez entrer des commandes Dagger :
```
container | from alpine | with-exec whoami
```

Vous venez de demander à Dagger d'exécuter un container basé sur l'image Docker Alpine et dans ce container, d'exécuter la commande `whoami`. Oui, mais on obtient ce résultat :
```
✔ container | from alpine | with-exec whoami 0.0s
Container@xxh3:83a64a1566b1bdba
```

Où est le résultat de la commande `whoami` ?

Ce comportement est normal. Il faut explicitement demander à Dagger de récupérer le résultat (sortie standard) de la commande et de l'afficher. Pour ce faire :
```
container | from alpine | with-exec whoami | stdout
```

En ajoutant  `| stdout`, Dagger affiche la sortie standard de la commande.

# Exécuter le shell Dagger dans... un shell 😮

Vous pouvez créer un script pour le shell Dagger, comme la commande `cat` :
```bash
dagger <<EOF
container \
| from alpine \
| with-exec -- sh -c "echo hello" \
| stdout
EOF
```

Le résultat sera:
```
✔ connect 0.2s
✔ loading type definitions 0.2s

✔ container: Container! 0.0s
$ .from(address: "alpine"): Container! 0.6s CACHED
$ .withExec(args: ["sh", "-c", "echo hello"]): Container! 0.0s CACHED
✔ .stdout: String! 0.0s

hello
```

Ou alors, vous pouvez utiliser la commande `echo` :
```bash
echo 'container | from alpine | with-exec -- sh -c "echo hello" | stdout' | dagger
```

Et même écrire les commandes Dagger dans un fichier :  
```bash
cat > hello.dsh <<EOF
container \
| from alpine \
| with-exec -- sh -c "echo hello" \
| stdout
EOF

dagger hello.dsh
```

# Monter un dossier ou un fichier dans le container exécuté par Dagger

Nous souhaitons créer un fichier dans le répertoire courant du container :
```bash
dagger <<EOF
container \
| from alpine \
| with-exec -- sh -c "echo titi > /data/test" \
| with-directory /data/ .
EOF
```

L'ennuie, c'est que vous allez obtenir ce message :
```
✔ connect 0.2s
✔ detect module: . 0.1s
✔ loading type definitions 0.3s
✘ with-directory 0.2s ERROR
✘ withExec sh -c 'echo titi > /data/test' 0.2s ERROR
sh: can't create /data/test: nonexistent directory                                                                                                                                                                                                                          
! exit code: 1

$ container: Container! 0.0s CACHED
$ .from(address: "alpine"): Container! 0.0s CACHED
✘ withExec sh -c 'echo titi > /data/test' 0.2s ERROR
! exit code: 1
✔ .withDirectory(
  ┆ path: "/data/"
  ┆ source: Address.directory: Directory!
  ): Container! 0.0s
```

Oui, car l'ordre des paramètres est important. C'est un peu comme enchainer des commandes dans un shell. Intervertissez les lignes `with-directory` et `with-exec`:
```bash
dagger <<EOF
container \
| from alpine \
| with-directory /data/ . \
| with-exec -- sh -c "echo titi > /data/test"
EOF
```

Si vous allez dans le répertoire `/workspaces/enter-the-daggerverse/hello-sh` aucun fichier `test` n'est présent.
Lorsque vous monter un répertoire dans un container, il est en quelque sorte copié dans le container.
Il faut donc exporter le fichier avec `export` pour le récupérer dans le répertoire `hello-sh` :
```bash
dagger <<EOF
container \
| from alpine \
| with-directory /data/ . \
| with-exec -- sh -c "echo titi > /data/test" \
| export ./test
EOF
```

Maintenant que nous avons testé le shell Dagger, nous allons utiliser un module : [Utiliser un module publié sur Daggerverse avec le shell Dagger](./06-utiliser-module-daggerverse-avec-shell-dagger.md).
