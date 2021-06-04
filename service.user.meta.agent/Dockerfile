FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8002

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8001", \
                "-db-srv=mongodb://useradminstorage:secure@192.168.0.177:27017"]