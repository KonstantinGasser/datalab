VERSION_TAG = $(shell git rev-parse --short HEAD)
deploy:
	# ============ Vue Build: STARTING ============
	npm run build

	# ============ Docker Build: STARTING ============
	sudo docker buildx build --platform linux/arm64 -t datalab-registry.dev:5000/frontend:$(VERSION_TAG) --load .

	# ============ Push To Registry ============
	sudo docker push datalab-registry.dev:5000/frontend:$(VERSION_TAG)

	# ============ Deployment <$(VERSION_TAG)> DONE ============