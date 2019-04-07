#!/bin/sh

source ./hack/common.sh

docker build -t "$IMAGE" .
