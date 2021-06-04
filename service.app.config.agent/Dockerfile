FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8005

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8005", \
                "-db-srv=mongodb://appconfigstorage:secure@192.168.178.103:27019"]