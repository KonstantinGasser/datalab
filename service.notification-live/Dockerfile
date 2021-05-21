FROM scratch

WORKDIR /usr/bin

COPY ./bin/service .

EXPOSE 8008

ENTRYPOINT ["/usr/bin/service", \ 
                "-host=:8008", \
                "-token-srv=service-user-authentication:8002"]
