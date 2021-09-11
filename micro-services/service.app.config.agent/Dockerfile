FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8005

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8005", \
                "-db-srv=mongodb://mongo-datalab-user:secure@dev-mongo-central:27017"]
