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
DOCKER_COMPILER ?= ${DOCKER_COMPILER_EL7}
REDIS_TARBALL = redis-${REDIS_VERSION}.tar.gz
OS_DISTRO ?= 

ifeq (${OS_VERSION},el7)
	DOCKER_COMPILER := ${DOCKER_COMPILER_EL7}
	OS_DISTRO = centos7
else ifeq (${OS_VERSION},el8) 
	DOCKER_COMPILER := ${DOCKER_COMPILER_EL8}
	OS_DISTRO = centos8
endif

GOARCH ?= x86_64
RPMBUILD_TARGET = x86_64
ifeq (${GOARCH},arm64)
    PMBUILD_TARGET = aarch64
endif

list_redis_version:
	curl ${REDIS_SRC_VERSIONS_URL} | grep "redis"  | awk -F"redis." '{print $$2}' | awk -F".tar.gz" '{print $$1}'
	
download_redis_src_code:
	docker run -v ${shell pwd}:/workbench ${DOCKER_COMPILER} /bin/sh -c "cd /tmp \
		&& curl -o ./${REDIS_TARBALL} ${REDIS_SRC_CODE_URL} \
		&& mv ${REDIS_TARBALL} /workbench/.  "

docker_build_redis: download_redis_src_code
	docker run -v ${shell pwd}:/workbench ${DOCKER_COMPILER} /bin/sh -c "mkdir -p ${REDIS_SRC_DIR} \
		&& tar -xf /workbench/${REDIS_TARBALL} -C ${REDIS_SRC_DIR} --strip-components 1 \
		&& cd ${REDIS_SRC_DIR}/deps \
			&& make -j 12 hiredis lua jemalloc linenoise \
			&& rm -fr /workbench/${REDIS_TARBALL} \
		&& cd .. \
			&& make -j 12 \
  		 	&& make PREFIX=/opt/redis-${REDIS_VERSION} install \
		&& cd /opt \
			&& cd redis-${REDIS_VERSION} && mv bin redis-${REDIS_VERSION} \
			&& tar -zcf /tmp/redis_bin.tar.gz  redis-${REDIS_VERSION} \
			&& mv /tmp/redis_bin.tar.gz /workbench/redis-${REDIS_VERSION}-${OS_DISTRO}.tar.gz "
