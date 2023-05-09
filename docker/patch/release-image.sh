#!/bin/bash
set -ex
REPO_NAME=${1}
IMAGE_NAME=ghcr.io/${OWNER}/${REPO_NAME}:$GIT_COMMIT_SHORT_SHA
RELEASE_IMAGE_NAME=ghcr.io/${OWNER}/${REPO_NAME}:${2:-dev}

sealos tag "${IMAGE_NAME}" "${RELEASE_IMAGE_NAME}"
sealos push "${RELEASE_IMAGE_NAME}"