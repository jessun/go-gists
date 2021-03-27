override DOCKER         = $(shell which docker)
DOCKER_IMAGE  ?= docker-registry:5000/actiontech/universe-compiler-go1.14.13-centos6

install:
	go build

docker_build:
	$(DOCKER) run -v $(shell pwd):/universe --rm $(DOCKER_IMAGE) -c "cd /universe && make install && ./test-go"
