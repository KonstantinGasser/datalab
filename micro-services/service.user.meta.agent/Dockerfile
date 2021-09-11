FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8002

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8001", \
                "-db-srv=mongodb://mongo-datalab-user:secure@dev-mongo-central:27017"]
