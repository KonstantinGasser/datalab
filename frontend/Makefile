VERSION = $(shell git rev-parse --short HEAD)
deploy:
	# ============ Vue Build: STARTING ============
	cd ./webapp && npm run build

	# ============ Docker Build: STARTING ============
	sudo docker buildx build --platform linux/arm64 -t datalab-registry.dev:5000/frontend:$(VERSION)_$(TAG) .

	# ============ Push To Registry ============
	sudo docker push datalab-registry.dev:5000/frontend:$(VERSION)_$(TAG)
	
	# ============ Deployment <$(VERSION)_$(TAG)> DONE ============
