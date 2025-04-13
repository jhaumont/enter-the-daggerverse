#!/bin/sh

mkdir -p "$HOME/.local/bin"
curl -fsSL https://dl.dagger.io/dagger/install.sh | DAGGER_VERSION=0.18.2 BIN_DIR=$HOME/.local/bin sh
sudo su <<EOF
echo "alias ll='ls -al'" >> $HOME/.zshrc
echo "PATH=$PATH:$HOME/.local/bin" >> $HOME/.zshrc
echo "autoload -U compinit" >> $HOME/.zshrc
echo "compinit -i" >> $HOME/.zshrc
echo "sudo chmod a+rwx /var/run/docker.sock" >> $HOME/.zshrc
$HOME/.local/bin/dagger completion zsh > /usr/share/zsh/site-functions/_dagger
echo "[alias]" >> $HOME/.gitconfig
echo "co = checkout" >> $HOME/.gitconfig
echo "br = branch" >> $HOME/.gitconfig
echo "ci = commit" >> $HOME/.gitconfig
echo "st = status" >> $HOME/.gitconfig
echo 'l = log --graph --pretty=format:\"%C(yellow)%h%Creset%C(red)%d%Creset %C(cyan)%cr%Creset %s %Creset%C(green)%an\"' >> $HOME/.gitconfig
echo "lf = log --pretty=fuller" >> $HOME/.gitconfig
echo "unadd = reset HEAD^" >> $HOME/.gitconfig
echo "graph = log --graph --abbrev-commit --decorate --format=format:'%C(yellow)%h%C(reset)%C(auto)%d%C(reset) %C(normal)%s%C(reset) %C(dim magenta)%an%C(reset) %C(dim blue)(%ar) %C(reset)' --all" >> $HOME/.gitconfig
EOF
