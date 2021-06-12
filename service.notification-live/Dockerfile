FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8008

ENTRYPOINT ["/usr/bin/service", \ 
                "-host=:8008", \
                "-token-srv=service-user-auth-agent:8002", \
                "-db-srv=mongodb://notify-live-agent:secure@notify-live-agent-db:27022"]
