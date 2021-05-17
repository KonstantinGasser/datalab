FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8002

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8002", \
                "-db-srv=mongodb://userauthstorage:secure@192.168.0.179:27021"]