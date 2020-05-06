set -e

WORKDIR=src/kitchen

ROOT_DIR=$WORKSPACE/$WORKDIR
export GOPATH=$WORKSPACE
export GO111MODULE=on

source "${ROOT_DIR}/ci/common.sh"

build_all

run_tests
