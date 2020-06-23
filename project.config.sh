#!/usr/bin/env bash
PROJECT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export GCP_PROJECT_ID="inspector-widget"
export GCP_FUNCTIONS_REGION="europe-west3"
export FIRESTORE_COLLECTION="osowieclog"
export API_SECRET_KEYS=""

source "${PROJECT_DIR}/secret.config.sh"
