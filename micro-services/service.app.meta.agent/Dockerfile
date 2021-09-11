FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8003

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8003", \
                "-config-srv=app-config-agent:8005", \
                "-apptoken-srv=app-token-agent:8006", \
                "-db-srv=mongodb://mongo-datalab-user:secure@dev-mongo-central:27017"]
