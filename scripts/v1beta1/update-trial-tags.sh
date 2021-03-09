#!/bin/bash

# Copyright 2021 The Kubeflow Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script is used to update tags in the Trial templates for all Katib examples.
# Run ./scripts/v1beta1/update-trial-template-tags.sh -t <image-tag> to execute it.

usage() {
  echo "Usage: $0 [-t <TAG>]" 1>&2
  exit 1
}

while getopts ":t:" opt; do
  case $opt in
  t)
    TAG=${OPTARG}
    ;;
  *)
    usage
    ;;
  esac
done

if [[ -z "$TAG" ]]; then
  echo "Image registry and tag must be set"
  echo "Usage: $0 [-t <TAG>]" 1>&2
  exit 1
fi

echo "Updating Trial template images..."
echo "Image tag: ${TAG}"

IMAGE_MXNET_MNIST="docker.io/kubeflowkatib\/mxnet-mnist"
IMAGE_PYTORCH_MNIST="docker.io/kubeflowkatib/pytorch-mnist"
IMAGE_ENAS_GPU="docker.io/kubeflowkatib/enas-cnn-cifar10-gpu"
IMAGE_ENAS_CPU="docker.io/kubeflowkatib/enas-cnn-cifar10-cpu"
IMAGE_DARTS="docker.io/kubeflowkatib/darts-cnn-cifar10"

# MXNet mnist.
find ./examples -regex ".*\.yaml" -exec sed -i '' -e "s@${IMAGE_MXNET_MNIST}:.*@${IMAGE_MXNET_MNIST}:${TAG}@" {} \;

# PyTorch mnist.
find ./examples -regex ".*\.yaml" -exec sed -i '' -e "s@${IMAGE_PYTORCH_MNIST}:.*@${IMAGE_PYTORCH_MNIST}:${TAG}@" {} \;

# ENAS cifar-10 GPU.
find ./examples -regex ".*\.yaml" -exec sed -i '' -e "s@${IMAGE_ENAS_GPU}:.*@${IMAGE_ENAS_GPU}:${TAG}@" {} \;

# ENAS cifar-10 CPU.
find ./examples -regex ".*\.yaml" -exec sed -i '' -e "s@${IMAGE_ENAS_CPU}:.*@${IMAGE_ENAS_CPU}:${TAG}@" {} \;

# DARTS cifar-10.
find ./examples -regex ".*\.yaml" -exec sed -i '' -e "s@${IMAGE_DARTS}:.*@${IMAGE_DARTS}:${TAG}@" {} \;

echo "Trial template images has been updated"
