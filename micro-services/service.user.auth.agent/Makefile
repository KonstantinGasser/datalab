VERSION = $(shell git rev-parse --short HEAD)
deploy:
	# ============ Compiling for LINUX/AMD64 ============
	GOOS=linux GOARCH=arm64 go build -o build/service cmd/main.go

	# ============ Docker Build: STARTING ============
	sudo docker buildx build --platform linux/arm64 -t datalab-registry.dev:5000/service.user.auth.agent:$(VERSION)_$(tag) --load .

	# ============ Push To Registry ============
	sudo docker push datalab-registry.dev:5000/service.user.auth.agent:$(VERSION)_$(tag)
	echo '* BUILD DONE FOR\n\tcreated on: $(shell date)\n\tversion:    $(VERSION)_$(tag)\n' >> build/version
	# ============ Deployment <$(VERSION)_$(tag)> DONE ============


go = $(shell which go)
docker = $(shell which docker)

deploy-stack:
	echo '====== BUILDING: SERVICE: USER-AUTH-AGENT ======'
	${go} version
	GOOS=linux GOARCH=arm64 ${go} build -o build/service cmd/main.go
	${docker} build -t user.auth.agent:latest  .