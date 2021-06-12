FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8002

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8002", \
                "-db-srv=mongodb://user-auth-agent:secure@user-auth-agent-db:27017"]