##
# Build Docker images used to generate ccache output testdata
##

docker-all: docker-debian docker-ubuntu

docker-debian: \
	docker-debian-11 \
	docker-debian-10 \
	docker-debian-9

docker-debian-%: Dockerfile
	docker build \
		--build-arg DOCKER_IMAGE=debian \
		--build-arg DOCKER_TAG=$*-slim \
		-t local/ccache:debian-$* \
		-f Dockerfile \
		.

docker-ubuntu: \
	docker-ubuntu-22.04 \
	docker-ubuntu-20.04 \
	docker-ubuntu-18.04

docker-ubuntu-%: Dockerfile
	docker build \
		--build-arg DOCKER_IMAGE=ubuntu \
		--build-arg DOCKER_TAG=$*\
		-t local/ccache:ubuntu-$* \
		-f Dockerfile \
		.

##
# Run software builds with ccache enabled to generate output testdata
##

ccache-testdata-all: ccache-testdata-debian ccache-testdata-ubuntu

ccache-testdata-debian: \
	ccache-testdata-debian-11 \
	ccache-testdata-debian-10 \
	ccache-testdata-debian-9

ccache-testdata-debian-%:
	./run.sh local/ccache debian-$* $@

ccache-testdata-ubuntu: \
	ccache-testdata-ubuntu-22.04 \
	ccache-testdata-ubuntu-20.04 \
	ccache-testdata-ubuntu-18.04

ccache-testdata-ubuntu-%:
	./run.sh local/ccache ubuntu-$* $@
