# e.g. REDIS_VERSION=5.0.7
REDIS_VERSION=
REDIS_SRC_CODE_URL=http://download.redis.io/releases/redis-${REDIS_VERSION}.tar.gz
REDIS_SRC_VERSIONS_URL=http://download.redis.io/releases/
SRC_DIR=/usr/local/src/redis-src

list_redis_version:
	curl ${REDIS_SRC_VERSIONS_URL} | grep "redis"  | awk -F"redis." '{print $$2}' | awk -F".tar.gz" '{print $$1}'
	
