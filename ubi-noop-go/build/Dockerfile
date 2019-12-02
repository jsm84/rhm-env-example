FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

LABEL name="UBI NoOp Operator" \
      vendor="Red Hat" \
      version="v0.0.1" \
      release="1" \
      summary="An example operator that exists purely for educational purposes" \
      description="An operator that deploys a desired number of sleeping Red Hat Universal Base Image (UBI) pods"

ENV OPERATOR=/usr/local/bin/ubi-noop-go \
    USER_UID=1001 \
    USER_NAME=ubi-noop-go

COPY licenses/ /licenses

# install operator binary
COPY build/_output/bin/ubi-noop-go ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
