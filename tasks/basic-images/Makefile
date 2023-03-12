DOCKER_REPO := raggupta/alpine-devops-base
OC_TAG := supervised-v1-oc 
SIMPLE_TAG := basic-v1
TINI_TAG := tini-v1
TAG := supervised-v1

docker-login:
	docker login https://index.docker.io/v1/

# using supervised
oc-build: docker-login
	cd supervised/oc && docker build -t ${DOCKER_REPO}:${OC_TAG} .

oc-push: oc-build 
	docker push ${DOCKER_REPO}:${OC_TAG}

build: docker-login
	cd supervised/basic && docker build -t ${DOCKER_REPO}:${TAG} .

push: build 
	docker push ${DOCKER_REPO}:${TAG}


# with only bash 
simple-build: docker-login
	cd simple && docker build -t ${DOCKER_REPO}:${SIMPLE_TAG} .

simple-push: simple-build 
	docker push ${DOCKER_REPO}:${SIMPLE_TAG}

# using tini
tini-build: docker-login
	cd tini && docker build -t ${DOCKER_REPO}:${TINI_TAG} .

tini-push: tini-build 
	docker push ${DOCKER_REPO}:${TINI_TAG}