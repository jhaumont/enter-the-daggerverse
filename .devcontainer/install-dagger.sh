#!/bin/sh

mkdir -p "$HOME/.local/bin"
curl -fsSL https://dl.dagger.io/dagger/install.sh | DAGGER_VERSION=0.15.1 BIN_DIR=$HOME/.local/bin sh
sudo su <<EOF
chmod 777 /var/run/docker.sock
echo "alias ll='ls -al'" >> $HOME/.zshrc
echo "PATH=$PATH:$HOME/.local/bin" >> $HOME/.zshrc
echo "autoload -U compinit" >> $HOME/.zshrc
echo "compinit" -i >> $HOME/.zshrc
$HOME/.local/bin/dagger completion zsh > /usr/share/zsh/site-functions/_dagger
EOF
