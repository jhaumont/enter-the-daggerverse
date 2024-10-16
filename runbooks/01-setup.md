# Setup

Pour réaliser ce codelab, vous avez 2 possibilités :

- Utiliser le GitHub Codespace mis à disposition (**méthode recommandée**)
- Installer les prérequis et Dagger sur votre machine (voir section *Installation locale* de ce document)

## Codespace

Cliquez sur le bouton ci-dessous et laissez vous porter par la magie :

[![Click to open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/jhaumont/enter-the-daggerverse)

## Installation locale

Voici les éléments à mettre en place pour réaliser le codelab sur votre machine :

- Avec Windows l'utilisation de WSL2 est recommandée (vous pouvez toutefois utiliser un système de virtualisation comme [Virtualbox](https://www.virtualbox.org))
- Avoir Docker installé sur votre machine (via **Docker for Windows/Mac** ou **Rancher Desktop** pour Windows et Mac OS)
- Installer un client [Git](https://git-scm.com/)
- Installer la version `1.23.x` du langage [Go](https://go.dev/doc/install)
- Installer le client [Dagger](https://docs.dagger.io/quickstart/cli/) version `0.13.5`
  - Exemple pour linux:
    ```bash
    curl -fsSL https://dl.dagger.io/dagger/install.sh | DAGGER_VERSION=0.13.5 $HOME/.local/bin sh
    ```

> [!TIP]
> En cas de soucis avec les identifiants docker sur WSL2, essayer la solution suivante https://forums.docker.com/t/docker-credential-desktop-exe-executable-file-not-found-in-path-using-wsl2/100225/5

### Getting Started with Dagger

Une fois prêts, démarrer le allez à la page [Getting started with Dagger](02-getting-started-with-Dagger.md)
