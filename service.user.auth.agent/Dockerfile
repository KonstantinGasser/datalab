FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8002

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8002", \
                "-db-srv=mongodb://userauthstorage:secure@192.168.178.103:27021"]