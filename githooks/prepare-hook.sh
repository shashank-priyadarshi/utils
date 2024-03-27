#!/bin/bash

parent_dir=$(pwd)
installation_dir=/tmp/bin

mkdir -p $installation_dir
cd $installation_dir || exit

go_version=go1.22.1.linux-amd64.tar.gz
nodejs_version=v20.12.0

go_download_url=https://go.dev/dl/$go_version
gopls_download_url=golang.org/x/tools/gopls@latest
goimports_download_url=golang.org/x/tools/cmd/goimports@latest
godef_download_url=github.com/rogpeppe/godef@latest
golint_download_url=golang.org/x/lint/golint@latest
staticcheck_download_url=honnef.co/go/tools/cmd/staticcheck@v0.4.7
golangci_lint_download_url=github.com/golangci/golangci-lint/cmd/golangci-lint@latest
nvm_installation_script_url=https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh

if ! command -v go &> /dev/null; then

    wget $go_download_url
    tar -C $installation_dir -xzf $installation_dir/$go_version
    rm $go_version

    # shellcheck disable=SC1090
    echo "export GOROOT=$installation_dir/go;export GOBIN=$installation_dir/go/bin;export PATH=$PATH:$installation_dir:$GOROOT:$GOBIN" >> ~/.bashrc

    # shellcheck disable=SC1090
    source ~/.bashrc

fi

if ! command -v gopls &> /dev/null; then
    go install $gopls_download_url
fi

if ! command -v goimports &> /dev/null; then
    go install $goimports_download_url
fi

if ! command -v godef &> /dev/null; then
    go install $godef_download_url
fi

if ! command -v golint &> /dev/null; then
    go install $golint_download_url
fi

if ! command -v staticcheck &> /dev/null; then
    go install $staticcheck_download_url
fi

if ! command -v golangci-lint &> /dev/null; then
    go install $golangci_lint_download_url
fi

if ! command -v npm &> /dev/null; then

    nodejs_installation_dir=$installation_dir/nodejs

    if ! command -v node &> /dev/null; then

#       TODO: Fork this script, save nvm to /tmp/bin and all other dependencies to this directory, then install node and npm to this directory as well
        curl -o- "$nvm_installation_script_url" | bash

        # shellcheck disable=SC1090
        source ~/.bashrc

        nvm install "$nodejs_version"

#        curl -o "$nodejs_installation_dir/$nodejs_version" $nodejs_download_url
#
#        mkdir -p $nodejs_installation_dir
#        tar -xf "$nodejs_installation_dir/$nodejs_version" -C $nodejs_installation_dir --strip-components=1

    fi

    NODE_HOME=$nodejs_installation_dir
    PATH="$NODE_HOME/bin:$PATH"

fi

if ! command -v commitlint &> /dev/null; then
  npm install --global commitlint@latest @commitlint/cli@latest @commitlint/config-conventional
fi

cd "$parent_dir" || exit

cp "$parent_dir/githooks/commit-msg" "$parent_dir/.git/hooks/"
chmod +x "$parent_dir/.git/hooks/commit-msg"

cp "$parent_dir/githooks/pre-commit" "$parent_dir/.git/hooks/pre-commit"
chmod +x "$parent_dir/.git/hooks/pre-commit"
