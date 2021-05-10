FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8003

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8003", \
                "-user-srv=userservice:8001", \
                "-token-srv=tokenservice:8002", \
                "-db-srv=mongodb://appstorage:secure@192.168.0.177:27018"]