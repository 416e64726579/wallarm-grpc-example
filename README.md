
# Path Traversal gRPC example

The very straightforward and silly example of parsing and detecting attacks in a gRPC stream by the Wallarm node.

## Usage

While in the `ptrav` directory, run the following commands:
```sh
export GO111MODULE=on
export PATH="$PATH:$(go env GOPATH)/bin"
( cd ../cmd/protoc-gen-go-grpc && go install . )
protoc \
  --go_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:. \
  --go-grpc_out=Mgrpc/service_config/service_config.proto=/internal/proto/grpc_service_config:. \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  ptrav/ptrav.proto
```

Compile the binaries
```sh
GOOS=linux go build -o grpc-client client/main.go
GOOS=linux go build -o grpc-server server/main.go
```

Bake images and convey them to the Docker Hub
```sh
docker build -t awallarm/grpc-client -f Dockerfile.client .
docker build -t awallarm/grpc-wlrm -f Dockerfile.server .
docker push awallarm/grpc-client
docker push awallarm/grpc-wlrm
```

Apply terraform state to deploy containers
```sh
terraform init
terraform apply
```

Send gRPC via the client locally (`172.17.0.3` is an IP address of Wallarm Node)
```sh
docker run -ti --rm awallarm/grpc-client grpc-client 172.17.0.3:80 ../../../../etc/passwd
docker run -ti --rm awallarm/grpc-client grpc-client 172.17.0.3:80 /etc/shadow
docker run -ti --rm awallarm/grpc-client grpc-client 172.17.0.3:80 ../../../../../../../../../../../etc/group
```

remotely
```sh
go run client/main.go 127.0.0.1:5082 ../../../../../../../../../../proc/1/environ
```