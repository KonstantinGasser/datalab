FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8004

ENTRYPOINT ["/usr/bin/service", \ 
                "-host=:8004", \ 
                "-apptoken-srv=app-token-agent:8006", \
                "-config-srv=app-config-agent:8005"]
