
PROTOC_VERSION := 3.9.1

ARCH := $(shell uname -m)
ifeq ($(ARCH),arm64)
	PROTOC_ARCH := aarch_64
endif
ifeq ($(ARCH),x86_64)
	PROTOC_ARCH := x86_64
endif

ifndef PROTOC_ARCH
$(error Unsupported arch: $(ARCH))
endif

PROTOC_ZIP := protoc-$(PROTOC_VERSION)-linux-x86_64.zip

BUILDX_INSTALLED := $(shell docker buildx version 2> /dev/null)
ifeq ($(BUILDX_INSTALLED),)
	BUILD_CMD := docker build --build-arg protoc_arch=$(PROTOC_ARCH)
else
	BUILD_CMD := docker buildx build --load --build-arg protoc_arch=$(PROTOC_ARCH)
endif

.PHONY: all
all: v2 v3

.PHONY: buildv2
buildv2: Dockerfile
	$(BUILD_CMD) --tag sensu/core:v2 .

.PHONY: buildv3
buildv3: Dockerfile
	$(BUILD_CMD) --tag sensu/core:v3 \
		--build-arg protoc_version=3.9.1 \
		.

.PHONY: v2
v2: buildv2
	docker run \
		-v $(CURDIR):/go/src/github.com/sensu/core \
		-w /go/src/github.com/sensu/core/v2 \
		sensu/core:v2 \
		go generate .

.PHONY: v3
v3: buildv3
	docker run \
		-v $(CURDIR):/go/src/github.com/sensu/core \
		-w /go/src/github.com/sensu/core/v3 \
		sensu/core:v3 \
		go generate .
