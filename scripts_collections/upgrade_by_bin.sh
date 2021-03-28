#!/bin/bash
# please run in project dir
# argument: project-name


# please set ip and user
REMOTE_SERVER_ADDR=
REMOTE_SERVER_USER=root

# environment
REMOTE=$REMOTE_SERVER_USER@$REMOTE_SERVER_ADDR
CONTAINER='udp-1 udp-2 udp-3'
SOURCE_PARENT_PATH=./bin

# check 
if [ -z $REMOTE_SERVER_ADDR ]
	then 
		echo "Please set REMOTE_SERVER_ADDR in script!"
		exit 1
fi


# exit when any command fails
set -e

cd "$PWD"

# guess project name
# way1: get project name from Makefile
MAKEFILE_PROJECT=$(grep -e "^PROJECT_NAME" Makefile|tr -d 'PROJECT_NAME'|tr -d '='|tr -d ' ')
# way2: get project name from directory name
DIR_PROJECT=$(basename "$PWD")
# If there is no PROJECT_NAME in Makefile, use directory name instead
PROJECT=${MAKEFILE_PROJECT:-$DIR_PROJECT}

if [ "$PROJECT" == "urman" ] || [ "$PROJECT" == "uguard" ]
	then
	    # check argument
		if [ -z "$1" ]
		then
	    	echo "The project has agent or manager, please use ./upgrade_by_bin.sh agent|mgr"
	    	exit 1
		elif [ "$1" == "agent" ] || [ "$1" == "mgr" ]
		then
			PROJECT=$PROJECT-$1
		else
			echo "The project has agent or manager, please use ./upgrade_by_bin.sh agent|mgr"
		    exit 1
		fi
fi

PROJECT_PATH=/opt/$PROJECT

# define script 
remote_container_script='#!/bin/bash
	function kill_and_wait {
	        pidfile=\$1
	        if [ -e \$pidfile ]; then
	                kill \$(cat \$pidfile) &>/dev/null
	        fi
	        for i in {1..10}; do
	                if [ ! -e \$pidfile ]; then
	                        return 0
	                fi
	                kill -0 \$(cat \$pidfile) &>/dev/null
	                if [ \$? -ne 0 ]; then
	                        return 0
	                fi
	                sleep 1
	        done
	        return 1
	}
	'"project_path=$PROJECT_PATH
	project=$PROJECT"'

	kill_and_wait \$project_path/\$project.pid
	if [ \$? -ne 0 ]; then
	    echo \"wait pid shutdown timeout\"
	    exit 1
	fi
	caps=\$(getcap \$project_path/bin/\$project |cut -d = -f 2)
	mv \$project_path/bin/\$project  \$project_path/bin/\$project.bak
	cp /tmp/\$project \$project_path/bin/\$project
	chown actiontech-universe:actiontech \$project_path/bin/\$project
	setcap \$caps \$project_path/bin/\$project
	cd \$project_path && nohup ./bin/\$project > logs/std.log 2>&1 & 
'

declare -a containers=$CONTAINER
remote_host_script='
	for container in '"${containers[@]}"'
	do
        docker exec $container test -d '"$PROJECT_PATH"'
        if [ $? -ne 0 ]; then
        	continue
        fi
        docker cp /tmp/'"$PROJECT"' $container:/tmp/'"$PROJECT"'
        docker exec $container bash -c '"'"'cat <<EOF > /tmp/remote_container_script.sh 
'"$remote_container_script"'
EOF'"'"'

		docker exec $container chmod +x /tmp/remote_container_script.sh
        docker exec $container /tmp/remote_container_script.sh
        if [ $? -ne 0 ]; then
        	echo "run remote_script wrong: "
	    	exit 1
        fi
        echo upgrade:$container
	done
'

# run script
case $PROJECT in 
	urman-mgr)
		SOURCE_PATH=$SOURCE_PARENT_PATH/manager
		make install/mgr
		;;
	uguard-mgr)
		SOURCE_PATH=$SOURCE_PARENT_PATH/manager
		make mgr
		;;
	urman-agent)
		SOURCE_PATH=$SOURCE_PARENT_PATH/agent
		make install/agent
		;;
	uguard-agent)
		SOURCE_PATH=$SOURCE_PARENT_PATH/agent
		make agent
		;;
	*)
		SOURCE_PATH=$SOURCE_PARENT_PATH/$PROJECT
		# urman makefile not support non-linux os for now
        make install
		;;
esac

scp $SOURCE_PATH $REMOTE:/tmp/$PROJECT
ssh $REMOTE "$remote_host_script"
if [ $? -ne 0 ]; then
	echo "run ssh wrong "
	exit 1
fi

