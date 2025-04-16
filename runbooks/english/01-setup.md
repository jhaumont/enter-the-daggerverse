# Setup

To run this codelab, you have 2 choices:

- Use GitHub Codespace by clicking on the button bellow (**recommended**)
- Install all prerequisites and Dagger on your computer (see section [Local installation](#local-installation) of this document)

## Fork this repository

Please, fork this repository. With your fork, you can commit & push to save your work.

## Codespace

Click on the button bellow:

[![Click to open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/jhaumont/enter-the-daggerverse)

In the codespace's creation's page, keep in mind to **change the repository name to use your fork**. To be more comfortable, select 4-core machine type:

![](../codespace.png)

## Local installation

To run this codelab, your need install all of components bellow:

- Under Windows, WSL2 is highly recommended (but you can use an other virtualization manager like [Virtualbox](https://www.virtualbox.org))
- Install **Docker for Windows/Mac** or **Rancher Desktop**
- Install [Git](https://git-scm.com/)
- Install [Go](https://go.dev/doc/install) language `1.23.x`
- Install [Dagger](https://docs.dagger.io/quickstart/cli/) version `0.18.3`
  - Example for linux:
    ```bash
    curl -fsSL https://dl.dagger.io/dagger/install.sh | DAGGER_VERSION=0.18.3 $HOME/.local/bin sh
    ```

> [!TIP]
> If you have any issue with Docker credential by using WSL2, you can try this fix it by following instruction given on this website https://forums.docker.com/t/docker-credential-desktop-exe-executable-file-not-found-in-path-using-wsl2/100225/5

## Getting Started with Dagger

Now, you are ready to go to next page [Getting started with Dagger](02-getting-started-with-Dagger.md)
