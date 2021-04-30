VERSION_TAG = $(shell git rev-parse --short HEAD)
deploy:
	# ============ Vue Build: STARTING ============
	cd ./webapp && npm run build

	# ============ Docker Build: STARTING ============
	sudo docker build -t datalab-registry.dev:5000/frontend:$(VERSION_TAG) .

	# ============ Push To Registry ============
	sudo docker push datalab-registry.dev:5000/frontend:$(VERSION_TAG)

	# ============ Deployment <$(VERSION_TAG)> DONE ============
