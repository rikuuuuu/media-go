#!/usr/bin/env bash

APP_ROOT=$(dirname $0)/..

gcloud kms encrypt \
    --location=asia-northeast1 \
    --keyring=mediadev \
    --key=config \
    --plaintext-file=${APP_ROOT}/config/serviceAccount.json \
    --ciphertext-file=${APP_ROOT}/config/dev-secret.json.enc