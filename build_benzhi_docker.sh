#!/usr/bin/env sh
set -eu
docker build --platform "${TARGETPLATFORM:-linux/amd64}" -f benzhi.Dockerfile -t "${IMAGE_TAG:-ai-course-cert-benzhi}" .
