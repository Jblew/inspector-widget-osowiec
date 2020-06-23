#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"
set -e

echo "# Installing"
echo "# Deploying function"
./function/deploy.sh
echo "# Function deploy done"

echo "# Installation done"
