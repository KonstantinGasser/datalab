FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8003

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8003", \
                "-permissions-srv=service-user-auth-agent:8002", \
                "-config-srv=service-app-config-agent:8005", \
                "-apptoken-srv=service-app-token-agent:8006", \
                "-db-srv=mongodb://app-meta-agent:secure@app-meta-agent-db:27021"]