DOCKER_IMAGE = reg.actiontech.com/actiontech/universe-compiler-go$(go_version)-centos6
CTN_NAME     = universe-centos6-compiler


override DOCKER_BIN = $(shell which docker)

build:
	$(DOCKER_BIN) build --build-arg go_version=$(go_version) --network host -t $(DOCKER_IMAGE) .

push:
	$(DOCKER_BIN) push $(DOCKER_IMAGE)

run:
	$(DOCKER_BIN) run -t -d --privileged --name $(CTN_NAME) --volume $(PWD):/workbench --network host $(DOCKER_IMAGE)

stop:
	$(DOCKER_BIN) stop $(CTN_NAME) && $(DOCKER_BIN) rm $(CTN_NAME)

# b: build run
# 	$(DOCKER_BIN) logs $(CTN_NAME)
#
