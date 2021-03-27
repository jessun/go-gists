override DOCKER         = $(shell which docker)
DOCKER_IMAGE  ?= docker-registry:5000/actiontech/universe-compiler-go1.14.13-centos6

all: docker_build upload

install:
	go build

docker_build:
	$(DOCKER) run -v $(shell pwd):/universe --rm $(DOCKER_IMAGE) -c "cd /universe && make install"
	
upload:
	scp -i ~/MEGAsync/actionsky/ssh/id_rsa ./test-go root@10.186.60.24:/tmp

