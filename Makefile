
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

.PHONY: buildv2
buildv2: Dockerfile
	docker build --tag sensu/core:v2 .


.PHONY: buildv3
buildv3: Dockerfile
	docker build --tag sensu/core:v3 \
		--build-arg protoc_release=https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-linux-x86_64.zip \
		.
