#/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
RELEASE=${1}
DIR=/tmp/cluster-version-util/${RELEASE}
URL=http://curl-paste.example.com:9990

function mgp() {
      curl -F 'title=Post request from a Curl command' -F 'paste=<-' ${URL}
}

if [ -d "${DIR}/release-manifests" ]; then
	cd ${SCRIPT_DIR} && \
	./cluster-version-util task-graph ${DIR} | dot -Tsvg | mgp && \
	cd ..
else
	cd ${SCRIPT_DIR} && \
	mkdir -p ${DIR} && \
	./oc image extract -a /pull-secret quay.io/openshift-release-dev/ocp-release:${RELEASE}-x86_64 --path /:${DIR} --confirm && \
	./cluster-version-util task-graph ${DIR} | dot -Tsvg | mgp && \
	cd ..
fi
