FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8080

ENTRYPOINT ["/usr/bin/service", \ 
                "-host=:8080", \ 
                "-user-srv=user-meta-agent:8001", \  
                "-app-srv=app-meta-agent:8003", \
                "-apptoken-srv=app-token-agent:8006", \
                "-config-srv=app-config-agent:8005", \
                "-auth-srv=user-auth-agent:8002", \
                "-notify-srv=notify-live-agent:8008"]
