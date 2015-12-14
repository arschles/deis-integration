SHORT_NAME := deis-integration

export GO15VENDOREXPERIMENT=1

# dockerized development environment variables
REPO_PATH := github.com/arschles/${SHORT_NAME}
DEV_ENV_IMAGE := quay.io/deis/go-dev:0.2.0
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_PREFIX := docker run --rm -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR}
DEV_ENV_CMD := ${DEV_ENV_PREFIX} ${DEV_ENV_IMAGE}

LDFLAGS := "-s -X main.version=${VERSION}"
BINDIR := ./rootfs/bin

REGISTRY ?= ${DEV_REGISTRY}
IMAGE_PREFIX ?= arschles
VERSION ?= git-$(shell git rev-parse --short HEAD)

IMAGE := ${REGISTRY}${IMAGE_PREFIX}/${SHORT_NAME}:${VERSION}

bootstrap:
	${DEV_ENV_CMD} glide up

build:
	mkdir -p ${BINDIR}
	${DEV_ENV_PREFIX} -e CGO_ENABLED=0 ${DEV_ENV_IMAGE} go build -a -installsuffix cgo -ldflags '-s' -o $(BINDIR)/boot || exit 1

docker-build:
	# build the main image
	docker build --rm -t ${IMAGE} rootfs

docker-push: docker-build
	docker push ${IMAGE}
