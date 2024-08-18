#!/bin/bash

# Function to perform setup activities
setup() {
  if ! command -v svu &> /dev/null; then
    printf "\nsvu not available. Installing latest svu for semantic versioning...\n"
    go install github.com/caarlos0/svu@latest
  else
    printf "\nsvu is already installed.\n"
  fi

  if ! command -v cargo &> /dev/null; then
    printf "\ncargo not available. Skipping install of git-cliff. Installing goreleaser & chglog...\n"
  elif ! command -v git-cliff &> /dev/null; then
    printf "\nInstalling git-cliff for changelog generation...\n"
    cargo install git-cliff
    git cliff --init
    printf "\nInstalling typos-cli for changelog typo detection & autofix...\n"
    cargo install typos-cli
  else
    printf "\ngit-cliff and cargo are already installed.\n"
  fi
}

# Arguments: setup, prep-release, release
if [ $# -ne 1 ]; then
  echo "Please provide the required argument: action: setup, prep-release, or release."
  exit 1
fi

action=$1

if [ "$action" == "setup" ]; then
  setup

elif [ "$action" == "release" ]; then
  printf "\nReleasing with ... ...\n"

else
  echo "Invalid action provided: $action. Please use setup or release."
  exit 1
fi
