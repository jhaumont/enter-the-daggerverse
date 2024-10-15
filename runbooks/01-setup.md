# Setup

Pour réaliser ce codelab, vous avez 2 possibilités :

- Utiliser le GitHub Codespace mis à disposition (**c'est la méthode recommandée**)
- Installer les prérequis et Dagger sur votre machine

## Codespace

[![Click to open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/jhaumont/enter-the-daggerverse)

## On prem

Voici les éléments à mettre en place pour le réaliser sur votre machine:

- Pour Windows, utiliser WSL
- Avoir Docker installer sur votre machine
  - ***Docker for Windows/Mac*** ou ***Rancher Desktop*** pour Windows et Mac OS
- Installer Git
- Installer la version `1.23.x` de GO https://go.dev/doc/install
- Installer Dagger https://docs.dagger.io/quickstart/cli/
  - Privilégier la version `0.13.5` avec laquelle le codelab a été préparé. Exemple pour linux:
    ```bash
    curl -fsSL https://dl.dagger.io/dagger/install.sh | DAGGER_VERSION=0.13.5 $HOME/.local/bin sh
    ```

> [!TIP]
> En cas de soucis avec les credentials docker sur WSL, essayer cette solution https://forums.docker.com/t/docker-credential-desktop-exe-executable-file-not-found-in-path-using-wsl2/100225/5

### Getting Started with Dagger

Une fois prêts, démarrer le codelab [ici](02-getting-started-with-Dagger.md)
