FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8003

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8003", \
                "-permissions-srv=user-auth-agent:8002", \
                "-config-srv=app-config-agent:8005", \
                "-apptoken-srv=app-token-agent:8006", \
                "-db-srv=mongodb://dev-datalab-user:secure@192.168.0.177:27018"]
