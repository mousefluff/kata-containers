# Copyright (c) 2020-2021 Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0
#

include ../../../utils.mk

ifeq ($(ARCH), ppc64le)
     override ARCH = powerpc64le
endif

.DEFAULT_GOAL := default
default: build

build:
	@RUSTFLAGS="$(EXTRA_RUSTFLAGS) --deny warnings" cargo build --target $(TRIPLE) --$(BUILD_TYPE)

static-checks-build:
	@echo "INFO: static-checks-build do nothing.."

clean:
	cargo clean

vendor:
	cargo vendor

test:
	@cargo test --all -- --nocapture

install:

check: standard_rust_check

.PHONY: \
	build \
	check \
	clean \
	install \
	test \
	vendor
