#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"
set -e

echo "# Installing"
echo "# Deploying functions"
./functions/deploy.sh
echo "# Functions deploy done"

echo "# Installation done"
