FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8006

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8006", \
                "-db-srv=mongodb://dev-datalab-user:secure@192.168.0.177:27018"]
