FROM scratch

WORKDIR /usr/bin

COPY build/service .

EXPOSE 8080

ENTRYPOINT ["/usr/bin/service", \ 
                "-host=:8080", \ 
                "-user-srv=service-user-meta-agent:8001", \  
                "-app-srv=service-app-meta-agent:8003", \
                "-apptoken-srv=service-app-token-agent:8006", \
                "-config-srv=service-app-config-agent:8005", \
                "-token-srv=service-user-auth-agent:8002", \
                "-notify-srv=service-notify-live-agent:8008"]