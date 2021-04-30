VERSION_TAG = $(shell git rev-parse --short HEAD)
deploy:
	# ============ Compiling for LINUX/AMD64 ============
	GOOS=linux GOARCH=arm64 go build -o bin/service

	# ============ Docker Build: STARTING ============
	sudo docker build -t datalab-registry.dev:5000/userservice-$(VERSION_TAG) .
	
	# ============ Push To Registry ============
	sudo docker push datalab-registry.dev:5000/userservice-$(VERSION_TAG)

	# ============ Deployment <$(VERSION_TAG)> DONE ============
