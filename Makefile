build:
	DOCKER_BUILDKIT=0 docker build --no-cache -f docker/server/Dockerfile -t server .
	DOCKER_BUILDKIT=0 docker build --no-cache -f docker/client/Dockerfile -t client .
