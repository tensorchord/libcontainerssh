FROM ubuntu:20.04

RUN apt-get update && apt-get install -y libssl-dev

COPY containerssh /

VOLUME /etc/containerssh
VOLUME /var/secrets

ENTRYPOINT ["/containerssh"]
CMD ["--config", "/etc/containerssh/config.yaml"]
