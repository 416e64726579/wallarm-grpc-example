IMAGE_REPO ?= docker.io/awallarm

build:
	@echo "Building the docker image..."
	@GOOS=linux go build -o $(CURDIR)/ptrav/grpc-client $(CURDIR)/ptrav/client/main.go
	@GOOS=linux go build -o $(CURDIR)/ptrav/grpc-server $(CURDIR)/ptrav/server/main.go
	@docker build -t $(IMAGE_REPO)/grpc-wlrm:latest -f $(CURDIR)/ptrav/Dockerfile.server $(CURDIR)/ptrav
	@docker build -t $(IMAGE_REPO)/grpc-client:latest -f $(CURDIR)/ptrav/Dockerfile.client $(CURDIR)/ptrav

push: build
	@echo "Pushing the docker image..."
	@docker push $(IMAGE_REPO)/grpc-wlrm:latest
	@docker push $(IMAGE_REPO)/grpc-client:latest