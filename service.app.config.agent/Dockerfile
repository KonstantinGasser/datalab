FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8005

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8005", \
                "-db-srv=mongodb://dev-datalab-user:secure@192.168.0.177:27018"]
