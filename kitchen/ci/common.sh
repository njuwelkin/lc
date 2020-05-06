#!/bin/bash

ERROR() {
  echo "[ERROR] $1"
}

INFO () {
  echo "[INFO] $1"
}

build_all() {
  go clean -cache -testcache  ./...

  if [[ -z $HTTPS_PROXY ]]; then
    https_proxy=${HTTPS_PROXY} go get ./...
  else
    go get ./...
  fi

  go install ./...
}

run_tests() {
  # build and run tests with coverage
  INFO "Running tests with coverage."
    # For PR, go here
  go test -coverprofile=coverage.out ./...
}


