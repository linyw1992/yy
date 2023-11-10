DOCKER := $(shell which docker)
DOCKER_IMAGE = common-contracts
DOCKER_TAG = 0.0.1
DOCKER_HOUSE = registry.shangchain.net:5443



docker:
	$(DOCKER) build -t ${DOCKER_IMAGE}:${DOCKER_TAG} .

upload:
	$(DOCKER) tag ${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_HOUSE}/${DOCKER_IMAGE}:${DOCKER_TAG}
	$(DOCKER) push ${DOCKER_HOUSE}/${DOCKER_IMAGE}:${DOCKER_TAG}
