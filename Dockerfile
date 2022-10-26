FROM golang AS build

WORKDIR /src
COPY . /src
RUN CGO_ENABLED=0 go build -o instup-bot main.go && \
    git clone https://github.com/openshift/cluster-version-operator && \
    cd ./cluster-version-operator && go build ./hack/cluster-version-util && cd .. && \
    mv ./cluster-version-operator/cluster-version-util ./cluster-version-util/ && \
    curl -# https://mirror.openshift.com/pub/openshift-v4/clients/ocp/stable/openshift-client-linux.tar.gz | tar xz oc && \
    mv oc ./cluster-version-util/ && \
    chmod +x ./cluster-version-util/cluster-version-util.sh ./cluster-version-util/oc ./cluster-version-util/cluster-version-util

FROM ubi8-minimal
RUN microdnf install -y graphviz
WORKDIR /
COPY --from=build /src/instup-bot /src/config.toml /src/pull-secret /
COPY --from=build /src/cluster-version-util/oc /src/cluster-version-util/cluster-version-util /src/cluster-version-util/cluster-version-util.sh /cluster-version-util/
ENTRYPOINT ["/instup-bot"]
