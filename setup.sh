#!/bin/bash

parent_dir=$(pwd)
installation_dir=/tmp/bin

mkdir -p $installation_dir
cd $installation_dir || exit

go_download_url=https://go.dev/dl/$go_version

if ! command -v go &> /dev/null; then

    # shellcheck disable=SC2059
    printf "\ngo not available. installing go$go_version...\n"
    wget $go_download_url
    tar -C $installation_dir -xzf $installation_dir/$go_version
    rm $go_version

    # shellcheck disable=SC1090
    echo "export GOROOT=$installation_dir/go;export GOBIN=$installation_dir/go/bin;export PATH=$PATH:$installation_dir:$GOROOT:$GOBIN" >> ~/.bashrc

    # shellcheck disable=SC1090
    source ~/.bashrc

fi

cd "$parent_dir" || exit

# Setup githooks
printf "\nsetting up git hooks\n"
. ./.githooks/prepare-hook.sh

printf "\nsetting up git commit message template\n"
git config commit.template ./.gitmessage
git config commit.cleanup strip

