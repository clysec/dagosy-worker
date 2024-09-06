#!/bin/bash
VERSION=$1
COMMENT=$2

git add . && \
git commit -m "$COMMENT" && \
git tag "$VERSION" -m "Version $VERSION" && \
git push && \
git push --tags