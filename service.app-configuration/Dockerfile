FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8005

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8005", \
                "-db-srv=mongodb://appconfigstorage:secure@192.168.0.179:27019"]