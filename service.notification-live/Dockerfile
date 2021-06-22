FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8008

ENTRYPOINT ["/usr/bin/service", \ 
                "-host=:8008", \
                "-token-srv=user-auth-agent:8002", \
                "-db-srv=mongodb://dev-datalab-user:secure@192.168.0.177:27018"]
