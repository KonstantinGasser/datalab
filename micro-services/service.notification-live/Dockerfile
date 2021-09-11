FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8008

ENTRYPOINT ["/usr/bin/service", \ 
                "-host=:8008", \
                "-token-srv=user-auth-agent:8002", \
                "-db-srv=mongodb://mongo-datalab-user:secure@dev-mongo-central:27017"]
