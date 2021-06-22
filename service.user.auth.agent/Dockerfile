FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8002

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8002", \
                "-db-srv=mongodb://dev-datalab-user:secure@192.168.0.177:27018"]
