MAKEFLAGS=--no-builtin-rules --no-builtin-variables --always-make
ROOT := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export PATH := $(ROOT)/scripts:$(PATH)

gen_server:
	scripts/gen_server.sh

encrypt_dev:
	scripts/encrypt.sh

decrypt_dev:
	scripts/decrypt.sh

check:
	scripts/check.sh