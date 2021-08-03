VERSION = $(shell git rev-parse --short HEAD)
deploy:
	# ============ Vue Build: STARTING ============
	npm run build

	# ============ Docker Build: STARTING ============
	sudo docker buildx build --platform linux/arm64 -t datalab-registry.dev:5000/service.forntend.web:$(VERSION)_$(tag) .

	# ============ Push To Registry ============
	sudo docker push datalab-registry.dev:5000/service.forntend.web:$(VERSION)_$(tag)
	# ============ Clean-Up of dist files ============
	rm -r dist
	echo '* BUILD DONE FOR\n\tcreated on: $(shell date)\n\tversion:    $(VERSION)_$(tag)\n' >> version
	# ============ Deployment <$(VERSION)_$(tag)> DONE ============

npm = $(shell which npm)
docker = $(shell which docker)

deploy-stack:
	echo '====== BUILDING: SERVICE: FRONTEND-WEB ======'
	${npm} -v
	${npm} run build
	${docker} build -t frontend.web:latest  .