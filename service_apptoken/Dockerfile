FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8005

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8006", \
                "-db-srv=mongodb://configstorage:secure@192.168.0.177:27019"]