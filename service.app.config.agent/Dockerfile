FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8005

ENTRYPOINT ["/usr/bin/service", \
                "-host=:8005", \
                "-db-srv=mongodb://app-config-agent:secure@app-config-agent-db:27019"]