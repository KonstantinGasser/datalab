FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8001

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8001", \
                "-db-srv=mongodb://useradminstorage:secure@192.168.0.179:27017"]   