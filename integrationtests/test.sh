#!/bin/bash

# List packages to be tested
export DATA=false
export DATABASE=false
export LOGGER=false
export NETWORK=false
export PUBSUB=false
export SECURITY=false
export WORKER=true

# Set environment variables for each package

go run integrationtests/main.go
