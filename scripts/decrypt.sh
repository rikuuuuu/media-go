#!/usr/bin/env bash

APP_ROOT=$(dirname $0)/..

gcloud --quiet config set project media-dev-303813

gcloud kms decrypt \
    --location=asia-northeast1 \
    --keyring=mediadev \
    --key=config \
    --plaintext-file=${APP_ROOT}/config/dev-secret.json \
    --ciphertext-file=${APP_ROOT}/config/dev-secret.json.enc

# sops --decrypt ${APP_ROOT}/config/dev-secret.enc.txt > ${APP_ROOT}/config/dev-secret.raw.txt
