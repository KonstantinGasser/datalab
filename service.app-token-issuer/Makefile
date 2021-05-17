VERSION = $(shell git rev-parse --short HEAD)
deploy:
	# ============ Compiling for LINUX/AMD64 ============
	GOOS=linux GOARCH=arm64 go build -o bin/service

	# ============ Docker Build: STARTING ============
	sudo docker buildx build --platform linux/arm64 -t datalab-registry.dev:5000/service.app.tokenissuer:$(VERSION)_$(tag) --load .

	# ============ Push To Registry ============
	sudo docker push datalab-registry.dev:5000/service.app.tokenissuer:$(VERSION)_$(tag)
	
	# ============ Deployment <$(VERSION)_$(tag)> DONE ============


protobuf:
	protoc -I. --go_out=$(GOPATH)/src service.app-token-issuer/proto/app-token-issuer.proto
	protoc -I. --go-grpc_out=$(GOPATH)/src service.app-token-issuer/proto/app-token-issuer.proto
