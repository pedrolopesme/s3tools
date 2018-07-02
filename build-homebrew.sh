#!/bin/bash

#
# build-homebrew.sh
#
# Shell script used during Home Brew installation. Please, do not change it.
#

set -eux

export GOPATH="$(pwd)/.gobuild"
SRCDIR="${GOPATH}/src/github.com/pedrolopesme/s3tools"

[ -d ${GOPATH} ] && rm -rf ${GOPATH}
mkdir -p ${GOPATH}/{src,pkg,bin}
mkdir -p ${SRCDIR}
cp -R * ${SRCDIR}
(
    echo ${GOPATH}
    cd ${SRCDIR}
    go get .
    go install .
)
