DATE = $(shell date +"%Y%m%y")
DOCKER_IMAGE = reg.actiontech.com/actiontech/universe-unittest-with-ucore:$(DATE)

build_docker_image:
	docker build -t $(DOCKER_IMAGE) .

push_docker_image:
	echo xi3mdE1NZEfGaT3NgnbR7dQ1Kkwlw0bW | docker login reg.actiontech.com --username 'robot$$actiontech-ci-agent' --password-stdin
	docker push $(DOCKER_IMAGE)

default: build_docker_image

