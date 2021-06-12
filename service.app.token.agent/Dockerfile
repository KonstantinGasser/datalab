FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8006

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8006", \
                "-db-srv=mongodb://app-token-agent:secure@app-token-agent-db:27020"]