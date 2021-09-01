# Constants
REDIS_SRC_CODE_URL=http://download.redis.io/releases/redis-${REDIS_VERSION}.tar.gz
REDIS_SRC_VERSIONS_URL=http://download.redis.io/releases/
DOCKER_COMPILER_EL7 = docker-registry:5000/actiontech/compiler-centos7-gcc-9.4
DOCKER_COMPILER_EL8 = 
REDIS_SRC_DIR=/usr/local/src/redis-src

# Variables
# e.g. REDIS_VERSION=5.0.7
REDIS_VERSION ?= 6.0.15
OS_VERSION ?= el7 
DOCKER_COMPILER ?= docker-registry:5000/actiontech/compiler-centos7
REDIS_TARBALL = redis-${REDIS_VERSION}.tar.gz

ifeq (${OS_VERSION},el7)
	DOCKER_COMPILER := ${DOCKER_COMPILER_EL7}
else ifeq (${OS_VERSION},el8) 
	DOCKER_COMPILER := ${DOCKER_COMPILER_EL8}
endif

list_redis_version:
	curl ${REDIS_SRC_VERSIONS_URL} | grep "redis"  | awk -F"redis." '{print $$2}' | awk -F".tar.gz" '{print $$1}'
	
download_redis_src_code:
	 curl -C - -o ./src/${REDIS_TARBALL} ${REDIS_SRC_CODE_URL}

docker_build_redis: download_redis_src_code
	docker run -v ${shell pwd}/src:/tmp/redis-tarball ${DOCKER_COMPILER} /bin/sh -c "mkdir -p ${REDIS_SRC_DIR} \
		&& tar -xf /tmp/redis-tarball/redis-${REDIS_VERSION}.tar.gz  -C ${REDIS_SRC_DIR} --strip-components 1 \
		&& cd ${REDIS_SRC_DIR}/deps \
			&& make -j4  hiredis lua jemalloc linenoise \
		&& cd .. \
			&& make -j4 \
			&& make PREFIX=/opt/redis-${REDIS_VERSION} install \
		&& cd /opt \
			&& tar -zcf /tmp/redis-tarball/redis-${REDIS_VERSION}-glibc2.28-el8.tar.gz redis-${REDIS_VERSION}/"
