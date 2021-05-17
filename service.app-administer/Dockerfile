FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8003

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8003", \
                "-user-srv=service-user-administer:8001", \
                "-permissions-srv=service-user-authentication:8002", \
                "-config-srv=service-app-configuration:8005", \
                "-db-srv=mongodb://appadminstorage:secure@192.168.0.179:27018"]